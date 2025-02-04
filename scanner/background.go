package scanner

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	dbm "github.com/readygo586/LiquidationBot/db"
	"github.com/syndtr/goleveldb/leveldb/util"
	"time"
)

func (s *Scanner) SyncAccountsBelow1P0Loop() {
	defer s.wg.Done()
	db := s.db

	t := time.NewTimer(0)
	defer t.Stop()

	count := 1
	for {
		select {
		case <-s.quitCh:
			return
		case <-t.C:
			logger.Printf("%vth background sync t @ %v\n", count, time.Now())
			count++

			var accounts []common.Address
			iter := db.NewIterator(util.BytesPrefix(dbm.LiquidationBelow1P0Prefix), nil)
			for iter.Next() {
				accounts = append(accounts, common.BytesToAddress(iter.Value()))
			}
			iter.Release()
			fmt.Printf("SyncAccountsBelow1P0Loop: %v\n", accounts)
			s.topAccountSyncCh <- accounts

			t.Reset(time.Second * SyncIntervalBelow1P0)
		}
	}
}

func (s *Scanner) SyncAccountsBackgroundLoop() {
	defer s.wg.Done()
	db := s.db

	t := time.NewTimer(0)
	defer t.Stop()

	count := 1
	for {
		select {
		case <-s.quitCh:
			return
		case <-t.C:
			logger.Printf("%vth background sync t @ %v\n", count, time.Now())
			count++

			var accounts []common.Address
			iter := db.NewIterator(util.BytesPrefix(dbm.BorrowersPrefix), nil)
			for iter.Next() {
				accounts = append(accounts, common.BytesToAddress(iter.Value()))
			}
			iter.Release()
			fmt.Printf("SyncAccountsBackgroundLoop: %v\n", accounts)
			s.lowAccountSyncCh <- accounts

			t.Reset(time.Second * SyncIntervalBackGround)
		}
	}
}
