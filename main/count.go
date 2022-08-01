package main

import (
	"fmt"
	"unicode"
)

/**
 * @Author: tang
 * @mail: yuetang2
 * @Date: 2022/7/31 14:10
 * @Desc:
 */

func WordsCount(s string) (int, int, int) {
	//ss, _ := ioutil.ReadFile("./describe/demo.txt")
	wordcount, hancount, punctcount := 0, 0, 0
	var sl = make([]string, 0, 10)
	var word string
	word = ""
	//plainWords := strings.Fields(s)
	//for _, word := range plainWords {
	//	runeCount := utf8.RuneCountInString(word)
	//	if len(word) == runeCount {
	//		wordcount++
	//	}
	//}
	for _, w := range s {
		if !unicode.IsLetter(w) {
			if word != "" {
				sl = append(sl, word)
			}
			word = ""
		} else {
			if !unicode.Is(unicode.Han, rune(w)) {
				word = fmt.Sprintf("%s%c", word, w)
			}
		}
	}
	if word != "" {
		sl = append(sl, word)
	}
	wordcount = len(sl)

	for _, v := range s {
		if unicode.Is(unicode.Han, v) {
			hancount++
		}
	}

	for _, v := range s {
		if unicode.IsPunct(v) {
			punctcount++
		}
	}
	return wordcount, hancount, punctcount

}

//func main() {
//	fmt.Println(WordsCount("hello,playground"))
//}

//fixme
// {"en1", "hello,playground", 2, 0, 1},
// {"encn1", "Hello你好世界", 1, 4, 0},
