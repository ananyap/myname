package managers

import (
	"database/sql"
	"github.com/myname/names"
	"github.com/myname/numbers"
)

type DbManager struct {
	db             *sql.DB
	realNames      []names.RealName
	numbersMiracle []numbers.Number
}

func NewDbManager(myDb *sql.DB) *DbManager {
	return &DbManager{db: myDb}
}

func (manager *DbManager) SetNumbersMiracle() {

	results, err := manager.db.Query("SELECT pairnumber, pairtype, pairpoint FROM numbers order by pairnumberid ")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var myNumber numbers.Number
		err = results.Scan(&myNumber.PairNumber, &myNumber.PairType, &myNumber.PairPoint)
		if err != nil {
			panic(err.Error())
		}

		manager.numbersMiracle = append(manager.numbersMiracle, myNumber)

	}
}

func (manager *DbManager) GetNumberMiracle() []numbers.Number {
	return manager.numbersMiracle
}

func (manager *DbManager) SetRealNames() {

	results, err := manager.db.Query("SELECT thainame, reangthai, leksat_thai, shadow FROM realname order by nameid")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var myName names.RealName
		err = results.Scan(&myName.ThaiName, &myName.ReangThai, &myName.LeksatThai, &myName.Shadow)
		if err != nil {
			panic(err.Error())
		}

		manager.realNames = append(manager.realNames, myName)

	}
}

func (manager *DbManager) GetRealNames() []names.RealName {
	return manager.realNames
}
