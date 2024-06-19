package file_data_base

import (
	"encoding/json"
	conf "headquarters/config_manager"
	"io"
	"os"
	"sync"
	"time"
)

const dateLayout = "2006-01-02"

type Record struct {
	UserId   string    `json:"userId"`
	Time     time.Time `json:"time"`
	Address  string    `json:"address"`
	Attempts int       `json:"attempts"`
}

type DataBaseInterface interface {
	AddUser(user *conf.User)
	GetUser(userId string) *conf.User
	AddRecord(record *Record)
}

type DataBase struct {
	StatsFileName string
	UserFileName  string
	mutex         sync.RWMutex
	records       map[string][]*Record
	statsFile     *os.File
	usersConfig   *conf.ConfigManager
}

func NewDataBase(userFileName string, statsFileName string) (*DataBase, error) {
	var db = new(DataBase)
	db.StatsFileName = statsFileName
	db.UserFileName = userFileName
	db.mutex = sync.RWMutex{}
	db.records = make(map[string][]*Record)
	db.usersConfig = conf.NewConfigManager(db.UserFileName)
	err := db.usersConfig.ReadConfig()

	if err != nil {
		return nil, err
	}

	err = db.readStats()
	if err != nil {
		return nil, err
	} else {
		return db, nil
	}
}

func (db *DataBase) readStats() error {
	file, err := os.OpenFile(db.StatsFileName, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&db.records)

	if err != nil && err != io.EOF {
		return err
	}

	return file.Close()
}

func (db *DataBase) updateStatsFile() error {
	file, err := os.OpenFile(db.StatsFileName, os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	var res []byte
	res, err = json.MarshalIndent(db.records, "", "\t")
	if err != nil {
		return err
	}
	_, err = file.Write(res)
	if err != nil {
		return err
	}
	return file.Close()
}

func (db *DataBase) AddRecord(record *Record) error {

	date := record.Time.Format(dateLayout)

	db.mutex.Lock()
	recordsForThisDate := db.records[date]

	if recordsForThisDate == nil {
		db.records[date] = []*Record{record}
	} else {
		db.records[date] = append(recordsForThisDate, record)

	}

	err := db.updateStatsFile()

	db.mutex.Unlock()
	return err
}

func (db *DataBase) GetUser(userId string) *conf.User {
	return db.usersConfig.GetUser(userId)
}

func (db *DataBase) AddUser(user *conf.User) error {
	db.mutex.Lock()

	db.usersConfig.AddUser(*user)
	err := db.usersConfig.WriteConfig()

	db.mutex.Unlock()

	return err
}
