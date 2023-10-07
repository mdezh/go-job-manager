package manager

import (
	"context"
	"sync"
	"time"
)

type Job func(context.Context) error

type DurationProvider func() time.Duration

func StaticDurationProvider(t time.Duration) DurationProvider {
	return func() time.Duration {
		return t
	}
}

type JobConfiguration struct {
	Interval DurationProvider
	Timeout  DurationProvider
}

type empty struct{}

type logger interface {
	Println(args ...any)
}

type jobRecord struct {
	name    string
	job     Job
	cfg     JobConfiguration
	mx      sync.RWMutex
	working bool
}

type manager struct {
	ctx       context.Context
	logger    logger
	startOnce sync.Once
	stopOnce  sync.Once
	started   chan empty
	stopped   chan empty
	done      chan empty
	wg        sync.WaitGroup
	jobs      []*jobRecord
}

func NewManager(ctx context.Context, logger logger) *manager {
	return &manager{
		ctx:     ctx,
		logger:  logger,
		started: make(chan empty),
		stopped: make(chan empty),
		done:    make(chan empty),
	}
}
