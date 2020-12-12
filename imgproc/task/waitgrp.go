package task

import (
	"fmt"
	"path"
	"path/filepath"
	"sync"

	"training.go/imgproc/filter"
)

//WaitGrpTask struct
type WaitGrpTask struct {
	dirCtx
	Filter filter.Filter
}

//NewWaitGrpTask nouvelle instance
func NewWaitGrpTask(srcDir, dstDir string, filter filter.Filter) Tasker {

	return &WaitGrpTask{
		Filter: filter,
		dirCtx: dirCtx{
			SrcDir: srcDir,
			DstDir: dstDir,
			files:  buildFilesList(srcDir),
		},
	}
}

//Process implémentation de Tasker
func (w *WaitGrpTask) Process() error {

	var wg sync.WaitGroup
	size := len(w.files)
	for i, f := range w.files {
		filename := filepath.Base(f)
		dst := path.Join(w.DstDir, filename)
		wg.Add(1)
		go w.applyFilter(f, dst, &wg, i+1, size)
	}
	wg.Wait()
	fmt.Println("Traitement des fichiers terminés")
	return nil
}

func (w *WaitGrpTask) applyFilter(src, dst string, wg *sync.WaitGroup, i, size int) {
	w.Filter.Process(src, dst)
	fmt.Printf("Traitement [%d/%d] %v => %v \n", i, size, src, dst)
	wg.Done()

}
