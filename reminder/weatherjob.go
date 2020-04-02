package reminder

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/axiaoxin-com/cronweibo"
	"github.com/axiaoxin-com/wttrin"
	"github.com/spf13/viper"
)

// 定时更新天气全局变量
func (r *Reminder) weatherJob() cronweibo.WeiboJob {
	return cronweibo.WeiboJob{
		Name:     "wttrin",
		Schedule: viper.GetString("reminder.wttrin_refresh_schedule"),
		Run:      r.wttrinRun,
	}
}

// 生成天气信息
func (r *Reminder) wttrinRun() (string, io.Reader) {
	// 默认每天7点半和17点半预报预报天气
	viper.SetDefault("reminder.wttrin_refresh_schedule", "0 30 7,17 * * *")
	lang := viper.GetString("reminder.wttrin_lang")
	loc := viper.GetString("reminder.wttrin_location")
	// 获取微博提醒昵称列表，空格分隔
	nicknameList := strings.Fields(viper.GetString("reminder.remind_list"))
	remindList := []string{}
	for _, nickname := range nicknameList {
		if !strings.HasPrefix(nickname, "@") {
			remindList = append(remindList, "@"+nickname)
		} else {
			remindList = append(remindList, nickname)
		}
	}
	remindStr := strings.Join(remindList, " ")
	// 获取天气图片
	log.Println("[DEBUG] wttrinRun start getting Image weather")
	img, err := wttrin.Image(lang, loc, "FpmM2")
	if err == nil {
		log.Println("[DEBUG] wttrinRun got the wttrin Image weather")
	} else {
		log.Println("[ERROR] wrttinRun get image weather error", err)
	}
	now := r.cronWeibo.Now().Format("15:04")
	text := fmt.Sprintf("%s %s 天气预报 %s", loc, now, remindStr)
	return text, img
}
