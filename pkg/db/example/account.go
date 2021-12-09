package example

import "github.com/dk-sirius/db-builder/pkg/db/example/tmp"

type A func()

// Account 账户
//@def primary f_id
//@def unique_index i_userID f_userID
//@def index i_name f_name
//@def unique_index i_userID_name f_userID f_name
type Account struct {
	tmp.TimestampDeep
	TimestampAt
	AccountID
	Name     string `db:"f_name,size=50,default=''" json:"name"`
	Password string `db:"f_password"`
	UserID   uint64 `db:"f_userID"`
	Nickname string `db:"f_nick_name,size=90,default=''"`
}

type AccountID struct {
	ID uint64 `db:"f_id,autoincrement"`
}

// constant
const name = 123
