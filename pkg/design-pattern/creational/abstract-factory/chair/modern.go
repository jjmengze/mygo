package chair

import "fmt"

type ModernChair struct{}

func (*ModernChair) SitOn() {
	fmt.Print("it is google modern chair please sit it on \n")
}

func (*ModernChair) Leave() {
	fmt.Print("modern chair hope you enjoy it \n")
}

func (*ModernChair) HasLeg() {
	fmt.Print("modern chair have no any leg! \n")
}
