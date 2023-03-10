// Code generated by Kitex v0.4.4. DO NOT EDIT.

package tiktokcommentservice

import (
	server "github.com/cloudwego/kitex/server"
	comment "github.com/ozline/tiktok/kitex_gen/tiktok/comment"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler comment.TiktokCommentService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
