package example

type TimestampAt struct {
	CreatedTime uint64 `db:"created_time,default='0'"`
	UpdatedTime uint64 `db:"updated_time,default='0'"`
	DeleteTime  uint64 `db:"delete_time,default='0'"`
}
