package metrics

import (
	"fmt"
	"testing"
)

type testData struct {
	key   string
	value float64
}

var mergeValueTests = []testData{
	testData{"key01", 10},
	testData{"key02", 20},
	testData{"key03", 30},
}

func TestValues(t *testing.T) {
	v := make(Values)
	v["key00"] = 1
	for _, dt := range mergeValueTests {
		v.Merge(Values{dt.key: dt.value})
	}

	fmt.Println("map: ", v)
}
