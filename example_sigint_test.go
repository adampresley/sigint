package sigint_test

import (
	"log"
	"os"

	"github.com/adampresley/sigint"
)

/*
This example sets up a SIGINT handler that simply
exits the application.
*/
func ExampleListenForSIGINT() {
	sigint.ListenForSIGINT(func() {
		log.Println("Shutting down...")
		os.Exit(0)
	})
}