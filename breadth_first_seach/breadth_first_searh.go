package breadth_first_seach

import (
	"fmt"
	"reflect"
)

type Workable interface {
	Work()
}

type Friendly interface {
	GetFriends() []WorkableFriendly
	Add(f WorkableFriendly)
}

type WorkableFriendly interface {
	Workable
	Friendly
}

type Friend struct {
	friends []WorkableFriendly
}

type Programmer struct {
	Friend
}

type Trader struct {
	Friend
}

type Manager struct {
	Friend
}

type Tester struct {
	Friend
}

func (f *Friend) GetFriends() []WorkableFriendly {
	return f.friends
}

func (f *Friend) Add(friend WorkableFriendly) {
	f.friends = append(f.friends, friend)
}

func (t *Trader) Work() {
	fmt.Println("i am trader")
}

func (p *Programmer) Work() {
	fmt.Println("i am programmer")
}

func (m *Manager) Work() {
	fmt.Println("i am manager")
}

func (t *Tester) Work() {
	fmt.Println("i am tester")
}

func SearchWorker(originWorker WorkableFriendly, searchWorker Workable) (w Workable) {
	var queue = originWorker.GetFriends()
	if len(queue) == 0 {
		return
	}
	for i := 0; i < len(queue); i++ {
		if isSearchable(queue[i], searchWorker) {
			return queue[i]
		} else {
			subFriends := queue[i].GetFriends()
			for j := range subFriends {
				found := false
				for end := i; end != 0; end-- {
					if queue[end] == subFriends[j] {
						found = true
						break
					}
				}
				if found == false {
					queue = append(queue, subFriends[j])
				}
			}
		}
	}
	return nil
}

func isSearchable(worker, searchWorker interface{}) bool {
	return reflect.TypeOf(worker) == reflect.TypeOf(searchWorker)
}
