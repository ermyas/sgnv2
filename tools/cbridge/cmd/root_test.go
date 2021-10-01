package cmd

import "testing"

func TestGetTokenAddr(t *testing.T) {
	cfg := &OneChainConfig{
		USDT: "0x12345678",
	}
	t.Error(cfg.GetTokenAddr("USD"))
}
