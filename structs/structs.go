package main

import (
	"fmt"
)

func main() {
	//var DockerDeepDive courseMeta
	//DockerDeepDive := new(courseMeta)

	DockerDeepDive := courseMeta{
		Author: "Nigel Poulton",
		Level:  "Intermediate",
		Rating: 5,
	}

	fmt.Println(DockerDeepDive)
	fmt.Println("\nDocker Deep Dive author is:", DockerDeepDive.Author)

	DockerDeepDive.Rating = 1

	fmt.Println("\nDocker Deep Dive Rating is:", DockerDeepDive.Rating)

	myCourse := newCourseMeta()

	fmt.Println(myCourse)

	myCourse2 := newCourseMeta2("rmasson", "pro", 2)

	fmt.Println(myCourse2.Author)

	myCourse2.printCourse()

	ecm := enhancedCourseMeta{}
	ecm.courseMeta = courseMeta{
		Author: "Nigel Poulton",
		Level:  "Intermediate",
		Rating: 5,
	}
	ecm.printCourse()

	ecm2 := enhancedCourseMeta{}
	ecm2.Author = "Ray"
	ecm2.Level = "noob"
	ecm2.Rating = 10
	ecm2.printCourse()
}

type courseMeta struct {
	Author string
	Level  string
	Rating float64
}

type enhancedCourseMeta struct {
	courseMeta
}

func newCourseMeta() *courseMeta {
	courseMeta := courseMeta{}
	courseMeta.Author = "Ray"
	courseMeta.Level = "Beginner"
	courseMeta.Rating = 5

	return &courseMeta
}

func newCourseMeta2(author, level string, rating float64) *courseMeta {
	courseMeta := courseMeta{}
	courseMeta.Author = author
	courseMeta.Level = level
	courseMeta.Rating = rating

	return &courseMeta
}

func (cm *courseMeta) printCourse() {
	fmt.Printf("\nCourse author: %s, level: %s, rating: %f", cm.Author, cm.Level, cm.Rating)
}
