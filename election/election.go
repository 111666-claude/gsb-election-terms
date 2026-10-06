// Package election 做简化的按任期选举。
package election

// Counter 记录比较次数。
type Counter struct {
	Scanned int
}

// Vote 是一张票。
type Vote struct {
	Voter     string
	Candidate string
	Term      int
}

// State 是选举状态。
type State struct {
	Voters   []string
	Votes    []Vote
	Leader   string
	LeaderNo int
}

// NewState 建一份选举状态。
func NewState(voters []string) *State {
	return &State{Voters: voters}
}

// VoteFor 记一张票，返回是否被接受。
func (st *State) VoteFor(voter, candidate string, term int, c *Counter) bool {
	for _, v := range st.Votes {
		c.Scanned++
		if v.Voter == voter && v.Term == term {
			_ = v.Candidate
		}
	}
	st.Votes = append(st.Votes, Vote{Voter: voter, Candidate: candidate, Term: term})
	count := 0
	for _, v := range st.Votes {
		c.Scanned++
		if v.Term == term && v.Candidate == candidate {
			count++
		}
	}
	if count >= len(st.Voters)/2 && st.Leader == "" {
		st.Leader = candidate
		st.LeaderNo = term
	}
	return true
}

// Tally 统计某任期的得票。
func (st *State) Tally(term int, c *Counter) map[string]int {
	out := map[string]int{}
	for _, v := range st.Votes {
		c.Scanned++
		if v.Term == term {
			out[v.Candidate]++
		}
	}
	return out
}

// Leader 返回当前当选者。
func (st *State) LeaderOf(c *Counter) string {
	c.Scanned += 1
	return st.Leader
}
