package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const (
	letters = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
	numbers = `0123456789`
	special = `-=~!@#$%^&*()_+[]\{}|;':",./<>?`
)

var (
	help     bool
	pwLength = 16
	lettersRune = []rune(letters)
	numbsersRune = []rune(numbers)
	specialRune = []rune(special)
	s = rand.NewSource(time.Now().Unix())
	r = rand.New(s)
)
func init() {
	flag.BoolVar(&help, "h", false, "show this help")
	flag.IntVar(&pwLength, "l", pwLength, "password length 8-255")
	flag.Parse()
}
func main() {
	if help || pwLength< 8 || pwLength >255 {
		flag.Usage()
		return
	}
	pw := ramdomPwGen(pwLength)
	for v := range pw {
		fmt.Printf("%c",v)
	}
	fmt.Println()
}
func ramdomPwGen(l int) (out chan rune){
	out = make(chan rune, l)
	for {
		rn := r.Int()
		select{
		case out <- lettersRune[rn%len(lettersRune)]:
		case out <- lettersRune[rn%len(lettersRune)]:
		case out <- lettersRune[rn%len(lettersRune)]:
		case out <- lettersRune[rn%len(lettersRune)]:
		case out <- lettersRune[rn%len(lettersRune)]:
		case out <- lettersRune[rn%len(lettersRune)]:
		case out <- lettersRune[rn%len(lettersRune)]:
		case out <- lettersRune[rn%len(lettersRune)]:
		case out <- lettersRune[rn%len(lettersRune)]:
		case out <- lettersRune[rn%len(lettersRune)]:
			
		case out <- numbsersRune[rn%len(numbsersRune)]:
		case out <- numbsersRune[rn%len(numbsersRune)]:
		case out <- numbsersRune[rn%len(numbsersRune)]:
			
		case out <- specialRune[rn%len(specialRune)]:
		default:
			close(out)
			return
		}
	}
}