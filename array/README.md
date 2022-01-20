## 数组
- no 31 给定一个数列，返回正好比该数列大的一个数：1.从后向前找到第一个递减的位置；2.从后向前找到比递减位置大的元素，交换；3.从递减+1位置，反转数组
- no 27 原地移除指定的元素：从后往前遍历，遇到相同的则和最后位置交换；最后位置用一个变量标记
- no 169 找到数组中个数超过一半的元素 o(1) 空间
- > 1. map记录 2. 排序然后找到中间元素 3.随机：从数组随机取一个数，遍历数组判断出现的次数 4.摩尔投票算法：需要候选者和计数器；计数器为0则更换候选者；候选者和当前元素一样，则计数器增加，否则计数器减少
- no 217 判断数组是否有重复值：1.暴力循环；2.排序然后遍历；3.hashmap
- 283 原地将0移动到末尾，并不改变元素相对顺序： 双指针：慢指针记录0的位置（没有0则保持和快指针一致）；块指针遇到非0的值与慢指针交换
- 128 求数组中的最长连续子序列，要求o(n):
- > 1.并查集：建立集合保存值和父亲；遍历数组：若存在num[i]-1的元素，则设置为num[i]父亲；遍历数组进行路径压缩，同一集合的父亲为同一个值（最小值）；返回最大集合的个数
- > 2.借助集合：若集合中找不到比当前小1的，则作为一个新的子串起点；循环遍历比当前大1的值是否能够找到，更新长度等。
- 150 返回数组中子数组中乘积最大的值：记录：当前最大值，当前最小值，和当前值比较，取最大和最小更新当前最大和最小值；最大的当前最大值则为返回结果
- > 1.当所有数位正数时，则所有数相乘位最大值，即当前最大值；2.当包含负数时，最大值可能变位最小值，同理，此处为记录最小值的意义；3.当包含0时，最大和最小均为0；当前元素即为最大值。
- 189 从右边开始旋转k次数组
- > 解法1：按题要求旋转k次，外面循环控制次数，内部循环控制数组整体位移1（从n-2 到 0） o(kn) 空间o（1）
- > 解法2：首先将数组整体倒排（左右交换），然后对[0,k-1]和[k:n-1]分别倒排  o（2n） o(1)
- > 解法3: copy整个数组，在原数组上操作
- 215 返回第K大的元素：1.大顶堆；2.排序取第K个；3.快速选择：只借用快排的partition，无需递归；对于选择最大的k个，调整上下边界，使得分解点正好位于倒数第K个
- 300 求数组的最长增长序列长度：
- > 1.耐心排序：一次遍历数组；若没有桶则建桶；若有桶，从左往右遍历，找打第一个大于等于当前元素的桶，将数据放入桶中；未找打则新建桶；桶的个数即最长增长子序列的长度；
- > 2.dp[i] 表示以i为结尾的最长子串；初始dp都为1；dp[i] = dp[j] +1 当且仅当nums[i]>nums[j] 且dp[i] < dp[j] +1 0<j<i
- > 3.	
- 334 求数组中是否有3个的递增子序列：同300 ：1.动态规划超时；2.耐心排序：通过；3.small 和big初始为最大整数；small记录最小值；big记录中间值；找到n大于small和big，则返回true

### TOPK问题： heap（NlogN） 或者快速选择算法
- 347 数组中出现的最频繁的K个数字：1.堆排序：NlogN；2.快速选择：最好o(N),最坏o(N^2)
### 数组合并
- 有序数组合并：1.找到两个数组的最大或最小位置，依次比较然后选择一个填入目标位置 o（n）
- 插入新的集合：1.边界处理：插入最前和最后的直接处理；2.获取插入位置；3.将前一个位置与待插入元素计算集合入队；，一次遍历剩下集合，右交集则更新队尾；无则入队
- 集合删除：1.处理边界条件：删除开头和结尾的情况；2.查询删除的起始位置；3.从起始位置开始，比较计算：4中情况：1.没有交集；2.交集导致集合分裂为两个；3.删除集合左边部分，4删除集合右边部分
- 集合查询：1.边界情况处理；2.遍历查询起始位置，查到则跳出循环，判断该集合是否包含要查询集合
- no 88 已经有序的2个数组原地合并：关键的原地合并：必须倒序合并，则不需要交换元素
- no 977 有序数组每个元素平方后输出到排序数组：取正数组和负数组，转换为两个有序数组合并；必须借用新的空间
- no 715 集合的删除、插入和查询
### 多数组相加问题
- no 454 四个数组相加，返回组合中和等于0 的组合个数；数组长度相等：可转换为o（n^2），利用hashmap
- > 数组两两相加；前一组数组的和存在map里，并记录和的个数；后一组相加：等于0则从map中取0值；否则从map取-sum的值；
- no 1 数组中任意两数相加等于target，返回对应数字的索引：使用map记录数组的所有值；遍历数组，在map中寻找target-num[i]；记录索引然后返回
- no 15 返回数组中三个数相加等于0的所有组合，同一位置的值不能重复使用；
- > 1.排序数组，可以直接筛选不符合条件的；比如最后一个数小于0，则无解
- > 2.遍历数组，计算差值；使用双指针（0，n），遍历数组，计算和是否等于差；
- > 3.需要考虑重复解的问题；
- 560 和为k的子数组个数：特殊思路：求出前缀和数组，利用sum[j]-sum[i]得到i+1到j的和等于k则计数加1；变种：sum[j] - k = sum[i] 将sum[i] 存入map
- 494 在数字前面放置+或者-号，让所有数字的和等于target；返回可能的组合个数：使用回溯算法;
### 集合
- 后面结果仅依赖前面结果或者后面的计算会改变前面的计算结果的-单向依赖算法模式：包含：1.后面的计算结果和前面无关，则向队列末尾插入；2.后面的计算依赖前面计算结果，则更新结果
```
ret := [][]int{arr[0]} //首元素入队
for i := 1; i < len(arr); i++ { //从1开始遍历元数据
	if 后面不依赖前面计算结果 { 
		ret = append(ret, intervals[i]) //新元素入队
	} else { //依赖前面计算结果，则更新队列尾部
		ret[len(ret)-1] = 新的值
	}
}
```
- no 986有序集合合并（取交集）：从最小或最大端开始遍历两个集合数组；不相交则跨过排序靠前的集合；相交：则取交集，跨过上边界较小的集合
- no 56 将数组中有交集的集合合并:1.首先按start坐标排序，将问题转化为前后依赖关系；2.后面的计算结果依赖前面的计算结果，利用队列计算	
- no 57 在有序集合的数组中插入一个集合（有交集则合并）：1.找到插入位置（位置+1为要插入的元素）；2.因为有序，所以采用通用算法求解；需要保证起始时至少有一个元素 ；3.特殊情况处理插入位置小于0的情况
- no 495 反复放毒，求敌人的总中毒时间: 转换成集合，按照标准解法解题
- no 1386 每排10个座位，共n排，求解能分配多少个4人连坐；：集合思想：1.针对已经占用作为集合，计算每排的有效集合占用情况；2.根据集合占用情况，计算实际可分配的座位数量
- no 735 小行星碰撞：1.不可改变顺序；2.负数表示向左，正数表示向右；3.只有左边是正数，右边是负数的情况才会碰撞，其他情况下均不会碰撞；4.碰撞情况分析：a.左边绝对值大，什么都不做；b.绝对值相等，则需要删除左边最后一个元素；c:左边绝对值小，则需要从后向前遍历找到第一个（绝对值大的）或（相等的）或（值小于0的）或（所有的都小于）

- no 350 求两个数组的交集：使用hashmap
			 
### 二分查找：数组必须有序
- 模式
```
// 闭区间搜索
初始值： left, right := 0, len(nums)-1
循环条件： for left <= right  //表示是闭区间搜索[a,b]              
中值： mid := left + (right-left)/2
条件1 ： nums[mid] < target 则left = mid + 1                        
条件2：nums[mid] > target 则right = mid - 1                         right = mid
条件3：nums[mid] == target ： 寻找左值：则压迫右值right = mid - 1，寻找右值则压迫左值：left = mid + 1
结束条件：直接返回-1
未找到左值：left >= len(nums) || nums[left] != target ；
未找到右值：right < 0 || nums[right] != target
//半开区间搜索
初始值： left, right := 0, len(nums)
循环条件：  for left < right  表示半开区间搜索[a,b)
中值： mid := left + (right-left)/2
条件1 ： nums[mid] < target 则left = mid + 1                        
条件2：nums[mid] > target 则right = mid
条件3：nums[mid] == target ： 寻找左值：则压迫右值right = mid，寻找右值则压迫左值：left = mid + 1
结束条件： 需要判断nums[left] == target ? left;-1;
未找到左值：if (left == nums.length) return -1; 
未找到右值：if (left == 0) return -1; return nums[left-1] == target ? (left-1) : -1;
```
	 
- no 704 34 35
- 162 找到一个峰值的索引，峰值比相邻元素大的值，要求o(logn) 
- > 寻找峰值有三种情况，1.递增序列，则最右值为结果；2.递减序列相反；3.序列存在多个峰值
```
l, r := 0, len(nums)-1
for l < r { //采用小于，表面至少两个元素
	mid := l + (r-l)/2
	if nums[mid] > nums[mid+1] { //值和数组本身比较，则最少需要两个元素，否则会越界
		r = mid
	} else {
		l = mid + 1
	}

}
return l
```
#### 有序的旋转数组查找 o(logn)
- 旋转数组特性： 
- > 1.任意取中间值，则至少右一半的区间是递增区间
- > 2.均可采用二分查重处理
- > 3.使用中间值和最右（最左）比较，是算法的核心
```
//最小值算法：中间值与最右值比较								
lo, hi, mid := 0, len(nums)-1, 0
for lo <= hi {
	mid = lo + (hi-lo)/2 
	if nums[mid] > nums[hi] { //中间值大于最右值，则右边非递增区间，最小值落在右边区域
		lo = mid + 1
	} else if nums[mid] < nums[hi] { //中间值小于最右值，则最小值值一定在左边，且有可能就是中间值
		hi = mid
	} else { //对于中间值等于最右值的情况是因为数组中有重复元素；此时不确定在哪边，则最右索引减一
		hi--
	}
}
return nums[lo]	
//查询目标值算法，无重复值
//首先上面算法查到最小值的索引，即偏移，将算法转换为二分查找				      
rot := lo //等价于一个偏移，即旋转的次数				    
lo, hi = 0, len(nums)-1
for lo <= hi {
	mid := lo + (hi-lo)/2
	realMid := (mid + rot) % len(nums)
	if nums[realMid] == target {
		return realMid
	} else if nums[realMid] < target {
		lo = mid + 1
	} else {
		hi = mid - 1
	}
}

// 查询目标值算法：标准解法，适用于重复和不重复的
for lo <= hi {
	mid := lo + (hi-lo)/2
	if nums[mid] == target { //遇到相等则返回
		return true
	}

	if nums[mid] == nums[lo] && nums[mid] == nums[hi] { //处理重复值的场景：中间值等于左右的值，则不确定在哪个区间
		hi--
		lo++
	} else if nums[mid] <= nums[hi] { //中间值小于最右值，则表示，mid-> hi是一个递增区间，适合二分查找

		if nums[mid] < target && nums[hi] >= target { //目标落在递增区间
			lo = mid + 1
		} else {                         //不在递增区间，则在另一半
			hi = mid - 1
		}
	} else { //右边不是递增区间，则左边是递增区间
		if nums[mid] > target && nums[lo] <= target { //目标在递增区间
			hi = mid - 1
		} else {   //目标不在递增区间
			lo = mid + 1
		}
	}

}					
```
- no 33 数组无重复元素
- no 81 数组有重复元素
- no 153 在旋转数组中查找最小值，无重复元素，要求o（logn），
- no 154 在旋转数组中查找最小值，有重复元素，要求o（logn）
#### 无序数组，找丢失的整数问题或重复值问题 要求o（n）o(1)
- 利用数组的位置做最小值计数
- 取负数法
- 亦或法							 
- no 41 在无序数组中，找到丢失的最小正整数: 数组大小为n，可以保存1+n的整数；两种情况：1-n的整数丢失可通过整数归位（num[num[i]-1] = num[i]）判断；1-n整数都在，则丢失的是最大数（n+1）
- > 将nums[i]=4 放到数组第4-1的位置；循环直到所有数组都正确放置；结论：找到第一个nums[i]!=i+1的值；循环里嵌套循环（第二层循环并非全量循环，算做o（1）
- no 268 給定n个数字（0-n），找到丢失的数字：1.将数字放入对应的索引位置；未放置的位置即未丢失的数字；2.ret^i^nums[i]
- no 448 给定1-n的数字，有重复，返回丢失的k个数字：1.将数字放入对应索引的位置；标记重复值的位置为-1，统计-1的位置为结果；2.将nums[i]值对应位置的值置为负数，最后统计为正数的位置
- 287 找到数组中重复的元素  注意条件1 <= nums[i] <= n，数组中只有一个重复值
- > 算法1：快慢指针求和问题：因为条件1 <= nums[i] <= n，数组长度为n+1，因此数组不会越界；可以把数组的值当作索引，构建列表；首先快慢指针获取相遇点；然后慢指针从0开始，快指针从相遇点开始，求解相遇点；相遇点即重复值；
- > 2 排序然后遍历
- > 3.map 统计值出现的次数
- > 4.类似41，将整数归位，并*-1，转为负数，重复的值则会转换为正数
- > 5 0的位置不会被映射值；考虑不断的交换0-i,将0位置的值放置到正确的位置；直到遇到重复
### 双索引法
- 345 数组部分反转：操作快排，首尾同时遍历；条件满足则交换；否则需要继续移动坐标
- 11 数组区间，求哪两个柱子之间可以灌最多的水？
- > 面积最大的灌水最多，遍历所有的柱子组合：面积计算：（j-i）* min(a[i],a[j])
- > 计算第一个和最后一个的面积。
- > 如何得到更大的面积？提高最短边的长度。从两侧分别遍历，交替提高短边的高度