// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package mysql

import ()

type Post struct {
	ID     string
	Title  string
	Body   string
	UserID string
}

type User struct {
	ID   string
	Name string
}