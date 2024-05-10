type Trie struct {
    Tries [26]*Trie
    End bool
}


func Constructor() Trie {
    return Trie{}
}


func (this *Trie) Insert(word string)  {
    node := this

    for _, c := range word {
        idx := c - 'a'
        if node.Tries[idx] == nil {
            node.Tries[idx] = &Trie{}
        }
        node = node.Tries[idx]
    }
    node.End = true
}


func (this *Trie) Search(word string) bool {
    node := this.SearchNode(word)
    return node != nil && node.End
}


func (this *Trie) StartsWith(prefix string) bool {
    node := this.SearchNode(prefix)
    return node != nil
}

func (this *Trie) SearchNode(word string) *Trie {
    node := this

    for _, c := range word {
        idx := c - 'a'
        if node.Tries[idx] == nil {
            return nil
        }
        node = node.Tries[idx]
    }

    return node
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */