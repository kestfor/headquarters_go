package tests

import (
	conf "headquarters/config_manager"
	"testing"
)

func TestConfig(t *testing.T) {
	var config = conf.NewConfigManager("users.json")
	err := config.ReadConfig()

	if err != nil {
		t.Fatalf("can't create config: %s", err.Error())
		return
	}
	newUser := conf.User{UserId: "123", UserName: "1235"}
	config.AddUser(newUser)
	if !config.InConfig(newUser.UserId) {
		t.Fatalf("new user wasn't added")
	}

	err = config.WriteConfig()
	if err != nil {
		t.Fatalf(err.Error())
	}
}
