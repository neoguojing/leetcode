## 排序
### 排序总结
- 冒泡排序： 内正外正，  n-n^2     稳定
- 选择： 内正外正， n^2  不稳定
- 插入：内逆外正， n-n^2  稳定
- 希尔排序： 三层循环，同插入  nlogn-n^2   不稳定
- 归并排序： 左右排序，然后merge， 左右比较加入排序   nlogn  稳定  空间 n
- - 归并排序，每轮合并，左右子数组都是有序的
  - 对于问题315，r[j] < l[i], 则r[j] 左侧的数都比l[i]小，记个数为rc
  - 同理，由于左侧元素是有序的，l[i] < l[i+1], 则有l[i+1]>r[j] ,右侧至少有rc个数比l[i+1]小
- 快速排序： 先分，然后左右排序，左右比较交换， nlogn-n^2 不稳定 空间 n
- 对排序： 先建堆，然后交换首尾，调整堆；建堆从len/2开始顺序调整堆；逆序遍历执行交换首尾  nlogn 不稳定
- 计数排序： 计算数据范围，构建count数组，对每个数据计数，生成累计数组；倒序填充新数组，count数组获取位置，原数组获取值； n+k， 空间k
- 桶排序： 依据桶的个数创建二位数组，作为桶；每个桶排序然后放入最终数组    n 空间 n+m    关键： num-min将索引归零
- 基数排序：按照数据位排序；计算数据最大位数；遍历最大位数，未每一位建立桶；桶为二维数组，一维长度为10（0-9）；将数据放入桶；遍历桶，塞入原数组； n*k  n+k
### 列表排序 一般使用归并排序 o(nlogn) o(1)
- 列表获取中值：添加一个虚拟头，两个指针，一个移动2步，一个移动一步；条件是：快指针不等于nil和next不能nil； 注意，最终需要拿到的是中值的前值，然后需要断掉中值直接的链接
- 归并：添加一个虚拟头，遍历依次顺着虚拟头插入；需要一个指针记录当前位置
-  no 128 对一个列表进行排序：归并排序，获取中止，递归，然后合并
- no 147 列表的插入排序：需要一个虚拟头；一个p指向当前位置；注意：每轮循环结束，已旋转列表的头需要重置
### 有限元素类型的排序问题
- no 75 三种颜色，0，1，2，将相同颜色连在一起，且按照0，1，2的顺序排序
- > 因为值的个数是3种，我们只关心0和2（首尾），排好序后，中间的就是1
- > 算法实现： 两个变量，记录当前最后一个0和最前一个2的位置；遍历数组，遇到2交换到后面；重新判断当前位置为1则继续，为0则交换到前面的位置继续
