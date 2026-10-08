下面按 **Hash / 数组计数 / 前缀后缀 / 堆** 归类，重点保留**题意 + 核心思路 + 关键代码**，方便面试快速回忆。

# 数组 & Hash 算法

## 1. Valid Anagram

**题意**：判断两个字符串是否由相同字符组成。

**思路**：字符计数，`s +1`，`t -1`，最后全部为 0。

```go
var cnt [26]int
for i := range s {
    cnt[s[i]-'a']++
    cnt[t[i]-'a']--
}
```

**关键点**：固定小字符集 → `[26]int` 比 `map` 更简单。

---

## 2. Two Sum

**题意**：找两个数，使 `nums[i] + nums[j] = target`。

**核心公式**

```text
need = target - nums[i]
```

**思路**：HashMap 保存 `值 → 下标`。

```go
for i, n := range nums {
    if j, ok := mp[target-n]; ok && i != j {
        return []int{i, j}
    }
    mp[n] = i
}
```

**注意**：可以边遍历边查，避免自己匹配自己。

---

## 3. Group Anagrams

**题意**：把字母组成相同的字符串分组。

**思路**：每个字符串生成唯一签名：

```text
"eat" → [1,0,0,0,1,...,1]
```

```go
var cnt [26]int
for _, c := range str {
    cnt[c-'a']++
}
mp[cnt] = append(mp[cnt], str)
```

**关键点**：

```go
map[[26]int][]string
```

数组可以作为 Go 的 map key。

---

## 4. Top K Frequent Elements

**题意**：找出现频率最高的 K 个元素。

**两步**：

```text
① HashMap：统计频率
② MinHeap：维护大小为 K 的堆
```

核心：

```go
for n, cnt := range mp {
    heap.Push(h, [2]int{n, cnt})

    if h.Len() > k {
        heap.Pop(h)
    }
}
```

**为什么最小堆？**

```text
堆顶 = 当前 K 个元素中频率最低的
超过 K → 删除堆顶
```

复杂度：

```text
O(n) 统计
O(m log k) 堆
```

---

## 5. Encode and Decode Strings

**题意**：字符串数组 → 一个字符串，并能够无歧义还原。

**核心问题**：不能直接用 `#` 分隔，因为原字符串本身可能包含 `#`。

**解决**：记录长度：

```text
3#abc5#hello
```

格式：

```text
length + "#" + string
```

Encode：

```go
res.WriteString(strconv.Itoa(len(str)))
res.WriteByte('#')
res.WriteString(str)
```

Decode：

```go
j := i
for encoded[j] != '#' {
    j++
}

length, _ := strconv.Atoi(encoded[i:j])
i = j + 1

res = append(res, encoded[i:i+length])
i += length
```

**关键思想**：`长度 + 分隔符` → 无歧义序列化。

---

# 数组技巧

## 6. Product of Array Except Self

**题意**：`res[i] = 除 nums[i] 外所有元素的乘积`，不能使用除法。

**核心公式**

```text
res[i] = 左边乘积 × 右边乘积
```

例如：

```text
nums = [1,2,3,4]

left  = [1,1,2,6]
right = [24,12,4,1]

res   = [24,12,8,6]
```

关键：

```go
prefix[i] = prefix[i-1] * nums[i-1]
suffix[i] = suffix[i+1] * nums[i+1]

res[i] = prefix[i] * suffix[i]
```

**面试点**：进一步可把 `prefix` / `suffix` 压缩成 `O(1)` 额外空间。

---

## 7. Valid Sudoku

**题意**：判断数独当前状态是否合法。

**三个维度检查**：

```text
row[r] → 第 r 行
col[c] → 第 c 列
box[b] → 第 b 个 3×3 宫格
```

用 bitmask：

```go
val := board[r][c] - '1'
bit := 1 << val

if rows[r]&bit != 0 ||
   cols[c]&bit != 0 ||
   boxes[idx]&bit != 0 {
    return false
}
```

标记：

```go
rows[r] |= bit
cols[c] |= bit
boxes[idx] |= bit
```

宫格编号：

```go
idx := (r/3)*3 + c/3
```

**核心思想**：

```text
一个 int 的 9 个 bit → 记录 1~9 是否出现
```

---

## 8. Longest Consecutive Sequence

**题意**：无序数组中找最长连续数字长度。

例如：

```text
[100,4,200,1,3,2]
→ 1,2,3,4
→ 4
```

**核心思路**：HashSet 判断数字是否存在。

关键判断：

```go
if _, ok := set[num-1]; !ok {
    // num 是连续序列起点
}
```

然后：

```go
cur := num
for set[cur+1] {
    cur++
}
```

**为什么只从起点开始？**

```text
num-1 不存在 → num 才可能是序列起点
```

复杂度：

```text
O(n)
```

---

### 你的代码使用的核心 Pattern

| 题目                  | Pattern           | 一句话记忆           |
| ------------------- | ----------------- | --------------- |
| Valid Anagram       | 数组计数              | `+1/-1`         |
| Two Sum             | HashMap           | `target - x`    |
| Group Anagrams      | HashMap + 签名      | `[26]int` 做 key |
| Top K Frequent      | HashMap + MinHeap | 堆维护 Top K       |
| Encode/Decode       | 编解码               | `长度#字符串`        |
| Product Except Self | Prefix/Suffix     | 左 × 右           |
| Valid Sudoku        | Bitmask           | 行/列/宫格记 bit     |
| Longest Consecutive | HashSet           | 只从起点扩展          |

### 最重要的 5 个记忆模板

```text
① 计数：
map / [26]int

② 查补数：
need = target - x

③ 分组：
构造 signature → map[signature][]value

④ Top K：
频率 map → 大小 K 的 MinHeap

⑤ 连续序列：
HashSet + 只从 num-1 不存在的位置开始
```

**整体理解**：这一组题本质上主要不是“数组操作”，而是**利用 Hash 把查找从 O(n) 降到 O(1)**，再配合计数、签名、前缀/后缀、堆等技巧解决不同问题。

如果你继续按这个方式整理，**Arrays & Hashing 基本就可以压缩成上面 5 个模板**，刷题时优先判断“这题属于哪个模板”，而不是重新想完整算法。
