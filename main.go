package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"os/exec"
	"strconv"
)

// CountryData holds the data previously known alongside the name of the country. Must be filled in manually
type CountryData struct {
	Name          string
	PreviousCases []float64
	Population    float64
	AverageRate   float64
}

// ResultData is the format of the infection results that will occur in the coming days
type ResultData struct {
	Name     string
	NewCases []float64
}

// averageAroundInfected obtains the average number of people someone infected is exposed to
func averageAroundInfected(data CountryData) (averageRate float64) {
	increaseRatio := data.PreviousCases[0] / data.PreviousCases[1]
	populationRatio := data.PreviousCases[0] / data.Population
	averageRate = (increaseRatio - 1) / (1 - populationRatio)
	return averageRate
}

func predictNextDay(previous float64, data CountryData) float64 {
	p := (1 - previous/data.Population)
	return math.Round(previous * (1 + data.AverageRate*p))
}

func main() {
	// Create results and images folder
	_, err := os.Stat("./results")
	if os.IsNotExist(err) {
		os.MkdirAll("./results", 0755)
	}

	var data []CountryData
	b, _ := ioutil.ReadFile("./data.json")
	json.Unmarshal(b, &data)

	var results []ResultData
	for i, country := range data {
		data[i].AverageRate = averageAroundInfected(country)
		results = append(results, ResultData{
			Name:     country.Name,
			NewCases: []float64{country.PreviousCases[0]},
		})
	}

	for i := 1; i <= 60; i++ {
		dayText := strconv.Itoa(i) + " day(s) from now:\n"

		for j, country := range data {
			newResult := predictNextDay(results[j].NewCases[i-1], country)

			results[j].NewCases = append(results[j].NewCases, newResult)
			dayText += strconv.FormatFloat(newResult/country.Population*100, 'f', 2, 64) + "% infected (" + strconv.FormatFloat(newResult, 'f', 0, 64) + " people) in " + country.Name + "\n"
		}

		fmt.Println(dayText)
	}

	for _, result := range results {
		json, _ := json.Marshal(result.NewCases)
		ioutil.WriteFile("./results/"+result.Name+".json", json, 0644)
	}

	_, err = exec.Command("python", "plot.py").Output()
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Created graph!")
	}
}
