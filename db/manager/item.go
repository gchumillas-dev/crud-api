package manager

import (
	"database/sql"
	"errors"
	"fmt"

	utils "github.com/gchumillas/crud-api/utils"
)

// Item manages items.
type Item struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// NewItem returns a new item.
func NewItem(ID ...int64) *Item {
	var id int64
	if len(ID) > 0 {
		id = ID[0]
	}

	return &Item{ID: id}
}

// CreateItem creates a new item.
func (item *Item) CreateItem(db *sql.DB) error {
	res, err := db.Exec(`
	insert into item(title, description)
	values(?, ?)`, item.Title, item.Description)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	item.ID = id

	return nil
}

// ReadItem reads an item.
func (item *Item) ReadItem(db *sql.DB) error {
	stmt, err := db.Prepare(`
	select id, title, description
	from item where id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	switch err := stmt.QueryRow(item.ID).Scan(&item.ID, &item.Title, &item.Description); {
	case err == sql.ErrNoRows:
		return err
	case err != nil:
		panic(err)
	}

	return nil
}

// UpdateItem updates an item.
func (item *Item) UpdateItem(db *sql.DB) error {
	query := "update item set title = ?, description = ? where id = ?"
	if _, err := db.Exec(query, item.Title, item.Description, item.ID); err != nil {
		return err
	}

	return nil
}

// DeleteItem deletes an item.
func (item *Item) DeleteItem(db *sql.DB) error {
	query := "delete from item where id = ?"
	if _, err := db.Exec(query, item.ID); err != nil {
		return err
	}

	return nil
}

// GetItems gets all items.
// TODO: rename count by rowsPerPage
func GetItems(db *sql.DB, offset int, count int, sortCol string, sortDir string) ([]Item, error) {
	items := []Item{}

	allowedCols := []string{"id", "title"}
	allowedDirs := []string{"asc", "desc", ""}
	if !utils.InArray(sortCol, allowedCols) || !utils.InArray(sortDir, allowedDirs) {
		return nil, errors.New("Invalid sorting")
	}

	query := fmt.Sprintf(`
	select id, title, description
	from item
	order by %s %s
	limit ?, ?`, sortCol, sortDir)
	rows, err := db.Query(query, offset, count)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		item := Item{}
		if err := rows.Scan(&item.ID, &item.Title, &item.Description); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

// GetNumItems returns the number of items.
func GetNumItems(db *sql.DB) (int, error) {
	query := "select count(*) from item"
	row := db.QueryRow(query)

	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

// SearchItem searches a give item and returns the page.
func SearchItem(db *sql.DB, ID int64, rowsPerPage int, sortCol string, sortDir string) (int, error) {
	query := fmt.Sprintf(`
	select id
	from item
	order by %s %s`, sortCol, sortDir)
	rows, err := db.Query(query)
	if err != nil {
		return 0, err
	}

	var pos int
	for pos = 0; rows.Next(); pos++ {
		var itemID int64
		if err := rows.Scan(&itemID); err != nil {
			return 0, err
		}

		if itemID == ID {
			break
		}
	}

	return pos / rowsPerPage, nil
}
