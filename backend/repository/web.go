package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/voyagegroup/treasure-2020-b/model"
)

func CreateWeb(db *sqlx.Tx, w *model.Web) (result sql.Result, err error) {
	stmt, err := db.Prepare(`
INSERT INTO website (owner_id, title, profile, theme, content) 
VALUES (?, ?, ?, ?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()
	return stmt.Exec(
		w.OwnerID,
		w.Title,
		w.Profile,
		w.Theme,
		w.Content,
	)
}

func AllWeb(db *sqlx.DB, ownerID int64) ([]model.Web, error) {
	webs := make([]model.Web, 0)
	if err := db.Select(&webs, `SELECT id, owner_id, title, profile, theme, content 
FROM website 
WHERE owner_id = ?
`, ownerID); err != nil {
		return nil, err
	}
	return webs, nil
}

func GetWebByID(db *sqlx.DB, id int64) (*model.Web, error) {
	web := model.Web{}
	if err := db.Get(&web, `SELECT id, owner_id, title, profile, theme, content 
FROM website 
WHERE id = ?
`, id); err != nil {
		return nil, err
	}
	return &web, nil
}
