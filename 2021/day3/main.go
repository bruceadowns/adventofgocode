package main

import (
	"bufio"
	"io"
	"log"
	"math"
	"os"
)

type diagnostic [12]int
type Report []diagnostic

func Part1(in Report) int {
	var gammaRate int
	var epsilonRate int

	for i := 0; i < 12; i++ {
		var zeros int
		var ones int
		for _, diag := range in {
			switch diag[i] {
			case 0:
				zeros++
			case 1:
				ones++
			}
		}

		if ones > zeros {
			gammaRate += int(math.Pow(2, float64(11-i)))
		} else {
			epsilonRate += int(math.Pow(2, float64(11-i)))
		}
	}

	return gammaRate * epsilonRate
}

func Part2(in Report) (res int) {
	var oxygenGeneratorRating int
	var co2ScrubberRating int

	oxygenGeneratorReport := in
	for i := 0; i < 12; i++ {
		var zeros int
		var ones int
		for _, diag := range oxygenGeneratorReport {
			switch diag[i] {
			case 0:
				zeros++
			case 1:
				ones++
			}
		}

		var oxygenGeneratorReportNew Report
		if zeros > ones {
			for _, diag := range oxygenGeneratorReport {
				if diag[i] == 0 {
					oxygenGeneratorReportNew = append(oxygenGeneratorReportNew, diag)
				}
			}
		} else {
			for _, diag := range oxygenGeneratorReport {
				if diag[i] == 1 {
					oxygenGeneratorReportNew = append(oxygenGeneratorReportNew, diag)
				}
			}
		}

		if len(oxygenGeneratorReportNew) == 1 {
			for k, v := range oxygenGeneratorReportNew[0] {
				if v == 1 {
					oxygenGeneratorRating += int(math.Pow(2, float64(11-k)))
				}
			}
			break
		}

		oxygenGeneratorReport = oxygenGeneratorReportNew
	}

	oxygenGeneratorReport = in
	for i := 0; i < 12; i++ {
		var zeros int
		var ones int
		for _, diag := range oxygenGeneratorReport {
			switch diag[i] {
			case 0:
				zeros++
			case 1:
				ones++
			}
		}

		var oxygenGeneratorReportNew Report
		if zeros > ones {
			for _, diag := range oxygenGeneratorReport {
				if diag[i] == 1 {
					oxygenGeneratorReportNew = append(oxygenGeneratorReportNew, diag)
				}
			}
		} else {
			for _, diag := range oxygenGeneratorReport {
				if diag[i] == 0 {
					oxygenGeneratorReportNew = append(oxygenGeneratorReportNew, diag)
				}
			}
		}

		if len(oxygenGeneratorReportNew) == 1 {
			for k, v := range oxygenGeneratorReportNew[0] {
				if v == 1 {
					co2ScrubberRating += int(math.Pow(2, float64(11-k)))
				}
			}
			break
		}

		oxygenGeneratorReport = oxygenGeneratorReportNew
	}

	return oxygenGeneratorRating * co2ScrubberRating
}

func In(r io.Reader) (res Report) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 12 {
			log.Fatalf("invalid input: %s", line)
		}

		var curr = diagnostic{}
		for k, v := range line {
			switch v {
			case '0':
				curr[k] = 0
			case '1':
				curr[k] = 1
			default:
				log.Fatalf("invalid input: %s", line)
			}
		}
		res = append(res, curr)
	}

	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
