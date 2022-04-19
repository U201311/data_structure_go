package trietree

/*leetcode 208 实现一个Trie树*/
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}

/*leetocde 79 dfs版本*/
func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, i, j, 0, word, &visited) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i, j, index int, word string, visited *[][]bool) bool {
	if index >= len(word) {
		return true
	}
	if i >= 0 && i < len(board) && j >= 0 && j < len(board[0]) && board[i][j] == byte(word[index]) && (*visited)[i][j] == false {
		(*visited)[i][j] = true
		f1 := dfs(board, i, j+1, index+1, word, visited)
		f2 := dfs(board, i, j-1, index+1, word, visited)
		f3 := dfs(board, i+1, j, index+1, word, visited)
		f4 := dfs(board, i-1, j, index+1, word, visited)
		(*visited)[i][j] = false
		return f1 || f2 || f3 || f4
	}
	return false
}

/*leetcode 212 Trie Tree
step1:构建Trie树
step2:递归查找string（dfs）遍历
step3:将找到的string放到set集合中

*/
func findWords(board [][]byte, words []string) []string {
	res := []string{}
	m := len(board)
	n := len(board[0])
	t := new(Trie)
	for _, word := range words {
		t.Insert(word)
	}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			visited[i] = make([]bool, n)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs_word(board, i, j, "", &res, &visited, t)
		}
	}
	return res
}

func dfs_word(board [][]byte, i, j int, str string, res *[]string, visited *[][]bool, t *Trie) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return
	}
	if (*visited)[i][j] {
		return
	}
	str = str + string(str)
	if !t.Search(str) {
		return
	}
	if t.Search(str) && set(res, str) == true {
		*res = append(*res, str)
	}
	(*visited)[i][j] = true
	dfs_word(board, i-1, j, str, res, visited, t)
	dfs_word(board, i+1, j, str, res, visited, t)
	dfs_word(board, i, j-1, str, res, visited, t)
	dfs_word(board, i, j+1, str, res, visited, t)

}

func set(res *[]string, str string) bool {
	for i := 0; i < len(*res); i++ {
		for (*res)[i] == str {
			return false
		}
	}
	return true
}

type Trie struct {
	children [60]*Trie
	word     string
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'A'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.word = word
}

func exist(board [][]byte, word string) bool {

	//方向数组
	dx := [4]int{-1, 0, 0, 1}
	dy := [4]int{0, -1, 1, 0}

	//1.建立tire树 插入进去
	t := &Trie{}
	t.Insert(word)

	m, n := len(board), len(board[0])
	visited := make([][]bool, m)

	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	var ans bool
	//枚举每个节点
	var dfs func(node *Trie, x, y int)
	dfs = func(node *Trie, x, y int) {
		ch := board[x][y]

		node = node.children[ch-'A']
		if node == nil {
			return
		}
		//
		if node.word != "" {
			ans = true
		}

		for i := 0; i < 4; i++ {
			nx, ny := x+dx[i], y+dy[i]

			if nx < 0 || ny < 0 || nx >= m || ny >= n {
				continue
			}
			if visited[nx][ny] == true {
				continue
			}
			//已经访问过
			visited[nx][ny] = true

			dfs(node, nx, ny)
			//还原现场
			visited[nx][ny] = false

		}

	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			visited[i][j] = true
			dfs(t, i, j)
			visited[i][j] = false
		}
	}

	return ans

}
