package client

import (
	"encoding/json"
	"fmt"
	"github.com/yadavparmatma/git_master/config"
	"github.com/yadavparmatma/git_master/model"
	"io/ioutil"
	"net/http"
)

type Client interface {
	CreateUrl(*config.Config, string) string
	Fetch(string) []model.Repo
}

type GitHub struct {
	Client
}

func (g *GitHub) CreateUrl(config *config.Config, user string) string {
	return fmt.Sprintf("%v/%v/%v?per_page=20", config.Host, user, config.Parameter)
}

func (g *GitHub) Fetch(url string) []model.Repo {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err.Error())
		return []model.Repo{}
	}

	var repos []model.Repo
	err = json.Unmarshal(bytes, &repos)
	return repos
}
