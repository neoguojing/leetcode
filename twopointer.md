下面按你前面整理算法的风格，压缩成一个**方便快速回忆的「双指针 Pattern」文档**。核心不是记 5 道题，而是记住：**左右指针为什么移动**。

# 双指针 Two Pointers

## 1. 核心思想

> **用两个指针从两端或同向移动，利用数据的有序性/单调性，减少枚举。**

常见两类：

| 类型        | 典型问题                                       | 核心          |
| --------- | ------------------------------------------ | ----------- |
| **相向双指针** | Valid Palindrome、Two Sum II、3Sum、Container | `l →`，`r ←` |
| **单向双指针** | 快慢指针、滑动窗口                                  | `l →`，`r →` |

本组题目主要是**相向双指针**。

---

# 2. 相向双指针统一模板

```go
l, r := 0, len(nums)-1

for l < r {
    // 根据当前状态判断

    if ... {
        l++
    } else if ... {
        r--
    } else {
        // 找到答案
        l++
        r--
    }
}
```

### 最关键的问题

不是记代码，而是问：

> **当前应该移动哪一边？为什么移动这一边不会错过答案？**

---

# 3. Valid Palindrome

### 题意

忽略非字母/数字以及大小写，判断字符串是否回文。

### 核心思路

```text
l →          ← r
```

两端字符比较：

* 左边非法字符 → `l++`
* 右边非法字符 → `r--`
* 都合法 → 比较
* 不相等 → `false`
* 相等 → `l++, r--`

### 关键代码

```go
for l < r {
    for l < r && !isAlphaNum(rune(s[l])) {
        l++
    }
    for l < r && !isAlphaNum(rune(s[r])) {
        r--
    }

    if unicode.ToLower(rune(s[l])) !=
       unicode.ToLower(rune(s[r])) {
        return false
    }

    l++
    r--
}
return true
```

### 记忆点

> **两边向中间：跳过无效字符 → 比较 → 同时收缩。**

---

# 4. Two Sum II

### 题意

**有序数组**中找两个数，使：

```text
nums[l] + nums[r] = target
```

### 核心公式

```text
sum > target → r--
sum < target → l++
sum = target → 找到
```

### 为什么？

因为数组有序：

```text
nums[l] ↑       nums[r] ↓
```

如果：

```text
nums[l] + nums[r] > target
```

当前 `r` 太大，所以：

```text
r--
```

反之：

```text
sum < target → l++
```

### 关键代码

```go
for l < r {
    sum := numbers[l] + numbers[r]

    if sum > target {
        r--
    } else if sum < target {
        l++
    } else {
        return []int{l + 1, r + 1}
    }
}
```

### 记忆

> **有序 + 两数和：大了右减，小了左加。**

---

# 5. 3Sum

### 题意

找所有：

```text
a + b + c = 0
```

### 核心思想

**3 个数 → 固定一个 + Two Sum。**

```text
for i:
    固定 nums[i]

    l = i + 1
    r = end

    用双指针找：
    nums[i] + nums[l] + nums[r] = 0
```

### 状态迁移

```text
sum > 0 → r--
sum < 0 → l++
sum = 0 → 记录 + l++ + r--
```

### 为什么必须排序？

排序以后才能利用：

```text
sum 大 → 减小 r
sum 小 → 增大 l
```

同时方便去重。

### 两层去重

固定元素：

```go
if i > 0 && nums[i] == nums[i-1] {
    continue
}
```

找到答案后：

```go
l++
r--

for l < r && nums[l] == nums[l-1] {
    l++
}
```

### 关键代码

```go
sort.Ints(nums)

for i := 0; i < len(nums); i++ {
    if nums[i] > 0 {
        break
    }

    if i > 0 && nums[i] == nums[i-1] {
        continue
    }

    l, r := i+1, len(nums)-1

    for l < r {
        sum := nums[i] + nums[l] + nums[r]

        if sum < 0 {
            l++
        } else if sum > 0 {
            r--
        } else {
            res = append(res,
                []int{nums[i], nums[l], nums[r]})

            l++
            r--

            for l < r && nums[l] == nums[l-1] {
                l++
            }
        }
    }
}
```

### 记忆

> **3Sum = 排序 + 固定一个 + Two Sum + 去重。**

---

# 6. Container With Most Water

### 题意

两根柱子组成容器：

```text
area = min(height[l], height[r]) × (r-l)
```

### 最关键的思维

宽度：

```text
r-l
```

每次移动指针，宽度一定减少。

所以：

> **必须移动较短的那根柱子。**

例如：

```text
height[l] < height[r]

面积受 height[l] 限制

l++
```

因为移动较高的柱子：

```text
min(height[l], height[r])
```

仍然最多受原来的矮柱限制，同时宽度还变小，**不可能得到更优解**。

### 公式

```text
area = min(h[l], h[r]) × (r-l)

h[l] <= h[r] → l++
h[l] >  h[r] → r--
```

### 关键代码

```go
for l < r {
    area := min(height[l], height[r]) * (r-l)
    res = max(res, area)

    if height[l] <= height[r] {
        l++
    } else {
        r--
    }
}
```

### 记忆

> **容器：谁矮移动谁。**

这是这道题最重要的一句话。

---

# 7. Trapping Rain Water

### 题意

每个位置能装多少水：

```text
water[i] = min(leftMax, rightMax) - height[i]
```

### 双指针核心

维护：

```text
leftMax
rightMax
```

如果：

```text
leftMax < rightMax
```

那么左边当前位置的水量已经可以确定：

```text
water = leftMax - height[l]
```

所以：

```text
l++
```

反之：

```text
rightMax <= leftMax
→ r--
```

### 核心公式

```text
leftMax < rightMax
    → l++
    → water = leftMax - height[l]

leftMax >= rightMax
    → r--
    → water = rightMax - height[r]
```

### 关键代码

```go
l, r := 0, len(height)-1
leftMax, rightMax := height[l], height[r]

for l < r {
    if leftMax < rightMax {
        l++
        leftMax = max(leftMax, height[l])
        res += leftMax - height[l]
    } else {
        r--
        rightMax = max(rightMax, height[r])
        res += rightMax - height[r]
    }
}
```

### 记忆

> **接雨水：比较左右最高墙，较小的一侧先结算。**

---

# 8. 这 5 道题统一记忆

| 题目                      | 双指针规则                         |
| ----------------------- | ----------------------------- |
| **Palindrome**          | 两边比较，不合法就跳过                   |
| **Two Sum II**          | `sum大 → r--`，`sum小 → l++`     |
| **3Sum**                | 固定一个，剩余部分做 Two Sum            |
| **Container**           | **矮的移动**                      |
| **Trapping Rain Water** | **较小的 leftMax/rightMax 一侧结算** |

---

# 9. 最重要的 Pattern

真正需要记住的是下面这张表：

```text
双指针
│
├── 两端比较
│   └── Palindrome
│
├── 有序数组 + 目标值
│   ├── 2Sum
│   └── 3Sum
│
├── 最大化面积
│   └── Container
│       └── 矮的移动
│
└── 根据边界确定当前位置
    └── Trapping Rain Water
        └── 较小 Max 的一侧先结算
```

### 面试时的快速判断

看到题目先问 3 个问题：

```text
① 数组是否有序？
       ↓
   能否利用大小关系移动指针？

② 两端是否可以同时向中间收缩？
       ↓
   相向双指针

③ 移动哪一个不会错过最优解？
       ↓
   找到“决定性约束”
```

**一句话总记忆：**

> **双指针不是“两个指针一起走”，而是利用单调性判断“哪一个指针可以安全移动”。**
