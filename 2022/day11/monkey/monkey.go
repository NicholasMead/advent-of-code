package monkey

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Monkey interface {
	GetItems() []uint64

	PopNext() (uint64, bool)
	Push(item uint64)
	Inspect(item uint64) uint64
	GetTarget(item uint64) int

	getDivider() uint64
}

func CommonBase(monkeys []Monkey) uint64 {
	var base uint64 = 1
	for _, m := range monkeys {
		base *= m.getDivider()
	}
	return base
}

func ParseMonkey(text string) (Monkey, error) {
	m := monkey{}

	// name
	m.name = regexp.MustCompile(`Monkey \d+`).FindString(text)[7:]

	// starting items
	itemsRaw := regexp.MustCompile(`<?[:,]\s\d+`).FindAllString(text, -1)
	for _, r := range itemsRaw {
		if item, err := strconv.Atoi(r[2:]); err != nil {
			return nil, err
		} else {
			m.items = append(m.items, uint64(item))
		}
	}

	//opperation
	oppRaw := regexp.MustCompile(`[\+\*]\s[\d\w]+`).FindString(text)
	oppRawSplit := strings.Split(oppRaw, " ")
	m.inspSign = rune(oppRaw[0])
	m.inspVal = oppRawSplit[1]

	//target div
	trgDivRaw := regexp.MustCompile(`divisible by \d+`).FindString(text)[13:]
	m.trgDiv, _ = strconv.ParseUint(trgDivRaw, 10, 64)

	//target mon
	trgMonTrueRaw := regexp.MustCompile(`If true: throw to monkey \d+`).FindString(text)[25:]
	trgMonFalseRaw := regexp.MustCompile(`If false: throw to monkey \d+`).FindString(text)[26:]
	m.trgMon[0], _ = strconv.Atoi(trgMonTrueRaw)
	m.trgMon[1], _ = strconv.Atoi(trgMonFalseRaw)

	return &m, nil
}

func MustParseMonkey(text string) Monkey {
	if m, err := ParseMonkey(text); err != nil {
		panic(err)
	} else {
		return m
	}
}

type monkey struct {
	name  string
	items []uint64

	//inspection
	inspSign rune
	inspVal  string

	//target
	trgDiv uint64
	trgMon [2]int
}

// GetTgtDiv implements Monkey
func (m *monkey) getDivider() uint64 {
	return m.trgDiv
}

// DecideTarget implements Monkey
func (m *monkey) GetTarget(item uint64) int {
	if item%m.trgDiv == 0 {
		return m.trgMon[0]
	} else {
		return m.trgMon[1]
	}
}

// GetItems implements Monkey
func (m *monkey) GetItems() []uint64 {
	return append([]uint64{}, m.items...)
}

// Inspect implements Monkey
func (m *monkey) Inspect(item uint64) uint64 {
	var val uint64
	if m.inspVal == "old" {
		val = item
	} else {
		val, _ = strconv.ParseUint(m.inspVal, 10, 64)
	}

	switch m.inspSign {
	case '*':
		val = item * val
	case '+':
		val = item + val
	default:
		panic("Unknown sign")
	}

	if val < item {
		panic("val decreased")
	}

	return val
}

// PopNext implements Monkey
func (m *monkey) PopNext() (uint64, bool) {
	if len(m.items) > 0 {
		item := m.items[0]
		m.items = m.items[1:]
		return item, true
	} else {
		return 0, false
	}
}

// Push implements Monkey
func (m *monkey) Push(item uint64) {
	m.items = append(m.items, item)
}

func (m monkey) String() string {
	template := `
Monkey %s:
  items: %d
  Operation: new = old %c %s
  Test: divisible by %d
    If true: throw to monkey %d
    If false: throw to monkey %d`

	return fmt.Sprintf(template, m.name, m.items, m.inspSign, m.inspVal, m.trgDiv, m.trgMon[0], m.trgMon[1])
}
