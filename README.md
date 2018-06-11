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
