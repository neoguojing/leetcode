你这组题非常适合整理成一张**“滑动窗口快速记忆表”**。核心其实只有两类：

1. **固定窗口**：窗口长度固定，例如 `checkInclusion`
2. **可变窗口**：右边扩张，条件不满足时左边收缩，例如 `lengthOfLongestSubstring`、`characterReplacement`、`minWindow`

另外 `maxSlidingWindow` 是滑动窗口 + 堆，属于一个特殊变体。

# 滑动窗口算法

## 1. 核心模板

### 可变窗口

```go
l := 0

for r := 0; r < len(s); r++ {
    // 加入 s[r]

    for !valid() {
        // 移除 s[l]
        l++
    }

    // 更新答案
}
```

**记忆：**

> `r` 扩张窗口 → 条件失效 → `l` 收缩 → 更新答案

---

### 固定窗口

```go
l := 0

for r := 0; r < len(s); r++ {
    // 加入 s[r]

    if r-l+1 == k {
        // 处理窗口

        // 移除 s[l]
        l++
    }
}
```

**记忆：**

> 窗口达到 `k` → 处理 → 左移一格

---

# 2. 五类典型题

| 问题                                  | 窗口类型 | 核心                 |
| ----------------------------------- | ---- | ------------------ |
| Best Time to Buy/Sell Stock         | 特殊窗口 | 左边维护最低价            |
| Longest Substring Without Repeating | 可变窗口 | 无重复                |
| Character Replacement               | 可变窗口 | `窗口长度 - 最大频次 <= k` |
| Permutation in String               | 固定窗口 | 比较字符频次             |
| Minimum Window Substring            | 可变窗口 | 满足需求后尽量缩小          |
| Sliding Window Maximum              | 固定窗口 | 窗口 + 最大堆           |

---

# 3. Best Time to Buy and Sell Stock

### 思路

不是严格意义上的普通滑动窗口，而是：

> `l` = 当前最低买入价，`r` = 今天卖出

如果：

```text
prices[r] > prices[l]
```

计算利润。

如果：

```text
prices[r] <= prices[l]
```

说明今天更便宜：

```text
l = r
```

### 核心公式

```text
profit = prices[r] - minPrice
```

### 关键代码

```go
l := 0
ret := 0

for r := 1; r < len(prices); r++ {
    if prices[r] > prices[l] {
        ret = max(ret, prices[r]-prices[l])
    } else {
        l = r
    }
}
```

**记忆：**

> 买最低，卖最高；遇到更低价格就换买点。

时间 `O(n)`，空间 `O(1)`。

---

# 4. Longest Substring Without Repeating Characters

### 目标

找最长的：

```text
无重复字符窗口
```

### 窗口条件

```text
window 中不能有重复字符
```

### 思路

`r` 不断扩张。

如果 `s[r]` 已经存在：

```text
不断删除 s[l]
直到 s[r] 不重复
```

### 核心代码

```go
l := 0
set := map[byte]bool{}
res := 0

for r := 0; r < len(s); r++ {

    for set[s[r]] {
        delete(set, s[l])
        l++
    }

    set[s[r]] = true

    res = max(res, r-l+1)
}
```

### 关键记忆

```text
重复 → l++
不重复 → 更新最大长度
```

模板：

```go
for invalid {
    remove(s[l])
    l++
}
```

---

# 5. Longest Repeating Character Replacement

### 目标

最多修改 `k` 个字符，使窗口全部变成同一个字符。

关键：

假设窗口：

```text
A A A B B
```

最多频次：

```text
maxFreq = 3
```

窗口长度：

```text
5
```

需要替换：

```text
5 - 3 = 2
```

所以窗口合法条件：

```text
windowLen - maxFreq <= k
```

### 核心公式

```text
需要修改数量 = 窗口长度 - 窗口内最高频字符数量
```

### 关键代码

```go
count := [26]int{}
l := 0
maxFreq := 0

for r := 0; r < len(s); r++ {
    count[s[r]-'A']++
    maxFreq = max(maxFreq, count[s[r]-'A'])

    for r-l+1-maxFreq > k {
        count[s[l]-'A']--
        l++
    }

    res = max(res, r-l+1)
}
```

### 最重要记忆点

```text
合法：
窗口长度 - 最大频次 <= k

不合法：
窗口长度 - 最大频次 > k
→ l++
```

这是非常典型的**“根据窗口代价收缩”**。

---

# 6. Permutation in String

例如：

```text
s1 = "ab"
s2 = "eidbaooo"
```

寻找：

```text
"ab" 的排列
```

也就是：

```text
"ab"
"ba"
```

### 本质

窗口长度固定：

```text
len(s1)
```

然后比较：

```text
count(window) == count(s1)
```

### 固定窗口模板

```go
count1 := [26]int{}
count2 := [26]int{}

for _, c := range s1 {
    count1[c-'a']++
}

l := 0

for r := 0; r < len(s2); r++ {
    count2[s2[r]-'a']++

    if r-l+1 == len(s1) {

        if count1 == count2 {
            return true
        }

        count2[s2[l]-'a']--
        l++
    }
}
```

### 记忆

> **排列 = 固定长度 + 字符频次相同**

---

# 7. Minimum Window Substring

这是滑动窗口里面最重要的一题。

### 目标

找到 `s` 中包含 `t` 所有字符的**最短窗口**。

例如：

```text
s = ADOBECODEBANC
t = ABC
```

答案：

```text
BANC
```

---

## 核心思想

和前面的题反过来：

前面很多题：

```text
窗口不合法 → 收缩
```

这里：

```text
窗口合法 → 尽可能收缩
```

### 两个状态

```go
have
need
```

例如：

```text
t = "AABC"
```

需求：

```text
A:2
B:1
C:1
```

当窗口满足所有需求：

```text
have == need
```

就开始：

```text
l++
```

尝试寻找更短答案。

### 核心代码

```go
for r := 0; r < len(s); r++ {

    // 加入窗口
    win[s[r]]++

    // 某字符刚好满足需求
    if win[s[r]] == need[s[r]] {
        have++
    }

    // 窗口已经满足要求
    for have == need {

        // 更新最小答案
        res = min(res, r-l+1)

        // 移除左边
        win[s[l]]--

        if win[s[l]] < need[s[l]] {
            have--
        }

        l++
    }
}
```

### 最重要的记忆

```text
不满足 → r 扩张
满足   → l 收缩
```

与最长子串正好形成对照：

| 问题     | 合法后            |
| ------ | -------------- |
| 最长无重复  | 更新答案，继续扩张      |
| 最长字符替换 | 更新答案，继续扩张      |
| 最小覆盖子串 | **不断收缩，寻找最小值** |

---

# 8. Sliding Window Maximum

这个问题和前面不同。

例如：

```text
nums = [1,3,-1,-3,5,3,6,7]
k = 3
```

窗口：

```text
[1,3,-1] → 3
[3,-1,-3] → 3
[-1,-3,5] → 5
...
```

目标：

```text
每个固定窗口求最大值
```

你这里用了：

```text
滑动窗口 + MaxHeap
```

### 核心

堆里面保存：

```go
[value, index]
```

为什么要保存 `index`？

因为需要判断堆顶元素是否已经离开窗口：

```go
(*maxHeap)[0][1] <= i-k
```

离开窗口：

```go
heap.Pop(maxHeap)
```

### 核心代码

```go
for i := 0; i < len(nums); i++ {

    heap.Push(heap, [2]int{nums[i], i})

    if i >= k-1 {

        // 清理过期元素
        for heap[0][1] <= i-k {
            heap.Pop(heap)
        }

        res = append(res, heap[0][0])
    }
}
```

### 记忆

> **固定窗口 + 要求最大/最小值 → 堆可以做。**

复杂度：

```text
O(n log k)
```

---

# 9. 滑动窗口最终记忆框架

真正面试时不要背 6 套代码，只记下面这张图：

```text
                滑动窗口
                    │
          ┌─────────┴─────────┐
          │                   │
       固定窗口              可变窗口
          │                   │
      长度 = k          r 扩张，条件判断
          │                   │
     ┌────┴────┐        ┌─────┴─────┐
     │         │        │           │
  字符频次   最大值    不满足收缩   满足收缩
     │         │        │           │
Permutation   Heap    Longest      Min Window
                       Substring
                       Character
                       Replacement
```

## 最核心的三个模板

### ① 固定窗口

```go
for r := 0; r < n; r++ {
    add(s[r])

    if r-l+1 == k {
        answer()

        remove(s[l])
        l++
    }
}
```

**关键词：**

> `窗口达到 k → 处理 → 左移`

---

### ② 最长窗口

```go
for r := 0; r < n; r++ {
    add(s[r])

    for !valid {
        remove(s[l])
        l++
    }

    res = max(res, r-l+1)
}
```

**关键词：**

> **不合法才收缩，合法就更新最大值**

---

### ③ 最短窗口

```go
for r := 0; r < n; r++ {
    add(s[r])

    for valid {
        res = min(res, r-l+1)

        remove(s[l])
        l++
    }
}
```

**关键词：**

> **合法就收缩，直到不合法**

---

# 10. 面试快速判断

看到题目先问自己三个问题：

### Q1：窗口大小固定吗？

```text
是 → 固定窗口
否 → 可变窗口
```

### Q2：求最长还是最短？

```text
最长：
不合法 → 收缩
合法   → 更新 max

最短：
不合法 → 扩张
合法   → 收缩并更新 min
```

### Q3：窗口状态怎么维护？

常见：

```text
字符是否存在 → map/set
字符频次     → [26]int / map
最大频次     → maxFreq
覆盖需求     → have / need
窗口最大值   → heap / monotonic deque
```

**一句话记忆：**

> **滑动窗口 = `r` 负责扩张，`l` 负责收缩；先确定窗口条件，再决定什么时候移动 `l`，最后决定更新 max 还是 min。**
