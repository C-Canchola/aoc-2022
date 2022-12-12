package main

import (
	"aoc-2022/utils"
	"fmt"
	"strings"
)

type Sln struct {
}

var _ utils.Solution[int, int] = Sln{}

func (Sln) ExampleSolutionPartOne() int {
	return 10605
}
func (Sln) SolvePartOne(lines []string) int {
	monkeys := parseMonkeys(lines)
	for i := 0; i < 20; i++ {
		advanceMonkeys(monkeys, true)
	}

	inspectedCounts := utils.MapFn(monkeys, func(m *monkey) int { return m.inspectedItems })
	utils.SortFromPredicate(inspectedCounts, func(i, j int) bool { return j < i })
	return inspectedCounts[0] * inspectedCounts[1]
}

func (Sln) ExampleSolutionPartTwo() int {
	return 2713310158
}

func (Sln) SolvePartTwo(lines []string) int {
	monkeys := parseMonkeys(lines)
	idx := 0
	fmt.Println(*(monkeys[idx]))
	fmt.Println(getMappedStartingItems(monkeys[0], false))
	for i := 0; i < 200; i++ {
		advanceMonkeys(monkeys, false)
		fmt.Println(*(monkeys[idx]))
		fmt.Println(getMappedStartingItems(monkeys[idx], false))
	}
	fmt.Println(*(monkeys[idx]))
	inspectedCounts := utils.MapFn(monkeys, func(m *monkey) int { return m.inspectedItems })
	utils.SortFromPredicate(inspectedCounts, func(i, j int) bool { return j < i })
	return inspectedCounts[0] * inspectedCounts[1]
}

func main() {

	utils.Solve[int, int](Sln{})

}

type monkey struct {
	startingItems  []int
	opType         string
	opAmt          string
	divTest        int
	testTrueIdx    int
	testFalseIdx   int
	inspectedItems int
}

// fail if not old or parseable as int
func checkOpAmt(s string) {
	if s == "old" {
		return
	}
	_ = utils.ParseIntMust(s)
}

func parseLastWordAsInt(s string) int {
	splt := utils.NewStringSplitter(s, " ")
	return utils.ParseIntMust(splt.GetItem(-1))
}

func parseMonkey(linesPtr *[]string) monkey {
	lines := *linesPtr

	itemLine := lines[1]
	opLine := lines[2]
	testLine := lines[3]
	testTrueLine := lines[4]
	testFalseLine := lines[5]

	itemsSplt := strings.Split(itemLine, ": ")[1]
	nums := strings.Split(itemsSplt, ", ")
	numsParsed := utils.MapFn(nums, utils.ParseIntMust)

	opSplt := strings.Split(opLine, " ")
	opType := opSplt[len(opSplt)-2]
	if !(opType == "+" || opType == "*") {
		panic("opType must be + or *. was " + opType)
	}

	opAmt := opSplt[len(opSplt)-1]

	divTest := parseLastWordAsInt(testLine)
	testTrueIdx := parseLastWordAsInt(testTrueLine)
	testFalseIdx := parseLastWordAsInt(testFalseLine)

	retLines := []string{}
	if len(lines) > 6 {
		retLines = lines[7:]
	}

	*linesPtr = retLines

	return monkey{
		startingItems: numsParsed,
		opType:        opType,
		opAmt:         opAmt,
		divTest:       divTest,
		testTrueIdx:   testTrueIdx,
		testFalseIdx:  testFalseIdx,
	}

}
func getMapper(m *monkey, forDivThree bool) func(int) int {

	if m.opAmt == "old" {
		switch m.opType {
		case "*":
			if forDivThree {
				return func(i int) int { return i * i }
			}
			return func(i int) int { return i }
		case "+":
			if forDivThree {
				return func(i int) int { return i + i }
			}
			return func(i int) int { return multiplyNumByPrimes(i, 2) }
		default:
			panic("invalid opType " + m.opType)
		}
	}

	amt := utils.ParseIntMust(m.opAmt)
	switch m.opType {
	case "*":
		if forDivThree {
			return func(i int) int { return i * amt }
		}
		return func(i int) int { return multiplyNumByPrimes(i, amt) }

	case "+":
		if forDivThree {
			return func(i int) int { return i + amt }
		}
		return func(i int) int { return getProductOfPrimeFactors(i) + amt }
	default:
		panic("invalid opType " + m.opType)
	}
}

func getMapperDiv3(m *monkey) func(int) int {
	return func(i int) int {
		return int(getMapper(m, true)(i) / 3)
	}
}
func getMappedStartingItems(m *monkey, divThree bool) []int {
	if divThree {
		return utils.MapFn(m.startingItems, getMapperDiv3(m))
	}
	nonDivThreeMapper := func(i int) int {
		v := getMapper(m, false)(i)
		return v
	}
	_ = nonDivThreeMapper
	return utils.MapFn(m.startingItems, nonDivThreeMapper)
	return utils.MapFn(m.startingItems, getMapper(m, true))
}

func getMonkeyThrowMap(m *monkey, divThree bool) map[int][]int {
	mapped := getMappedStartingItems(m, divThree)
	mp := map[int][]int{}

	mp[m.testTrueIdx] = []int{}
	mp[m.testFalseIdx] = []int{}

	for _, n := range mapped {
		if n%m.divTest == 0 {
			mp[m.testTrueIdx] = append(mp[m.testTrueIdx], n)
			continue
		}
		mp[m.testFalseIdx] = append(mp[m.testFalseIdx], n)
	}
	return mp
}

func parseMonkeys(lines []string) []*monkey {
	monkeys := []*monkey{}
	for len(lines) != 0 {

		m := parseMonkey(&lines)
		monkeys = append(monkeys, &m)
	}
	return monkeys
}

func advanceMonkeyInRound(m *monkey, monkeys []*monkey, divThree bool) {
	inc := len(m.startingItems)
	throwMap := getMonkeyThrowMap(m, divThree)

	for k, values := range throwMap {
		monkeys[k].startingItems = append(monkeys[k].startingItems, values...)
	}

	m.startingItems = []int{}
	m.inspectedItems = m.inspectedItems + inc
}

func advanceMonkeys(monkeys []*monkey, divThree bool) {
	for _, m := range monkeys {
		advanceMonkeyInRound(m, monkeys, divThree)
	}
}

// Notes:
/*
Believe removing the div three gives incorrect answer due to overflow

Need way to apply mapping which results in the same div check results without strong
increase in numerical values.

Additions should be ok at 10,000 rounds.

One interesting thing of note is that all divisible rules are prime numbers and
all multiply rules (minus square of self) are multiplied by prime numbers.

Rather than keep the numbers, the idea is that only the prime factors should be kept for each number
in the case of multiplication.

Rather, if multiply by prime number, that prime factor gets added to the list of prime factors.
if old * old, keep current list of prime factors

if add, then multiply all prime factors and THEN add the number. Then recalculate prime factors
*/
var primeMap map[int]bool = map[int]bool{}
var onlyPrimeMap map[int]bool = map[int]bool{}

func getNIsPrime(n int) bool {
	isPrime, ok := primeMap[n]
	if ok {
		return isPrime
	}
	primeMap[n] = true
	for k := range onlyPrimeMap {
		if n%k == 0 {
			primeMap[n] = false
			break
		}
	}

	if !primeMap[n] {
		for i := 2; i < n; i++ {
			if n%i == 0 {
				primeMap[n] = false
				break
			}
		}
	}

	if primeMap[n] {
		onlyPrimeMap[n] = true
	}

	return primeMap[n]
}

var primeFactorMap map[int]map[int]bool = map[int]map[int]bool{}

func getPrimeFactors(n int) map[int]bool {
	primes, ok := primeFactorMap[n]
	if ok {
		return primes
	}
	primes = map[int]bool{}
	for k := range onlyPrimeMap {
		if n%k != 0 {
			continue
		}
		otherPrimes := getPrimeFactors(int(n / k))
		copyM := make(map[int]bool)
		copyM[k] = true
		for copyK, copyV := range otherPrimes {
			copyM[copyK] = copyV
		}
		primeFactorMap[n] = copyM
		break
	}

	if _, ok := primeFactorMap[n]; ok {
		return primeFactorMap[n]
	}

	for i := 2; i <= n; i++ {
		if n%i != 0 {
			continue
		}
		if !getNIsPrime(i) {
			continue
		}

		primes[i] = true
	}
	primeFactorMap[n] = primes
	return primeFactorMap[n]
}

func getProductOfPrimeFactors(n int) int {
	factors := getPrimeFactors(n)
	prod := 1
	for k := range factors {
		prod *= k
	}
	return prod
}

func multiplyNumByPrimes(n int, m int) int {
	mult := n * m
	_, ok := primeFactorMap[mult]
	if !ok {
		primeFactorMap[mult] = utils.CombineMaps(getPrimeFactors(n), getPrimeFactors(m))
	}
	return getProductOfPrimeFactors(mult)

}

// returns true even if less. only because all starting numbers are greater than all checks and all ops increase numbers
func isDivisible(n int, m int) bool {
	nPrimes := getPrimeFactors(n)
	mPrimes := getPrimeFactors(m)

	for k := range mPrimes {
		if nPrimes[k] {
			continue
		}
		return false
	}
	return true
}
