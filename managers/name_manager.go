package managers

import (
	"bytes"
	"strconv"
)

type NameManager struct {
	thaiName       string
	numberOfChar   map[string]string
	numberOfShadow map[string]string
	reangName      string
	leksat         int
	shadow         int
}

func NewNameManager(thaiName string) *NameManager {
	return &NameManager{
		thaiName: thaiName,
		numberOfChar: map[string]string{
			"ก": "1", "ด": "1", "ถ": "1", "ท": "1", "ภ": "1", "ฤ": "1", "ฦ": "1", "า": "1", "ำ": "1", "ุ": "1", "่": "1",
			"ข": "2", "ง": "2", "ช": "2", "บ": "2", "ป": "2", "เ": "2", "แ": "2", "ู": "2", "้": "2",
			"ฆ": "3", "ต": "3", "ฑ": "3", "ฒ": "3", "๋": "3",
			"ค": "4", "ธ": "4", "ญ": "4", "ร": "4", "ษ": "4", "ะ": "4", "โ": "4", "ั": "4", "ิ": "4",
			"ฉ": "5", "ฌ": "5", "ณ": "5", "น": "5", "ม": "5", "ห": "5", "ฎ": "5", "ฮ": "5", "ฬ": "5", "ึ": "5",
			"จ": "6", "ล": "6", "ว": "6", "อ": "6", "ใ": "6",
			"ซ": "7", "ศ": "7", "ส": "7", "๊": "7", "ี": "7", "ื": "7",
			"ผ": "8", "ฝ": "8", "พ": "8", "ฟ": "8", "ย": "8", "็": "8",
			"ฏ": "9", "ฐ": "9", "ไ": "9", "์": "9",
			"a": "1", "i": "1", "j": "1", "q": "1", "y": "1",
			"A": "1", "I": "1", "J": "1", "Q": "1", "Y": "1",
			"b": "2", "k": "2", "r": "2",
			"B": "2", "K": "2", "R": "2",
			"c": "3", "g": "3", "l": "3", "s": "3",
			"C": "3", "G": "3", "L": "3", "S": "3",
			"d": "4", "m": "4", "t": "4",
			"D": "4", "M": "4", "T": "4",
			"e": "5", "h": "5", "n": "5", "x": "5",
			"E": "5", "H": "5", "N": "5", "X": "5",
			"u": "6", "v": "6", "w": "6",
			"U": "6", "V": "6", "W": "6",
			"o": "7", "z": "7",
			"O": "7", "Z": "7",
			"f": "8", "p": "8",
			"F": "8", "P": "8",
		},
		numberOfShadow: map[string]string{
			"อ": "6", "ะ": "6", "า": "6", "ิ": "6", "ี": "6", "ุ": "6", "ู": "6", "เ": "6", "โ": "6",
			"ก": "15", "ข": "15", "ค": "15", "ฆ": "15", "ง": "15",
			"จ": "8", "ฉ": "8", "ช": "8", "ซ": "8", "ฌ": "8", "ญ": "8",
			"ฎ": "17", "ฏ": "17", "ฐ": "17", "ฑ": "17", "ฒ": "17", "ณ": "17",
			"บ": "19", "ป": "19", "ผ": "19", "ฝ": "19", "พ": "19", "ฟ": "19", "ภ": "19", "ม": "19",
			"ศ": "21", "ษ": "21", "ส": "21", "ห": "21", "ฬ": "21", "ฮ": "21",
			"ด": "10", "ต": "10", "ถ": "10", "ท": "10", "ธ": "10", "น": "10",
			"ย": "12", "ร": "12", "ล": "12", "ว": "12",
		},
	}

}

func (manager *NameManager) GetThaiName() string {
	return manager.thaiName
}

func (manager *NameManager) GetReangName() string {
	return manager.reangName
}
func (manager *NameManager) SetReangName() {
	var b bytes.Buffer
	namex := []rune(manager.thaiName)
	for _, name := range namex {
		for k, number := range manager.numberOfChar {
			if string(name) == k {
				b.WriteString(number)
				break
			}
		}
	}

	manager.reangName = b.String()
}

func (manager *NameManager) SetLeksat() {
	manager.leksat = 0
	for _, charX := range []rune(manager.reangName) {
		num, _ := strconv.Atoi(string(charX))
		manager.leksat += num
	}
}

func (manager *NameManager) GetLeksat() int {
	return manager.leksat
}

func (manager *NameManager) SetShadow() {
	manager.shadow = 0
	namex := []rune(manager.thaiName)
	for _, name := range namex {
		for k, number := range manager.numberOfShadow {
			if string(name) == k {
				numShar, _ := strconv.Atoi(number)
				manager.shadow += numShar
				break
			}
		}
	}

}

func (manager *NameManager) GetShadow() int {
	return manager.shadow
}
