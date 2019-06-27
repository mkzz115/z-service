package log

import "testing"

func TestInfo(t *testing.T) {
	pth := "/Users/mackie/goproj/src/github.com/mkzz115/z-service/output/test.log"
	err := Init(pth)
	t.Logf("err:%v\n", err)
	Info("uid[0]")

}
