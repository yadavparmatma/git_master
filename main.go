package main

import (
	"fmt"
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
		Users:     users,
	}

	response := make(chan []model.Repo, len(users))
	quit := make(chan int)

	go func() {
		task := &executor.Task{
			Config: c,
			Users:  users,
		}
		go task.FetchRepositories(response)
	}()

	for {
		select {
		case repos := <-response:
			printer.PrintRepo(repos, quit)
		case <-time.After(60 * time.Second):
			fmt.Println("Fetch timeout..")
			close(response)
			close(quit)
			return
		case <-quit:
			fmt.Println("Done...")
			close(response)
			close(quit)
			return
		}
	}
}
