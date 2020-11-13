package main

import "fmt"
import "bufio"
import "os"
import "log"

import "github.com/bsdpunk/dndshell/shell"

//import "./shell"

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		for scanner.Scan() {
			fmt.Println("Feature to come later")

		}
		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
	} else {
		shell.Shell(os.Args)
	}

}
