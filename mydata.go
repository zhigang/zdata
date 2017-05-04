package zdata

import (
	"math/rand"
	"strconv"
	"time"
)

// FetchLists is return test data.
// The argument is unimportant things.
func FetchLists(applyID string) map[string]interface{} {

	time.Sleep(30 * time.Millisecond)

	timespan := time.Now()
	rand := rand.New(rand.NewSource(timespan.UnixNano()))
	cnt := rand.Intn(100)
	var lists string

	for i := 0; i < cnt; i++ {
		lists += "item-" + strconv.Itoa(i) + ","
	}

	return map[string]interface{}{
		"ApplyID": applyID,
		"Time":    timespan,
		"Count":   cnt,
		"Lists":   lists,
	}
}
