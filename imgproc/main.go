package main

import (
	"flag"
	"fmt"
	"time"

	"training.go/imgproc/filter"
	"training.go/imgproc/task"
)

func main() {

	var srcDir = flag.String("src", "", "Input Directory")
	var dstDir = flag.String("dst", "", "Output Directory")
	var filterType = flag.String("filter", "grayscale", "grayscale/blur")
	flag.Parse()

	var f filter.Filter
	switch *filterType {
	case "grayscale":
		f = filter.Grayscale{}
	case "blur":
		f = filter.Blur{}
	}

	t := task.NewChanTask(*srcDir, *dstDir, f, 16)

	start := time.Now()
	t.Process()
	elapsed := time.Since(start)

	fmt.Printf("Le traitement à pris %v  secondes \n", elapsed)

}
