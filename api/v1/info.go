package v1

import (
	database "ginsdgo/databse"
	"ginsdgo/model"
	"ginsdgo/sdgo"

	"github.com/gin-gonic/gin"
)


func Info(c *gin.Context) {
	// get phone
	value, _ := c.Get("phone")
	user,_ := value.(string)
	_,u := sdgo.GetUserInfo(user)
	// send the json of u to the client
	c.JSON(200, gin.H{
		"status":  200,
		"data":    u,
		"message": "",
	})
}


func Matches(c *gin.Context) {
	// get matches list
	value, _ := c.Get("name")
	username,_ := value.(string)
	
	//sql.DB
	db := database.GetDB()
	ret := model.MatchList{}
	query := `SELECT * FROM ghistory WHERE USER = ? ORDER BY gameid DESC`
	rows, err := db.Query(query, username)
	success := true
	if err==nil{
		for rows.Next() {
			var m model.Match
			if err := rows.Scan(&m.Gameid, &m.User, &m.Botid, &m.Atk, &m.Def, &m.Skill, &m.Kills, &m.Death, &m.Support, &m.Win, &m.Types, &m.P1, &m.P2, &m.P3, &m.P4, &m.Performance); err != nil {
				success = false
				break
			}
			ret.List = append(ret.List, m)
		}
		rows.Close()
	}else{
		success = false
	}
	ret.Success = success

	c.JSON(200, gin.H{
		"status":  200,
		"data":    ret,
		"message": "",
	})
}