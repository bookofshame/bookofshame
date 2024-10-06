package location

import "strconv"

type Division struct {
	Id     int    `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	BnName string `db:"bnName" json:"bnName"`
}

func (d Division) Key() string {
	return strconv.Itoa(d.Id)
}

func (d Division) Value() string {
	return d.Name + " (" + d.BnName + ")"
}

type District struct {
	Id         int     `db:"id" json:"id"`
	Name       string  `db:"name" json:"name"`
	BnName     string  `db:"bnName" json:"bnName"`
	DivisionId int     `db:"divisionId" json:"divisionId"`
	Lat        float64 `db:"lat" json:"lat"`
	Long       float64 `db:"long" json:"long"`
}

func (d District) Key() string {
	return strconv.Itoa(d.Id)
}

func (d District) Value() string {
	return d.Name + " (" + d.BnName + ")"
}

type Upazila struct {
	Id         int    `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	BnName     string `db:"bnName" json:"bnName"`
	DistrictId int    `db:"districtId" json:"districtId"`
}

func (u Upazila) Key() string {
	return strconv.Itoa(u.Id)
}

func (u Upazila) Value() string {
	return u.Name + " (" + u.BnName + ")"
}

type Union struct {
	Id        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	BnName    string `db:"bnName" json:"bnName"`
	UpazilaId int    `db:"upazilaId" json:"upazilaId"`
}

func (u Union) Key() string {
	return strconv.Itoa(u.Id)
}

func (u Union) Value() string {
	return u.Name + " (" + u.BnName + ")"
}
