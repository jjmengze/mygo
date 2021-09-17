package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/jjmengze/mygo/internal/delivery/graph/generated"
	"github.com/jjmengze/mygo/internal/delivery/graph/model"
)

func (r *queryResolver) Users(ctx context.Context) (*model.User, error) {
	return nil, nil
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	//id := 10
	//name := string("happy")
	age := 20
	return &model.User{
		//ID:   id,
		//Name: &name,
		Age: &age,
		//if return null friends
		Friends: []*model.User{
			{
				Age: &age,
			},
			//	{
			//		Age: &age,
			//	},
			//	{
			//		Age: &age,
			//	},
			//	{
			//		Age: &age,
			//	},
		},
	}, nil
}

func (r *queryResolver) Hello(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, name string) (*model.User, error) {
	//fmt.Println("Hello")
	age := rand.Int()
	var w float64 = 80
	var h float64 = 180
	return &model.User{
		//ID:   id,
		Name:   &name,
		Age:    &age,
		Weight: &w,
		Height: &h,
		Friends: []*model.User{
			{
				Age: &age,
			},
		},
	}, nil
}

func (r *userResolver) Friends(ctx context.Context, obj *model.User) ([]*model.User, error) {
	fmt.Println("parent obj age", *obj.Age)
	age := rand.Int()
	var w float64 = 80
	var h float64 = 180
	f := []*model.User{
		&model.User{
			Name: obj.Name,
			Age:  &age,
			//Friends: ,
			Height: &h,
			Weight: &w,
		},
	}
	return f, nil
}

func (r *userResolver) Height(ctx context.Context, obj *model.User, unit *model.HeightUnit) (*float64, error) {
	var h float64
	switch *unit {
	case model.HeightUnitFoot:
		h = 180 / 30.48
	case model.HeightUnitMetre:
		h = 180 / 100
	case model.HeightUnitCentimetre:
		h = 180
	}
	return &h, nil
}

func (r *userResolver) Weight(ctx context.Context, obj *model.User, unit *model.WeightUnit) (*float64, error) {
	var w float64
	switch *unit {
	case model.WeightUnitGram:
		w = 70 * 100
	case model.WeightUnitPound:
		w = 70 / 0.45359237
	case model.WeightUnitKilogram:
		w = 70
	}
	return &w, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
