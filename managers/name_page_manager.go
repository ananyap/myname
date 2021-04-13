package managers

import (
	"log"
	"math"
)

type NamePageManager struct {
	showPerPage int
	amountItems int
	totalPage   int
}

func NewNamePageManager(amountItems int, showPerPage int) *NamePageManager {
	return &NamePageManager{amountItems: amountItems, showPerPage: showPerPage}
}

func (manager *NamePageManager) SetTotalPage() {
	//total := 0
	var pageRare float64
	pageRare = float64(manager.amountItems) / float64(manager.showPerPage)
	manager.totalPage = int(math.Ceil(pageRare))

	log.Println("showPerPages: ", manager.showPerPage)
	log.Println("amountItems: ", manager.amountItems)
	log.Println("pageRare: ", pageRare)
	log.Println("page: ", manager.totalPage)

}
