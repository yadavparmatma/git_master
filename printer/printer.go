package printer

import (
	"fmt"
	"github.com/yadavparmatma/git_master/model"
)

func PrintRepo(repos []model.Repo, quit chan int) {
	fmt.Println("Printing Repos......")
	for i := range repos {
		repo := repos[i]
		fmt.Println(i+1, ".")
		fmt.Println("	", repo.Name)
		fmt.Println("	", repo.Language)
	}
	defer done(quit)
}

func done(quit chan int) {
	quit <- 0
}
