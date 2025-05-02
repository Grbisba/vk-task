package subpub

import (
	"errors"
)

var (
	errNoSubscriber = errors.New("no subscribers found")
)
