package main

import (
	"bufio"
	"fmt"
	"mime"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	decoder := mime.WordDecoder{}

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		decodedLine, err := decoder.DecodeHeader(line)

		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("%s\n", err.Error()))
			os.Exit(1)
		}

		fmt.Print(decodedLine)
	}
}
