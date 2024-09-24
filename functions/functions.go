package functions

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// this function for rading files

func ReadFile(filename string) []string {
	Content, err := os.ReadFile(filename)
	// HANDLE ERR
	if err != nil {
		fmt.Println("ereur open file : ", err)
		os.Exit(1)
	}
	arr := strings.Fields(string(Content))

	return arr
}

// this function to parse int

func Parse_Float(text []string) []float64 {
	numbers := []float64{}
	for _, num := range text {
		flaot, err := strconv.ParseFloat(num, 64)
		if err != nil {
			fmt.Println("ereur : ", err)
		}
		numbers = append(numbers, flaot)
	}
	return numbers
}

// this function to calculate the average 

func Average(numbs []float64) int {
	sum := 0.0
	n := len(numbs)
	for _, num := range numbs {
		sum += (num)
	}
	average := sum/float64(n)

	average = math.Round(average)
	return int(average)
}

// this function to calculate the median
func Median(numbs []float64) int {
	// sort our array
	sort.Float64s(numbs)
	median := 0.0
	n := len(numbs)
	// check for median
	if n % 2 == 0 {
		median = (numbs[n/2-1] + numbs[n/2]) /2
		median = math.Round(median)

	}else{
		median = math.Round(numbs[n/2])
		
	}

	return int(median)
}

// this function to calculate variance
func Variance(numbs []float64,avr int) int {
	
	variance := 0.0
	n := len(numbs)
	for _, num := range numbs {
		variance += (num-float64(avr)) * (num-float64(avr))
	}

	// devid for the length
	variance = variance / float64(n)
	variance = math.Round(variance)
	return int(variance)
}

// standard devitions
func Standar_Devision(variance int) int {
	
	standar_devision := math.Sqrt(float64(variance))
	standar_devision = math.Round(standar_devision)
	return int(standar_devision)
}
