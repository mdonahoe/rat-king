package main

import (
	"bufio"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/google/shlex"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

// hold the command strings
type Commands []string

func parseCommandsFromFile(f *os.File) Commands {
	// Prep an array of command strings
	cmds := make(Commands, 0, 10)

	// Wrapping the unbuffered `*os.File` with a buffered reader gives us
	// the `ReadString` method which will let us read entire lines.
	rdr := bufio.NewReader(f)

	// `ReadString` returns the next string from the
	// input up to the given separator byte.
	for {
		// Switch on the err after reading a single line from the input file.
		switch line, err := rdr.ReadString('\n'); err {

		// The err is nil, so the read succeeded.
		case nil:
			cmds = append(cmds, line)

		// End of input.
		case io.EOF:
			return cmds

		// There was a problem!
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	}
}

func main() {
	cmds := parseCommandsFromFile(os.Stdin)
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond) // Build our new spinner
	for _, line := range cmds {
		trimmedLine := strings.TrimRight(line, "\n")
		s.Suffix = " " + trimmedLine
		s.Color("green", "bold")
		s.Start()

		// parse and run the cmd, printing output
		cmdline, err3 := shlex.Split(line)
		if err3 != nil {
			log.Fatal(err3)
		}
		cmd := exec.Command(cmdline[0], cmdline[1:]...)
		out, err := cmd.Output()
		s.Stop()
		if err != nil {
			d := color.New(color.FgRed, color.Bold)
			d.Printf("%s\n", trimmedLine)
			fmt.Printf("%s\n\n", err)
		} else {
			d := color.New(color.FgGreen, color.Bold)
			d.Printf("%s\n", trimmedLine)
			fmt.Printf("%s\n", out)
		}
	}
}
