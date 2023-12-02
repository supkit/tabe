package config

import (
	"fmt"
	"testing"
)

func TestWatch(t *testing.T) {
	err := Watch("tabe.yaml")
	if err != nil {
		fmt.Printf("config error: %v\n", err)
		return
	}
	fmt.Printf("config: %+v\n", config)
}
