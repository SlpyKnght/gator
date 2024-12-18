package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct{
	DbUrl string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read()Config{
	path, err := getConfigFilePath()
	if err != nil{
		fmt.Println(err)
		os.Exit(2)
	}
	file, err := os.Open(path)
	if err != nil{
		fmt.Println(err)
		os.Exit(3)
	}
	defer file.Close()
	var conf Config
	json.NewDecoder(file).Decode(&conf)

	return conf
}

func (conf *Config)SetUser(user string)error{
	conf.CurrentUserName = user
	path, err := getConfigFilePath()
	if err != nil{
		return err
	}
	file, err := os.OpenFile(path, os.O_RDWR | os.O_CREATE, os.ModePerm)
	if err != nil{
		return err
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(&conf)
	if err != nil{
		fmt.Printf("failed to write updated config to file:\n%v\n", err)
		return err
	}
	return nil
}

func getConfigFilePath()(string,error){
	home, err := os.UserHomeDir()
	if err != nil{
		return "", err
	}
	path := home + "/" + configFileName
	return path, nil
}
