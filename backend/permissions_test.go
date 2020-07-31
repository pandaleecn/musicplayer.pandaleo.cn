package main

import (
	"github.com/kataras/iris"
	"testing"
)

func TestPermissions(t *testing.T) {
	getMore(t, "permissions", iris.StatusOK, true, "操作成功")
}

func TestImportPermissions(t *testing.T) {
	bImport(t, "permissions/import", iris.StatusOK, true, "成功导入18项数据", nil)
}
