package interactor

import (
	"clean-golang/app/interfaces/repository"
)

var repo repository.Repository

func Index() {
	_, usersErr := repo.All()
	if usersErr != nil {
		return
	}
}
