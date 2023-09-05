package main

import (
	"fmt"
	"io"
	"os"
	"unicode"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	/* a := profile.Start(profile.CPUProfile, profile.ProfilePath("."))
	defer a.Stop() */

	args := os.Args
	f, err := os.Open(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	b := make([]byte, 1)
	//var b bytes.Buffer

	var inWord bool
	var wordCount int64
	//r := bufio.NewReader(f)
	for {

		//b, err := r.ReadByte()
		_, err := f.Read(b)
		if err == io.EOF {
			break
		}
		if unicode.IsSpace(rune(b[0])) && inWord {
			wordCount++
			inWord = false
		} else {
			inWord = true
		}
	}
	fmt.Println("total words:", wordCount)
}
