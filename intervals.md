可以。下面把 **Intervals** 压缩成一份适合刷题前快速回忆的版本，并给每个 Pattern 补上最关键的 Go 代码片段。

# Intervals 区间算法快速记忆

## 1. 核心判断

### 区间重叠

对于：

```text
A = [a,b]
B = [c,d]
```

如果**端点相接不算重叠**：

```text
A 与 B 重叠 ⇔ c < b && a < d
```

常见写法：

```go
cur[0] < prev[1]
```

如果**端点相接也算重叠**：

```go
cur[0] <= prev[1]
```

> ⚠️ 先看题目对 `[1,2]` 和 `[2,3]` 的定义。

---

### 区间合并

重叠后：

```text
start = min(start1, start2)
end   = max(end1, end2)
```

关键代码：

```go
last[1] = max(last[1], cur[1])
```

---

# 2. Merge Intervals

### 题意

给一组区间，合并所有重叠区间。

### 思路

**按 start 排序 → 看前一个 end → 重叠就扩展**

```go
sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][0] < intervals[j][0]
})

res := [][]int{}

for _, cur := range intervals {
    if len(res) == 0 || res[len(res)-1][1] < cur[0] {
        res = append(res, cur)
    } else {
        res[len(res)-1][1] = max(
            res[len(res)-1][1],
            cur[1],
        )
    }
}
```

### 记忆

> **Merge：start 排序，重叠就扩大 end。**

---

# 3. Insert Interval

### 题意

已经排序且不重叠的区间中插入一个新区间。

### 思路

分三段：

```text
左边不重叠 → 中间合并 → 右边不重叠
```

关键判断：

```go
// 左边：cur.end < new.start
for i < len(intervals) &&
    intervals[i][1] < newInterval[0] {

    res = append(res, intervals[i])
    i++
}

// 中间：cur.start <= new.end
for i < len(intervals) &&
    intervals[i][0] <= newInterval[1] {

    newInterval[0] = min(newInterval[0], intervals[i][0])
    newInterval[1] = max(newInterval[1], intervals[i][1])
    i++
}

// 放入合并后的区间
res = append(res, newInterval)

// 右边
res = append(res, intervals[i:]...)
```

### 记忆

> **Insert：左 → 合并 → 右**
> 左边看 `end`，中间看 `start`。

---

# 4. Meeting Rooms

### 题意

判断所有会议能否在一个会议室举行。

### 思路

按开始时间排序，只需要检查相邻会议：

```text
前一个 end <= 当前 start
```

说明没有冲突。

```go
sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][0] < intervals[j][0]
})

for i := 1; i < len(intervals); i++ {
    if intervals[i][0] < intervals[i-1][1] {
        return false
    }
}

return true
```

### 记忆

> **一个房间：start 排序，看相邻。**

---

# 5. Meeting Rooms II

### 题意

求最少需要多少会议室。

### 思路

核心变成：

> **当前会议开始时，有多少个会议还没结束？**

按 `start` 排序。

用 **MinHeap 保存所有房间的 end**：

```text
最小 end = 最早释放的房间
```

如果：

```go
cur.start >= heap[0]
```

就可以复用房间。

关键代码：

```go
sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][0] < intervals[j][0]
})

h := &MinHeap{}
heap.Init(h)

for _, cur := range intervals {
    if h.Len() > 0 && cur[0] >= (*h)[0] {
        heap.Pop(h)
    }

    heap.Push(h, cur[1])
}

return h.Len()
```

MinHeap：

```go
type MinHeap []int

func (h MinHeap) Len() int      { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
    *h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}
```

### 记忆

> **Rooms II：start 进入，最早 end 回收。**

---

# 6. Non-overlapping Intervals

### 题意

最少删除多少区间，使剩余区间互不重叠。

### 关键转换

```text
最少删除
↓
最多保留
↓
每次保留结束最早的区间
```

### 贪心

**按 end 升序排序。**

```go
sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][1] < intervals[j][1]
})

count := 0
prevEnd := math.MinInt

for _, cur := range intervals {
    if cur[0] < prevEnd {
        // 冲突，删除当前
        count++
    } else {
        prevEnd = cur[1]
    }
}

return count
```

### 为什么？

结束越早：

```text
给后面的区间留下的空间越大
```

所以冲突时：

```text
保留 end 更小的
```

### 记忆

> **删除最少 = 保留最多 = end 最早。**

---

# 7. Interval List Intersections

### 题意

求两个已经排序的区间列表的所有交集。

### 交集公式

```text
start = max(A.start, B.start)
end   = min(A.end, B.end)
```

如果：

```text
start <= end
```

存在交集。

### 双指针

```go
start := max(A[i][0], B[j][0])
end := min(A[i][1], B[j][1])

if start <= end {
    res = append(res, []int{start, end})
}

if A[i][1] < B[j][1] {
    i++
} else {
    j++
}
```

### 为什么移动 end 更小的？

因为：

```text
end 小的区间已经结束
```

不可能再与后面的区间产生新的交集。

### 记忆

> **Intersection：max start + min end；谁 end 小，谁移动。**

---

# 8. Minimum Interval to Include Each Query

### 题意

对每个 `q`，找包含它的最短区间：

```text
left <= q <= right
```

长度：

```text
length = right - left + 1
```

### 思路

三个动作：

```text
左边加入
右边淘汰
长度最小
```

按：

```text
interval.left ↑
query ↑
```

排序。

MinHeap 按 `length` 排序。

```go
for _, q := range queries {

    // 1. left <= q 的区间加入
    for i < len(intervals) &&
        intervals[i][0] <= q {

        heap.Push(h, Node{
            length: intervals[i][1] - intervals[i][0] + 1,
            right:  intervals[i][1],
        })

        i++
    }

    // 2. right < q 的区间失效
    for h.Len() > 0 && (*h)[0].right < q {
        heap.Pop(h)
    }

    // 3. 堆顶就是最短合法区间
    if h.Len() > 0 {
        ans[qIndex] = (*h)[0].length
    } else {
        ans[qIndex] = -1
    }
}
```

### 记忆

> **Query：left 加入，right 淘汰，heap 取最短。**

---

# 9. Remove Covered Intervals

### 题意

如果：

```text
A.start <= B.start
A.end   >= B.end
```

则 `A` 覆盖 `B`。

例如：

```text
[1,5] 覆盖 [2,4]
```

### 思路

排序：

```text
start ↑
如果 start 相同：end ↓
```

维护历史最大 `end`：

```go
sort.Slice(intervals, func(i, j int) bool {
    if intervals[i][0] == intervals[j][0] {
        return intervals[i][1] > intervals[j][1]
    }
    return intervals[i][0] < intervals[j][0]
})

maxEnd := 0
count := 0

for _, cur := range intervals {
    if cur[1] <= maxEnd {
        count++ // 被覆盖
    } else {
        maxEnd = cur[1]
    }
}

return len(intervals) - count
```

### 为什么同 start 时 end 要降序？

例如：

```text
[1,5]
[1,3]
```

先处理 `[1,5]`，才能判断 `[1,3]` 被覆盖。

### 记忆

> **Covered：start ↑，同 start 时 end ↓，维护 maxEnd。**

---

# 10. Find Right Interval

### 题意

对每个区间：

```text
[i.start, i.end]
```

寻找另一个区间，使：

```text
other.start >= i.end
```

并且 `other.start` 尽可能小。

### 思路

```text
所有 start 排序
+
lower_bound(end)
```

核心：

```go
pos := sort.Search(len(starts), func(i int) bool {
    return starts[i] >= intervalEnd
})
```

如果：

```go
pos == len(starts)
```

则不存在。

### 记忆

> **Right Interval = 找第一个 start >= 当前 end。**

---

# 11. Intervals 总结表

| 问题               | 排序                 | 核心              |
| ---------------- | ------------------ | --------------- |
| Merge Intervals  | `start ↑`          | 重叠 → 合并         |
| Insert Interval  | `start ↑`          | 左 → 合并 → 右      |
| Meeting Rooms    | `start ↑`          | 相邻检查            |
| Meeting Rooms II | `start ↑`          | MinHeap(end)    |
| Non-overlapping  | `end ↑`            | 贪心保留最早 end      |
| Intersection     | 已排序                | 双指针             |
| Minimum Interval | `left ↑ + query ↑` | MinHeap(length) |
| Remove Covered   | `start ↑, end ↓`   | `maxEnd`        |
| Find Right       | `start ↑`          | lower_bound     |

---

# 12. 最终记忆框架

看到 **Intervals**，先问自己：

```text
① 要合并？
   → start 排序

② 要判断冲突？
   → start 排序

③ 要最少删除？
   → 转成最多保留
   → end 排序 + 贪心

④ 要最少房间？
   → start 排序 + MinHeap(end)

⑤ 求两个区间交集？
   → max(start) + min(end)
   → 双指针

⑥ Query 找最短覆盖区间？
   → left 加入
   → right 淘汰
   → MinHeap(length)

⑦ 找右侧第一个可用区间？
   → start 排序 + lower_bound
```

## 5 个必须记住的代码模式

### ① Merge

```go
if last[1] < cur[0] {
    res = append(res, cur)
} else {
    last[1] = max(last[1], cur[1])
}
```

### ② Insert

```go
for left...
for overlap...
for right...
```

### ③ Greedy Delete

```go
sort by end
if cur.start < prevEnd {
    delete++
} else {
    prevEnd = cur.end
}
```

### ④ Meeting Rooms II

```go
sort by start

if start >= minEnd {
    pop()
}
push(end)
```

### ⑤ Intersection

```go
start = max(a.start, b.start)
end   = min(a.end, b.end)

if start <= end {
    add()
}

move smaller end
```

> **Intervals 本质就 5 个工具：`start排序`、`end排序贪心`、`MinHeap`、`双指针`、`二分`。**
