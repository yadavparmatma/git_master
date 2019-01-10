package main

import (
	"github.com/yadavparmatma/git_master/config"
	"github.com/yadavparmatma/git_master/executor"
	"github.com/yadavparmatma/git_master/model"
	"github.com/yadavparmatma/git_master/printer"
	"time"
)

func main() {
	users := []string{"yadavparmatma", "last-stand", "adwin"}

	//TODO: Read config from config file and populate configuration
	c := &config.Config{
		Host:      "https://api.github.com/users",
		Parameter: "repos",
	}

	repoPrinter := &printer.RepoPrinter{}
	response := make(chan []model.Repo)

	task := &executor.Task{
		Config: c,
	}

	for _, user := range users {
		go task.Execute(user, response)
	}

	for {
		select {
		case repos := <-response:
			repoPrinter.Print(repos)
		case <-time.After(time.Second * 10):
			return
		}
	}
}
