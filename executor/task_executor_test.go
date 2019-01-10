package executor

import (
	"github.com/stretchr/testify/assert"
	"github.com/yadavparmatma/git_master/config"
	"github.com/yadavparmatma/git_master/mocks"
	"github.com/yadavparmatma/git_master/model"
	"testing"
)

func TestExecute(t *testing.T) {
	client := new(mocks.Client)
	task := &Task{
		Config: &config.Config{
			Host:      "http://host",
			Parameter: "repo",
		},
		Client: client,
	}

	expectedUrl := "http://host/yadav/repo?per_page=20"
	var expectedRepos []model.Repo
	expectedRepos = append(expectedRepos, model.Repo{"turner", "java"})

	client.On("CreateUrl", task.Config, "yadav").Return(expectedUrl)
	client.On("Fetch", expectedUrl).Return(expectedRepos)

	response := task.Execute("yadav")
	assert.Equal(t, expectedRepos, response)
}
