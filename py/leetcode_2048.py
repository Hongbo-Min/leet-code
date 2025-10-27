"""
如果整数  x 满足：对于每个数位 d ，这个数位 恰好 在 x 中出现 d 次。那么整数 x 就是一个 数值平衡数 。

给你一个整数 n ，请你返回 严格大于 n 的 最小数值平衡数 。
n = 1
输出: 22
n = 1000
输出：1333
"""

class Solution:
    def nextBeautifulNumber(self, n: int) -> int:
        while True:
            n += 1
            if self.is_beautiful_number(n):
                return n
    
    def is_beautiful_number(self, n: int) -> bool:
        """
        判断整数 n 是否是数值平衡数
        """
        digit_count = {}
        for digit in str(n):
            if digit not in digit_count:
                digit_count[digit] = 0
            digit_count[digit] += 1
        for digit, count in digit_count.items():
            if int(digit) != count:
                return False
        return True

s = Solution()
print(s.nextBeautifulNumber(1))
print(s.nextBeautifulNumber(1000))
