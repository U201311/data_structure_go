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
