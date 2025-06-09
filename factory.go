package logger

import (
	"os"
	"strconv"
	"strings"

	"ebizzone.com/logger/console"
	"ebizzone.com/logger/log"
)

func New(name string) log.Logger {
	_, err := os.Stat(".logger.conf")
	if err == nil {
		// if config file exists
		logger := readConfigFile(name)
		return logger
	} else {
		return console.NewLogger(name)
	}
}

func readConfigFile(name string) log.Logger {
	data, err := os.ReadFile(".logger.conf")
	var loggerType = "console"
	if err == nil {
		lines := strings.Split(string(data), "\n")
		for l := range lines {
			line := strings.TrimSpace(lines[l])

			// if this is comment line, ignore
			if strings.HasPrefix(line, "#") {
				continue
			}
			// if this is empty line, ignore
			if line == "" {
				continue
			}
			i := strings.Index(line, "=")
			// if not key specified, ignore
			if i < 0 {
				continue
			} else {
				// fmt.Println(line)
				key := line[0:i]
				value := line[i+1:]

				// parse the config file
				switch key {
				case "LEVEL":
					v, err := strconv.ParseUint(value, 10, 8)
					if err == nil {
						log.CurrentLevel = uint8(v)
					}
				case "TYPE":
					loggerType = value
				}
			}
		}
	}

	if loggerType == "console" {
		return console.NewLogger(name)
	} else {
		// currently this is the only implementation
		return console.NewLogger(name)
	}
}
