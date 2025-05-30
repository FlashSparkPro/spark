import { TokenPubkey, TokenPubkeyAnnouncement } from "@buildonspark/lrc20-sdk";
import {
  ListAllTokenTransactionsCursor,
  OperationType,
} from "@buildonspark/lrc20-sdk/proto/rpc/v1/types";
import {
  NetworkError,
  SparkWallet,
  SparkWalletProps,
} from "@buildonspark/spark-sdk";
import {
  decodeSparkAddress,
  encodeSparkAddress,
} from "@buildonspark/spark-sdk/address";
import { OutputWithPreviousTransactionData } from "@buildonspark/spark-sdk/proto/spark";
import { ConfigOptions } from "@buildonspark/spark-sdk/services/wallet-config";
import {
  bytesToHex,
  bytesToNumberBE,
  hexToBytes,
} from "@noble/curves/abstract/utils";
import { TokenFreezeService } from "./services/freeze.js";
import { IssuerTokenTransactionService } from "./services/token-transactions.js";
import { GetTokenActivityResponse, TokenDistribution } from "./types.js";
import { convertTokenActivityToHexEncoded } from "./utils/type-mappers.js";
import { NotImplementedError } from "@buildonspark/spark-sdk";

const BURN_ADDRESS = "02".repeat(33);

export type IssuerTokenInfo = {
  tokenPublicKey: string;
  tokenName: string;
  tokenSymbol: string;
  tokenDecimals: number;
  maxSupply: bigint;
  isFreezable: boolean;
};

export class IssuerSparkWallet extends SparkWallet {
  private issuerTokenTransactionService: IssuerTokenTransactionService;
  private tokenFreezeService: TokenFreezeService;

  public static async initialize(options: SparkWalletProps) {
    const wallet = new IssuerSparkWallet(options.options);

    const initResponse = await wallet.initWallet(options.mnemonicOrSeed);
    return {
      wallet,
      ...initResponse,
    };
  }

  protected constructor(configOptions?: ConfigOptions) {
    super(configOptions);
    this.issuerTokenTransactionService = new IssuerTokenTransactionService(
      this.config,
      this.connectionManager,
    );
    this.tokenFreezeService = new TokenFreezeService(
      this.config,
      this.connectionManager,
    );
  }

  public async getIssuerTokenBalance(): Promise<{
    balance: bigint;
  }> {
    const publicKey = await super.getIdentityPublicKey();
    const balanceObj = await this.getBalance();

    if (!balanceObj.tokenBalances || !balanceObj.tokenBalances.has(publicKey)) {
      return {
        balance: 0n,
      };
    }
    return {
      balance: balanceObj.tokenBalances.get(publicKey)!.balance,
    };
  }

  public async getIssuerTokenInfo(): Promise<IssuerTokenInfo | null> {
    const lrc20Client = await this.lrc20ConnectionManager.createLrc20Client();

    try {
      const tokenInfo = await lrc20Client.getTokenPubkeyInfo({
        publicKeys: [hexToBytes(await super.getIdentityPublicKey())],
      });

      const info = tokenInfo.tokenPubkeyInfos[0];
      return {
        tokenPublicKey: bytesToHex(info.announcement!.publicKey!.publicKey),
        tokenName: info.announcement!.name,
        tokenSymbol: info.announcement!.symbol,
        tokenDecimals: Number(bytesToNumberBE(info.announcement!.decimal)),
        isFreezable: info.announcement!.isFreezable,
        maxSupply: bytesToNumberBE(info.announcement!.maxSupply),
      };
    } catch (error) {
      throw new NetworkError("Failed to get token info", {
        operation: "getIssuerTokenInfo",
        errorCount: 1,
        errors: error instanceof Error ? error.message : String(error),
      });
    }
  }

  public async mintTokens(tokenAmount: bigint): Promise<string> {
    var tokenPublicKey = await super.getIdentityPublicKey();

    const tokenTransaction =
      await this.issuerTokenTransactionService.constructMintTokenTransaction(
        hexToBytes(tokenPublicKey),
        tokenAmount,
      );

    return await this.issuerTokenTransactionService.broadcastTokenTransaction(
      tokenTransaction,
    );
  }

  public async burnTokens(
    tokenAmount: bigint,
    selectedOutputs?: OutputWithPreviousTransactionData[],
  ): Promise<string> {
    const burnAddress = encodeSparkAddress({
      identityPublicKey: BURN_ADDRESS,
      network: this.config.getNetworkType(),
    });
    return await this.transferTokens({
      tokenPublicKey: await super.getIdentityPublicKey(),
      tokenAmount,
      receiverSparkAddress: burnAddress,
      selectedOutputs,
    });
  }

  public async freezeTokens(
    sparkAddress: string,
  ): Promise<{ impactedOutputIds: string[]; impactedTokenAmount: bigint }> {
    await this.syncTokenOutputs();
    const tokenPublicKey = await super.getIdentityPublicKey();
    const decodedOwnerPubkey = decodeSparkAddress(
      sparkAddress,
      this.config.getNetworkType(),
    );
    const response = await this.tokenFreezeService!.freezeTokens(
      hexToBytes(decodedOwnerPubkey),
      hexToBytes(tokenPublicKey),
    );

    // Convert the Uint8Array to a bigint
    const tokenAmount = bytesToNumberBE(response.impactedTokenAmount);

    return {
      impactedOutputIds: response.impactedOutputIds,
      impactedTokenAmount: tokenAmount,
    };
  }

  public async unfreezeTokens(
    sparkAddress: string,
  ): Promise<{ impactedOutputIds: string[]; impactedTokenAmount: bigint }> {
    await this.syncTokenOutputs();
    const tokenPublicKey = await super.getIdentityPublicKey();
    const decodedOwnerPubkey = decodeSparkAddress(
      sparkAddress,
      this.config.getNetworkType(),
    );
    const response = await this.tokenFreezeService!.unfreezeTokens(
      hexToBytes(decodedOwnerPubkey),
      hexToBytes(tokenPublicKey),
    );
    const tokenAmount = bytesToNumberBE(response.impactedTokenAmount);

    return {
      impactedOutputIds: response.impactedOutputIds,
      impactedTokenAmount: tokenAmount,
    };
  }

  public async getIssuerTokenActivity(
    pageSize: number = 100,
    cursor?: ListAllTokenTransactionsCursor,
    operationTypes?: OperationType[],
    beforeTimestamp?: Date,
    afterTimestamp?: Date,
  ): Promise<GetTokenActivityResponse> {
    const lrc20Client = await this.lrc20ConnectionManager.createLrc20Client();

    try {
      const transactions = await lrc20Client.listTransactions({
        tokenPublicKey: hexToBytes(await super.getIdentityPublicKey()),
        cursor,
        pageSize,
        beforeTimestamp,
        afterTimestamp,
        operationTypes,
      });

      return convertTokenActivityToHexEncoded(transactions);
    } catch (error) {
      throw new NetworkError("Failed to get token activity", {
        operation: "listTransactions",
        errorCount: 1,
        errors: error instanceof Error ? error.message : String(error),
      });
    }
  }

  public async getIssuerTokenDistribution(): Promise<TokenDistribution> {
    throw new NotImplementedError("Token distribution is not yet supported");
  }

  public async announceTokenL1(
    tokenName: string,
    tokenTicker: string,
    decimals: number,
    maxSupply: bigint,
    isFreezable: boolean,
    feeRateSatsPerVb: number = 4.0,
  ): Promise<string> {
    await this.lrc20Wallet!.syncWallet();

    const tokenPublicKey = new TokenPubkey(this.lrc20Wallet!.pubkey);

    const announcement = new TokenPubkeyAnnouncement(
      tokenPublicKey,
      tokenName,
      tokenTicker,
      decimals,
      maxSupply,
      isFreezable,
    );

    try {
      const tx = await this.lrc20Wallet!.prepareAnnouncement(
        announcement,
        feeRateSatsPerVb,
      );

      return await this.lrc20Wallet!.broadcastRawBtcTransaction(
        tx.bitcoin_tx.toHex(),
      );
    } catch (error) {
      throw new NetworkError(
        "Failed to broadcast announcement transaction on L1",
        {
          operation: "broadcastRawBtcTransaction",
          errorCount: 1,
          errors: error instanceof Error ? error.message : String(error),
        },
      );
    }
  }
}
