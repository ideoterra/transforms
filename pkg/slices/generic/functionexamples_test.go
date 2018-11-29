package generic_test

import (
	"strings"

	"github.com/ideoterra/transforms/pkg/slices/generic"
)

func Example_mapReduce() {
	sentance := "it was the best of times, it was the worst of times"

	words := []interface{}{}
	for _, word := range strings.Split(sentance, " ") {
		generic.Append(&words, word)
	}

	wordLength := func(word interface{}) interface{} {
		return len(word.(string))
	}

	wordLengths := generic.Map(&words, wordLength)

	totalLength := generic.Reduce(wordLengths, func(a, acc interface{}) interface{} {
		return a.(int) + acc.(int)
	})

}
