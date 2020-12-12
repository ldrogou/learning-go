package main

import (
	"flag"
	"fmt"
	"time"

	"training.go/imgproc/filter"
	"training.go/imgproc/task"
)

func main() {

	flags
	var f filter.Filter = filter.Grayscale{}
	//t := task.NewWaitGrpTask("./input", "output", f)
	t := task.NewChanTask("./input", "output", f, 16)

	start := time.Now()
	t.Process()
	elapsed := time.Since(start)

	fmt.Printf("Le traitement à pris %v  secondes \n", elapsed)

}
