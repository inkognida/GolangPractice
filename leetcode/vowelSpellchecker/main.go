package main

import (
	"fmt"
	"strings"
)

func checkVowel(word string, query string, vowels map[rune]rune) bool {
	if len(word) != len(query) {
		return false
	}

	word = strings.Map(func(r rune) rune {
		if u, ok := vowels[r]; ok {
			return u
		}
		return r
	}, word)

	query = strings.Map(func(r rune) rune {
		if u, ok := vowels[r]; ok {
			return u
		}
		return r
	}, query)

	return word == query
}

func spellchecker(wordlist []string, queries []string) []string {
	r := make([]string, 0)
	rmap := make(map[string]struct{}, 0)
	vowel := map[rune]rune{'a': '*', 'e': '*', 'i': '*', 'o': '*', 'u': '*'}

	for _, v := range wordlist {
		rmap[v] = struct{}{}
	}
Main:
	for _, v := range queries {
		_, ok := rmap[v]
		if ok {
			r = append(r, v)
		} else {
			for _, w := range wordlist {
				wordl := strings.ToLower(w)
				querl := strings.ToLower(v)

				if wordl == querl || checkVowel(wordl, querl, vowel) {
					//fmt.Println(wordl == querl, checkVowel(wordl, querl, vowel))

					r = append(r, w)
					continue Main
				}
			}
			r = append(r, "")
			continue Main
		}
	}

	return r
}

func main() {
	wl := []string{"KiTe", "kite", "hare", "Hare"}
	qu := []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}

	t1 := []string{"YellOw"}
	t2 := []string{"yollow"}
	fmt.Println(spellchecker(wl, qu))
	fmt.Println(spellchecker(t1, t2))

	fmt.Println(strings.ToLower("kite") == strings.ToLower("keti"))
}
