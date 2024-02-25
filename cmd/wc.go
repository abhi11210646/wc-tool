package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Stats struct {
	File  string
	Lines int
	Words int
	Bytes int
}

type WC struct {
	LineFlag  bool
	WordFlag  bool
	ByteFlag  bool
	FlagCount int
	FilePath  []string
	Stats     []Stats
}

func (wc *WC) calculateFileStats(reader io.Reader, file string) {
	scanner := bufio.NewScanner(reader)
	fileStats := Stats{Lines: -1, Bytes: -1, File: file} // new line count adjustment
	for scanner.Scan() {
		text := scanner.Text()
		fileStats.Bytes += len(text) + 1
		fileStats.Words += len(strings.Fields(text))
		fileStats.Lines++
	}
	wc.Stats = append(wc.Stats, fileStats)
}

func main() {
	wc := WC{Stats: []Stats{}}
	wc.loadArgs()

	if len(wc.FilePath) > 0 {
		for _, file := range wc.FilePath {
			reader, err := os.Open(file)
			if err != nil {
				fmt.Printf("wc: %s: No such file or directory\n", file)
				continue
			}
			defer reader.Close()
			wc.calculateFileStats(reader, file)
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		wc.calculateFileStats(reader, "")
	}
	wc.printStats()
}
