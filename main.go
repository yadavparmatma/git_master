package main

import (
	"github.com/yadavparmatma/git_master/client"
	"github.com/yadavparmatma/git_master/config"
	"github.com/yadavparmatma/git_master/executor"
	"github.com/yadavparmatma/git_master/printer"
	"sync"
)

func main() {
	users := []string{"yadavparmatma", "last-stand", "adwin"}

	//TODO: Read config from config file and populate configuration
	c := &config.Config{
		Host:      "https://api.github.com/users",
		Parameter: "repos",
	}

	repoPrinter := &printer.RepoPrinter{}
	task := &executor.Task{
		Config: c,
		Client: new(client.GitHub),
	}

	var wg sync.WaitGroup
	for _, user := range users {
		wg.Add(1)
		go func(u string) {
			execute := task.Execute(u)
			repoPrinter.Print(execute)
			defer wg.Done()
		}(user)
	}
	wg.Wait()
}
