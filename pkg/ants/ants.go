package ants

import (
	"errors"
	"sync"

	"github.com/panjf2000/ants"
)

var (
	AntsPool *ants.PoolWithFunc
	Wg       sync.WaitGroup
)

func Init() {
	ants_Pool, err := ants.NewPoolWithFunc(500, func(payload interface{}) {
		defer Wg.Done()
	})
	if err != nil {
		panic(errors.New("[ants goroutine init error]"))
	}
	AntsPool = ants_Pool
}
