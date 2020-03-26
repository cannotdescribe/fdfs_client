package fdfs_client

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	config, err := NewConfig("fdfs.conf")
	if err != nil {
		fmt.Println(err)
		return;
	}
	fmt.Println(config.trackerAddr)
	fmt.Println(config.maxConns)
}