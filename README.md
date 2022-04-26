##### 项目规范
1. cmd: 放main.go和配置文件, 作为启动入口
2. conf: 放配置文件对应的golang struct, 使用的是toml
3. model: 放结构体, 比如Http参数转换用的struct, DB存储对应的struct, 各层之间传递用的struct
4. dao: data access object, 数据库访问方法, redis, memcached访问方法, 还有一些RPC调用也放在这里面
5. http: 提供http服务, 主要是提供协议转换, 聚合. 逻辑还是再service层做.
6. service: 提供服务和逻辑实现.
7. grpc: 提供rpc服务.

### 队列查询命令
1. bucket
   * zrange dmq:bucket:2 0 -1 withscores
2. pool
   * get dmq:jobpool:1
3. ready queue
   * 
### TODO
1. ~~push api~~
   * ~~push逻辑~~
2. ~~ticker~~
   * ~~扫描bucket~~
3. 日志组件
   * 按天分割
4. ~~消费者~~
5. ~~http client~~
5. ~~消费次数限制~~
5. ~~container参数验证~~
6. blpop会影响安全退出，暂时用lpop
