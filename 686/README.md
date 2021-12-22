给定两个字符串 a 和 b，寻找重复叠加字符串 a 的最小次数，使得字符串 b 成为叠加后的字符串 a 的子串，如果不存在则返回 -1。

注意：字符串 "abc" 重复叠加 0 次是 ""，重复叠加 1 次是 "abc"，重复叠加 2 次是 "abcabc"。

 

示例 1：

输入：a = "abcd", b = "cdabcdab"
输出：3
解释：a 重复叠加三遍后为 "abcdabcdabcd", 此时 b 是其子串。
示例 2：

输入：a = "a", b = "aa"
输出：2
示例 3：

输入：a = "a", b = "a"
输出：1
示例 4：

输入：a = "abc", b = "wxyz"
输出：-1
 

提示：

1 <= a.length <= 104
1 <= b.length <= 104
a 和 b 由小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/repeated-string-match
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


从题目来看字符串a只需要扩容到比b长，那么就可以进行子串的匹配，但是不可能无限扩容去匹配，因为扩容到一定长度就会覆盖所有情况。

举个栗子🌰：

- a = "aabaa"
- b = "aaab"

1. aaab 匹配 aaba  ❌
2. aaab 匹配 abaa  ❌

扩容成 aabaaaabaa

此时我们已经匹配过 前面2个a了 
aa baaaabaa

从 b开始匹配 

3. aaab 匹配  baaa ❌
4. aaab 匹配  aaaa ❌
5. aaab 匹配  aaab ✅

其实往后再看下来 最终会回复到 aaba这个重复匹配上 所以

字符串扩容的下界就是 a > b 的扩容数量 

上界就是 下界+1 

