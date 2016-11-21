package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	/*firstRank := 39
	secondRank := 614

	if firstRank < secondRank {
		fmt.Println("\nFirst course is doing better" +
			" than second course")
	} else if firstRank > secondRank {
		fmt.Println("\noh dear.... your first course must be doing abysmally")
	} else {
		fmt.Println("\nBoth courses are either the same or something weird is going on")
	}*/

	if firstRank, secondRank := 120, 614; firstRank < secondRank {
		fmt.Println("\nFirst course is doing better" +
			" than second course")
		if firstRank > 100 {
			fmt.Println("Errr, you may wanna consider another job :-D")
		}
	} else if firstRank > secondRank {
		fmt.Println("\noh dear.... your first course must be doing abysmally")
		if secondRank > 100 {
			fmt.Println("Errr, you may wanna consider another job :-D")
		}
	} else {
		fmt.Println("\nBoth courses are either the same or something weird is going on")
	}

	/*switch "docker" {
	case "linux":
		fmt.Println("\nHere are some recommended Linux courses...")
	case "docker":
		fmt.Println("\nHere are some recommended Docker courses...")
		fallthrough
	case "windows":
		fmt.Println("\nHere are some recommended Windows courses...")
	default:
		fmt.Println("\nSorry, we couldn't find a match, why not try our top 100 list!")
	}*/

	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("We got an even number -", tmpNum)
	case 1, 3, 5, 7, 9:
		fmt.Println("We got an odd number -", tmpNum)
	}

	_, err := os.Open("C:\\Temp\\logstash-file.log")

	if err != nil {
		fmt.Println("Error returned was:", err)
	} else {
		fmt.Println("No Error opening the file")
	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
