package zdata

import (
	"fmt"
	"testing"
)

func TestData(t *testing.T) {
	var data = FetchLists("999")
	if data == nil {
		t.FailNow()
	}
	fmt.Println(data)
}
