package service

import (
	"backend-svc-go/dao"
	"backend-svc-go/global"
	"backend-svc-go/model"
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/baidubce/bce-qianfan-sdk/go/qianfan"
)

const accessKey = "ALTAKHHpIkWJCbcSOh1IwTbXpw"
const secretKey = "7d727d160885493aa59085e37193310f"

func UserPhoneLogin(phone, invitor_code string) (*model.User, error) {
	user, err := dao.UserPhoneLogin(phone, invitor_code)
	return user, err
}

func GetUserByInvitor(invitor_code string) ([]model.UserRelation, error) {
	user, err := dao.GetUserByInvitor(invitor_code)
	return user, err
}

func GetUserWeixinInvitors(invitor, invitorType int) ([]model.UserWeixin, error) {
	user, err := dao.GetUserWeixinInvitors(invitor, invitorType)
	return user, err
}

func UpdateAvatar(user_id int, avatar string) error {
	return dao.UpdateAvatar(user_id, avatar)
}

func UserWeixinLogin(code, invitor_code string) (*model.User, error) {
	fmt.Println("code is " + code)
	fmt.Println("invitor code is " + invitor_code)
	res, err := global.GetOpenidByCode(code)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	openid := res.Get("openid").MustString()
	accessToken := res.Get("access_token").MustString()
	fmt.Println("openid is", openid)
	user, err := dao.GetUserByOpenid(openid, accessToken, invitor_code)

	return user, err
}

func SendMessage(message string) (*qianfan.ModelResponse, error) {
	qianfan.GetConfig().AccessKey = accessKey
	qianfan.GetConfig().SecretKey = secretKey

	chat := qianfan.NewChatCompletion(
		qianfan.WithModel("ERNIE-Bot"),
	)
	resp, err := chat.Do(
		context.TODO(),
		&qianfan.ChatCompletionRequest{
			Messages: []qianfan.ChatCompletionMessage{
				qianfan.ChatCompletionUserMessage(message),
			},
		},
	)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	fmt.Println(resp.Result)
	return resp, nil
}

func GetFunction(category string) ([]*model.Function, error) {
	functions, err := dao.GetFunction(category)
	return functions, err
}

func GetBaiduAccessToken() (*model.BaiduAccessToken, error) {
	return dao.GetBaiduAccessToken()
}

func GetCategory() ([]*model.Category, error) {
	return dao.GetCategory()
}

func GetAppByCategoryID(category_id int) ([]*model.App, error) {
	return dao.GetAppByCategoryID(category_id)
}

func GetHotApp() ([]*model.App, error) {
	return dao.GetHotApp()
}

func Predict() (*model.Predict, error) {

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	randomInt := r.Intn(80) + 30

	var predict = model.Predict{}

	predict.Money = strconv.Itoa(randomInt)

	l := []string{}

	l = append(l, "新的一年，新的开始。愿你不忘初心，继续前行，用努力和坚持迎接属于自己的辉煌！")

	l = append(l, "每一步的努力都不会白费，新的一年，愿你从容面对每一个挑战，勇敢追逐自己的梦想！")

	l = append(l, "新年，新的希望，新的目标。愿你在每一个日升月落间，不断积累，迎接更美好的未来！")

	l = append(l, "过去的一年可能有遗憾，但新的一年，属于你的一切还未开始。加油，努力，勇敢地迎接挑战！")

	l = append(l, "时光不负有心人，愿你在新的一年里，不畏风雨，向着梦想的方向不停奔跑，迎接属于你的精彩！")

	l = append(l, "新的一年，新的你。让每一个日出都充满动力，用热情和行动书写属于自己的精彩人生！")

	l = append(l, "新年是一张白纸，画上你希望的图景。愿你在这片白纸上涂抹出更加绚丽的色彩，成就属于自己的辉煌！")

	l = append(l, "无论过去一年多么不易，都请相信，新的一年，你会更加坚韧、自信，走得更远，飞得更高！")

	l = append(l, "新的一年，愿你用心去做每一件事，努力去追求每一个目标，让梦想在脚下生根发芽，开花结果！")

	l = append(l, "在这个全新的开始里，给自己一个全新的目标，给梦想一个新的起点，让我们一起在新的一年中拼搏奋斗！")

	randomIndex := rand.Intn(len(l))

	predict.Text = l[randomIndex]

	return &predict, nil

}

func ShenJia() (*model.Predict, error) {

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	randomInt := r.Intn(1000) + 500

	var predict = model.Predict{}

	predict.Money = strconv.Itoa(randomInt)

	l := []string{}

	l = append(l, "你是一个有商业头脑的人, 在赚钱的门道上很有自己的一套, 你能抓住成功的机会, 并能很好的处理各种问题, 相信你一定可以更加成功的")

	l = append(l, "每一步的努力都不会白费，新的一年，愿你从容面对每一个挑战，勇敢追逐自己的梦想！")

	l = append(l, "新年，新的希望，新的目标。愿你在每一个日升月落间，不断积累，迎接更美好的未来！")

	randomIndex := rand.Intn(len(l))

	predict.Text = l[randomIndex]

	return &predict, nil

}
