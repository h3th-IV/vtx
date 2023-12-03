package sub

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Assignment() {

	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatalf("%e", err)
	}
	defer file.Close()
	cSlice := make([]int, 0)
	var total int

	Reader := bufio.NewScanner(file)
	for Reader.Scan() {
		line := Reader.Text()
		tish := strings.Split(line, "XXX")
		for _, item := range tish {
			for _, char := range item {
				num, err := strconv.Atoi(string(char))
				if err == nil {
					// fmt.Printf("%T", num)
					//cSlice = append(cSlice, num)
					total += num
				}

			}
		}
	}
	if err := Reader.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(cSlice)
	fmt.Println(total)
}
