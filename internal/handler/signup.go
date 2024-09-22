package handler

import (
	"encoding/json"
	"github.com/jexroid/gopi/api"
	"github.com/jexroid/gopi/pkg"
	"github.com/jexroid/gopi/pkg/utils"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func userinit(u *api.User) (*api.User, error) {
	u.UUID = pkg.Uuid()
	hashpass, err := utils.GenerateHash(u.Password)
	if err != nil {
		logrus.Error("[...1...] error in making hash ", err)
		utils.Telegram("[...1...] error in hashing")
		anotherTry, anotherError := utils.GenerateHash(u.Password)
		if anotherError != nil {
			utils.Telegram("[...2...] error in hashing")
			logrus.Error("[...2...] error in making hash ", err)
			return u, err
		}
		logrus.Info("[..1..] second time was successfully")
		u.Password = anotherTry
		return u, nil
	}
	u.Password = hashpass
	return u, nil
}

func (db Database) Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")

	var user api.User
	json.NewDecoder(r.Body).Decode(&user)
	if valid := utils.SignupCheker(user); valid != nil {
		logrus.Error(valid)
		w.WriteHeader(http.StatusNotAcceptable)
		jsonresponse, _ := json.Marshal(&api.RegisterResponse{Ok: false})
		w.Write(jsonresponse)
		return
	}

	_, err := userinit(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal error"))
		return
	}


	resault := db.DB.Table("User").Create(&user)
	logrus.Error(resault.Error)
	if nil != resault.Error {
		if resault.Error.Error() == `ERROR: duplicate key value violates unique constraint "uni_User_phone" (SQLSTATE 23505)` {
			jr, _ := json.Marshal(&api.RegisterResponse{
				Ok:        true,
				Token:     "",
				UserExist: true,
			})
			w.Write(jr)
			return
		}
		logrus.Error(resault.Error)
		//	TODO: Telegram Bot. (Read the telegram Code)
	}
	jwt := pkg.CreateToken(user.UUID, user.Phone)

	cookie := &http.Cookie{
		Name:     "identity",
		Value:    jwt,
		Path:     "/",
		HttpOnly: true,
		Secure:   os.Getenv("GO_ENV") == "production",
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, cookie)
	jsonresponse, _ := json.Marshal(&api.RegisterResponse{Ok: true, Token: jwt})
	w.Write(jsonresponse)
	return
}
