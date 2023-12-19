package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type WC struct {
	lineFlag bool
	wordFlag bool
	charFlag bool
	filePath []string
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
		case "-w":
			wc.wordFlag = true
		case "-m":
			wc.charFlag = true
		default:
			wc.filePath = append(wc.filePath, args[i])
		}
	}
}

func main() {
	wc := WC{}
	wc.loadArgs()

	if len(wc.filePath) > 0 {
		total := 0
		for _, file := range wc.filePath {
			reader, _ := os.Open(file)
			defer reader.Close()
			scanner := bufio.NewScanner(reader)
			l := 0
			w := 0
			c := 0
			for scanner.Scan() {
				text := scanner.Text()
				c += len(text)
				w += len(strings.Fields(text))
				l++
			}
			total += l
			fmt.Println(l, w, c, file)
		}
		if len(wc.filePath) > 1 {
			fmt.Println(total, "total")
		}
	} else {
		r := bufio.NewReader(os.Stdin)
		scanner := bufio.NewScanner(r)
		l := 0
		w := 0
		c := 0
		for scanner.Scan() {
			text := scanner.Text()
			c += len(text)
			w += len(strings.Fields(text))
			l++
		}
		fmt.Println(l, w, c)
	}
}
