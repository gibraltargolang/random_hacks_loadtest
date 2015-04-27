/**
 * @author: thomasmodeneis
 */

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {

	//configure your target
	//1) nodejs 0.10.x -> https://desolate-harbor-2201.herokuapp.com/hello
	//2) golang 1.3 https://limitless-basin-5531.herokuapp.com/hello
	//3) java vertex https://vertx-simple-json-endpoint.herokuapp.com/hello
	//4) play scala https://sheltered-tor-1754.herokuapp.com/hello
	//5) golang 1.2.1 https://gentle-citadel-4686.herokuapp.com/hello (offline)

	urlTarget := "https://desolate-harbor-2201.herokuapp.com/hello"

	//configure if POST or GET maybe?
	m := "POST"

	rate := uint64(100)         // per second
	duration := 4 * time.Second //4 seconds

	//read json from file
	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	//configure your target
	targeter := vegeta.NewStaticTargeter(&vegeta.Target{
		Method: m,
		URL:    urlTarget,
		Body:   file,
	})
	attacker := vegeta.NewAttacker()

	// define results
	var results vegeta.Results
	for res := range attacker.Attack(targeter, rate, duration) {
		results = append(results, res)
	}

	//configure to generate a json report
	rep := vegeta.ReportJSON
	jsondata, _ := rep.Report(results)
	err := ioutil.WriteFile("./results/results.json", jsondata, 0644)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(jsondata))

	//configure to generate a html report
	rep = vegeta.ReportPlot
	htmldata, _ := rep.Report(results)
	err = ioutil.WriteFile("./results/results.html", htmldata, 0644)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	//configure to generate a text report
	rep = vegeta.ReportText
	textdata, _ := rep.Report(results)
	err = ioutil.WriteFile("./results/results.txt", textdata, 0644)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

}
