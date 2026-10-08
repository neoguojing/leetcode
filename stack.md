基于你这几道题，可以把**栈算法压缩成 3 类核心 Pattern**：普通栈、单调栈、辅助栈。重点记“**什么时候入栈、什么时候出栈、栈里维护什么**”。

# 栈算法快速记忆

## 1. 栈的核心思想

**LIFO：后进先出**

Go：

```go
s := []int{}

s = append(s, x)          // push
x := s[len(s)-1]          // top
s = s[:len(s)-1]          // pop
```

栈题首先问：

> **当前元素来了，它需要和“最近的哪个元素”比较？**

如果答案是“最近一个满足条件的元素”，通常考虑栈。

---

# 一、普通栈：处理嵌套 / 顺序依赖

## 1. Valid Parentheses

**题意：** 判断括号是否匹配。

### 思路

遇到左括号 → 入栈
遇到右括号 → 必须匹配栈顶，然后弹出。

```go
closeToOpen := map[rune]rune{
    ')': '(',
    ']': '[',
    '}': '{',
}

for _, c := range s {
    if open, ok := closeToOpen[c]; ok {
        if len(stack) == 0 || stack[len(stack)-1] != open {
            return false
        }
        stack = stack[:len(stack)-1]
    } else {
        stack = append(stack, c)
    }
}

return len(stack) == 0
```

### 记忆

> **左括号入栈，右括号匹配栈顶。**

---

# 二、辅助栈：用额外状态维护信息

## 2. Min Stack

**题意：** `Push / Pop / Top / GetMin` 都要求 O(1)。

普通栈：

```text
stack = [5, 3, 7]
```

如果想 O(1) 得到最小值，就维护一个 `minStack`：

```text
stack    = [5, 3, 7]
minStack = [5, 3, 3]
```

每一层保存：

> **当前栈中最小值**

```go
func (s *MinStack) Push(val int) {
    s.stack = append(s.stack, val)

    minv := val
    if len(s.min) > 0 && s.min[len(s.min)-1] < minv {
        minv = s.min[len(s.min)-1]
    }

    s.min = append(s.min, minv)
}
```

### 记忆

> **主栈存数据，辅助栈存状态。**

类似思想还可以用于：

* Min Stack
* Max Stack
* 需要 O(1) 查询某种聚合状态的问题

---

# 三、栈模拟计算：Reverse Polish Notation

## Evaluate Reverse Polish Notation

**题意：**

```text
["2","1","+","3","*"]
```

相当于：

```text
(2 + 1) * 3
```

### 思路

数字 → 入栈
运算符 → 弹出两个数字计算，再把结果入栈。

关键：

```go
a := s[len(s)-1] // 右操作数
b := s[len(s)-2] // 左操作数

s = s[:len(s)-2]

switch t {
case "+":
    s = append(s, b+a)
case "-":
    s = append(s, b-a)
case "*":
    s = append(s, b*a)
case "/":
    s = append(s, b/a)
}
```

### 最容易错

```text
- 和 /
```

顺序一定是：

```text
b op a
```

而不是：

```text
a op b
```

### 记忆

> **数字 push，运算符 pop 两个 → 计算 → push。**

---

# 四、单调栈：栈题最重要的 Pattern

单调栈解决的核心问题：

> **寻找左/右边第一个比当前元素大/小的元素。**

典型：

* Daily Temperatures
* Next Greater Element
* Largest Rectangle in Histogram
* Stock Span 等

---

# 4. Daily Temperatures

**题意：**

找到右边第一个比当前温度高的日期。

例如：

```text
73 74 75 71 69 72 76
↓
 1  1  4  2  1  1  0
```

### 核心

栈里存的是：

```text
还没有找到答案的元素下标
```

当前温度 `t` 来了：

```go
for len(stack) > 0 &&
    temperatures[stack[len(stack)-1]] < t {

    idx := stack[len(stack)-1]
    stack = stack[:len(stack)-1]

    ret[idx] = i - idx
}

stack = append(stack, i)
```

### 为什么出栈？

因为：

```text
当前 t > 栈顶元素
```

说明当前元素就是栈顶元素的：

> **第一个更大元素**

所以栈顶可以结算。

### 记忆

> **找右边第一个更大 → 单调递减栈。**

---

# 五、Car Fleet

这题本质也是：

> **单调栈 / 单调序列思想。**

先按照距离终点的远近排序：

```text
离终点近 → 离终点远
```

计算每辆车到终点时间：

```text
time = (target - position) / speed
```

维护到目前为止的“最大到达时间”。

```go
for _, car := range cars {
    time := float64(car.distance) / float64(car.speed)

    if len(stack) == 0 || time > stack[len(stack)-1] {
        stack = append(stack, time)
    }
}
```

如果：

```text
当前 time <= 前车 time
```

说明当前车最终会追上前车：

```text
→ 同一个 Fleet
```

### 记忆

> **按位置排序 → 算到达时间 → 时间追上就合并。**

---

# 六、Largest Rectangle：单调栈高级应用

**题意：**

柱状图中寻找最大矩形面积。

核心不是直接算面积，而是找：

> **当前柱子左右第一个比它矮的位置。**

例如：

```text
      ┌───┐
  ┌───┤   │
  │   │   │
──┴───┴───┴──
```

对于每个 `i`：

```text
left[i]  = 左边第一个 < heights[i] 的位置
right[i] = 右边第一个 < heights[i] 的位置
```

那么：

```text
宽度 = right[i] - left[i] - 1

面积 = heights[i] * 宽度
```

### 你的代码采用的是

**边界跳跃 DP 思想**：

```go
for i := 1; i < n; i++ {
    p := i - 1

    for p >= 0 && heights[p] >= heights[i] {
        p = left[p]
    }

    left[i] = p
}
```

右边同理。

### 更重要的理解

对于柱子 `i`：

```text
      i
      ↓
... [矮] [高] [高] [i] [高] [矮] ...
      ↑                    ↑
    left                 right
```

矩形可以向左右扩展，直到遇到：

```text
比 heights[i] 矮的柱子
```

所以：

```text
area = height × 最大宽度
```

### 单调栈版本的核心

```go
for len(stack) > 0 &&
    heights[stack[len(stack)-1]] >= heights[i] {

    mid := stack[len(stack)-1]
    stack = stack[:len(stack)-1]

    left := -1
    if len(stack) > 0 {
        left = stack[len(stack)-1]
    }

    right := i

    area := heights[mid] * (right - left - 1)
}
stack = append(stack, i)
```

### 记忆

> **最大矩形 = 找每根柱子的左右第一个更矮位置。**

---

# 七、单调栈统一模板

遇到：

> **“右边第一个更大/更小”**

优先想到：

```go
stack := []int{}

for i := 0; i < n; i++ {

    for len(stack) > 0 && 条件成立 {

        idx := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        // i 是 idx 的答案
    }

    stack = append(stack, i)
}
```

关键在于：

### 找右边第一个更大

```go
nums[i] > nums[stackTop]
```

通常维护：

```text
单调递减栈
```

### 找右边第一个更小

```go
nums[i] < nums[stackTop]
```

通常维护：

```text
单调递增栈
```

---

# 八、栈题 Pattern 总表

| 类型    | 典型题                | 核心              |
| ----- | ------------------ | --------------- |
| 普通栈   | Valid Parentheses  | 最近匹配            |
| 辅助栈   | Min Stack          | 栈 + 状态          |
| 栈模拟   | Evaluate RPN       | pop → 计算 → push |
| 单调递减栈 | Daily Temperatures | 右边第一个更大         |
| 单调序列  | Car Fleet          | 时间追赶/合并         |
| 单调栈   | Largest Rectangle  | 左右第一个更矮         |

---

# 最后只记这 4 句话

### ① 括号

> **左括号 push，右括号匹配 top。**

### ② 表达式

> **数字 push，运算符 pop 两个再 push 结果。**

### ③ 单调栈

> **找“下一个更大/更小” → 单调栈。**

### ④ 最大矩形

> **每根柱子找左右第一个更矮 → 高 × 宽。**

**栈题最核心的判断：**

> 如果当前元素需要和“之前最近的、还没解决的元素”比较，优先考虑栈；如果要求“第一个更大/更小”，优先考虑单调栈。

如果按你现在刷题的方式，**栈这一章不需要再刷很多题**：把 `Valid Parentheses → RPN → MinStack → Daily Temperatures → Car Fleet → Largest Rectangle` 这 6 个 Pattern 吃透，基本就覆盖了面试里最常见的栈题。
