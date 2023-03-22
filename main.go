package main

import (
	"fmt"
	"io/ioutil"
	functions "math-skills/functions"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := os.Args[1]
	dataInFile := ReadFile(fileName)
	tabDataInFile := TabDataInFile(dataInFile)
	Moyenne := (functions.Average(tabDataInFile))
	Mediane := (functions.Median(tabDataInFile))
	Variance := (functions.Variance(tabDataInFile))
	EcartType := (functions.StandardDeviation(tabDataInFile))

	fmt.Printf("Average: %0.0f\n", float64(Moyenne))
	fmt.Printf("Median: %0.0f\n", float64(Mediane))
	fmt.Printf("Variance: %0.0f\n", float64(Variance))
	fmt.Printf("Standard Deviation: %0.0f\n", float64(EcartType))

}

func TabDataInFile(dataInFile string) []float64 {
	tabDataInFile := strings.Split(dataInFile, "\n")
	var tabDataFloat []float64
	var num float64

	for _, value := range tabDataInFile {
		if value != "" {
			num, _ = strconv.ParseFloat(value, 64)
			tabDataFloat = append(tabDataFloat, num)
		}
	}
	return tabDataFloat
}

func ReadFile(fileName string) string {
	dataInFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(dataInFile)
}
