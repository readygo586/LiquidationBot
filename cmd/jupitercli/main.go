package main

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/readygo586/LiquidationBot/config"
	dbm "github.com/readygo586/LiquidationBot/db"
	"github.com/readygo586/LiquidationBot/scanner"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
	"github.com/syndtr/goleveldb/leveldb/util"
	"math/big"
	"os"
)

func main() {
	cobra.EnableCommandSorting = false
	rootCmd := &cobra.Command{
		Use:   "jupitercli",
		Short: "jupiter liquidation bot client",
	}

	rootCmd.AddCommand(queryCmd())
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("err:%v", err)
		os.Exit(1)
	}
}

// StartCmd runs the service passed in, either stand-alone or in-process with
// Tendermint.
func queryCmd() *cobra.Command {
	var configFile string
	cmd := &cobra.Command{
		Use:   "query",
		Short: "querying subcommand",
	}

	cmd.AddCommand(
		accountCommand(&configFile),
		listCommand(&configFile),
		heightCommand(&configFile),
	)
	cmd.PersistentFlags().StringVarP(&configFile, "config", "f", "../config.yml", "config file (default is ../config.yml)")
	return cmd
}

func heightCommand(configFile *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "height",
		Short: "syncing height",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.New(*configFile)
			if err != nil {
				return err
			}

			if !fileExists(cfg.DB) {
				return fmt.Errorf("db does not exist")
			}

			db, err := dbm.NewDB(cfg.DB)
			if err != nil {
				return err
			}

			bz, err := db.Get(dbm.LatestHandledHeightStoreKey(), nil)
			if err != nil {
				return err
			}

			fmt.Printf("current syncing height:%v\n", big.NewInt(0).SetBytes(bz).Int64())
			return nil
		},
	}
	return cmd
}

func accountCommand(configFile *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account [0x...]",
		Short: "account info",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.New(*configFile)
			if err != nil {
				return err
			}

			if !fileExists(cfg.DB) {
				return fmt.Errorf("db does not exist")
			}

			db, err := dbm.NewDB(cfg.DB)
			if err != nil {
				return err
			}

			accountBytes := common.HexToAddress(args[0]).Bytes()
			bz, err := db.Get(dbm.AccountStoreKey(accountBytes), nil)
			if err != nil {
				fmt.Printf("can not found account:%v\n", args[0])
				return err
			}

			var info scanner.AccountInfo
			err = json.Unmarshal(bz, &info)
			if err != nil {
				return err
			}

			fmt.Printf("account:%v\n :%+v\n", args[0], info)
			return nil
		},
	}
	return cmd
}

func listCommand(configFile *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list [1.0]",
		Short: "list account whose health factor below assigned level",
		Long: `list account whose health factor below assigned level, currently the following levels are provided
               x<1.0, 1.0 <= x < 1.1, 1.1 <= x < 1.5, 1.5 <= x < 2.0, x > 2.0, x=255 for nonProfit`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.New(*configFile)
			if err != nil {
				return err
			}

			if !fileExists(cfg.DB) {
				return fmt.Errorf("db does not exist")
			}

			db, err := dbm.NewDB(cfg.DB)
			if err != nil {
				return err
			}

			level, err := decimal.NewFromString(args[0])
			if err != nil {
				return fmt.Errorf("invalid parameter")
			}

			var prefix []byte
			if level.Cmp(scanner.DecimalNonProfit) == 0 {
				prefix = dbm.LiquidationNonProfitPrefix
			} else {
				if level.Cmp(scanner.Decimal1P0) != 1 {
					prefix = dbm.LiquidationBelow1P0Prefix
				} else if level.Cmp(scanner.Decimal1P1) != 1 {
					prefix = dbm.LiquidationBelow1P1Prefix
				} else if level.Cmp(scanner.Decimal1P5) != 1 {
					prefix = dbm.LiquidationBelow1P5Prefix
				} else if level.Cmp(scanner.Decimal2P0) != 1 {
					prefix = dbm.LiquidationBelow2P0Prefix
				} else {
					prefix = dbm.LiquidationAbove2P0Prefix
				}
			}

			iter := db.NewIterator(util.BytesPrefix(prefix), nil)
			defer iter.Release()
			count := 0
			fmt.Printf("account below%v:\n", args[0])
			for iter.Next() {
				fmt.Printf("%v,", common.BytesToAddress(iter.Value()))
				count++
			}
			fmt.Printf("\n total:%v\n", count)
			return nil
		},
	}
	return cmd
}

// filesExists reports whether the named file or directory exists.
func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
