package src

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

type Ch1 struct {
	name string
	count int
}

func Challenge1() {
	fmt.Println("------------------ Challenge 1 ------------------")
	// Read txt file
	content, err := os.ReadFile("challenge_input/challenge1.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	fileContent := string(content)

	// Split string into array
	fileContentArray := strings.Split(fileContent, " ")
	getData := []Ch1{}

	for i := 0; i < len(fileContentArray); i++ {
		worldSecret := fileContentArray[i]
		if len(getData) == 0 {
			getData = append(getData, Ch1{worldSecret, 1})
		} else {
			for j := 0; j < len(getData); j++ {
				if getData[j].name == worldSecret {
					getData[j].count++
					break
				} else if j == len(getData) - 1 {
					getData = append(getData, Ch1{worldSecret, 1})
					break
				}
			}
		}
	}
	fmt.Println(getData)
	fmt.Println("------------------ Solution ------------------")
	var bufferOut bytes.Buffer
	for i := 0; i < len(getData); i++ {
		bufferOut.WriteString(fmt.Sprintf("%s%d", getData[i].name, getData[i].count)) 
	}
	fmt.Println(bufferOut.String())
	
}