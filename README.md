# wfuStruct
使用go实现微服务框架
--------------  
客户端:  
1.基于普罗米修斯实现中间件监控请求状态  
2.使用计数器、漏桶、令牌桶等算法进行限流  
3.基于hystrix实现中间件进行熔断  
4.基于jaeger实现中间件进行分布式追踪  
5.使用etcd作为服务注册中心，实现服务发现中间件  
6.使用轮询/加权轮询、随机/加权随机等算法，实现负载均衡中间件  
7.基于grpc框架进行远程调用  
8.基于模板，实现客户端代码自动生成  
  
服务端:  
1.使用etcd作为服务注册中心  
2.使用普罗米修斯实现中间件监控请求状态  
3.基于jaeger实现中间件进行分布式追踪  
4.使用计数器、漏桶、令牌桶等算法进行限流  
5.基于模板，实现服务端代码自动生成  

文件结构:
 - clientMidware      (客户端中间件)  
 - clientService      (客户端处理主流程)  
 - clientUtil         (客户端元数据处理)  
 - codeGenerate       (自动代码生成模块)  
 - logs               (日志库)  
 - midware            (服务端中间件)  
 - output             (自动代码生成结果)  
   - conf             (配置文件)  
   - controller       (需要实现的服务端逻辑)  
   - generate         (protobuff文件生成的代码)  
   - model            (需要实现的model包)  
   - router           (服务端grpc路由逻辑)  
   - scripts          (需要实现的脚本)  
   - client           (客户端调用逻辑)  
     - clientTool     (客户端内部调用逻辑)  
	 - client_main.go (需要实现的客户端调用逻辑)  
   - server           (需要实现的服务端调用逻辑)  
 - register           (服务注册/服务发现管理模块)  
   - etcdRegister     (etcd服务注册/服务发现模块)  
 - requsetBalance     (负载均衡算法实现)  
 - server             (服务端配置文件解析、服务端处理主流程)  
 - util               (用以自动代码生成时，检查调用逻辑文件是否存在)  