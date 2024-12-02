package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solve() {
    file, err := os.Open("days/day1/input.txt")
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err)
		return
	}
	defer file.Close()

    var col1 []int
	var col2 []int

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        parts := strings.Fields(line)
        if len(parts) != 2 {
            fmt.Println("Invalid line: ", line)
            continue
        }

        num1, err1 := strconv.Atoi(parts[0])
        num2, err2 := strconv.Atoi(parts[1])
        if err1 != nil || err2 != nil {
            fmt.Println("Invalid numbers in line: ", line)
            continue
        }

        col1 = append(col1, num1)
        col2 = append(col2, num2)
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("Failed to read file: %s\n", err)
        return
    }

    sort.Ints(col1)
    sort.Ints(col2)

    diff := 0

    for i := 0; i < len(col1); i++ {
        // calcualte absolute value of difference
        abs := col1[i] - col2[i]
        println(abs)
        if abs < 0 {
            diff = diff + (-abs)
        } else {
            diff = diff + abs
        }
        
    }
    println(diff)
}