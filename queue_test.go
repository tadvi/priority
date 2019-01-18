package priority

import (
	"fmt"
	"strings"
	"testing"
)

type QItem struct {
	name     string
	priority int
}

func (qi *QItem) Less(other Item) bool {
	return qi.priority < other.(*QItem).priority
}

func TestQueue(t *testing.T) {
	const want = "&{bobby 1},&{bob 0},&{frank 2},&{crab 4},&{frog 6},&{frank 8}"

	var pit Queue
	pit.Push(&QItem{"bobby", 1})
	pit.Push(&QItem{"frank", 8})
	pit.Push(&QItem{"crab", 4})
	pit.Push(&QItem{"frog", 6})
	pit.Push(&QItem{"bob", 0})
	pit.Push(&QItem{"frank", 2})

	var out []string
	for len(pit) > 0 {
		out = append(out, fmt.Sprint(pit.Pop()))
	}

	got := strings.Join(out, ",")
	if want != got {
		t.Fatalf("got %q, want %q", want, got)
	}
}
