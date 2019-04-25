package main

import "math/rand"
import "time"
import "flag"
import "strconv"

func main() {

	difficulty := flag.Int("diff", 1, "Difficulty for processing challenge")
	length := flag.Int("l", 200000, "Length of random text to include in message")
	outputMethod := flag.String("o", "stdout", "Where to output data [stdout|kafka|bus]")
	destination := flag.String("dest", "", "Destination of output method")
	count := flag.Int("c", 1, "Number to produce (0 for unlimited)")

	flag.Parse()

	for i := 0; *count == 0 || i < *count; i++ {

		test := RandStringRunes(*length)

		data := strconv.Itoa(*difficulty) + ":" + test

		var output output
		if *outputMethod == "stdout" {
			output = stdout{data: data}
		} else if *outputMethod == "kafka" {
			output = kafka{data: data}
		} else {
			output = bus{data: data, destination: *destination}
		}
		output.out()
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
