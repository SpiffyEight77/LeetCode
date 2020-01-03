package main

import (
	"fmt"
	"strings"
)

func replaceWords(dict []string, sentence string) string {
	//s := []byte(sentence)
	s := strings.Split(sentence, " ")
	//fmt.Println(s)
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(dict); j++ {
			if len(dict[j]) == 1 {
				t := strings.Split(s[i], "")
				if dict[j] == t[0] {
					s[i] = dict[j]
					break
				}
			} else if dict[j] == s[i] {
				break
			} else {
				f := true
				t := strings.Split(s[i], "")
				ss := strings.Split(dict[j], "")

				//l := math.Min(float64(len(t)),float64(len(ss)))

				for k := 0; k < len(ss); k++ {
					if k == len(t) && len(t) != len(ss) || ss[k] != t[k] {
						f = false
						break
					}
				}
				if !f {
					continue
				}
				s[i] = dict[j]
				break
			}
		}
	}
	res := strings.Join(s, " ")
	return res
}

func main() {
	dict := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	//dict := []string{"a", "b", "c"}
	//sentence := "aadsfasf absbs bbab cadsfafs"
	fmt.Println(replaceWords(dict, sentence))
}
