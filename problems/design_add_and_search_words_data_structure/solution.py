from collections import deque

class TrieNode:
    def __init__(self):
        self.end = False
        self.childs = [None] * 26

class WordDictionary:

    def __init__(self):
        self.root = TrieNode()

    def charToIndex(self, c):
        return ord(c) - ord('a')

    def addWord(self, word: str) -> None:
        cur = self.root
        for c in word:
            idx = self.charToIndex(c)
            if not cur.childs[idx]:
                cur.childs[idx] = TrieNode()
            cur = cur.childs[idx]
        cur.end = True

    def search(self, word: str) -> bool:
        queue = deque()
        queue.appendleft(self.root)
        for i in range(len(word)):
            c = word[i]
            for j in range(len(queue)):
                node = queue.pop()
                indexes = [i for i in range(26)] if c == '.' else [self.charToIndex(c)]
                for idx in indexes:
                    if node.childs[idx]:
                        queue.appendleft(node.childs[idx])

        for i in range(len(queue)):
            node = queue.pop()
            if node.end:
                return True
        
        return False
        


# Your WordDictionary object will be instantiated and called as such:
# obj = WordDictionary()
# obj.addWord(word)
# param_2 = obj.search(word)