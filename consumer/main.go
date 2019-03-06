package main

import "fmt"
import "os"
import "strings"
import "crypto/sha256"
import "encoding/hex"
import "strconv"

func main() {

	//get input - check for stdin for testing otherwse default to kafka configuration
	var input string
	if len(os.Args) > 1 {
		input = strings.Join(os.Args[1:], " ")
	} else {
		//todo integrate kafka
		input = "6:hello world"
	}

	splitLoc := strings.Index(input, ":")
	if splitLoc < 0 {
		panic("Input " + input + " does not contain a colon")
	}
	difficulty := input[0:splitLoc]
	content := input[splitLoc+1:]

	diffint, err := strconv.Atoi(difficulty)
	if err != nil {
		panic(err)
	}

	target := strings.Repeat("0", diffint)

	match := false
	nounce := 0
	var sha string
	for !match {

		hasher := sha256.New()
		hasher.Write([]byte(content + strconv.Itoa(nounce)))
		sha = hex.EncodeToString(hasher.Sum(nil))

		if sha[0:diffint] == target {
			match = true
		} else {
			nounce++
		}
	}

	fmt.Println(input + strconv.Itoa(nounce))
	fmt.Println(sha)

}
