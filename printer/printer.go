package printer

import (
	"fmt"
	"github.com/yadavparmatma/git_master/model"
)

type Printer interface {
	Print([]model.Repo)
}

type RepoPrinter struct {
	Printer
}

func (repoPrinter *RepoPrinter) Print(repos []model.Repo) {
	fmt.Println("Printing Repos......")
	for i := range repos {
		repo := repos[i]
		fmt.Println(i+1, ".")
		fmt.Println("	", repo.Name)
		fmt.Println("	", repo.Language)
	}
}
