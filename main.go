package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type color map[[3]int]string

func main() {
	colorInput := color{
		{132, 53, 148}:  "Purple",
		{245, 152, 177}: "Pink",
	}

	rgbs := make([][3]int, 0, 10)
	colors := make([]string, 0, 10)
	for rgb, color := range colorInput {
		rgbs = append(rgbs, rgb)
		colors = append(colors, color)

	}

	fmt.Println(colors)
	fmt.Println(rgbs)

	var message string

	file, err := os.Create("colors.txt")
	if err != nil {
		fmt.Println("Can't create a=the file due to the following error:", err)
		return
	}

	for i := 0; i < len(rgbs); i++ {
		rgbStr := strings.Join([]string{
			strconv.Itoa(rgbs[i][0]),
			strconv.Itoa(rgbs[i][1]),
			strconv.Itoa(rgbs[i][2]),
		}, ",")
		message = fmt.Sprintf("\"%s\": \"%s\"\n", rgbStr, colors[i])
		_, err = file.WriteString(message)
		if err != nil {
			fmt.Println("Failed to write your message to the file:", err)
			return
		}
	}

}
