package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

const topSizeNum = 10

var wordRe = regexp.MustCompile(`[a-zA-Zа-яА-Я_\-]+`)

type Word struct {
	word  string
	count int
}

func main() {
	text := `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы -- сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	-- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
громко:"Пу-ух! Пу-ух!"-- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни -- так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь -- теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

	results := Top10(text)

	fmt.Println(results)

}

func Top10(text string) (result []string) {
	wordsGroup := ParseText(text)

	result = GetTopWords(wordsGroup)
	return result
}

func ParseText(text string) map[string]int {
	words := strings.Split(text, " ")
	wordsGroup := make(map[string]int)

	for _, word := range words {
		if word == "" {
			continue
		}
		clearWord := wordRe.FindString(word)

		wordsGroup[strings.ToLower(clearWord)]++
	}
	return wordsGroup
}

func GetTopWords(wordsGroup map[string]int) []string {
	var words = make([]Word, 0, len(wordsGroup))
	for key, value := range wordsGroup {
		words = append(words, Word{word: key, count: value})
	}

	sort.Slice(words, func(i, j int) bool {
		return words[i].count > words[j].count
	})

	var result = make([]string, 0, topSizeNum)
	var length int
	if length = topSizeNum; len(words) < topSizeNum {
		length = len(words)
	}
	for _, word := range words[0:length] {
		result = append(result, word.word)
	}
	return result
}
