package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type color map[[3]int]string

func extractRGBandColors(colorInput color) ([]string, []string) {
	rgbs := make([]string, 0, 10)
	colors := make([]string, 0, 10)
	for rgb, color := range colorInput {
		rgbStr := strings.Join([]string{
			strconv.Itoa(rgb[0]),
			strconv.Itoa(rgb[1]),
			strconv.Itoa(rgb[2]),
		}, ",")
		rgbs = append(rgbs, rgbStr)
		colors = append(colors, color)
	}
	return rgbs, colors
}

func createFile() *os.File {
	file, err := os.Create("colors.txt")
	if err != nil {
		fmt.Println("Can't create the file due to the following error:", err)
		return nil
	}
	return file
}

func writeToTXT(file *os.File, rgbs, colors []string) {
	var message string
	for i := 0; i < len(rgbs); i++ {
		message = fmt.Sprintf("\"%s\": \"%s\"\n", rgbs[i], colors[i])
		_, err := file.WriteString(message)
		if err != nil {
			fmt.Println("Failed to write your message to the file:", err)
			return
		}
	}
}
func main() {
	colorInput := color{
		{132, 53, 148}:  "Purple",
		{245, 152, 177}: "Pink",
	}

	rgbs, colors := extractRGBandColors(colorInput)
	file := createFile()
	writeToTXT(file, rgbs, colors)
	//fmt.Println(colors)
	//fmt.Println(rgbs)

}
