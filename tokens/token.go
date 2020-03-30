package tokens

type Keyword string

const (
	ID Keyword = "$id"
)

func (s Keyword) String() string {
	return string(s)
}
