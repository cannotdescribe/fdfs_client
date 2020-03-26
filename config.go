package fdfs_client

import (
	"bufio"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type config struct {
	trackerAddr []string
	maxConns    int
}

func NewConfig(configName string) (*config, error) {
	return newConfig(configName)
}

func newConfig(configName string) (*config, error) {
	config := &config{}
	f, err := os.Open(configName)
	if err != nil {
		return nil, err
	}
	splitFlag := "\n"
	if runtime.GOOS == "windows" {
		splitFlag = "\r\n"
	}
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return config, nil
			}
			return nil, err
		}
		line = strings.TrimSuffix(line, splitFlag)
		str := strings.SplitN(line, "=", 2)
		key := strings.TrimSpace(str[0])
		value := strings.TrimSpace(str[1])
		switch key {
		case "tracker_server":
			config.trackerAddr = append(config.trackerAddr, value)
		case "maxConns":
			config.maxConns, err = strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
		}
	}
	return config, nil
}
