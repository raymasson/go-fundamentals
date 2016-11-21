//Define this file as executable program
package main

//Import some packages
import (
	"fmt"
	"runtime"
)

//Entry point
func main() {
	fmt.Println("Hello from", runtime.GOOS)
}
