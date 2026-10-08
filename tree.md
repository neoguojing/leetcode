可以，把它压缩成真正适合**刷题前快速回忆**的版本：只保留**题意 + 核心思路 + 状态公式 + 关键代码**。

# 🌳 Tree 树算法速记

## 1. 基础 DFS

### Maximum Depth

**题意：** 求树的最大深度。

**思路：** 后序 DFS，当前高度 = 左右子树最大高度 + 1。

```go
if root == nil { return 0 }

l := dfs(root.Left)
r := dfs(root.Right)

return max(l, r) + 1
```

**公式：**

```text
height(root) = max(height(L), height(R)) + 1
```

---

### Invert Binary Tree

**题意：** 左右子树交换。

**思路：** 递归翻转左右，再交换。

```go
l := invertTree(root.Left)
r := invertTree(root.Right)

root.Left, root.Right = r, l
```

**记忆：**

> 递归 → swap。

---

## 2. 后序 DFS + 全局答案

### Diameter of Binary Tree

**题意：** 求最长路径的边数。

**关键：**

```text
经过 root 的路径 = l + r
返回父节点 = max(l,r) + 1
```

```go
l := dfs(root.Left)
r := dfs(root.Right)

ans = max(ans, l+r)

return max(l, r) + 1
```

⭐ **核心：答案可以左右都要，但返回父节点只能选一边。**

---

### Binary Tree Maximum Path Sum

**题意：** 最大路径和，节点可以有负数。

**关键：**

```text
负贡献不要
当前答案 = l + root + r
返回父节点 = root + max(l,r)
```

```go
l := max(0, dfs(root.Left))
r := max(0, dfs(root.Right))

ans = max(ans, l+root.Val+r)

return root.Val + max(l, r)
```

⭐ 和 Diameter 是同一个模板。

---

### Balanced Binary Tree

**题意：** 任意节点左右高度差 ≤ 1。

**技巧：**

```text
正常 → 返回高度
失衡 → 返回 -1
```

```go
l := dfs(root.Left)
if l == -1 { return -1 }

r := dfs(root.Right)
if r == -1 { return -1 }

if abs(l-r) > 1 { return -1 }

return max(l,r) + 1
```

⭐ **一个 DFS 同时完成高度计算 + 平衡判断。**

---

# 3. 两棵树 / 结构比较

### Same Tree

**题意：** 两棵树结构和值都相同。

```go
if p == nil && q == nil { return true }
if p == nil || q == nil { return false }
if p.Val != q.Val { return false }

return same(p.Left,q.Left) &&
       same(p.Right,q.Right)
```

**记忆：**

```text
nil / nil → true
一个 nil → false
值不同 → false
否则比较左右
```

---

### Subtree

**题意：** `subRoot` 是否是 `root` 的子树。

**思路：**

```text
遍历 root
    ↓
每个节点尝试 sameTree
```

```go
if sameTree(root, subRoot) {
    return true
}

return isSubtree(root.Left, subRoot) ||
       isSubtree(root.Right, subRoot)
```

⭐ **Subtree = DFS 找位置 + SameTree 判断。**

---

# 4. LCA

### Lowest Common Ancestor

**题意：** 找 p、q 最近公共祖先。

**核心逻辑：**

```text
root == p/q → 返回 root

左、右都有结果 → root 就是 LCA

只有一边 → 返回这一边
```

```go
l := dfs(root.Left)
r := dfs(root.Right)

if l != nil && r != nil {
    return root
}

if l != nil {
    return l
}
return r
```

⭐ **左右各找到一个 → 当前节点就是答案。**

---

# 5. BFS

### Level Order

**题意：** 按层遍历。

**核心：**

```go
size := len(q)   // 当前层节点数

for i := 0; i < size; i++ {
    node := q[0]
    q = q[1:]

    // 处理 node
    // 加入左右孩子
}
```

⭐ **`size = 当前层大小` 是层序遍历关键。**

---

### Right Side View

**题意：** 每层最右边的节点。

本质：

```text
Level Order + 每层最后一个
```

```go
if i == size-1 {
    res = append(res, node.Val)
}
```

---

# 6. 自顶向下 DFS

### Good Nodes

**题意：** 从 root 到当前节点路径上，当前值 ≥ 历史最大值。

**状态：**

```text
maxVal = 当前路径最大值
```

```go
if node.Val >= maxVal {
    ans++
}

maxVal = max(maxVal, node.Val)

dfs(node.Left, maxVal)
dfs(node.Right, maxVal)
```

⭐ 和前面的后序不同：

> **父节点信息传给孩子。**

---

# 7. BST

### Validate BST

**核心性质：**

> BST 的中序遍历严格递增。

```go
dfs(root.Left)

if root.Val <= prev {
    return false
}

prev = root.Val

dfs(root.Right)
```

**记忆：**

```text
BST → 中序 → 递增
```

---

### Kth Smallest

**题意：** BST 第 k 小。

因为：

```text
BST 中序 = 从小到大
```

所以：

```go
dfs(root.Left)

k--
if k == 0 {
    ans = root.Val
}

dfs(root.Right)
```

**记忆：**

> BST + 中序 + 第 k 个。

---

# 8. 树的构造 / 序列化

### Construct Binary Tree

**题意：** preorder + inorder 构造树。

核心：

```text
preorder[0] = root
```

在 inorder 找 root：

```text
左边 → 左子树
右边 → 右子树
```

```go
rootVal := preorder[0]
idx := find(inorder, rootVal)

leftCount := idx
```

⭐ **Preorder 找 Root，Inorder 切左右。**

---

### Serialize / Deserialize

**思路：** Preorder + `N` 表示 nil。

例如：

```text
1,2,N,N,3,N,N
```

Serialize：

```go
res = append(res, value)
dfs(left)
dfs(right)
```

遇到空：

```go
res = append(res, "N")
```

Deserialize：

```go
if vals[i] == "N" {
    i++
    return nil
}

node := newNode(vals[i])
i++

node.Left = dfs()
node.Right = dfs()
```

⭐ **关键：序列化必须记录 nil，否则结构无法恢复。**

---

# 🌳 最终只记这张表

| 类型     | 题目           | 核心                                |
| ------ | ------------ | --------------------------------- |
| 后序     | Max Depth    | `max(l,r)+1`                      |
| 后序+全局  | Diameter     | `ans=l+r`，返回 `max(l,r)+1`         |
| 后序+全局  | Max Path Sum | `ans=l+root+r`，返回 `root+max(l,r)` |
| 后序+状态  | Balanced     | 高度 / `-1`                         |
| 双树 DFS | Same Tree    | nil + value + 左右                  |
| DFS    | Subtree      | DFS 找位置 + Same Tree               |
| DFS    | LCA          | 左右都有 → root                       |
| BFS    | Level Order  | `size=len(q)`                     |
| BFS    | Right View   | 每层最后一个                            |
| 前序     | Good Nodes   | 传 `maxVal`                        |
| 中序     | Validate BST | 严格递增                              |
| 中序     | Kth Smallest | 第 k 个                             |
| 递归构造   | Build Tree   | Preorder 找 root，Inorder 切分        |
| 前序     | Serialize    | Preorder + `N`                    |

## ⭐ 树题最重要的 4 个模板

```text
① 高度：
return max(l,r) + 1

② 路径：
ans = l + root + r
return root + max(l,r)

③ BST：
中序 → 严格递增

④ BFS：
size = len(queue)
→ 一层一层处理
```

**真正做题时，先问一句：**

> **“我的 `dfs(root)` 到底要返回什么？”**

这比记每一道树题的代码更重要。
