下面整理成一份**面试/刷题快速回忆版**，重点放在：**Pattern → 状态 → 选择 → 递归 → 撤销**。

# 回溯算法 Backtracking

## 1. 核心思想

回溯 = **DFS + 撤销选择**

```text
选择
 ↓
递归
 ↓
撤销
```

通用模板：

```go
func dfs(state) {
    if 满足结束条件 {
        保存答案
        return
    }

    for 每个可选项 {
        if 不合法 {
            continue
        }

        做选择
        dfs(下一状态)
        撤销选择
    }
}
```

核心问题只有 4 个：

```text
① 每一层代表什么？
② 当前有哪些选择？
③ 下一层状态是什么？
④ 什么时候结束？
```

---

# 2. 最重要的 Pattern 分类

## ① Subsets：子集

**每个元素：选 / 不选**

```text
nums = [1,2,3]

              []
           /      \
         [1]       []
        /   \     /   \
      [1,2] [1] [2]   []
```

### 状态

```go
dfs(start)
```

### `for` 版本

```go
func dfs(start int) {
    res = append(res, append([]int{}, path...))

    for i := start; i < len(nums); i++ {
        path = append(path, nums[i])
        dfs(i + 1)
        path = path[:len(path)-1]
    }
}
```

**记忆：**

> 子集：当前 path 本身就是答案，所以进入 DFS 就保存。

---

# 3. Combination：组合

例如从 `1,2,3,4` 中选 `2` 个。

特点：

> **无顺序、不能重复使用。**

### Pattern

```go
for i := start; i < n; i++ {
    path = append(path, nums[i])
    dfs(i + 1)
    path = path[:len(path)-1]
}
```

核心：

```text
start
 ↓
只能选择后面的元素
```

避免：

```text
[1,2]
[2,1]  ← 不需要
```

---

# 4. Combination Sum：组合总和

特点：

> **元素可以重复使用。**

例如：

```text
[2,3,6,7], target=7

[2,2,3]
[7]
```

关键区别：

```go
dfs(i, remain-nums[i])
```

而不是：

```go
dfs(i+1, ...)
```

### 记忆

```text
可以重复 → dfs(i)
不能重复 → dfs(i+1)
```

核心：

```go
for i := start; i < len(nums); i++ {
    if nums[i] > remain {
        break
    }

    path = append(path, nums[i])
    dfs(i, remain-nums[i]) // 可以再次使用
    path = path[:len(path)-1]
}
```

---

# 5. Combination Sum II

特点：

1. 每个元素只能使用一次
2. 输入可能有重复值
3. 结果不能重复

所以：

```go
dfs(i + 1, remain-nums[i])
```

并且：

```go
sort.Ints(nums)

if i > start && nums[i] == nums[i-1] {
    continue
}
```

### 最重要的理解

```text
同一层重复 → 跳过
不同层重复 → 可以
```

例如：

```text
[1,1,6]

第一层：
  选第一个 1 ✓
  第二个 1 → 跳过

但：
第一层选 1
    ↓
第二层仍然可以选另一个 1
```

---

# 6. Permutation：排列

排列和 Combination 最大区别：

> **顺序重要，每一层都可以从任意未使用元素中选择。**

因此没有 `start`，使用：

```go
used[i]
```

### Pattern

```go
func dfs() {
    if len(path) == len(nums) {
        save()
        return
    }

    for i := 0; i < len(nums); i++ {
        if used[i] {
            continue
        }

        used[i] = true
        path = append(path, nums[i])

        dfs()

        path = path[:len(path)-1]
        used[i] = false
    }
}
```

### 记忆

```text
组合：
    后面选 → start

排列：
    任意没用过的 → used[]
```

---

# 7. Generate Parentheses

特点：

> 每一步选择 `(` 或 `)`，但必须满足合法性。

状态：

```text
left  = 已使用左括号
right = 已使用右括号
```

选择：

```go
if left < n {
    // 加 "("
}

if right < left {
    // 加 ")"
}
```

结束：

```go
if right == n {
    save()
}
```

### 记忆公式

```text
( ：left < n
) ：right < left
```

本质仍然是：

```text
选择 → DFS → 撤销
```

---

# 8. Letter Combinations

例如：

```text
"23"

2 → abc
3 → def
```

特点：

> **每一层处理一个 digit。**

状态：

```go
dfs(pos)
```

### Pattern

```go
chars := digitToChar[digits[pos]]

for i := 0; i < len(chars); i++ {
    path = append(path, chars[i])
    dfs(pos + 1)
    path = path[:len(path)-1]
}
```

结束：

```go
if pos == len(digits) {
    save()
}
```

### 记忆

> `pos` = 当前处理第几个 digit。

---

# 9. Palindrome Partitioning

例如：

```text
"aab"

["a","a","b"]
["aa","b"]
```

特点：

> 每层选择一个**子串**作为下一段。

状态：

```go
dfs(start)
```

选择：

```go
for i := start + 1; i <= len(s); i++ {
    sub := s[start:i]

    if !isPalindrome(sub) {
        continue
    }

    path = append(path, sub)
    dfs(i)
    path = path[:len(path)-1]
}
```

### 记忆

```text
start = 当前切分位置
s[start:i] = 当前选择的子串
```

---

# 10. Word Search

这是**网格回溯**。

每层：

> 当前选择一个格子。

状态：

```go
dfs(i, j, k)
```

表示：

```text
(i,j) = 当前格子
k     = 当前匹配 word[k]
```

选择：

```text
上
下
左
右
```

### 核心

```go
ch := board[i][j]
board[i][j] = '#'

found := dfs(i+1,j,k+1) ||
         dfs(i-1,j,k+1) ||
         dfs(i,j+1,k+1) ||
         dfs(i,j-1,k+1)

board[i][j] = ch
```

### 记忆

> **网格回溯 = DFS + visited + 恢复**

---

# 11. N-Queens

每层：

> **处理一行。**

状态：

```go
dfs(row)
```

选择：

```text
当前行的每一列
```

合法条件：

```text
同列没有 Queen
左上没有 Queen
右上没有 Queen
```

核心：

```go
for col := 0; col < n; col++ {
    if !isSafe(row, col) {
        continue
    }

    board[row][col] = 'Q'

    dfs(row + 1)

    board[row][col] = '.'
}
```

### 为什么只检查上面？

因为：

```text
row 0 → row 1 → row 2 → ...
```

Queen 是从上往下放的。

下面还没有 Queen。

### 记忆

> **N-Queens = 每层一行，每次选一列。**

---

# 12. 一张表记住所有题

| 题型                       | 每层处理  | 选择方式         | 核心状态              |
| ------------------------ | ----- | ------------ | ----------------- |
| **Subsets**              | 元素    | 选/不选         | `start`           |
| **Combination**          | 元素    | 后面的元素        | `start`           |
| **Combination Sum**      | 元素    | 后面的元素，可重复    | `start`           |
| **Combination Sum II**   | 元素    | 后面的元素，不重复    | `start + 去重`      |
| **Permutation**          | 位置    | 任意未使用元素      | `used[]`          |
| **Parentheses**          | 字符    | `(` / `)`    | `left/right`      |
| **Letter Combinations**  | digit | 对应字符         | `pos`             |
| **Palindrome Partition** | 子串    | `s[start:i]` | `start`           |
| **Word Search**          | 网格    | 上下左右         | `i,j,k + visited` |
| **N-Queens**             | 行     | 当前行的列        | `row`             |

---

# 13. 最重要的判断口诀

遇到回溯题，先问：

### ① 顺序重要吗？

```text
是 → Permutation → used[]

否 → Combination / Subsets → start
```

### ② 元素能重复吗？

```text
能 → dfs(i)

不能 → dfs(i+1)
```

### ③ 有重复输入导致重复答案吗？

```text
有 → sort + 同层去重
```

### ④ 是不是网格？

```text
是 → DFS + visited + 恢复
```

### ⑤ 是不是约束满足？

```text
是 → 每层一个决策 + isValid
```

例如 N-Queens：

```text
每层 = 一行
选择 = 列
约束 = 三个方向
```

---

# 14. 最终只记这个万能模板

```go
func dfs(state) {
    if finished {
        save(path)
        return
    }

    for choice := choices; choice; choice++ {
        if !valid(choice) {
            continue
        }

        // 选择
        path = append(path, choice)

        // 递归
        dfs(nextState)

        // 撤销
        path = path[:len(path)-1]
    }
}
```

然后根据题目替换三个东西：

```text
① state：我现在在哪？
② choices：我下一步能选什么？
③ valid：什么情况下不能选？
```

**回溯算法真正需要掌握的不是 10 道题，而是这 4 个核心 Pattern：**

```text
start      → 子集 / 组合
used[]     → 排列
约束条件   → 括号 / N-Queens
visited    → 网格 DFS
```

其余基本都是这几个模板的变形。
