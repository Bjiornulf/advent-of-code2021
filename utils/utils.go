package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)



func ReadInts(filename string) []int {
	result := ReadLines(filename);
	ints := make([]int, len(result));
	var err error;
	for i := 0; i < len(result); i++ {
		ints[i], err = strconv.Atoi(result[i]);
		if (err != nil) {
			log.Fatal(err);
		}
	}
	return ints;
}

func ReadLines(filename string) []string {
	f, err := os.Open(filename);
	defer f.Close() // never forget to close the file at the end of the function
	if (err != nil) {
		panic(err);
	}
	scanner := bufio.NewScanner(f);
	result := make([]string, 0);
	for scanner.Scan() {
		result = append(result, scanner.Text());
	}
	return result;
}
