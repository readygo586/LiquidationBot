package dbm

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/syndtr/goleveldb/leveldb/util"
	"math/big"
	"os"
	"testing"
)

func TestAccessDB(t *testing.T) {
	db, err := NewDB("testdb1")
	require.NoError(t, err)

	defer db.Close()
	defer os.RemoveAll("testdb1")

	//latest handled height
	db.Put(LatestHandledHeightStoreKey(), big.NewInt(12561332).Bytes(), nil)
	bz, _ := db.Get(LatestHandledHeightStoreKey(), nil)
	height := big.NewInt(0).SetBytes(bz)
	require.EqualValues(t, 12561332, height.Uint64())

	//borrowers
	for i := 0; i < 10; i++ {
		db.Put(BorrowersStoreKey([]byte(fmt.Sprintf("account%v", i))), []byte(fmt.Sprintf("account%v", i)), nil)
	}

	iter0 := db.NewIterator(util.BytesPrefix(BorrowersPrefix), nil)
	defer iter0.Release()
	t.Logf("borrows address")
	for iter0.Next() {
		fmt.Printf("%v\n", string(iter0.Value()))
	}

	for i := 10; i < 20; i++ {
		db.Put(LiquidationBelow1P0StoreKey([]byte(fmt.Sprintf("account%v", i))), []byte(fmt.Sprintf("account%v", i)), nil)
	}

	iter1 := db.NewIterator(util.BytesPrefix(LiquidationBelow1P0Prefix), nil)
	defer iter1.Release()
	t.Logf("liquidation below 1 address")
	for iter1.Next() {
		fmt.Printf("%v\n", string(iter1.Value()))
	}

}

func Test_MarketMember(t *testing.T) {
	db, err := NewDB("testdb1")
	require.NoError(t, err)

	defer db.Close()
	defer os.RemoveAll("testdb1")

	marketVBTC := []byte("market_vBTC")
	marketVETH := []byte("market_vETH")
	//market member
	for i := 0; i < 10; i++ {
		account := []byte(fmt.Sprintf("account%v", i))
		db.Put(MarketMemberStoreKey(marketVBTC, account), account, nil)
	}

	for i := 5; i < 15; i++ {
		account := []byte(fmt.Sprintf("account%v", i))
		db.Put(MarketMemberStoreKey(marketVETH, account), account, nil)
	}

	prefixVBTC := append(MarketMemberPrefix, marketVBTC...)
	iter1 := db.NewIterator(util.BytesPrefix(prefixVBTC), nil)
	defer iter1.Release()
	t.Logf("%s", marketVBTC)
	for iter1.Next() {
		fmt.Printf("%v\n", string(iter1.Value()))
	}

	prefixVETH := append(MarketMemberPrefix, marketVETH...)
	iter2 := db.NewIterator(util.BytesPrefix(prefixVETH), nil)
	defer iter2.Release()
	t.Logf("%s", marketVBTC)
	for iter2.Next() {
		fmt.Printf("%v\n", string(iter2.Value()))
	}

}
