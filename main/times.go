package main

import (
	"fmt"
	"io/ioutil"
	"unicode"
)

/**
 * @Author: tang
 * @mail: yuetang2
 * @Date: 2022/5/29 15:55
 * @Desc: 统计文本文件中所有单词出现的次数: 1）读取文件可以用ioutil.ReadFile(); 2) 字符串操作可以引用strings包
 */

func Count() {
	//读取文件，仅适用于英文文档
	ss, _ := ioutil.ReadFile("words.txt")

	//sl := strings.Fields(string(ss))
	//sl := strings.Split(string(ss), " ")
	//strings包里没有合适的方法过滤符号

	//初始化slice
	var sl = make([]string, 0, 10)
	//定义word用于存放单词
	var word string
	word = ""
	//遍历ss读取字母到w
	for _, w := range ss {
		//判断w如果不是字母字符，即表示单词结束：可能是空格、回车、标点等，此时将单词存入word，并存入slice内，重置word继续循环
		if !unicode.IsLetter(rune(w)) {
			//为防止多个非字母字符相连，会存在word为空值插入slice，过滤word为空的情况
			if word != "" {
				sl = append(sl, word)
			}
			word = ""
			continue

		} else {
			//如果w一直是字母字符，则拼到word后面，直到遇到非字母字符，表示一个单词
			word = fmt.Sprintf("%s%c", word, w)
		}
	}
	// 初始化一个map用来存储结果、统计次数
	wordsMap := make(map[string]int, 10)
	for _, keyWord := range sl {
		wordsMap[keyWord]++
	}
	//打印单词和次数
	for key, value := range wordsMap {
		fmt.Printf("%s: %d\n", key, value)
	}
}

func main() {
	Count()
}
