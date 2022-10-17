package sender

import "github.com/amir/raidman"

type RiemannClient interface {
	Close() error
	Send(event *raidman.Event) error
}
