package classify

// UnionFindSet ...
type UnionFindSet struct {
	Set map[interface{}]*UnionFind
}

// MakeSet ...
func MakeSet() *UnionFindSet {
	return &UnionFindSet{
		Set: make(map[interface{}]*UnionFind),
	}
}

// UnionFind ...
// 并查集是一种树型的数据结构，用于处理一些不交集（Disjoint Sets）的合并及查询问题
type UnionFind struct {
	Parent *UnionFind
	Val    int
	Rank   int
}

//NewElem ...
// 创建单元素集合
func (s *UnionFindSet) NewElem(x int) *UnionFind {
	ret := &UnionFind{
		Val:  x,
		Rank: 0,
	}

	ret.Parent = ret
	s.Set[x] = ret

	return ret
}

// Find ...
// 确定元素属于哪一个子集。它可以被用来确定两个元素是否属于同一子集。
func (s *UnionFindSet) Find(x int) *UnionFind {
	elem := s.Set[x]
	if elem == nil {
		return nil
	}

	if elem.Parent.Val != elem.Val {
		// 路径压缩
		elem.Parent = s.Find(elem.Parent.Val)
	}

	return elem.Parent
}

// UnionByRank ...
//将两个子集合并成同一个集合
func (s *UnionFindSet) UnionByRank(x int, y int) *UnionFind {
	xRoot := s.Find(x)
	yRoot := s.Find(y)

	if xRoot.Rank < yRoot.Rank {
		xRoot.Parent = yRoot
	} else if xRoot.Rank > yRoot.Rank {
		yRoot.Parent = xRoot
	} else {
		yRoot.Parent = xRoot
		xRoot.Rank++
	}
	return yRoot
}

// Union ...
// y->x
func (s *UnionFindSet) Union(x int, y int) *UnionFind {
	xRoot := s.Find(x)
	yRoot := s.Find(y)

	yRoot.Parent = xRoot
	return xRoot
}
