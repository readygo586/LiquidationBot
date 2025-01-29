package config

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConfig(t *testing.T) {
	config, err := New("../config_test.yml")
	require.NoError(t, err)
	assert.Equal(t, "testnet", config.Network)
	t.Logf("rpc_url:%+v\n", config.RpcUrl)
	t.Logf("network:%+v\n", config.Network)
	t.Logf("oracle:%+v\n", config.Oracle)
	t.Logf("comptroller:%+v\n", config.Comptroller)
	t.Logf("vai_controller:%+v\n", config.VaiController)
	t.Logf("vai:%+v\n", config.Vai)
	t.Logf("wbtc:%+v\n", config.WBTC)
	t.Logf("weth:%+v\n", config.WETH)
	t.Logf("private_key:%+v\n", config.PrivateKey)
	t.Logf("db:%+v\n", config.DB)
	t.Logf("StartHeight:%v\n", config.StartHeight)
	t.Logf("override:%v\n", config.Override)
}
