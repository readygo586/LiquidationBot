package scanner

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/readygo586/LiquidationBot/config"
	dbm "github.com/readygo586/LiquidationBot/db"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_Approve(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	assert.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)
	tx, err := s.doApprove()
	assert.NoError(t, err)
	fmt.Printf("tx: %v\n", tx.Hash())

}
