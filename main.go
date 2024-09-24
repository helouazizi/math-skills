package main

import (
	"fmt"
	"math-skills/functions"
)



func main(){

// our data 
data := functions.ReadFile("data.txt")
nums := functions.Parse_Float(data)
avrg := functions.Average(nums)
median := functions.Median(nums)
variance := functions.Variance(nums,avrg)
standar_devision := functions.Standar_Devision(variance)
// print our results
fmt.Println("Average: ",avrg)
fmt.Println("Median: ",median)
fmt.Println("Variance: ",variance)
fmt.Println("Standard Deviation: ",standar_devision)


}