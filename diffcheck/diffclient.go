package diffcheck

import (
	"time"
)

// DiffClient holds the configuration for diff operations.
type DiffClient struct {
	DiffTimeout time.Duration
}

func New() *DiffClient {
	return &DiffClient{
		DiffTimeout: time.Second,
	}
}
