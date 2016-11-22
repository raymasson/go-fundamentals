package main

import (
	"fmt"
)

func main() {
	leagueTitles := make(map[string]int)
	leagueTitles["Sunderland"] = 6
	leagueTitles["Newcastle"] = 4

	recentHead2Head := map[string]int{
		"Sunderland": 5,
		"Newcastle":  0,
	}

	fmt.Printf("\nLeague titles: %v\nRecent head to heads: %v\n", leagueTitles, recentHead2Head)

	testMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}

	/*for key, value := range testMap {
		fmt.Printf("\nKey is: %v Value is: %v", key, value)
	}*/

	testMap["A"] = 100
	testMap["F"] = 1973
	fmt.Println(testMap)
	fmt.Println(testMap["A"])

	delete(testMap, "F")
	fmt.Println(testMap)
}
