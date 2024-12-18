package main

import (
	"fmt"
	"os"

	"github.com/slpyknght/gator/internal/config"
)

type state struct{
	conf *config.Config
}

type command struct{
	name string
	arguments []string
}

type commands struct{
	cmds map[string]func(*state, command) error
}

func main(){
	conf := config.Read()
	s := &state{conf: &conf}
	cmds := commands{cmds: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	args := os.Args
	if len(args) < 2{
		fmt.Println("invalid number of arguments")
		os.Exit(1)
	}
	cmd := command{name: args[1], arguments: args[2:]}
	err := cmds.run(s, cmd)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}

func handlerLogin(s *state, cmd command)error{
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("invalid number of arguments")
	}
	if err := s.conf.SetUser(cmd.arguments[0]); err != nil{
		return err
	}
	fmt.Println("user name has been updated.")
	return nil
}

func (c *commands)register(name string, f func(*state, command) error){
	c.cmds[name] = f
}

func (c *commands)run(s *state, cmd command)error{
	if fn, ok := c.cmds[cmd.name]; ok{
		return fn(s, cmd)
	}
	return fmt.Errorf("command does not exist")
}
