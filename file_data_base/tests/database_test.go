package tests

import (
	"headquarters/file_data_base"
	"testing"
	"time"
)

func TestAddRecord(t *testing.T) {
	db, err := file_data_base.NewDataBase("users.json", "stats.json")
	if err != nil {
		t.Fatalf(err.Error())
		return
	}

	var newRecord = &file_data_base.Record{UserId: "1234", Time: time.Now(), Address: "testAddr", Attempts: 1}
	err = db.AddRecord(newRecord)
	if err != nil {
		t.Fatalf(err.Error())
	}
}
