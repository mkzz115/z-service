package idgenutil

import (
	"testing"
)

func TestCreateId(t *testing.T) {
	for i := 1; i <= 10; i++ {
		testi(t, i)
	}
}

func testi(t *testing.T, i int) {
	u, err := CreateId()
	t.Logf("test_%d\t, u[%d],err[%v]\n", i, u, err)
}
