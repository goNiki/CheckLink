package domain

import "sync/atomic"

var isDraining atomic.Bool

func SetDraining(draining bool) {
	isDraining.Store(draining)
}

func IsDraining() bool {
	return isDraining.Load()
}
