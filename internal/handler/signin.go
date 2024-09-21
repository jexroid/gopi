package handler

import (
	"encoding/json"
	"github.com/jexroid/gopi/api"
	"github.com/jexroid/gopi/pkg"
	"github.com/jexroid/gopi/pkg/utils"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (db Database) Signin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")

	var form api.LoginRequest
	json.NewDecoder(r.Body).Decode(&form)

	if valid := utils.SigninCheker(form); valid != nil {
		logrus.Info(valid)
		w.WriteHeader(http.StatusNotAcceptable)
		jsonresponse, _ := json.Marshal(&api.LoginResponse{Ok: false})
		w.Write(jsonresponse)
		return
	}

	var user api.User
	resault := db.DB.Table("User").First(&user, "phone = ?", strconv.Itoa(form.Phone))
	if nil != resault.Error {
		if resault.Error.Error() == "record not found" {
			jr, _ := json.Marshal(&api.RegisterResponse{
				Ok:        true,
				Token:     "",
				UserExist: true,
			})
			w.Write(jr)
			return
		}
		//	TODO: add database error handling
	}

	isPassValid, _ := utils.ComparePasswordAndHash(form.Password, user.Password)
	logrus.Info(isPassValid)

	if isPassValid {
		jwt := pkg.CreateToken(user.UUID, user.Phone)
		cookie := &http.Cookie{
			Name:  "identity",
			Value: jwt,
			Path:  "/",
			//Domain:   ""
			HttpOnly: true,
			//Secure:   false,
			SameSite: http.SameSiteLaxMode,
		}

		http.SetCookie(w, cookie)

		jsonresponse, _ := json.Marshal(&api.LoginResponse{Ok: true, Valid: true, Token: jwt})
		w.Write(jsonresponse)
		return
	}

	response := &api.LoginResponse{}
	jsonresponse, _ := json.Marshal(response)

	w.Write(jsonresponse)
	return
}
