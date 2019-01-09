package executor

import (
	"github.com/yadavparmatma/git_master/client"
	"github.com/yadavparmatma/git_master/config"
	"github.com/yadavparmatma/git_master/model"
	"sync"
)

type TaskExecutor interface {
	Execute(chan []model.Repo)
}

type Task struct {
	Config *config.Config
	Users  []string
}

func (task *Task) Execute(responseChannel chan []model.Repo) {
	urlChannel := make(chan string, len(task.Users))
	var wg sync.WaitGroup
	wg.Add(len(task.Users) * 2)

	gc := &client.GitHub{}
	for _, user := range task.Users {
		go gc.CreateUrl(task.Config, user, urlChannel, &wg)
		go gc.Fetch(<-urlChannel, responseChannel, &wg)
	}
	wg.Wait()
}
