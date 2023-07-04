# my-help
支持一些简单的辅助功能，基于golang实现，可以方便实现跨平台的使用，避免在多个平台找工具;
准备支持的简单功能：

1. 下载文件
2. 计算md5
3. unzip文件
4. 时间的换算
5. 简单的加减乘除的计算
6. 进制转换，比如10进制转二进制等



## 编译

**windows：**

```
set GO111MODULE=on

.\win_compile.bat
```



**linux下：**

```
export GO111MODULE=on

make tool
```



## 例子：

my-help --help