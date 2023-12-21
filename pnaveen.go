package pnaveen

import (
	"fmt"

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

func From10() {
	fmt.Println("Im from version 1.0.0")
}
func From11() {
	fmt.Println("Im from version 1.2.0")
}
