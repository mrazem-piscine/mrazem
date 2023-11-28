package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func String(text string) string {
	if len(text) == 0 {
		return ""
	}

	lines := strings.Split(text, "\n")
	result := make([]string, 0)

	for _, line := range lines {
		if len(line) == 0 {
			result = append(result, "")
			continue
		}

		chars := []byte(line)
		s := getArt(chars[0])

		for i := 1; i < len(chars); i++ {
			s = fuse(s, getArt(chars[i]))
		}

		result = append(result, s...)
	}

	return strings.Replace(strings.Join(result, "\n"), "$", " ", -1)
}

func fuse(left, right []string) []string {
	var step []int

	for j := 0; j < 6; j++ {
		k := len(left[j]) + len(right[j]) -
			len(strings.TrimRight(left[j], " ")) -
			len(strings.TrimLeft(right[j], " "))
		step = append(step, k)
	}

	sort.Ints(step)
	c := step[0]
	for j := 0; j < 6; j++ {
		cc := ""
		for i := 1; i <= c; i++ {
			lc := left[j][len(left[j])-(c-i+1) : len(left[j])-(c-i)]
			rc := right[j][i-1 : i]
			if lc == " " {
				cc += rc
			} else {
				cc += lc
			}
		}
		left[j] = left[j][:len(left[j])-c] + cc + right[j][c:]
	}
	return left
}

func main() {
	fmt.Println("Enter the string:")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	text = strings.ReplaceAll(text, "\\n", "\n")

	asciiArt := String(text)

	fmt.Println(asciiArt)
}
