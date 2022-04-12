package main

import (
	"fmt"
	"os"
)

func main() {
	//sample constant
	const c = 186000
	fmt.Println("The speed of light is", c, "miles per second.")

	//print all enviroment variables
	// too messy: fmt.Println(os.Environ())

	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}
