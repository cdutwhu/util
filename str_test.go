package util

import (
	"fmt"
	"testing"
)

var pln = fmt.Println

func TestRemovePrefix(t *testing.T) {
	pln(Str("sif.abc").RemovePrefix("sif."))
}

func TestRemoveSuffix(t *testing.T) {
	pln(Str("sif.abc").RemoveSuffix("abc"))
}

func TestRemoveTailFromLast(t *testing.T) {
	pln(Str("a.sif.abc").RemoveTailFromLast("."))
}

func TestKeyValueMap(t *testing.T) {
	pln(Str(`<abc a="dd"  c=fff>>>>>`).KeyValueMap(' ', '=', '>'))
	pln(Str(`<abc a="dd"  c=fff>>>>>`).KeyValueMap(' ', '=', '|'))
}

func TestMakeQuotes(t *testing.T) {
	pln(Str("abc").MakeQuotes(QDouble))
}

func TestRemoveQuotes(t *testing.T) {
	pln(Str("'abc'").RemoveQuotes())
}

func TestMakeBrackets(t *testing.T) {
	pln(Str("abc").MakeBrackets(BCurly))
}

func TestRemoveBrackets(t *testing.T) {
	pln(Str("<abc>").RemoveBrackets())
}
