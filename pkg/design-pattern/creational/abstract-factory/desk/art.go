package desk

import "fmt"

type ArtDesk struct{}

func (*ArtDesk) GetDrawer() {
	fmt.Print("it is Drawer of beautifully Desk !!!! \n")
}

func (*ArtDesk) AutoRising() {
	fmt.Print("beautifully is beautifully Rising !!!!!!\n")
}
