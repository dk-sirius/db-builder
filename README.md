db-builder

db-builder 主要完成的是通过对象定义数据库，然后转换成postgresl（当前仅支持pg），创建表生成的schema，包括，主键、索引、字段等等。并在对应文件目录下生成目标文件__generate.go 文件，该文件转换后是所有sql信息以及对象关联的方法（表名称、字段映射的数据库字段、migate方法等）

一、数据表抽象定义

./pkg/db/table

    /**
        // default each field value is not null
    
    	// comments
        // @def constraintKey  constraintValues ...
        // @def index indexName RelativeField ....
        // @def unique_index indexName RelativeField ...
    	type FieldDef struct {
    		Field1 type `db:"fieldName,constraint(size=10\default=''...)"`   // see detail token.DbToken
    		....
        }
    */

二、数据表token 定义

    const (
    	ddlKeyBeg DbToken = iota
    	Indent
    	Db
    	Def
    	Index
    	Unique
    	ddlKeyEnd
    
    	ddlCheckBeg
    	Primary
    	Size
    	ValueDefault
    
    	ddlSequenceBeg
    	AutoIncr
    	ddlSequenceEnd
    	ddlCheckEnd
    )
     ……
    // sequence
    var dTokens = [...]string{
    	Indent:       "Indent",
    	Db:           "db",
    	Def:          "@def",
    	Index:        "index",
    	Unique:       "unique_index",
    	Primary:      "primary",
    	Size:         "size",
    	ValueDefault: "default",
    	AutoIncr:     "autoincrement",
    }
    ……

三、使用（github/dk-sirius/db-decl）

数据库表对象定义文件：account.go ,目录结构如下

    /example
        /tmp
        			tmp.go
        account.go ----------- 表定义对象
        base_db.go
      
定义对象内容

    //go:generate db-decl gen -n sirius -m Account
    
    // Account 账户
    //@def primary f_id
    //@def unique_index i_userID f_userID
    //@def index i_name f_name
    //@def unique_index i_userID_name f_userID f_name
    type Account struct {
    	tmp.TimestampDeep
    	TimestampAt
    	AccountID
    	Name     string `db:"f_name,size=50,default=''"`
    	Password string `db:"f_password"`
    	UserID   uint64 `db:"f_userID"`
    	Nickname string `db:"f_nick_name,size=90,default=''"`
    }
    
    type AccountID struct {
    	ID uint64 `db:"f_id,autoincrement"`
    }

执行后生成目录

    /example
        /tmp
        			tmp.go
        account.go ----------- 表定义对象
        account__generated.go
        base_db.go

执行后生成文件 

    package example
    
    import "strings"
    import "database/sql"
    import "github.com/jmoiron/sqlx"
    
    var _account = `CREATE SEQUENCE t_account_f_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1;
    ALTER SEQUENCE t_account_f_id_seq OWNED BY t_account.f_id;
    
    CREATE TABLE IF NOT EXISTS t_account ( f_deep bigint NOT NULL DEFAULT '0',created_time bigint NOT NULL DEFAULT '0',updated_time bigint NOT NULL DEFAULT '0',delete_time bigint NOT NULL DEFAULT '0',f_name varchar(50) NOT NULL DEFAULT '',f_password text NOT NULL,f_userID bigint NOT NULL,f_nick_name varchar(90) NOT NULL DEFAULT '',f_id bigint NOT NULL DEFAULT nextval('t_account_f_id_seq'::regclass),PRIMARY KEY(f_id) );
    
    ALTER TABLE t_account ADD IF NOT EXISTS f_deep bigint NOT NULL DEFAULT '0';
    ALTER TABLE t_account ADD IF NOT EXISTS created_time bigint NOT NULL DEFAULT '0';
    ALTER TABLE t_account ADD IF NOT EXISTS updated_time bigint NOT NULL DEFAULT '0';
    ALTER TABLE t_account ADD IF NOT EXISTS delete_time bigint NOT NULL DEFAULT '0';
    ALTER TABLE t_account ADD IF NOT EXISTS f_name varchar(50) NOT NULL DEFAULT '';
    ALTER TABLE t_account ADD IF NOT EXISTS f_password text NOT NULL;
    ALTER TABLE t_account ADD IF NOT EXISTS f_userID bigint NOT NULL;
    ALTER TABLE t_account ADD IF NOT EXISTS f_nick_name varchar(90) NOT NULL DEFAULT '';
    ALTER TABLE t_account ADD IF NOT EXISTS f_id bigint NOT NULL DEFAULT nextval('t_account_f_id_seq'::regclass);
    
    CREATE UNIQUE INDEX IF NOT EXISTS i_userID ON t_account (f_userID);
    CREATE INDEX IF NOT EXISTS i_name ON t_account (f_name);
    CREATE UNIQUE INDEX IF NOT EXISTS i_userID_name ON t_account (f_userID,f_name);
    
    ALTER UNIQUE INDEX IF NOT EXISTS i_userID ON t_account (f_userID);
    ALTER INDEX IF NOT EXISTS i_name ON t_account (f_name);
    ALTER UNIQUE INDEX IF NOT EXISTS i_userID_name ON t_account (f_userID,f_name);`
    
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
    func (t Account) IndexNameValue() []string {
    	return t.ToIndexSlice("f_name")
    }
    func (Account) UniqueIndexUserIDKey() string {
    	return "i_userID"
    }
    func (Account) UniqueIndexUserIDNameKey() string {
    	return "i_userID_name"
    }
    func (t Account) UniqueIndexUserIDValue() []string {
    	return t.ToIndexSlice("f_userID")
    }
    func (t Account) UniqueIndexUserIDNameValue() []string {
    	return t.ToIndexSlice("f_userID,f_name")
    }
    func (Account) ToIndexSlice(s string) []string {
    	return strings.Split(s, ",")
    }
    
