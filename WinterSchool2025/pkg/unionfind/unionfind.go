package unionfind

type DisjointSet struct {
	parent []int
	rank   []int
	// или size []int — по желанию
}

func (myDis *DisjointSet) NewDisjoinSet(n int) *DisjointSet {
	parent := make([]int, n)
	rank := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
		rank[0] = 0
	}

	return &DisjointSet{
		parent: parent,
		rank:   rank,
	}

}
