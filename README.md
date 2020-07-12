# mysql-to-proto

使用过grpc的同学都知道，写proto文件比较繁琐，尤其是写message，对应很多字段，为此写了一个简单的从mysql直接读取表结构，生成proto文件的工具。

工具的使用很简单，需要简单的配置，即可运行生成proto文件。

### 准备

使用前需先安装依赖包go-sql-driver/mysql

```
$ go get -u github.com/go-sql-driver/mysql
```



### 使用说明：
修改数据配置和要装换的表，一次执行一张表的转换，生成多个proto文件后

window 下去执行pb.bat,可以一次转换多个proto文件的不同语言服务端，这里展示的生成的golang ，

如果需要编译其他语言，修改相应的可执行文件和bat文件dos命令即可 



## LICENSE

BSD License <http://creativecommons.org/licenses/BSD/>

Book License: [CC BY-SA 3.0 License](http://creativecommons.org/licenses/by-sa/3.0/)