可以。下面给你一版**完整但紧凑的 1D DP 文档**，把前面遗漏的内容补齐，并统一成「**状态公式 + 解释 + 关键代码 + 记忆点**」格式。

# 1D DP 一维动态规划

## 0. 通用解题框架

做一维 DP，固定问：

```text
① dp[i] 表示什么？
② 当前状态最后一步是什么？
③ 从哪些旧状态转移？
④ 初始化是什么？
⑤ 遍历方向是什么？
⑥ 最终答案在哪里？
```

核心：

> **定义状态 → 找最后一步 → 写状态转移。**

---

# 1. Climbing Stairs

## 状态

```text
dp[i] = 到达第 i 阶的方法数
```

最后一步：

```text
i-1 → i
i-2 → i
```

## 公式

$$
dp[i]=dp[i-1]+dp[i-2]
$$

## 初始化

```text
dp[0] = 1
dp[1] = 1
```

## 代码

```go
dp := make([]int, n+1)

dp[0] = 1
dp[1] = 1

for i := 2; i <= n; i++ {
    dp[i] = dp[i-1] + dp[i-2]
}

return dp[n]
```

## 记忆

> **最后一步有多种方式 → 方案数相加。**

---

# 2. Min Cost Climbing Stairs

## 状态

```text
dp[i] = 到达第 i 个台阶的最小成本
```

最后一步来自：

```text
i-1
i-2
```

## 公式

$$
dp[i]=\min(dp[i-1],dp[i-2])+cost[i]
$$

## 初始化

```text
dp[0] = cost[0]
dp[1] = cost[1]
```

## 代码

```go
dp[0] = cost[0]
dp[1] = cost[1]

for i := 2; i < n; i++ {
    dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
}

return min(dp[n-1], dp[n-2])
```

## 记忆

> **最后一步有多个来源 → 取最小成本。**

---

# 3. House Robber

## 状态

```text
dp[i] = 偷 0..i 号房子的最大收益
```

第 `i` 个房子：

```text
不偷 → dp[i-1]

偷 → dp[i-2] + nums[i]
```

## 公式

$$
dp[i]=\max(dp[i-1],dp[i-2]+nums[i])
$$

## 初始化

```text
dp[0] = nums[0]
dp[1] = max(nums[0], nums[1])
```

## 代码

```go
dp[0] = nums[0]
dp[1] = max(nums[0], nums[1])

for i := 2; i < n; i++ {
    dp[i] = max(dp[i-1], dp[i-2]+nums[i])
}

return dp[n-1]
```

## 记忆

> **最后一个元素：选 or 不选。**

---

# 4. House Robber II

房子变成环：

```text
0 ... n-1
↑       ↓
└───────┘
```

`0` 和 `n-1` 不能同时偷。

拆成两个线性问题：

```text
① 不偷第一个 → [1 ... n-1]
② 不偷最后一个 → [0 ... n-2]
```

## 公式

$$
answer=\max(rob(0,n-2),rob(1,n-1))
$$

## 代码

```go
func rob(nums []int) int {
    n := len(nums)

    if n == 1 {
        return nums[0]
    }

    return max(
        robRange(nums, 0, n-2),
        robRange(nums, 1, n-1),
    )
}

func robRange(nums []int, left, right int) int {
    prev2, prev1 := 0, 0

    for i := left; i <= right; i++ {
        cur := max(prev1, prev2+nums[i])
        prev2 = prev1
        prev1 = cur
    }

    return prev1
}
```

## 记忆

> **环 → 排除一个端点 → 两次 House Robber。**

---

# 5. Decode Ways

## 状态

```text
dp[i] = 前 i 个字符的解码方案数
```

最后一个字符：

```text
单独解码 → dp[i-1]
```

最后两个字符：

```text
一起解码 → dp[i-2]
```

## 公式

如果 `s[i-1] != '0'`：

$$
dp[i] += dp[i-1]
$$

如果最后两位属于 `[10,26]`：

$$
dp[i] += dp[i-2]
$$

所以：

$$
dp[i]
=
合法单字符?dp[i-1]:0
+
合法双字符?dp[i-2]:0
$$

## 初始化

```text
dp[0] = 1
dp[1] = 1
```

`dp[0] = 1` 是为了提供“空前缀”这个基础状态。

## 代码

```go
dp := make([]int, n+1)

dp[0] = 1
dp[1] = 1

for i := 2; i <= n; i++ {
    if s[i-1] != '0' {
        dp[i] += dp[i-1]
    }

    num := int(s[i-2]-'0')*10 +
        int(s[i-1]-'0')

    if num >= 10 && num <= 26 {
        dp[i] += dp[i-2]
    }
}

return dp[n]
```

## 记忆

> **计数 DP → 把所有合法路径 `+=`。**

---

# 6. Word Break

## 状态

推荐定义：

```text
dp[i] = s[0:i] 是否可以被拆分
```

如果：

```text
dp[i] == true
```

并且：

```text
s[i:end] ∈ wordDict
```

那么：

```text
dp[end] = true
```

## 公式

$$
dp[i]\land word(i,end)
\Rightarrow dp[end]
$$

## 初始化

```text
dp[0] = true
```

## 代码

```go
dp := make([]bool, n+1)
dp[0] = true

for i := 0; i < n; i++ {
    if !dp[i] {
        continue
    }

    for _, w := range wordDict {
        end := i + len(w)

        if end <= n && s[i:end] == w {
            dp[end] = true
        }
    }
}

return dp[n]
```

## 记忆

> **字符串看成位置：能从 `i` 走到 `end`，就标记 `dp[end]=true`。**

---

# 7. Longest Increasing Subsequence

这是非常典型的**一维 DP**。

## 状态

```text
dp[i] = 以 nums[i] 结尾的最长递增子序列长度
```

注意：

> 必须是“**以 i 结尾**”，否则无法判断前面的序列能不能接到 `nums[i]`。

枚举：

```text
j < i
```

如果：

```text
nums[j] < nums[i]
```

那么 `nums[i]` 可以接到 `dp[j]` 后面。

## 公式

$$
dp[i]
=
\max_{\substack{j<i\\nums[j]<nums[i]}}
(dp[j]+1)
$$

## 初始化

每个数字自己就是一个长度为 `1` 的递增子序列：

```text
dp[i] = 1
```

## 代码

```go
func lengthOfLIS(nums []int) int {
    n := len(nums)

    dp := make([]int, n)

    for i := 0; i < n; i++ {
        dp[i] = 1

        for j := 0; j < i; j++ {
            if nums[j] < nums[i] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
    }

    ans := 0

    for _, v := range dp {
        ans = max(ans, v)
    }

    return ans
}
```

## 为什么答案是 `max(dp)`？

因为最长递增子序列不一定以最后一个元素结束。

例如：

```text
[9,1,4,2,3,3,7]
```

```text
dp = [1,1,2,2,3,3,4]
```

答案：

```text
max(dp) = 4
```

## 记忆

> **“以 i 结尾的最长……” → 枚举 `j < i` → 能接就 `dp[j]+1`。**

---

# 8. Maximum Product Subarray

## 为什么一个 dp 不够？

因为：

```text
负 × 负 = 正
```

当前最小值可能变成下一步最大值。

例如：

```text
-10 × -5 = 50
```

所以需要同时保存最大、最小。

## 状态

```text
maxDP[i] = 以 nums[i] 结尾的最大乘积
minDP[i] = 以 nums[i] 结尾的最小乘积
```

## 公式

设：

```text
x = nums[i]
```

$$
maxDP[i]
=
\max(x,maxDP[i-1]x,minDP[i-1]x)
$$

$$
minDP[i]
=
\min(x,maxDP[i-1]x,minDP[i-1]x)
$$

## 代码

```go
maxDP[0] = nums[0]
minDP[0] = nums[0]

for i := 1; i < n; i++ {
    x := nums[i]

    maxDP[i] = max(
        x,
        maxDP[i-1]*x,
        minDP[i-1]*x,
    )

    minDP[i] = min(
        x,
        maxDP[i-1]*x,
        minDP[i-1]*x,
    )
}
```

最终：

```text
answer = max(maxDP)
```

## 记忆

> **乘法遇到负数 → 最大、最小一起维护。**

---

# 9. 一维 DP 模式总结

## 模式 1：相邻状态

```text
dp[i] ← dp[i-1], dp[i-2]
```

典型：

```text
Climbing Stairs
Min Cost Climbing Stairs
House Robber
```

---

## 模式 2：以 i 结尾

```text
dp[i] = 以 i 结尾的最优结果
```

然后：

```text
枚举 j < i
```

典型：

```text
LIS
Maximum Product Subarray
```

---

## 模式 3：前缀 DP

```text
dp[i] = 前 i 个元素的结果
```

典型：

```text
Decode Ways
Word Break
```

---

# 10. DP 运算符速记

| 问题    | 运算      |   |   |
| ----- | ------- | - | - |
| 求方案数  | `+=`    |   |   |
| 求最大值  | `max()` |   |   |
| 求最小值  | `min()` |   |   |
| 求是否可行 | `       |   | ` |

例如：

```text
Decode Ways
dp[i] += ...

House Robber
dp[i] = max(...)

Coin Change
dp[i] = min(...)

Word Break
dp[i] = dp[i] || ...
```

---

# 11. 最终速记表

| 题目                  | `dp` 状态                        | 状态转移                             | 答案           |
| ------------------- | ------------------------------ | -------------------------------- | ------------ |
| **Climbing Stairs** | `dp[i]` = 到 i 的方案数             | `dp[i-1]+dp[i-2]`                | `dp[n]`      |
| **Min Cost Stairs** | `dp[i]` = 到 i 的最小成本            | `min(dp[i-1],dp[i-2])+cost[i]`   | `min(last2)` |
| **House Robber**    | `dp[i]` = 0..i 最大收益            | `max(dp[i-1],dp[i-2]+nums[i])`   | `dp[n-1]`    |
| **House Robber II** | 同 House Robber                 | 两个区间分别求                          | `max()`      |
| **Decode Ways**     | `dp[i]` = 前 i 个方案数             | `+= dp[i-1]/dp[i-2]`             | `dp[n]`      |
| **Word Break**      | `dp[i]` = 前缀是否可拆               | `dp[i] → dp[end]`                | `dp[n]`      |
| **LIS**             | `dp[i]` = 以 i 结尾最长长度           | `nums[j]<nums[i] → max(dp[j]+1)` | `max(dp)`    |
| **Max Product**     | `max/minDP[i]` = 以 i 结尾最大/最小乘积 | `max/min(x,max*x,min*x)`         | `max(maxDP)` |

---

# 12. 一维 DP 最终心法

看到一道新题，不要直接想代码：

```text
                    新题
                     ↓
              dp[i] 表示什么？
                     ↓
             当前状态最后一步？
                     ↓
          ┌──────────┼──────────┐
          ↓          ↓          ↓
        前1/2个     前面j个     前缀
          ↓          ↓          ↓
      Climbing      LIS       Decode
      Robber                  Word Break
          ↓
       写公式
          ↓
      初始化 + 遍历
          ↓
       找最终答案
```

最重要的一句话：

> **DP 不是“背公式”，而是先定义 `dp[i]`，然后问：得到 `dp[i]` 的最后一步有哪些可能？**
