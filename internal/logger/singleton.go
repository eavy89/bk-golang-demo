package logger

import (
	"backend-go-demo/internal/config"
	"fmt"
	"sync"
)

var (
	instance Logger
	once     sync.Once
)

// Init initializes the singleton logger
func Init(config LogConfig) error {
	var err error
	once.Do(func() {
		instance, err = New(config) // by changing the instance we can change the logger implementation
	})
	return err
}

// Get returns the singleton logger instance
func Get() Logger {
	if instance == nil {
		output := config.GetLogFilename()
		devel := false
		if output == "" {
			output = StdOut
			devel = true
		}

		err := Init(LogConfig{
			Level:       Debug,
			OutputPath:  output,
			JSONFormat:  false,
			Development: devel, // Add this flag for better console output
		})

		if err != nil {
			panic(fmt.Sprintf("failed to initialize logger: %v", err))
		}

	}
	return instance
}
