package tries

// SimpleTrie only support `a` to `z`
type SimpleTrie struct {
	end      bool
	children [26]*SimpleTrie
}

func NewSimpleTrie() SimpleTrie {
	return SimpleTrie{}
}

const a = 'a'

func (this *SimpleTrie) Insert(word string) {
	if len(word) == 0 {
		this.end = true
		return
	}
	if this.children[word[0]-a] == nil {
		this.children[word[0]-a] = &SimpleTrie{}
	}
	this.children[word[0]-a].Insert(word[1:])
}

func (this *SimpleTrie) Search(word string) bool {
	if len(word) == 0 {
		return this.end
	}
	if this.children[word[0]-a] == nil {
		return false
	}
	return this.children[word[0]-a].Search(word[1:])
}

func (this *SimpleTrie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	if this.children[prefix[0]-a] == nil {
		return false
	}
	return this.children[prefix[0]-a].StartsWith(prefix[1:])
}
