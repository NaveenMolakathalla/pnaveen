package pnaveen

import (
	"github.com/NaveenMolakathalla/dnaveen"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dnaveen.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dnaveen.WhenGrownUp(Barks())
}
