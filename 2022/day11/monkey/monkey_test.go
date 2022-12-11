package monkey

import "testing"

var testMonkeyInput0 string = `
Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3
`
var testMonkeyInput1 string = `
Monkey 1:
  Starting items: 79, 98
  Operation: new = old + old
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 3
`

func TestMonkey(t *testing.T) {
	var monkey Monkey

	t.Run("ParseMonkey", func(t *testing.T) {
		m, err := ParseMonkey(testMonkeyInput0)
		if err != nil {
			t.Error(err)
			t.FailNow()
		} else {
			monkey = m
		}
	})

	t.Run("GetItems", func(t *testing.T) {
		expectItems(t, monkey, 79, 98)
	})

	t.Run("PopNext", func(t *testing.T) {
		pop, next := monkey.PopNext()
		expectValue(t, next, true)
		expectValue(t, pop, 79)
		expectItems(t, monkey, 98)

		pop, next = monkey.PopNext()
		expectValue(t, next, true)
		expectValue(t, pop, 98)
		expectItems(t, monkey)

		pop, next = monkey.PopNext()
		expectValue(t, next, false)
		expectValue(t, pop, 0)
		expectItems(t, monkey)
	})

	t.Run("Push", func(t *testing.T) {
		expectItems(t, monkey)
		monkey.Push(79)
		expectItems(t, monkey, 79)
		monkey.Push(98)
		expectItems(t, monkey, 79, 98)
	})

	t.Run("Inpect*19", func(t *testing.T) {
		for inp, exp := range map[uint64]uint64{
			79: 1501,
			98: 1862,
		} {
			out := monkey.Inspect(inp)
			expectValue(t, out, exp)
		}
	})

	t.Run("Inpect+old", func(t *testing.T) {
		monkey, err := ParseMonkey(testMonkeyInput1)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		for inp, exp := range map[uint64]uint64{
			79: 158,
			98: 196,
		} {
			out := monkey.Inspect(inp)
			expectValue(t, out, exp)
		}
	})

	t.Run("GetTarget", func(t *testing.T) {
		for inp, exp := range map[uint64]int{
			79: 3,
			98: 3,
			23: 2,
		} {
			out := monkey.GetTarget(monkey.Inspect(inp))
			expectValue(t, out, exp)
		}
	})

	t.Run("CommonBase", func(t *testing.T) {
		monkeys := []Monkey{
			MustParseMonkey(testMonkeyInput0),
			MustParseMonkey(testMonkeyInput1),
		}
		expectValue(t, CommonBase(monkeys), 19*23)
	})
}

func expectValue[T comparable](t *testing.T, value T, expect T) {
	if value != expect {
		t.Error("Expected", expect, "got", value)
	}
}

func expectItems(t *testing.T, m Monkey, expect ...uint64) {
	items := m.GetItems()
	if len(items) != len(expect) {
		t.Errorf("Expected %d items got %d", len(expect), len(items))
		return
	}
	for in, item := range items {
		expectValue(t, item, expect[in])
	}
}
