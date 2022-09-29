//go:generate easyjson -all .

package domain

const (
	UNREAD_STATUS   = "0"
	ARCHIVED_STATUS = "1"
)

type Topic struct {
	ItemID     string `json:"item_id"`
	GivenURL   string `json:"given_url"`
	Status     string `json:"status"`
	TimeToRead int64  `json:"time_to_read"`
}

type Response struct {
	List map[string]Topic `json:"list"`
}
