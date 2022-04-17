按值传递
所见即所得

CPU核心分配逻辑处理器 得到线程
goroutine 执行路径

数据段 栈 堆

内存帧 frame 
goroutine 只能对内存进行直接访问 只能对帧进行操作 沙盒（隔离）

机制与语义
如何运行与实际行为

go 正确性 和 缩减程序运行时的开销

escape analysis

