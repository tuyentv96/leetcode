class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        
        for c in s:
            if c == '[' or c == '(' or c == '{':
                stack.append(c)
            else:
                if len(stack) == 0:
                    return False
                p = stack.pop()
                if c == '}':
                    if p != '{':
                        return False
                elif c == ')':
                    if p != '(':
                        return False
                elif c == ']':
                    if p != '[':
                        return False
        
        return True if len(stack) == 0 else False
                