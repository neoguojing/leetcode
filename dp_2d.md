这里为您汇总了整份完整的、包含**二维空间状态定义、状态转移方程、初始化、遍历方向及完整二维代码片段**的“公式驱动的 DP 速查文档”：

---

# 🚀 公式驱动的 DP 速查文档（完整版）

> **DP 解题固定黄金流程**：
> `状态定义` $\rightarrow$ `最后一步（状态转移源）` $\rightarrow$ `状态转移方程` $\rightarrow$ `初始化` $\rightarrow$ `遍历方向`

---

## 1. 0-1 背包（核心基准）

* **状态定义**：`dp[i][j]` = 前 $i$ 个物品中，在容量不超过 $j$ 的情况下，所能获得的**最大价值**。
* **最后一步**：考虑第 $i$ 个物品（重量 $w$，价值 $v$）选或不选：
* **不选**：等价于前 $i-1$ 个物品放入容量 $j$ 的背包，即 `dp[i-1][j]`。
* **选**：前提是 $j \ge w$，价值为前 $i-1$ 个物品放入容量 $j-w$ 的背包价值加上当前物品价值，即 `dp[i-1][j-w] + v`。


* **状态转移方程**：

$$\text{dp}[i][j] = \max(\text{dp}[i-1][j], \text{dp}[i-1][j-w] + v)$$


* **初始化**：`dp[0][j] = 0`（无物品时价值为0），`dp[i][0] = 0`（容量为0时价值为0）。
* **遍历方向**：外层正序遍历物品，内层正序遍历容量。
* **核心代码片段（二维）**：
```go
dp := make([][]int, n+1)
for i := range dp {
    dp[i] = make([]int, capacity+1)
}

for i := 1; i <= n; i++ {
    w, v := weights[i-1], values[i-1]
    for j := 0; j <= capacity; j++ {
        dp[i][j] = dp[i-1][j] // 不选当前物品
        if j >= w {
            dp[i][j] = max(dp[i][j], dp[i-1][j-w] + v) // 选当前物品
        }
    }
}

```



---

## 2. 0-1 背包：是否能够组成 Target (Subset Sum)

* **状态定义**：`dp[i][j]` = 前 $i$ 个数字中，**是否**能够选出若干个数字，使其和恰好等于 $j$（布尔值）。
* **状态转移方程**：

$$\text{dp}[i][j] = \text{dp}[i-1][j] \parallel \text{dp}[i-1][j - \text{num}]$$


* **初始化**：`dp[i][0] = true`（目标和为0总是可以由空集组成），其余为 `false`。
* **遍历方向**：外层正序遍历数字，内层正序遍历目标和。
* **核心代码片段（二维）**：
```go
dp := make([][]bool, n+1)
for i := range dp {
    dp[i] = make([]bool, target+1)
    dp[i][0] = true
}

for i := 1; i <= n; i++ {
    num := nums[i-1]
    for j := 1; j <= target; j++ {
        dp[i][j] = dp[i-1][j] // 不选当前数
        if j >= num {
            dp[i][j] = dp[i][j] || dp[i-1][j-num] // 选当前数
        }
    }
}

```



---

## 3. 0-1 背包：求方案数 (Target Sum)

* **状态定义**：`dp[i][j]` = 前 $i$ 个数字中，能够组合成和为 $j$ 的**方案总数**。
* **状态转移方程**：

$$\text{dp}[i][j] = \text{dp}[i-1][j] + \text{dp}[i-1][j - \text{num}]$$


* **初始化**：`dp[0][0] = 1`。
* **遍历方向**：外层正序遍历数字，内层正序遍历目标和。
* **核心代码片段（二维）**：
```go
dp := make([][]int, n+1)
for i := range dp {
    dp[i] = make([]int, target+1)
}
dp[0][0] = 1

for i := 1; i <= n; i++ {
    num := nums[i-1]
    for j := 0; j <= target; j++ {
        dp[i][j] = dp[i-1][j] // 不选
        if j >= num {
            dp[i][j] += dp[i-1][j-num] // 选
        }
    }
}

```



---

## 4. 0-1 背包 vs 完全背包（核心对比）

| 特性维度 | 0-1 背包 (0-1 Knapsack) | 完全背包 (Unbounded Knapsack) |
| --- | --- | --- |
| **物品使用次数** | 每件物品**只能用一次** | 每件物品可以**无限次使用** |
| **二维状态转移** | $\text{dp}[i][j] = \max(\text{dp}[i-1][j], \mathbf{\text{dp}[i-1]}[j-w] + v)$ | $\text{dp}[i][j] = \max(\text{dp}[i-1][j], \mathbf{\text{dp}[i]}[j-w] + v)$ |
| **当前状态来源** | 依赖**上一行**（上方） | 依赖**当前行**（左侧已更新的值） |
| **一句话记忆** | **0-1：上一行** | **完全：同一行** |

---

## 5. Coin Change II

* **状态定义**：`dp[i][j]` = 使用前 $i$ 种硬币组成金额 $j$ 的**组合数量**。
* **状态转移方程**：

$$\text{dp}[i][j] = \text{dp}[i-1][j] + \text{dp}[i][j - \text{coin}]$$



*(注意：因为硬币无限使用，选当前硬币时，状态从本行 `dp[i]` 转移而来)*
* **初始化**：`dp[i][0] = 1`。
* **核心代码片段（二维）**：
```go
dp := make([][]int, len(coins)+1)
for i := range dp {
    dp[i] = make([]int, amount+1)
    dp[i][0] = 1
}

for i := 1; i <= len(coins); i++ {
    coin := coins[i-1]
    for j := 1; j <= amount; j++ {
        dp[i][j] = dp[i-1][j] // 不用当前硬币
        if j >= coin {
            dp[i][j] += dp[i][j-coin] // 使用当前硬币（来自本行 dp[i]）
        }
    }
}

```



---

## 6. Edit Distance (编辑距离)

* **状态定义**：`dp[i][j]` = `word1` 的前 $i$ 个字符变成 `word2` 的前 $j$ 个字符所需要的**最少操作数**。
* **状态转移方程**：

$$\text{dp}[i][j] = \begin{cases} \text{dp}[i-1][j-1] & \text{if } word1[i-1] == word2[j-1] \\ \min(\text{dp}[i-1][j-1], \text{dp}[i-1][j], \text{dp}[i][j-1]) + 1 & \text{otherwise} \end{cases}$$


* **初始化**：`dp[0][j] = j`, `dp[i][0] = i`。
* **核心代码片段（二维）**：
```go
dp := make([][]int, m+1)
for i := range dp {
    dp[i] = make([]int, n+1)
    dp[i][0] = i
}
for j := 0; j <= n; j++ {
    dp[0][j] = j
}

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

$$\text{dp}[i][j] = \begin{cases} \text{dp}[i-1][j-1] + 1 & \text{if } text1[i-1] == text2[j-1] \\ \max(\text{dp}[i-1][j], \text{dp}[i][j-1]) & \text{otherwise} \end{cases}$$


* **初始化**：`dp[0][j] = 0`, `dp[i][0] = 0`。
* **核心代码片段（二维）**：
```go
dp := make([][]int, len1+1)
for i := range dp {
    dp[i] = make([]int, len2+1)
}

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
* **状态转移方程**：

$$\text{dp}[i][j] = \text{dp}[i-1][j] + \text{dp}[i][j-1]$$


* **初始化**：第一行和第一列全部设为 `1`。
* **核心代码片段（二维）**：
```go
dp := make([][]int, m)
for i := range dp {
    dp[i] = make([]int, n)
    dp[i][0] = 1
}
for j := 0; j < n; j++ {
    dp[0][j] = 1
}

for i := 1; i < m; i++ {
    for j := 1; j < n; j++ {
        dp[i][j] = dp[i-1][j] + dp[i][j-1]
    }
}

```



---

## 9. Stock + Cooldown (股票买卖含冷冻期)

* **状态定义**：
* `dp[i][0]` = 第 $i$ 天结束时不持有股票的最大利润
* `dp[i][1]` = 第 $i$ 天结束时持有股票的最大利润


* **状态转移方程**：

$$\text{dp}[i][0] = \max(\text{dp}[i-1][0], \text{dp}[i-1][1] + \text{price})$$


$$\text{dp}[i][1] = \max(\text{dp}[i-1][1], \text{dp}[i-2][0] - \text{price})$$


* **核心代码片段（二维状态矩阵）**：
```go
dp := make([][]int, n)
for i := range dp {
    dp[i] = make([]int, 2)
}
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
* **状态转移方程**：若邻居 `(ni, nj)` 满足 `matrix[ni][nj] > matrix[i][j]`：

$$\text{dp}[i][j] = 1 + \max(\text{dp}[ni][nj])$$


* **初始化**：`dp[i][j] = 1`（路径至少包含自身）。
* **核心代码片段（二维 DP / 记忆化 DFS）**：
```go
rows, cols := len(matrix), len(matrix[0])
memo := make([][]int, rows)
for i := range memo { memo[i] = make([]int, cols) }

var dfs func(r, c int) int
dfs = func(r, c int) int {
    if memo[r][c] != 0 { return memo[r][c] }

    res := 1
    dirs := [][2]int{{0,1}, {0,-1}, {1,0}, {-1,0}}
    for _, d := range dirs {
        nr, nc := r + d[0], c + d[1]
        if nr >= 0 && nr < rows && nc >= 0 && nc < cols && matrix[nr][nc] > matrix[r][c] {
            res = max(res, 1 + dfs(nr, nc))
        }
    }
    memo[r][c] = res
    return res
}

```



---

## 11. 🎯 DP 公式核心速查总表

| 经典 Pattern | 状态定义 `dp[i][j]` | 核心转移公式 |
| --- | --- | --- |
| **0-1 背包 (最大价值)** | `dp[i][j]` | $\max(\text{dp}[i-1][j], \text{dp}[i-1][j-w] + v)$ |
| **0-1 背包 (是否可达)** | `dp[i][j]` | $\text{dp}[i-1][j] \parallel \text{dp}[i-1][j-w]$ |
| **0-1 背包 (方案数)** | `dp[i][j]` | $\text{dp}[i-1][j] + \text{dp}[i-1][j-w]$ |
| **完全背包 (方案数)** | `dp[i][j]` | $\text{dp}[i-1][j] + \text{dp}[i][j-w]$ |
| **LCS (最长公共子序列)** | `dp[i][j]` | 相同：`左上 + 1`；不同：$\max(\text{上}, \text{左})$ |
| **Edit Distance (编辑距离)** | `dp[i][j]` | $\min(\text{左上}, \text{上}, \text{左}) + 1$ |
| **Unique Paths (网格路径)** | `dp[i][j]` | $\text{上} + \text{左}$ |
| **Stock (股票买卖)** | `dp[i][state]` | 状态机转移（依赖 `i-1` 或 `i-2`） |
| **Increasing Path (矩阵递增)** | `dp[i][j]` | $1 + \max(\text{合法更大邻居})$ |

---

## 12. 💡 最终高频考点速记口诀

* **二维状态依赖**：
* 0-1 背包系列：依赖**上一行**（`i-1`）。
* 完全背包系列：依赖**当前行**（`i`）。


* **字符串对齐**：
* LCS：相同 $\rightarrow$ 左上 + 1；不同 $\rightarrow$ `max(上, 左)`
* 编辑距离：相同 $\rightarrow$ 左上；不同 $\rightarrow$ `min(左上, 上, 左) + 1`


* **聚合操作映射**：
* 能否达到 $\rightarrow$ `OR` ($\parallel$)
* 方案数量 $\rightarrow$ 求和 (`+`)
* 最优解（最大/最小） $\rightarrow$ `max` / `min`
