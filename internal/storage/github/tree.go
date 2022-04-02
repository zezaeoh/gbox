package github

import (
	"github.com/google/go-github/v43/github"
	"github.com/shivamMg/ppds/tree"
	"strings"
)

const slash = "/"

type Node struct {
	data     string
	children []*Node
}

func (n Node) Data() interface{} {
	return n.data
}

func (n Node) Children() []tree.Node {
	c := make([]tree.Node, len(n.children))
	for i, child := range n.children {
		c[i] = tree.Node(child)
	}
	return c
}

func (n *Node) append(node *Node) {
	if n.children == nil {
		n.children = []*Node{node}
		return
	}
	n.children = append(n.children, node)
}

func gitTreeToNode(t *github.Tree) (*Node, error) {
	root := Node{
		data:     slash,
		children: []*Node{},
	}
	ts := map[string]*Node{}
	bs := map[string]*Node{}

	for _, e := range t.Entries {
		switch e.GetType() {
		case "tree":
			ts[e.GetPath()] = &Node{data: toTreeData(e.GetPath())}
		case "blob":
			if strings.HasSuffix(e.GetPath(), gboxSuffix) {
				bs[e.GetPath()] = &Node{data: toBlobData(e.GetPath())}
			}
		}
	}
	for k, v := range ts {
		ss := strings.Split(k, slash)
		// 1 depth
		if len(ss) == 1 {
			root.append(v)
			continue
		}
		// n depth
		p := strings.Join(ss[:len(ss)-1], slash)
		if n, ok := ts[p]; ok {
			n.append(v)
		}
	}
	for k, v := range bs {
		ss := strings.Split(k, slash)
		// 1 depth
		if len(ss) == 1 {
			root.append(v)
			continue
		}
		// n depth
		p := strings.Join(ss[:len(ss)-1], slash)
		if n, ok := ts[p]; ok {
			n.append(v)
		}
	}
	return &root, nil
}

func toTreeData(s string) string {
	ss := strings.Split(s, slash)
	return ss[len(ss)-1] + slash
}

func toBlobData(s string) string {
	ss := strings.Split(s, slash)
	fn := ss[len(ss)-1]
	return fn[:len(fn)-len(gboxSuffix)]
}
