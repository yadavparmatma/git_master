package main

import (
	"fmt"
	"github.com/yadavparmatma/git_master/config"
	"github.com/yadavparmatma/git_master/executor"
	"github.com/yadavparmatma/git_master/model"
	"github.com/yadavparmatma/git_master/printer"
)

func main() {
	users := []string{"yadavparmatma", "last-stand", "adwin"}

	//TODO: Read config from config file and populate configuration
	c := &config.Config{
		Host:      "https://api.github.com/users",
		Parameter: "repos",
		Users:     users,
	}

	response := make(chan []model.Repo, len(users))
	quit := make(chan string)

	go func() {
		task := &executor.Task{
			Config: c,
			Users:  users,
		}
		task.FetchRepositories(response)
		defer Done(quit)
	}()

	for {
		select {
		case repos := <-response:
			printer.PrintRepo(repos)
		case done := <-quit:
			fmt.Println(done)
			close(response)
			close(quit)
			return
		}
	}
}

func Done(quit chan string) {
	quit <- "Done"
}
