# HomeWork Title

# Question 1
题目：1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

## 机器信息
青云 CentOS 8.2 的机器，1核2G

## Redis 信息
5.0.3

## 测试指令
备注：默认端口6379

```shell

redis-cli -p 6379 flushall
redis-benchmark -h 127.0.0.1 -q -t set,get -r 1000000 -n 500000 -d 10

redis-cli -p 6379 flushall
redis-benchmark -h 127.0.0.1 -q -t set,get -r 1000000 -n 500000 -d 20

redis-cli -p 6379 flushall
redis-benchmark -h 127.0.0.1 -q -t set,get -r 1000000 -n 500000 -d 50

redis-cli -p 6379 flushall
redis-benchmark -h 127.0.0.1 -q -t set,get -r 1000000 -n 500000 -d 100

redis-cli -p 6379 flushall
redis-benchmark -h 127.0.0.1 -q -t set,get -r 1000000 -n 500000 -d 200

redis-cli -p 6379 flushall
redis-benchmark -h 127.0.0.1 -q -t set,get -r 1000000 -n 500000 -d 1024

redis-cli -p 6379 flushall
redis-benchmark -h 127.0.0.1 -q -t set,get -r 1000000 -n 500000 -d 5120

```

## 测试结果

|     | SET (requests / s)| GET (requests / s) |
|  ----  | ---- | ---- |
| 10  | 42673.04 | 43828.89|
| 20  | 42484.49 | 43940.59|
| 50  | 42466.45 | 44161.81|
| 100 | 42201.22 | 44487.94|
| 200 | 41760.63 | 42837.56|
| 1k  | 38705.68 | 42767.94|
| 5k  | 27022.64 | 24284.81|


## 参考资料
1. [Redis性能测试工具redis-benchmark使用](https://blog.csdn.net/ccccsy99/article/details/107823849)


# Question 2
写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息, 分析上述不同 value 大小下，平均每个 key 的占用内存空间。