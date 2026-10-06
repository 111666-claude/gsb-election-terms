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

	voted   map[string]map[int]bool
	maxTerm map[string]int
	tally   map[int]map[string]int
}

// NewState 建一份选举状态。
func NewState(voters []string) *State {
	return &State{
		Voters:  voters,
		voted:   map[string]map[int]bool{},
		maxTerm: map[string]int{},
		tally:   map[int]map[string]int{},
	}
}

// VoteFor 记一张票，返回是否被接受。
// 同一投票者在同一任期最多一票；见过更高任期后，低任期投票一律拒绝。
func (st *State) VoteFor(voter, candidate string, term int, c *Counter) bool {
	c.Scanned++
	if seen, ok := st.maxTerm[voter]; ok && term < seen {
		return false
	}

	byTerm := st.voted[voter]
	if byTerm == nil {
		byTerm = map[int]bool{}
		st.voted[voter] = byTerm
	}
	c.Scanned++
	if byTerm[term] {
		return false
	}

	byTerm[term] = true
	st.maxTerm[voter] = term
	st.Votes = append(st.Votes, Vote{Voter: voter, Candidate: candidate, Term: term})

	termTally := st.tally[term]
	if termTally == nil {
		termTally = map[string]int{}
		st.tally[term] = termTally
	}
	termTally[candidate]++

	count := termTally[candidate]
	if count > len(st.Voters)/2 && (st.Leader == "" || term > st.LeaderNo) {
		st.Leader = candidate
		st.LeaderNo = term
	}
	return true
}

// Tally 统计某任期的得票。
func (st *State) Tally(term int, c *Counter) map[string]int {
	c.Scanned++
	out := map[string]int{}
	for candidate, count := range st.tally[term] {
		out[candidate] = count
	}
	return out
}

// Leader 返回当前当选者。
func (st *State) LeaderOf(c *Counter) string {
	c.Scanned += 1
	return st.Leader
}
