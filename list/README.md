## 链表
> 修改链表结构，需要保存prev指针
- no 328 将列表的第奇数个节点放在前面偶数放在后面：三个指针：1.偶数节点的头；2.奇数节点的遍历指针；3.偶数节点的遍历指针； 注意判断条件
- no 25 以k为一组反转链表，不够n则不反转：k=0或1 则不需要反转：反转列表函数，返回开头和结尾指针；使用递归，以k为分组进行反转；
### 快慢指针
- > p为快指针，循环结束条件（p != nil）；所有指针初始化指向head（head包含有效值）；cnt初始为1
- > k为倒数的节点个数：1.若需要找到倒数第k个元素，则当cnt>k时，慢指针开始移动；2.若需要找到倒数第k+1个元素，则cnt > k+1,慢指针开始移动
- > 边界条件:cnt==k+1时，正好有k个元素；cnt<k+1 则不够k个元素
- > 寻找中间节点：结束条件为fast!=nil && fast.next != nil,此时找到的是单数列表的对称中心，双数列表的两个中心节点的右边（第二个）节点even != nil && even.Next != nil
- no 19移除列表倒数第n个元素
- > 如何快速找到倒数第n个元素？移除倒数第n个元素，要找到倒数第n+1个元素
- > cnt > k+1,慢指针开始移动，cnt==k+1，则移除开头元素，cnt<k+1,则不够k个元素
- no 1721 交换正数和倒数第k个元素的值 ：  cnt > k 慢指针开始移动；cnt<k+1则元素

### 链表合并					       
- no 21 双列表合并 ：
- > 1. 递归法：传入3个参数：l1，l2，和新的头节点；结束条件：某个列表为nil，返回；递归结构 v1 < v2,则挂到新列表，merge(l1.Next,l2,head),否则
- > 2. 遍历法：先循环a列表，直到找到比b小的，修改a的指针指向b，然后遍历b：
- > 以 a != nil 作为循环结束条件；需要分别保存前序节点；
```
<!-- 编码细节	 -->
for a != nil && b != nil {
	for a != nil && b != nil && a.Val <= b.Val {
		prevA = a
		a = a.Next
	}
	if prevA != nil {
		prevA.Next = b
		prevA = nil
	}
	for b != nil && a != nil && b.Val.(int) <= a.Val.(int) {
		prevB = b
		b = b.Next
	}
	if prevB != nil {
		prevB.Next = a
		prevB = nil
	}
}

```
- no 23 k个列表合并 ： 每两个列表合并，合并完之后放入一个数组，递归调用					     
### 链表copy
- no 138 有随机指针的列表深度copy：问题的关键在于随机指针是跳跃的，无法在重建列表时找到目标指针
- > map建立旧表指针到新表指针的映射；第一次循环建立新列表同时建立map映射；第二次循环：根据random指针通过map找到新表random指针，为新表的节点赋值random指针
	