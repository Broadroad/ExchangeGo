package scheduler

import (
	"database/sql"
	"net/http"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ExchangeGo/exchange/fcoin"
)

var db *DB

func init() {
	db, err = sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
		return
	}
	defer db.Close()
}
func TestNewFCoin(t *testing.T) {
	fc := fcoin.NewFCoin(&http.Client{}, "", "")
	sconfig := SchedulerConfig{fc: fc}
	NewScheduler(sconfig)
}
