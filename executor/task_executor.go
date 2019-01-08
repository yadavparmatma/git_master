package executor

import (
	"github.com/yadavparmatma/git_master/client"
	"github.com/yadavparmatma/git_master/config"
	"github.com/yadavparmatma/git_master/model"
	"sync"
)

type TaskExecutor interface {
	FetchRepositories(resp chan []model.Repo)
}

type Task struct {
	Config *config.Config
	Users  []string
}

func (task *Task) FetchRepositories(response chan []model.Repo) {
	urls := make(chan string, len(task.Users))
	var wg sync.WaitGroup
	wg.Add(len(task.Users) * 2)

	gitClient := &client.GitHub{}
	for _, user := range task.Users {
		go gitClient.CreateUrl(task.Config, user, urls, &wg)
		go gitClient.FetchRepositories(<-urls, response, &wg)
	}
	wg.Wait()
}