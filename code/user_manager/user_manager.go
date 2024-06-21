package user_manager

import (
	"encoding/json"
	"io"
	"os"
)

type User struct {
	UserId   int64  `json:"userId"`
	UserName string `json:"userName"`
}

type Config struct {
	Users []User `json:"users"`
}

type configMap interface {
	InConfig(userId string) bool
	AddUser(user User)
	GetUser(userId string) *User
}

type ConfigInterface interface {
	configMap
	ReadConfig() error
	WriteConfig() error
}

type ConfigManager struct {
	config   Config
	FileName string
	file     *os.File
}

func NewConfigManager(fileName string) *ConfigManager {
	return &ConfigManager{FileName: fileName, file: nil}
}

func (configManager *ConfigManager) ReadConfig() error {
	file, err := os.OpenFile(configManager.FileName, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configManager.config)
	if err != nil && err != io.EOF {
		return err
	}
	return file.Close()
}

func (configManager *ConfigManager) WriteConfig() error {
	file, err := os.OpenFile(configManager.FileName, os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	var res []byte
	res, err = json.MarshalIndent(configManager.config, "", "\t")
	if err != nil {
		return err
	}
	_, err = file.Write(res)
	if err != nil {
		return err
	}
	return file.Close()
}

func (configManager *ConfigManager) AddUser(user User) {
	configManager.config.Users = append(configManager.config.Users, user)
}

func (configManager *ConfigManager) GetUser(userId int64) *User {
	for _, user := range configManager.config.Users {
		if user.UserId == userId {
			return &user
		}
	}
	return nil
}

func (configManager *ConfigManager) InConfig(userId int64) bool {
	for _, user := range configManager.config.Users {
		if user.UserId == userId {
			return true
		}
	}
	return false
}
