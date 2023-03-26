package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/yoshikawa-river/ChatApp/adapter/http/response"
	"github.com/yoshikawa-river/ChatApp/usecase"
	"github.com/yoshikawa-river/ChatApp/usecase/models"
)

type userHandler struct {
	userInteractor usecase.IUserInteractor
}

func NewUserHandler(ui usecase.IUserInteractor) *userHandler {
	return &userHandler{userInteractor: ui}
}

func (uh *userHandler) createUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, read_err := io.ReadAll(r.Body)
	if read_err != nil {
		panic(read_err)
	}
	defer r.Body.Close()

	var u *models.User
	if err := json.Unmarshal(body, &u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	user, err := uh.userInteractor.Create(ctx, u)

	var rs response.SimpleResponse
	if err != nil {
		rs = response.SimpleResponse{
			Status:  http.StatusOK,
			Message: "failed to create user.",
			Error:   err.Error(),
		}
	} else {
		rs = response.SimpleResponse{
			Status: http.StatusOK,
			Detail: response.Detail{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
			},
		}
	}

	response.RespondJson(w, rs, rs.Status)
}
