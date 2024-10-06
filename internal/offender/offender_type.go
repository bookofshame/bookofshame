package offender

import (
	"strings"

	"github.com/google/uuid"
)

type Offender struct {
	Id             int     `db:"id" json:"id"`
	FullName       string  `db:"fullName" json:"fullName"`
	Address        string  `db:"address" json:"address"`
	DivisionId     int     `db:"divisionId" json:"divisionId"`
	DistrictId     int     `db:"districtId" json:"districtId"`
	UpazilaId      *int    `db:"upazilaId" json:"upazilaId"`
	UnionId        *int    `db:"unionId" json:"unionId"`
	IsOrganization bool    `db:"isOrganization" json:"isOrganization"`
	IsEnabler      bool    `db:"isEnabler" json:"isEnabler"`
	IsPerpetrator  bool    `db:"isPerpetrator" json:"isPerpetrator"`
	Photo          *string `db:"photo" json:"photo"`
	Metadata       string  `db:"metadata" json:"omitempty"`
	CreatedAt      string  `db:"createdAt" json:"createdAt"`
	CreatedBy      *string `db:"createdBy,omitempty" json:"createdBy"`
}

func (c *Offender) NormalizedName() string {
	return strings.ReplaceAll(c.FullName, " ", "_")
}

func (c *Offender) GeneratePhotoName() string {
	return c.NormalizedName() + "_" + uuid.New().String() + ".jpg"
}
