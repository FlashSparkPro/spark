package so

import (
	"context"
	"database/sql/driver"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/XSAM/otelsql"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/lightsparkdev/spark/common"
	pb "github.com/lightsparkdev/spark/proto/spark"
	"github.com/lightsparkdev/spark/so/middleware"
	"github.com/lightsparkdev/spark/so/utils"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"gopkg.in/yaml.v3"
)

var (
	defaultPoolMinConns          = 4
	defaultPoolMaxConns          = 32
	defaultPoolMaxConnLifetime   = 5 * time.Minute
	defaultPoolMaxConnIdleTime   = 30 * time.Second
	defaultPoolHealthCheckPeriod = 15 * time.Second
)

// Config is the configuration for the signing operator.
type Config struct {
	// Index is the index of the signing operator.
	Index uint64
	// Identifier is the identifier of the signing operator, which will be index + 1 in 32 bytes big endian hex string.
	// Used as shamir secret share identifier in DKG key shares.
	Identifier string
	// IdentityPrivateKey is the identity private key of the signing operator.
	IdentityPrivateKey []byte
	// SigningOperatorMap is the map of signing operators.
	SigningOperatorMap map[string]*SigningOperator
	// Threshold is the threshold for the signing operator.
	Threshold uint64
	// SignerAddress is the address of the signing operator.
	SignerAddress string
	// authzEnforced determines if authorization checks are enforced
	authzEnforced bool
	// DKGCoordinatorAddress is the address of the DKG coordinator.
	DKGCoordinatorAddress string
	// SupportedNetworks is the list of networks supported by the signing operator.
	SupportedNetworks []common.Network
	// BitcoindConfigs are the configurations for different bitcoin nodes.
	BitcoindConfigs map[string]BitcoindConfig
	// ServerCertPath is the path to the server certificate.
	ServerCertPath string
	// ServerKeyPath is the path to the server key.
	ServerKeyPath string
	// Lrc20Configs are the configurations for different LRC20 nodes and
	// token transaction withdrawal parameters.
	Lrc20Configs map[string]Lrc20Config
	// DKGLimitOverride is the override for the DKG limit.
	DKGLimitOverride uint64
	// RunDirectory is the base directory for resolving relative paths
	RunDirectory string
	// If true, return the details of the error to the client instead of just 'Internal Server Error'
	ReturnDetailedErrors bool
	// If true, return the details of the panic to the client instead of just 'Internal Server Error'
	ReturnDetailedPanicErrors bool
	// RateLimiter is the configuration for the rate limiter
	RateLimiter RateLimiterConfig
	// Tracing configuration
	Tracing common.TracingConfig
	// Database is the configuration for the database.
	Database DatabaseConfig
}

// DatabaseDriver returns the database driver based on the database path.
func (c *Config) DatabaseDriver() string {
	if strings.HasPrefix(c.Database.URI, "postgresql") {
		return "postgres"
	}
	return "sqlite3"
}

// OperatorConfig contains the configuration for a signing operator.
type OperatorConfig struct {
	// Bitcoind is a map of bitcoind configurations per network.
	Bitcoind map[string]BitcoindConfig `yaml:"bitcoind"`
	// Lrc20 is a map of addresses of lrc20 nodes per network
	Lrc20 map[string]Lrc20Config `yaml:"lrc20"`
	// Tracing is the configuration for tracing
	Tracing common.TracingConfig `yaml:"tracing"`
	// Database is the configuration for the database
	Database *DatabaseConfig `yaml:"database"`
	// ReturnDetailedErrors determines if detailed errors should be returned to the client
	ReturnDetailedErrors bool `yaml:"return_detailed_errors"`
	// ReturnDetailedPanicErrors determines if detailed panic errors should be returned to the client
	ReturnDetailedPanicErrors bool `yaml:"return_detailed_panic_errors"`
}

// BitcoindConfig is the configuration for a bitcoind node.
type BitcoindConfig struct {
	Network        string `yaml:"network"`
	Host           string `yaml:"host"`
	User           string `yaml:"rpcuser"`
	Password       string `yaml:"rpcpassword"`
	ZmqPubRawBlock string `yaml:"zmqpubrawblock"`
}

type Lrc20Config struct {
	// DisableRpcs turns off external LRC20 RPC calls for token transactions.
	// Useful to unblock token transactions in the case LRC20 nodes behave unexpectedly.
	// Although this is primarily intended for testing, even in a production environment
	// transfers can still be validated and processed without LRC20 communication,
	// although exits for resulting outputs will be blocked until the data is backfilled.
	DisableRpcs bool `yaml:"disablerpcs"`
	// DisableL1 removes the ability for clients to move tokens on L1.  All tokens minted in this Spark instance
	// must then stay within this spark instance. It disables SO chainwatching for withdrawals and disables L1 watchtower logic.
	// Note that it DOES NOT impact the need for announcing tokens on L1 before minting.
	// The intention is that if this config value is set in an SO- that any tokens minted do not have Unilateral Exit or L1 deposit capabilities.
	DisableL1                     bool   `yaml:"disablel1"`
	Network                       string `yaml:"network"`
	Host                          string `yaml:"host"`
	RelativeCertPath              string `yaml:"relativecertpath"`
	WithdrawBondSats              uint64 `yaml:"withdrawbondsats"`
	WithdrawRelativeBlockLocktime uint64 `yaml:"withdrawrelativeblocklocktime"`
	// TransactionExpiryDuration is the duration after which started token transactions expire
	// after which the tx will be cancelled and the input TTXOs will be reset to a spendable state.
	TransactionExpiryDuration time.Duration `yaml:"transaction_expiry_duration"`
	GRPCPageSize              uint64        `yaml:"grpcspagesize"`
	GRPCPoolSize              uint64        `yaml:"grpcpoolsize"`
}

type DatabaseConfig struct {
	URI                       string         `yaml:"uri"`
	IsRDS                     bool           `yaml:"is_rds"`
	PoolMinConns              *int           `yaml:"pool_min_conns"`
	PoolMaxConns              *int           `yaml:"pool_max_conns"`
	PoolMaxConnLifetime       *time.Duration `yaml:"pool_max_conn_lifetime"`
	PoolMaxConnIdleTime       *time.Duration `yaml:"pool_max_conn_idle_time"`
	PoolHealthCheckPeriod     *time.Duration `yaml:"pool_health_check_period"`
	PoolMaxConnLifetimeJitter *time.Duration `yaml:"pool_max_conn_lifetime_jitter"`
}

// RateLimiterConfig is the configuration for the rate limiter
type RateLimiterConfig struct {
	// Enabled determines if rate limiting is enabled
	Enabled bool `yaml:"enabled"`
	// Window is the time window for rate limiting
	Window time.Duration `yaml:"window"`
	// MaxRequests is the maximum number of requests allowed in the window
	MaxRequests int `yaml:"max_requests"`
	// Methods is a list of methods to rate limit
	// Note: This does not set up rate limiting across methods by IP,
	// nor does it provide configuration for custom per-method rate limiting.
	Methods []string `yaml:"methods"`
}

// NewConfig creates a new config for the signing operator.
func NewConfig(
	configFilePath string,
	index uint64,
	identityPrivateKeyFilePath string,
	operatorsFilePath string,
	threshold uint64,
	signerAddress string,
	databasePath string,
	authzEnforced bool,
	dkgCoordinatorAddress string,
	supportedNetworks []common.Network,
	aws bool,
	serverCertPath string,
	serverKeyPath string,
	dkgLimitOverride uint64,
	runDirectory string,
	rateLimiter RateLimiterConfig,
) (*Config, error) {
	identityPrivateKeyHexStringBytes, err := os.ReadFile(identityPrivateKeyFilePath)
	if err != nil {
		return nil, err
	}
	identityPrivateKeyBytes, err := hex.DecodeString(strings.TrimSpace(string(identityPrivateKeyHexStringBytes)))
	if err != nil {
		return nil, err
	}

	signingOperatorMap, err := LoadOperators(operatorsFilePath)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	var operatorConfig OperatorConfig
	if err := yaml.Unmarshal(data, &operatorConfig); err != nil {
		return nil, err
	}

	setLrc20Defaults(operatorConfig.Lrc20)

	identifier := utils.IndexToIdentifier(index)

	if dkgCoordinatorAddress == "" {
		dkgCoordinatorAddress = signingOperatorMap[identifier].Address
	}

	// If new database config is not provided, create it from the old config.
	// TODO(mhr): Deprecate DatabasePath and AWS in favor of DatabaseConfig (LPT-384).
	if operatorConfig.Database == nil {
		operatorConfig.Database = &DatabaseConfig{
			URI:   databasePath,
			IsRDS: aws,
		}
	}

	return &Config{
		Index:                     index,
		Identifier:                identifier,
		IdentityPrivateKey:        identityPrivateKeyBytes,
		SigningOperatorMap:        signingOperatorMap,
		Threshold:                 threshold,
		SignerAddress:             signerAddress,
		authzEnforced:             authzEnforced,
		DKGCoordinatorAddress:     dkgCoordinatorAddress,
		SupportedNetworks:         supportedNetworks,
		BitcoindConfigs:           operatorConfig.Bitcoind,
		Lrc20Configs:              operatorConfig.Lrc20,
		ServerCertPath:            serverCertPath,
		ServerKeyPath:             serverKeyPath,
		DKGLimitOverride:          dkgLimitOverride,
		RunDirectory:              runDirectory,
		ReturnDetailedErrors:      operatorConfig.ReturnDetailedErrors,
		ReturnDetailedPanicErrors: operatorConfig.ReturnDetailedPanicErrors,
		RateLimiter:               rateLimiter,
		Tracing:                   operatorConfig.Tracing,
		Database:                  *operatorConfig.Database,
	}, nil
}

func (c *Config) IsNetworkSupported(network common.Network) bool {
	for _, supportedNetwork := range c.SupportedNetworks {
		if supportedNetwork == network {
			return true
		}
	}
	return false
}

func NewRDSAuthToken(ctx context.Context, uri *url.URL) (string, error) {
	awsRegion := os.Getenv("AWS_REGION")
	if awsRegion == "" {
		return "", fmt.Errorf("AWS_REGION is not set")
	}
	awsRoleArn := os.Getenv("AWS_ROLE_ARN")
	if awsRoleArn == "" {
		return "", fmt.Errorf("AWS_ROLE_ARN is not set")
	}
	awsWebIdentityTokenFile := os.Getenv("AWS_WEB_IDENTITY_TOKEN_FILE")
	if awsWebIdentityTokenFile == "" {
		return "", fmt.Errorf("AWS_WEB_IDENTITY_TOKEN_FILE is not set")
	}
	podName := os.Getenv("POD_NAME")
	if podName == "" {
		return "", fmt.Errorf("POD_NAME is not set")
	}

	dbUser := uri.User.Username()
	dbEndpoint := uri.Host

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return "", err
	}

	client := sts.NewFromConfig(cfg)
	awsCreds := aws.NewCredentialsCache(stscreds.NewWebIdentityRoleProvider(
		client,
		awsRoleArn,
		stscreds.IdentityTokenFile(awsWebIdentityTokenFile),
		func(o *stscreds.WebIdentityRoleOptions) {
			o.RoleSessionName = podName
		}))

	token, err := auth.BuildAuthToken(ctx, dbEndpoint, awsRegion, dbUser, awsCreds)
	if err != nil {
		return "", err
	}

	return token, nil
}

var OtelSQLSpanOptions = otelsql.SpanOptions{
	OmitConnResetSession: true,
	OmitConnPrepare:      true,
}

type DBConnector struct {
	uri    *url.URL
	isRDS  bool
	driver driver.Driver
	pool   *pgxpool.Pool
}

func NewDBConnector(ctx context.Context, dbCfg *DatabaseConfig) (*DBConnector, error) {
	uri, err := url.Parse(dbCfg.URI)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database path: %w", err)
	}

	otelWrappedDriver := otelsql.WrapDriver(stdlib.GetDefaultDriver(),
		otelsql.WithAttributes(semconv.DBSystemPostgreSQL),
		otelsql.WithSpanOptions(OtelSQLSpanOptions),
	)

	connector := &DBConnector{
		uri:    uri,
		isRDS:  dbCfg.IsRDS,
		driver: otelWrappedDriver,
	}

	// Only create pool for PostgreSQL
	if strings.HasPrefix(dbCfg.URI, "postgresql") {
		config, err := pgxpool.ParseConfig(dbCfg.URI)
		if err != nil {
			return nil, fmt.Errorf("failed to parse pool config: %w", err)
		}
		config.ConnConfig.Tracer = otelpgx.NewTracer()

		config.MinConns = int32(defaultPoolMinConns)
		if dbCfg.PoolMinConns != nil {
			config.MinConns = int32(*dbCfg.PoolMinConns)
		}

		config.MaxConns = int32(defaultPoolMaxConns)
		if dbCfg.PoolMaxConns != nil {
			config.MaxConns = int32(*dbCfg.PoolMaxConns)
		}

		config.MaxConnLifetime = defaultPoolMaxConnLifetime
		if dbCfg.PoolMaxConnLifetime != nil {
			config.MaxConnLifetime = *dbCfg.PoolMaxConnLifetime
		}

		config.MaxConnIdleTime = defaultPoolMaxConnIdleTime
		if dbCfg.PoolMaxConnIdleTime != nil {
			config.MaxConnIdleTime = *dbCfg.PoolMaxConnIdleTime
		}

		config.HealthCheckPeriod = defaultPoolHealthCheckPeriod
		if dbCfg.PoolHealthCheckPeriod != nil {
			config.HealthCheckPeriod = *dbCfg.PoolHealthCheckPeriod
		}

		if dbCfg.PoolMaxConnLifetimeJitter != nil {
			config.MaxConnLifetimeJitter = *dbCfg.PoolMaxConnLifetimeJitter
		}

		if dbCfg.IsRDS {
			config.BeforeConnect = func(ctx context.Context, cfg *pgx.ConnConfig) error {
				token, err := NewRDSAuthToken(ctx, uri)
				if err != nil {
					return fmt.Errorf("failed to get RDS auth token: %w", err)
				}
				cfg.Password = token
				return nil
			}
		}

		pool, err := pgxpool.NewWithConfig(ctx, config)
		if err != nil {
			return nil, fmt.Errorf("failed to create connection pool: %w", err)
		}
		connector.pool = pool
	}

	return connector, nil
}

func (c *DBConnector) Connect(ctx context.Context) (driver.Conn, error) {
	if !c.isRDS {
		return c.driver.Open(c.uri.String())
	}
	uri := c.uri
	token, err := NewRDSAuthToken(ctx, c.uri)
	if err != nil {
		return nil, err
	}
	uri.User = url.UserPassword(uri.User.Username(), token)
	return c.driver.Open(uri.String())
}

func (c *DBConnector) Driver() driver.Driver {
	return c.driver
}

func (c *DBConnector) Pool() *pgxpool.Pool {
	return c.pool
}

func (c *DBConnector) Close() {
	if c.pool != nil {
		c.pool.Close()
	}
}

// LoadOperators loads the operators from the given file path.
func LoadOperators(filePath string) (map[string]*SigningOperator, error) {
	operators := make(map[string]*SigningOperator)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var yamlObj interface{}
	if err := yaml.Unmarshal(data, &yamlObj); err != nil {
		return nil, err
	}

	jsonStr, err := json.Marshal(yamlObj)
	if err != nil {
		return nil, err
	}

	var operatorList []*SigningOperator
	if err := json.Unmarshal(jsonStr, &operatorList); err != nil {
		return nil, err
	}

	for _, operator := range operatorList {
		operators[operator.Identifier] = operator
	}
	return operators, nil
}

// GetSigningOperatorList returns the list of signing operators.
func (c *Config) GetSigningOperatorList() map[string]*pb.SigningOperatorInfo {
	operatorList := make(map[string]*pb.SigningOperatorInfo)
	for _, operator := range c.SigningOperatorMap {
		operatorList[operator.Identifier] = operator.MarshalProto()
	}
	return operatorList
}

// AuthzEnforced returns whether authorization is enforced
func (c *Config) AuthzEnforced() bool {
	return c.authzEnforced
}

func (c *Config) IdentityPublicKey() []byte {
	return c.SigningOperatorMap[c.Identifier].IdentityPublicKey
}

func (c *Config) GetRateLimiterConfig() *middleware.RateLimiterConfig {
	return &middleware.RateLimiterConfig{
		Window:      c.RateLimiter.Window,
		MaxRequests: c.RateLimiter.MaxRequests,
		Methods:     c.RateLimiter.Methods,
	}
}

const (
	defaultTokenTransactionExpiryDuration = 3 * time.Minute
)

// setLrc20Defaults sets default values for Lrc20Config fields if they are zero.
func setLrc20Defaults(lrc20Configs map[string]Lrc20Config) {
	for k, v := range lrc20Configs {
		if v.TransactionExpiryDuration == 0 {
			slog.Info("TokenTransactionExpiryDuration not set, using default value", "default_duration", defaultTokenTransactionExpiryDuration)
			v.TransactionExpiryDuration = defaultTokenTransactionExpiryDuration
		}
		lrc20Configs[k] = v
	}
}
