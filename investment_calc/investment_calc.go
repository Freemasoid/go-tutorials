package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	investmentAmount := 1000.0
	expectedReturnRate := 5.5
	years := 10.0

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calcInvestments(investmentAmount, expectedReturnRate, years)

	//futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

func calcInvestments(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	futVal := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	realFutVal := futVal / math.Pow(1+inflationRate/100, years)
	return futVal, realFutVal
}
