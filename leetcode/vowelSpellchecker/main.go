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
	lmap := make(map[string]string, 0)
	vmap := make(map[string]string, 0)
	vowels := map[rune]rune{'a': '*', 'e': '*', 'i': '*', 'o': '*', 'u': '*'}

	for _, v := range wordlist {
		rmap[v] = struct{}{}

		if _, ok := lmap[strings.ToLower(v)]; !ok {
			lmap[strings.ToLower(v)] = v
		}

		if _, ok := vmap[strings.Map(func(r rune) rune {
			if u, ok := vowels[r]; ok {
				return u
			}
			return r
		}, v)]; !ok {
			vmap[strings.Map(func(r rune) rune {
				if u, ok := vowels[r]; ok {
					return u
				}
				return r
			}, strings.ToLower(v))] = v
		}
	}

	for _, v := range queries {
		if _, ok := rmap[v]; ok {
			r = append(r, v)
			continue
		}
		value, ok := lmap[strings.ToLower(v)]
		if ok {
			r = append(r, value)
			continue
		}

		value, ok = vmap[strings.Map(func(r rune) rune {
			if u, ok := vowels[r]; ok {
				return u
			}
			return r
		}, strings.ToLower(v))]

		if ok {
			r = append(r, value)
			continue
		}
		r = append(r, "")

	}

	return r
}

func main() {
	//wl := []string{"KiTe", "kite", "hare", "Hare"}
	//qu := []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}

	t1 := []string{"ae", "aa"}
	t2 := []string{"UU"}
	//fmt.Println(spellchecker(wl, qu))
	fmt.Println(spellchecker(t1, t2))

}
