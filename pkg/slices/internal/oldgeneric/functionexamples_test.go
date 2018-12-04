package oldgeneric_test

import (
	"fmt"
	"strings"

	"github.com/ideoterra/transforms/pkg/slices/internal/oldgeneric"
)

func Example_mapReduce() {
	sentance := "it was the best of times, it was the worst of times"

	words := []interface{}{}
	for _, word := range strings.Split(sentance, " ") {
		oldgeneric.Append(&words, word)
	}

	wordLength := func(word interface{}) interface{} {
		return len(word.(string))
	}

	wordLengths := oldgeneric.Map(words, wordLength)

	totalLength := oldgeneric.Reduce(wordLengths, func(a, acc interface{}) interface{} {
		return a.(int) + acc.(int)
	})

	fmt.Println(totalLength)
	//Output:[40]
}
