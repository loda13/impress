package main

import "testing"

/**
 * @Author: tang
 * @mail: yuetang2
 * @Date: 2022/7/31 16:30
 * @Desc:
 */
func TestWordsCount(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantword  int
		wanthan   int
		wantpunct int
	}{
		{"en1", "hello,playground", 2, 0, 1},
		{"en2", "hello, playground", 2, 0, 1},
		{"cn1", "你好世界", 0, 4, 0},
		{"encn1", "Hello你好世界", 1, 4, 0},
		{"encn2", "Hello 你好世界", 1, 4, 0},
		{"en3", "Catch the star that holds your destiny, the one that forever twinkles within your heart. Take advantage of precious opportunities while they still sparkle before you. Always believe that your ultimate goal is attainable as long as you commit yourself to it.", 42, 0, 4},
		{"cn2", "追随能够改变你命运的那颗星，那颗永远在你心中闪烁的明星。当它在你面前闪耀时，抓住这宝贵的机会。请谨记，只要你坚持不懈，最终的目标总能实现。", 0, 62, 7},
		{"encn3", "This is Faith at Faith Radio Online-Simply to Relax. Find the star that twinkles in your heart for you alone are capable of making your brightest dreams come true. Give your hopes everything you've got and you will catch the star that holds your destin.您正在收听的是faith轻松之声。寻找心中那颗闪耀的明星，因为只有你自己才能够让美好的梦想变成现实。满怀希望并全力以赴，你就能摘下改变命运的那颗星。", 48, 64, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _, _ := WordsCount(tt.input); got != tt.wantword {
				t.Errorf("单词：WordsCount() = %v, want %v", got, tt.wantword)
			}
			if _, got, _ := WordsCount(tt.input); got != tt.wanthan {
				t.Errorf("汉字：WordsCount() = %v, want %v", got, tt.wanthan)
			}
			if _, _, got := WordsCount(tt.input); got != tt.wantpunct {
				t.Errorf("标点：WordsCount() = %v, want %v", got, tt.wantpunct)
			}
		})
	}
}
