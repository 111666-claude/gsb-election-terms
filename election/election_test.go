package election

import "testing"

func TestFirstVoteAccepted(t *testing.T) {
	st := NewState([]string{"a", "b", "c"})
	if !st.VoteFor("a", "X", 1, &Counter{}) {
		t.Fatal("首票应被接受")
	}
}

func TestLeaderEmptyInitially(t *testing.T) {
	st := NewState([]string{"a", "b", "c"})
	if st.LeaderOf(&Counter{}) != "" {
		t.Fatal("初始不该有当选者")
	}
}

func TestVotersStored(t *testing.T) {
	st := NewState([]string{"a", "b"})
	if len(st.Voters) != 2 {
		t.Fatalf("应存下投票者：%+v", st.Voters)
	}
}

func TestTallyUnknownTermEmpty(t *testing.T) {
	st := NewState([]string{"a"})
	if got := st.Tally(9, &Counter{}); len(got) != 0 {
		t.Fatalf("未知任期应无票：%+v", got)
	}
}

func TestTallyRecordsVote(t *testing.T) {
	st := NewState([]string{"a", "b", "c"})
	c := &Counter{}
	st.VoteFor("a", "X", 1, c)
	if got := st.Tally(1, c)["X"]; got != 1 {
		t.Fatalf("X 应有 1 票，得到 %d", got)
	}
}

func TestCounterCounts(t *testing.T) {
	st := NewState([]string{"a"})
	c := &Counter{}
	st.VoteFor("a", "X", 1, c)
	if c.Scanned == 0 {
		t.Fatal("应统计比较次数")
	}
}
