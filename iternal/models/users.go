package models

type User struct {
	Id   uint64 `bun:"id,pk,autoincrement"`
	Name string `bun:"name,notnull"`
	Age  uint8  `bun:"age"`
	// Email    string `bun:"email,unique"`//TODO: вернуть. мешать дебагу будет
	Email    string `bun:"email"`
	Password string `bun:"password,notnull"`
}
