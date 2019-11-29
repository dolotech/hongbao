// Copyright 2012-2016 Apcera Inc. All rights reserved.

package signal

import (
	"os"
	"os/signal"
)

// Signal Handling
func HandleSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			os.Exit(0)
		}
	}()
}
