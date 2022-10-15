# MySQL必知必会

## ch1 了解SQL

**数据库**：一个文件或者一组文件。

数据库软件称为DBMS（数据库管理系统）。数据库 是通过DBMS创建和操纵的容器。

**表**：存储在表中的数据是一种类型的数据或一个 清单。（唯一）

模式（schema）： 关于数据库和表的布局及特性的信息。

**列**（column）： 表中的一个字段。所有表都是由一个或多个列组 成的。数据库中每个列都有相应的数据类型。

**行**：表中的一个记录，表中的数据是按行存储的。

**主键**：一列（或一组列），其值能够唯一区分表 中每个行。（必须要有且要不同）

- 任意两行都不具有相同的主键值；

- 每个行都必须具有一个主键值（主键列不允许NULL值）。

**SQL：**一种与数据库通信的语言。



## ch2 MySQL简介

**MySQL**:MySQL是一种DBMS，即它是一种数据库软件。是基于客户机—服 务器的数据库。

- 服务器软件为MySQL DBMS。你可以在本地安装的副本上运行， 也可以连接到运行在你具有访问权的远程服务器上的一个副本。
- 客户机可以是MySQL提供的工具、脚本语言（如Perl）、Web应用 开发语言（如ASP、ColdFusion、JSP和PHP）、程序设计语言（如 C、C++、Java）等。

**MySQL工具（客户机）：**mysql命令行实用程序、MySQL Administrator、MySQL Query Browser。 



## ch3 使用MySQL





## 附录A： MySQL入门

### 需要什么

为使用MySQL和学习本书中各章的内容，你需要访问MySQL服务器 和客户机应用（用来访问服务器的软件）副本。

**服务器**：（不一定需要自己安装MySQL副本，但需要访问服务器。基本上有 下面两种选择。）

- 访问一个已有的MySQL服务器，或许是你的公司或许是商用的或 院校的服务器。为使用这个服务器，你需要得到一个服务器账号 （一个登录名和一个口令。）
- 下载MySQL服务器的一个免费副本，安装在你自己的计算机上。

**客户及软件：**（用来实际运行 MySQL命令的程序。）

- mysql命令行实用程序
- MySQL Adiminstrator
- MySQL Query Browser

### 获得软件

学习更多的mysql知识：http://dev.mysql.com/

下载服务器的一个副本：http://dev.mysql.com/downloads/

MySQL Adiminstrator和MySQL Query Browser不作为MySQL的核心 部分安装，必须从http://dev.mysql.com/downloads/下载。

### 安装软件

你要安装一个本地MySQL服务器，应该在安装可选的MySQL实 用程序之前进行。

所有安装都会提示你 输入需要的信息，包括：（不确定要指定什么，默认就好）

- 安装位置
- root用户的口令
- 端口、服务或进程名等

多个MySQL服务器的副本可安装在单 台机器上，只要每个服务器使用不同的端口即可。

## 附录B：样例表

### 样例表

随身物品推销商使用的**订单录入系统**

使用6张表完成任务：

- 管理供应商
- 管理产品目录
- 管理顾客列表
- 录入顾客订单



**vendors表**：

![img](vendors.png)

- 表示存储销售产品的供应商
- vend_id 列用来匹配产品和供应商
- vend_id 作为主键。vend_id为一个自动增量字段。

**products表**：

![img](products.png)

- 产品目录，每行一个产品
- prod_id 作为主键
- 为实施引用完整性，在vend_id 上定义一个外键，关联到vendors的vend_id。

**customers表：**

![img](customers.png)

- 存储所有顾客的信息
- cust_id 为主键，是一个自动增量字段。

**orders表：**

![img](orders.png)

- 存储顾客订单（不是订单细节）
- order_num为主键，是一个自动增量字段。
- 为实施引用完整性，在cust_id 上定义一个外键，关联到customers的cust_id。

**orderitems表：**

![img](orderitems.png)

- 存贮每个订单中的实际物品，每个订单的每个物品占一行。
- order_num和order_item 作为主键。
- 为实施引用完整性，在order_num上定义外键，关联它到 orders表的order_num，在prod_id上定义外键，关联它到products 表 的prod_id。

**productnotes表：**

![img](productnotes.png)

- 存储与特定产品有关的注释。
- note_id 为主键。
- 列note_text必须为FULLTEXT搜索进行索引。
- 由于这个表使用全文本搜索，因此必须指定ENGINE=MyISAM。

### 创建样例表

下载相关：http://www.forta.com/books/0672327120/

步骤（mysql命令行实用程序）：

1. 创建一个数据源
2. 选择这个数据源，用`USE` 命令。
3. 执行`create.sql`脚本，`source create.sql`(指定create.sql文件的完全路径)。
4. 重复前面的步骤，用populate.sql文件填充各个新表。

## 附录C：MySQL语句的语法

最常使用 的MySQL语句的语法。每条语句以简要的描述开始，然后给出它的语法。 为增加方便性，还给出对讲授相应语句的章的交叉引用。



## 附录D:MySQL数据类型

### 串数据类型

两种基本的串数据类型：

- 定长串

长度固定，长度在创建表时指定。

- 变长串

长度可变，`TEXT`属于变长串类型

![img](string.png)

**既然变长数据类型这样灵活，为什么还要使用定长数据类型？**

答：因为性能。MySQL处理定长列远比处理变长列快得多。此外，MySQL不 允许对变长列（或一个列的可变部分）进行索引。这也会极大地影响性能。

**注：**不管使用何种形式的串数据类型，串值都必须括在 引号内（通常单引号更好）。电话号码和邮政编码使用串类型存储。

### 数值数据类型

![img](numeric.png)

### 日期和时间数据类型

![img](date_time.png)

### 二进制数据类型

二进制数据类型可存储任何数据（甚至包括二进制信息），如图像、 多媒体、字处理文档等

![img](binary.png)