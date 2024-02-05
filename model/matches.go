package model


type Match struct{
	Gameid int
	User string
	Botid int
	Atk int
	Def int
	Skill int
	Kills int
	Death int
	Support int
	Win int
	Types int
	P1 int
	P2 int
	P3 int
	P4 int
	Performance int
}

type MatchList struct{
	Success bool
	List []Match
}