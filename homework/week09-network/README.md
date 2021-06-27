# HomeWork Title

# Question 1
总结几种 socket 粘包的解包方式: fix length/delimiter based/length field based frame decoder。尝试举例其应用

1. fix length：固定长度大小
    - 客户端在发送数据包的时候，每个包都固定长度，比如1024个字节大小，如果客户端发送的数据长度不足1024个字节，则通过补充空格的方式补全到指定长度
2. delimiter based：分隔符做结束标志
    - 在包尾增加回车换行符进行分割
3. length field based frame decoder：消息头指定消息长度
    - 将消息分成消息头和消息体，消息头中包含表示消息(体)总长度的字段
    
补充：
4. line based：基于行来进行消息粘包拆包


## 参考资料
1. [netty学习笔记一：TCP粘包拆包](https://www.cnblogs.com/magotzis/p/9527024.html)
2. [Golang解决TCP粘包拆包问题](https://blog.csdn.net/cqims21/article/details/104740507)


# Question 2
实现一个从 socket connection 中解码出 goim 协议的解码器。