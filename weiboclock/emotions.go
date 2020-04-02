package weiboclock

import (
	"math/rand"
	"strings"
	"time"

	"github.com/pkg/errors"
)

var (
	// TollTails 标点小尾巴
	TollTails = []string{
		"!", "~", ".", "?",
	}

	// ClockEmoji 整点emoji
	ClockEmoji = []string{"🕛", "🕐", "🕑", "🕒", "🕓", "🕔", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚", "🕛"}

	// TextEmotions 颜文字表情
	TextEmotions = []string{
		"(°ー°〃)", "_(:з」∠)_ ", "o(*≧▽≦)ツ┏━┓", "๑乛◡乛๑ ", "(σ‘・д･)σ", "( ＿ ＿)ノ｜", "┑(￣Д ￣)┍",
		"(＃°Д°)", "(-ω- )", "(′・ω・`)", "( ^ω^)", "乀(ˉεˉ乀)", "ʕ•̫͡•ʕ•̫͡•ʔ•̫͡•ʔ", "(๑‾᷆д‾᷇๑)Fightᵎᵎ", "꒰｡⁻௰⁻｡꒱",
		"\\(╯-╰)/", "\\(▔▽▔)/", "^(oo)^", "(O^~^O)", "（╯＾╰）", "(*ﾟ∀ﾟ*)", "(శωశ)", "(*´∀`)~♥", "σ`∀´)σ",
		"(〃∀〃)", "(^_っ^)", "(｡◕∀◕｡)", "ヽ(✿ﾟ▽ﾟ)ノ", "ε٩(๑> ₃ <)۶з", "(σ′▽‵)′▽‵)σ", "σ ﾟ∀ ﾟ) ﾟ∀ﾟ)σ",
		"｡:.ﾟヽ(*´∀`)ﾉﾟ.:｡", "(✪ω✪)", "(∂ω∂)", "─=≡Σ((( つ•̀ω•́)つ", "(๑´ڡ`๑)", "(´▽`ʃ♡ƪ)", "(❛◡❛✿)", "(灬ºωº灬)",
		"(￣▽￣)/", "╰(*°▽°*)╯", "(๑•̀ㅂ•́)و✧", "( ^ω^)", "٩(｡・ω・｡)و", "( ～'ω')～", "(๑ơ ₃ ơ)♥", "(ﾉ◕ヮ◕)ﾉ*:･ﾟ✧",
		"o(☆Ф∇Ф☆)o", "(￫ܫ￩)", "(♥д♥)", "✧◝(⁰▿⁰)◜✧", "(ᗒᗨᗕ)/", "(=´ω`=)", "(｢･ω･)｢", "(*´д`)", "Σ>―(〃°ω°〃)♡→",
		"(▰˘◡˘▰)", "ヾ(´ε`ヾ)", "(っ●ω●)っ", "◥(ฅº￦ºฅ)◤", "ヽ( ° ▽°)ノ", "(　ﾟ∀ﾟ) ﾉ♡", "✧*｡٩(ˊᗜˋ*)و✧*｡",
		"⁽⁽◟(∗ ˊωˋ ∗)◞ ⁾⁾", "ヾ(´︶`*)ﾉ♬", "ヾ(*´∀ ˋ*)ﾉ", "(๑•̀ω•́)ノ", "ヾ (o ° ω ° O ) ノ゙ ", "╮(╯_╰)╭", "(๑•́ ₃ •̀๑)",
		"(´･_･`)", "(ㆆᴗㆆ)", "┐(´д`)┌", "( ˘･з･)", "( ´•︵•` )", "(｡ŏ_ŏ)", "(◞‸◟)", "( ˘•ω•˘ )", "(눈‸눈)", "(´･ω･`)",
		"(*´艸`*)", "(〃∀〃)", "(つд⊂)", "(๑´ㅂ`๑)", "ε٩(๑> ₃ <)۶з", "(๑´ڡ`๑)", "(灬ºωº灬)", "(๑• . •๑)",
		"(๑ơ ₃ ơ)♥", "(●｀ 艸´)", ",,Ծ‸Ծ,,", "(〃ﾟдﾟ〃)", "(๑´ㅁ`)", "(๑¯∀¯๑)", "(〃´∀｀)", "(⋟﹏⋞)", "(ノдT)",
		"(T_T)", "：ﾟ(｡ﾉω＼｡)ﾟ･｡", "(TдT)", "(☍﹏⁰)", "(╥﹏╥)", "｡ﾟ(ﾟ´ω`ﾟ)ﾟ｡", "இдஇ", "｡ﾟヽ(ﾟ´Д`)ﾉﾟ｡", "。･ﾟ･(つд`ﾟ)･ﾟ･",
		"(ﾟд⊙)", "(‘⊙д-)", "Σ(*ﾟдﾟﾉ)ﾉ", "(((ﾟДﾟ;)))", "(((ﾟдﾟ)))", "(☉д⊙)", "(|||ﾟдﾟ)",
		"(´⊙ω⊙`)", "ฅ(๑*д*๑)ฅ!!", "(゜ロ゜)", "(✘﹏✘ა)", "(✘Д✘๑ )", "(╬☉д⊙)", "(／‵Д′)／~ ╧╧", "(╯‵□′)╯︵┴─┴",
		"(◓Д◒)✄╰⋃╯", "(ﾒﾟДﾟ)ﾒ", "(`へ´≠)", "(#ﾟ⊿`)凸", "(╬▼дﾟ)", "(ᗒᗣᗕ)՞", "( ิ◕㉨◕ ิ)", "(❍ᴥ❍ʋ)", "(◕ܫ◕)", "(ΦωΦ)",
		"ก็ʕ•͡ᴥ•ʔ ก้", "(=´ω`=)", "(⁰⊖⁰)", "(=´ᴥ`)", "ฅ●ω●ฅ", "( ° ͜ʖ͡°)╭∩╮", "(⌐▀͡ ̯ʖ▀)", "(･ิω･ิ)", "ʕ•̀ω•́ʔ✧", "٩(♡ε♡ )۶", "٩(๑´3｀๑)۶",
		"(๑•̀ㅁ•́๑)✧", "•̀.̫•́✧", "⁽⁽٩(๑˃̶͈̀ ᗨ ˂̶͈́)۶⁾⁾", "( •̀ᄇ• ́)ﻭ✧", "(▭-▭)✧", "(▭-▭)✧",
		"(⸝⸝⸝ᵒ̴̶̥́ ⌑ ᵒ̴̶̣̥̀⸝⸝⸝)", "((̵̵́ ̆͒͟˚̩̭ ̆͒)̵̵̀)ﾞ", "o̖⸜((̵̵́ ̆͒͟˚̩̭ ̆͒)̵̵̀)⸝o̗", "٩(ˊᗜˋ*)و", "(∩ᵒ̴̶̷̤⌔ᵒ̴̶̷̤∩)", "₍₍ ◟꒰ ‾᷅д̈ ‾᷄ ╬꒱", "ଘ(੭ˊ꒳ˋ)੭✧", "ʕ·͡·̫͖ʕ⁎̯͡⁎ʔ",
	}

	// WeiboEmotions 微博官方表情，weiboclock Run方法调用时进行初始化
	WeiboEmotions = []string{}

	// Emotions 全部表情 TextEmotions + 初始化后的WeiboEmotions
	Emotions = []string{}
)

// PickOneEmotion 随机选择一个表情
func PickOneEmotion() string {
	rand.Seed(time.Now().Unix())
	return Emotions[rand.Intn(len(Emotions))]
}

// TollTail 随机获取标点小尾巴~
func TollTail(count int) string {
	rand.Seed(time.Now().Unix())
	tail := TollTails[rand.Intn(len(TollTails))]
	return strings.Repeat(tail, count)
}

// InitEmotions 初始化表情，返回表情总数
func (clock *WeiboClock) InitEmotions() (int, error) {
	// reset
	Emotions = []string{}
	WeiboEmotions = []string{}

	// 获取微博官方表情
	vb := clock.cronWeibo.WeiboClient()
	token := clock.cronWeibo.Token()
	language := "cnname"
	emotionType := "face"
	emotions, err := vb.Emotions(token.AccessToken, emotionType, language)
	if err != nil {
		return 0, errors.Wrap(err, "weiboclock InitWeiboEmotions Emotions error")
	}
	for _, emotion := range *emotions {
		WeiboEmotions = append(WeiboEmotions, emotion.Phrase)
		Emotions = append(Emotions, emotion.Phrase)
	}
	Emotions = append(Emotions, TextEmotions...)
	return len(Emotions), nil
}
