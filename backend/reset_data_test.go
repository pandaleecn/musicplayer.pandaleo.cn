package main

import (
	"github.com/kataras/iris"
	"testing"
)

func TestResetData(t *testing.T) {
	t.Skip()
	getOnAuth(t, "resetData", iris.StatusOK, true, "重置数据成功")
}
