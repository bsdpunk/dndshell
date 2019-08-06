package main

import "fmt"
import "bufio"
import "os"
import "log"
import "./shell"

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
	} else {
		shell.Shell()
	}

}
