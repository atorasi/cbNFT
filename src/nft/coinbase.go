package nft

import (
	"github/atorasi/cbnft/src/account"
	"github/atorasi/cbnft/src/constants"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewMinter(client *ethclient.Client, wallet account.Wallet) Minter {
	return Minter{
		Client:  client,
		Wallet:  wallet,
		Account: account.Account{},
	}
}

func (app Minter) MintCoinbase() (*types.Receipt, error) {
	contractAddr := common.HexToAddress(constants.CONTRACTADDR)
	contractAbi, _ := app.Account.ReadAbi(constants.MINTER_ABI)

	dataArg := struct {
		Proof                  [][32]byte
		QuantityLimitPerWallet *big.Int
		PricePerToken          *big.Int
		Currency               common.Address
	}{
		Proof:                  [][32]byte{},
		QuantityLimitPerWallet: big.NewInt(2),
		PricePerToken:          big.NewInt(0),
		Currency:               common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"),
	}

	calldata, err := contractAbi.Pack("claim",
		app.Wallet.PublicKey,
		big.NewInt(0),
		big.NewInt(1),
		common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"),
		big.NewInt(0),
		dataArg,
		[]byte{},
	)
	if err != nil {
		return nil, err
	}

	txHash, err := app.Account.SendTransaction(app.Client, app.Wallet, contractAddr, calldata, big.NewInt(0), "BASE")
	if err != nil {
		return nil, err
	}

	return txHash, nil
}
