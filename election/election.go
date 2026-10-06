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

	seen     map[string]map[int]bool
	lastTerm map[string]int
	tally    map[int]map[string]int
}

// NewState 建一份选举状态。
func NewState(voters []string) *State {
	return &State{
		Voters:   voters,
		seen:     map[string]map[int]bool{},
		lastTerm: map[string]int{},
		tally:    map[int]map[string]int{},
	}
}

// VoteFor 记一张票，返回是否被接受。
func (st *State) VoteFor(voter, candidate string, term int, c *Counter) bool {
	if st.seen == nil {
		st.seen = map[string]map[int]bool{}
		st.lastTerm = map[string]int{}
		st.tally = map[int]map[string]int{}
	}
	c.Scanned++
	if term < st.lastTerm[voter] {
		return false
	}
	c.Scanned++
	if st.seen[voter][term] {
		return false
	}
	if st.seen[voter] == nil {
		st.seen[voter] = map[int]bool{}
	}
	st.seen[voter][term] = true
	st.lastTerm[voter] = term
	st.Votes = append(st.Votes, Vote{Voter: voter, Candidate: candidate, Term: term})
	if st.tally[term] == nil {
		st.tally[term] = map[string]int{}
	}
	st.tally[term][candidate]++
	c.Scanned++
	if st.tally[term][candidate] > len(st.Voters)/2 && (st.Leader == "" || term > st.LeaderNo) {
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
