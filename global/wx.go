package global

import (
	"fmt"
	"io/ioutil"
	"net/http"

	simplejson "github.com/bitly/go-simplejson"
)

func GetAccessToken() {

}

func GetOpenidByCode(code string) (*simplejson.Json, error) {
	var URL = "https://api.weixin.qq.com/sns/oauth2/access_token"
	appid := "wx6d9de6152a5a0aee"
	secret := "18bb93961be74a61bbd56325cba56b3a"

	grant_type := "authorization_code"

	apiUrl := URL + "?appid=" + appid + "&secret=" + secret + "&code=" + code + "&grant_type=" + grant_type
	fmt.Println(apiUrl)
	resp, err := http.Get(apiUrl)

	if err != nil {
		fmt.Println("调用API接口失败:", err)
		return nil, err
	}

	defer resp.Body.Close()

	contents, _ := ioutil.ReadAll(resp.Body)
	result, _ := simplejson.NewJson(contents)
	fmt.Println(result)
	return result, nil

}

func GetUserinfoByAccessToken(access_token, openid string) (*simplejson.Json, error) {
	var URL = "https://api.weixin.qq.com/sns/userinfo"

	apiUrl := URL + "?access_token=" + access_token + "&openid=" + openid
	fmt.Println(apiUrl)
	resp, err := http.Get(apiUrl)

	if err != nil {
		fmt.Println("调用API接口失败:", err)
		return nil, err
	}

	defer resp.Body.Close()

	contents, _ := ioutil.ReadAll(resp.Body)
	result, _ := simplejson.NewJson(contents)
	fmt.Println(result)
	return result, nil

}
