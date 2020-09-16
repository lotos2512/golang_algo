package breadth_first_seach

import (
	"github.com/stfsy/golang-assert"
	"testing"
)

var programmer1 = Programmer{}
var programmer2 = Programmer{}
var programmer3 = Programmer{}
var programmer4 = Programmer{}

var trader1 = Trader{}
var trader2 = Trader{}
var manager1 = Manager{}
var manager2 = Manager{}

func init() {
	programmer1.Add(&programmer2)
	programmer1.Add(&programmer3)
	programmer1.Add(&programmer4)

	programmer2.Add(&programmer1)
	programmer2.Add(&manager1)
	programmer2.Add(&manager2)

	trader1.Add(&trader2)
	trader2.Add(&trader1)

	manager1.Add(&manager2)
	manager2.Add(&manager1)

	manager2.Add(&trader1)
	manager2.Add(&trader2)
}

func TestSearchWorkerTrader(t *testing.T) {
	result := SearchWorker(&programmer1, &Trader{})
	assert.Equal(t, result, &trader1, "result")
}

func TestSearchWorkerManager(t *testing.T) {
	result := SearchWorker(&programmer1, &Manager{})
	assert.Equal(t, result, &manager1, "result")
}

func TestSearchWorkerNotFound(t *testing.T) {
	result := SearchWorker(&programmer1, &Tester{})
	assert.Equal(t, result, nil, "result")
}

func TestSearchWorkerTested(t *testing.T) {
	manager := Manager{}
	tester := Tester{}
	manager.Add(&tester)
	programmer1.Add(&manager)
	result := SearchWorker(&programmer1, &Tester{})
	assert.Equal(t, result, &tester, "result")
}
