package main

import (
	"fmt"
	"github.com/yadavparmatma/git_master/client"
	"github.com/yadavparmatma/git_master/config"
	"github.com/yadavparmatma/git_master/model"
	"sync"
)

func main() {
	gitClient := &client.GitHub{}
	users := []string{"yadavparmatma", "last-stand", "adwin"}

	//TODO: Read config from config file and populate configuration
	c := &config.Config{
		Host:      "https://api.github.com/users",
		Parameter: "repos",
	}

	var wg sync.WaitGroup
	wg.Add((len(users) * 2) + 1)

	urls := make(chan string, len(users))
	response := make(chan []model.Repo, len(users))

	go func() {
		for _, user := range users {
			//TODO: Execute these inside a task
			go gitClient.CreateUrl(c, user, urls, &wg)
			go gitClient.FetchRepositories(<-urls, response, &wg)
		}
		defer wg.Done()
	}()

	//TODO: Print response using goroutines
	fmt.Println("Repository Description....")
	for i := 0; i < len(users); i++ {
		repos := <-response
		for i := range repos {
			repo := repos[i]
			fmt.Println(i+1, ".")
			fmt.Println("	", repo.Name)
			fmt.Println("	", repo.Language)
		}
	}
	wg.Wait()
	close(urls)
	close(response)
}
