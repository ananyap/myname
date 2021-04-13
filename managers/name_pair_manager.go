package managers

import (
	"github.com/myname/names"
	"github.com/myname/numbers"
	"sort"
	"strconv"
)

type NamePairManager struct {
	name          names.RealName
	namesTable    []names.RealName
	numbersMira   []numbers.Number
	plusShaNumD   []names.PlusTitle
	plusTitleNumD []names.PlusTitle
	listSatDShaD  []names.PlusTitle
}

func NewNamePairManager(myName names.RealName, namesTable []names.RealName, numbersMira []numbers.Number) *NamePairManager {
	return &NamePairManager{
		name:        myName,
		namesTable:  namesTable,
		numbersMira: numbersMira,
	}
}
func (manager *NamePairManager) GetListSatDShaD() []names.PlusTitle {

	return manager.listSatDShaD
}

func (manager *NamePairManager) SetListSatDShaD() {

	for _, NameObj := range manager.namesTable {
		shaPlus := NameObj.Shadow + manager.name.Shadow
		satPlus := NameObj.LeksatThai + manager.name.LeksatThai

		for _, numMiraObj := range manager.numbersMira {

			if strconv.Itoa(satPlus) == numMiraObj.PairNumber {

				if numMiraObj.PairType == "D10" || numMiraObj.PairType == "D8" || numMiraObj.PairType == "D5" {

					for _, numMiraObjCheckShaD := range manager.numbersMira {

						if strconv.Itoa(shaPlus) == numMiraObjCheckShaD.PairNumber {

							if numMiraObjCheckShaD.PairType == "D10" || numMiraObjCheckShaD.PairType == "D8" || numMiraObjCheckShaD.PairType == "D5" {
								manager.listSatDShaD = append(manager.listSatDShaD, names.PlusTitle{
									Surname:      NameObj,
									NumSatPlus:   satPlus,
									NumShaPlus:   shaPlus,
									ScoreSatPlus: numMiraObj.PairPoint,
								})

							}
							break
						}
					}

				}

				break
			}
		}
	}

	sort.SliceStable(manager.listSatDShaD, func(i, j int) bool {
		return manager.listSatDShaD[i].ScoreSatPlus > manager.listSatDShaD[j].ScoreSatPlus
	})

}

func (manager *NamePairManager) SetPlusShaNumD() {

	for _, NameObj := range manager.namesTable {
		shaPlus := NameObj.Shadow + manager.name.Shadow

		for _, numMiraObj := range manager.numbersMira {

			if strconv.Itoa(shaPlus) == numMiraObj.PairNumber {

				if numMiraObj.PairType == "D10" || numMiraObj.PairType == "D8" || numMiraObj.PairType == "D5" {
					manager.plusShaNumD = append(manager.plusShaNumD, names.PlusTitle{
						Surname:    NameObj,
						NumSatPlus: shaPlus,
					})
				}
				break
			}
		}
	}

	sort.SliceStable(manager.plusShaNumD, func(i, j int) bool {
		return manager.plusShaNumD[i].ScoreSatPlus > manager.plusShaNumD[j].ScoreSatPlus
	})

}

func (manager *NamePairManager) SetPlusTitleNumD() {
	for _, NameObj := range manager.namesTable {
		xPlus := NameObj.LeksatThai + manager.name.LeksatThai

		for _, numMiraObj := range manager.numbersMira {

			if strconv.Itoa(xPlus) == numMiraObj.PairNumber {

				if numMiraObj.PairType == "D10" || numMiraObj.PairType == "D8" || numMiraObj.PairType == "D5" {
					manager.plusTitleNumD = append(manager.plusTitleNumD, names.PlusTitle{
						Surname:    NameObj,
						NumSatPlus: xPlus,
					})

				}
				break
			}
		}
	}

	sort.SliceStable(manager.plusTitleNumD, func(i, j int) bool {
		return manager.plusTitleNumD[i].ScoreSatPlus > manager.plusTitleNumD[j].ScoreSatPlus
	})
}

func (manager *NamePairManager) GetPlusTitleNumD() []names.PlusTitle {
	return manager.plusTitleNumD
}

func (manager *NamePairManager) GetPlusShaNumD() []names.PlusTitle {
	return manager.plusShaNumD
}
