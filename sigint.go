// Copyright 2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

/*
Package sigint provides a method to handle CTRL+C (SIGINT) in
command line applications. This is useful for server style
applications, like web and TCP servers.

This package is provided under the MIT license. Please see
the LICENSE file for more information.
*/
package sigint

import (
	"os"
	"os/signal"
)

/*
Sets up a listener to execute a function when the SIGINT
event is fired (usually CTRL+C). The callback function
has no arguments and does not need to return anything.
*/
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