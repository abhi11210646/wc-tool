package main

import (
	"bufio"
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
	Stats := Stats{Lines: -1, Bytes: -1, File: file} // new line count adjustment
	for scanner.Scan() {
		text := scanner.Text()
		Stats.Bytes += len(text) + 1
		Stats.Words += len(strings.Fields(text))
		Stats.Lines++
	}
	wc.Stats = append(wc.Stats, Stats)
}

func main() {
	wc := WC{Stats: []Stats{}}
	wc.loadArgs()

	if len(wc.FilePath) > 0 {
		for _, file := range wc.FilePath {
			reader, err := os.Open(file)
			if err != nil {
				panic(err)
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
