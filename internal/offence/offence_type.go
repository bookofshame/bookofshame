package offence

type Offence struct {
	Id          int    `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Address     string `db:"address" json:"address"`
	Thana       string `db:"thana" json:"thana"`
	District    string `db:"district" json:"district"`
	Division    string `db:"division" json:"division"`
	Metadata    string `db:"metadata" json:"omitempty"`
	DateTime    string `db:"dateTime" json:"dateTime"`
	CreatedAt   string `db:"createdAt" json:"createdAt"`
}
