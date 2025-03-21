package main

import (
	"fmt"
	"math"
)
const inflationRate float64 = 2.5

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := getFutureValue(investmentAmount, expectedReturnRate, years)
	futureRealValue := getFutureRealValue(futureValue, years)

	fmt.Printf("Future value is: %v\n", futureValue)
	fmt.Printf("Future real value is: %v\n", futureRealValue)
}

func getFutureValue (investmentAmount float64, expectedReturnRate float64, years float64) float64 {
	return investmentAmount * math.Pow(1+expectedReturnRate/100, years)
}
func getFutureRealValue (futureValue float64, years float64) float64 {
	return futureValue / math.Pow(1+inflationRate/100, years)
}