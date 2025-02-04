package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/readygo586/LiquidationBot/config"
	"github.com/readygo586/LiquidationBot/scanner"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	cobra.EnableCommandSorting = false

	rootCmd := &cobra.Command{
		Use:   "jupiterd",
		Short: "jupiter liquidation bot Daemon (server)",
	}

	rootCmd.AddCommand(StartCmd())
	rootCmd.AddCommand(VersionCmd())
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("err:%v", err)
		os.Exit(1)
	}

}

// StartCmd runs the service passed in, either stand-alone or in-process with
// Tendermint.
func StartCmd() *cobra.Command {
	var configFile string
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Run the jupiter liquidation bot server",
		Long:  `Run the jupiter liquidation bot server`,
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Info("starting jupiter liquidation bot")
			cfg, err := config.New(configFile)
			if err != nil {
				panic(err)
			}

			scanner.Start(cfg)
			return nil
		},
	}
	cmd.PersistentFlags().StringVarP(&configFile, "config", "f", "../../config.yml", "config file (default is ../config.yaml)")
	return cmd
}

// StartCmd runs the service passed in, either stand-alone or in-process with
// Tendermint.
func VersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Get jupiter liquidation bot version",
		Long:  `Get jupiter liquidation bot version`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("jupiter liquidation bot v0.1")
			return nil
		},
	}
	return cmd
}
