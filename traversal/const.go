package traversal

type State int
const (
	Continue State = iota
	Return State  = iota
	Match State = iota
	MatchReturn State =iota

)
