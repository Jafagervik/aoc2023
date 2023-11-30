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

	for sc.Scan(){
		data, err := strconv.Atoi(sc.Text())

		if err != nil {
		}
	}
	fmt.Println(maxCalories)
}
