package main

import (
	"bufio"
	"fmt"
	"os"

	wc "github.com/abhi11210646/wc-tool/internal"
)

func main() {
	wc := wc.NewWordCount()
	wc.LoadArgs()

	if len(wc.FilePath) > 0 {
		for _, file := range wc.FilePath {
			reader, err := os.Open(file)
			if err != nil {
				fmt.Printf("wc: %s: No such file or directory\n", file)
				continue
			}
			defer reader.Close()
			wc.CalculateFileStats(reader, file)
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		wc.CalculateFileStats(reader, "")
	}
	wc.PrintStats()
}
