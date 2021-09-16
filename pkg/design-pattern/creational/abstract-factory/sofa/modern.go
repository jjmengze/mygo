package sofa

import "fmt"

type ModernSofa struct{}

func (*ModernSofa) SitOn() {
	fmt.Print("it is google modern sofa please sit it on \n")
}

func (*ModernSofa) Leave() {
	fmt.Print("modern sofa hope you enjoy it \n")
}

func (*ModernSofa) GetSoftLevel() {
	fmt.Print("Victorian sofa is Modern Modern Modern! \n")
}
