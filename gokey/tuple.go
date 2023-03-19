package gokey

import "time"

type tuple struct {
	ttl       time.Duration
	createdAt time.Time
	value     []byte
}

func newTuple() {

}
