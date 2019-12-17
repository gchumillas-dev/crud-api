package manager

import "database/sql"

// Item manages items.
type Item struct {
	ID          int64
	Title       string
	Description string
}

// CreateItem creates a new item.
func (item *Item) CreateItem(db *sql.DB) {
	res, err := db.Exec(`insert into item values(title, description)`, item.Title, item.Description)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	item.ID = id
}

// ReadItem reads an item.
func (item *Item) ReadItem(db *sql.DB, ID string) (found bool) {
	stmt, err := db.Prepare(`
		select id, title, description
		from item where id = ?`)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	switch err := stmt.QueryRow(ID).Scan(&item.ID, &item.Title, &item.Description); {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		panic(err)
	}

	return true
}

// UpdateItem updates an item.
func (item *Item) UpdateItem(db *sql.DB) {
	_, err := db.Exec(`
		update item set title = ?, description = ?
		where id = ?`, item.Title, item.Description)
	if err != nil {
		panic(err)
	}
}

// DeleteItem deletes an item.
func (item *Item) DeleteItem(db *sql.DB) {
	_, err := db.Exec(`delete from item where id = ?`, item.ID)
	if err != nil {
		panic(err)
	}
}

// GetItems gets all items.
func GetItems(db *sql.DB) []Item {
	items := []Item{}

	sql := `
	select id, title, description
	from item
	order by id desc`
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		item := Item{}
		if err := rows.Scan(&item.ID, &item.Title, &item.Description); err != nil {
			panic(err)
		}

		items = append(items, item)
	}

	return items
}
