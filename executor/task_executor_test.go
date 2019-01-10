package executor

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yadavparmatma/git_master/config"
	"github.com/yadavparmatma/git_master/model"
	"testing"
)

type GitHub struct {
	mock.Mock
}

func (m *GitHub) Fetch(url string) []model.Repo {
	res := m.Called(url)
	return res.Get(0).([]model.Repo)
}

func (m *GitHub) CreatUrl(config *config.Config, user string) string {
	res := m.Called(config, user)
	return res.Get(0).(string)
}

func TestExecute(t *testing.T) {
	responseChannel := make(chan []model.Repo)

	task := &Task{
		Config: &config.Config{
			Host:      "http://host",
			Parameter: "repo",
		},
	}

	expectedUrl := "http://host/yadav/repo?per_page=20"
	var expectedRepos []model.Repo
	hub := GitHub{}

	hub.On("CreateUrl", task.Config, "yadav").Return(expectedUrl)
	hub.On("Fetch", expectedUrl).Return(expectedRepos)

	task.Execute("yadav", responseChannel)

	assert.Equal(t, expectedRepos, <-responseChannel)
}
