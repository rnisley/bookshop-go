package db

func CreateBook(title, author string, price float32) (int, error) {
	database := Connect().Db

	database.Exec(`
		INSERT INTO Books (title, author, price)
		VALUES (?, ?, ?);
	`, title, author, price)

	return GetBookId(title, author)
}

func GetBookId(title, author string) (int, error) {
	database := Connect().Db

	rows, err := database.Query(`
		SELECT (id) FROM Books
		WHERE title = ? and author = ?;
	`, title, author)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var bid int
	for rows.Next() {
		err = rows.Scan(&bid)
		if err != nil {
			return 0, err
		}
	}
	return bid, nil
}

func GetBookPrice(bid int) (float32, error) {
	database := Connect().Db

	rows, err := database.Query(`
		SELECT (price) FROM Books
		WHERE id = ?;
	`, bid)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var price float32
	for rows.Next() {
		err = rows.Scan(&price)
		if err != nil {
			return 0, err
		}
	}
	return price, nil
}
