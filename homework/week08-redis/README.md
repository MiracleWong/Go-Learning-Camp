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

## 测试结论

随着value增大，set/get吞吐降低；在测试数据集中时，降低不太明显。

## 参考资料
1. [Redis性能测试工具redis-benchmark使用](https://blog.csdn.net/ccccsy99/article/details/107823849)


# Question 2
题目：写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息, 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

```shell
默认端口：6379

# 最初的info memory
flushdb
info memory

# 执行命令后的info memory
redis-benchmark -h 127.0.0.1 -q -t set -r 100000 -n 500000 -d 10  (10 可变)

info memory
key *
```


| value字节大小/B | 写入前used_memory_dataset/B | 写入后used_memory_dataset/B | key数量 | 平均每个key的占用内存空间/B |
| --- | --- | --- | --- | --- |
| 10 | 11202 | 3468962 | 86444 | 40.00 |
| 20 | 11354 | 4171274 | 86665 | 40.00 |
| 50 | 12482 | 6939666 | 86602 | 79.99 |
| 100 | 11834 | 11783858 | 86559 | 136.00 |
| 200 | 11986 | 21432738 | 86374 | 248.00 |
| 1k | 12138 | 112716858 | 86430 | 1304.00 |
| 5k | 12290 | 533291402 | 86459 | 6168.00 |