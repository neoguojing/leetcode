可以，按你之前的算法速记文档风格，压缩成 **“题意 → 模型 → 关键思路 → 关键代码 → 记忆”**。

# Longest / Common 类问题速记

## 1. Maximum Subarray

**题意**：连续子数组的最大和。

**模型**：一维 DP / Kadane

**状态**

```text
cur = 以当前元素结尾的最大子数组和
```

**关键思路**

当前元素只有两种选择：

```text
nums[i]              // 重新开始
cur + nums[i]        // 接着之前的
```

**关键代码**

```go
cur = max(nums[i], cur+nums[i])
ans = max(ans, cur)
```

**记忆**

> 累计变差，就从当前重新开始。

---

## 2. Longest Increasing Subsequence

**题意**：最长递增**子序列**，元素不要求连续。

```text
[9,1,4,2,3,7] → [1,2,3,7]
```

**模型**：一维 DP

**状态**

```text
dp[i] = 以 nums[i] 结尾的最长递增子序列长度
```

**转移**

如果：

```text
j < i && nums[j] < nums[i]
```

那么 `nums[i]` 可以接到 `j` 后面：

```text
dp[i] = max(dp[i], dp[j]+1)
```

**关键代码**

```go
for i := 0; i < n; i++ {
    dp[i] = 1

    for j := 0; j < i; j++ {
        if nums[j] < nums[i] {
            dp[i] = max(dp[i], dp[j]+1)
        }
    }
}

ans = max(dp...)
```

**注意**

答案是：

```text
max(dp)
```

不是 `dp[n-1]`。

**记忆**

> 以 i 结尾 → 向前找 j → 能接就 `dp[j]+1`。

---

## 3. Longest Common Subsequence

**题意**：两个字符串的最长公共**子序列**。

```text
abcde
ace

→ ace
```

**模型**：二维 DP

**状态**

```text
dp[i][j] =
text1 前 i 个
text2 前 j 个
的最长公共子序列长度
```

**关键转移**

相等：

```go
if text1[i-1] == text2[j-1] {
    dp[i][j] = dp[i-1][j-1] + 1
}
```

不相等：

```go
else {
    dp[i][j] = max(
        dp[i-1][j],
        dp[i][j-1],
    )
}
```

**为什么？**

不相等时：

```text
text1 当前字符不要
        OR
text2 当前字符不要
```

**记忆**

> 两个字符串 → 二维 DP
> 相等 → 左上 + 1
> 不等 → 上、左取 max

---

## 4. Longest Substring Without Repeating Characters

**题意**：最长**连续子串**，不能有重复字符。

**模型**：滑动窗口

窗口：

```text
[l ... r]
```

保证窗口内没有重复。

**关键代码**

```go
for r := 0; r < len(s); r++ {
    for set[s[r]] {
        delete(set, s[l])
        l++
    }

    set[s[r]] = true
    ans = max(ans, r-l+1)
}
```

**记忆**

> 连续 + 窗口必须满足条件 → 滑动窗口。

---

## 5. Longest Repeating Character Replacement

**题意**：最长连续子串，最多修改 `k` 个字符后变成相同字符。

**模型**：滑动窗口

关键公式：

```text
窗口长度 - 窗口最高频次数 <= k
```

**关键代码**

```go
count[s[r]]++
maxFreq = max(maxFreq, count[s[r]])

for r-l+1-maxFreq > k {
    count[s[l]]--
    l++
}

ans = max(ans, r-l+1)
```

**记忆**

> 窗口中除了最高频字符，其余都需要替换。

---

## 6. Longest Consecutive Sequence

**题意**：找最长的**数值连续**序列。

```text
[100,4,200,1,3,2]

→ 1,2,3,4
```

**模型**：Hash

注意这里的“连续”不是数组下标连续，而是：

```text
1,2,3,4
```

**关键思路**

只从序列起点开始：

```text
num-1 不存在 → num 是起点
```

然后：

```text
num+1
num+2
...
```

**更容易记的关键代码**

```go
set := map[int]bool{}

for _, x := range nums {
    set[x] = true
}

for _, x := range nums {
    if !set[x-1] { // 起点
        y := x
        for set[y] {
            y++
        }

        ans = max(ans, y-x)
    }
}
```

你的代码使用的是更高级的**区间边界合并**：

```go
left := mp[num-1]
right := mp[num+1]

sum := left + right + 1

mp[num] = sum
mp[num-left] = sum
mp[num+right] = sum
```

面试速记时，优先记：

> Hash 找起点 → 向后找连续数字。

---

## 7. Lowest Common Ancestor

**题意**：二叉树中找 `p、q` 的最近公共祖先。

**模型**：递归 + 后序信息汇总

**关键思路**

递归问左右子树：

```text
左边找到谁？
右边找到谁？
```

**关键代码**

```go
if root == nil || root == p || root == q {
    return root
}

l := lowestCommonAncestor(root.Left, p, q)
r := lowestCommonAncestor(root.Right, p, q)

if l != nil && r != nil {
    return root
}

if l != nil {
    return l
}
return r
```

**记忆**

> 左右都找到 → 当前 root 是 LCA。
> 只有一边找到 → 返回那一边。

---

# 最终区分

```text
Maximum Subarray
→ 连续 + 最大和
→ Kadane

LIS
→ 一个数组 + 不连续 + 递增
→ 一维 DP

LCS
→ 两个字符串 + 不连续 + 公共
→ 二维 DP

Longest Substring
→ 连续 + 无重复
→ 滑动窗口

Character Replacement
→ 连续 + 修改 k 个
→ 滑动窗口

Longest Consecutive
→ 数值连续
→ Hash

LCA
→ 二叉树 + 公共祖先
→ 递归
```

### 最关键的判断

> **先区分 Substring / Subarray / Subsequence：**
>
> * `Substring / Subarray` → **连续**
> * `Subsequence` → **可以跳**
> * 两个 `Sequence` → 通常考虑 **二维 DP**
> * “数值连续” → **Hash**
> * 树上的“公共” → **递归**
