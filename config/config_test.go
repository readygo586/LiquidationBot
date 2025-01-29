package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConfig(t *testing.T) {
	config, err := New("../config.yml")
	require.NoError(t, err)
	t.Logf("rpc_url:%+v\n", config.RpcUrl)
	t.Logf("oracle:%+v\n", config.Oracle)
	t.Logf("comptroller:%+v\n", config.Comptroller)
	t.Logf("vai_controller:%+v\n", config.VaiController)
	t.Logf("vai:%+v\n", config.Vai)
	t.Logf("private_key:%+v\n", config.PrivateKey)
	t.Logf("db:%+v\n", config.DB)
	t.Logf("StartHeight:%v\n", config.StartHeight)
	t.Logf("override:%v\n", config.Override)
}
