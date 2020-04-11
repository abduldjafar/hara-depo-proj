package otp

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	user3 "hara-depo-proj/controller/Account/ownerotlet/user"
	"hara-depo-proj/model"
	"hara-depo-proj/redis"
	"hara-depo-proj/util"
	"io/ioutil"
	"net/http"
	"strconv"
)

func cleanRedis(user model.OtpCode) {
	redis.DeleteDataFromRedis(user.KodeUser + "-register")
	redis.DeleteDataFromRedis(user.KodeUser + "-register-login-count")
}

func OtpUserConfirm(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	user := model.OtpCode{}
	userOtlet := model.UserOtlet{}

	body, err1 := ioutil.ReadAll(r.Body)

	err1 = json.Unmarshal(body, &user)
	err2 := json.Unmarshal(body, &userOtlet)
	_ = err2
	_ = err1

	defer r.Body.Close()

	err := redis.GetData(user.KodeUser + "-register")
	if err != nil {
		fmt.Println(err.Error())
	}

	values := redis.GetDataFromRedis(user.KodeUser + "-register")
	fmt.Println(user.KodeOtp, values)

	if values == user.KodeOtp {
		err := user3.FindUserStatus(db, user.Hp, "Enabled")
		if err != nil {
			response := model.ResponseOtpNull{}
			err2 := user3.FindUserStatus(db, user.Hp, "Disabled")
			if err2 != nil {
				if err := db.Save(&userOtlet).Error; err != nil {
					util.RespondError(w, http.StatusInternalServerError, err.Error())
					return
				}
				redis.DeleteDataFromRedis(user.KodeUser + "-register-login-count")
				response.Messages = "register"
				response.User = nil
				response.TokoUser = nil
				util.RespondJSON(w, 200, response)
			} else {
				userOtlet := model.UserOtlet{}
				toko := model.Toko{}
				response := model.ResponseOtp{}

				if err := db.Where("user_otlet.kode_user=? AND user_otlet.hp=?", user.KodeUser, user.Hp).Find(&userOtlet).Error; err != nil {
					util.RespondError(w, http.StatusInternalServerError, err.Error())
					return
				}
				if err := db.Where("toko.kode_user=?", user.KodeUser).Find(&toko).Error; err != nil {
					util.RespondError(w, http.StatusInternalServerError, err.Error())
					return
				}
				response.Messages = "home"
				response.User = userOtlet
				response.TokoUser = toko
				redis.DeleteDataFromRedis(user.KodeUser + "-register-login-count")
				util.RespondJSON(w, 200, response)
			}
		} else {
			cleanRedis(user)
			userOtlet := model.UserOtlet{}
			toko := model.Toko{}
			response := model.ResponseOtp{}

			if err := db.Where("user_otlet.kode_user=? AND user_otlet.hp=?", user.KodeUser, user.Hp).Find(&userOtlet).Error; err != nil {
				util.RespondError(w, http.StatusInternalServerError, err.Error())
				return
			}
			if err := db.Where("toko.kode_user=?", user.KodeUser).Find(&toko).Error; err != nil {
				util.RespondError(w, http.StatusInternalServerError, err.Error())
				return
			}
			response.Messages = "home"
			response.User = userOtlet
			response.TokoUser = toko
			util.RespondJSON(w, 200, response)
		}

	} else {
		errLogincount := redis.GetData(user.KodeUser + "-register-login-count")
		if errLogincount != nil {
			redis.SaveData(user.KodeUser+"-register-login-count", strconv.Itoa(1))
			util.RespondError(w, http.StatusConflict, "Login failed")
		} else {
			value := redis.GetDataFromRedis(user.KodeUser + "-register-login-count")
			fmt.Println(value)
			countFromRedis, _ := strconv.Atoi(value)
			countFromRedis = 1 + countFromRedis
			redis.SaveData(user.KodeUser+"-register-login-count", strconv.Itoa(countFromRedis))

			if countFromRedis > 3 {
				util.RespondError(w, http.StatusInternalServerError, "PercobaanKonfirmasi sudah mencapai Batas Maksimal")
				redis.DeleteDataFromRedis(user.KodeUser + "-register-login-count")
			} else {
				//redis.DeleteDataFromRedis(user.KodeUser + "-register-login-count")
				util.RespondError(w, http.StatusConflict, "Login failed")
			}
		}
	}
}
