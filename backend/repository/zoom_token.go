package repository

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-2020-b/model"
)

func CreateZoomToken(db *sqlx.Tx, z *model.ZoomToken) (sql.Result, error) {
	log.Println("CreateZoomTokenNNNNNNNNNNN", z)
	stmt, err := db.Prepare(`
INSERT INTO zoom (owner_id, access_token, refresh_token)
VALUE (?, ?, ?)
ON DUPLICATE KEY
UPDATE access_token = ?, refresh_token = ?
`)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
			log.Println(err)
		}
	}()

	return stmt.Exec(z.OwnerID, z.AcccessToken, z.RefreshToken, z.AcccessToken, z.RefreshToken)
}

func GetZoomTokenByOwnerID(db *sqlx.DB, ownerID int64) (*model.ZoomToken, error) {
	z := model.ZoomToken{}
	if err := db.Get(&z, `
SELECT id, owner_id, access_token, refresh_token
FROM zoom WHERE owner_id = ?
`, ownerID); err != nil {
		return nil, err
	}
	return &z, nil
}
