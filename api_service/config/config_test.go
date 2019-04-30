package config

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestFromToml(t *testing.T) {
	wd, _ := os.Getwd()
	t.Logf("path:%s", wd)
	cfg, err := FromToml(wd + "/.test/test1.toml")
	require.Nil(t, err)
	t.Logf("result:%+v, %v\n", cfg, err)
}
