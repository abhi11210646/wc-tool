package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Stats struct {
	lines int
	words int
	bytes int
}

type WC struct {
	lineFlag  bool
	wordFlag  bool
	byteFlag  bool
	flagCount int
	filePath  []string
	stats     []Stats
}

func (wc *WC) fileStats(reader io.Reader) Stats {
	scanner := bufio.NewScanner(reader)
	stats := Stats{lines: -1, bytes: -1} // new line count adjustment
	for scanner.Scan() {
		text := scanner.Text()
		if wc.byteFlag {
			stats.bytes += len(text) + 1
		}
		if wc.wordFlag {
			stats.words += len(strings.Fields(text))
		}
		if wc.lineFlag {
			stats.lines++
		}
	}
	wc.stats = append(wc.stats, stats)
	return stats
}

func main() {
	wc := WC{stats: []Stats{}}
	wc.loadArgs()

	if len(wc.filePath) > 0 {
		for i, file := range wc.filePath {
			reader, err := os.Open(file)
			if err != nil {
				panic(err)
			}
			defer reader.Close()
			stats := wc.fileStats(reader)
			wc.printStats(stats)
			fmt.Printf(" %s\n", wc.filePath[i])
		}
		// if len(wc.stats) > 1 {
		// fmt.Printf("total")
		// }
	} else {
		reader := bufio.NewReader(os.Stdin)
		stats := wc.fileStats(reader)
		wc.printStats(stats)
		fmt.Println()
	}
}
