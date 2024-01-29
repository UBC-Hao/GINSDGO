package sdgo

import (
	"encoding/json"
	"io/ioutil"
	"sync"
	"time"
)

type SDGOUser struct{
	Lastlogin time.Time
	QQ string `json:"QQ"`
	CustomName string `json:"customname"`
	Icon int
	Name  string `json:"name"`
	UID   uint32 `json:"uid"`
	D1 uint8
	D2 uint8
	Rank  uint8
	RankL uint8
	Exp   uint32
	Hej 	uint32
	//Special int
	BP    uint32
	GP    uint32
	Slot  uint32
	Grid     *Grid       `json:"grid"`
	PingBi string
}


type Grid struct{
	Robot      []*Robot     `json:"list"`
	GO         int
	PageCount  int
	NextUUID   uint64       `json:"nextUUID"`
}

type Robot struct {
	ID   HexBotID
	Pos  uint16
	UUID HexUint64
	C8     []HexUint16 `json:",omitempty"`
	Wing   uint8
	WingLv HexByte //[]byte // 4 byte
	Sess   uint32
	Lv     uint8
	Skill  HexUint32
}


func CheckLogin(phone string, username string) bool {
	ok, user := GetUserInfo(phone)
	if !ok {
		return false
	}
	if user.Name != username {
		return false
	}
	return true
}

var (
	// map to store ACC to *SDGOUser
	//initialize with 0
	ACCMap map[string]*SDGOUser = make(map[string]*SDGOUser,1000)
	// mutex
	accMapMutex sync.RWMutex
)


func GetUserInfo(phone string) (bool, *SDGOUser){
	accMapMutex.RLock()
	if user, ok := ACCMap[phone]; ok {
		accMapMutex.RUnlock()
		return true, user
	}
	accMapMutex.RUnlock()

	//read file
	data, err := ioutil.ReadFile("acc/" +phone +".json")
	if err != nil{
		return false, nil
	}
	// Unmarshal json
	var user SDGOUser
	err = json.Unmarshal(data, &user)
	if err != nil {
		//fmt.Println(err)
		return false, nil
	}
	accMapMutex.Lock()
	ACCMap[phone] = &user
	accMapMutex.Unlock()
	return true, &user
}

func init(){
	// flush ACCMap every 10 minutes
	go func(){
		for{
			accMapMutex.Lock()
			ACCMap = make(map[string]*SDGOUser,1000)
			accMapMutex.Unlock()
			time.Sleep(10 * time.Minute)
		}
	}()
}