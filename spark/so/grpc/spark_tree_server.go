package grpc

import (
	"context"

	pb "github.com/lightsparkdev/spark/proto/spark_tree"
	"github.com/lightsparkdev/spark/so"
	"github.com/lightsparkdev/spark/so/ent"
	tree "github.com/lightsparkdev/spark/so/tree"
)

// SparkTreeServer is the grpc server for the Spark protocol.
// It will be used by the user or Spark service provider.
type SparkTreeServer struct {
	pb.UnimplementedSparkTreeServiceServer
	config *so.Config
	scorer tree.Scorer
}

// NewSparkTreeServer creates a new SparkTreeServer.
func NewSparkTreeServer(config *so.Config, dbClient *ent.Client) *SparkTreeServer {
	scorer := tree.NewPolarityScorer(dbClient)
	go scorer.Start()
	return &SparkTreeServer{config: config, scorer: scorer}
}

// GetLeafDenominationCounts returns the number of leaves for each denomination.
func (*SparkTreeServer) GetLeafDenominationCounts(ctx context.Context, req *pb.GetLeafDenominationCountsRequest) (*pb.GetLeafDenominationCountsResponse, error) {
	return tree.GetLeafDenominationCounts(ctx, req)
}

// ProposeTreeDenominations proposes the denominations for a new tree.
func (*SparkTreeServer) ProposeTreeDenominations(ctx context.Context, req *pb.ProposeTreeDenominationsRequest) (*pb.ProposeTreeDenominationsResponse, error) {
	return tree.ProposeTreeDenominations(ctx, req)
}

// FetchPolarityScores fetches the polarity scores for a given SSP.
func (s *SparkTreeServer) FetchPolarityScores(req *pb.FetchPolarityScoreRequest, stream pb.SparkTreeService_FetchPolarityScoresServer) error {
	return s.scorer.FetchPolarityScores(req, stream)
}
