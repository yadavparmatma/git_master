package executor

import (
	"github.com/yadavparmatma/git_master/client"
	"github.com/yadavparmatma/git_master/config"
	"github.com/yadavparmatma/git_master/model"
	"sync"
)

type TaskExecutor interface {
	fetchRepositories(chan []model.Repo)
}

type Task struct {
	Config *config.Config
	Users  []string
}

func Execute(task *Task, resp chan []model.Repo) {
	task.fetchRepositories(resp)
}

func (task *Task) fetchRepositories(response chan []model.Repo) {
	urlChannel := make(chan string, len(task.Users))
	var wg sync.WaitGroup
	wg.Add(len(task.Users) * 2)

	gitClient := &client.GitHub{}
	for _, user := range task.Users {
		go client.CreateUrl(gitClient, task.Config, user, urlChannel, &wg)
		go client.Fetch(gitClient, <-urlChannel, response, &wg)
	}
	wg.Wait()
}
