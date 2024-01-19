package server

import (
	yahooapi "dashboard/yahooApi"
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	*sql.DB

	path string
}

func (self *Database) InitDb() (err error) {
	_, fileErr := os.Stat(self.path)
	db, err := sql.Open("sqlite3", self.path)
	if err != nil {
		return
	}

	self.DB = db

	if fileErr != nil {
		err = self.Setup()
	}

	return
}

func (self *Database) Setup() (err error) {
	log.Println("setting up new database")
	setup := `
	CREATE TABLE stocks (
		symbol		TEXT		PRIMARY KEY,
		shortname	TEXT,
		longname	TEXT,
		exchange	TEXT,
		instrument	TEXT,
		score		REAL,

		notify		BOOLEAN		DEFAULT FAlSE
	)
	`

	_, err = self.Exec(setup)

	return
}

func (self *Database) NotificationSymbols() (symbols []yahooapi.Symbol, err error) {
	query := `
	select * from stocks
	where notify = true
	`
	rows, err := self.Query(query)
	defer rows.Close()

	for rows.Next() {
		var symbol yahooapi.Symbol
		err = rows.Scan(&symbol.Symbol, &symbol.ShortName, &symbol.LongName, &symbol.Exchange, &symbol.InstrumentType, &symbol.Score, &symbol.DoNotify)
		if err != nil {
			return
		}
		symbols = append(symbols, symbol)
	}

	return
}

func (self *Database) AddSymbol(symbol yahooapi.Symbol) (err error) {
	query := `
	insert into stocks (symbol, shortname, longname, exchange, instrument, score)
	values (?, ?, ?, ?, ?, ?)
	`

	_, err = self.Exec(
		query,
		symbol.Symbol,
		symbol.ShortName,
		symbol.LongName,
		symbol.Exchange,
		symbol.InstrumentType,
		symbol.Score,
	)

	return
}

func (self *Database) Symbols() (symbols []yahooapi.Symbol, err error) {
	query := `select * from stocks`

	rows, err := self.Query(query)
	defer rows.Close()

	for rows.Next() {
		var symbol yahooapi.Symbol
		err = rows.Scan(&symbol.Symbol, &symbol.ShortName, &symbol.LongName, &symbol.Exchange, &symbol.InstrumentType, &symbol.Score, &symbol.DoNotify)
		if err != nil {
			return
		}

		symbols = append(symbols, symbol)
	}

	return
}
