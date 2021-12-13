package example

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

var _account = `CREATE TABLE IF NOT EXISTS t_account ( f_deep bigint NOT NULL DEFAULT '0',created_time bigint NOT NULL DEFAULT '0',updated_time bigint NOT NULL DEFAULT '0',delete_time bigint NOT NULL DEFAULT '0',f_name varchar(50) NOT NULL DEFAULT '',f_password text NOT NULL,f_userID bigint NOT NULL,f_nick_name varchar(90) NOT NULL DEFAULT '',f_id bigserial NOT NULL,PRIMARY KEY(f_id) );

ALTER TABLE t_account ADD IF NOT EXISTS f_deep bigint NOT NULL DEFAULT '0';
ALTER TABLE t_account ADD IF NOT EXISTS created_time bigint NOT NULL DEFAULT '0';
ALTER TABLE t_account ADD IF NOT EXISTS updated_time bigint NOT NULL DEFAULT '0';
ALTER TABLE t_account ADD IF NOT EXISTS delete_time bigint NOT NULL DEFAULT '0';
ALTER TABLE t_account ADD IF NOT EXISTS f_name varchar(50) NOT NULL DEFAULT '';
ALTER TABLE t_account ADD IF NOT EXISTS f_password text NOT NULL;
ALTER TABLE t_account ADD IF NOT EXISTS f_userID bigint NOT NULL;
ALTER TABLE t_account ADD IF NOT EXISTS f_nick_name varchar(90) NOT NULL DEFAULT '';
ALTER TABLE t_account ADD IF NOT EXISTS f_id bigserial NOT NULL;

CREATE UNIQUE INDEX IF NOT EXISTS i_userID ON t_account (f_userID);
CREATE INDEX IF NOT EXISTS i_name ON t_account (f_name);
CREATE UNIQUE INDEX IF NOT EXISTS i_userID_name ON t_account (f_userID,f_name);`

func (Account) TableName() string {
	return "t_account"
}
func (Account) Migrate(db *sqlx.DB) sql.Result {
	return db.MustExec(_account)
}
func (Account) Schema() string {
	return _account
}
func (Account) PrimaryKeys() []string {
	return []string{"f_id"}
}
func (Account) FieldDeepKey() string {
	return "f_deep"
}
func (Account) FieldCreatedTimeKey() string {
	return "created_time"
}
func (Account) FieldUpdatedTimeKey() string {
	return "updated_time"
}
func (Account) FieldDeleteTimeKey() string {
	return "delete_time"
}
func (Account) FieldNameKey() string {
	return "f_name"
}
func (Account) FieldPasswordKey() string {
	return "f_password"
}
func (Account) FieldUserIDKey() string {
	return "f_userID"
}
func (Account) FieldNickNameKey() string {
	return "f_nick_name"
}
func (Account) FieldIdKey() string {
	return "f_id"
}
func (Account) IndexNameKey() string {
	return "i_name"
}
func (Account) IndexNameValue() string {
	return "f_name"
}
func (Account) UniqueIndexUserIDKey() string {
	return "i_userID"
}
func (Account) UniqueIndexUserIDNameKey() string {
	return "i_userID_name"
}
func (Account) UniqueIndexUserIDValue() string {
	return "f_userID"
}
func (Account) UniqueIndexUserIDNameValue() string {
	return "f_userID,f_name"
}
func (Account) FieldKeys() []string {
	return []string{"f_deep", "created_time", "updated_time", "delete_time", "f_name", "f_password", "f_userID", "f_nick_name", "f_id"}
}
func (Account) AutoFieldKeys() map[string]bool {
	return map[string]bool{"f_id": true}
}
