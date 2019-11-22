package grapql

import (
	context "context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/SamurAI-Software/gyfHub/models"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/crypto/bcrypt"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	DB           *sqlx.DB
	EnterRequest map[string]map[string]string
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
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Subscription() SubscriptionResolver {
	return &subscriptionResolver{r}
}

type followerResolver struct{ *Resolver }

func (r *followerResolver) User(ctx context.Context, obj *Follower) (*User, error) {
	find, err := models.Users(qm.Where("username = ?", obj.Username)).One(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	user := &User{Email: find.Email, Mobile: find.Mobile, Avatar: string(find.Avatar), Username: find.Username}
	return user, nil
}

type hubResolver struct{ *Resolver }

func (r *hubResolver) User(ctx context.Context, obj *Hub) (*User, error) {
	find, err := models.Users(qm.Where("username = ?", obj.Username)).One(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	user := &User{Email: find.Email, Mobile: find.Mobile, Avatar: string(find.Avatar), Username: find.Username}
	return user, nil
}

type messageResolver struct{ *Resolver }

func (r *messageResolver) User(ctx context.Context, obj *Message) (*User, error) {
	find, err := models.Users(qm.Where("username = ?", obj.Username)).One(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	user := &User{Email: find.Email, Mobile: find.Mobile, Avatar: string(find.Avatar), Username: find.Username}
	return user, nil
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) MoveGif(ctx context.Context, uid string, currCat string, distCat string, gifID []string) ([]*Gif, error) {
	var passback []*Gif
	for _, v := range gifID {
		find, err := models.UserFavorites(qm.Where("user_id = ?", uid), qm.And("gif_id = ?", v)).One(context.Background(), r.DB)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		find.Category = "distCat"
		_, err = find.Update(context.Background(), r.DB, boil.Whitelist(models.UserFavoriteColumns.Category))
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	found, err := models.UserFavorites(qm.Select(models.UserFavoriteColumns.GifID), qm.Where("user_id = ?", uid), qm.And("category = ?", currCat)).All(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for _, v := range found {
		gif, err := models.Gifs(qm.Where("id = ?", v.GifID)).One(context.Background(), r.DB)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		g := &Gif{ID: gif.ID, Gif: string(gif.GifData)}
		passback = append(passback, g)
	}
	return passback, nil
}
func (r *mutationResolver) ToggleFollower(ctx context.Context, uid string, followerName string) ([]*Follower, error) {
	var passback []*Follower

	fmt.Println("check user exist")
	find, err := models.Users(qm.Where("username = ?", followerName)).One(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("check follower in list")
	exist, err := models.Followers(qm.Where("user_id = ?", uid), qm.And("follower_id = ?", find.ID)).Exists(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if exist {
		_, err := models.Followers(qm.Where("user_id = ?", uid), qm.And("follower_id = ?", find.ID)).DeleteAll(context.Background(), r.DB)
		if err != nil {
			return nil, err
		}
	} else {
		id, err := uuid.NewV4()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		uf := &models.Follower{ID: id.String(), UserID: uid, FollowerID: find.ID}
		err = uf.Insert(context.Background(), r.DB, boil.Infer())
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	allfollowers, err := models.Followers(qm.Where("user_id = ?", uid)).All(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, v := range allfollowers {
		user, err := v.User().One(context.Background(), r.DB)
		if err != nil {
			fmt.Println(err)
		}
		f := &Follower{Username: user.Username, Isfollowing: true}
		passback = append(passback, f)
	}
	return passback, nil
}
func (r *mutationResolver) ChangeAvatar(ctx context.Context, uid string, avatar string) (bool, error) {
	find, err := models.FindUser(context.Background(), r.DB, uid)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	find.Avatar = []byte(avatar)
	_, err = find.Update(context.Background(), r.DB, boil.Whitelist(models.UserColumns.Avatar))
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}
func (r *mutationResolver) ChangePass(ctx context.Context, uid string, oldPass string, newPass string) (bool, error) {
	find, err := models.FindUser(context.Background(), r.DB, uid)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	err = bcrypt.CompareHashAndPassword(find.PasswordHash, []byte(oldPass))
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}

	find.PasswordHash = newPasswordHash
	_, err = find.Update(context.Background(), r.DB, boil.Whitelist(models.UserColumns.PasswordHash))
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	return true, nil
}

func ImageToID(img string) string {
	h := sha1.New()
	h.Write([]byte(img))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (r *mutationResolver) AddGif(ctx context.Context, uid string, gifData string) ([]*Gif, error) {
	id := ImageToID(gifData)
	exist, err := models.GifExists(context.Background(), r.DB, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if !exist {
		g := &models.Gif{ID: id, GifData: []byte(gifData)}
		err = g.Insert(context.Background(), r.DB, boil.Infer())
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	exist, err = models.UserGifs(qm.Where("user_id = ?", uid), qm.And("gif_id = ?", id)).Exists(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if !exist {
		ugID, err := uuid.NewV4()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		ug := &models.UserGif{ID: ugID.String(), UserID: uid, GifID: id}

		err = ug.Insert(context.Background(), r.DB, boil.Infer())
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	gifs, err := models.UserGifs(qm.Where("user_id = ?", uid)).All(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var passback []*Gif
	for _, v := range gifs {
		gif, err := v.Gif().One(context.Background(), r.DB)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		g := &Gif{ID: gif.ID, Gif: string(gif.GifData)}
		passback = append(passback, g)
	}
	return passback, nil
}
func (r *mutationResolver) RmCustGif(ctx context.Context, uid string, gifID string) ([]*Gif, error) {
	exist, err := models.UserGifs(qm.Where("user_id = ?", uid), qm.And("gif_id = ?", gifID)).Exists(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if exist {
		_, err := models.UserGifs(qm.Where("user_id = ?", uid), qm.And("gif_id = ?", gifID)).DeleteAll(context.Background(), r.DB)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	gifs, err := models.UserGifs(qm.Where("user_id = ?", uid)).All(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var passback []*Gif
	for _, v := range gifs {
		gif, err := v.Gif().One(context.Background(), r.DB)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		g := &Gif{ID: gif.ID, Gif: string(gif.GifData)}
		passback = append(passback, g)
	}
	return passback, nil
}
func (r *mutationResolver) RmFavoriteGif(ctx context.Context, uid string, gifID string) ([]*Gif, error) {
	exist, err := models.UserFavorites(qm.Where("user_id = ?", uid), qm.And("gif_id = ?", gifID)).Exists(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if !exist {
		return nil, err
	}

	uf, err := models.UserFavorites(qm.Where("user_id = ?", uid), qm.And("gif_id = ?", gifID)).One(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	category := uf.Category

	_, err = uf.Delete(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	gifs, err := models.UserFavorites(qm.Where("user_id = ?", uid), qm.And("category = ?", category)).All(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var passback []*Gif
	for _, v := range gifs {
		gif, err := v.Gif().One(context.Background(), r.DB)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		g := &Gif{ID: gif.ID, Gif: string(gif.GifData)}
		passback = append(passback, g)
	}
	return passback, nil
}
func (r *mutationResolver) FavoriteGif(ctx context.Context, uid string, gifData string, category string) ([]*Gif, error) {
	id := ImageToID(gifData)
	exist, err := models.GifExists(context.Background(), r.DB, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if !exist {
		gif := &models.Gif{ID: id, GifData: []byte(gifData)}
		err = gif.Insert(context.Background(), r.DB, boil.Infer())
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	exist, err = models.UserFavorites(qm.Where("user_id = ?", uid), qm.And("gif_id = ?", id)).Exists(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if !exist {
		ufID, err := uuid.NewV4()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		uf := &models.UserFavorite{ID: ufID.String(), UserID: uid, GifID: id, Category: category}
		err = uf.Insert(context.Background(), r.DB, boil.Infer())
	}

	gifs, err := models.UserFavorites(qm.Where("user_id = ?", uid), qm.And("category = ?", category)).All(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var passback []*Gif
	for _, v := range gifs {
		gif, err := v.Gif().One(context.Background(), r.DB)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		g := &Gif{ID: gif.ID, Gif: string(gif.GifData)}
		passback = append(passback, g)
	}

	return passback, nil
}
func (r *mutationResolver) CreateHub(ctx context.Context, uid string, hubName string, status bool) ([]*Hub, error) {
	exist, err := models.HubExists(context.Background(), r.DB, hubName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if exist {
		fmt.Println("hub exist!")
		return nil, err
	}

	hub := &models.Hub{ID: hubName, Logo: []byte(""), UserID: uid, IsPrivate: status, IsClose: false}
	err = hub.Insert(context.Background(), r.DB, boil.Infer())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	exist, err = models.HubUsers(qm.Where("hub_id = ?", hubName), qm.And("user_id = ?", uid)).Exists(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if exist {
		fmt.Println("user already have this hub")
		return nil, err
	}

	id, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	hu := &models.HubUser{ID: id.String(), HubID: hubName, UserID: uid}
	err = hu.Insert(context.Background(), r.DB, boil.Infer())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	hubs, err := models.HubUsers(qm.Where("user_id = ?", uid)).All(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var passback []*Hub
	for _, v := range hubs {
		hub, err = v.Hub().One(context.Background(), r.DB)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		hubers, err := hub.HubUsers().Count(context.Background(), r.DB)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		owner, err := hub.User().One(context.Background(), r.DB)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		t := time.Now()
		// latestMsg, err := hub.ChatMSGS(qm.Where("create_at < ?", t), qm.Limit(1)).One(context.Background(), r.DB)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return nil, err
		// }
		h := &Hub{ID: hub.ID, Logo: string(hub.Logo), Hubers: int(hubers), Status: hub.IsPrivate, Close: hub.IsClose, LatestActive: t, Username: owner.Username}
		passback = append(passback, h)
	}
	return passback, nil
}
func (r *mutationResolver) AddJoinedHub(ctx context.Context, uid string, hubName string) ([]*Hub, error) {
	exist, err := models.HubExists(context.Background(), r.DB, hubName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if !exist {
		fmt.Println("hub is not exist!")
		return nil, err
	}

}
func (r *mutationResolver) ExitJoinedHub(ctx context.Context, uid string, hubName string) ([]*Hub, error) {
	panic("not implemented")
}
func (r *mutationResolver) PermitHubers(ctx context.Context, huberName string, hubName string) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) ChangeLogo(ctx context.Context, uid string, logo string) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) ToggleClosingHub(ctx context.Context, uid string, hubID string) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) ToggleHubStatus(ctx context.Context, uid string, hubID string) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) AddChatGif(ctx context.Context, uid string, gifData string, hubID string) (*Message, error) {
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
func (r *queryResolver) EnterHub(ctx context.Context, userID string, hubID string) ([]*Message, error) {
	panic("not implemented")
}
func (r *queryResolver) LoadMoreChat(ctx context.Context, userID string, hubID string, createAt time.Time) ([]*Message, error) {
	panic("not implemented")
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) ChatGifAdded(ctx context.Context, hubID string) (<-chan *Message, error) {
	panic("not implemented")
}
