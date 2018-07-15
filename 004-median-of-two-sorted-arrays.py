from utils.input import input_int_list


def quick_sort(nums, left, right):
    '''inplace sorting'''
    if len(nums) < 0:
        return nums

    i = left
    j = right
    x = nums[(i + j) // 2]

    while i < j:
        while nums[i] < x:
            i += 1

        while nums[j] > x:
            j -= 1

        if i <= j:
            nums[i], nums[j] = nums[j], nums[i]
            i += 1
            j -= 1

    if i < right:
        quick_sort(nums, i, right)
    if j > left:
        quick_sort(nums, left, j)

    return nums


class Solution:
    def findMedianSortedArrays(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """
        nums = nums1 + nums2
        if len(nums) == 0:
            return 0

        quick_sort(nums, 0, len(nums) - 1)
        length = len(nums)

        if length % 2 == 0:
            return (nums[length // 2] + nums[(length // 2) - 1]) / 2
        else:
            return nums[length // 2]


if __name__ == '__main__':
    solver = Solution()
    # Test Case:
    # [1,3]
    # [2]
    list_a = input_int_list('list a: ')
    list_b = input_int_list('list b: ')
    print(solver.findMedianSortedArrays(list_a, list_b))
