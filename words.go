package words

import "fmt"
import "strings"

func WordCount(s string) map[string]int{
	m := make(map[string]int)
	words := strings.Fields(s)
	fmt.Println("BeE")
	for _, word := range words {
		m[word]  += 1
		//fmt.Println(word)
		//fmt.Printf("%v\n", m)
	}
	
	//fmt.Printf("%v", m)
	return m
}