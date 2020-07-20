package classify

// UnionFind ...
// 并查集是一种树型的数据结构，用于处理一些不交集（Disjoint Sets）的合并及查询问题
type UnionFind struct {
	Parent *UnionFind
	Val    interface{}
}

//MakeSet ...
// 创建单元素集合
func MakeSet(x interface{}) *UnionFind {
	ret := &UnionFind{
		Val: x,
	}

	ret.Parent = ret
	return ret
}

// Find ...
// 确定元素属于哪一个子集。它可以被用来确定两个元素是否属于同一子集。
func Find(x *UnionFind) *UnionFind {
	if x.Parent.Val != x.Val {
		// 路径压缩
		x.Parent = Find(x.Parent)
	}

	return x.Parent
}

// Union ...
//将两个子集合并成同一个集合
func Union(x *UnionFind, y *UnionFind) *UnionFind {
	xRoot := Find(x)
	yRoot := Find(y)

	xRoot.Parent = yRoot
	return yRoot
}
