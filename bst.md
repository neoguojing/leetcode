对，旋转数组这里最关键的就是**“如何判断哪一边有序”**。补充后建议记成下面这个版本：

# 二分查找：快速回忆版

## 1. 普通二分 —— 找值

**题意**：有序数组找 `target`。

```go
l, r := 0, len(nums)-1
for l <= r {
    m := l + (r-l)/2

    if nums[m] == target {
        return m
    } else if nums[m] < target {
        l = m + 1
    } else {
        r = m - 1
    }
}
return -1
```

> **找值 → `l <= r`**

---

## 2. Search Matrix —— 二维转一维

**思路**：把矩阵当成一维有序数组。

```go
row := m / cols
col := m % cols
```

然后普通二分。

> **二维有序 → 拉平二分**

---

## 3. 二分答案 —— Koko

**题意**：找最小满足条件的 `k`。

```text
k 越大 → 时间越少
× × × √ √ √
      ↑
    最小可行
```

```go
if hours <= h {
    r = m
} else {
    l = m + 1
}
```

> **最小可行值 → `check(mid)`**

---

# 4. 旋转数组最小值

例如：

```text
[4 5 6 7 0 1 2]
        ↑
       最小
```

核心：**比较 `nums[m]` 和 `nums[r]`。**

```go
if nums[m] < nums[r] {
    r = m
} else {
    l = m + 1
}
```

### 为什么？

```text
nums[m] < nums[r]
```

说明 `[m ... r]` 是正常递增的，最小值不会在 `m` 右边：

```go
r = m
```

否则：

```go
l = m + 1
```

> **找最小值 → `mid` 和 `right` 比**

---

# 5. 旋转数组找值 ⭐

例如：

```text
[4 5 6 7 0 1 2]
    l   m     r
```

关键问题：

> **如何判断哪一边有序？**

### 判断左边是否有序

```go
if nums[l] <= nums[m] {
    // [l ... m] 有序
}
```

因为旋转点不在 `[l,m]` 中。

例如：

```text
4 5 6 7 | 0 1 2
l     m
```

`4 <= 7`，所以左边有序。

---

### 否则右边有序

```go
else {
    // [m ... r] 有序
}
```

例如：

```text
4 5 6 | 7 0 1 2
      m       r
```

此时：

```text
nums[l] > nums[m]
```

说明旋转点在左边，右边 `[m...r]` 有序。

---

## 然后判断 target 是否在有序区间

### 左边有序

```go
if nums[l] <= nums[m] {
    if nums[l] <= target && target < nums[m] {
        r = m - 1
    } else {
        l = m + 1
    }
}
```

### 右边有序

```go
else {
    if nums[m] < target && target <= nums[r] {
        l = m + 1
    } else {
        r = m - 1
    }
}
```

### 完整核心

```go
if nums[m] == target {
    return m
}

if nums[l] <= nums[m] {          // 左边有序
    if nums[l] <= target && target < nums[m] {
        r = m - 1
    } else {
        l = m + 1
    }
} else {                         // 右边有序
    if nums[m] < target && target <= nums[r] {
        l = m + 1
    } else {
        r = m - 1
    }
}
```

### ⭐ 旋转数组必须记住

```text
nums[l] <= nums[m]
        ↓
左边有序

nums[l] > nums[m]
        ↓
右边有序
```

然后：

```text
有序的一边
    ↓
target 是否在这个范围？
    ↓
是 → 去有序区间
否 → 去另一边
```

> **旋转数组找值 = 先判断哪边有序 → 再判断 target 是否在有序范围。**

---

# 6. TimeMap —— 边界二分

**题意**：找 `timestamp <= target` 的最新值。

转换成：

> **找第一个 `> target`，然后 `idx - 1`**

```go
idx := sort.Search(len(pairs), func(i int) bool {
    return pairs[i].timestamp > timestamp
})

return pairs[idx-1].value
```

> **最后一个 ≤ → 第一个 > − 1**

---

# 7. Median —— 第 K 小

**思路**：比较两个数组的第 `k/2` 个元素，较小的一侧前 `k/2` 个一定不可能成为第 K 小。

```go
if a[aStart+i-1] > b[bStart+j-1] {
    // 删除 B 前 j 个
    k -= j
} else {
    // 删除 A 前 i 个
    k -= i
}
```

> **第 K 小 → 每次排除约 `k/2` 个**

---

# 二分最终记忆表

| 类型       | 核心判断                                 |
| -------- | ------------------------------------ |
| 普通二分     | `nums[m]` vs `target`                |
| 二分答案     | `check(mid)`                         |
| 旋转最小值    | `nums[m]` vs `nums[r]`               |
| **旋转找值** | **`nums[l] <= nums[m]` → 左有序，否则右有序** |
| 边界二分     | 第一个满足条件                              |
| 第 K 小    | 排除 `k/2` 个                           |

### 最重要的 5 句话

```text
普通二分：找 target

二分答案：找最小/最大可行值

旋转最小：mid 和 right 比

旋转找值：nums[l] <= nums[m] → 左边有序
          否则 → 右边有序

边界问题：最后一个 <= x → 第一个 > x - 1
```

这几个基本就是你这批二分题需要形成的**核心 Pattern**。
