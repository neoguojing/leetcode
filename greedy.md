可以，把它压缩成**刷题快速回忆版**：只保留「题意 → 核心思路/公式 → 关键代码 → 记忆点」。

# Greedy 贪心算法｜快速回忆版

## 1. Greedy 核心

**每一步做一个局部最优选择，并证明这个选择不会破坏全局最优。**

常见证明：

```text
交换论证：换成我的选择，不会更差
状态支配：A 永远比 B 好 → 丢掉 B
不可逆：一旦越界/失败 → 永远无法恢复
```

---

## 2. Assign Cookies

**题意**：每个孩子有最小需求，每块饼干有大小，最多满足多少孩子？

**思路**：排序后，小需求匹配最小够用的饼干。

```go
sort.Ints(g)
sort.Ints(s)

i, j := 0, 0
for i < len(g) && j < len(s) {
    if s[j] >= g[i] {
        i++
    }
    j++
}
return i
```

**记忆**：小需求 → 小资源。

---

## 3. Non-overlapping Intervals

**题意**：删除最少区间，使剩余区间互不重叠。

**思路**：按 `right` 升序，优先保留结束最早的区间。

```go
sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][1] < intervals[j][1]
})

end, remove := math.MinInt, 0

for _, in := range intervals {
    if in[0] >= end {
        end = in[1]
    } else {
        remove++
    }
}
return remove
```

**记忆**：结束越早，给后面空间越大。

---

## 4. Maximum Subarray

**题意**：找连续子数组最大和。

**公式**：

```text
cur = max(nums[i], cur + nums[i])
```

含义：

```text
继续之前的子数组
    OR
从当前重新开始
```

```go
cur, ans := nums[0], nums[0]

for i := 1; i < len(nums); i++ {
    cur = max(nums[i], cur+nums[i])
    ans = max(ans, cur)
}
return ans
```

**记忆**：累计变负，历史就丢掉。

---

## 5. Jump Game

**题意**：能否到达最后位置？

**状态**：

```text
maxReach = 当前最远可达位置
```

```go
maxReach := 0

for i := 0; i < len(nums); i++ {
    if i > maxReach {
        return false
    }
    maxReach = max(maxReach, i+nums[i])
}
return true
```

**记忆**：更远支配更近。

---

## 6. Jump Game II

**题意**：到达最后位置最少跳几次？

**思路**：当前跳跃覆盖一个区间，在区间内找下一跳最远位置。

```text
farthest = max(farthest, i + nums[i])
```

到当前边界：

```text
i == curEnd
→ 必须再跳一次
→ curEnd = farthest
```

```go
jumps, curEnd, farthest := 0, 0, 0

for i := 0; i < len(nums)-1; i++ {
    farthest = max(farthest, i+nums[i])

    if i == curEnd {
        jumps++
        curEnd = farthest
    }
}
return jumps
```

**记忆**：当前层找最远，边界到了再跳。

---

## 7. Gas Station

**题意**：环形路线，每站获得 `gas`，行驶消耗 `cost`，找能走完整圈的起点。

**公式**：

```text
diff[i] = gas[i] - cost[i]
```

先判断：

```text
Σgas < Σcost → 无解
```

扫描：

```text
total += diff[i]

total < 0
→ 当前起点失败
→ start = i + 1
→ total = 0
```

```go
total, start := 0, 0

for i := range gas {
    total += gas[i] - cost[i]

    if total < 0 {
        total = 0
        start = i + 1
    }
}
return start
```

**记忆**：累计变负，换起点。

---

## 8. Hand of Straights

**题意**：把牌分成 `groupSize` 张连续牌。

**思路**：排序后，当前最小未使用牌必须作为一组开头。

```text
v → v+1 → ... → v+groupSize-1
```

```go
sort.Ints(hand)

cnt := map[int]int{}
for _, v := range hand {
    cnt[v]++
}

for _, v := range hand {
    if cnt[v] == 0 {
        continue
    }

    for x := v; x < v+groupSize; x++ {
        if cnt[x] == 0 {
            return false
        }
        cnt[x]--
    }
}
return true
```

**记忆**：最小牌没有退路。

---

## 9. Partition Labels

**题意**：切分字符串，使同一个字符只出现在一个区间，并且区间尽可能多。

**状态**：

```text
last[c] = c 最后出现位置
right = 当前区间必须到达的最远位置
```

```text
right = max(right, last[c])

i == right → 可以切
```

```go
last := make([]int, 26)

for i, c := range s {
    last[c-'a'] = i
}

start, right := 0, 0
res := []int{}

for i, c := range s {
    right = max(right, last[c-'a'])

    if i == right {
        res = append(res, i-start+1)
        start = i + 1
    }
}
return res
```

**记忆**：覆盖所有字符的最后位置，才能切。

---

## 10. Merge Triplets to Form Target

**题意**：通过 `max` 合并三元组，能否得到 `target=[x,y,z]`？

**关键性质**：

```text
max 只能变大
```

所以：

```text
a > x || b > y || c > z
→ 永远不可能参与
```

剩余三元组只需要分别覆盖：

```text
a == x
b == y
c == z
```

```go
found := [3]bool{}

for _, t := range triplets {
    if t[0] > target[0] ||
       t[1] > target[1] ||
       t[2] > target[2] {
        continue
    }

    if t[0] == target[0] { found[0] = true }
    if t[1] == target[1] { found[1] = true }
    if t[2] == target[2] { found[2] = true }
}

return found[0] && found[1] && found[2]
```

**记忆**：超过目标不可逆，先淘汰。

---

# 贪心快速判断

看到题目先问：

```text
① 能不能排序后做选择？
② 有没有“最远 / 最早 / 最小 / 最大”？
③ 某些状态能不能证明永远没用？
④ 失败后能不能永久换掉当前选择？
```

最核心：

> **贪心 = 找到可以永久淘汰其他选择的理由。**
