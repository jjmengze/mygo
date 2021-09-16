package abstract_factory

import (
	"github.com/jjmengze/mygo/pkg/design-pattern/creational/abstract-factory/chair"
	"github.com/jjmengze/mygo/pkg/design-pattern/creational/abstract-factory/desk"
	"github.com/jjmengze/mygo/pkg/design-pattern/creational/abstract-factory/sofa"
)

type FurnishingFactory interface {
	CreateChair() chair.Chair
	CreateDesk() desk.Desk
	CreateSofa() sofa.Sofa
}

type ArtFurnishingFactory struct {}

var _ FurnishingFactory = &ArtFurnishingFactory{}

func (a ArtFurnishingFactory) CreateChair() chair.Chair {
	return &chair.ArtChair{}
}

func (a ArtFurnishingFactory) CreateDesk() desk.Desk {
	return &desk.ArtDesk{}
}

func (a ArtFurnishingFactory) CreateSofa() sofa.Sofa {
	return &sofa.ArtSofa{}
}

type ModernFurnishingFactory struct {}

var _ FurnishingFactory = &ModernFurnishingFactory{}

func (a ModernFurnishingFactory) CreateChair() chair.Chair {
	return &chair.ModernChair{}
}

func (a ModernFurnishingFactory) CreateDesk() desk.Desk {
	return &desk.ModernDesk{}
}

func (a ModernFurnishingFactory) CreateSofa() sofa.Sofa {
	return &sofa.ModernSofa{}
}

var _ FurnishingFactory = &VictorianFurnishingFactory{}

type VictorianFurnishingFactory struct {}

func (a VictorianFurnishingFactory) CreateChair() chair.Chair {
	return &chair.VictorianChair{}
}

func (a VictorianFurnishingFactory) CreateDesk() desk.Desk {
	return &desk.VictorianDesk{}
}

func (a VictorianFurnishingFactory) CreateSofa() sofa.Sofa {
	return &sofa.VictorianSofa{}
}
