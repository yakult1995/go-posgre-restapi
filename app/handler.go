package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

// "Hello World!!"をリターン
func Root(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	type RootMessage struct {
		Message string `json:"message"`
	}

	response := RootMessage{
		Message: "Hello World!!",
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

// Userリストをリターン
func ListUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// ユーザ数をカウント
	var userCount int
	if err := db.QueryRow("SELECT COUNT(*) FROM users LIMIT 1000").Scan(&userCount); err != nil {
		panic(err)
	}

	// 0ならば空jsonを返す
	if userCount == 0 {
		emp := []string{}
		if err := json.NewEncoder(w).Encode(emp); err != nil {
			panic(err)
		}
		return
	}

	// 全ユーザ抽出
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	var users []User
	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	if err := rows.Close(); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
}

// ユーザ情報を更新する
func UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId := p.ByName("id")

	// パラメータを受け取る
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 情報の更新
	sqlQuery, err := db.Prepare("UPDATE users SET name = $2, email = $3 WHERE id = $1 RETURNING *")
	err = sqlQuery.QueryRow(userId, user.Name, user.Email).
		Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		// 対象ユーザがいないとき　
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
			emp := []string{}
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest)
			if err := json.NewEncoder(w).Encode(emp); err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

// ユーザの情報を取得
func DetailUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId := p.ByName("id")

	var user User
	row := db.QueryRow("SELECT * FROM users WHERE id = $1", userId)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		// 対象ユーザがいないとき
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
			emp := []string{}
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest)
			if err := json.NewEncoder(w).Encode(emp); err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}

}

// ユーザ作成
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ユーザ作成
	query, err := db.Prepare("INSERT INTO users(name, email) VALUES($1, $2) RETURNING *")
	if err != nil {
		panic(err)
	}

	// 作成したユーザの情報取得
	err = query.QueryRow(user.Name, user.Email).
		Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

// ユーザ消去
func DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId := p.ByName("id")
	_ = db.QueryRow("DELETE FROM users WHERE id = $1", userId)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNoContent)
}
