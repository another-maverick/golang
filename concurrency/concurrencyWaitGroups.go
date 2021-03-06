package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var i int = -1
	var file string
	for i, file = range os.Args[1:] {
		wg.Add(1)
		go func(filename string) {
			compress(filename)
			wg.Done()
		}(file)
	}
	wg.Wait()
	fmt.Println("compressed files", i+1)

}

func compress(filename string) error {

	in, err := os.Open(filename)
	if err != nil {
		fmt.Println("error is ", err)
	}
	defer in.Close()

	out, err := os.Create(filename + ".gz")
	if err != nil {
		fmt.Println("unable to create the gz file ", err)
	}
	defer out.Close()

	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	gzout.Close()
	return err
}
