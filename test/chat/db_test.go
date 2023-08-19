package test

import (
	"testing"

	"github.com/ozline/tiktok/cmd/video/dal/db"
)

func testDB(t *testing.T) {
	t.Log("----------BEGIN------------")
	db.Init()
	if db.DB==nil{
		t.Fail()
	}
	t.Log("----------SUCCESS----------")
}
