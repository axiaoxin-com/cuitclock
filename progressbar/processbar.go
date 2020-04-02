package progressbar

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/schollz/progressbar/v2"
)

var saucerAndPaddings = [][]string{
	{"░", "▒"},
	{"⬛️", "⬜️"},
	{"❌", "⭕️"},
	{"⚫️", "⚪️"},
	{"🖤", "🤍"},
	{"🤍", "❤️"},
}

// ProgressBar 返回静态进度条
func ProgressBar(theme progressbar.Theme, width, total, current int) string {
	buf := strings.Builder{}
	bar := progressbar.NewOptions(
		total,
		progressbar.OptionSetTheme(theme),
		progressbar.OptionSetPredictTime(false),
		progressbar.OptionSetWidth(width),
		progressbar.OptionSetWriter(&buf),
	)

	if err := bar.Set(current); err != nil {
		log.Println("[ERROR] progressbar Set error", err)
	}
	return strings.TrimSpace(buf.String())
}

// DayProgressBar 今日使用进度
func DayProgressBar(t time.Time) string {
	// 使用day作为seed，一天内使用相同主题
	day := t.Day()
	// 0点进度条撑满，需要回退一天保持之前的进度
	if t.Hour() == 0 {
		day--
	}
	rand.Seed(int64(day))
	saucerAndPadding := saucerAndPaddings[rand.Intn(len(saucerAndPaddings))]
	theme := progressbar.Theme{Saucer: saucerAndPadding[0], SaucerHead: "", SaucerPadding: saucerAndPadding[1], BarStart: "", BarEnd: ""}

	hour := t.Hour()
	if hour == 0 {
		hour = 24
	}
	bar := ProgressBar(theme, 10, 24, hour)
	// 替换 [hour:24] 为 [hour/24]
	bar = strings.Replace(bar, ":", "/", 1)
	return bar
}

func isLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// YearProgressBar 今年使用进度
func YearProgressBar(t time.Time) string {
	// 使用year作为seed，一年内使用相同主题
	// 一年最后一天凌晨0点撑满进度条，回退一年
	year := t.Year()
	if int(t.Month()) == 1 && t.Day() == 1 && t.Hour() == 0 {
		year--
	}
	rand.Seed(int64(year))
	saucerAndPadding := saucerAndPaddings[rand.Intn(len(saucerAndPaddings))]
	theme := progressbar.Theme{Saucer: saucerAndPadding[0], SaucerHead: "", SaucerPadding: saucerAndPadding[1], BarStart: "", BarEnd: ""}

	dayCount := 365
	if isLeap(t.Year()) {
		dayCount = 366
	}

	bar := ProgressBar(theme, 15, dayCount, t.YearDay())
	bar = strings.Replace(bar, ":", "/", 1)
	return bar
}
