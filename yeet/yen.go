package yeet

import (
	_ "embed"
)

//go:embed richard.txt
var richard string

func Do() string {
	return richard
}
