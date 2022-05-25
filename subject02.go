package main

/**
根据活动参与人数 两种设计
	场景01(参加人数小于1w) 排行数据不入库，使用redis的zset数据结构实现，活动结束之后根据比赛分数计算逻辑将排行信息入redis
	优势：
		1、减少了对数据库的操作。
		2、基于redis的单线程，避免了变更排名信息的加锁
	缺点：
		1、参与人数过多，导致的存储过大，效率降低。
		2、数据安全性低，redis数据过期丢失后，需要重新计算。
	场景02(参与人数大于1w) 基于 redis + mysql 分数计算后入库，前几百名数据缓存到redis中
		查询：
			系统提供玩家名次查询接⼝，玩家能够查询⾃⼰名次前后10位玩家的分数和名次。 这个查询操作暂时没有想到很好的方法

*/

//数据结构
type score struct {
	id          int    //自增id
	uid         int    // uid
	vitality    int    //排序使用的value
	params      int    //分数的第二排序规则
	source      string //活动场次区分
	note        string //一些备注扩展信息
	add_time    int
	update_time int
}