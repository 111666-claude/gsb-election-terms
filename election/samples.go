package election

import (
	"encoding/json"
	"fmt"
)

// Sample 是一个固定场景，用来复现题面里的锚点。
type Sample struct {
	Name string
	Run  func() map[string]any
}

func tallyJSON(m map[string]int) any {
	b, _ := json.Marshal(m)
	return json.RawMessage(b)
}

// Samples 返回全部场景。
func Samples() []Sample {
	return []Sample{
		{Name: "dupvote", Run: func() map[string]any {
			st := NewState([]string{"a", "b", "c"})
			c := &Counter{}
			st.VoteFor("a", "X", 1, c)
			st.VoteFor("a", "X", 1, c)
			return map[string]any{"tally": tallyJSON(st.Tally(1, c))}
		}},
		{Name: "majority", Run: func() map[string]any {
			st := NewState([]string{"a", "b", "c", "d"})
			c := &Counter{}
			st.VoteFor("a", "X", 1, c)
			st.VoteFor("b", "X", 1, c)
			return map[string]any{"leader": st.LeaderOf(c)}
		}},
		{Name: "stale", Run: func() map[string]any {
			st := NewState([]string{"a", "b", "c"})
			c := &Counter{}
			st.VoteFor("a", "X", 1, c)
			st.VoteFor("b", "X", 1, c)
			st.VoteFor("a", "Y", 2, c)
			st.VoteFor("b", "Y", 2, c)
			st.VoteFor("c", "Y", 2, c)
			return map[string]any{"leader": st.LeaderOf(c)}
		}},
		{Name: "reject", Run: func() map[string]any {
			st := NewState([]string{"a", "b", "c"})
			c := &Counter{}
			st.VoteFor("a", "X", 2, c)
			return map[string]any{"accepted": st.VoteFor("a", "Y", 1, c)}
		}},
	}
}

// Work 跑规模线场景。
func Work(n int) map[string]any {
	st := NewState([]string{"a", "b", "c", "d", "e"})
	c := &Counter{}
	for i := 0; i < n; i++ {
		st.VoteFor(fmt.Sprintf("v%04d", i), "X", i, c)
	}
	return map[string]any{"votes": n, "scanned": c.Scanned}
}
