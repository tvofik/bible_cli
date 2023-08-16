package main

import "github.com/tvofik/biblecli/cmd"
import "time"
import "fmt"

func main() {
	start := time.Now()
	cmd.Execute()
	elapsed := time.Since(start)
	fmt.Printf("Mine took %s", elapsed)
}
