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
	MakeMap(units, "/")
	BFS(units)

}

func BFS(units []Unit) Node {
	var root = Unit{"/", ""}
	fmt.Println(Unitsmap)
	var node Node
	var queue Queue
	PutToQoueue(root.Id, &queue, -1)
	node.Unit = root
	fmt.Println(queue)

	PutToNodeFromQueue(&node, &queue)
	fmt.Println(node)

	return node
}

func MakeMap(units []Unit, parentId string) {
	for _, uu := range units {

		for _, u := range units {
			if parentId == u.ParentId && HasNo(Unitsmap[parentId], u) {
				Unitsmap[parentId] = append(Unitsmap[parentId], u)
			}
		}
		parentId = uu.ParentId
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
		x.ParentId = "nil"
	}

	return x
}

func PutToQoueue(ParentId string, queue *Queue, iterator int) {
	for _, u := range Unitsmap[ParentId] {
		if HasNo(queue.values, u) {
			queue.Push(u)
		}
	}
	iterator++
	if iterator < len(queue.values) {
		PutToQoueue(queue.values[iterator].Id, queue, iterator)
	}
}

func PutToNodeFromQueue(node *Node, queue *Queue) {
	iterator := -1
	tempNode := Node{}
	check := true
	for check {
		if queue.Peek().ParentId == node.Unit.Id {
			tempNode.Unit = queue.Pop()
			node.Children = append(node.Children, tempNode)
			fmt.Println(node)
		} else {
			check = false
		}
	}
	iterator++
	for iterator < len(queue.values) {
		if len(node.Children) > 0 {
			PutToNodeFromQueue(&node.Children[iterator], queue)
		}
	}
}
