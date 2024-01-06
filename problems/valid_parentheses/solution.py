class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        pairs = {
            '{': '}',
            '(': ')',
            '[': ']'
        }

        for c in s:
            if c in pairs:
                stack.append(c)
            elif len(stack) == 0 or pairs[stack.pop()] != c:
                return False
        
        return len(stack) == 0