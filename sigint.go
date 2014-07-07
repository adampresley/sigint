// Copyright 2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package sigint

import (
	"os"
	"os/signal"
)

func ListenForSIGINT(handler func()) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	go func() {
		for _ = range done {
			handler()
			return
		}
	}()
}