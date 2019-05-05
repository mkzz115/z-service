package confutil

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFromToml1(t *testing.T) {
	wd, _ := os.Getwd()
	t.Logf("path:%s", wd)
	cfg, err := FromToml(wd + "/.test/test1.toml")
	require.Nil(t, err)
	t.Logf("result:%+v, %v\n", cfg, err)
}
