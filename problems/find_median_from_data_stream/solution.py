from heapq import *

class MedianFinder:

    def __init__(self):
        self.smalls = []
        self.larges = []

    def addNum(self, num: int) -> None:
        if len(self.smalls) == len(self.larges):
            heappush(self.smalls, -num)
            heappush(self.larges, -heappop(self.smalls))
        else:
            heappush(self.larges, num)
            heappush(self.smalls, -heappop(self.larges))

    def findMedian(self) -> float:
        if len(self.smalls) == len(self.larges):
            return (self.larges[0] - self.smalls[0]) / 2
        else:
            return self.larges[0]
        

# Your MedianFinder object will be instantiated and called as such:
# obj = MedianFinder()
# obj.addNum(num)
# param_2 = obj.findMedian()