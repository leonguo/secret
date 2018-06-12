# 高性能开发框架

基础路由框架:echo

纯JSON输出接口压测结果： QPS 5万+
```
    wrk -t12 -c100 -d30s http://localhost:8999/test
    Running 30s test @ http://localhost:8999/test
      12 threads and 100 connections
      Thread Stats   Avg      Stdev     Max   +/- Stdev
        Latency     2.04ms    2.65ms  93.95ms   94.62%
        Req/Sec     4.64k   662.15    11.86k    78.83%
      1663838 requests in 30.05s, 285.62MB read
    Requests/sec:  55362.51
    Transfer/sec:      9.50MB
```

框架结构
  - app --应用目录
       - middleware    -- 基础组件
       - controllers   -- 业务逻辑
       - lib           -- 业务基础库
       - models        -- 数据模型
       - public        -- 资源
       - app.go        -- 应用入口
       - route.go      -- 路由
  - config --服务配置文件
       - config.toml  -- 配置文件
  - db --数据库统一入口
       - gorm  -- ORM
       - redis -- redis操作
  - util  --全局公用类
  - tests --单元测试
  - server.go --框架启动文件