package server

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/readygo67/LiquidationBot/config"
	dbm "github.com/readygo67/LiquidationBot/db"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"
)

const DefaultStartHeigt = uint64(13000000)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)

func Start(cfg *config.Config) error {
	client, err := ethclient.Dial(cfg.RpcUrl)
	if err != nil {
		return err
	}

	db, err := dbm.NewDB(cfg.DB)
	if err != nil {
		return err
	}
	defer db.Close()

	startHeight := DefaultStartHeigt
	var storedHeight uint64
	exist, err := db.Has(dbm.LatestHandledHeightStoreKey(), nil)
	if exist {
		bz, err := db.Get(dbm.LatestHandledHeightStoreKey(), nil)
		if err != nil {
			return err
		}
		storedHeight = big.NewInt(0).SetBytes(bz).Uint64()
		startHeight = storedHeight
	}
	logger.Printf("startHeight:%v, storedHeight:%v, configHeight:%v\n", startHeight, storedHeight, cfg.StartHeight)
	if cfg.Override {
		startHeight = cfg.StartHeight
	}
	err = db.Put(dbm.LatestHandledHeightStoreKey(), big.NewInt(0).SetUint64(startHeight).Bytes(), nil)
	if err != nil {
		panic(err)
	}

	syncer := NewSyncer(client, db, cfg.Comptroller, cfg.Oracle, cfg.PancakeRouter, cfg.Liquidator, cfg.PrivateKey)

	syncer.Start()

	waitExit()

	syncer.Stop()
	return nil
}

func waitExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	i := <-c
	log.Printf("Received interrupt[%v], shutting down...\n", i)
}
