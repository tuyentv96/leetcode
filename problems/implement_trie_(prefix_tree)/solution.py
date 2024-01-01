class TrieNode:
    def __init__(self):
        self.childs = [None] * 26
        self.end_word = False

class Trie:

    def __init__(self):
        self.root = TrieNode()

    def char_to_index(self, c):
        return ord(c) - ord('a')
        
    def insert(self, word: str) -> None:
        node = self.root
        for c in word:
            idx = self.char_to_index(c)
            if not node.childs[idx]:
                node.childs[idx] = TrieNode()
            node = node.childs[idx]

        node.end_word = True

    def search(self, word: str) -> bool:
        node = self.root
        for c in word:
            idx = self.char_to_index(c)
            if not node.childs[idx]:
                return False
            node = node.childs[idx]
        return node.end_word == True

    def startsWith(self, prefix: str) -> bool:
        node = self.root
        for c in prefix:
            idx = self.char_to_index(c)
            if not node.childs[idx]:
                return False
            node = node.childs[idx]
        return True     


# Your Trie object will be instantiated and called as such:
# obj = Trie()
# obj.insert(word)
# param_2 = obj.search(word)
# param_3 = obj.startsWith(prefix)