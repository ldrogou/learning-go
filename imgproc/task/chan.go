package task

import (
	"fmt"
	"path"
	"path/filepath"

	"training.go/imgproc/filter"
)

//ChanTask structure des channel asynchrone
type ChanTask struct {
	dirCtx
	Filter   filter.Filter
	PoolSize int
}

func NewChanTask(srcDir, dstDir string, filter filter.Filter, poolSize int) Tasker {

	return &ChanTask{
		Filter: filter,
		dirCtx: dirCtx{
			SrcDir: srcDir,
			DstDir: dstDir,
			files:  buildFilesList(srcDir),
		},
		PoolSize: poolSize,
	}
}

type jobReq struct {
	src string
	dst string
}

//Process fonction Process
func (c *ChanTask) Process() error {
	size := len(c.files)
	jobs := make(chan jobReq, size)
	results := make(chan string, size)

	for w := 1; w <= c.PoolSize; w++ {
		go worker(w, c, jobs, results)
	}

	// strat
	for _, f := range c.files {
		filename := filepath.Base(f)
		d := path.Join(c.DstDir, filename)

		jobs <- jobReq{src: f,
			dst: d}
	}
	close(jobs)

	for range c.files {
		fmt.Printf(<-results)
	}
	return nil

}

func worker(id int, chanTask *ChanTask, jobs <-chan jobReq, result chan<- string) {

	for j := range jobs {
		fmt.Printf("worker %d, start job %v \n", id, j)
		chanTask.Filter.Process(j.src, j.dst)
		fmt.Printf("worker %d, finished job %v \n", id, j)
		result <- j.dst

	}
}
