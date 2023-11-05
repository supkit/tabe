package config

import (
	"fmt"
	"testing"
)

func TestWatch(t *testing.T) {
	config, err := Watch()
	if err != nil {
		fmt.Printf("config error: %v\n", err)
		return
	}
	fmt.Printf("config: %+v\n", config)
}
