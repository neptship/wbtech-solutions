package main

import "fmt"

func main() {
	temprature := []float32{-25.4, -27.0, 13.0, 15.5, 24.5, -21.0, 32.5}
	tempSort := make(map[int][]float32)
	for _, temp := range temprature {
		tempT := int(temp) - int(temp)%10
		tempSort[tempT] = append(tempSort[tempT], temp)
	}

	fmt.Println(tempSort)
}
