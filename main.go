package main

import (
    //"fmt"
	"github.com/briandowns/spinner"
	"time"
)

func main() {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)  // Build our new spinner
	s.Start()                                                    // Start the spinner
    s.Prefix = "Running "
    s.Color("green", "bold")
	time.Sleep(4 * time.Second)                                  // Run for some time to simulate work
	s.Stop()
}
