package main

import (
	"fmt"
	"time"
)

func main() {
	for timer := 10; timer >= 0; timer-- {
		if timer%2 != 0 {
			continue
		}
		if timer == 0 {
			fmt.Println("BOOM!")
			break
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}

	coursesInProg := []string{
		"Docker Deep Dive",
		"Docker Clustering",
		"Docker and Kubernetes",
	}

	coursesCompleted := []string{
		"Docker Deep Dive",
		"Go Fundamentals",
		"Puppet Fundamentals",
	}

	for _, i := range coursesInProg {
		fmt.Println(i)

		for _, j := range coursesCompleted {
			if i == j {
				fmt.Println("\nHey we found a clash.", i, "is in both lists")
			}
		}
	}
}
