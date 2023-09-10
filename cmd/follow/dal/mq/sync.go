package mq

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
)

type SyncFollow struct {
}

func (s *SyncFollow) SyncFollowMQ(ctx context.Context) error {
	defer FollowMQCli.ReleaseRes()

	msgs, err := FollowMQCli.Consume(ctx)
	if err != nil {
		klog.Error(err)
		return err
	}

	var forever chan struct{}

	go func() {
		for msg := range msgs {
			klog.Infof("Resolve msg: %s", msg.Body)

			err := FollowMQCli.FollowActionInsert(ctx, msg)
			if err != nil {
				klog.Errorf("Insert follow action: %s", err)
				continue
			}

			err = msg.Ack(false)
			if err != nil {
				klog.Error(err)
				continue
			}
		}
	}()

	<-forever

	return nil
}
