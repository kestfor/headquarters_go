package user_manager

import (
	"encoding/json"
	"io"
	"os"
)

type User interface {
	UserId() int64
	UserName() string
}

type TelegramUser struct {
	Id   int64  `json:"userId"`
	Name string `json:"userName"`
}

func (u *TelegramUser) UserId() int64 {
	return u.Id
}

func (u *TelegramUser) UserName() string {
	return u.Name
}

func NewTelegramUser(userId int64, userName string) *TelegramUser {
	return &TelegramUser{userId, userName}
}

type Config struct {
	Users []TelegramUser `json:"users"`
}

type configMap interface {
	InConfig(userId string) bool
	AddUser(user User)
	GetUser(userId string) *User
	Users() []TelegramUser
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

func (configManager *ConfigManager) AddUser(user TelegramUser) {
	configManager.config.Users = append(configManager.config.Users, user)
}

func (configManager *ConfigManager) GetUser(userId int64) *TelegramUser {
	for _, user := range configManager.config.Users {
		if user.UserId() == userId {
			return &user
		}
	}
	return nil
}

func (configManager *ConfigManager) InConfig(userId int64) bool {
	for _, user := range configManager.config.Users {
		if user.UserId() == userId {
			return true
		}
	}
	return false
}

func (configManager *ConfigManager) Users() []TelegramUser {
	res := make([]TelegramUser, len(configManager.config.Users), len(configManager.config.Users))
	copy(res, configManager.config.Users)
	return res
}
