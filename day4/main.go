package main

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("input %s\n", line)

		var found5, found6 bool
		for idx := 1; ; idx++ {
			buf := bytes.NewBufferString(scanner.Text())
			buf.WriteString(strconv.Itoa(idx))

			sum := fmt.Sprintf("%x", md5.Sum(buf.Bytes()))

			if !found5 && sum[:5] == "00000" {
				fmt.Printf("%d begins with at least 5 zeros [%s]\n", idx, sum)
				found5 = true
			}

			if !found6 && sum[:6] == "000000" {
				fmt.Printf("%d begins with at least 6 zeros [%s]\n", idx, sum)
				found6 = true
			}

			if found5 && found6 {
				break
			}
		}
	}
}
