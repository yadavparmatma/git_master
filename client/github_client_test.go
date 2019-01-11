package client

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	c "github.com/yadavparmatma/git_master/config"
	"github.com/yadavparmatma/git_master/model"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGitHub_CreateUrl(t *testing.T) {
	config := &c.Config{
		Host:      "http://host",
		Parameter: "repo",
	}

	user := "yadav"
	expectedUrl := "http://host/yadav/repo?per_page=20"
	hub := GitHub{}
	actual := hub.CreateUrl(config, user)

	assert.Equal(t, expectedUrl, actual)
}

func TestGitHub_Fetch_Response_Success(t *testing.T) {
	repo := model.Repo{Name: "A", Language: "B",}
	expected := append(make([]model.Repo, 0), repo)
	body, _ := json.Marshal(expected)

	server := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				n, _ := io.WriteString(w, string(body))
				fmt.Println(n)
			}))

	defer server.Close()

	gc := &GitHub{}
	actual := gc.Fetch(server.URL)

	assert.Equal(t, expected, actual)
}
