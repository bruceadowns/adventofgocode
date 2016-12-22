package main

import (
	"bufio"
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

		var password [8]byte
		for idx := 0; ; idx++ {
			hasher := md5.New()
			hasher.Write([]byte(fmt.Sprintf("%s%d", line, idx)))
			hash := hex.EncodeToString(hasher.Sum(nil))
			//log.Printf("%d hash: %s %c", i, hash, hash[5])

			if 0 == strings.Compare("00000", hash[:5]) {
				offset := hash[5] - 48
				if offset >= 0 && offset <= 7 {
					if password[offset] == 0 {
						password[offset] = hash[6]
					}
				}

				log.Printf("%s %d '%s' [%d]", hash[:7], offset, password, idx)
			}

			count := 0
			for _, b := range password {
				if b != 0 {
					count++
				}
			}
			if count == 8 {
				break
			}
		}

		log.Printf("password: %s", password)
	}
}
