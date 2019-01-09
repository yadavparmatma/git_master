package client

import (
	"encoding/json"
	"fmt"
	"github.com/yadavparmatma/git_master/config"
	"github.com/yadavparmatma/git_master/model"
	"io/ioutil"
	"net/http"
	"sync"
)

type Client interface {
	CreateUrl(*config.Config, string, chan string, *sync.WaitGroup)
	Fetch(string, chan []model.Repo, *sync.WaitGroup)
}

type GitHub struct {
}

func (g *GitHub) CreateUrl(config *config.Config, user string, urlChannel chan string, wg *sync.WaitGroup) {
	url := fmt.Sprintf("%v/%v/%v?per_page=20", config.Host, user, config.Parameter)
	urlChannel <- url
	defer wg.Done()
}

func (g *GitHub) Fetch(url string, resp chan []model.Repo, wg *sync.WaitGroup) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err.Error())
		resp <- []model.Repo{}
	}

	var repos []model.Repo
	err = json.Unmarshal(bytes, &repos)
	resp <- repos
	defer wg.Done()
}
