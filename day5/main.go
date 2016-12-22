package main

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("door id: %s", line)

		password := bytes.Buffer{}
		idx := 0
		for ; ; idx++ {
			hasher := md5.New()
			hasher.Write([]byte(fmt.Sprintf("%s%d", line, idx)))
			hash := hex.EncodeToString(hasher.Sum(nil))
			//log.Printf("%d hash: %s %c", i, hash, hash[5])

			if 0 == strings.Compare("00000", hash[:5]) {
				password.WriteByte(hash[5])
				//log.Printf("next char: %c", hash[5])
			}

			if password.Len() == 8 {
				break
			}
		}

		log.Printf("password: %s [%d]", password.String(), idx)
	}
}
