package sofa

import "fmt"

type VictorianSofa struct{}

func (*VictorianSofa) SitOn() {
	fmt.Print("it is Victorian!!!! \n")
}

func (*VictorianSofa) Leave() {
	fmt.Print("oh why you leave Victorian \n")
}

func (*VictorianSofa) GetSoftLevel() {
	fmt.Print("Victorian sofa is soooooooooft! \n")
}
