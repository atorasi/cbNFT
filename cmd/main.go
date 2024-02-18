package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github/atorasi/cbnft/src/account"
	"github/atorasi/cbnft/src/constants"
	"github/atorasi/cbnft/src/run"
	"github/atorasi/cbnft/utils"
)

func main() {
	clearTerminal()
	fmt.Printf("%s\n\n", constants.LOGO)
	log.Println("------- t.me/tripleshizu t.me/tripleshizu -------")
	log.Println("Donate - 0x4163dfa9eE4A25e950ce1a0A2221FafA29fe2df6 - Any EVM")
	fmt.Println()

	walletSlice, err := account.SliceOfAccs()
	if err != nil {
		log.Fatal(err)
	}
	proxySlice, err := utils.ReadFile(`..\data\proxy.txt`)
	if err != nil {
		log.Fatal(err)
	}

	for index, wallet := range walletSlice {
		client, err := account.NewClient(index, proxySlice)
		if err != nil {
			log.Println("An error with RPC connection, check your node", err)
		}
		module, err := run.RunMinter(wallet.Index, wallet, client)
		if err != nil {
			log.Printf("Acc.%d | An error was occured with %s: %v", wallet.Index, module, err)
		}

		if constants.SETTINGS.NeedDelayAcc {
			acc := account.Account{}
			delay := acc.RandomInt(constants.SETTINGS.DelayAccMin, constants.SETTINGS.DelayAccMax)
			log.Printf("Acc.%d | sleep for %d seconds before the next account", wallet.Index, delay)
			time.Sleep(time.Duration(delay) * time.Second)

		}
	}

	log.Println("The software has shut down. Press Enter to exit.")
	fmt.Scanln()

	log.Println("------- t.me/tripleshizu t.me/tripleshizu -------")
	log.Println("Donate - 0x4163dfa9eE4A25e950ce1a0A2221FafA29fe2df6 - Any EVM")
}

func clearTerminal() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
	default:
		cmd = exec.Command("clear")
	}

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Ошибка при очистке терминала: %v", err)
	}
}
