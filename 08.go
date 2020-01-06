package main

type treeNode2 struct {
	val    string
	parent *treeNode2
	left   *treeNode2
	right  *treeNode2
}

func findNext(node *treeNode2) *treeNode2 {

	if node == nil {
		return nil
	}
	if node.right != nil {
		// 有右节点，则找到右子树的最左节点
		t := node.right
		for t.left != nil {
			t = t.left
		}
		return t
	}

	if node.parent != nil {
		// 非根节点
		t := node
		p := node.parent

		for p != nil && t == p.right {
			t = p
			p = p.parent
		}
		return p
	}

	return nil
}

/*
func main() {
	a := &treeNode2{val: "a"}
	b := &treeNode2{val: "b"}
	c := &treeNode2{val: "c"}
	a.left = b
	a.right = c
	b.parent = a
	c.parent = a
	d := &treeNode2{val: "d"}
	e := &treeNode2{val: "e"}
	b.left = d
	b.right = e
	d.parent = b
	e.parent = b
	f := &treeNode2{val: "f"}
	g := &treeNode2{val: "g"}
	c.left = f
	c.right = g
	f.parent = c
	g.parent = c
	h := &treeNode2{val: "h"}
	i := &treeNode2{val: "i"}
	h.parent = e
	i.parent = e
	e.left = h
	e.right = i

	res := findNext(a)
	fmt.Printf("node: %v, res: %v\n", a.val, res.val)

	res = findNext(b)
	fmt.Printf("node: %v, res: %v\n", b.val, res.val)

	res = findNext(d)
	fmt.Printf("node: %v, res: %v\n", d.val, res.val)

	res = findNext(f)
	fmt.Printf("node: %v, res: %v\n", f.val, res.val)

	res = findNext(i)
	fmt.Printf("node: %v, res: %v\n", i.val, res.val)
}
*/
