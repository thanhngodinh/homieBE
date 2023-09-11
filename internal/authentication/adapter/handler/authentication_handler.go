package handler

import (
	"context"
	"net/http"
	"reflect"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/core-go/core"

	"hostel-service/internal/authentication/domain"
	"hostel-service/internal/authentication/service"
)

func NewAuthenticationHandler(find func(context.Context, interface{}, interface{}, int64, ...int64) (int64, string, error), service service.AuthenticationService, status core.StatusConfig, logError func(context.Context, string, ...map[string]interface{}), validate func(context.Context, interface{}) ([]core.ErrorMessage, error), action *core.ActionConfig) *HttpAuthenticationHandler {
	modelType := reflect.TypeOf(domain.User{})
	params := core.CreateParams(modelType, &status, logError, validate, action)
	return &HttpAuthenticationHandler{service: service, Params: params}
}

type HttpAuthenticationHandler struct {
	service service.AuthenticationService
	*core.Params
}

// Login godoc
// @Summary      Login
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        User body      domain.User true "User to login"
// @Success      201  {object}  util.Response{value=int,data=domain.User}
// @Failure      400  {string}  string      "Invalid character 's' looking for beginning of value"
// @Failure      500  {string}  string      "Internal Server Error"
// @Router       /auth/login [post]
func (h *HttpAuthenticationHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	er1 := core.Decode(w, r, &credentials)
	if er1 == nil {
		user, er2 := h.service.GetByUsername(r.Context(), credentials.Username)
		if er2 != nil {
			core.HasError(w, r, nil, er2, h.Status.Error, h.Error, h.Log, h.Resource, *h.Action.Load)
			return
		}
		if user == nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)) != nil {
			core.Respond(w, r, http.StatusNotFound, nil, nil, nil, nil)
		}
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Audience:  user.Id,
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		})
		token, er3 := claims.SignedString(domain.SECRET_KEY)
		if er3 != nil {
			core.HasError(w, r, nil, er3, h.Status.Error, h.Error, h.Log, h.Resource, *h.Action.Load)
		}
		accessToken := domain.AccessToken{
			TokenString: token,
		}
		core.RespondModel(w, r, accessToken, nil, nil, nil)
	}
}

// CreateUser godoc
// @Summary      Create a User
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        User body     domain.User  true "User to create"
// @Success      201  {object} util.Response{value=int,data=domain.User}
// @Failure      400  {string} string       "Invalid character 's' looking for beginning of value"
// @Failure      500  {string} string       "Internal Server Error"
// @Router       /auth/register [post]
func (h *HttpAuthenticationHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	er1 := core.Decode(w, r, &user)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &user)
		if !core.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, h.Log, h.Resource, h.Action.Create) {
			user.Id = uuid.New().String()
			hashedPassword, er3 := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			if er3 != nil {
				core.HasError(w, r, nil, er3, h.Status.Error, h.Error, h.Log, h.Resource, *h.Action.Load)
				return
			}
			user.Password = string(hashedPassword)
			now := time.Now()
			user.CreatedAt = &now
			res, er3 := h.service.Create(r.Context(), &user)
			core.AfterCreated(w, r, &user, res, er3, h.Status, h.Error, h.Log, h.Resource, h.Action.Create)
		}
	}
}