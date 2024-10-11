package modles

type User struct {
	Id       uint64 `bun:"id,pk,autoincrement"`
	Name     string `bun:"name,notnull"`
	Age      uint8  `bun:"age"`
	Email    string `bun:"email,unique"`
	Password string `bun:"password,notnull"`
}
