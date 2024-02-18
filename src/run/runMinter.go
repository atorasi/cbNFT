package run

import (
	"github/atorasi/cbnft/src/account"
	"github/atorasi/cbnft/src/constants"
	"github/atorasi/cbnft/src/deposit"
	"github/atorasi/cbnft/src/nft"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

func RunMinter(index int, wallet account.Wallet, client *ethclient.Client) (string, error) {
	var module string
	if constants.SETTINGS.NeedOkx {
		module, err := okxWithdrawal(index, "BASE", client, wallet)
		if err != nil {
			return module, err
		}
	}

	module, err := mintNft(index, client, wallet)
	if err != nil {
		return module, err
	}

	return "", nil
}

func okxWithdrawal(index int, sideChain string, client *ethclient.Client, wallet account.Wallet) (string, error) {
	balanceSideStart, err := account.Account.NativeBalance(account.Account{}, client, wallet)
	if err != nil {
		return "Get balance", err
	}

	log.Printf("Acc.%d | Preparing to Okx withdrawal", index)
	depositClient := deposit.NewDepositApp(wallet)
	if _, err := depositClient.OkxWithdraw(sideChain); err != nil {
		return "OKX", err
	}
	log.Printf("Acc.%d | Succesfully withdrew from OKX, waiting for funds", index)

	for {
		newBalance, err := account.Account.NativeBalance(account.Account{}, client, wallet)
		if err != nil {
			return "Get balance", err
		}
		if newBalance != balanceSideStart {
			log.Printf("Acc.%d | Funds deposited", index)
			break
		}
		time.Sleep(time.Duration(30) * time.Second)
		log.Printf("Acc.%d | Didnt get funds yet, sleep 30 seconds.", index)

	}

	return "", nil
}

func mintNft(index int, client *ethclient.Client, wallet account.Wallet) (string, error) {
	log.Printf("Acc.%d | Preparing to Mint Coinbase NFT", index)
	minter := nft.NewMinter(client, wallet)
	if _, err := minter.MintCoinbase(); err != nil {
		return "Coinbase", err
	}

	return "", nil
}
