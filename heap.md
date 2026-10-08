下面整理成一份**面试快速回忆版 Heap 文档**。重点只保留：**什么时候用、核心思路、公式、关键代码、容易错的点**；Heap 本身的 Go 实现省略。

# Heap（堆）算法快速记忆

## 1. Heap 是解决什么问题？

Heap 的核心不是排序，而是：

> **动态维护当前最重要的元素。**

通常是：

```text
最大值 → Max Heap
最小值 → Min Heap
```

典型场景：

```text
反复取最大/最小
Top K
动态中位数
多个有序数据流合并
```

---

# 2. Heap Pattern 总表

| 场景          | Heap         | 核心思路           |
| ----------- | ------------ | -------------- |
| 反复取最大       | Max Heap     | 最大值始终在堆顶       |
| 反复取最小       | Min Heap     | 最小值始终在堆顶       |
| Top K 最大    | **Min Heap** | 只保留 K 个最大，淘汰最小 |
| Top K 最小    | **Max Heap** | 只保留 K 个最小，淘汰最大 |
| 动态中位数       | Max + Min    | 左半最大、右半最小      |
| K-way Merge | Min/Max Heap | 每个序列维护一个当前候选   |

---

# 3. Pattern ①：反复取最大 / 最小

## 典型题：Last Stone Weight

### 题意

每次：

```text
取当前最大的两个 x >= y
    ↓
x == y → 都删除
x > y  → Push(x-y)
```

### 思路

```text
反复取最大两个
      ↓
   Max Heap
```

### 关键代码

```go
for h.Len() > 1 {
    x := heap.Pop(h).(int)
    y := heap.Pop(h).(int)

    if x > y {
        heap.Push(h, x-y)
    }
}

if h.Len() == 1 {
    return (*h)[0]
}
return 0
```

### 记忆

> **反复取最大 → Max Heap；处理后的结果重新入堆。**

---

# 4. Pattern ②：Top K

这是 Heap 最重要的 Pattern 之一。

## Kth Largest

例如：

```text
nums = [3,2,1,5,6,4]
k = 2
```

答案：

```text
5
```

---

## 为什么找第 K 大却使用 Min Heap？

因为我们只需要保留：

```text
最大的 K 个
```

使用：

```text
Min Heap
```

让堆顶始终是：

> **当前 Top-K 中最小的元素。**

一旦超过 K：

```text
Push
 ↓
size > K
 ↓
Pop 最小
```

### 关键代码

```go
minHeap := &MinHeap{}

for _, num := range nums {
    heap.Push(minHeap, num)

    if minHeap.Len() > k {
        heap.Pop(minHeap)
    }
}

return (*minHeap)[0]
```

最终：

```text
Min Heap 中只剩最大的 K 个

例如 K=3：

[8, 10, 15]

       8
      / \
    10   15
```

堆顶 `8`：

```text
= Top-K 中最小
= 第 K 大
```

---

## Top K 口诀

```text
找 K 个最大
    ↓
Min Heap
    ↓
超过 K
    ↓
Pop 最小
```

反过来：

```text
找 K 个最小
    ↓
Max Heap
    ↓
超过 K
    ↓
Pop 最大
```

### 一定记住

> **Top K 最大 → 小顶堆。**
> **Top K 最小 → 大顶堆。**

---

# 5. Pattern ③：K Closest Points

### 题意

找距离原点最近的 K 个点。

距离：

```math
d=x^2+y^2
```

### 思路

我们需要：

```text
反复取距离最小的点
        ↓
    Min Heap
```

Heap 中可以保存：

```text
[dist, x, y]
```

按照 `dist` 比较。

### 关键代码

```go
for _, p := range points {
    x, y := p[0], p[1]
    dist := x*x + y*y

    heap.Push(minHeap, []int{dist, x, y})
}

res := [][]int{}

for i := 0; i < k; i++ {
    p := heap.Pop(minHeap).([]int)
    res = append(res, []int{p[1], p[2]})
}
```

### 为什么不用 `sqrt`？

因为：

```text
sqrt(a) < sqrt(b)
⇔
a < b
```

所以直接比较：

```text
x² + y²
```

即可。

### 记忆

> **K Closest = 按距离取最小 → Min Heap。**

---

# 6. Pattern ④：动态中位数

## MedianFinder

这是 Heap 中最重要的“双堆”问题。

把数据分成两半：

```text
1 2 3 | 5 8 10
←────→|←──────→
 small    large
```

其中：

```text
small = Max Heap
large = Min Heap
```

---

## 为什么？

### small 用 Max Heap

需要快速获得：

```text
左半边最大值
```

即：

```go
(*small)[0]
```

### large 用 Min Heap

需要快速获得：

```text
右半边最小值
```

即：

```go
(*large)[0]
```

所以：

```text
small[0] | large[0]
    ↑          ↑
左边最大    右边最小
```

这两个数就是中间位置。

---

## 两个不变量

### ① 数量平衡

```text
|small.Len() - large.Len()| <= 1
```

### ② 左边不大于右边

```text
max(small) <= min(large)
```

即：

```text
small[0] <= large[0]
```

---

## AddNum

### 第一步：决定放哪边

```go
if large.Len() == 0 || num <= (*large)[0] {
    heap.Push(small, num)
} else {
    heap.Push(large, num)
}
```

### 第二步：重新平衡

```go
if small.Len() > large.Len()+1 {
    heap.Push(large, heap.Pop(small))
}

if large.Len() > small.Len()+1 {
    heap.Push(small, heap.Pop(large))
}
```

---

## FindMedian

```go
if small.Len() > large.Len() {
    return float64((*small)[0])
}

if large.Len() > small.Len() {
    return float64((*large)[0])
}

return float64((*small)[0]+(*large)[0]) / 2
```

### 记忆

> **左大顶，右小顶；数量平衡，中间取值。**

---

# 7. Pattern ⑤：K-way Merge

## 典型题：Design Twitter

Twitter 的 News Feed 本质上不是普通的 Heap，而是：

> **多个已经有序的数据流进行合并。**

例如：

```text
用户 A：A1 → A2 → A3
用户 B：B1 → B2
用户 C：C1 → C2 → C3
```

每个用户自己的 Tweet 已经按时间有序。

现在需要得到：

```text
A1 B1 C1 A2 C2 B2 ...
```

即全局最新的 10 条。

---

## 核心思路

不要把所有 Tweet 都放进 Heap。

只放：

```text
每个用户当前最新的一条
```

例如：

```text
A1
B1
C1
 ↓
 Heap
```

Pop 出全局最新：

```text
A1
```

那么 A 的下一条：

```text
A2
```

进入 Heap：

```text
B1
C1
A2
```

继续。

---

## 核心模板

```text
每个有序序列
      ↓
取当前第一个元素
      ↓
Push Heap
      ↓
Pop 全局最小/最大
      ↓
从对应序列取下一个
      ↓
Push
      ↓
继续
```

这就是：

> **K-way Merge**

---

## Twitter 中 Heap 元素

你的设计：

```go
[count, tweetId, followeeId, index]
```

分别表示：

```text
count      → 时间顺序
tweetId    → Tweet
followeeId  → 来自哪个用户
index      → 这个用户下一条 Tweet
```

Pop：

```go
curr := heap.Pop(h).([]int)

count := curr[0]
tweetId := curr[1]
followeeId := curr[2]
index := curr[3]
```

输出 Tweet：

```go
res = append(res, tweetId)
```

然后从同一个用户取下一条：

```go
if index >= 0 {
    tweet := tweets[index]

    heap.Push(h, []int{
        tweet[0],
        tweet[1],
        followeeId,
        index - 1,
    })
}
```

### 记忆

> **多路有序数据 → 每路放一个候选 → Heap 选全局最优 → 从该路补下一个。**

---

# 8. Task Scheduler：特殊情况

这题容易被误认为一定需要 Heap。

实际上题目只有：

```text
A-Z
```

所以直接统计频率即可。

设：

```text
maxFreq  = 最大任务频率
countMax = 达到最大频率的任务数
```

公式：

```math
(maxFreq-1)(n+1)+countMax
```

最终：

```go
return max(
    len(tasks),
    (maxFreq-1)*(n+1)+countMax,
)
```

### 思路

最高频任务：

```text
A _ _ A _ _ A
```

它决定框架。

其他任务：

```text
填空隙
```

填不满：

```text
idle
```

填得满：

```text
直接连续执行
```

所以：

```text
answer = max(任务总数, 理论框架长度)
```

### 记忆

> **Task Scheduler：最高频任务决定骨架，其他任务填空隙。**

---

# 9. 这些题放在一起

```text
                    Heap
                      │
        ┌─────────────┼─────────────┐
        ↓             ↓             ↓
    取最大/最小       Top K        动态中位数
        │             │             │
     Max/Min       反向 Heap      双 Heap
        │             │             │
 Last Stone       Kth Largest   MedianFinder
 K Closest                        │
                                  │
                              Max + Min
```

另外：

```text
多个有序数据流
       ↓
K-way Merge
       ↓
Twitter / Merge K Lists
```

---

# 10. Heap 题识别口诀

刷题时先问自己：

### Q1：是不是一直在取最大/最小？

```text
是
↓
Heap
```

### Q2：是直接取，还是只保留 Top K？

```text
直接取最大 → MaxHeap
直接取最小 → MinHeap

Top K 最大 → MinHeap
Top K 最小 → MaxHeap
```

### Q3：是不是动态中位数？

```text
是
↓
MaxHeap + MinHeap
```

### Q4：是不是多个有序数据流？

```text
是
↓
K-way Merge
↓
Heap 保存每一路当前候选
```

---

# 11. 最终面试速记版

```text
┌────────────────────────────────────┐
│              HEAP                  │
├────────────────────────────────────┤
│ 取最大        → MaxHeap            │
│ 取最小        → MinHeap            │
│                                    │
│ Top K 最大    → MinHeap(size K)    │
│                 超 K → Pop 最小    │
│                                    │
│ Top K 最小    → MaxHeap(size K)    │
│                 超 K → Pop 最大    │
│                                    │
│ 动态中位数                           │
│   左 → MaxHeap                     │
│   右 → MinHeap                     │
│   两边 size ≤ 1                    │
│                                    │
│ K-way Merge                        │
│   每路放一个候选                    │
│   Pop 全局最优                      │
│   从对应路补下一个                  │
└────────────────────────────────────┘
```

### 目前这几道题，只需要牢牢记住：

```text
Last Stone
→ 取最大两个
→ MaxHeap

K Closest
→ 取最小距离
→ MinHeap

Kth Largest
→ 保留 K 个最大
→ MinHeap
→ 超过 K 淘汰最小

MedianFinder
→ 左 Max + 右 Min
→ 维护中间

Twitter
→ 多个有序 Tweet 流
→ K-way Merge + Heap

Task Scheduler
→ 最高频任务定骨架
→ 数学公式
```

**一句话总纲：**

> **Heap = 动态维护“当前最重要的边界元素”；直接取最值、维护 Top-K、维护中间位置、合并多路有序数据，是最核心的 4 个 Pattern。**
