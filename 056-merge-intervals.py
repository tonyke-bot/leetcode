class Interval:
    @classmethod
    def from_list(cls, list):
        return cls(list[0], list[1])

    def __init__(self, s=0, e=0):
        self.start = s
        self.end = e

    def __repr__(self):
        return "{} -> {}".format(self.start, self.end)


class Solution:
    def has_overlap(self, int1, int2):
        return int1.end >= int2.start

    def merge(self, intervals):
        """
        :type intervals: List[Interval]
        :rtype: List[Interval]
        """
        if len(intervals) < 2:
            return intervals

        intervals = sorted(intervals, key=lambda x: x.start)
        current = intervals[0]
        length = len(intervals)
        result = []
        i = 1

        while True:
            if self.has_overlap(current, intervals[i]):
                current_end = max(current.end, intervals[i].end)
                current = Interval(current.start, current_end)
                if i == length - 1:
                    result.append(current)
                    break

            else:
                result.append(current)
                current = intervals[i]

            if i == length - 1:
                result.append(intervals[i])
                break

            i += 1

        return result


if __name__ == '__main__':
    solver = Solution()
    # Test Case:
    #
    print(solver.has_overlap(Interval(1, 4), Interval(1, 3)))
    print(solver.has_overlap(Interval(1, 4), Interval(1, 5)))
    print(solver.has_overlap(Interval(1, 4), Interval(2, 3)))
    print(solver.has_overlap(Interval(1, 4), Interval(2, 5)))

    input_list = [[1, 2], [1, 3]]
    intervals = [Interval(*int) for int in input_list]
    result = solver.merge(intervals)
    print(result)