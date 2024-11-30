package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		sum   float64
		years int64
		rate  int64
	)

	fmt.Println("Enter sum:")
	fmt.Scan(&sum)
	fmt.Println("Enter years:")
	fmt.Scan(&years)
	fmt.Println("Enter rate:")
	fmt.Scan(&rate)

	if sum < 100 || sum > 1000000 {
		fmt.Println("Incorrect sum")
		return
	}
	if years < 1 || years > 100 {
		fmt.Println("Incorrect years")
		return
	}
	if rate < 1 || rate > 20 {
		fmt.Println("Incorrect rate")
		return
	}

	result := sum * math.Pow(1+float64(rate)/100, float64(years))
	fmt.Printf("Total result is %.2f\n", result)
}
