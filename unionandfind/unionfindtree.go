package unionandfind

import "errors"

/*
leetcode 200 并查集
*/
type UnionFind struct {
	parent []int
	rank   []int //优化并查集,rank低的parent指向rank大的
	count  int
}

//并查集构造函数 初始化并查集
func NewUnionFind(grid [][]byte) *UnionFind {
	m := len(grid)
	n := len(grid[0])
	buf1 := make([]int, m*n)
	buf2 := make([]int, m*n)
	count1 := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				buf1[i*n+j] = i*n + j
				count1 += 1
			} else {
				buf1[i*n+j] = -1
			}
			buf2[i*n+j] = 1
		}
	}
	return &UnionFind{parent: buf1, rank: buf2, count: count1}
}

//并查集寻找根节点值，并加上压缩算法
func (set *UnionFind) getRoot(i int) int {
	if i != set.parent[i] {
		// set.parent[p]=set.parent[set.parent[p]]   //压缩算法1
		// p=set.parent[p]
		set.parent[i] = set.getRoot(set.parent[i]) //压缩算法2递归
	}
	// return p //压缩算法1
	return set.parent[i] //压缩算法2,本题就选压缩算法2了
}

//并查集合并
func (set *UnionFind) Union(p, q int) error {
	if p < 0 || p > len(set.parent) || q < 0 || q > len(set.parent) {
		return errors.New(
			"error: index is illegal.")
	}

	pRoot := set.getRoot(p)
	qRoot := set.getRoot(q)

	if pRoot != qRoot {
		if set.rank[pRoot] < set.rank[qRoot] {
			set.parent[pRoot] = qRoot

		} else if set.rank[qRoot] < set.rank[pRoot] {
			set.parent[qRoot] = pRoot
		} else {
			set.parent[pRoot] = qRoot
			set.rank[qRoot] += 1
		}
		set.count -= 1
	}
	return nil
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	row := len(grid)
	col := len(grid[0])

	uf := NewUnionFind(grid)
	var directions [4][2]int = [4][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} //定义一个方向数组。用于遍历上下左右相邻的字符，用于方法二，使代码更简洁，对比方法一可以看出

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '0' {
				continue
			}
			for _, d := range directions {

				nr, nc := i+d[0], j+d[1]
				if nr >= 0 && nc >= 0 && nr < row && nc < col && grid[nr][nc] == '1' {
					uf.Union(i*col+j, nr*col+nc)
				}
			}
		}
	}
	return uf.count

}
