package nft

import (
	"github/atorasi/cbnft/src/account"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Minter struct {
	Client  *ethclient.Client
	Wallet  account.Wallet
	Account account.Account
}
