package wc

import (
	"bufio"
	"io"
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

func NewWordCount() *WC {
	return &WC{Stats: []Stats{}}
}

func (wc *WC) CalculateFileStats(reader io.Reader, file string) {
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
