package traversal

type State int

const (
	// Continue iterating with no evaluation of the current field or key
	Continue    State = iota
	// Return stop iterating and return up
	Return      State = iota
	// Match evaluate any field or key, will continue iterating after
	Match       State = iota
	// MatchReturn evaluate any field or key and then return up
	MatchReturn State = iota
)
