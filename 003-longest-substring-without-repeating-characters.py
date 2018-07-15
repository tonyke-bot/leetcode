class Solution:
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        set_hash = set()
        loggest = 0
        current_substr_length = 0
        length = len(s)
        for i in range(length):
            for j in range(i, length):
                if s[j] in set_hash:
                    loggest = max(loggest, current_substr_length)
                    set_hash.clear()
                    current_substr_length = 0
                    break
                set_hash.add(s[j])
                current_substr_length += 1

            if loggest > length - i + 1:
                break

        return max(loggest, current_substr_length)


if __name__ == '__main__':
    solver = Solution()
    # Test Case:
    # "abcabcbb"
    # input_str = input()

    test_str = 'c'
    print(test_str, solver.lengthOfLongestSubstring(test_str))
    test_str = 'abcabcbb'
    print(test_str, solver.lengthOfLongestSubstring(test_str))
    test_str = ''
    print(test_str, solver.lengthOfLongestSubstring(test_str))
    test_str = 'bbbbbb'
    print(test_str, solver.lengthOfLongestSubstring(test_str))
    test_str = 'askdhjklsadhfjkdahsfb'
    print(test_str, solver.lengthOfLongestSubstring(test_str))
    test_str = '1985hkjdsfn'
    print(test_str, solver.lengthOfLongestSubstring(test_str))
    test_str = '123abcdaddd'
    print(test_str, solver.lengthOfLongestSubstring(test_str))
