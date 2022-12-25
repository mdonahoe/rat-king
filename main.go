package main

import (
	"fmt"
	"github.com/briandowns/spinner"
	"time"
)

func main() {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond) // Build our new spinner
	s.Start()                                                    // Start the spinner
	s.Prefix = "Lexing "
	s.Color("green", "bold")
	time.Sleep(1 * time.Second) // Run for some time to simulate work
	s.Stop()
	fmt.Println("\rLexed.")
	s.Color("cyan")
	s.Prefix = "Compiling "
	s.Start()
	time.Sleep(1 * time.Second) // Run for some time to simulate work
	s.Stop()
	fmt.Println("Compiled.")
	s.Prefix = "Linking "
	s.Color("red", "bold")
	s.Start()
	time.Sleep(1 * time.Second) // Run for some time to simulate work
	s.Stop()
	fmt.Println("Linked.")
	s.Prefix = "Installing "
	s.Color("magenta", "bold")
	s.Start()
	time.Sleep(1 * time.Second) // Run for some time to simulate work
	s.Stop()
	fmt.Println("Installed!")
}
