package config

import (
	"fmt"
	"testing"
)

func TestWatch(t *testing.T) {
	config, err := Watch()
	fmt.Println(config)
	fmt.Println(err)
}
