package main

import (
	"fmt"
	"os"
)

func outputHelp() {
	fmt.Println(`
	USAGE:  wc [OPTIONS] FILE(S)
		-c      print the byte counts
		-l      print the newline counts
		-w      print the word counts  
	EXAMPLE: 
		wc -l logs.txt
		wc -w logs.txt
		wc -l logs.txt logs2.txt
	  `)
}

func (wc *WC) loadArgs() {
	args := os.Args[1:]
	for i, a := range args {
		if a == "-h" {
			outputHelp()
			return
		}
		switch a {
		case "-l":
			wc.lineFlag = true
			wc.flagCount++
		case "-w":
			wc.wordFlag = true
			wc.flagCount++
		case "-c":
			wc.byteFlag = true
			wc.flagCount++
		default:
			wc.filePath = append(wc.filePath, args[i])
		}
	}
	if wc.flagCount == 0 { // if no flag then set all
		wc.lineFlag = true
		wc.wordFlag = true
		wc.byteFlag = true
	}
}

func (wc *WC) printStats(stats Stats) {
	if wc.lineFlag {
		fmt.Printf("%2d", stats.lines)
	}
	if wc.wordFlag {
		fmt.Printf("%3d", stats.words)
	}
	if wc.byteFlag {
		fmt.Printf("%3d", stats.bytes)
	}
}
