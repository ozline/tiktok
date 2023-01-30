package main

import (
	"context"

	kitex_gen "github.com/ozline/tiktok/services/user/kitex_gen/kitex_gen"
)

// KitexProtoBufImpl implements the last service interface defined in the IDL.
type KitexProtoBufImpl struct{}

// MyHandT1 implements the KitexProtoBufImpl interface.
func (s *KitexProtoBufImpl) MyHandT1(ctx context.Context, req *kitex_gen.Request1) (resp *kitex_gen.Response, err error) {
	resp = &kitex_gen.Response{}
	resp.Message = req.Message
	return
}
