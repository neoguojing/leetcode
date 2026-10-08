这是一份为您精心整理、结构化的“公式驱动的 DP 速查文档”**。在原有的核心框架基础上，我们为每一道经典题目**补充了详细的公式解释、状态设计直觉以及核心 Go/Python 风格的代码片段，方便您在面试或刷题时快速查阅与复习。

---

# 🚀 公式驱动的 DP 速查文档

> **DP 解题固定黄金流程**：
> `状态定义` $\rightarrow$ `最后一步（状态转移源）` $\rightarrow$ `状态转移方程` $\rightarrow$ `初始化` $\rightarrow$ `遍历方向`

---

## 1. 0-1 背包（核心基准）

* **状态定义**：`dp[i][j]` = 前 $i$ 个物品中，在容量不超过 $j$ 的情况下，所能获得の**最大价值**。
* **最后一步**：考虑第 $i$ 个物品（重量 $w$，价值 $v$）选或不选：
* **不选**：等价于前 $i-1$ 个物品放入容量 $j$ 的背包，即 `dp[i-1][j]`。
* **选**：前提是 $j \ge w$，价值为前 $i-1$ 个物品放入容量 $j-w$ 的背包价值加上当前物品价值，即 `dp[i-1][j-w] + v`。


* **状态转移方程**：

$$\text{dp}[i][j] = \max(\text{dp}[i-1][j], \text{dp}[i-1][j-w] + v)$$


* **初始化**：`dp[0][j] = 0`（无物品时价值为0），`dp[i][0] = 0`（容量为0时价值为0）。
* **遍历方向**：外层正序遍历物品，内层**倒序**遍历容量（空间优化到一维）。
* **一句话理解**：`0-1 = 上一行 = 倒序`。
* **核心代码片段**：
```go
// 一维空间优化版本
for i := 1; i <= n; i++ {
    w, v := weights[i-1], values[i-1]
    for j := capacity; j >= w; j-- {
        dp[j] = max(dp[j], dp[j-w] + v)
    }
}

```



---

## 2. 0-1 背包：是否能够组成 Target (Subset Sum)

* **状态定义**：`dp[i][j]` = 前 $i$ 个数字中，**是否**能够选出若干个数字，使其和恰好等于 $j$（布尔值）。
* **状态转移方程**：

$$\text{dp}[i][j] = \text{dp}[i-1][j] \parallel \text{dp}[i-1][j - \text{num}]$$


* **初始化**：`dp[0][0] = true`，其余为 `false`。
* **遍历方向**：外层正序遍历数字，内层**倒序**遍历目标和。
* **典型应用**：*Partition Equal Subset Sum*（先判断 `sum % 2 != 0` 直接返回 false，令 $\text{target} = \text{sum} / 2$，转化为 0-1 背包）。
* **核心代码片段**：
```go
dp[0] = true
for _, num := range nums {
    for j := target; j >= num; j-- {
        dp[j] = dp[j] || dp[j-num]
    }
}

```



---

## 3. 0-1 背包：求方案数 (Target Sum)

* **状态定义**：`dp[i][j]` = 前 $i$ 个数字中，能够组合成和为 $j$ 的**方案总数**。
* **状态转移方程**：

$$\text{dp}[i][j] = \text{dp}[i-1][j] + \text{dp}[i-1][j - \text{num}]$$


* **初始化**：`dp[0][0] = 1`（和为0的方案数有一种，即什么都不选）。
* **遍历方向**：外层正序遍历数字，内层**倒序**遍历目标和。
* **典型应用**：*Target Sum*（转换数学关系：设正数和为 $P$，负数和绝对值为 $N$。则 $P - N = \text{target}$，又 $P + N = \text{sum}$，得 $P = (\text{sum} + \text{target}) / 2$，转为求组合数）。
* **核心代码片段**：
```go
dp[0] = 1
for _, num := range nums {
    for j := target; j >= num; j-- {
        dp[j] += dp[j-num]
    }
}

```



---

## 4. 0-1 背包 vs 完全背包（核心对比）

| 特性维度 | 0-1 背包 (0-1 Knapsack) | 完全背包 (Unbounded Knapsack) |
| --- | --- | --- |
| **物品使用次数** | 每件物品**只能用一次** | 每件物品可以**无限次使用** |
| **二维状态转移** | $\text{dp}[i][j] = \max(\text{dp}[i-1][j], \mathbf{\text{dp}[i-1]}[j-w] + v)$ | $\text{dp}[i][j] = \max(\text{dp}[i-1][j], \mathbf{\text{dp}[i]}[j-w] + v)$ |
| **当前状态来源** | 依赖**上一行**（左上方/正上方） | 依赖**当前行**（左侧已更新的值） |
| **一维降维遍历** | `for j := capacity; j >= w; j--` (**倒序**) | `for j := w; j <= capacity; j++` (**正序**) |
| **一句话记忆** | **0-1：上一行，倒序** | **完全：同一行，正序** |

---

## 5. Coin Change II

* **状态定义**：`dp[i][j]` = 使用前 $i$ 种硬币组成金额 $j$ 的**组合数量**。
* **状态转移方程**：

$$\text{dp}[i][j] = \text{dp}[i-1][j] + \text{dp}[i][j - \text{coin}]$$


* **初始化**：`dp[0] = 1`（金额为0时方案数为1）。
* **遍历方向**：外层正序遍历硬币种类，内层**正序**遍历金额（完全背包）。
* **核心代码片段**：
```go
dp[0] = 1
for _, coin := range coins {
    for j := coin; j <= amount; j++ {
        dp[j] += dp[j-coin]
    }
}

```



---

## 6. Edit Distance (编辑距离)

* **状态定义**：`dp[i][j]` = `word1` 的前 $i$ 个字符变成 `word2` 的前 $j$ 个字符所需要的**最少操作数**。
* **状态转移方程**：
* 若 `word1[i-1] == word2[j-1]`：无需操作 $\rightarrow$ $\text{dp}[i][j] = \text{dp}[i-1][j-1]$
* 若不等，取三种操作的最小值：

$$\text{dp}[i][j] = \min(\text{dp}[i-1][j-1], \text{dp}[i-1][j], \text{dp}[i][j-1]) + 1$$




* **初始化**：`dp[0][j] = j`（空串变到长为 $j$ 的字符串需要 $j$ 次插入），`dp[i][0] = i`（删除 $i$ 次）。
* **一句话记忆**：`左上 = 替换`，`上 = 删除`，`左 = 插入`。
* **核心代码片段**：
```go
for i := 1; i <= m; i++ {
    for j := 1; j <= n; j++ {
        if word1[i-1] == word2[j-1] {
            dp[i][j] = dp[i-1][j-1]
        } else {
            dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
        }
    }
}

```



---

## 7. Longest Common Subsequence (最长公共子序列 - LCS)

* **状态定义**：`dp[i][j]` = `text1` 前 $i$ 个字符与 `text2` 前 $j$ 个字符的 **LCS 长度**。
* **状态转移方程**：
* 字符相同时：$\text{dp}[i][j] = \text{dp}[i-1][j-1] + 1$
* 字符不同时：$\text{dp}[i][j] = \max(\text{dp}[i-1][j], \text{dp}[i][j-1])$


* **初始化**：`dp[0][j] = 0`, `dp[i][0] = 0`。
* **一句话记忆**：`相同 → 左上 + 1`，`不同 → 上/左 max`。
* **核心代码片段**：
```go
for i := 1; i <= len1; i++ {
    for j := 1; j <= len2; j++ {
        if text1[i-1] == text2[j-1] {
            dp[i][j] = dp[i-1][j-1] + 1
        } else {
            dp[i][j] = max(dp[i-1][j], dp[i][j-1])
        }
    }
}

```



---

## 8. Unique Paths (不同路径)

* **状态定义**：`dp[i][j]` = 从起点 `(0,0)` 到达网格点 `(i,j)` 的**路径数量**。
* **状态转移方程**：由于只能向右或向下走，到达 `(i,j)` 的上一步只能来自上方 `(i-1, j)` 或左方 `(i, j-1)`。

$$\text{dp}[i][j] = \text{dp}[i-1][j] + \text{dp}[i][j-1]$$


* **初始化**：第一行和第一列的所有格子初始化为 `1`。
* **一句话记忆**：`路径数 = 上 + 左`。
* **核心代码片段**：
```go
for i := 1; i < m; i++ {
    for j := 1; j < n; j++ {
        dp[i][j] = dp[i-1][j] + dp[i][j-1]
    }
}

```



---

## 9. Stock + Cooldown (股票买卖含冷冻期)

* **状态定义**：
* `dp[i][0]` = 第 $i$ 天结束时**不持有**股票的最大利润。
* `dp[i][1]` = 第 $i$ 天结束时**持有**股票的最大利润。


* **状态转移方程**：

$$\text{dp}[i][0] = \max(\text{dp}[i-1][0], \text{dp}[i-1][1] + \text{price})$$


$$\text{dp}[i][1] = \max(\text{dp}[i-1][1], \text{dp}[i-2][0] - \text{price})$$



*(注意：冷冻期导致买入需要参考 $i-2$ 天)*
* **核心直觉**：`买 = -price`，`卖 = +price`，`冷冻期看 i-2`。
* **核心代码片段**：
```go
dp[0][0], dp[0][1] = 0, -prices[0]
for i := 1; i < n; i++ {
    dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
    if i >= 2 {
        dp[i][1] = max(dp[i-1][1], dp[i-2][0] - prices[i])
    } else {
        dp[i][1] = max(dp[i-1][1], -prices[i])
    }
}

```



---

## 10. Longest Increasing Path (矩阵中的最长递增路径)

* **状态定义**：`dp[i][j]` = 从矩阵格子 `(i,j)` 出发能够获得的最长递增路径长度。
* **状态转移方程**：遍历四个方向的邻居 `(ni, nj)`，若 `matrix[ni][nj] > matrix[i][j]`：

$$\text{dp}[i][j] = \max(\text{dp}[i][j], 1 + \text{dp}[ni][nj])$$


* **解法特征**：严格递增保证了不可能形成环，因此使用 **DFS + 记忆化搜索 (Memoization)**。
* **核心代码片段**：
```go
var dfs func(r, c int) int
dfs = func(r, c int) int {
    if memo[r][c] != 0 { return memo[r][c] }
    memo[r][c] = 1 // 至少包含自己
    for _, d := range dirs {
        nr, nc := r + d[0], c + d[1]
        if nr >= 0 && nr < rows && nc >= 0 && nc < cols && matrix[nr][nc] > matrix[r][c] {
            memo[r][c] = max(memo[r][c], 1 + dfs(nr, nc))
        }
    }
    return memo[r][c]
}

```



---

## 11. 🎯 DP 公式核心速查总表

| 经典 Pattern | 状态定义 `dp[...]` | 核心转移公式 |
| --- | --- | --- |
| **0-1 背包 (最大价值)** | `dp[i][j]` | $\max(\text{dp}[i-1][j], \text{dp}[i-1][j-w] + v)$ |
| **0-1 背包 (是否可达)** | `dp[i][j]` | $\text{dp}[i-1][j] \parallel \text{dp}[i-1][j-w]$ |
| **0-1 背包 (方案数)** | `dp[i][j]` | $\text{dp}[i-1][j] + \text{dp}[i-1][j-w]$ |
| **完全背包 (方案数)** | `dp[i][j]` | $\text{dp}[i-1][j] + \text{dp}[i][j-w]$ |
| **LCS (最长公共子序列)** | `dp[i][j]` | 相同：`左上 + 1`；不同：$\max(\text{上}, \text{左})$ |
| **Edit Distance (编辑距离)** | `dp[i][j]` | $\min(\text{左上}, \text{上}, \text{左}) + 1$ |
| **Unique Paths (网格路径)** | `dp[i][j]` | $\text{上} + \text{左}$ |
| **Stock (股票买卖)** | `dp[i][state]` | 基于买入、卖出、冷冻期的状态机转移 |
| **Increasing Path (矩阵递增)** | `dp[i][j]` | $1 + \max(\text{合法更大邻居})$ |

---

## 12. 💡 最终高频考点速记口诀

* **背包优化方向**：
* 0-1 背包：`dp[i-1][j-w]` $\rightarrow$ **倒序遍历**
* 完全背包：`dp[i][j-w]` $\rightarrow$ **正序遍历**


* **字符串对齐**：
* LCS：相同 $\rightarrow$ 左上 + 1；不同 $\rightarrow$ `max(上, 左)`
* 编辑距离：相同 $\rightarrow$ 左上；不同 $\rightarrow$ `min(左上, 上, 左) + 1`


* **聚合操作映射**：
* 能否达到 $\rightarrow$ `OR` ($\parallel$)
* 方案数量 $\rightarrow$ 求和 (`+`)
* 最优解（最大/最小） $\rightarrow$ `max` / `min`
* 矩阵无环搜索 $\rightarrow$ `DFS + Memo`
