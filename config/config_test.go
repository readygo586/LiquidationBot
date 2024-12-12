package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConfig(t *testing.T) {
	config, err := New("../config.yml")
	require.NoError(t, err)
	logger.Printf("db:%+v\n", config.DB)
	logger.Printf("comptroller:%+v\n", config.Comptroller)
	logger.Printf("override:%v\n", config.Override)
	logger.Printf("StartHeight:%v\n", config.StartHeight)
	logger.Printf("PrivateKey:%v\n", config.PrivateKey)
}
