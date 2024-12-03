package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

type aocFunc func([]byte) (int, error)

var funcs = map[string]aocFunc{
	"11": d1p1,
	"12": d1p2,

	"31": d3p1,
	"32": d3p2,
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Please specify a day and part!")
	}
	day := os.Args[1]
	part := os.Args[2]

	log.Println("Reading input...")
	input, err := os.ReadFile("d" + day + ".txt")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Starting execution...")
	funcToExec := funcs[day+part]
	start := time.Now()
	output, err := funcToExec(input)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Execution complete!\nOutput: %s\nExecution time: %s", strconv.FormatInt(int64(output), 10), time.Since(start))
}
