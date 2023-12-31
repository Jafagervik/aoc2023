package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	input,_ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	maxCalories := 0
	currentElfCalories := 0

	for sc.Scan(){
		snack, err := strconv.Atoi(sc.Text())
		currentElfCalories += snack

		if err != nil {
			if currentElfCalories > maxCalories{
				maxCalories = currentElfCalories
			}
			currentElfCalories = 0
		}
	}
	fmt.Println(maxCalories)
}
