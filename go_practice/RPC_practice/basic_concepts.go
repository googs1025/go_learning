package main

/*
	硬件层
	*:docker + k8s

	通信层
	*:网络传输(RPC 用远程调用)
	1.HTTP传输. GET POST PUT DELETE
	2.基于TCP更靠近底层，RPC基于TCP
	*:需要知道谁调用谁 ->服务注册与服务发现
		1.需要分布式数据同步:etcd consul zk
	三者间的传递
	client server registry

	应用平台层
	云管理平台、监控平台、日志管理平台、服务治理平台

	微服务层
	实现业务逻辑

	微服务讲解:
	1.微服务架构
	解耦各模块逻辑、对外API网关

	2.服务注册与服务发现
	*:客户端做:需要一个注册中心，记录服务地址，知道具体微服务实例映射位置
	*:服务端做:简单、服务启动时，自动注册即可
	*:基本都是客户端做服务发现
	etcd用于解决分布式一致性场景:raft算法
	etcd场景:
	注册发现、共享配置、分布式锁、leader选举

	3.rpc调用与服务监控
	RPC相关内容
	#:数据传输: JSON XML Protobuf
	#:负载:随机算法 轮询 一致性hash 加权
	#:异常容错:健康检查 容断 限流
	服务监控
	#:日志收集
	#:打点采样


	Raft算法
	1.提供集群系统中分布式状态机通用方法，确保集群中的每个节点都同意一系列的状态转换
	2.一个Raft集群包含若干个服务器节点，通常为5个。
	3.节点有三种状态:follower(跟随者)、candidate(候选人)、leader(领导者)
	raft一致性算法:
	通过选出一个leader简化日志副本管理:日志项只准许leader流向follower
	可将raft算法分为三个子问题:
	leader election(领导选举):原来的leader挂掉后，需要选出新的leader
	log replication(日志复制):leader从客户端接收日志，并复制给整个集群
	safety(安全性)



 */


