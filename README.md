# leetcode
## 贪心算法：贪心算法不从整体最优上加以考虑，而是一步一步进行，每一步只以当前情况为基础，根据某个优化测度做出局部最优选择
- 贪⼼选择性质：一个问题的全局最优解可以通过一系列局部最优解（贪心选择）来得到。
- 最优子结构：指的是一个问题的最优解包含其子问题的最优解
- 判断一个问题是否通过贪心算法求解，是需要进行严格的数学证明的
- 分发饼干
- 重叠区间（反向思维）问题转换：当选择结束时间最早的区间之后，再在剩下的时间内选出最多的区间

## 递归算法： 种通过重复将原问题分解为同类的子问题而解决的方法
- 使用缓存避免重复计算
## 分治算法： 把规模大的问题不断分解为子问题，使得问题规模减小到可以直接求解为止。
- 分治算法从实现方式上来划分，可以分为两种：「递归算法」和「迭代算法」。
- 迭代算法： 如傅里叶变换、二分查找等
- 归并排序： 属于迭代算法
- 
## 回溯法：递归结束条件，剪枝约束条件和
- 一种能避免不必要搜索的穷举式的搜索算法
- 走不通就退回再走的技术称为「回溯法」，而满足回溯条件的某个状态的点称为「回溯点」
- 1.找到一个可能存在的正确答案；
- 2.在尝试了所有可能的分布方法之后宣布该问题没有答案。
- 全排列、子集、n皇后
```
明确所有选择：画出搜索过程的决策树，根据决策树来确定搜索路径。
明确终止条件：推敲出递归的终止条件，以及递归终止时的要执行的处理方法。
将决策树和终止条件翻译成代码：
定义回溯函数（明确函数意义、传入参数、返回结果等）。
书写回溯函数主体（给出约束条件、选择元素、递归搜索、撤销选择部分）。
明确递归终止条件（给出递归终止条件，以及递归终止时的处理方法）。

```
### 通用模式：
```
func Backword() []string{
  backword(cond1,cond2,oneResult,&allResult)
}

<!-- cond1为回溯过滤条件；cond2为递归结束条件；oneResult为临时计算的值，allResult为所有计算的值 -->
func backword(cond1,cond2,oneResult,*allResult){
  <!--  填充本轮计算的值  -->
  if cond2 {
    *allResult = *append(allResult,oneResult)
  }
  <!--  n为每个位置可选择的值的个数，循环控制在每个位置依次尝试不同的值  -->
  for i<n {
    <!-- 过滤回溯条件   -->
    if cond1 {
     continue
    }
    <!--    选择i位置的值    -->
    backword(cond1,cond2,oneResult,*allResult)
     <!--    此处回滚上一步的操作，进行下一次尝试    -->
  }
}

```
### 算法题
- 22 根据数字生成对应对数的圆括号的组合，圆括号必须合法，先左后右
- > 问题：1.如何判定合法？即约束条件：已经放置的左括号数量大于右括号数量；2.结束条件，n对括号全部打印
- > 回溯/DFS，因为每个位置只有两种情况，因此不需要循环，而是递归两次即可
- 
## 动态规划 是一种求解多阶段决策过程最优化问题的方法
- 自顶向下的记忆化搜索方法: 是使用递归的，栈可能溢出，可以处理复杂问题
- 自底向上的递推方法： 使用dp数组的方法，使用循环
- 适用于动态规划求解的问题，在分解之后得到的子问题往往是相互联系的，会出现若干个重叠子问题
- 使用动态规划方法会将这些重叠子问题的解保存到表格里，供随后的计算查询使用，从而避免大量的重复计算。
- 关键点在于「如何状态设计」和「推导状态转移条件」，还有各种各样的「优化方法」
- 最优子结构性质：指的是一个问题的最优解包含其子问题的最优解
- 重叠子问题性质：指的是在求解子问题的过程中，有大量的子问题是重复的
- 无后效性：指的是子问题的解（状态值）只与之前阶段有关，而与后面阶段无关
- 求走法，种类等问题：如91，62 70等，到达某个解的种类是固定的比如1步还是两步，比如向下还是向右，比如一个字母还是两个字母，通常类型dp[i] = dp[i-1]+dp[i-2]
- 求最短，求最长等，通常涉及到max和min的求解
- 数组范围：需要分左和右dp两部分求解
### 记忆优化搜索 是一种通过存储已经遍历过的状态信息，从而避免对同一状态重复遍历的搜索算法。
- 问题复杂，递推关系比明显
- 需要一个缓存保存问题的解；
- 使用递归： 例如：目标和（p+n=sum(nums);p-n=target）; 第N个fib数
### 线性动态规划
- 背包问题、区间 DP、数位 DP 等都属于线性DP
- 一维线性 DP 问题、二维线性 DP 问题，以及多维线性 DP 问题。
- 单串线性 DP 问题：问题的输入为单个数组或单个字符串的线性 DP 问题； dp[i] ； 如最长增长子串，最大子数组和
- - 最长的斐波那契子序列的长度： 需要考虑两个结束位置；dp[i][j] 表示以 arr[i] 和 arr[j] 为结尾的斐波那契式子序列的最大长度。
- 双串线性 DP 问题：问题的输入为两个数组或两个字符串的线性 DP 问题。状态一般可定义为 dp[i][j]
- - 最长公共子序列：dp[i][j] 为：「以 text1 的前 i 个元素组成的子字符串」与「以 text2 的前 j 个元素组成的子字符串」的最长公共子序列长度。第0个元素标识空字符串
  - 最长重复子数组： 同上，状态方程有部分不同
  - 编辑距离：同上，状态方程需要考虑两个分支，4总情况
- 矩阵线性DP问题：问题的输入为二维矩阵的线性 DP 问题；状态为dp[i][j]
- - 最小路径和：dp[i][j] 表示从左上角 (0, 0) 到位置 (i, j) 的路径上的最小和；可以使用一维数组节省空间；其中d[j] 标识上一行的状态，d[j-1] 标识当前行的状态
  - dp[i][j] 表示 以矩阵位置 (i, j) 为右下角，且包含 1 的最大正方形的边长。
- 无串线性 DP 问题：问题的输入不是显式的数组或字符串，但依然可分解为若干子问题的线性 DP 问题
- - 整数拆分：dp[i]将正整数 i 拆分为至少 2 个正整数的和后，这些正整数的最大乘积；分为2次拆分和多次拆
  - 只有两个键的键盘: dp[i] = min(dp[i], dp[j] + i/j, dp[i/j] + j);i%j=0; 举例：6/2=3 (copy,paste)=2 （copy，paste，paste）= 6

#### 区间动态规划：属于线性动态规划
- 先在小区间内得到最优解，再利用小区间的最优解合并，从而得到大区间的最优解，最终得到整个区间的最优解。
- 单个区间从中间向两侧更大区间转移的区间 DP 问题：回文串划分问题，矩阵链乘法：dp[i][j]=max{dp[i+1][j−1],dp[i+1][j],dp[i][j−1]}+cost[i][j],i≤j
- 多个小区间转移到大区间的区间 DP 问题：石子合并，括号匹配
- 最长回文子序列: 字符相等dp[i+1][j−1]+2}  字符不等max{dp[i][j−1],dp[i−1][j]}
- 挫气球：左右添加气球，改造问题；假设最后戳破的气球为k
- 切棍子的最小成本：切点排序，cost := dp[i][k] + dp[k][j] + (cuts[j] - cuts[i])
#### 树形dp问题
- 定根
- - 二叉树的最大路径和：路径可以穿过和不穿过根节点；需要考虑为负时舍弃
  - 相邻字符不同的最长路径：构建邻接表表示图，寻找图的最长路径，过程中过滤相邻节点相同的路径
- 不定根
- - 最小高度树：二次遍历与换根法法； 寻找树的中心节点；对于树来说，最多只能有两个相邻的节点作为中心

#### 状态压缩dp
- 小规模数据」的数组 / 字符串上，结合「二进制」的性质来进行状态定义与状态转移的动态规划方法。
- 对于一个长度为 n 的集合 S，可以利用「二进制」的性质枚举其所有子集
- 我们通过一个n 位长度的二进制数」来表示「由 n 个物品所组成的集合中所有物品的选择状态」。
- 对于元素个数不超过 n 的集合来说，一共会出现 2 n 个状态数量。因为在 n 变大时会呈现指数级增长，所以状态压缩 DP 只适用于求解小数据规模问题（通常 n≤20
- 总状态数量：1 << n
- 统计状态中1的个数，标识已经选择的状态数：需要函数计算
- mask&(1<<i) ： 判断i位是否被选择
- mask | (1 << j)： 设置j位为1
- 
- 数组的低位是状态的低位，两厢逆序
- - 两个数组最小的异或值之和：
  - - dp[newMask] = min(dp[newMask], dp[mask] + (nums1[count] ^ nums2[j]))
    - mask为所有总状态和即1<<n,从0开始遍历
    - count为已经匹配过的元素的个数nums1[count] 当前要匹配的元素；
    - newMask := mask | (1 << j)： j为当前未匹配的num2中的元素下标；newMask为选择j进行匹配后的状态
    - dp[mask]为之前的状态
  - 数组的最大与和：对于每一个状态 mask，枚举当前待分配的数 nums[i]，尝试将其放入任意一个合法的篮子中
  - - dp[newMask]=max(dp[newMask],dp[mask]+(nums[count]&j))
    - mask为篮子的状态数，每个篮子占两位，一共n个篮子，1 << (numSlots * 2) 
    - 计算编号为j的篮子的分配情况：(mask >> ((j - 1) * 2)) & 3  j>0
    - 标记 篮子j的某个位置已经被访问： mask + (1 << ((j - 1) * 2))
    - num[count] 标识要分配的数
    - count为已经分配的数的个数
    - j标识篮子编号
    - totalStates := 1 << (numSlots * 2) // 每个篮子有 2 位，numSlots 个篮子
#### 计数类dp 动态规划方法来统计可行方案数目的问题
- 解决这个问题有多少种方法
- 不同路径：dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
- 整数拆分：max(j×(i−j),j×dp[i−j]) 分为可以拆分和不能继续拆分两种情况
#### 数位dp 一种与数位相关的一类计数类动态规划问题
- 求解给定区间 [left,right] 中，满足特定条件的数值个数
- 求解满足特定条件的第 k 小数
- 前缀和：sum[l,r]=prefixSum[r]−prefixSum[l−1]
- 将区间数字拆分为数位，然后逐位进行确定
- 统计特殊整数：过于复杂
- - 状态： dp[pos][tight][mask]
  - tight 为 1，当前位数字只能选 0∼digits[pos]
  - tight 为 0，当前位数字可以选 0∼9
  - res += dfs(pos+1, tight && (d == limit), newMask)
- 至少有 1 位重复的数字： 逆向计算没有重复数字的个数
- - dp法： 数字转换为字符串，状态是【pos,mask,tight】的数组；使用递归
  - - pos：当前正在处理的数位位置；pos达到数字尾部则结束
    - mask：记录当前数位是否已被使用（通过二进制掩码表示）；newMask |= (1 << d) 标记新状态；
    - tight：当前是否仍然受 n 数值限制；比如123，第一位只能选择0或者1；有限制则tight为1；若tight等于1则需要计算新的limit，limit默认为9
    - tight = true 且 d == limit： tight状态传导
- - 组合法：每个数位拆解放入数组
  - - 计算没有重复的排列数量：第一位取9个值，第二位只能取8个值；第一位不能使用0；pem(m,n) 从m个数字中选出n个进行排列
    - 计算小于 length 位数的数字没有重复的总数：9×P(9,i−1)，忽略第一位，则第二位可以选择9个
    - 计算长度为 length 且不超过 n 的没有重复数字的数量：遍历数位个数 和 每位能够选择的数字；noDupCount += perm(9-i, length-i-1)
### 普通
- no 91 1->A 26->Z 由数字组成的字符串，求有多少种解码方法：dp[i]表示字符串长度为i有多少种编码法：类似楼梯，两个字符，要么1个字符：dp[i] = dp[i-1]+ dp[i-2] 两个字符有一定限制
- no 62 机器人从左上角走到右下角的总数，自能向下和向右：dp[i][j] 表示走到（i，j）有多少种走法；由于只能向下和向右，所以第0列和第0行只有一种走法：当i=0，dp[0][j] = 1 同理dp[i][0]=1
- > dp[i][j] = dp[i-1][j] + dp[i][j-1] 到（i,j) 的走法总数，等于从上边走的总数+从左边过来的总数
- no 63 同62,格子中加了障碍，遇到障碍，则dp[i][j] = 0
- no 64 同62，求所有路径中的最小路径：dp[i][j] 表示最小路径长度 ：dp[i][j] = min(dp[i][j-1],dp[i-1][j])+grid[i][j]
- no 70 爬梯子，一步或两步，求总共有多少趴法：dp[0] = 0, dp[1] = 1 dp[2] = 2; dp[i]表示到i的爬法有多少种：到达i，要么爬两阶，要么爬1阶，则dp[i] = dp[i-1] + dp[i-2]
- no 139 s是否能够完全分裂为wordDict中的单词；s不能有剩余的非dict单词:
- > 动态规划：dp[i]  = ( dp[j] && set[s[j:i]]);初始dp[0] = 0
- > 递归：set[s[0:i]] && helper(s[i:],set)
- no 140 将s分解为带空格的句子，句子中每个单词都是wordDict中的单词
- > 动态规划：dp[i] = []string{} 保存所有以i为结尾的句子；dp[i] = dp[j] + s[j:i];当j==0 或者 dp[j] != nil
- no 96 给定1-n，求独立的二分查找树有多少个？ ：G(n)=G(0)∗G(n−1)+G(1)∗(n−2)+...+G(n−1)∗G(0)
- > dp[i] += dp[j] * dp[i-j] i>=j dp[0] = 1 dp[1] = 1
- no 198 房子盗窃问题：数组表示每个房子的现金数量；相邻房子不能抢；求能盗取的最大现金数；
- > dp[i] : 盗窃i房子能够得到的最大现金数：dp[i] = max(dp[i-1],dp[i-2]+nums[i])
- > dp[0] = nums[0] ,dp[1] = max(nums[0],nums[1])
- no 213 房子盗窃II，房子形成环，条件同198：形成环的房子，首尾也有可能邻接；[0,n-1],[1,n] 风别调用198的函数
- no 337 二叉树状的房子：递归左子树和右子树，返回两个值：0：表示不包含当前节点的值:leftMax+rightMax;2.表示包含当前root.Val.(int) + right[0] + left[0]

### 正则匹配
- "*"的正确用法是：zo*  等价于z zo 或者zooo
- no 10 判定s和p是否匹配：p中可能包含("."代表任意一个字符)，"*"代表之前的0个或多个字符
- > dp[len(s)+1][len(p)] ,dp[0][0] = true;dp[0][i] && p[i]=='*' 则dp[0][i+1] = true
- > 若s[i] == p[j] && dp[i][j] ,则dp[i+1][j+1] = true
- > 若p[j] == '.' && dp[i][j] ,则dp[i+1][j+1] = true
- > 若p[j] == '*' ：
- > 1.  p.charAt(j-1) == s.charAt(i) or p.charAt(j-1) == '.' : dp[i+1][j+1] = (dp[i+1][j] || dp[i][j+1] || dp[i+1][j-1]) // （aa,a*）|| (aaa,a*) || ( a，a*)  
- > 2. p.charAt(j-1) != s.charAt(i) or p.charAt(j-1) != '.' : dp[i+1][j+1] = dp[i+1][j-1]  // (aa,aab*）
### 回文
- 5 最大回文子串 ：必要条件：(s[start] == s[end] && (end - start <= 2 || dp[start + 1][end - 1])
- > 子问题：dp[i][j]是回文需要满足什么条件？0<=i<n,i<=j<n
- > 1.s[i] == s[j] 且 dp[i+1][j-1]是回文串
- > 2.s[i] == s[j] j - i <= 2
- 516 最大回文子序列，不需要连续字符
  ```
  Input: s = "bbbab"
  Output: 4 (bbbb)
  ```
- > 子问题：dp[i][j]=k，i，j所代表的子串有k个字符可构成回文；0<=i<n,i<=j<n;k<=j-i+1
- > 1j-i == 0; dp[i][j] = 1
- > 2 j-i == 1; 且 s[i] == s[j]; dp[i][j] = 2
- > 3.j-i > 1 ,s[i] == s[j];dp[i][j] = dp[i+1][j-1] + 2
- > 4.s[i] != s[j]; dp[i][j] = max(dp[i+1][j],dp[i][j-1])
- > 5.2和3综合为：s[i] == [j],则;dp[i][j] = dp[i+1][j-1] + 2
- no 214 给定一个字符串在前面加字符构成最小的回文串
- > 暴力法：反转s为rev，s和rev在中间部分i位置一定会相等；rev[0:i]即为要补充的字符串
- no 125 判断字符串是否是回文：开始和结尾比较，不相等则返回false；去掉除数字和字母外的其他字符
- 630 去掉最多一个字符是否是回文?: 双指针，遇到不相等的i+1 或j-1 ，使用递归判断子串
- 131 字符串进行分组，每个分组都是回文，问字符串有多少种分组：核心问题是组合问题，使用回溯（dfs）算法
- > dfs：关键：传递一个起始位置，分组保存和最终结果；1.遍历字符串从起始位置开始；取子串：[起始位置：i+1]；2.每轮的剪枝条件：新的子串是否是回文；3.可以使用备忘录记录已经计算过的回文
- > dfs 和dp的结合：dp参考no 5
- 132 将字符串切割为回文，求可以最少切割多少次？
- > dfs + dp 会超时
- > dp: dp[i][j] 判断是否是回文；cut[k] 表示0->k需要的最少切割次数；1.无回文的情况下cut[k] = k(每个字符切割1次)；2.有回文的情况下：若dp[i][j]是回文，则cut[j] = cut[i-1] + 1;当i为0是，cut[j] = 0
### 基于左右的dp
- 42 数字代表墙，求墙体的凹槽可以装多少单位的水？
- 分析：当前墙体可以乘水的单位由左右两边的墙体高度决定;
- >求解i能够乘水的单位trap[i]，求和得到能够乘水的最大值
- >i的乘水量 trap[i]由左边和右边最低的决定:trap[i] = min(left[i],right[i])-height[i]
- > trap[i] <=0.则trap[i] = 0 ；否则为trap[i]
- >left[i]和right[i]求解：动态规划
- 子问题：left[i]的值如何计算？
- >max(left[i-1],height[i-1]) : 左边墙高度：由左边第一个值和左边的maxleft决定
- >right同理
- 238 构建一个数组，值是除自己以外的其他数的乘积,要求O(n)，不使用除法
- > 子问题：pro[i] = left[i-1] * right[i+1] ,分为左右两边求解，同42
-> left[i] = num[i-1] * left[i-1]

### 背包问题
- 01问题只选定的值只有1个和0个的区别，空间优化第二重循环必须逆序
- 完全问题：选定的值可以有0个和无限个，空间优化第二重必须正序
- 求最小值问题，dp需要初始化为特大值

#### 01背包：n个物品，w[i]为第i个物品的的重量，v[i]第i个物品的价值，背包承重w，要求放入的价值最大：
- > 状态转换：dp[i][j] 为前i个物品(1到i)放入重量为j的背包的最大价值：两种情况决定最大价值，i放入和不放入的最大值（i-1实际代表第i个物品）
- > 1.i不放入背包的最大价值：dp[i-1][j]
- > 2.i放入背包，则背包承重变为j-w[i]:dp[i-1][j-w[i-1]]+v[i-1];此处i-1因为物品不能重复放入
- > dp[i][j] = max(dp[i-1][j],dp[i-1][j-w[i-1]]+v[i-1]);j>=w[i-1]
- > 优化空间j必须反向遍历
  > 一维状态数组公式：dp[w] = max(dp[w], dp[w - weight[i - 1]] + value[i - 1])
- 分割等和子集：等价为子集为数组和的一半；nums[i] 可以同时为价值和重量
#### 完全背包问题：与01背包的不同在于同一个物体可以无限重复放入 （硬币问题）
- > 状态转换： dp[i][j] = max(dp[i-1][j],dp[i][j-w[i-1]]+v[i-1]);j>=w[i-1]；与01背包的不同在于dp[i][j-w[i-1]]+v[i-1])，第一维索引为i，表示i可以重复放入
- > 优化空间j必须正向遍历
- 多重背包问题：不同点在于引入了n[i] 表示第i总物品有多少个
- > 状态转换：考虑装入i物品的件数：k为装入i物品的件数（0...n[i]）= min(n[i],j/w[i])
- > dp[i][j] = max(dp[i-1][j-k*w[i]]+ k*v[i] for every k)
  > 简化后：dp[i][j] = max(dp[i-1][j], dp[i][j-weights[i-1]] + values[i-1])
  > 一维数组公式：  dp[w] = max(dp[w], dp[w - weight[i - 1]] + value[i - 1])
- no 518 amount = 5, coins = [1,2,5]，求组合数
- >子问题：dp[i][j] 为前i个coins组合成j的组合总数；两种情况：coins[i-1]为对应第i个硬币，使用和不使用第i个的组合数相加，决定了总的组合数
- > 1. 不使用第i个硬币的组合j的情况：j<coin[i-1];dp[i][j] = dp[i-1][j];
- > 2. 使用第i个硬币的情况，由于相同硬币可以重复使用（dp[i][j-coins[i-1]]第一维索引为i）：
- >  j>=coin[i-1]；dp[i][j] = dp[i-1][j] +dp[i][j-coins[i-1]] ; j-coins[i-1] 表示j已经加入了一个硬币i
- no 322 求银币组成amount值的需要的最小银币个数：dp[i][j] = min(dp[i-1][j],dp[i][j-coin[i]]+1 )；注意求最小值，dp要初始化为特别大的值
#### 多重背包问题：物品个数有限
- 选择物品个数的时候要在w/w[i-1] 和count[i-1]中选择更小的哪个
- dp[j] = max(dp[j], dp[j - k * weights[i]] + k * values[i])  (0 ≤ k ≤ num[i], 且 j ≥ k * weights[i])
- dp[i][j] = max(dp[i][j], dp[i-1][j - k * weights[i-1]] + k * values[i-1])   (j ≥ k * weights[i-1])
- 可通过二进制优化转换为0-1 背包问题
#### 混合背包问题：
- 组合0-1 和完全背包问题；多重背包问题转换为0-1背包问题
#### 分组背包问题
#### 二维费用背包问题
#### 背包问题变种：
- 恰好装满背包的最大价值：dp[w]的初始状态1<w<n的情况下为负无穷没有合法解
- 求方案总数：背包问题是求最大值，现在是求和
- 求最优方案：结合了最大价值和总方案数
- 求具体方案：

### 其他
- no 1524 获取数组中所有子数组中和为奇数的子数组的个数 (子数组元素必须连续)；递归法超时
- > dp0[i] 表示以[i:n]子数组的和为偶数的子数组长度；dp1[i] 相反
- > 状态转换：
- > 1.第i个元素为奇数，则dp0[i] = dp1[i+1]（奇数数组尽数转换为偶数） dp1[i] = dp0[i+1] + 1(1为第i个元素本身)
- > 2.第i个元素为偶数，则dp1[i] = dp0[i+1] dp0[i] = dp1[i+1] + 1(1为第i个元素本身)
- > 3.结果为dp1[i]的和
- 416 是否能够将集合分成两个组，每组的和相等？ 所有值都是正数;：数组整体和为奇数的无法分割直接返回false;问题转化为：是否有任意个数字的和等于sum/2
- > 动态规划： 类似 518 银币组合问题：dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]);其中dp[i][0] = true 0<=i<n;dp[0][j]=false; 1<=j<n
- > 递归：求解sum/2的组合，分为包含num[i] 和不包含num[i]两种情况


## 分治策略
- no 395  字符串子串中包含最少k个重复字符的最长子串长度；子串中的所有字符的重复字符必须超过k
- >1.分治策略： 统计字符串中所有的字符个数；若个数小于k则，分别计算两侧的字符串
- >2.滑动窗口：

## 贪心策略
### 跳跃游戏：一个数组，每个值表示当前能够跳跃的最大步数，一般问题：能不能跳到最后？跳到最后需要的步数（最小，最大）
- 考虑从目的地开始，从后向前
- 贪心策略总是假设最终结果成立
- 贪心策略不是所有问题都有最优解，大多数问题都能得到整体最优解
- 自顶向下的求解过程
- no 55：问是否能跳到最后
- > 解法一：从0开始找打第一个能跳到最后的索引（num[i] >= end -i）,然后更新end值为i；若未找到则退出循环
```
for i := 0; i < positon; i++ { //一个循环里负责遍历i的同时动态更新positon，缩小循环的范围
	if nums[i] >= (positon - i) {
		positon = i
		isUpdate = true
		break
	}
}
```
- > 解法二：动态规划：dp[i] 表示能到达最后；从后往前;初始化：dp[n-1]=true; dp[i] = dp[i+k] && i+k<n（从后向前遍历nums）；0<=k<=nums[i](一重循环遍历nump[i]的值) ; 结论：dp[0] == true
- > 解法三：贪心策略：找到每个位置能够跳到的最远的位置；若最远的位置比当前的位置还小(这个条件不好找)，则返回false；farest = max(farest,nums[i]+i)
- no 45 求最少的跳数
-> 解法一：从后往前，每次找到最远的位置
```
for positon != 0 {
	for i := 0; i < positon; i++ {
		if nums[i] >= (positon - i) { //i 从0开始，保证每次找到了跳的最远的位置
			positon = i
			steps++
			break
		}
	}
}
```
-> 解法二： 动态规划：dp[i]为i到达末尾的最小跳数;dp[i] = min(dp[i+k]+1,dp[i]) i+k<len(nums) && 0=<k<=nums[i]；结论：dp[0]
-> 解法三：贪心算法：求每步的能够跳的最远距离；												  
- no 630 课程调度：优先级队列；优先截止时间近的和优先替换课程时间最长的课程，都体现了贪心策略
- > 贪心策略：1.首先以截止时间排序，防止替换时需要考虑截止时间；2.遍历课程，课程时间入队，并累加时间线；若时间线大于当前截止时间，则出队课程时间最长的课程；
- no 134 gas存的时每个加油站的油量；cost存的是从i->i+1需要消耗的油量；路是环路；问能不能环绕一圈,返回起点
- > 暴力法：计算gas[i]-cost[i] >= 0 的起点;遍历所有合适的起点，从起点开始走：每次(路径加1）%n,并计算邮箱的剩余油量，小于0或者回到起点则退出；判断是否可以抵达一圈
- > 总结算法：1.若总油量大于总cost，总会有一个解，否则没有解；2.A无法到达B，则A和B中间的都无法到达；可以直接跳到B作为新的起点
