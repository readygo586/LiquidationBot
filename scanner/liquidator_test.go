package scanner

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/readygo586/LiquidationBot/config"
	dbm "github.com/readygo586/LiquidationBot/db"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_Liquidate(t *testing.T) {
	cfg, err := config.New("../config.yml")
	c, err := ethclient.Dial(cfg.RpcUrl)

	db, err := dbm.NewDB("testdb1")
	assert.NoError(t, err)
	defer db.Close()
	defer os.RemoveAll("testdb1")

	account := common.HexToAddress("0x4e3CC26bce18b0F420155DCE102c976aF057867E")
	s := NewScanner(c, db, cfg.Comptroller, cfg.VaiController, cfg.Vai, cfg.Oracle, cfg.PrivateKey)
	err = s.syncOneAccount(account)
	assert.NoError(t, err)

	bz, err := s.db.Get(dbm.AccountStoreKey(account.Bytes()), nil)
	assert.NoError(t, err)

	var info AccountInfo
	err = json.Unmarshal(bz, &info)
	assert.NoError(t, err)
	assert.Equal(t, account, info.Account)
	assert.Equal(t, 1, len(info.Assets))
	fmt.Printf("info: %+v\n", info.toReadable())

	bz, err = s.db.Get(dbm.LiquidationBelow1P0StoreKey(account.Bytes()), nil)
	assert.Equal(t, account.Bytes(), bz)

	err = s.liquidate(&info)
	assert.NoError(t, err)
}
