package executor

import (
	"github.com/yadavparmatma/git_master/client"
	"github.com/yadavparmatma/git_master/config"
	"github.com/yadavparmatma/git_master/model"
)

type TaskExecutor interface {
	Execute(string, chan []model.Repo)
}

type Task struct {
	TaskExecutor
	Config *config.Config
}

func (task *Task) Execute(user string, responseChannel chan []model.Repo) {
	gc := &client.GitHub{}
	url := gc.CreateUrl(task.Config, user)
	repos := gc.Fetch(url)
	responseChannel <- repos
}
