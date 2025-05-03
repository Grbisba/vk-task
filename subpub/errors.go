package subpub

import (
	"errors"
)

var (
	errNoSubscriber   = errors.New("no subscribers found")
	errTimeoutToWrite = errors.New("timeout to write into queue")
)
