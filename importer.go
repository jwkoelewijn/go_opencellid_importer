package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseCountryCodes(codes string) []string {
	return strings.Split(codes, ",")
}

func contains(collection []string, item string) bool {
	for _, elem := range collection {
		if elem == item {
			return true
		}
	}
	return false
}

func main() {

	infilePtr := flag.String("inputFile", "cell_towers.csv", "Input file downloaded from OpenCellID")
	outfilePtr := flag.String("outputFile", "selected_cell_towers.csv", "Output file with a selection of celltowers")
	countriesPtr := flag.String("countryCodes", "204,206,262", "MCC (Mobile Country Codes) to select from inputFile")

	flag.Parse()

	countries := parseCountryCodes(*countriesPtr)
	fmt.Println(fmt.Sprintf("Selecting country codes %+v", countries))

	inFile, err := os.Open(*infilePtr)
	if err != nil {
		panic(fmt.Errorf("Could not open file: %+v", err))
	}
	defer inFile.Close()

	outFile, err := os.Create(*outfilePtr)
	if err != nil {
		panic(fmt.Errorf("Could not open file: %+v", err))
	}
	defer outFile.Close()

	var total int64 = 0

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	_, err = writeHeader(outFile)
	if err != nil {
		panic(fmt.Sprintf("Could not write header %+v", err))
	}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		country := parts[1]
		if contains(countries, country) {
			_, err := writeLine(outFile, line)
			if err != nil {
				fmt.Printf("Could not write line: %s, Reason: %+v\n", line, err)
			} else {
				if total%100 == 0 {
					fmt.Print(".")
				}
				total += 1
			}
		}
	}
	fmt.Printf("\nFound %d towers\n", total)
}

func writeHeader(file *os.File) (int, error) {
	header := "mcc,mnc,lac,cell,longitude,latitude,range,samples,average_signal\n"
	return io.WriteString(file, header)
}

func writeLine(file *os.File, line string) (int, error) {
	fields := strings.Split(line, ",")
	newLine := fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s\n", fields[1], fields[2], fields[3], fields[4], fields[6], fields[7], fields[8], fields[9], fields[13])
	cell_id, err := strconv.Atoi(fields[4])
	if err != nil || cell_id > math.MaxInt32 {
		fmt.Printf("\nSkipping due to invalid data for cell_id: %s\n", fields[4])
		return 0, nil
	}
	return io.WriteString(file, newLine)
}
