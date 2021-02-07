package main


import (
	"sort"
	"strings"
	"fmt"
)

// WalkFn is used when walking the tree. Takes a
// key and value, returning if iteration should
// be terminated.
//在遍历树时使用/ WalkFn。需要一个
//键和值，如果迭代应该返回
//被终止。
type WalkFn func(s string, v interface{}) bool

// leafNode is used to represent a value
type leafNode struct {
	key string
	val interface{}
}

// edge is used to represent an edge node
// edge用于表示一个边缘节点
type edge struct {
	label byte
	node  *node
}

type node struct {
	// leaf is used to store possible leaf// leaf用于存储可能的leaf
	leaf *leafNode

	// prefix is the common prefix we ignore //前缀是我们通常忽略的前缀
	prefix string

	// Edges should be stored in-order for iteration.
	// We avoid a fully materialized slice to save memory,
	// since in most cases we expect to be sparse
	//edges应该按迭代的顺序存储。
	//我们避免一个完全物化的片，以节省内存，
	//因为在大多数情况下，我们希望资源很少
	edges edges
}

func (n *node) isLeaf() bool {
	return n.leaf != nil
}

func (n *node) addEdge(e edge) {
	n.edges = append(n.edges, e)
	n.edges.Sort()
}

func (n *node) updateEdge(label byte, node *node) {
	num := len(n.edges)
	idx := sort.Search(num, func(i int) bool {
		return n.edges[i].label >= label
	})
	if idx < num && n.edges[idx].label == label {
		n.edges[idx].node = node
		return
	}
	panic("replacing missing edge")
}
//根据字节ac吗获取对应的node
func (n *node) getEdge(label byte) *node {
	num := len(n.edges)
	//查找 0-len(edges)书中，一旦找到满足条件的数则返回
	// 他的arc 码大于等于要查找的值
	idx := sort.Search(num, func(i int) bool {
		return n.edges[i].label >= label
	})
	fmt.Println("idx",idx)
	if idx < num && n.edges[idx].label == label {
		return n.edges[idx].node
	}
	return nil
}

func (n *node) delEdge(label byte) {
	num := len(n.edges)
	idx := sort.Search(num, func(i int) bool {
		return n.edges[i].label >= label
	})
	if idx < num && n.edges[idx].label == label {
		copy(n.edges[idx:], n.edges[idx+1:])
		n.edges[len(n.edges)-1] = edge{}
		n.edges = n.edges[:len(n.edges)-1]
	}
}

type edges []edge

func (e edges) Len() int {
	return len(e)
}

func (e edges) Less(i, j int) bool {
	return e[i].label < e[j].label
}

func (e edges) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e edges) Sort() {
	sort.Sort(e)
}

// Tree implements a radix tree. This can be treated as a
// Dictionary abstract data type. The main advantage over
// a standard hash map is prefix-based lookups and
// ordered iteration,
type Tree struct {
	root *node
	size int
}

// New returns an empty Tree
func New() *Tree {
	return NewFromMap(nil)
}

// NewFromMap returns a new tree containing the keys
// from an existing map
func NewFromMap(m map[string]interface{}) *Tree {
	t := &Tree{root: &node{}}
	for k, v := range m {
		t.Insert(k, v)
	}
	return t
}

// Len is used to return the number of elements in the tree
func (t *Tree) Len() int {
	return t.size
}

// longestPrefix finds the length of the shared prefix
// of two strings
func longestPrefix(k1, k2 string) int {
	max := len(k1)
	if l := len(k2); l < max {
		max = l
	}
	var i int
	for i = 0; i < max; i++ {
		if k1[i] != k2[i] {
			break
		}
	}
	return i
}

// Insert is used to add a newentry or update
// an existing entry. Returns if updated.
//插入用于添加新条目或更新
//现有的条目。返回如果更新。
func (t *Tree) Insert(s string, v interface{}) (interface{}, bool) {
	var parent *node
	n := t.root
	search := s
	for {
		// Handle key exhaution 处理 key 耗尽
		if len(search) == 0 {
			//修改值
			if n.isLeaf() { //修改值的情况，是否单独处理更新。是因为树的搜索，比较麻烦，在添加中同时处理跟新比较方便
				old := n.leaf.val
				n.leaf.val = v
				return old, true
			}

			n.leaf = &leafNode{//没有值
				key: s,
				val: v,
			}
			t.size++
			fmt.Println("增加节点",search)
			return nil, false
		}

		// Look for the edge  寻找边缘
		parent = n
		n = n.getEdge(search[0])//得到第一个字节的

		// No edge, create one  没有则创建
		if n == nil {
			e := edge{
				label: search[0],
				node: &node{
					leaf: &leafNode{
						key: s,
						val: v,
					},
					prefix: search,
				},
			}
			parent.addEdge(e)
			t.size++
			return nil, false
		}

		// Determine longest prefix of the search key on match
		//在匹配时确定搜索键的最长前缀
		commonPrefix := longestPrefix(search, n.prefix)
		if commonPrefix == len(n.prefix) {
			search = search[commonPrefix:]
			continue
		}
		//重新拿到公共节点
		// Split the node//分割节点
		t.size++
		//处理第一个节点（原有节点增加层级），用的是更新方式。和
		child := &node{
			prefix: search[:commonPrefix],
		}
		parent.updateEdge(search[0], child)

		// Restore the existing node 恢复现有节点
		child.addEdge(edge{
			label: n.prefix[commonPrefix],
			node:  n,
		})
		n.prefix = n.prefix[commonPrefix:]
		//处理分割后的第二个节点，直接新增，
		// Create a new leaf node //创建一个新的叶子节点
		leaf := &leafNode{
			key: s,
			val: v,
		}

		// If the new key is a subset, add to to this node
		////如果新键是一个子集，添加到这个节点
		search = search[commonPrefix:]
		if len(search) == 0 {//待验证，为何
			//r.Insert("abb",1)
			//r.Insert("acb",2)
			//r.Insert("ac",2)//修改叶子节点最后一条插入会走此逻辑
			child.leaf = leaf
			fmt.Println("插入ac")
			return nil, false
		}

		// Create a new edge for the node 为node 创建一个数枝
		child.addEdge(edge{
			label: search[0],
			node: &node{
				leaf:   leaf,
				prefix: search,
			},
		})
		return nil, false
	}
}

// Delete is used to delete a key, returning the previous
// value and if it was deleted
func (t *Tree) Delete(s string) (interface{}, bool) {
	var parent *node
	var label byte
	n := t.root
	search := s
	for {
		// Check for key exhaution
		if len(search) == 0 {
			if !n.isLeaf() {
				break
			}
			goto DELETE
		}

		// Look for an edge
		parent = n
		label = search[0]
		n = n.getEdge(label)
		if n == nil {
			break
		}

		// Consume the search prefix
		if strings.HasPrefix(search, n.prefix) {
			search = search[len(n.prefix):]
		} else {
			break
		}
	}
	return nil, false

DELETE:
// Delete the leaf
	leaf := n.leaf
	n.leaf = nil
	t.size--

	// Check if we should delete this node from the parent
	if parent != nil && len(n.edges) == 0 {
		parent.delEdge(label)
	}

	// Check if we should merge this node
	if n != t.root && len(n.edges) == 1 {
		n.mergeChild()
	}

	// Check if we should merge the parent's other child
	if parent != nil && parent != t.root && len(parent.edges) == 1 && !parent.isLeaf() {
		parent.mergeChild()
	}

	return leaf.val, true
}

// DeletePrefix is used to delete the subtree under a prefix
// Returns how many nodes were deleted
// Use this to delete large subtrees efficiently
func (t *Tree) DeletePrefix(s string) int {
	return t.deletePrefix(nil, t.root, s)
}

// delete does a recursive deletion
func (t *Tree) deletePrefix(parent, n *node, prefix string) int {
	// Check for key exhaustion
	if len(prefix) == 0 {
		// Remove the leaf node
		subTreeSize := 0
		//recursively walk from all edges of the node to be deleted
		recursiveWalk(n, func(s string, v interface{}) bool {
			subTreeSize++
			return false
		})
		if n.isLeaf() {
			n.leaf = nil
		}
		n.edges = nil // deletes the entire subtree

		// Check if we should merge the parent's other child
		if parent != nil && parent != t.root && len(parent.edges) == 1 && !parent.isLeaf() {
			parent.mergeChild()
		}
		t.size -= subTreeSize
		return subTreeSize
	}

	// Look for an edge
	label := prefix[0]
	child := n.getEdge(label)
	if child == nil || (!strings.HasPrefix(child.prefix, prefix) && !strings.HasPrefix(prefix, child.prefix)) {
		return 0
	}

	// Consume the search prefix
	if len(child.prefix) > len(prefix) {
		prefix = prefix[len(prefix):]
	} else {
		prefix = prefix[len(child.prefix):]
	}
	return t.deletePrefix(n, child, prefix)
}

func (n *node) mergeChild() {
	e := n.edges[0]
	child := e.node
	n.prefix = n.prefix + child.prefix
	n.leaf = child.leaf
	n.edges = child.edges
}

// Get is used to lookup a specific key, returning
// the value and if it was found
func (t *Tree) Get(s string) (interface{}, bool) {
	n := t.root
	search := s
	for {
		// Check for key exhaution
		if len(search) == 0 {
			if n.isLeaf() {
				return n.leaf.val, true
			}
			break
		}

		// Look for an edge
		n = n.getEdge(search[0])
		if n == nil {
			break
		}

		// Consume the search prefix
		if strings.HasPrefix(search, n.prefix) {
			search = search[len(n.prefix):]
		} else {
			break
		}
	}
	return nil, false
}

// LongestPrefix is like Get, but instead of an
// exact match, it will return the longest prefix match.
func (t *Tree) LongestPrefix(s string) (string, interface{}, bool) {
	var last *leafNode
	n := t.root
	search := s
	for {
		// Look for a leaf node
		if n.isLeaf() {
			last = n.leaf
		}

		// Check for key exhaution
		if len(search) == 0 {
			break
		}

		// Look for an edge
		n = n.getEdge(search[0])
		if n == nil {
			break
		}

		// Consume the search prefix
		if strings.HasPrefix(search, n.prefix) {
			search = search[len(n.prefix):]
		} else {
			break
		}
	}
	if last != nil {
		return last.key, last.val, true
	}
	return "", nil, false
}

// Minimum is used to return the minimum value in the tree
func (t *Tree) Minimum() (string, interface{}, bool) {
	n := t.root
	for {
		if n.isLeaf() {
			return n.leaf.key, n.leaf.val, true
		}
		if len(n.edges) > 0 {
			n = n.edges[0].node
		} else {
			break
		}
	}
	return "", nil, false
}

// Maximum is used to return the maximum value in the tree
func (t *Tree) Maximum() (string, interface{}, bool) {
	n := t.root
	for {
		if num := len(n.edges); num > 0 {
			n = n.edges[num-1].node
			continue
		}
		if n.isLeaf() {
			return n.leaf.key, n.leaf.val, true
		}
		break
	}
	return "", nil, false
}

// Walk is used to walk the tree
func (t *Tree) Walk(fn WalkFn) {
	recursiveWalk(t.root, fn)
}

// WalkPrefix is used to walk the tree under a prefix
func (t *Tree) WalkPrefix(prefix string, fn WalkFn) {
	n := t.root
	search := prefix
	for {
		// Check for key exhaution
		if len(search) == 0 {
			recursiveWalk(n, fn)
			return
		}

		// Look for an edge
		n = n.getEdge(search[0])
		if n == nil {
			break
		}

		// Consume the search prefix
		if strings.HasPrefix(search, n.prefix) {
			search = search[len(n.prefix):]

		} else if strings.HasPrefix(n.prefix, search) {
			// Child may be under our search prefix
			recursiveWalk(n, fn)
			return
		} else {
			break
		}
	}

}

// WalkPath is used to walk the tree, but only visiting nodes
// from the root down to a given leaf. Where WalkPrefix walks
// all the entries *under* the given prefix, this walks the
// entries *above* the given prefix.
func (t *Tree) WalkPath(path string, fn WalkFn) {
	n := t.root
	search := path
	for {
		// Visit the leaf values if any
		if n.leaf != nil && fn(n.leaf.key, n.leaf.val) {
			return
		}

		// Check for key exhaution
		if len(search) == 0 {
			return
		}

		// Look for an edge
		n = n.getEdge(search[0])
		if n == nil {
			return
		}

		// Consume the search prefix
		if strings.HasPrefix(search, n.prefix) {
			search = search[len(n.prefix):]
		} else {
			break
		}
	}
}

// recursiveWalk is used to do a pre-order walk of a node
// recursively. Returns true if the walk should be aborted
func recursiveWalk(n *node, fn WalkFn) bool {
	// Visit the leaf values if any
	if n.leaf != nil && fn(n.leaf.key, n.leaf.val) {
		return true
	}

	// Recurse on the children
	for _, e := range n.edges {
		if recursiveWalk(e.node, fn) {
			return true
		}
	}
	return false
}

// ToMap is used to walk the tree and convert it into a map
func (t *Tree) ToMap() map[string]interface{} {
	out := make(map[string]interface{}, t.size)
	t.Walk(func(k string, v interface{}) bool {
		out[k] = v
		return false
	})
	return out
}
func main() {
				//var s string
				//fmt.Printf("选择一个从0到100的整数.\n")
				//answer := sort.Search(100, func(i int) bool {
				//	fmt.Printf("这个数字 <= %d? ", i)
				//	fmt.Scanf("%s", &s)
				//	return s != "" && s[0] == 'y'
				//})
				//fmt.Printf("你的数字是 %d.\n", answer)

	r := New()
	//r.Insert("abb",1)
	//r.Insert("acb",2)
	//fmt.Println(r.root)
	//
	////r.Insert("acd",2)//修改公共节点了
	//r.Insert("ac",2)//修改叶子节点
	//r.Delete("ac")
	//r.Insert("acbc",3)

	r.Insert("abb",1)
	r.Insert("acb",2)
	r.Insert("acd",2)//修改叶子节点
	r.Delete("acb")

	leaf := r.root.leaf
	nodes := r.root.edges[0].node
	edges1 :=  r.root.edges[0].node.edges[0]
	edges2 :=  r.root.edges[0].node.edges[1].node.edges[0]

	fmt.Println(leaf)
	fmt.Println(nodes,edges1,edges2)
	fmt.Println(r.Get("foo"))
}
