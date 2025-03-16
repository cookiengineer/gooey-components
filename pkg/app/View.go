//go:build wasm

package app

import "github.com/cookiengineer/gooey/pkg/dom"

type View interface {
	Init()
	GetElement(string) *dom.Element
	SetElement(string, *dom.Element)
	RemoveElement(string) bool
	Enter() bool
	Leave() bool
}

