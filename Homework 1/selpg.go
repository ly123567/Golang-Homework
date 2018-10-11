package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"

	flag "github.com/spf13/pflag"
)

type Params struct {
	start      int
	end        int
	length     int
	filename   string
	outputFile string
	pageType   bool
	help       bool
}

var param Params

func init() {
	flag.IntVarP(&param.start, "start", "s", -1, "Start page of showing pages")
	flag.IntVarP(&param.end, "end", "e", -1, "End page of showing pages")
	flag.IntVarP(&param.length, "length", "l", 72, "The line-length of each pages")
	flag.StringVarP(&param.outputFile, "dest", "d", "", "Select the output file path")
	flag.BoolVarP(&param.pageType, "type", "f", false, "Divide pages by page break")
	flag.BoolVarP(&param.help, "help", "h", false, "Show this message")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: ./selpg [-s start] [-e end] [optional...] [filepath]\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()
}

func check() {
	if param.help || param.start == -1 || param.end == -1 || param.start > param.end {
		flag.Usage()
	}

	if param.start == -1 {
		fmt.Fprintf(os.Stderr, "\nMissing necessary argument: -s\n")
		os.Exit(1)
	}

	if param.end == -1 {
		fmt.Fprintf(os.Stderr, "\nMissing necessary argument: -e\n")
		os.Exit(1)
	}

	if param.start > param.end {
		fmt.Fprintf(os.Stderr, "Start page should less than end page!\n")
		os.Exit(2)
	}

	if param.length < 1 {
		fmt.Fprintf(os.Stderr, "The line number can not less than 1!\n")
		os.Exit(3)
	}

	if flag.NArg() > 0 { //Non-flag parameter number
		param.filename = flag.Arg(0)
		_, err := os.Stat(param.filename) //check if file exist
		if err != nil {
			fmt.Fprintf(os.Stderr, "Please input filepath correctly!\n")
			os.Exit(4)
		}
	}
}

func process() string {
	lineCount := 0
	pageCount := 1

	input := os.Stdin
	response := ""

	if param.filename != "" {
		err := errors.New("")
		input, err = os.Open(param.filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Open file failed!\n")
			os.Exit(5)
		}
		defer input.Close()
	}

	readLine := bufio.NewReader(input)
	if !param.pageType { // -l type
		for {
			line, err := readLine.ReadString('\n')
			if err == io.EOF {
				response += line
				break
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "Read file error!\n")
				os.Exit(6)
			}
			lineCount++
			if lineCount > param.length {
				pageCount++
				lineCount = 1
			}
			if pageCount >= param.start && pageCount <= param.end {
				response += line
			}
			if pageCount > param.end {
				break
			}
		}
	} else { // -f type
		for {
			page, err := readLine.ReadString('\f')
			if err == io.EOF {
				response += page
				break
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "Read file error!\n")
				os.Exit(6)
			}
			if pageCount >= param.start && pageCount <= param.end {
				response += page
			}
			pageCount++

			if pageCount > param.end {
				break
			}
		}
	}

	return response
}

func output(response string) {
	if param.outputFile != "" { // -d is not null
		cmd := exec.Command("/usr/bin/lp", fmt.Sprintf("-d%s", param.outputFile))
		reader, _, err := os.Pipe()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Create pipe error! \n")
			os.Exit(7)
		}
		cmd.Stdin = reader
		cmd.Run()
	}
	fmt.Printf(response)
}

func main() {
	check()
	output(process())
}
