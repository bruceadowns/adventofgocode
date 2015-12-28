package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

func countArray(in []interface{}) (res float64) {
	for _, v := range in {
		switch v := v.(type) {
		case string:
		case float64:
			res += v
		case []interface{}:
			res += countArray(v)
		case map[string]interface{}:
			res += countObject(v)
		default:
			log.Fatalf("unknown type")
		}
	}

	return
}

func countObject(in map[string]interface{}) (res float64) {
	for _, v := range in {
		switch v := v.(type) {
		case string:
			if v == "red" {
				return 0
			}
		case float64:
			res += v
		case []interface{}:
			res += countArray(v)
		case map[string]interface{}:
			res += countObject(v)
		default:
			log.Fatalf("unknown type")
		}
	}

	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// scan for json array
		{
			var in []interface{}
			if err := json.Unmarshal(scanner.Bytes(), &in); err == nil {
				//log.Printf("sum of numbers in array %.0f [%s]", countArray(in), scanner.Text())
				log.Printf("sum of numbers in array %.0f", countArray(in))
				continue
			}
		}

		// scan for json object
		{
			var in map[string]interface{}
			if err := json.Unmarshal(scanner.Bytes(), &in); err == nil {
				//log.Printf("sum of numbers in object %.0f [%s]", countObject(in), scanner.Text())
				log.Printf("sum of numbers in object %.0f", countObject(in))
				continue
			}
		}

		log.Fatalf("invalid input %s", scanner.Text())
	}
}
