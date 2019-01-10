package executor

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yadavparmatma/git_master/config"
	"github.com/yadavparmatma/git_master/mocks"
	"github.com/yadavparmatma/git_master/model"
	"testing"
)

func TestExecute(t *testing.T) {
	responseChannel := make(chan []model.Repo)
	cl := mocks.Client{}

	task := &Task{
		Config: &config.Config{
			Host:      "http://host",
			Parameter: "repo",
		},
	}

	expectedUrl := "http://host/yadav/repo?per_page=20"
	var expectedRepos []model.Repo

	cl.On("CreateUrl", mock.Anything, mock.Anything).Return(expectedUrl)
	cl.On("Fetch", mock.Anything).Return(expectedRepos)

	task.Execute("yadav", responseChannel)

	assert.Equal(t, expectedRepos, <-responseChannel)
}
