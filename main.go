package main

import "fmt"

type Unit struct {
	Id       string
	ParentId string
}
type Node struct {
	Unit     Unit
	Children []Node
}
type Queue struct {
	values []Unit
}

var Unitsmap = map[string][]Unit{}

func main() {

	var u1 = Unit{"1", "/"}
	var u2 = Unit{"2", "/"}
	var u3 = Unit{"3", "1"}
	var u4 = Unit{"4", "3"}
	var u5 = Unit{"5", "2"}
	var u6 = Unit{"6", "3"}
	var u7 = Unit{"7", "6"}
	var u8 = Unit{"8", "1"}

	units := []Unit{u1, u2, u3, u4, u5, u6, u7, u8}
	MakeMap(units)
	BFS()

}

func BFS() Node {
	var root = Unit{"/", ""}
	fmt.Println(Unitsmap)
	var node Node
	var queue Queue
	PutToQoueue(root.Id, &queue, -1)
	node.Unit = root
	fmt.Println(queue)

	PutToNodeFromQueue(queue.Peek(), &queue, &node)
	fmt.Println(node)

	return node
}

func MakeMap(units []Unit) {

	for _, uu := range units {
		if HasNo(Unitsmap[uu.ParentId], uu) {
			Unitsmap[uu.ParentId] = append(Unitsmap[uu.ParentId], uu)
		}
		for _, u := range units {
			if uu.Id == u.ParentId && HasNo(Unitsmap[u.ParentId], u) {
				Unitsmap[u.ParentId] = append(Unitsmap[u.ParentId], u)
			}
		}
		//parentId = uu.ParentId
	}

}

func HasNo(units []Unit, u Unit) bool {
	for _, uu := range units {
		if uu == u {
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

func PutToNodeFromQueue(u Unit, queue *Queue, node *Node) {
	iterator := 0
	tempNode := Node{}
	check := true
	for check {
		if node.Unit.Id == queue.Peek().ParentId {
			tempNode.Unit = queue.Pop()
			node.Children = append(node.Children, tempNode)
			//fmt.Println(node)
		} else {
			check = false
		}

	}
	for iterator < len(node.Children) && len(queue.values) > 0 {
		if len(node.Children) > 0 && queue.Peek().ParentId == node.Unit.Id {
			//todo:заменить 124 строку чтоб отправлять не ребенка а пик из queue
			PutToNodeFromQueue(queue.Peek(), queue, &node.Children[iterator])
			iterator++
		} else {
			PutToNodeFromQueue(queue.Peek(), queue, node)

		}
	}

}
