# leetcode

## 回溯法：递归结束条件，剪枝约束条件和
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
## 动态规划
- 求走法，种类等问题：如91，62 70等，到达某个解的种类是固定的比如1步还是两步，比如向下还是向右，比如一个字母还是两个字母，通常类型dp[i] = dp[i-1]+dp[i-2]
- 求最短，求最长等，通常涉及到max和min的求解
- 数组范围：需要分左和右dp两部分求解

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
- > 状态转换：dp[i][j] 为前i个物品放入重量为j的背包的最大价值：两种情况决定最大价值，i放入和不放入的最大值（i-1实际代表第i个物品）
- > 1.i不放入背包的最大价值：dp[i-1][j]
- > 2.i放入背包，则背包承重变为j-w[i]:dp[i-1][j-w[i-1]]+v[i-1];此处i-1因为物品不能重复放入
- > dp[i][j] = max(dp[i-1][j],dp[i-1][j-w[i-1]]+v[i-1]);j>=w[i-1]
- > 优化空间j必须反向遍历
#### 完全背包问题：与01背包的不同在于同一个物体可以无限重复放入 （硬币问题）
- > 状态转换： dp[i][j] = max(dp[i-1][j],dp[i][j-w[i-1]]+v[i-1]);j>=w[i-1]；与01背包的不同在于dp[i][j-w[i-1]]+v[i-1])，第一维索引为i，表示i可以重复放入
- > 优化空间j必须正向遍历
- 多重背包问题：不同点在于引入了n[i] 表示第i总物品有多少个
- > 状态转换：考虑装入i物品的件数：k为装入i物品的件数（0...n[i]）= min(n[i],j/w[i])
- > dp[i][j] = max(dp[i-1][j-k*w[i]]+ k*v[i] for every k)
- no 518 amount = 5, coins = [1,2,5]，求组合数
- >子问题：dp[i][j] 为前i个coins组合成j的组合总数；两种情况：coins[i-1]为对应第i个硬币，使用和不使用第i个的组合数相加，决定了总的组合数
- > 1. 不使用第i个硬币的组合j的情况：j<coin[i-1];dp[i][j] = dp[i-1][j];
- > 2. 使用第i个硬币的情况，由于相同硬币可以重复使用（dp[i][j-coins[i-1]]第一维索引为i）：
- >  j>=coin[i-1]；dp[i][j] = dp[i-1][j] +dp[i][j-coins[i-1]] ; j-coins[i-1] 表示j已经加入了一个硬币i
- no 322 求银币组成amount值的需要的最小银币个数：dp[i][j] = min(dp[i-1][j],dp[i][j-coin[i]]+1 )；注意求最小值，dp要初始化为特别大的值
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
