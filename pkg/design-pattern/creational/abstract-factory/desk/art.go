package chair

import "fmt"

type ArtChair struct{}

func (*ArtChair) SitOn() {
	fmt.Print("it is a beautifully art chair !!!! \n")
}

func (*ArtChair) Leave() {
	fmt.Print("Be careful not to break \n")
}

func (*ArtChair) HasLeg() {
	fmt.Print("ArtChair chair have 1 leg! \n")
}
