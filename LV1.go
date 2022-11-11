package main

import (
	"fmt"
	"strings"
)

var skill = map[int]string{
	1: "the world",
	2: "star platinum",
}
var words = map[int]string{}

func main() {
	var (
		n1 int
		n2 int
		s1 string
		s2 string
		s3 string
	)
	fmt.Println("技能：")
	for k, v := range skill {
		fmt.Println(k, v)
	}
	fmt.Println("是否有自定义技能和文字？")
	fmt.Scanln(&s1)
	if s1 == "是" {
		define()
		fmt.Println("文字：")
		for k, v := range words {
			fmt.Println(k, v)
		}
		fmt.Println("请输入文字编号")
		fmt.Scanln(&n1)
		s2 = words[n1]
		fmt.Println("技能：")
		for k, v := range skill {
			fmt.Println(k, v)
		}
		fmt.Println("请输入技能编号：")
		fmt.Scanln(&n2)
		s3 = skill[n2]
		ReleaseSkill(s3, func(s string) {
			fmt.Println(s2, s3)
		})
	} else {
		fmt.Println("请输入文字")
		fmt.Scanln(&s2)
		fmt.Println("请输入技能编号：")
		fmt.Scanln(&n2)
		s3 = skill[n2]
		ReleaseSkill(s3, func(s string) {
			fmt.Println(s2, s3)
		})
	}
}
func ReleaseSkill(s string, ReleaseSkillFunc func(string)) {
	ReleaseSkillFunc(s)
}
func define() {
	var (
		n1 int
		n2 int
		s1 string
		s2 string
	)
	fmt.Println("请分别输入技能编号和名称")
	fmt.Scanln(&n1)
	fmt.Scanln(&s1)
	skill[n1] = s1
	fmt.Println("请分别输入编号和文字")
	for {
		fmt.Scanln(&n2)
		fmt.Scanln(&s2)
		b := strings.Contains(s2, "逆天")
		if b == true {
			fmt.Println("文字内含敏感词”逆天“请重新输入")
		} else {
			break
		}
	}
	words[n2] = s2
}
