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

func Test_MarshalPendingLiquidation(t *testing.T) {
	tx1 := PendingLiquidation{
		Hash:   common.HexToHash("0x07f2ba8e0b76ab3140c534d7353785dc8df6151747c23a2863976b9688865e56"),
		Height: uint64(47797935),
	}

	bz, err := json.Marshal(tx1)
	assert.NoError(t, err)
	fmt.Printf("pendingTx: %v\n", string(bz))

	var tx2 PendingLiquidation
	err = json.Unmarshal(bz, &tx2)
	assert.NoError(t, err)
	assert.Equal(t, tx1.Hash, tx2.Hash)
	assert.Equal(t, tx1.Height, tx2.Height)
}

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
