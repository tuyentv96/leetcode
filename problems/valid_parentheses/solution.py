from collections import deque
class Solution:
    def isValid(self, s: str) -> bool:
        dq = deque()
        pairs = {
            '(': ')',
            '{': '}',
            '[': ']'
        }

        for c in s:
            if c in pairs:
                dq.append(c)
            elif len(dq) == 0 or pairs[dq.pop()] != c:
                return False

        return len(dq) == 0