package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"github.com/PrinceNarteh/gqlgen_cc/graph/generated"
	"github.com/PrinceNarteh/gqlgen_cc/graph/model"
)

// User is the resolver for the user field.
func (r *meetUpResolver) User(ctx context.Context, obj *model.MeetUp) (*model.User, error) {
	user := new(model.User)

	for _, u := range users {
		if u.ID == obj.UserId {
			user = u
			break
		}
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// CreateMeetUp is the resolver for the createMeetUp field.
func (r *mutationResolver) CreateMeetUp(ctx context.Context, input model.NewMeetUp) (*model.MeetUp, error) {
	panic(fmt.Errorf("not implemented"))
}

// Meetups is the resolver for the meetups field.
func (r *queryResolver) Meetups(ctx context.Context) ([]*model.MeetUp, error) {
	return meetups, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user := new(model.User)

	for _, u := range users {
		if u.ID == id {
			user = u
			break
		}
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// Meetups is the resolver for the meetups field.
func (r *userResolver) Meetups(ctx context.Context, obj *model.User) ([]*model.MeetUp, error) {
	meetupsSlice := make([]*model.MeetUp, 0)

	for _, meetup := range meetups {
		if meetup.UserId == obj.ID {
			meetupsSlice = append(meetupsSlice, meetup)
		}
	}

	return meetupsSlice, nil
}

// MeetUp returns generated.MeetUpResolver implementation.
func (r *Resolver) MeetUp() generated.MeetUpResolver { return &meetUpResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type meetUpResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var meetups = []*model.MeetUp{
	{
		ID:          "1",
		Title:       "First Meetup",
		Description: "First Meetup Description",
		UserId:      "1",
	},
	{
		ID:          "2",
		Title:       "Second Meetup",
		Description: "Second Meetup Description",
		UserId:      "2",
	},
}
var users = []*model.User{
	{
		ID:        "1",
		Email:     "john.doe@email.com",
		Username:  "JDoe",
		FirstName: "John",
		LastName:  "Doe",
	},
	{
		ID:        "2",
		Email:     "jane.doe@email.com",
		Username:  "JaneDoe",
		FirstName: "Jane",
		LastName:  "Doe",
	},
}

func (r *userResolver) MeetUps(ctx context.Context, obj *model.User) (*model.MeetUp, error) {
	panic("implement me")
}
