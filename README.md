# my-help
支持一些简单的辅助功能，基于golang实现，可以方便实现跨平台的使用，避免在多个平台找工具;
已经支持了的简单功能：

1. 下载文件
2. 计算md5
3. unzip文件
4. 时间的换算
5. 简单的加减乘除的计算
6. telnet
7. windows下的pstree



## 编译

**windows：**

```
set GO111MODULE=on

.\win_compile.bat
```



**linux：**

```
export GO111MODULE=on

make tool
```



## 例子：

my-help --help

.\bin\my-help.exe time -s "1136214245000000000"

.\bin\my-help.exe time -s "2006-01-02 15:04:05"

.\bin\my-help.exe calc -s "1888*1111111*1212"

.\bin\my-help.exe pstree