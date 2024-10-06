package gender

import "strconv"

type Gender struct {
	Id     int    `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	BnName string `db:"bnName" json:"bnName"`
}

func (g Gender) Key() string {
	return strconv.Itoa(g.Id)
}

func (g Gender) Value() string {
	return g.Name + " (" + g.BnName + ")"
}
