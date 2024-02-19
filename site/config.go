package site

import (
	"strings"
	"time"
)

var Config *config

type Keywords []string

func (k *Keywords) String() string { return strings.Join(*k, ",") }

type searchEngineConnection struct {
	Google string
}

type config struct {
	Port                   uint32
	Domain                 string
	Author                 string
	Title                  string
	Description            string
	Keywords               *Keywords
	DatePublished          time.Time
	DateModified           time.Time
	PhoneNumber            string
	SearchEngineConnection *searchEngineConnection
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

func Init() {
	c := &config{}
	c.Port = uint32(8019)
	c.Domain = "colagom.com"
	c.Author = "콜라곰"
	c.Title = "콜라곰의 강남유흥 여행"
	c.Description = "콜라곰과 함께 떠나는 강남의 유흥주점의 가격, 시스템, 위치정보 안내. 가라오케, 셔츠룸, 하이퍼블릭, 레깅스룸, 쩜오, 호빠, 클럽의 모든 정보"
	k := Keywords([]string{"콜라곰의 강남유흥 여행", "콜라곰", "강남유흥", "유흥", "유흥주점", "강남유흥주점", "강남룸빵", "룸빵", "가라오케", "셔츠룸", "하이퍼블릭", "레깅스룸", "쩜오", "호빠", "클럽"})
	c.Keywords = &k
	c.DatePublished = date(2023, 9, 6)
	c.DateModified = date(2024, 2, 18)
	// 업종마다 전화번호가 다른경우 store/store.go 파일의 setPhoneNumber 함수에서 하드코딩
	c.PhoneNumber = "010-6590-7589"
	c.SearchEngineConnection = &searchEngineConnection{
		Google: "_0O-P4S7tPNubMmy6jQikADwwAgFvJH5Ep0gWbFthYM",
	}
	Config = c
}
