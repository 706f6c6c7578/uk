package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		fmt.Println("Usage: uk < infile > outfile")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	replacer := strings.NewReplacer(
		"Ä", "AE", "Ö", "OE", "Ü", "UE", "ẞ", "SS",
		"ä", "ae", "ö", "oe", "ü", "ue", "ß", "ss",
	)

	var input strings.Builder
	for {
		part, err := reader.ReadString('\n')
		input.WriteString(part)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			return
		}
	}

	output := replacer.Replace(input.String())
	writer.WriteString(output)
	writer.Flush()
}
