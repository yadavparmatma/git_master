package printer

import (
	"fmt"
	"github.com/yadavparmatma/git_master/model"
)

type Printer interface {
	print([]model.Repo)
}

type RepoPrinter struct {
}

func Print(printer Printer, repos []model.Repo) {
	printer.print(repos)
}

func (repoPrinter RepoPrinter) print(repos []model.Repo) {
	fmt.Println("Printing Repos......")
	for i := range repos {
		repo := repos[i]
		fmt.Println(i+1, ".")
		fmt.Println("	", repo.Name)
		fmt.Println("	", repo.Language)
	}
}
