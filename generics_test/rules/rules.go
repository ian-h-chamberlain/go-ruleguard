package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

func MatchGeneric(m dsl.Matcher) {
	m.Import("github.com/quasilyte/go-ruleguard/generics_test")

	m.Match(`GetT($x)`).
		Where(m["x"].Type.Implements("generics_test.GenIntf")).
		Report(`found an implementor`)
}
