package gokey

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	client := newClient()

	_, err := client.cache.SAdd("some key", []any{1, 2, 3, 4, 5})
	fmt.Printf("%v", err)
}
