package main

import (
	"fmt"

	"github.com/rmrfslashbin/go-functionalOptions/thing"
)

func main() {
	myThing := thing.New(
		thing.SetColor("red"),
		thing.SetDescription("a red thing"),
		thing.SetLabel("red thing"),
		thing.SetSize(13),
		thing.SetVersion("1.0.1"),
	)

	fmt.Printf("Version %s of my %s thing is %d cm tall.\n\n", myThing.GetVersion(), myThing.GetLabel(), myThing.GetSize())

	defaultColorThing := thing.New(
		thing.SetDescription("another useful thing"),
		thing.SetLabel("default color thing"),
		thing.SetSize(5),
		thing.SetVersion("2.0.5"),
	)
	fmt.Printf("Version %s of my %s thing is %d cm tall.\n", defaultColorThing.GetVersion(), defaultColorThing.GetColor(), defaultColorThing.GetSize())
	fmt.Printf("I like to call it '%s'.\n", defaultColorThing.GetLabel())
	fmt.Printf("You can call it '%s'\n", defaultColorThing.GetDescription())
}
