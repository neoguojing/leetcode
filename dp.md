# 动态规划背包问题

## 1. 0-1背包问题

### 状态方程：
dp[j] = max(dp[j], dp[j - w_i] + v_i)

### 一句话描述：
在背包容量为 `j` 时，当前最大价值是考虑是否放入物品 `i` 后的最大值，即不放物品 `i` 时的最大价值与放入物品 `i` 后的最大价值（即剩余容量 `j - w_i` 的最大价值加上物品 `i` 的价值）之间的较大者。

---

## 2. 完全背包问题

### 状态方程：
dp[j] = max(dp[j], dp[j - w_i] + v_i)

### 一句话描述：
在背包容量为 `j` 时，当前最大价值是考虑放入物品 `i` 多次后的最大值，即不放物品 `i` 时的最大价值与放入物品 `i` 后的最大价值（即剩余容量 `j - w_i` 的最大价值加上物品 `i` 的价值）之间的较大者。

**完全背包与 0-1背包的区别**：完全背包允许每个物品多次放入背包，因此每次更新时都可以考虑再次放入相同物品。

---

## 3. 多重背包问题

### 状态方程（分解法）：
dp[j] = max(dp[j], dp[j - k * w_i] + k * v_i)
其中，`k` 表示物品 `i` 放入的次数，且 `0 ≤ k ≤ m_i`（物品 `i` 的最大数量限制）。

### 一句话描述：
在背包容量为 `j` 时，当前最大价值是考虑放入物品 `i` 多次（最多 `m_i` 次）后的最大值，即不放物品 `i` 时的最大价值与放入 `k` 次物品 `i` 后的最大价值（即剩余容量 `j - k * w_i` 的最大价值加上 `k` 次物品 `i` 的总价值）之间的较大者。

**多重背包与完全背包的不同**：每个物品的放入次数是有限制的，因此我们需要用 `k` 来控制放入次数。

---

## 4. 混合背包问题

### 状态方程：
dp[j] = max(dp[j], dp[j - w_i] + v_i)

### 一句话描述：
在背包容量为 `j` 时，当前最大价值是考虑同时使用 0-1 背包和完全背包方法后的最大值，即不放物品 `i` 时的最大价值与放入物品 `i` 后的最大价值（即剩余容量 `j - w_i` 的最大价值加上物品 `i` 的价值）之间的较大者。

**混合背包**：是一种折中的背包问题，通常我们有两种情况：某些物品只能放一次（类似 0-1 背包），而某些物品可以放多次（类似完全背包）。

---

## 总结

对于所有背包问题的状态方程，我们的核心思路都是：
> 在背包容量为 `j` 时，当前最大价值是考虑是否选择放入物品 `i`，并根据不同的背包类型（0-1、完全、多重或混合）来决定放入次数和剩余容量的最大价值。

---

## 重点区分

- **0-1背包**：每个物品只能放一次。
- **完全背包**：每个物品可以放无限次。
- **多重背包**：每个物品可以放若干次（有限次数）。
- **混合背包**：一部分物品只能放一次，另一部分物品可以放多次。

---

## 关键问题1：为什么使用 `capacity + 1` 的数组？

在动态规划中，`dp` 数组的大小通常是背包的容量加一，这是因为数组 `dp` 的索引代表的是背包的不同容量状态，从 0 到 `capacity`。

### 为什么是 `capacity + 1`？
- **背包容量范围**：背包的容量 `capacity` 是从 0 到 `capacity`，也就是一共有 `capacity + 1` 个可能的容量值。
  - 容量 0 表示背包没有放任何物品。
  - 容量 1 表示背包只能放下物品总重量为 1 的物品。
  - ...
  - 容量 `capacity` 表示背包最大容量，能够放下总重量为 `capacity` 的物品。

因此，我们声明 `dp := make([]int, capacity+1)`，是为了包含所有从 0 到 `capacity` 的容量情况。

---

## 关键问题2：为什么倒序更新避免了重复选择？而正序则重复选择了物品？

**从后往前更新 `dp` 数组** 是为了避免在同一轮更新中多次使用同一个物品。通过从后往前更新，我们确保每个物品只能在一个背包容量下被选择一次。

### 举例说明

假设背包的容量为 5，物品 1 的重量为 2，价值为 3。如果我们 **从前往后** 更新 `dp` 数组，可能会在更新过程中多次选择物品 1，导致计算错误。

例如，当我们在更新 `dp[4]` 时，物品 1 被放入背包一次，而当更新 `dp[5]` 时，物品 1 再次被放入背包，这样物品 1 被“重复”放入，导致错误的结果。

### 为什么倒序避免了重复选择？
从 **后往前** 更新 `dp[j]` 数组，可以确保每个物品只在一个背包容量下被选择一次。通过这样的更新方式，我们避免了在同一轮更新时，物品被多次放入背包，保证了每个物品只被考虑一次。

---

## 总结

- 通过倒序更新 `dp` 数组，我们确保每个物品只能在一个背包容量下被选择一次，避免了重复选择同一物品。
- 使用 `capacity + 1` 数组大小是因为背包容量从 0 到 `capacity`，因此需要 `capacity + 1` 个元素来记录每个容量状态。

