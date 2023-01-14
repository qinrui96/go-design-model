package pkg

type DbAdapter interface {
	WriteDB()
}

type SqlAdapter struct {
	sql *Sql
	DbAdapter
}

func (s *SqlAdapter) WriteDB() {
	s.sql.WriteSql()
}

type TiDBAdapter struct {
	tiDB *TiDB
	DbAdapter
}

func (t *TiDBAdapter) WriteDB() {
	t.tiDB.WriteTiDB()
}

func NewTiDBAdapter(db *TiDB) *TiDBAdapter {
	return &TiDBAdapter{tiDB: db}
}

func NewSqlAdapter(db *Sql) *SqlAdapter {
	return &SqlAdapter{sql: db}
}
