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
	c.Port = uint32(8018)
	c.Domain = "hamjayoung.com"
	c.Author = "함자영"
	c.Title = "함자영의 강남룸빵 대탐험"
	c.Description = "함자영실장이 소개하는 강남지역 모든 룸빵의 가격, 시스템, 위치정보 안내. 가라오케, 셔츠룸, 하이퍼블릭, 레깅스룸, 쩜오, 호빠, 클럽의 모든 정보"
	k := Keywords([]string{"함자영의 강남룸빵 대탐험", "함자영실장", "강남룸빵", "룸빵", "가라오케", "셔츠룸", "하이퍼블릭", "레깅스룸", "쩜오", "호빠", "클럽"})
	c.Keywords = &k
	c.DatePublished = date(2023, 8, 10)
	c.DateModified = date(2023, 8, 10)
	// 업종마다 전화번호가 다른경우 store/store.go 파일의 setPhoneNumber 함수에서 하드코딩
	c.PhoneNumber = "010-2781-9627"
	c.SearchEngineConnection = &searchEngineConnection{
		Google: "GOOGLE_SITE_VERIFICATION",
	}
	Config = c
}
