# 日志传输到ES

## 流程

- 1. 配置项加载，初始化kafka，es连接

- 2. 获取kafka中的topic，消费消息

- 3. 消息发送到管道中

- 4. 监听管道，将管道中的消息发送到es中
