package main

import "fmt"

type Unit struct {
	Id       string
	ParentId string
	Children []Unit
}

type Queue struct {
	values []Unit
}

var Unitsmap = map[string][]Unit{}

func main() {

	var u1 = Unit{Id: "1", ParentId: "/"}
	var u2 = Unit{Id: "2", ParentId: "/"}
	var u3 = Unit{Id: "3", ParentId: "1"}
	var u4 = Unit{Id: "4", ParentId: "3"}
	var u5 = Unit{Id: "5", ParentId: "2"}
	var u6 = Unit{Id: "6", ParentId: "3"}
	var u7 = Unit{Id: "7", ParentId: "6"}
	var u8 = Unit{Id: "8", ParentId: "1"}

	units := []Unit{u1, u2, u3, u4, u5, u6, u7, u8}
	MakeMap(units)
	fmt.Println(BFS())

}

func BFS() Unit {
	var root = Unit{Id: "/", ParentId: ""}
	fmt.Println(Unitsmap)
	var queue Queue
	PutToQoueue(root.Id, &queue, -1)
	for _ = range queue.values {
		PutToNodeFromQueue(&queue, &root)
	}

	return root
}

func MakeMap(units []Unit) {
	for _, u := range units {
		if HasNo(Unitsmap[u.ParentId], u) {
			Unitsmap[u.ParentId] = append(Unitsmap[u.ParentId], u)
		}
	}
}

func HasNo(units []Unit, u Unit) bool {
	for _, uu := range units {
		if uu.Id == u.Id {
			return false
		}
	}
	return true
}

func (queue *Queue) Push(value Unit) {
	queue.values = append(queue.values, value)
}

func (queue *Queue) Pop() Unit {
	x := queue.values[0]
	queue.values = queue.values[1:len(queue.values)]
	return x
}
func (queue *Queue) Peek() Unit {
	var x Unit
	if len(queue.values) > 0 {
		x = queue.values[0]
	} else {
		x.Id = "nil"
		x.ParentId = "niladasdasd"
	}

	return x
}

func PutToQoueue(ParentId string, queue *Queue, iterator int) {
	for _, u := range Unitsmap[ParentId] {
		queue.Push(u)
	}
	iterator++
	if iterator < len(queue.values) {
		PutToQoueue(queue.values[iterator].Id, queue, iterator)
	}
}

func PutToNodeFromQueue(queue *Queue, unit *Unit) {
	iterator := 0
	check := true
	for check {
		if unit.Id == queue.Peek().ParentId {
			unit.Children = append(unit.Children, queue.Pop())

		} else {
			check = false
		}
	}
	for iterator < len(unit.Children) {
		if len(unit.Children) > 0 {
			PutToNodeFromQueue(queue, &unit.Children[iterator])
			iterator++
		}
	}
}
