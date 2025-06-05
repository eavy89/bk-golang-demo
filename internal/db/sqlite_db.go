package db

import (
	"backend-go-demo/internal/model"
	"database/sql"
)

type SQLiteDB struct {
	conn *sql.DB
}

func (s *SQLiteDB) CreateUser(user *model.User) error {
	_, err := s.conn.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	return err
}

func (s *SQLiteDB) GetUserByUsername(username string) (*model.User, error) {
	row := s.conn.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	return user, err
}

func (s *SQLiteDB) CreatePurchase(p *model.Purchase) error {
	_, err := s.conn.Exec(
		`INSERT INTO purchases (user_id, item, quantity, price) VALUES (?, ?, ?, ?)`,
		p.UserID, p.Item, p.Quantity, p.Price,
	)
	return err
}

func (s *SQLiteDB) GetPurchasesByUserID(userID int) ([]model.Purchase, error) {
	rows, err := s.conn.Query(
		`SELECT id, item, quantity, price, created_at FROM purchases WHERE user_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var purchases []model.Purchase
	for rows.Next() {
		var p model.Purchase
		p.UserID = userID
		err := rows.Scan(&p.ID, &p.Item, &p.Quantity, &p.Price, &p.CreatedAt)
		if err != nil {
			continue
		}
		purchases = append(purchases, p)
	}

	return purchases, nil
}

func (s *SQLiteDB) Close() error {
	return s.conn.Close()
}

func initSchema(conn *sql.DB) error {

	// create a "users" table if not exists
	_, err := conn.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE,
		password TEXT
	)`)
	if err != nil {
		return err
	}

	// create a "purchases" table if not exists
	_, err = conn.Exec(`CREATE TABLE IF NOT EXISTS purchases (
   		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		item TEXT,
		quantity INTEGER,
		price REAL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`)
	if err != nil {
		return err
	}

	// TODO: To maintain a lot of sql tables we will use a "migration tool" like "GROM" or something else

	return nil
}

func NewSQLiteDB(path string) (*SQLiteDB, error) {
	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	// Initialize schema
	if err := initSchema(conn); err != nil {
		return nil, err
	}

	return &SQLiteDB{conn: conn}, nil
}
