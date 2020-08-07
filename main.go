package main

import (
	"fmt"
)

// NOTE: cannot just use channels without go routine  error: all goroutines are asleep - deadlock!

func main() {
	ore := make(chan string)
	smelt := make(chan string)
	doneChan := make(chan string)
	mine := [6]string{"diamond", "ruby", "gold", "silver", "coal", "platinum"}

	go func(mine [6]string) { // go thread that mines
		for _, v := range mine {
			fmt.Println("mined ", v)
			ore <- v
		}
		doneChan <- "mining successful"
	}(mine)

	go func() { // go thread breaks mined ore
		for breakedOre := range ore {
			// breakedOre := <-ore
			fmt.Println("breaking ore ", breakedOre)
			smelt <- breakedOre
		}
		doneChan <- "breaking ore successful"
	}()

	go func() { // go thread smelts mined ore
		for smeltedOre := range smelt {
			// smeltedOre := <-smelt
			fmt.Println("smelting ore ", smeltedOre)
		}
		doneChan <- "smelting ore successful"
	}()

	<-doneChan // blocks until all go routines have completed there tasks successfully
}
