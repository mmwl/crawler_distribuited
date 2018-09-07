package parser

import (
	"regexp"
	"strconv"
	"awesomeProject/engine"
	"awesomeProject/model"
)

//预先编译正则表达式
var nameRe = regexp.MustCompile(`<h1 class="ceiling-name ib fl fs24 lh32 blue">(.+[^<])</h1>`)
var genderRe  = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">(.+)</span></td>`)
var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heightRe  = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var weigthRe  = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">(.+[0-9])KG</span></td>`)
var incomeRe  = regexp.MustCompile(`<td><span class="label">月收入：</span>(.+[^<])</td>`)
var marriageRe  = regexp.MustCompile(`<td><span class="label">婚况：</span>(.+[^<])</td>`)
var educationRe  = regexp.MustCompile(`<td><span class="label">学历：</span>(.+[^<])</td>`)
var occupationRe  = regexp.MustCompile(`<td><span class="label">职业： </span>(.+[^<])</td>`)
var hokouRe  = regexp.MustCompile(`<td><span class="label">籍贯：</span>(.+[^<])</td>`)
var xinzuoRe  = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">(.+[^<])</span></td>`)
var houseRe  = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">(.+[^<])</span></td>`)
var carRe  = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">(.+[^<])</span></td>`)

//人物解析器
func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{
		Name:extractString(contents,nameRe),
		Gender:extractString(contents,genderRe),
		Age: func() int {
			if a,err := strconv.Atoi(extractString(contents,ageRe)); err == nil {
				return a
			}
			return 0
		}(),
		Height: func() int {
			if h,err := strconv.Atoi(extractString(contents,heightRe)); err == nil {
				return h
			}
			return 0
		}(),
		Weigth: func() int {
			if w, err := strconv.Atoi(extractString(contents, weigthRe)); err == nil {
				return w
			}
			return 0
		}(),
		Income:extractString(contents,incomeRe),
		Marriage:extractString(contents,marriageRe),
		Education:extractString(contents,educationRe),
		Occupation:extractString(contents,occupationRe),
		Hokou:extractString(contents,hokouRe),
		Xinzuo:extractString(contents,xinzuoRe),
		House:extractString(contents,houseRe),
		Car:extractString(contents,carRe),
	}
	result := engine.ParseResult{
		Items:[]interface{}{profile},
	}
	return result
}

func extractString(contents []byte,re *regexp.Regexp) string{
	submatch := re.FindSubmatch(contents)
	if submatch != nil && len(submatch) >= 2 {
		return string(submatch[1])
	}
	return ""
}