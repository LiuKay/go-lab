package main

import (
	"algorithmns/bubblesort"
	"algorithmns/qsort"
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

var inputFile *string = flag.String("i", "infile", "File contains values for sorting")
var outFile *string = flag.String("o", "outfile", "File to store sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func main() {
	flag.Parse()

	if inputFile != nil {
		fmt.Println("infile=", *inputFile, "outfile=", *outFile, "algorithm=", *algorithm)
	}

	values, err := readValues(*inputFile)
	if err != nil {
		t1 := time.Now()
		switch *algorithm {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("Sorting algorithm", *algorithm, "is unsupported.")
		}
		t2 := time.Now()
		fmt.Println("The sorting process costs", t2.Sub(t1), "time to complete.")

		writeValues(values, *outFile)
	} else {
		fmt.Println(err)
	}

}

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file ", infile)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	values = make([]int, 0)
	for {
		line, isPrefix, err1 := reader.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			return
		}
		if isPrefix {
			fmt.Println("A too lng line, seems unexpected.")
			return
		}

		str := string(line)

		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create file:", outfile)
		return err
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}
