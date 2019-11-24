package grapql

import (
	context "context"
	"crypto/sha1"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/SamurAI-Software/gyfHub/models"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/crypto/bcrypt"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

//Resolver is customized data struct
type Resolver struct {
	DB           *sqlx.DB
	EnterRequest map[string]map[string][]string
	mu           sync.Mutex
	Hubs         map[string]*Chathub
}

// Chathub is for tracking and subscribe the room
type Chathub struct {
	Name      string
	Messages  []Message
	Observers map[string]struct {
		Username string
		Message  chan *Message
	}
}

func getUserID(ctx context.Context) string {
	if userID, ok := ctx.Value("userID").(string); ok {
		return userID
	}
	return ""
}

// Follower is struct for follower resolver
func (r *Resolver) Follower() FollowerResolver {
	return &followerResolver{r}
}

// Hub is for Hub resolver struct
func (r *Resolver) Hub() HubResolver {
	return &hubResolver{r}
}

// Message is for message resolver struct
func (r *Resolver) Message() MessageResolver {
	return &messageResolver{r}
}

// Mutation is for mutation resolver struct
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

// Query is for Query resolver struct
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

// Subscription is for subscription resolver struct
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

// ImageToID is for generate image file into a uniqe hashed id
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
	t := time.Now()
	hub := &models.Hub{ID: hubName, Logo: []byte(""), LastActive: t, UserID: uid, IsPrivate: status, IsClose: false}
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

		h := &Hub{ID: hub.ID, Logo: string(hub.Logo), Hubers: int(hubers), Status: hub.IsPrivate, Close: hub.IsClose, LatestActive: hub.LastActive, Username: owner.Username}
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
		return nil, errors.New("HUB_NOT_EXIST_ISSUE")
	}

	exist, err = models.HubUsers(qm.Where("hub_id = ?", hubName), qm.And("user_id = ?", uid)).Exists(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if !exist {
		targetHub, err := models.Hubs(qm.Where("id = ?", hubName)).One(context.Background(), r.DB)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		// if it is public, add it to user hub list
		if !targetHub.IsPrivate {
			huID, err := uuid.NewV4()
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			hu := &models.HubUser{ID: huID.String(), HubID: targetHub.ID, UserID: uid}
			err = hu.Insert(context.Background(), r.DB, boil.Infer())
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
		} else {
			// deal with go map.......
			if m, exist := r.EnterRequest[targetHub.UserID]; exist {
				if list, exist := m[targetHub.ID]; exist {
					if _, exist := contains(list, uid); !exist {
						list = append(list, uid)
						m[targetHub.ID] = list
						r.EnterRequest[targetHub.UserID] = m
					}
				} else {
					m[targetHub.ID] = []string{uid}
					r.EnterRequest[targetHub.UserID] = m
				}
			} else {
				m := make(map[string][]string)
				m[targetHub.ID] = []string{uid}
				r.EnterRequest[targetHub.ID] = m
			}

		}
	}

	hubs, err := models.HubUsers(qm.Where("user_id = ?", uid)).All(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var passback []*Hub
	for _, v := range hubs {
		hub, err := v.Hub().One(context.Background(), r.DB)
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
		h := &Hub{ID: hub.ID, Logo: string(hub.Logo), Hubers: int(hubers), Status: hub.IsPrivate, Close: hub.IsClose, LatestActive: hub.LastActive, Username: owner.Username}
		passback = append(passback, h)
	}
	return passback, nil
}

func contains(list []string, item string) (int, bool) {
	for i, v := range list {
		if v == item {
			return i, true
		}
	}
	return -1, false
}

func (r *mutationResolver) ExitJoinedHub(ctx context.Context, uid string, hubName string) ([]*Hub, error) {
	exist, err := models.HubExists(context.Background(), r.DB, hubName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if !exist {
		fmt.Println("hub is not exist!")
		return nil, err
	}
	exist, err = models.HubUsers(qm.Where("hub_id = ?", hubName), qm.And("user_id = ?", uid)).Exists(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if !exist {
		fmt.Println("You are not in the hub!")
		return nil, err
	}
	_, err = models.HubUsers(qm.Where("hub_id = ?", hubName), qm.And("user_id = ?", uid)).DeleteAll(context.Background(), r.DB)
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
		hub, err := v.Hub().One(context.Background(), r.DB)
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
		h := &Hub{ID: hub.ID, Logo: string(hub.Logo), Hubers: int(hubers), Status: hub.IsPrivate, Close: hub.IsClose, LatestActive: hub.LastActive, Username: owner.Username}
		passback = append(passback, h)
	}
	return passback, nil

}

func (r *mutationResolver) PermitHubers(ctx context.Context, uid string, huberName string, hubName string) (bool, error) {
	//check user is the owner
	targetHub, err := models.FindHub(context.Background(), r.DB, hubName)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	if targetHub.UserID != uid {
		fmt.Println("Owner only")
		return false, errors.New("NOT_OWENER_ISSUE")
	}

	huber, err := models.Users(qm.Where("username = ?", huberName)).One(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	exist, err := models.HubUsers(qm.Where("user_id = ?", huber.ID), qm.And("hub_id = ?", targetHub.ID)).Exists(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	if exist {
		fmt.Println("The huber is already permitted")
		return false, errors.New("ALREADY_PERMITTED_ISSUE")
	}

	if m, exist := r.EnterRequest[uid]; exist {
		if list, exist := m[hubName]; exist {
			if i, exist := contains(list, huber.ID); exist {
				if len(list) == 1 {
					delete(m, hubName)
				} else if len(list) > 1 {
					list = append(list[:i], list[i+1:]...)
					m[hubName] = list
				}
				r.EnterRequest[uid] = m
			} else {
				fmt.Println("the request is not on the request list")
				return false, errors.New("NOT_ON_LIST_ISSUE")
			}
		} else {
			fmt.Println("the request is not on the request list")
			return false, errors.New("NOT_ON_LIST_ISSUE")
		}

	} else {
		fmt.Println("the request is not on the request list")
		return false, errors.New("NOT_ON_LIST_ISSUE")
	}

	id, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	hu := &models.HubUser{ID: id.String(), HubID: hubName, UserID: huber.ID}
	err = hu.Insert(context.Background(), r.DB, boil.Infer())
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) ChangeLogo(ctx context.Context, uid string, hubName string, logo string) (bool, error) {
	//check the person is the owner
	targetHub, err := models.FindHub(context.Background(), r.DB, hubName)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	if targetHub.UserID != uid {
		fmt.Println("invaliad request")
		return false, errors.New("NOT_HUB_OWNER_ISSUE")
	}

	targetHub.Logo = []byte(logo)
	_, err = targetHub.Update(context.Background(), r.DB, boil.Whitelist(models.HubColumns.Logo))
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) ToggleClosingHub(ctx context.Context, uid string, hubID string) (bool, error) {
	// check user is the owner
	targetHub, err := models.FindHub(context.Background(), r.DB, hubID)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	if targetHub.UserID != uid {
		fmt.Println("invaliad request")
		return false, errors.New("NOT_HUB_OWNER_ISSUE")
	}

	targetHub.IsClose = !targetHub.IsClose
	_, err = targetHub.Update(context.Background(), r.DB, boil.Whitelist(models.HubColumns.IsClose))
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) ToggleHubStatus(ctx context.Context, uid string, hubID string) (bool, error) {
	targetHub, err := models.FindHub(context.Background(), r.DB, hubID)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	if targetHub.UserID != uid {
		fmt.Println("invaliad request")
		return false, errors.New("NOT_HUB_OWNER_ISSUE")
	}

	targetHub.IsPrivate = !targetHub.IsPrivate
	_, err = targetHub.Update(context.Background(), r.DB, boil.Whitelist(models.HubColumns.IsPrivate))
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) AddChatGif(ctx context.Context, uid string, gifData string, hubID string) (*Message, error) {
	user, err := models.FindUser(context.Background(), r.DB, uid)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	r.mu.Lock()
	hub := r.Hubs[hubID]
	if hub == nil {
		hub = &Chathub{
			Name: hubID,
			Observers: map[string]struct {
				Username string
				Message  chan *Message
			}{},
		}
		r.Hubs[hubID] = hub
	}
	r.mu.Unlock()

	// check gif data
	gifID := ImageToID(gifData)
	exist, err := models.GifExists(context.Background(), r.DB, gifID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if !exist {
		newGif := &models.Gif{ID: gifID, GifData: []byte(gifData)}
		err = newGif.Insert(context.Background(), r.DB, boil.Infer())
	}

	// store message into database
	id, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	t := time.Now()
	chatMsg := &models.ChatMSG{ID: id.String(), HubID: hubID, GifID: gifID, CreateAt: t, UserID: uid}
	err = chatMsg.Insert(context.Background(), r.DB, boil.Infer())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	message := Message{
		ID:       id.String(),
		GifID:    gifID,
		Gif:      gifData,
		CreateAt: t,
		Username: user.Username,
	}
	hub.Messages = append(hub.Messages, message)

	r.mu.Lock()
	for _, observer := range hub.Observers {
		if observer.Username == "" || observer.Username == message.Username {
			observer.Message <- &message
		}
	}
	r.mu.Unlock()

	return &message, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetUserProfile(ctx context.Context, uid string) (*Profile, error) {
	user, err := models.FindUser(context.Background(), r.DB, uid)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	c, err := models.UserCategories(qm.Where("user_id = ?", uid)).All(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	categories := []string{}
	for _, v := range c {
		categories = append(categories, v.Category)
	}
	passback := &Profile{Email: user.Email, Mobile: user.Mobile, Username: user.Username, Avatar: string(user.Avatar), Categories: categories}
	return passback, nil
}
func (r *queryResolver) GetUserCategory(ctx context.Context, uid string, cat string) ([]*Gif, error) {
	gifs, err := models.UserFavorites(qm.Where("user_id = ?", uid), qm.And("category = ?", cat)).All(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	passback := []*Gif{}
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
func (r *queryResolver) GetUserFollower(ctx context.Context, uid string) ([]*Follower, error) {
	userFollower, err := models.Followers(qm.Select(models.FollowerColumns.FollowerID), qm.Where("user_id = ?", uid)).All(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	passback := []*Follower{}
	for _, v := range userFollower {
		follower, err := v.Follower().One(context.Background(), r.DB)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		f := &Follower{Username: follower.Username, Isfollowing: true}
		passback = append(passback, f)
	}
	return passback, nil
}
func (r *queryResolver) GetFollowerUser(ctx context.Context, uid string) ([]*Follower, error) {
	followerUser, err := models.Followers(qm.Where("follower_id = ?", uid)).All(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// get user following id
	userFollower, err := models.Followers(qm.Select(models.FollowerColumns.FollowerID), qm.Where("user_id = ?", uid)).All(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	userFollowerList := []string{}
	for _, v := range userFollower {
		userFollowerList = append(userFollowerList, v.FollowerID)
	}

	passback := []*Follower{}
	for _, v := range followerUser {
		follower, err := v.User().One(context.Background(), r.DB)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		_, isFollowing := contains(userFollowerList, follower.ID)
		f := &Follower{Username: follower.Username, Isfollowing: isFollowing}
		passback = append(passback, f)
	}

	return passback, nil

}
func (r *queryResolver) GetUserHub(ctx context.Context, uid string) ([]*Hub, error) {
	hubs, err := models.HubUsers(qm.Where("user_id = ?", uid)).All(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	passback := []*Hub{}
	for _, v := range hubs {
		hub, err := v.Hub().One(context.Background(), r.DB)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		hubers, err := models.HubUsers(qm.Where("hub_id = ?", hub.ID)).Count(context.Background(), r.DB)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		owner, err := hub.User(qm.Select(models.UserColumns.Username)).One(context.Background(), r.DB)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		h := &Hub{ID: hub.ID, Logo: string(hub.Logo), Hubers: int(hubers), Status: hub.IsPrivate, Close: hub.IsClose, LatestActive: hub.LastActive, Username: owner.Username}
		passback = append(passback, h)
	}

	return passback, nil
}
func (r *queryResolver) GetCustGif(ctx context.Context, uid string) ([]*Gif, error) {
	custGifs, err := models.UserGifs(qm.Where("user_id = ?", uid)).All(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	passback := []*Gif{}
	for _, v := range custGifs {
		custGif, err := v.Gif().One(context.Background(), r.DB)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		g := &Gif{ID: custGif.ID, Gif: string(custGif.GifData)}
		passback = append(passback, g)
	}
	return passback, nil
}
func (r *queryResolver) GetGlobalHubs(ctx context.Context, search string) ([]*Hub, error) {
	panic("not implemented")
}
func (r *queryResolver) GetOtherUser(ctx context.Context, username string) (*Follower, error) {
	panic("not implemented")
}
func (r *queryResolver) EnterHub(ctx context.Context, uid string, hubID string) ([]*Message, error) {
	exist, err := models.HubUsers(qm.Where("user_id = ?", uid), qm.And("hub_id = ?", hubID)).Exists(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if !exist {
		fmt.Println("Unauthorized access prohibited")
		return nil, errors.New("UNAUTHORIZED_ACCESS")
	}

	chatMsgs, err := models.ChatMSGS(qm.Where("hub_id = ?", hubID), qm.Limit(10), qm.OrderBy(models.ChatMSGColumns.CreateAt)).All(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	passback := []*Message{}
	for _, v := range chatMsgs {
		gif, err := v.Gif().One(context.Background(), r.DB)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		user, err := v.User().One(context.Background(), r.DB)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		cm := &Message{ID: v.ID, GifID: v.GifID, Gif: string(gif.GifData), CreateAt: v.CreateAt, Username: user.Username}
		passback = append(passback, cm)
	}
	return passback, nil
}
func (r *queryResolver) LoadMoreChat(ctx context.Context, uid string, hubID string, createAt time.Time) ([]*Message, error) {
	exist, err := models.HubUsers(qm.Where("user_id = ?", uid), qm.And("hub_id = ?", hubID)).Exists(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if !exist {
		fmt.Println("Unauthorized access prohibited")
		return nil, errors.New("UNAUTHORIZED_ACCESS")
	}

	chatMsgs, err := models.ChatMSGS(qm.Where("hub_id = ?", hubID), qm.And("create_at < ?", createAt), qm.Limit(10), qm.OrderBy(models.ChatMSGColumns.CreateAt)).All(context.Background(), r.DB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	passback := []*Message{}
	for _, v := range chatMsgs {
		gif, err := v.Gif().One(context.Background(), r.DB)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		user, err := v.User().One(context.Background(), r.DB)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		cm := &Message{ID: v.ID, GifID: v.GifID, Gif: string(gif.GifData), CreateAt: v.CreateAt, Username: user.Username}
		passback = append(passback, cm)
	}
	return passback, nil

}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) ChatGifAdded(ctx context.Context, hubID string) (<-chan *Message, error) {
	r.mu.Lock()
	hub := r.Hubs[hubID]
	if hub == nil {
		hub = &Chathub{
			Name: hubID,
			Observers: map[string]struct {
				Username string
				Message  chan *Message
			}{},
		}
		r.Hubs[hubID] = hub
	}
	r.mu.Unlock()

	id, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	events := make(chan *Message, 1)

	go func() {
		<-ctx.Done()
		r.mu.Lock()
		delete(hub.Observers, id.String())
		r.mu.Unlock()
	}()

	r.mu.Lock()
	hub.Observers[id.String()] = struct {
		Username string
		Message  chan *Message
	}{Username: getUserID(ctx), Message: events}
	r.mu.Unlock()

	return events, nil
}
