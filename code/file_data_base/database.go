package file_data_base

import (
	"bufio"
	"encoding/json"
	conf "headquarters/code/user_manager"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

const dateLayout = "2006-01-02"

type Record struct {
	UserId   int64     `json:"userId"`
	Time     time.Time `json:"time"`
	Address  string    `json:"address"`
	Attempts int       `json:"attempts"`
}

type User interface {
	UserId() int64
	UserName() string
}

type DataBaseInterface interface {
	AddUser(user *User) error
	GetUser(userId int64) *User
	AddRecord(record *Record) error
	AddPhrase(phrase string) error
	Users() []User
	Contains(userId int64) bool
}

type DataBase struct {
	StatsFileName   string
	UserFileName    string
	PhrasesFileName string
	mutex           sync.RWMutex
	records         map[string][]*Record
	statsFile       *os.File
	usersConfig     *conf.ConfigManager
	Phrases         []string
}

func NewDataBase(userFileName string, statsFileName string, phrasesFileName string) (*DataBase, error) {
	var db = new(DataBase)
	db.StatsFileName = statsFileName
	db.UserFileName = userFileName
	db.PhrasesFileName = phrasesFileName
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
	}

	err = db.readPhrases()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db *DataBase) Contains(userId int64) bool {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	return db.usersConfig.InConfig(userId)
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

func (db *DataBase) GetUser(userId int64) User {
	return db.usersConfig.GetUser(userId)
}

func (db *DataBase) AddUser(user User) error {
	db.mutex.Lock()

	if !db.usersConfig.InConfig(user.UserId()) {
		db.usersConfig.AddUser(conf.TelegramUser{user.UserId(), user.UserName()})
	}
	err := db.usersConfig.WriteConfig()

	db.mutex.Unlock()

	return err
}

func (db *DataBase) readPhrases() error {
	file, err := os.OpenFile(db.PhrasesFileName, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(file)
	for {
		line, _ := reader.ReadString('\n')
		if line == "" {
			break
		}
		endIndex := strings.Index(line, "\n")
		if endIndex == -1 {
			endIndex = len(line)
		}
		line = line[:endIndex]
		db.Phrases = append(db.Phrases, line)
	}
	return file.Close()
}

func (db *DataBase) writePhrases() error {
	file, err := os.OpenFile(db.PhrasesFileName, os.O_WRONLY|os.O_TRUNC, os.ModePerm)

	if err != nil {
		return nil
	}
	writer := bufio.NewWriter(file)
	for _, phrase := range db.Phrases {
		_, err := writer.WriteString(phrase + "\n")
		if err != nil {
			return err
		}
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return file.Close()
}

func (db *DataBase) AddPhrase(phrase string) error {
	db.mutex.Lock()

	db.Phrases = append(db.Phrases, phrase)
	err := db.writePhrases()

	db.mutex.Unlock()
	return err
}

func (db *DataBase) Users() []conf.TelegramUser {
	return db.usersConfig.Users()
}
