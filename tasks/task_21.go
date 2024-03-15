package main

/*
За основу взяты примеры из статьи
https://habr.com/ru/articles/765468/
А также
https://pkg.go.dev/github.com/alexandergrom/go-patterns/Structural/Adapter#section-readme
*/

import (
	"fmt"
)

// Структуры со своими методами, которые надо адаптировать
type PostgreSQL struct {
}
func (p *PostgreSQL) SelectFromPGSQL() string {
	return fmt.Sprintf("Data from Postgres")
}


type MySQL struct {
}
func (m *MySQL) SelectFromMYSQL() string {
	return fmt.Sprintf("Data from MySQL")
}



// Интерфейс для адаптируемых баз даннных
type DBadapter interface {
	Select() string
}


// адаптер для postgres
type PgsqlAdapter struct {
	*PostgreSQL
}
func (adapter *PgsqlAdapter) Select() string {
	return adapter.PostgreSQL.SelectFromPGSQL()
}
func NewPgsqlAdapter(pg *PostgreSQL) DBadapter {
	return &PgsqlAdapter{pg}
}


// адаптер для mysql
type MysqlAdapter struct {
	*MySQL
}
func (adapter *MysqlAdapter) Select() string {
	return adapter.MySQL.SelectFromMYSQL()
}
func NewMysqlAdapter(msql *MySQL) DBadapter {
	return &MysqlAdapter{msql}
}




func main() {
	databases := [2]DBadapter{
		NewPgsqlAdapter(&PostgreSQL{}),
		NewMysqlAdapter(&MySQL{}),
	}

	for _, db := range databases {
		result := db.Select()
		fmt.Println(result)
	}

}