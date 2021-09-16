package desk

import "fmt"

type ModernDesk struct{}

func (*ModernDesk) GetDrawer() {
	fmt.Print("it is Drawer of Modern Modern Modern Desk !!!! \n")
}

func (*ModernDesk) AutoRising() {
	fmt.Print("ModernDesk is Rising !!!!!!\n")
}
