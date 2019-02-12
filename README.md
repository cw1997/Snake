# Snake
一個golang開發的命令行貪食蛇小遊戲

# author 
昌维(cw1997)

# usage 
命令行执行 

	go run ./main.go

# 数据结构
- Position 坐标，保存x y二维坐标
- Snake 蛇单位，保存蛇头的朝向，用于在每次update操作的时候确定蛇的前进方向；用list双向链表保存蛇身，方便调用remove last node操作从尾部擦除已经往前移动的部分

# 全局变量
- world全局变量负责存储整个游戏地图（贪食蛇游戏的地图为一个矩形二维数组），值为当前所在坐标上的元素。由于golang目前还不支持enum枚举，因此此处使用字符串存储，food代表食物，snake代表蛇，其他值代表空。
- empty代表所有地图上的空位，用于随机产生食物使用。定义该数组是为了提升效率，防止每次随机生成食物的时候都要搜索整个world二维数组。该数组里面的值在render渲染的时候进行append。类型为Position，后面会讲整个类型是什么。
- snk为Snake的实例化对象（go里面应该叫做struct）
- foodPosition为food当前坐标，用于碰撞侦测。

# 核心思路
render负责将数据模型上的world游戏地图渲染到console上

input负责接收输入，根据输入的wsad上下左右动态修改蛇头的朝向（forward），由于不能让输出操作阻塞游戏的渲染线程和数据模型刷新线程，因此要在goroutine里面跑input线程

update负责游戏每一帧的数据更新，例如蛇的坐标移动，蛇头和food的碰撞侦测，碰撞食物后蛇头append一个node，蛇尾不动，代表整个蛇身长度增加1，并且generateFood调用后重新在随机位置生成一个新的food，然后原先的空位标记为空。碰撞墙壁（大于地图尺寸或者小于0）直接打印gameover然后exit主线程

# 现有不足
- 未检测蛇吃自己。解决方案：在update里面增加一下蛇与自身的list结构里面所有坐标的碰撞检测代码
- 输入有少许延迟。解决方案：可以适当减少input线程的tickrate时间，减少到不影响render和update，但是比render的tickrate小。
- 未打印墙壁。解决方案：对所有数组下表都要修改。

