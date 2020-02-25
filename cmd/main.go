package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/evzpav/crypto-arbitrage/pkg/arbitrage"
	"github.com/spf13/cobra"
)

const defaultConfigPath string = "./configs.yaml"

func main() {

	arbitrageApp, err := arbitrage.New(defaultConfigPath)
	if err != nil {
		log.Fatal(err)
	}
	arbitrageCmd := buildArbitrageCommand(arbitrageApp)
	if err := arbitrageCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func buildArbitrageCommand(arb *arbitrage.Arbitrage) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "arbitrage",
		Long: "Check Arbitrage Opportunity",
		Args: cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {

			coins := args[0]
			exchangeNames := args[1]
			targetAmount := args[2]

			target, err := strconv.ParseFloat(targetAmount, 64)
			if err != nil {
				log.Fatalf("Failed to parse target amount: %v", targetAmount)
			}
			arb.Run(convertToSlice(exchangeNames), convertToSlice(coins), target)
		},
	}

	return cmd
}

func convertToSlice(s string) []string {
	return strings.Split(s, ",")
}
