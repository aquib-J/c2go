package ast

type ArraySubscriptExpr struct {
	Address  string
	Position string
	Type     string
	Kind     string
	Children []interface{}
}

func parseArraySubscriptExpr(line string) *ArraySubscriptExpr {
	groups := groupsFromRegex(
		"<(?P<position>.*)> '(?P<type>.*?)' (?P<kind>.*)",
		line,
	)

	return &ArraySubscriptExpr{
		Address: groups["address"],
		Position: groups["position"],
		Type: groups["type"],
		Kind: groups["kind"],
		Children: []interface{}{},
	}
}