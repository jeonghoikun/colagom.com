package store

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/jeonghoikun/colagom.com/site"
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

type Category struct {
	Name   string
	Stores []*Store
}

func ListAllCategories() []*Category {
	list := []*Category{}
	for _, s := range ListAllStores() {
		ok := false
		for _, c := range list {
			if s.Type == c.Name {
				ok = true
				break
			}
		}
		if !ok {
			list = append(list, &Category{
				Name:   s.Type,
				Stores: []*Store{s},
			})
			continue
		}
		for _, c := range list {
			if s.Type == c.Name {
				c.Stores = append(c.Stores, s)
			}
		}
	}
	sort.Slice(list, func(i, j int) bool { return list[i].Name < list[j].Name })
	for _, x := range list {
		sort.Slice(x.Stores, func(i, j int) bool {
			return x.Stores[i].DatePublished.UnixNano() > x.Stores[j].DatePublished.UnixNano()
		})
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

type TimeType struct {
	// Has: 유무
	Has bool
	// Open: 오픈시간. ex) 18:00
	Open string
	// Closed: 마감시간. ex) 00:00
	Closed string
}

type Hour struct {
	// Part1: 1부
	Part1 *TimeType
	// Part2: 2부
	Part2 *TimeType
}

type Menu struct {
	// Part1Whisky: 1부 주대
	Part1Whisky int
	// Part2Whisky: 2부 주대
	Part2Whisky int
	// TC: 아가씨 티시
	TC int
	// RT: 룸비
	RT int
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
	// Hour: 영업시간 하드코딩
	Hour *Hour
	// Price: 가격 하드코딩
	Menu *Menu
	// PhoneNumber: 하드코딩 X.
	PhoneNumber string
	// 생성일
	DatePublished time.Time
	// 수정일
	DateModified time.Time
}

func (s *Store) IsModified() bool { return s.DatePublished.UnixNano() != s.DateModified.UnixNano() }

func storeDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

func initKaraoke() {
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "논현동",
			Address:      "151-30",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3164.8479106529085!2d127.03145169999998!3d37.5115051!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3f05b7c4407%3A0xbb44e0b5425b8a89!2z7ISc7Jq47Yq567OE7IucIOqwleuCqOq1rCDrhbztmITrj5kgMTUxLTMw!5e0!3m2!1sko!2skr!4v1660745693771!5m2!1sko!2skr",
		},
		Type:        STORE_TYPE_KARAOKE,
		Title:       "퍼펙트",
		Description: "강남 퍼펙트 가라오케는 화려한 분위기와 최신 노래 라이브러리로 여러분을 미쳐하게 할 최고의 유흥주점입니다. 노래를 부르고 춤추며 즐기는 가라오케 무대에서 또래와 즐거운 시간을 보내보세요!",
		Active: &Active{
			IsPermanentClosed: true,
			Reason:            "하이퍼블릭으로 업종 변경",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 350000, Part2Whisky: 160000, TC: 120000, RT: 50000},
		DatePublished: storeDate(2023, 9, 5),
		DateModified:  storeDate(2023, 10, 15),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "삼성동",
			Address:      "142-35",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.1043050533926!2d127.05085469999999!3d37.505458!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca411d5a288d7%3A0xca6681460caa4840!2s411%20Teheran-ro%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1662046616801!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_KARAOKE,
		Title:       "파티원",
		Description: "강남 파티원 가라오케는 최신 노래 라이브러리와 화려한 무대로 여러분을 미쳐하게 할 최고의 유흥주점입니다. 노래를 부르고 춤추며 즐기는 가라오케 무대에서 또래와 즐거운 시간을 보내보세요!",
		Active: &Active{
			IsPermanentClosed: true,
			Reason:            "정상폐업",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 150000, TC: 100000, RT: 50000},
		DatePublished: storeDate(2024, 2, 18),
		DateModified:  storeDate(2024, 2, 18),
	})
}

func initShirtRoom() {
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "삼성동",
			Address:      "142-35",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.1043050533926!2d127.05085469999999!3d37.505458!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca411d5a288d7%3A0xca6681460caa4840!2s411%20Teheran-ro%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1662046616801!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_SHIRTROOM,
		Title:       "디씨",
		Description: "강남 디씨 셔츠룸은 강남 지역의 유흥주점으로, 멋진 셔츠를 입고 즐기는 쇼와 엔터테인먼트를 제공하여 파티와 즐거운 시간을 보낼 수 있는 장소입니다.",
		Active: &Active{
			IsPermanentClosed: true,
			Reason:            "정상폐업",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 160000, Part2Whisky: 130000, TC: 120000, RT: 50000},
		DatePublished: storeDate(2023, 9, 5),
		DateModified:  storeDate(2024, 2, 15),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "잠원동",
			Address:      "18-9",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3164.7060647283693!2d127.0171104!3d37.514850200000005!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3dd364c8bc7%3A0x3ab4d058c71d79a8!2s18-9%20Jamwon-dong%2C%20Seocho-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1670862647642!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_SHIRTROOM,
		Title:       "유앤미",
		Description: "강남 유앤미 셔츠룸은 화려한 강남 지역에서 즐길 수 있는 역동적인 유흥주점으로, 멋진 셔츠룸 경험과 다양한 엔터테인먼트 옵션을 제공합니다.",
		Active: &Active{
			IsPermanentClosed: true,
			Reason:            "하이퍼블릭으로 업종 변경",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 160000, Part2Whisky: 130000, TC: 120000, RT: 50000},
		DatePublished: storeDate(2023, 9, 5),
		DateModified:  storeDate(2024, 6, 4),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "삼성동",
			Address:      "143-27",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.0354982629583!2d127.0543849!3d37.5070809!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca413c457ed95%3A0x2c8f79900d733d24!2s143-27%20Samseong-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1685329268008!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_SHIRTROOM,
		Title:       "씨엔엔",
		Description: "강남 씨엔엔 셔츠룸에서는 고급스러움과 맞춤 서비스로 귀하의 밤을 특별하게 만들어 드립니다. 독특한 분위기에서 최상의 경험을 제공하며, 강남의 밤을 더욱 빛내 줄 것입니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 160000, Part2Whisky: 130000, TC: 120000, RT: 50000},
		DatePublished: storeDate(2024, 2, 15),
		DateModified:  storeDate(2024, 2, 15),
	})
}

func initHighPublic() {
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "604-7",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.088372324827!2d127.0311099!3d37.5058338!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3fb63865cd7%3A0x31427b556da83644!2s604-7%20Yeoksam-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1662056274810!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_HIGHPUBLIC,
		Title:       "달토",
		Description: "강남의 달토 하이퍼블릭은 화려한 밤문화와 다채로운 엔터테인먼트를 즐길 수 있는 유흥주점 지구로 유명한 곳으로, 현장에서 무한한 즐거움을 찾을 수 있습니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 170000, Part2Whisky: 150000, TC: 120000, RT: 50000},
		DatePublished: storeDate(2023, 9, 5),
		DateModified:  storeDate(2024, 1, 26),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "604-7",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.088372324827!2d127.0311099!3d37.5058338!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3fb63865cd7%3A0x31427b556da83644!2s604-7%20Yeoksam-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1662056274810!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_HIGHPUBLIC,
		Title:       "런닝래빗",
		Description: "강남 런닝래빗 하이퍼블릭은 서울 강남의 인기 유흥주점으로, 화려한 분위기와 다채로운 엔터테인먼트를 즐길 수 있는 곳입니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 170000, Part2Whisky: 150000, TC: 120000, RT: 50000},
		DatePublished: storeDate(2023, 9, 28),
		DateModified:  storeDate(2024, 1, 26),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "831-42",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.6005044881886!2d127.03146729999997!3d37.4937527!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca15057aba5c3%3A0x3c39e1c32ad3bd0f!2s831-42%20Yeoksam-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1665731145337!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_HIGHPUBLIC,
		Title:       "트렌드",
		Description: "강남 트렌드 하이퍼블릭은 강남 지역에서 최신 트렌드를 반영한 현대적인 유흥주점으로, 열정적인 무드와 다채로운 엔터테인먼트로 방문객들에게 특별한 경험을 제공합니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 200000, Part2Whisky: 0, TC: 120000, RT: 50000},
		DatePublished: storeDate(2023, 9, 5),
		DateModified:  storeDate(2024, 1, 26),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "대치동",
			Address:      "890-38",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.150656732908!2d127.05328440000001!3d37.504364699999996!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca41055280155%3A0xc6516a6b77ef70c1!2z7ISc7Jq47Yq567OE7IucIOqwleuCqOq1rCDrjIDsuZjrj5kgODkwLTM4!5e0!3m2!1sko!2skr!4v1660489421580!5m2!1sko!2skr",
		},
		Type:        STORE_TYPE_HIGHPUBLIC,
		Title:       "사라있네",
		Description: "강남 사라있네 하이퍼블릭은 서울 강남에서 유명한 유흥주점으로, 화려한 분위기와 다양한 엔터테인먼트 옵션을 제공하여 즐거운 밤을 보낼 수 있는 곳입니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 170000, Part2Whisky: 150000, TC: 120000, RT: 50000},
		DatePublished: storeDate(2023, 9, 5),
		DateModified:  storeDate(2024, 1, 26),
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
		Title:       "메이커",
		Description: "강남 메이커 하이퍼블릭은 대한민국 서울의 강남 지역에서 유명한 유흥주점 중 하나입니다. 이 곳은 강남의 화려한 밤문화와 열기로운 분위기를 즐기고 싶은 사람들을 위한 완벽한 장소입니다.",
		Active: &Active{
			IsPermanentClosed: true,
			Reason:            "정상폐업",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 110000, RT: 50000},
		DatePublished: storeDate(2023, 9, 5),
		DateModified:  storeDate(2024, 11, 15),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "823-30",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.4051104401256!2d127.03307020000001!3d37.4983624!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca1565c22d639%3A0x1fcb22298cd33520!2z7ISc7Jq47Yq567OE7IucIOqwleuCqOq1rCDsl63sgrzrj5kgODIzLTMw!5e0!3m2!1sko!2skr!4v1693829638202!5m2!1sko!2skr",
		},
		Type:        STORE_TYPE_HIGHPUBLIC,
		Title:       "수목원",
		Description: "강남 수목원 하이퍼블릭은 강남 지역에서 독특하고 화려한 밤문화를 즐길 수 있는 최고의 장소 중 하나로 손꼽히며, 다채로운 액티비티와 흥겨운 분위기를 찾는 이들에게 확실한 만족을 제공합니다. 24시간 영업으로, 서울의 밤을 더욱 특별하게 만들어 줄 곳 중 하나입니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 250000, Part2Whisky: 0, TC: 130000, RT: 50000},
		DatePublished: storeDate(2023, 9, 5),
		DateModified:  storeDate(2024, 1, 26),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "논현동",
			Address:      "151-30",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3164.8479106529085!2d127.03145169999998!3d37.5115051!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3f05b7c4407%3A0xbb44e0b5425b8a89!2z7ISc7Jq47Yq567OE7IucIOqwleuCqOq1rCDrhbztmITrj5kgMTUxLTMw!5e0!3m2!1sko!2skr!4v1660745693771!5m2!1sko!2skr",
		},
		Type:        STORE_TYPE_HIGHPUBLIC,
		Title:       "퍼펙트",
		Description: "강남 퍼펙트 하이퍼블릭에서는 독보적인 서비스와 미학이 공존하는 공간에서 품격 있고 독특한 경험을 제공합니다. 도심 속에서 찾아보기 힘든 특별한 휴식과 여유를 선사하며, 각종 이벤트와 프로모션을 통해 방문객에게 새로운 즐거움과 설렘을 선사합니다. 서울 강남구의 중심지에서 품격과 스타일을 느껴보세요.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 150000, TC: 120000, RT: 50000},
		DatePublished: storeDate(2023, 10, 15),
		DateModified:  storeDate(2024, 1, 26),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "824-8",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.434083647925!2d127.0305156!3d37.4976789!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca15741b03c33%3A0xf28611c1cfc94af5!2z7ISc7Jq47Yq567OE7IucIOqwleuCqOq1rCDsl63sgrzrj5kgODI0LTg!5e0!3m2!1sko!2skr!4v1704324043037!5m2!1sko!2skr",
		},
		Type:        STORE_TYPE_HIGHPUBLIC,
		Title:       "워라벨",
		Description: "강남 워라벨 하이퍼블릭은 강남구에 위치한 유흥주점으로, 일과 삶의 균형을 중시하는 현대적인 분위기와 고급 서비스를 제공합니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 200000, Part2Whisky: 0, TC: 120000, RT: 50000},
		DatePublished: storeDate(2024, 1, 26),
		DateModified:  storeDate(2024, 1, 26),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "832-7",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.679446741483!2d127.02837221193238!3d37.49189017194145!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca1502738de7b%3A0x65a8ee648278baf2!2z7ISc7Jq47Yq567OE7IucIOqwleuCqOq1rCDsl63sgrzrj5kgODMyLTc!5e0!3m2!1sko!2skr!4v1704324092279!5m2!1sko!2skr",
		},
		Type:        STORE_TYPE_HIGHPUBLIC,
		Title:       "방탄",
		Description: "강남 방탄 하이퍼블릭은 서울 강남구의 인기 있는 유흥주점으로, 최신 음향 시스템과 현대적인 디자인을 자랑하며 다채로운 엔터테인먼트를 즐길 수 있습니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 200000, Part2Whisky: 0, TC: 130000, RT: 50000},
		DatePublished: storeDate(2024, 1, 26),
		DateModified:  storeDate(2024, 1, 26),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "잠원동",
			Address:      "18-9",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3164.7060647283693!2d127.0171104!3d37.514850200000005!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3dd364c8bc7%3A0x3ab4d058c71d79a8!2s18-9%20Jamwon-dong%2C%20Seocho-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1670862647642!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_HIGHPUBLIC,
		Title:       "유앤미",
		Description: "강남 유앤미 하이퍼블릭은 최신 음향과 조명을 갖춘 고급 유흥업소입니다. 세련된 인테리어와 다양한 엔터테인먼트로 특별한 밤을 즐길 수 있습니다. 지금 예약하세요!",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 170000, Part2Whisky: 150000, TC: 120000, RT: 50000},
		DatePublished: storeDate(2024, 6, 4),
		DateModified:  storeDate(2024, 6, 4),
	})
}

func initLeggingsRoom() {
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "삼성동",
			Address:      "144-10",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.0368804441946!2d127.0548939!3d37.5070483!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca413ea3ed99f%3A0xdd0a3d80af8a9047!2s144-10%20Samseong-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1662646930422!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_LEGGINGS,
		Title:       "하이킥",
		Description: "강남 하이킥 레깅스룸은 매력적인 레깅스걸들과 함께 즐길 수 있는 곳으로, 실력있는 댄서들의 화려한 퍼포먼스와 함께 음악과 춤, 술과 음식을 즐길 수 있습니다. 여기에서는 다양한 칵테일과 주류를 제공하며, 고객들은 친구들과 즐거운 시간을 보낼 수 있는 최적의 장소입니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "01:00"},
			Part2: &TimeType{Has: true, Open: "01:00", Closed: "15:00"},
		},
		Menu:          &Menu{Part1Whisky: 250000, Part2Whisky: 0, TC: 150000, RT: 50000},
		DatePublished: storeDate(2023, 9, 5),
		DateModified:  storeDate(2024, 1, 26),
	})
}

func initDot5() {
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "논현동",
			Address:      "204-4",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.069981650343!2d127.02487893188555!3d37.50626757076464!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3fb554ff02b%3A0x8d9e573a46ec1b7a!2s204-4%20Nonhyeon-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1679716196560!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "유니크",
		Description: "강남 유니크 쩜오는 강남 지역에서 유명한 유흥주점 중 하나로, 독특하고 특별한 분위기를 제공하는 곳입니다. 이 곳은 현대적이고 화려한 인테리어로 꾸며져 있으며, 고객들에게 다채로운 음악과 엔터테인먼트를 제공합니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2023, 9, 5),
		DateModified:  storeDate(2024, 1, 26),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "831",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.5641798788934!2d127.0297203!3d37.4946097!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca1508715f00d%3A0xf4d079a0f225c1b1!2s831%20Yeoksam-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1679397724056!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "831",
		Description: "강남 831 쩜오는 편안한 의자와 테이블이 배치된 모던하고 스타일리시한 인테리어로 꾸며져 있으며, DJ의 멋진 음악과 무대 퍼포먼스로 고객들에게 즐거운 시간을 선사합니다. 또한, 다양한 주류 메뉴와 칵테일을 제공하여, 고객들이 자신의 취향에 맞게 골라 즐길 수 있습니다.",
		Active: &Active{
			IsPermanentClosed: true,
			Reason:            "썸데이로 상호 변경",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2023, 9, 5),
		DateModified:  storeDate(2024, 9, 13),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "735-32",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.3760308056935!2d127.0341289!3d37.4990484!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca1560e5d6327%3A0x5c114aeb8260a643!2s735-32%20Yeoksam-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1679397562888!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "에이원",
		Description: "강남 에이원 쩜오는 강남 지역에서 손님들에게 다양한 엔터테인먼트와 음료를 제공하는 유흥주점입니다. 이 장소는 강남의 활기찬 나이트라이프와 열띤 분위기를 경험하고자 하는 분들에게 인기가 있습니다. 강남 에이원 쩜오는 친절하고 프로페셔널한 스태프들이 고객 서비스에 최선을 다하며, 파티나 이벤트를 개최하는 데에도 적합한 장소입니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2023, 9, 5),
		DateModified:  storeDate(2024, 1, 26),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "삼성동",
			Address:      "141-33",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.121539211326!2d127.04949690000001!3d37.5050515!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca40fc775ade5%3A0xdd9b10797e776ad1!2s141-33%20Samseong-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1678667592079!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "미라클",
		Description: "강남 미라클 쩜오는 대한민국 서울의 강남 지역에 위치한 유흥주점으로, 지역 내에서 손꼽히는 인기 있는 엔터테인먼트 장소 중 하나입니다. 이 곳은 한국의 대표적인 유흥 문화를 체험할 수 있는 곳 중 하나로, 다양한 엔터테인먼트 옵션과 즐거운 분위기로 손님들을 맞이합니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2023, 9, 5),
		DateModified:  storeDate(2024, 1, 26),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "701-2",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.204884148887!2d127.0430503!3d37.5030856!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca406fc7ff209%3A0x341d4adf49840962!2s701-2%20Yeoksam-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1678667437305!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "오키도키",
		Description: "강남 오키도키는 대한민국 서울의 강남 지역에 위치한 인기 유흥주점으로, 현지와 관광객 모두에게 인기 있는 엔터테인먼트 공간 중 하나입니다. 이 곳은 다양한 오락 요소와 파티 분위기를 즐길 수 있는 곳으로, 다음과 같은 특징을 가지고 있습니다.",
		Active: &Active{
			IsPermanentClosed: true,
			Reason:            "더글로리로 상호 변경",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2023, 9, 5),
		DateModified:  storeDate(2024, 4, 27),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "신사동",
			Address:      "561-30",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3164.5256239750274!2d127.0258308!3d37.5191051!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3ecf7b91b35%3A0x90e6eb4e73a5644e!2s561-30%20Sinsa-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1678606137375!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "인트로",
		Description: "강남 인트로 쩜오는 강남의 다른 유흥 시설과 가까워 위치적으로도 편리하며, 서울에서의 특별한 밤을 만들고자 하는 분들에게 추천하는 곳 중 하나입니다. 편안한 분위기와 환상적인 엔터테인먼트를 원한다면, 강남 인트로 쩜오를 방문해보세요.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2023, 9, 5),
		DateModified:  storeDate(2024, 1, 26),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "논현동",
			Address:      "248-7",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3164.741047626004!2d127.03369181564705!3d37.51402523489071!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3f415b07255%3A0x2162a0d614d3c110!2s640%20Eonju-ro%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1678605759071!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "머니볼",
		Description: "강남 머니볼 쩜오는 대한민국 서울 강남 지역에 위치한 유흥주점 중 하나로, 도심의 번화가에서 활기찬 분위기와 다양한 엔터테인먼트를 제공하는 곳입니다. 이곳은 주로 젊은 이들과 비즈니스 모임을 위한 장소로 알려져 있으며, 다양한 음료와 칵테일, 안주 메뉴를 즐길 수 있습니다.",
		Active: &Active{
			IsPermanentClosed: true,
			Reason:            "멀리건, 알파벳으로 상호 변경",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2023, 9, 5),
		DateModified:  storeDate(2024, 4, 27),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "731-11",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.401460674176!2d127.0436794!3d37.498448499999995!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca401a6b8183b%3A0xcbcd58a8b2cb7c50!2s731%20Yeoksam-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1678605118720!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "라이징",
		Description: "강남 라이징 쩜오는 강남 지역에 위치한 인기 있는 유흥주점으로, 그 신명을 거두고 있는 곳입니다. 이 곳은 화려하고 현대적인 분위기를 자랑하며, 주로 클럽 음악과 댄스 퍼포먼스로 알려져 있습니다. 라이징 쩜오는 고품질의 음료와 다채로운 칵테일 메뉴로 손님들을 맞이하며, 낮과 밤을 가리지 않고 엔터테인먼트와 즐거움을 제공합니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2023, 9, 5),
		DateModified:  storeDate(2024, 1, 26),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "736-17",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.3715755621406!2d127.03453809999999!3d37.4991535!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca15607cff005%3A0x9a314c8436603f9e!2s736-17%20Yeoksam-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1677802895674!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "임팩트",
		Description: "강남 임팩트 쩜오는 서울의 강남 지역에 위치한 역동적인 유흥주점으로, 이곳은 강남의 밤문화를 대표하는 곳 중 하나입니다. 임팩트 쩜오는 화려한 조명과 풍부한 음악으로 분위기를 고조시키며, 다채로운 칵테일과 음료를 제공하여 고객들에게 즐거운 시간을 제공합니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2023, 9, 5),
		DateModified:  storeDate(2024, 1, 26),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "824-7",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.429128349951!2d127.03037690000001!3d37.4977958!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca1576a139921%3A0xda0428a0d46a18b2!2s824-7%20Yeoksam-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1676634190100!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "스테이",
		Description: "강남 스테이 쩜오는 아름다운 인테리어와 조명, 고급스러운 분위기로 고객들에게 편안한 분위기를 제공합니다. 다양한 음료 메뉴와 간식을 통해 손님들에게 다채로운 맛을 제공하며, 프로페셔널한 바텐더들이 특별한 칵테일을 선사해줍니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2023, 9, 5),
		DateModified:  storeDate(2024, 1, 26),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "702-16",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.158957771797!2d127.0454229!3d37.504168899999996!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca405e2735e15%3A0xc330c6245a409809!2s702-16%20Yeoksam-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1661933858945!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "에프원",
		Description: "강남 에프원 쩜오는 강남 지역의 다양한 연령대와 취향을 고려하여 다양한 프로그램과 이벤트를 제공하며, 친구들과의 모임, 회식, 파티, 혹은 특별한 날을 기념하기에 최적의 장소 중 하나입니다. 강남 에프원 쩜오에서 잊지 못할 경험을 즐겨보세요!",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2023, 9, 5),
		DateModified:  storeDate(2023, 9, 5),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "677-22",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.2307457193765!2d127.03704181193267!3d37.50247557193869!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3f8acb4cd37%3A0xa46ef02bf086e82c!2z7ISc7Jq47Yq567OE7IucIOqwleuCqOq1rCDsl63sgrzrj5kgNjc3LTIy!5e0!3m2!1sko!2skr!4v1704324149895!5m2!1sko!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "킹스맨",
		Description: "강남 킹스맨 쩜오는 고급스러운 분위기의 강남구 유흥주점으로, 세련된 인테리어와 프리미엄 서비스로 강남의 밤문화를 대표하는 명소입니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2024, 1, 26),
		DateModified:  storeDate(2024, 1, 26),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "701-2",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.204884148887!2d127.0430503!3d37.5030856!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca406fc7ff209%3A0x341d4adf49840962!2s701-2%20Yeoksam-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1678667437305!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "더글로리",
		Description: "강남 더글로리 쩜오는 최상의 서비스와 고급스러운 분위기로 강남의 밤을 화려하게 장식합니다. 매혹적인 인테리어와 섬세한 요리가 어우러져 모든 순간을 더욱 특별하게 만들어 드립니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2024, 3, 23),
		DateModified:  storeDate(2024, 4, 27),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "논현동",
			Address:      "248-7",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3164.741047626004!2d127.03369181564705!3d37.51402523489071!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3f415b07255%3A0x2162a0d614d3c110!2s640%20Eonju-ro%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1678605759071!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "멀리건",
		Description: "강남 멀리건 쩜오는 고급스러운 분위기와 전문 바텐더가 제공하는 프리미엄 칵테일, 활기찬 라이브 음악으로 강남의 밤을 완벽하게 만듭니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2024, 4, 27),
		DateModified:  storeDate(2024, 4, 27),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "논현동",
			Address:      "248-7",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3164.741047626004!2d127.03369181564705!3d37.51402523489071!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3f415b07255%3A0x2162a0d614d3c110!2s640%20Eonju-ro%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1678605759071!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "알파벳",
		Description: "강남 알파벳 쩜오에서는 세련된 인테리어, 맞춤형 칵테일, 그리고 독특한 이벤트로 강남의 밤을 특별하게 만들어 드립니다. 매력적인 밤을 경험하세요!",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2024, 4, 27),
		DateModified:  storeDate(2024, 4, 27),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "831",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.5641798788934!2d127.0297203!3d37.4946097!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca1508715f00d%3A0xf4d079a0f225c1b1!2s831%20Yeoksam-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1679397724056!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "썸데이",
		Description: "강남 썸데이 쩜오는 세련된 인테리어와 고급스러운 분위기 속에서 다양한 프리미엄 음료를 즐길 수 있는 유흥주점입니다. 전문 바텐더가 제공하는 독창적인 칵테일과 주류는 고객의 취향을 만족시키며, 프라이빗한 공간에서 편안하게 시간을 보낼 수 있습니다. 차별화된 서비스와 세심한 배려로 강남의 밤을 더욱 특별하게 만들어주는 최적의 장소입니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2024, 9, 13),
		DateModified:  storeDate(2024, 9, 13),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "역삼동",
			Address:      "822-5",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.3849794120856!2d127.02926860000001!3d37.4988373!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca159d7d08f47%3A0x19ac7457d361928!2z7ISc7Jq47Yq567OE7IucIOqwleuCqOq1rCDthYztl6TrnoDroZwgMTEx!5e0!3m2!1sko!2skr!4v1661153125692!5m2!1sko!2skr",
		},
		Type:        STORE_TYPE_DOT5,
		Title:       "블렌딩",
		Description: "",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "05:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2024, 9, 20),
		DateModified:  storeDate(2024, 9, 20),
	})
}

func initClub() {
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "도산대로",
			Address:      "114",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3164.64159846425!2d127.02127!3d37.5163704!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3e9a9f07727%3A0x4fcde2f83452e564!2s114%20Dosan-daero%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1681189780772!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_CLUB,
		Title:       "사운드",
		Description: "강남 사운드 클럽은 세계적으로 유명한 디제이와 아티스트들을 초빙하여 매주 열리는 화려한 파티와 이벤트로 손님들을 매료시킵니다. 클럽 내부는 현대적이고 고급스러운 인테리어로 꾸며져 있으며, 강렬한 빛과 사운드 시스템이 함께 어우러져 춤을 춤으로써 파티 분위기를 한층 더 뜨겁게 만듭니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "23:00", Closed: "11:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2023, 9, 5),
		DateModified:  storeDate(2024, 1, 26),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "잠원동",
			Address:      "21-3",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3164.6962477822185!2d127.0192326!3d37.51508169999999!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca3e80fe94731%3A0xadedf946e74c560c!2s21-3%20Jamwon-dong%2C%20Seocho-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1681189358457!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_CLUB,
		Title:       "레이스",
		Description: "강남 레이스 클럽은 대한민국 서울의 중심인 강남 지역에 위치한 역동적이고 흥미로운 엔터테인먼트 공간입니다. 이 클럽은 클래식한 나이트클럽 분위기와 현대적인 디자인을 조화롭게 결합하여, 파티와 레이싱의 열기를 동시에 느낄 수 있는 곳으로 알려져 있습니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "23:00", Closed: "11:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 0, Part2Whisky: 0, TC: 0, RT: 0},
		DatePublished: storeDate(2023, 9, 5),
		DateModified:  storeDate(2024, 1, 26),
	})
}

func initHobba() {
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "삼성동",
			Address:      "143-35",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.0622628938677!2d127.05028567647602!3d37.5064496275705!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca4118576f5e1%3A0xbc745a3337004851!2s143-35%20Samseong-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1685329649613!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_HOBBA,
		Title:       "어게인",
		Description: "강남 어게인 호빠는 서울의 강남 지역에서 인기 있는 여성유흥주점 중 하나로, 화려한 분위기와 다양한 엔터테인먼트 옵션을 제공하는 곳입니다. 이곳은 파티와 이벤트, 친구들과의 모임, 비즈니스 만찬 및 여가 시간을 즐기기에 완벽한 장소로 손꼽힙니다.",
		Active: &Active{
			IsPermanentClosed: false,
			Reason:            "",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "15:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 180000, Part2Whisky: 0, TC: 60000, RT: 50000},
		DatePublished: storeDate(2023, 9, 5),
		DateModified:  storeDate(2024, 1, 26),
	})
	stores = append(stores, &Store{
		Location: &Location{
			Do:           "서울",
			Si:           "강남구",
			Dong:         "삼성동",
			Address:      "143-27",
			GoogleMapSrc: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3165.0354982629583!2d127.0543849!3d37.5070809!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x357ca413c457ed95%3A0x2c8f79900d733d24!2s143-27%20Samseong-dong%2C%20Gangnam-gu%2C%20Seoul!5e0!3m2!1sen!2skr!4v1685329268008!5m2!1sen!2skr",
		},
		Type:        STORE_TYPE_HOBBA,
		Title:       "씨엔엔",
		Description: "강남은 대한민국의 중심지 중 하나로, 그만큼 강남 씨엔엔 호빠의 엔터테인먼트 공간도 고품격이며 세련된 분위기를 자랑합니다. 강남 씨엔엔 호빠 역시 고급스러운 인테리어와 조명으로 공간을 아름답게 꾸며, 고객들에게 편안하고 럭셔리한 분위기를 제공합니다.",
		Active: &Active{
			IsPermanentClosed: true,
			Reason:            "셔츠룸으로 업종 변경",
		},
		Hour: &Hour{
			Part1: &TimeType{Has: true, Open: "18:00", Closed: "15:00"},
			Part2: &TimeType{Has: false},
		},
		Menu:          &Menu{Part1Whisky: 180000, Part2Whisky: 0, TC: 60000, RT: 50000},
		DatePublished: storeDate(2023, 9, 5),
		DateModified:  storeDate(2024, 2, 15),
	})
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
		case STORE_TYPE_DOT5:
			s.PhoneNumber = "010-2170-4981"
		case STORE_TYPE_CLUB, STORE_TYPE_HOBBA:
			s.PhoneNumber = "010-6590-7589"
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
	initKaraoke()
	initShirtRoom()
	initHighPublic()
	initLeggingsRoom()
	initDot5()
	initClub()
	initHobba()

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
