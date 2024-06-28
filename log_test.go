package colog_test

import (
	"github.com/HuXin0817/colog"
	"testing"
)

func TestLog(t *testing.T) {
	err := colog.OpenLog("test/test.log")
	if err != nil {
		panic(err)
	}
	colog.Info("hello", "world")
	//os.RemoveAll("test/")
}
