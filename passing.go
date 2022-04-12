package main

import (
	"fmt"
)

func main() {
	name := "Arza Henrie"
	course := "Go Fundamentals"

	fmt.Println("\nHi", name, "your current course is", course)
	updateCourse(&course)

	fmt.Println("You are currently watching", course)

}

func updateCourse(course *string) string {
	*course = "Getting Started with Docker"
	fmt.Println("Updated course to", *course)
	return *course
}
