package main

import "math/rand"
import "time"
import "flag"
import "strconv"

func main() {

	difficulty := flag.Int("d", 1, "Difficulty for processing challenge")
	length := flag.Int("l", 1234, "Length of random text to include in message")
	outputMethod := flag.String("o", "stdout", "Where to output data [stdout|kafka]")

	flag.Parse()
	test := RandStringRunes(*length)

	data := strconv.Itoa(*difficulty) + ":" + test

	var output output
	if *outputMethod == "stdout" {
		output = stdout{data: data}
	} else {
		output = kafka{data: data}
	}

	output.out()
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
