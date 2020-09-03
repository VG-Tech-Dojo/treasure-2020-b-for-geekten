package model

type ZoomToken struct {
	ID           int64  `db:"id" json:"id"`
	OwnerID      int64  `db:"owner_id" json:"owner_id"`
	AcccessToken string `db:"access_token" json:"access_token"`
	RefreshToken string `db:"refresh_token" json:"refresh_token"`
}
