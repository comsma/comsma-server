package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	database "github.com/comsma/comsma/database"
	"github.com/comsma/comsma/graph/generated"
	"github.com/comsma/comsma/graph/model"
	"github.com/comsma/comsma/internal/gallery"
	"github.com/comsma/comsma/internal/message"
)

// Description is the resolver for the description field.
func (r *galleryResolver) Description(ctx context.Context, obj *database.Gallery) (string, error) {
	panic(fmt.Errorf("not implemented: Description - description"))
}

// Imagelocation is the resolver for the Imagelocation field.
func (r *galleryResolver) Imagelocation(ctx context.Context, obj *database.Gallery) (string, error) {
	panic(fmt.Errorf("not implemented: Imagelocation - Imagelocation"))
}

// SendContactMessage is the resolver for the sendContactMessage field.
func (r *mutationResolver) SendContactMessage(ctx context.Context, input model.Contact) (*model.ContactResult, error) {
	message.SendEmail(input.Name, input.Message, input.Email)
	return nil, nil
}

// AddGallery is the resolver for the addGallery field.
func (r *mutationResolver) AddGallery(ctx context.Context, input model.NewGallery) (*database.Gallery, error) {
	panic(fmt.Errorf("not implemented: AddGallery - addGallery"))
}

// GetGalleries is the resolver for the getGalleries field.
func (r *queryResolver) GetGalleries(ctx context.Context) ([]*database.Gallery, error) {
	var galleries []database.Gallery

	gallery.GetGalleries(galleries)
	return galleries, nil
}

// Gallery returns generated.GalleryResolver implementation.
func (r *Resolver) Gallery() generated.GalleryResolver { return &galleryResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type galleryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
