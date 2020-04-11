package otp

import (
	"fmt"
	"hara-depo-proj/config"
	"io/ioutil"
	"net/http"
)

func OtpZensiva(to string, randomString string) {

	baseConfig := &config.Configuration{}
	config.GetConfig(baseConfig)
	userkey := baseConfig.Zensiva.Userkey
	passkey := baseConfig.Zensiva.Passkey
	message := "Berikut%20kode%20OTP%20HARA%20Depo%20Anda:%20" + randomString + ".%20Jangan%20barikan%20kode%20OTP%20ini%20ke%20siapapun."
	url := "https://alpha.zenziva.net/apps/smsapi.php?userkey=" + userkey +
		"&passkey=" + passkey + "&nohp=" + to + "&pesan=" + message +
		"&type=otp"

	req, _ := http.NewRequest("POST", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	fmt.Println(string(body))

}
