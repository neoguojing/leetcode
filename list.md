下面按**面试快速回忆**整理，只保留「题意 → 核心思路 → 关键代码」。

# 链表算法速记

## 1. 反转链表 — Reverse Linked List

**思路：三指针**

`prev ← cur → next`

核心就是保存 `next`，再把 `cur.Next` 指向 `prev`。

```go
var prev *ListNode

for head != nil {
    next := head.Next
    head.Next = prev
    prev = head
    head = next
}

return prev
```

**记忆：** `保存 → 反向 → 前进`

---

## 2. 合并两个有序链表

**思路：双指针 + dummy**

每次取较小节点接到结果链表。

```go
dummy := &ListNode{}
cur := dummy

for l1 != nil && l2 != nil {
    if l1.Val <= l2.Val {
        cur.Next = l1
        l1 = l1.Next
    } else {
        cur.Next = l2
        l2 = l2.Next
    }
    cur = cur.Next
}

if l1 != nil {
    cur.Next = l1
} else {
    cur.Next = l2
}

return dummy.Next
```

**记忆：** `dummy + 双指针 + 剩余直接接`

---

## 3. 环形链表 — Linked List Cycle

**思路：快慢指针**

* slow：一步
* fast：两步
* 有环 → 最终相遇
* 无环 → fast 到 nil

```go
slow, fast := head, head

for fast != nil && fast.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next

    if slow == fast {
        return true
    }
}
return false
```

**记忆：** 判断环 = `slow/fast 是否相遇`

---

## 4. 重排链表 — Reorder List

目标：

```text
1 → 2 → 3 → 4 → 5
↓
1 → 5 → 2 → 4 → 3
```

**三步：**

1. 找中点
2. 反转后半段
3. 两个链表交替合并

```go
// ① 找中点
slow, fast := head, head.Next
for fast != nil && fast.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
}

// ② 断开 + 反转后半段
second := slow.Next
slow.Next = nil
second = reverse(second)

// ③ 交替合并
first := head
for second != nil {
    n1, n2 := first.Next, second.Next

    first.Next = second
    second.Next = n1

    first, second = n1, n2
}
```

**记忆：** `中点 → 后半反转 → 交替合并`

---

## 5. 删除倒数第 N 个节点

**思路：dummy + 双指针保持 n 距离**

先让 `fast` 前进 `n` 步，再一起走。

```go
dummy := &ListNode{Next: head}
slow, fast := dummy, dummy

for n > 0 {
    fast = fast.Next
    n--
}

for fast.Next != nil {
    slow = slow.Next
    fast = fast.Next
}

slow.Next = slow.Next.Next

return dummy.Next
```

**为什么用 dummy？**

统一处理：

```text
删除头节点
删除中间节点
删除尾节点
```

**记忆：** `dummy + fast 先走 n 步`

---

## 6. 随机链表复制 — Copy List with Random Pointer

每个节点有：

```text
Next + Random
```

**思路：哈希表建立映射**

```text
old node → copy node
```

第一遍创建节点：

```go
oldToCopy := map[*Node]*Node{nil: nil}

for cur := head; cur != nil; cur = cur.Next {
    oldToCopy[cur] = &Node{Val: cur.Val}
}
```

第二遍连接指针：

```go
for cur := head; cur != nil; cur = cur.Next {
    copy := oldToCopy[cur]

    copy.Next = oldToCopy[cur.Next]
    copy.Random = oldToCopy[cur.Random]
}
```

**关键：**

```go
map[原节点]复制节点
```

`nil → nil` 可以避免额外判断。

---

## 7. 两数相加 — Add Two Numbers

链表存储的是**低位 → 高位**：

```text
2 → 4 → 3
5 → 6 → 4
↓
7 → 0 → 8
```

**思路：逐位相加 + carry**

公式：

```text
sum = a + b + carry
digit = sum % 10
carry = sum / 10
```

关键代码：

```go
sum := l1.Val + l2.Val + carry

cur.Next = &ListNode{Val: sum % 10}
carry = sum / 10
cur = cur.Next
```

最后：

```go
if carry != 0 {
    cur.Next = &ListNode{Val: carry}
}
```

**记忆：** 和小学竖式加法完全一样。

---

# 8. 找重复数字 — Find Duplicate Number

虽然输入是数组，但本质是**链表找环**。

因为：

```go
next = nums[i]
```

把数组看成：

```text
i → nums[i]
```

重复数字意味着存在环。

**Floyd 两阶段：**

```go
// ① 找环内相遇点
slow, fast := 0, 0

for {
    slow = nums[slow]
    fast = nums[nums[fast]]

    if slow == fast {
        break
    }
}

// ② 找环入口 = 重复数字
slow2 := 0
for slow != slow2 {
    slow = nums[slow]
    slow2 = nums[slow2]
}

return slow
```

**记忆：**

> 数组 → 映射成链表 → 找环 → 环入口就是重复数字

---

# 9. LRU Cache

这是链表题中非常重要的一题。

**核心结构：**

```text
HashMap + 双向链表
```

```text
map[key] → Node

left ⇄ ... ⇄ right
       ↑
    最近使用
```

约定：

* `right` 附近 = 最新
* `left` 附近 = 最旧

### Get

找到节点后：

```text
删除旧位置
→ 插入最新位置
→ 返回 value
```

```go
if node, ok := cache[key]; ok {
    remove(node)
    insert(node)
    return node.val
}
```

### Put

```text
已存在 → 删除旧节点
→ 插入新节点到 right
→ 超容量 → 删除 left.next
```

### 核心操作

```go
func remove(node *Node) {
    node.prev.next = node.next
    node.next.prev = node.prev
}

func insert(node *Node) {
    prev := right.prev
    prev.next = node
    right.prev = node
    node.prev = prev
    node.next = right
}
```

**记忆：**

> LRU = `HashMap O(1) 查找` + `双向链表 O(1) 移动/删除`

---

# 10. 合并 K 个有序链表

你的代码采用：

**分治 + 两两合并**

```text
[k个链表]
    ↓
左右拆分
 ↙     ↘
递归    递归
 ↘     ↙
 两两合并
```

核心：

```go
mid := left + (right-left)/2

l1 := divide(lists, left, mid)
l2 := divide(lists, mid+1, right)

return merge(l1, l2)
```

本质就是：

```text
Merge Sort
```

**记忆：**

> K 个 → 分成两半 → 递归 → 两个有序链表合并

复杂度：

```text
O(N log K)
```

---

# 11. K 个一组反转链表

**核心：递归 + 每次反转 K 个**

首先判断剩余节点是否够 `k` 个：

```go
cur := head
for i := 0; i < k; i++ {
    if cur == nil {
        return head
    }
    cur = cur.Next
}
```

然后：

```text
先处理后面的组
↓
再反转当前 K 个
```

反转核心：

```go
next := head.Next
head.Next = cur
cur = head
head = next
```

最终：

```go
return cur
```

**最重要的理解：**

```text
检查够不够 K 个
        ↓
递归处理后面的组
        ↓
当前组反转
        ↓
接到后面的结果
```

---

# 链表题 Pattern 总结

| Pattern         | 代表题             | 核心                  |
| --------------- | --------------- | ------------------- |
| **反转**          | Reverse List    | `prev / cur / next` |
| **双指针**         | Cycle           | `slow / fast`       |
| **间距双指针**       | Remove Nth      | `fast 先走 n`         |
| **中点+反转**       | Reorder         | `找中点 → 反转 → 合并`     |
| **Dummy**       | Merge / Remove  | 简化头节点处理             |
| **HashMap**     | Copy Random     | `old → copy`        |
| **模拟**          | Add Two Numbers | `sum + carry`       |
| **链表化**         | Find Duplicate  | 数组 → 环              |
| **HashMap+双链表** | LRU             | O(1) 查找+移动          |
| **分治**          | Merge K Lists   | `Divide + Merge`    |
| **递归反转**        | Reverse K Group | `K个一组`              |

### 最值得背的 6 个模板

```text
① 反转：prev / cur / next

② 快慢：slow / fast

③ 倒数第N：dummy + fast先走N

④ 重排：中点 → 反转 → 交错合并

⑤ LRU：map + 双向链表

⑥ Merge K：分治 + merge
```

这几类基本覆盖你这批题的核心思维。
