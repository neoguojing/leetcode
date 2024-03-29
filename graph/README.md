## 图
- 邻接表：适合稀疏图，求解入度和度不方便，空间复杂度o(v+e)或o(v+2e)：1.需要点集合（点的数组）,每个点保存一个列表，存储以该点为起点的所有边；2.边结构存储边的末尾节点，以及下一条边的指针
- 邻接矩阵：适合稠密矩阵；空间负责度o(v^2)
### 广度优先
### 最短路径算法（动态规划）（bellman-ford）：边的权重可以为负数，n为节点个数 时间复杂度O（VE）
- 输入：边集合：[i,j,w];点集合；距离集合Distant；起始点
- Distant[i]:为源点到i点的最短距离；初始时设置 Distant[原点] = 0，其他的为无穷大
- 松弛计算：对所有边，若Distant[j] > Distant[i] + w[i,j],则Distant[j]= Distant[i] + w[i,j]；i为边的起点，j为边的终点
- 计算最短路径：进行n-1轮的松弛计算；第一轮更新原点相隔一条表的节点；第二轮两条边可达的距离；n-1轮正好覆盖最差情况（n个节点串联）；若某轮没有更新Distant，则结束
- 负环路：权值之和为负数的环路；存在则无法求出最短路径；遍历所有边，若依然存在Distant[j] > Distant[i] + w[i,j]，则有环路，无法求解最短路径
```
dist := make([]int, n)//到原点的距离
edges [][]int{起始点，终点，权重} //边集合，邻接矩阵
	// n-1轮循环
for i := 0; i < n-1; i++ {
	    check := false
		   // 遍历所有边，进行松弛
	for _, v := range edges {
		if dist[v[1]] > dist[v[0]]+v[2] {
			dist[v[1]] = dist[v[0]] + v[2]
			check = true
		}
	}

	if !check {
		break
	}
}	
```
- no 787 找到最多K次中转的最便宜的航班路径
- > 状态转移：dp[k][j] 为需要i次中转的终点为j的价格；0<=k<=K+2;0<=j<=n
- > dp[k][j] = min(dp[k][j],dp[k-1][i]+w[i,j])
### Dijkstra（贪心策略） ： 处理单源最短路径，边的权重不能为负数，适用于有向图和无向图，图可以包含环路 o（v^2）
- 思路：每次选取未加入集合的，离原点最近的点；加入集合，并以该点为起点，更新到原点的距离
- 输入：原点v；顶点集合；边集合包含点和权重[i,j,w]；dist[n] 原点到节点的距离；prev[]记录当前点的前一个节点；已计算的的点集合s[]
- 初始化：初始化s[v] =1,其他为0；初始化prev和dist
- 计算：遍历节点，找出未放入s的节点中，找到dist最小的，将节点放入s； 以新加入的点为起始，更新dist和prev
- 通用模式
```
  dist := make([]int, n) //保存到原点距离
	path := make([]int, n) //保存到原点path，可选
	set := make(map[int]bool) // 判定点是否被选定
  adj := make([][]int, n) //邻接矩阵，值为权重
  // 初始化
  adj[src]可发现的边初始化
  set[src]初始化
  dist和path初始化
  // 计算
  for i := 0; i < n; i++ {
       min := INF //最小值保存
		   selectNode := 0 //选中节点保存
      //寻找节点              
      for j := 0; j < n; j++ {
        if !set[j] && dist[j] < min {
				min = dist[j]
				selectNode = j
			  }
      }
      // 保存节点                             
      set[selectNode] = true
      //从新选节点出发，更新        dist和path                     
      for j := 0; j < n; j++ {
        if !set[j] && min+adj[selectNode][j] < dist[j] {
          dist[j] = min + adj[selectNode][j]
          path[j] = selectNode
			  }
      }

  }
  
```
### Floyd：任意两个节点间的最短路径 o(v^3) 在任何图中使用，包括有向图、带负权边的图;邻接矩阵来储存边
- 思路： 对于每一对顶点 u 和 v，看看是否存在一个顶点 w 使得从 u 到 w 再到 v 比己知的路径更短
 ```
  For k←1 to n do // k为“媒介节点”
   For i←1 to n do
      For j←1 to n do
         if (dist(i,k) + dist(k,j) < dist(i,j)) then // 是否是更短的路径？
            dist(i,j) = dist(i,k) + dist(k,j)
 ```
### 拓扑排序 适用于DAG 有向无环图
- Kahn算法：优先将入度为0的节点从图上删除，排列，直到所有节点均排列完成；需要所有点的入度和一个队列；每次出对计数一次；若最终计数不等于总的节点数，则有环
```
q := NewQueue()
cnt := 0
//入度为0，入队
for i, in := range ins {
	if in == 0 {
		q.Push(i)
	}
}
<!-- 没有则有环 -->
if q.Empty() {
	return ret
}

for !q.Empty() {
	point := q.Pop()
	//cnt，统计访问次数
	cnt++
	ret = append(ret, point.(int))
	for _, edge := range prerequisites {
		if edge[1] == point {
			ins[edge[0]]--
			if ins[edge[0]] == 0 {
				q.Push(edge[0])
			}
		}
	}
}
<!-- 节点数和访问次数不相等，有环  -->
if cnt != numCourses {
	return []int{}
}
return ret
```
- no 329 求矩阵中的最长递增路径长度：因为递增不存在环路，DAG；使用拓扑排序算法，入队的层数即路径长度；
- DFS算法： 优先访问出度为0的节点
- > 1.建立数据的邻接表；2.建立已访问节点集合；3.遍历所有节点，调用dfs；4.dfs优先判断节点状态：0未访问，1 正在访问（该状态有环，直接返回false），2该节点的所有边都被访问（返回true）；5.dfs中遍历该节点的所有边，递归调用dfs；6.dfs中设置节点状态为2，并将当前节点放入结果集；7.结果集的逆向为顺序
```
主函数
res := []int{}
for i := 0; i < Vnum; i++ {
	if !dfs(graph, visited, i, &res) {
		return []int{}
	}
}
result := make([]int, 0)
for i := numCourses - 1; i >= 0; i-- {
	result = append(result, res[i])
}
	
dfs
func dfs(graph, visited []int, cur int, res *[]int) bool {
	if visited[cur] == 1 {
		return false
	} else if visited[cur] == 2 {
		return true
	} else {
		visited[cur] = 1
	}

	for i := 0; i < len(graph[cur]); i++ {
		if !dfs(graph, visited, graph[cur][i], res) {
			return false
		}
	}

	*res = append(*res, course)
	visited[cur] = 2
	return true
}
```
- no 207 210
- BFS算法 ： 在求解某些最小路径时会有效
```
q := NewQueue()
q.Push(begin) //源如队
for !q.Empty() { //非空
	qsize := q.Len() //获取当前队列长度
	for i := 0; i < qsize; i++ {  //遍历当前层，将当前层所有元素出队
		p := q.Pop().(string) 
		if 条件判断 {
			for i := 0; i < n; i++ {//遍历源数据集

			}
		}

		
	}
}					
``` 
- no 127 LadderLength word ladder 即有一个字符不一样的单词，从beginword开始，从wordList选仅一个字符不一样的word，直到endWord命中
