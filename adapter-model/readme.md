# adapter model
## introduce
配饰器模式是为了解决函数名不统一的问题。
## example
例如，一个团队想调用写数据库的函数，但是不用的数据库写命令的函数名不一样。写sql使用writeSql，写tiDB使用writeTiDB。