package sofa

import "fmt"

type ArtSofa struct{}

func (*ArtSofa) SitOn() {
	fmt.Print("it is a beautifully art sofa !!!! \n")
}

func (*ArtSofa) Leave() {
	fmt.Print("Be careful not to break \n")
}

func (*ArtSofa) GetSoftLevel() {
	fmt.Print("ArtSofa sofa is art art art! \n")
}
