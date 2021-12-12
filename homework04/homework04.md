## 作业1：
#### windows11平台

#### redis benchmark

###### 1. 10B get/set

```shell
redis-benchmark -h 127.0.0.1 -p 6379 -d 10 -t get,set
```

![10B](/10B.png)

###### 2. 20B get/set

```shell
redis-benchmark -h 127.0.0.1 -p 6379 -d 20 -t get,set
```

![20B](/20B.png)

###### 3. 50B get/set

```shell
redis-benchmark -h 127.0.0.1 -p 6379 -d 50 -t get,set
```

![50B](/50B.png)

###### 4. 100B get/set

```shell
redis-benchmark -h 127.0.0.1 -p 6379 -d 100 -t get,set
```

![100B](/100B.png)

###### 5. 200B get/set

```shell
redis-benchmark -h 127.0.0.1 -p 6379 -d 200 -t get,set
```

![200B](/200B.png)

###### 6. 1000B get/set

```shell
redis-benchmark -h 127.0.0.1 -p 6379 -d 1000 -t get,set
```

![1000B](/1000B.png)

###### 7. 5000B get/set

```shell
redis-benchmark -h 127.0.0.1 -p 6379 -d 5000 -t get,set
```

![5000B](/5000B.png)

#### 结论

随着字节变大，set 还是很平稳的，get 变化挺大。不过整体性能还是很好。



## 作业2：

#### info memory

![info memory](/info.png)

额，没有数据源，尴尬了😅

