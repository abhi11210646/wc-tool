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
			wc.LineFlag = true
			wc.FlagCount++
		case "-w":
			wc.WordFlag = true
			wc.FlagCount++
		case "-c":
			wc.ByteFlag = true
			wc.FlagCount++
		default:
			wc.FilePath = append(wc.FilePath, args[i])
		}
	}
	if wc.FlagCount == 0 { // if no flag then set all
		wc.LineFlag = true
		wc.WordFlag = true
		wc.ByteFlag = true
	}
}

func (wc *WC) printStats() {
	var t_l, t_w, t_b int
	for _, s := range wc.Stats {
		if wc.LineFlag {
			fmt.Printf("%-3d", s.Lines)
			t_l += s.Lines
		}
		if wc.WordFlag {
			fmt.Printf("%-3d", s.Words)
			t_w += s.Words
		}
		if wc.ByteFlag {
			fmt.Printf("%-5d", s.Bytes)
			t_b += s.Bytes
		}
		fmt.Printf("%s\n", s.File)
	}
	if len(wc.Stats) > 1 {
		if t_l != 0 {
			fmt.Printf("%-3d", t_l)
		}
		if t_w != 0 {
			fmt.Printf("%-3d", t_w)
		}
		if t_b != 0 {
			fmt.Printf("%-5d", t_b)
		}
		fmt.Printf("%s\n", "total")
	}
}
