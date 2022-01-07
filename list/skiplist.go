package list

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// CopyRandomList 对存在随机指针的列表进行深度copy，重新构建一个新的列表
// no 138
func CopyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	new := &Node{}
	cur := head
	p2I := map[*Node]*Node{}
	p := new
	for cur != nil {
		tmp := &Node{
			Val: cur.Val,
		}
		p.Next = tmp
		p = p.Next
		p2I[cur] = tmp
		cur = cur.Next
	}

	p = new.Next
	for head != nil {
		if head.Random != nil {
			rand := p2I[head.Random]
			p.Random = rand
		}
		head = head.Next
		p = p.Next
	}

	return new.Next
}
