package grapql

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/jmoiron/sqlx"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) Config {
	return Config{
		Resolvers: &Resolver{
			DB: db,
		},
		Directives: DirectiveRoot{
			User: func(ctx context.Context, obj interface{}, next graphql.Resolver, userID string) (res interface{}, err error) {
				return next(context.WithValue(ctx, "userID", userID))
			},
		},
	}
}

func getUserID(ctx context.Context) string {
	if userID, ok := ctx.Value("userID").(string); ok {
		return userID
	}
	return ""
}

func (r *Resolver) Follower() FollowerResolver {
	return &followerResolver{r}
}
func (r *Resolver) Hub() HubResolver {
	return &hubResolver{r}
}
func (r *Resolver) Message() MessageResolver {
	return &messageResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Profile() ProfileResolver {
	return &profileResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Subscription() SubscriptionResolver {
	return &subscriptionResolver{r}
}

type followerResolver struct{ *Resolver }

func (r *followerResolver) User(ctx context.Context, obj *Follower) (*User, error) {
	panic("not implemented")
}

type hubResolver struct{ *Resolver }

func (r *hubResolver) User(ctx context.Context, obj *Hub) (*User, error) {
	panic("not implemented")
}

type messageResolver struct{ *Resolver }

func (r *messageResolver) User(ctx context.Context, obj *Message) (*User, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) MoveGif(ctx context.Context, uid string, currCat string, distCat string, gifID []string) ([]*Gif, error) {
	panic("not implemented")
}
func (r *mutationResolver) ToggleFollower(ctx context.Context, uid string, followerID string) ([]*Follower, error) {
	panic("not implemented")
}
func (r *mutationResolver) ChangeAvatar(ctx context.Context, uid string, avatar string) (*Gif, error) {
	panic("not implemented")
}
func (r *mutationResolver) ChangePass(ctx context.Context, uid string, oldPass string, newPass string) (*SuccessMsg, error) {
	panic("not implemented")
}
func (r *mutationResolver) AddGif(ctx context.Context, uid string, gifData string) ([]*Gif, error) {
	panic("not implemented")
}
func (r *mutationResolver) RmCustGif(ctx context.Context, uid string, gifID string) ([]*Gif, error) {
	panic("not implemented")
}
func (r *mutationResolver) RmFavoriteGif(ctx context.Context, uid string, gifID string) ([]*Gif, error) {
	panic("not implemented")
}
func (r *mutationResolver) FavoriteGif(ctx context.Context, uid string, gifID string, category string) (*SuccessMsg, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateHub(ctx context.Context, uid string, hubName string, status bool) ([]*Hub, error) {
	panic("not implemented")
}
func (r *mutationResolver) ChangeLogo(ctx context.Context, uid string, logo string) (*SuccessMsg, error) {
	panic("not implemented")
}
func (r *mutationResolver) ToggleClosingHub(ctx context.Context, uid string, hubID string) (*SuccessMsg, error) {
	panic("not implemented")
}
func (r *mutationResolver) ToggleHubStatus(ctx context.Context, uid string, hubID string) (*SuccessMsg, error) {
	panic("not implemented")
}
func (r *mutationResolver) AddChatGif(ctx context.Context, uid string, gifData string, hubID string) (*Message, error) {
	panic("not implemented")
}

type profileResolver struct{ *Resolver }

func (r *profileResolver) Categories(ctx context.Context, obj *Profile) ([]*Category, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetUserProfile(ctx context.Context, uid string) (*Profile, error) {
	panic("not implemented")
}
func (r *queryResolver) GetUserCategory(ctx context.Context, uid string, cat string) ([]*Gif, error) {
	panic("not implemented")
}
func (r *queryResolver) GetUserFollower(ctx context.Context, uid string) ([]*Follower, error) {
	panic("not implemented")
}
func (r *queryResolver) GetFollowerUser(ctx context.Context, uid string) ([]*Follower, error) {
	panic("not implemented")
}
func (r *queryResolver) GetUserHub(ctx context.Context, uid string) ([]*Hub, error) {
	panic("not implemented")
}
func (r *queryResolver) GetCustGif(ctx context.Context, uid string) ([]*Gif, error) {
	panic("not implemented")
}
func (r *queryResolver) GetGlobalHubs(ctx context.Context, search string) ([]*Hub, error) {
	panic("not implemented")
}
func (r *queryResolver) GetOtherUser(ctx context.Context, userID string) (*Follower, error) {
	panic("not implemented")
}
func (r *queryResolver) FirstLoadChatHub(ctx context.Context, hubID string) ([]*Message, error) {
	panic("not implemented")
}
func (r *queryResolver) LoadMoreChat(ctx context.Context, hubID string, createAt time.Time) ([]*Message, error) {
	panic("not implemented")
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) ChatGifAdded(ctx context.Context, hubID string) (<-chan *Message, error) {
	panic("not implemented")
}
