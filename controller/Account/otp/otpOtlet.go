package otp

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	user3 "hara-depo-proj/controller/Account/ownerotlet/user"
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/otp"
	"hara-depo-proj/redis"
	"hara-depo-proj/util/customResponse"
	"io/ioutil"
	"net/http"
)

func OtpUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := mobile.UserOtlet{}
	user2 := mobile.UserOtlet{}

	body, err1 := ioutil.ReadAll(r.Body)

	err1 = json.Unmarshal(body, &user)
	err2 := json.Unmarshal(body, &user2)
	_ = err2
	_ = err1

	otpnumber := otp.RandomStringOTP(6)
	IdOtlet := otp.RandomString(4)
	user.KodeUser = IdOtlet

	err := user3.Finduser(db, user.Hp)
	if err != nil {
		otp.OtpZensiva(user.Hp, otpnumber)
		redis.SaveData(user.KodeUser+"-register", otpnumber)
		response := map[string]interface{}{"KodeUser": user.KodeUser}
		customResponse.RespondJSON(w, http.StatusAccepted, response)

		return
	} else {
		if err := db.Where("hp= ?", user.Hp).Find(&user2).Error; err != nil {
			var resp = map[string]interface{}{"status": false, "message": "Something Wrong"}
			fmt.Println(resp)
		}
		redis.SaveData(user2.KodeUser+"-register", otpnumber)
		response := map[string]interface{}{"KodeUser": user2.KodeUser, "Message": "User Registered"}
		otp.OtpZensiva(user.Hp, otpnumber)
		customResponse.RespondJSON(w, http.StatusAccepted, response)
		return
	}
}
