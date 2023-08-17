package store

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/jeonghoikun/hamjayoung.com/site"
)

const (
	STORE_TYPE_HIGHPUBLIC string = "하이퍼블릭"
	STORE_TYPE_SHIRTROOM  string = "셔츠룸"
	STORE_TYPE_KARAOKE    string = "가라오케"
	STORE_TYPE_LEGGINGS   string = "레깅스룸"
	STORE_TYPE_DOT5       string = "쩜오"
	STORE_TYPE_HOBBA      string = "호빠"
	STORE_TYPE_CLUB       string = "클럽"
)

var stores []*Store = []*Store{}

func Get(do, si, dong, storeType, title string) (o *Store, has bool) {
	for _, s := range stores {
		if s.Location.Do == do && s.Location.Si == si && s.Location.Dong == dong &&
			s.Type == storeType && s.Title == title {
			return s, true
		}
	}
	return nil, false
}

func ListAllStores() []*Store { return stores }

func ListStoresByDoSiAndStoreType(do, si, storeType string) []*Store {
	list := []*Store{}
	for _, s := range stores {
		if s.Location.Do == do && s.Location.Si == si && s.Type == storeType {
			list = append(list, s)
		}
	}
	return list
}

type Location struct {
	// Do: ex) 서울
	Do string
	// Si: ex) 강남구
	Si string
	// Dong: ex) 역삼동
	Dong string
	// Address: ex) 822-5
	Address string
	// GoogleMapSrc: iframe google map의 src속성 값
	GoogleMapSrc string
}

type Keywords []string

func (k *Keywords) String() string { return strings.Join(*k, ",") }

type Active struct {
	// IsPermanentClosed: 영업중=true 폐업=false
	IsPermanentClosed bool
	// Reason: 폐업상태일 경우에만 입력
	Reason string
}

type Store struct {
	Location *Location
	// Type: 업종 하드코딩
	Type string
	// Title: 가게이름 하드코딩
	Title string
	// Description: 가게 설명 하드코딩
	Description string
	// Keywords: 하드코딩 X. 서버 시작시 지역명, 가게이름, 업종 등으로 자동 초기화 됨
	Keywords Keywords
	// Active: 영업, 폐업 유무와 폐업사유 하드코딩
	Active *Active
	// PhoneNumber: 하드코딩 X.
	PhoneNumber string
	// 생성일
	DatePublished time.Time
	// 수정일
	DateModified time.Time
}

func storeDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

func initHighPublic() {
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "822-5",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.3849794120856!2d127.02926860000001!3d37.4988373!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca159d7d08f47%3A0x19ac7457d361928!2z7ISc7Jq47Yq567OE7IucIOqwleuCqOq1rCDthYztl6TrnoDroZwgMTEx!5e0!3m2!1sko!2skr!4v1661153125692!5m2!1sko!2skr",
		},
		Type:        STORE_TYPE_HIGHPUBLIC,
		Title:       "메이커",
		Description: "메이커 is ...",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		DatePublished: storeDate(2023, 8, 5),
		DateModified:  storeDate(2023, 8, 5),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "822-5",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.3849794120856!2d127.02926860000001!3d37.4988373!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca159d7d08f47%3A0x19ac7457d361928!2z7ISc7Jq47Yq567OE7IucIOqwleuCqOq1rCDthYztl6TrnoDroZwgMTEx!5e0!3m2!1sko!2skr!4v1661153125692!5m2!1sko!2skr",
		},
		Type:        STORE_TYPE_HIGHPUBLIC,
		Title:       "메이커2",
		Description: "메이커2 is ...",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		DatePublished: storeDate(2023, 8, 5),
		DateModified:  storeDate(2023, 8, 5),
	})
	// ..
}

func setStoreKeywords() {
	for _, s := range stores {
		s.Keywords = Keywords([]string{
			fmt.Sprintf("%s %s %s %s %s", s.Location.Do, s.Location.Si, s.Location.Dong, s.Type, s.Title),
			fmt.Sprintf("%s %s", s.Location.Do, s.Type),
			fmt.Sprintf("%s %s %s", s.Location.Do, s.Location.Si, s.Type),
			fmt.Sprintf("%s %s %s %s", s.Location.Do, s.Location.Si, s.Location.Dong, s.Type),
			fmt.Sprintf("%s %s", s.Title, s.Type),
			fmt.Sprintf("%s %s 가격", s.Title, s.Type),
			fmt.Sprintf("%s %s 시스템", s.Title, s.Type),
			fmt.Sprintf("%s %s 주소", s.Title, s.Type),
		})
	}
}

func setPhoneNumbers() {
	for _, s := range stores {
		switch s.Type {
		//		case STORE_TYPE_DOT5:
		//			s.PhoneNumber = "010-2222-2222"
		//		case STORE_TYPE_CLUB:
		//			s.PhoneNumber = "010-3333-3333"
		default:
			s.PhoneNumber = site.Config.PhoneNumber
		}
	}
}

func sortStores() {
	sort.Slice(stores, func(i, j int) bool {
		return stores[i].DatePublished.UnixNano() < stores[j].DatePublished.UnixNano()
	})
}

// 서버 시작시 vieiws/store directories 자동 생성
func createViewsDirectories() error {
	for _, s := range stores {
		dir := fmt.Sprintf("views/store/%s/%s/%s/%s",
			s.Location.Do, s.Location.Si, s.Location.Dong, s.Type)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

// 서버 시작시 views/store/../../{{store.Title}}.html 파일 자동 생성
func createHTMLFiles() error {
	for _, s := range stores {
		filepath := fmt.Sprintf("views/store/%s/%s/%s/%s/%s.html",
			s.Location.Do, s.Location.Si, s.Location.Dong, s.Type, s.Title)
		if _, err := os.Stat(filepath); err == nil {
			fmt.Println(err)
			continue
		}
		if err := os.WriteFile(filepath, []byte("write me!"), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

// 서버 시작시 store 이미지 디렉토리 자동 생성
func createStaticImgDirectories() error {
	for _, s := range stores {
		dir := fmt.Sprintf("static/img/store/%s/%s/%s/%s/%s",
			s.Location.Do, s.Location.Si, s.Location.Dong, s.Type, s.Title)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func Init() error {
	initHighPublic()
	// initShirtRoom()
	// initHobba()
	// ..

	sortStores()

	setStoreKeywords()
	setPhoneNumbers()

	if err := createViewsDirectories(); err != nil {
		return err
	}
	if err := createHTMLFiles(); err != nil {
		return err
	}
	if err := createStaticImgDirectories(); err != nil {
		return err
	}
	return nil
}
