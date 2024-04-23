package piscine

type QueueNode struct {
	node  *TreeNode
	level int
}

type Queue struct {
	items []*QueueNode
}

func (q *Queue) Enqueue(item *QueueNode) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() *QueueNode {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	queue := Queue{}
	queue.Enqueue(&QueueNode{node: root, level: 0})

	for len(queue.items) > 0 {
		currentNode := queue.Dequeue()
		f(currentNode.node.Data)

		if currentNode.node.Left != nil {
			queue.Enqueue(&QueueNode{node: currentNode.node.Left, level: currentNode.level + 1})
		}
		if currentNode.node.Right != nil {
			queue.Enqueue(&QueueNode{node: currentNode.node.Right, level: currentNode.level + 1})
		}
	}
}
