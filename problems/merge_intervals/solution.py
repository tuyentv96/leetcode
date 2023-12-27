class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        if len(intervals) < 2:
            return intervals

        result = []
        intervals.sort(key=lambda interval: interval[0])
        begin, end = intervals[0][0], intervals[0][1]

        for i in range(1, len(intervals)):
            if end < intervals[i][0]:
                result.append([begin, end])
                begin = intervals[i][0]
                end = intervals[i][1]
            else:
                end = max(end, intervals[i][1])
        
        result.append([begin, end])

        return result