学习笔记

### 递归模板

1. terminator
2. process
3. drill down
4. reverse the state

### 分治模板

1. terminator

2. prepare data

3. conquer subproblem

4. process and generate the final result

5. reverse if needed

   我觉得这里是和递归一样的

## 感触



- 人脑的递归非常低效，很累
- 找到最近最简单方法，将其拆分可重复解决的问题
- 数学归纳法思维（抵制人肉递归的诱惑）





==本质：寻找重复性->计算机指令集==

画递归树





## 动态规划

理解为动态递推



DP 和 分治是有内在联系的



关键点：

- 本质上和递归和分治没有根本上的区别（关键看有无最优的子结构）

- 共性：找到重复子问题
- 差异性：最优子结构，中途可以淘汰次优解