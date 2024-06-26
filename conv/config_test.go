package conv_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dymensionxyz/dymint/config"
	"github.com/dymensionxyz/dymint/conv"
	tmcfg "github.com/tendermint/tendermint/config"
)

func TestGetNodeConfig(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name        string
		input       func(*tmcfg.Config)
		ok          func(*config.NodeConfig) bool
		expectError bool
	}{
		{
			"empty",
			func(c *tmcfg.Config) {
				*c = tmcfg.Config{}
			},
			nil,
			true,
		},
		{
			"RootDir",
			func(c *tmcfg.Config) {
				c.BaseConfig.RootDir = "~/root"
			},
			func(nc *config.NodeConfig) bool { return nc.RootDir == "~/root" },
			false,
		},
		{
			"DBPath",
			func(c *tmcfg.Config) {
				c.BaseConfig.DBPath = "./database"
			},
			func(nc *config.NodeConfig) bool { return nc.DBPath == "./database" },
			false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var actual config.NodeConfig
			tmConfig := tmcfg.DefaultConfig()
			c.input(tmConfig)
			err := conv.GetNodeConfig(&actual, tmConfig)
			if c.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.True(t, c.ok(&actual))
			}
		})
	}
}
