package dbm

import (
	leveldb "github.com/syndtr/goleveldb/leveldb"
)

const HashLength = 32

// Hash to identify uniqueness
type Hash [HashLength]byte

var (
	KeyLatestHandledHeight     = []byte("latest_handled_height")
	BorrowersPrefix            = []byte("borrowers") //prefix with all borrowers
	AccountPrefix              = []byte("account")
	MarketPrefix               = []byte("market")
	MarketMemberPrefix         = []byte("market_member")
	JoinedMarketPrefix         = []byte("joined_market")
	LiquidationBelow1P0Prefix  = []byte("liquidation_below_1p0")
	LiquidationBelow1P1Prefix  = []byte("liquidation_below_1p1")
	LiquidationBelow1P5Prefix  = []byte("liquidation_below_1p5")
	LiquidationBelow2P0Prefix  = []byte("liquidation_below_2p0")
	LiquidationAbove2P0Prefix  = []byte("liquidation_above_2p0")
	LiquidationNonProfitPrefix = []byte("liquidation_non_profit") //
	BadLiquidationTxPrefix     = []byte("bad_liquidation_tx")
	PendingLiquidationTxPrefix = []byte("pending_liquidation_tx")
)

func LatestHandledHeightStoreKey() []byte {
	return KeyLatestHandledHeight
}

func BorrowersStoreKey(account []byte) []byte {
	return append(BorrowersPrefix, account...)
}

// record account joined markets
func JoinedMarketMemberStoreKey(account []byte, address []byte) []byte {
	bz := append(JoinedMarketPrefix, account...)
	return append(bz, address...)
}

func MarketStoreKey(market []byte) []byte {
	return append(MarketPrefix, market...)
}
func MarketMemberStoreKey(market []byte, account []byte) []byte {
	bz := append(MarketMemberPrefix, market...)
	return append(bz, account...)
}

func AccountStoreKey(account []byte) []byte {
	return append(AccountPrefix, account...)
}

func LiquidationBelow1P0StoreKey(address []byte) []byte {
	return append(LiquidationBelow1P0Prefix, address...)
}

func LiquidationBelow1P1StoreKey(address []byte) []byte {
	return append(LiquidationBelow1P1Prefix, address...)
}

func LiquidationBelow1P5StoreKey(address []byte) []byte {
	return append(LiquidationBelow1P5Prefix, address...)
}

func LiquidationBelow2P0StoreKey(address []byte) []byte {
	return append(LiquidationBelow2P0Prefix, address...)
}

func LiquidationAbove2P0StoreKey(address []byte) []byte {
	return append(LiquidationAbove2P0Prefix, address...)
}

func LiquidationNonProfitStoreKey(address []byte) []byte {
	return append(LiquidationNonProfitPrefix, address...)
}

func BadLiquidationTxStoreKey(address []byte) []byte {
	return append(BadLiquidationTxPrefix, address...)
}

func PendingLiquidationTxStoreKey(address []byte) []byte {
	return append(PendingLiquidationTxPrefix, address...)
}

func NewDB(path string) (*leveldb.DB, error) {
	return leveldb.OpenFile(path, nil)
}
