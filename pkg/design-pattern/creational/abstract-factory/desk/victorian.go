package desk

import "fmt"

type VictorianDesk struct{}

func (*VictorianDesk) GetDrawer() {
	fmt.Print("it is Drawer of Victorian Desk !!!! \n")
}

func (*VictorianDesk) AutoRising() {
	fmt.Print("Victorian can not Rising \n")
}
