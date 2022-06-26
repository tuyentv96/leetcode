from queue import Queue
class Solution:
    def floodFill(self, image: List[List[int]], sr: int, sc: int, color: int) -> List[List[int]]:
        if len(image) == 0:
            return []
        
        if image[sr][sc] == color:
            return image
        
        original = image[sr][sc]

        m = len(image)
        n = len(image[0])
        queue = Queue()
        queue.put((sr, sc))
        visited = set()
        
        while not queue.empty():
            csr, csc = queue.get()
            if (csr, csc) in visited:
                continue
            
            visited.add((csr, csc))
            if csr < 0 or csr >= m or csc < 0 or csc >= n:
                continue
            
            if image[csr][csc] != original:
                continue

            image[csr][csc] = color
            
            queue.put((csr + 1, csc))
            queue.put((csr - 1, csc))
            queue.put((csr, csc + 1))
            queue.put((csr, csc - 1))
            
        return image

            