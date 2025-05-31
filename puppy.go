package puppy

import (
	"github.com/alexpereap/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woff! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func SmallBark() string {
	return dog.WhenSmall(Bark())
}

func SmallBarks() string {
	return dog.WhenSmall(Barks())
}
