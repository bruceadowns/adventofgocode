package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var totalUnescaped, totalEscaped, totalCode int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		totalCode += len(line)

		unescapedLine, err := strconv.Unquote(line)
		if err != nil {
			log.Fatalf("invalid input %s", line)
		}
		totalUnescaped += len(unescapedLine)

		escapedLine := strconv.Quote(line)
		totalEscaped += len(escapedLine)
	}

	fmt.Printf("decode code-mem=%d [%d-%d]\n", totalCode-totalUnescaped, totalCode, totalUnescaped)
	fmt.Printf("re-encode code-mem=%d [%d-%d]\n", totalEscaped-totalCode, totalCode, totalEscaped)
}
