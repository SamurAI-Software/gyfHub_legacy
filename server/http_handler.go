package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	grapql "github.com/SamurAI-Software/gyfHub/grapql"
	"github.com/SamurAI-Software/gyfHub/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/gorilla/websocket"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/gofrs/uuid"
)

// SetupNewRouter is for control all the api router
func (c *API) SetupNewRouter() chi.Router {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	//public route
	r.Group(func(r chi.Router) {
		r.Post("/api/signup", c.SignUpHandler)
		r.Post("/api/signin", c.SignInHandler)
		r.Get("api/verification/{verifyToken}", c.VerifyHandler)
		r.Get("/api/signout", SignOutHandler)
	})

	//private route
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(auther()))

		r.Use(jwtauth.Authenticator)

		r.Handle("/", handler.Playground("GraphQL playground", "/query"))
		r.Handle("/query", handler.GraphQL(
			grapql.NewExecutableSchema(
				grapql.Config{
					Resolvers: &grapql.Resolver{
						DB: c.DBDriver.DB,
					},
					Directives: grapql.DirectiveRoot{
						User: func(ctx context.Context, obj interface{}, next graphql.Resolver, userID string) (res interface{}, err error) {
							return next(context.WithValue(ctx, "userID", userID))
						},
					},
				},
			),
			handler.WebsocketUpgrader(websocket.Upgrader{
				CheckOrigin: func(r *http.Request) bool { return true },
			}),
		))

	})

	return r
}

//VerifyHandler is for user verify their account
func (c *API) VerifyHandler(w http.ResponseWriter, r *http.Request) {
	verifyToken := chi.URLParam(r, "verifyToken")
	exist, err := models.Users(qm.Where("verify_token = ?", verifyToken)).Exists(context.Background(), c.DBDriver.DB)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !exist {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := models.Users(qm.Select(models.UserColumns.Verified), qm.Where("verify_token = ?", verifyToken)).One(context.Background(), c.DBDriver.DB)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	user.Verified = true
	_, err = user.Update(context.Background(), c.DBDriver.DB, boil.Whitelist(models.UserColumns.Verified))
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// SignUpHandler is for handling sign up action
func (c *API) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	user := &SignUpInput{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Println(err)
		return
	}

	// check username exist
	alreadyExist, err := models.Users(qm.Where("username=?", user.Username)).Exists(context.Background(), c.DBDriver.DB)
	if err != nil {
		fmt.Println(err)
		errResponse(w, err)
		return
	}

	if alreadyExist {
		fmt.Println(ErrUsernameExistIssue)
		errResponse(w, ErrUsernameExistIssue)
		return
	}

	// check email
	alreadyExist, err = models.Users(qm.Where("email=?", user.Email)).Exists(context.Background(), c.DBDriver.DB)
	if err != nil {
		fmt.Println(err)
		errResponse(w, err)
		return
	}

	if alreadyExist {
		fmt.Println(ErrEmailExistIssue)
		errResponse(w, ErrEmailExistIssue)
		return
	}

	// check mobile
	alreadyExist, err = models.Users(qm.Where("mobile=?", user.Mobile)).Exists(context.Background(), c.DBDriver.DB)
	if err != nil {
		fmt.Println(err)
		errResponse(w, err)
		return
	}

	if alreadyExist {
		fmt.Println(ErrMobileExistIssue)
		errResponse(w, ErrMobileExistIssue)
		return
	}

	verifyToken, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
		errResponse(w, err)
		return
	}

	resetToken, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
		errResponse(w, err)
		return
	}

	uid, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
		errResponse(w, err)
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		errResponse(w, err)
		return
	}

	newUser := &models.User{
		ID:             uid.String(),
		Username:       user.Username,
		Email:          user.Email,
		Mobile:         user.Mobile,
		PasswordHash:   hashedPass,
		VerifyToken:    verifyToken.String(),
		ResetPassToken: resetToken.String(),
		Verified:       false,
		Avatar:         []byte(""),
	}

	err = newUser.Insert(context.Background(), c.DBDriver.DB, boil.Infer())
	if err != nil {
		fmt.Println(err)
		errResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&SuccessResponse{Success: true, Message: "signed up"})
}

// SignInHandler is for handling sign in action
func (c *API) SignInHandler(w http.ResponseWriter, r *http.Request) {
	user := &SignInInput{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Println(err)
		return
	}

	// check user exist by username
	userObj, err := models.Users(qm.Where("username=?", user.Username)).One(context.Background(), c.DBDriver.DB)
	if err != nil {
		fmt.Println(err)
		errResponse(w, err)
		return
	}

	// check user is verified
	if !userObj.Verified {
		fmt.Println(ErrVerifiedIssue)
		errResponse(w, ErrVerifiedIssue)
		return
	}

	err = bcrypt.CompareHashAndPassword(userObj.PasswordHash, []byte(user.Password))
	if err != nil {
		fmt.Println(err)
		errResponse(w, err)
		return
	}

	// Setup Cookie
	expiration := time.Now().Add(7 * 24 * time.Hour)
	_, jwt, err := auther().Encode(jwt.MapClaims{"id": userObj.ID})
	if err != nil {
		errResponse(w, err)
		return
	}
	cookie := http.Cookie{Name: "jwt", Value: jwt, Expires: expiration, HttpOnly: true, Path: "/", Secure: false}
	http.SetCookie(w, &cookie)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&SuccessResponse{Success: true, Message: "signed in"})
}

// SignOutHandler is for sign out
func SignOutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "jwt",
		Path:     "/",
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&SuccessResponse{Success: true, Message: "signed out"})
}

func errResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&SuccessResponse{Success: false, Message: err.Error()})
}
