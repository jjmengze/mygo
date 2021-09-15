package chair

import "fmt"

type VictorianChair struct{}

func (*VictorianChair) SitOn() {
	fmt.Print("it is Victorian!!!! \n")
}

func (*VictorianChair) Leave() {
	fmt.Print("oh why you leave Victorian \n")
}

func (*VictorianChair) HasLeg() {
	fmt.Print("Victorian chair have 4 leg! \n")
}
