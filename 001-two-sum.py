from utils.input import input_int_list, input_int


class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        for i in range(len(nums) - 1):
            for j in range(i + 1, len(nums)):
                if nums[i] + nums[j] == target:
                    return [i, j]

        return []


if __name__ == "__main__":
    solver = Solution()

    nums = input_int_list("nums: ")
    target = input_int("target: ")
    print(solver.twoSum(nums, target))
