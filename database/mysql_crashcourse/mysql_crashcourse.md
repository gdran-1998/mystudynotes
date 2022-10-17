

# MySQL必知必会

## ch1 了解SQL

**数据库**：一个文件或者一组文件。

数据库软件称为DBMS（数据库管理系统）。数据库 是通过DBMS创建和操纵的容器。

**表**：存储在表中的数据是一种类型的数据或一个清单。（唯一）

模式（schema）： 关于数据库和表的布局及特性的信息。有时，模式用作数据库的同义词。

**列**（column）： 表中的一个字段。所有表都是由一个或多个列组成的。数据库中每个列都有相应的数据类型。

**行**：表中的一个记录，表中的数据是按行存储的。

**主键**：一列（或一组列），其值能够唯一区分表 中每个行。（必须要有且要不同）

- 任意两行都不具有相同的主键值；

- 每个行都必须具有一个主键值（主键列不允许NULL值）。

**SQL：**一种与数据库通信的语言。



## ch2 MySQL简介

DBMS可分为两类：一类为基于共享文件系统的DBMS，另一类为基 于客户机—服务器的DBMS。	

**MySQL**:MySQL是一种DBMS，即它是一种数据库软件。是基于客户机—服务器的数据库管理系统。

- 服务器软件为MySQL DBMS。你可以在本地安装的副本上运行， 也可以连接到运行在你具有访问权的远程服务器上的一个副本。
- 客户机可以是MySQL提供的工具、脚本语言（如Perl）、Web应用开发语言（如ASP、ColdFusion、JSP和PHP）、程序设计语言（如 C、C++、Java）等。

**MySQL工具（客户机）：**mysql命令行实用程序、MySQL Administrator、MySQL Query Browser。 



## ch3 使用MySQL

### 连接

```sql
mysql -h localhost -P 3306 -u root -p
```

### 选择数据库

```sql
USE crashcourse;
```

### 了解数据库和表

```sql
# 返回可用数据库的一个列表
SHOW DATABASES;
# 返回当前选择的数据库内可用表的列表
SHOW TABLES;
# SHOW 也可以用来显示表列
SHOW COLUMNS FROM customers;
# DESCRIBE 作为 SHOW COLUMNS FROM 的一种快捷方式
DESCRIBE customers;
#################################
# 所支持的其他SHOW语句还有:
# SHOW STATUS用于显示广泛的服务器状态信息
# SHOW CREATE DATABASE和SHOW CREATE TABLE，分别用来显示创建特定数据库或表的MySQL语句
# SHOW GRANTS，用来显示授予用户（所有用户或特定用户）的安全权限；
# SHOW ERRORS和SHOW WARNINGS，用来显示服务器错误或警告消息。
```



## ch4 检索数据

使用SELECT语句从表中检索一个或多个数据列。

### 检索单个列

所需的列名在SELECT关键字之后给出，FROM 关键字指出从其中检索数据的表名。

```SQl
SELECT prod_name
FROM products;
```

结果返回 表中所有的行；结果没有过滤，也没有排序。

### 检索多个列

SELECT 关键字后给出多个列名，列名之间以逗号分隔。

```SQL
SELECT prod_id,prod_name,prod_price
FROM products;
```

### 检索所有列

使用星号（*）通配符检索所有的列

```sql
SELECT *
FROM products;
```

缺点：检索不需要的列通常会降低检索和应 用程序的性能。

优点：由于不明确指定列 名（因为星号检索每个列），所以能检索出名字未知的列。

### 检索不同的行

```SQL
SELECT DISTINCT vend_id
FROM products;
```

SELECT DISTINCT vend_id告诉MySQL只返回不同（唯一）的 vend_id行。

DISTINCT关键字应用于所有列而不仅是前置它的列。

### 限制结果

使用LIMIT子句，返回第一行或前几行。

```sql
SELECT prod_name
FROM products
LIMIT 5;
```

从第0行开始的5行

```sql
SELECT prod_name
FROM products
LIMIT 5,5;
```

从第5行开始的5行

所以带一个值的LIMIT总是从第一行开始，给出的数为返回的行数。 带两个值的LIMIT可以指定从行号为第一个值的位置开始。

LIMIT 4 OFFSET 3意为从行3开始取4行，就像LIMIT 3, 4一样。

### 使用完全限定的表名

完全限定的列名

```sql
SELECT products.prod_name
FROM products;
```

完全限定的表名

```sql
SELECT products.prod_name
FROM crashcourse.products;
```



## ch5 排序检索数据

使用SELECT语句的ORDER BY子句，根据需要排序检索出的数据。

为了明确地排序用SELECT语句检索出的数据，可使用ORDER BY子句。 ORDER BY子句取一个或多个列的名字，据此对输出进行排序。

**用非检索的列排序数据是完全合法的。**

### 排序数据

```sql
SELECT prod_name
FROM products
ORDER BY prod_name;
```

对prod_name列以字母顺序排序数据

### 按多个列排序

```sql
SELECT prod_id,prod_price,prod_name
FROM products
ORDER BY prod_price,prod_name;
```

代码检索3个列，并按其中两个列对结果进行排序——首先按 价格，然后再按名称排序。

**重要的是理解在按多个列排序时，排序完全按所规定的顺序进行。 换句话说，对于上述例子中的输出，仅在多个行具有相同的prod_price 值时才对产品按prod_name进行排序。如果prod_price列中所有的值都是唯一的，则不会按prod_name排序。**

### 指定排序方向

降序：DESC  升序：ASC  默认为升序

按价格以降序排序产品（最贵的排在最前面）

```sql
SELECT prod_id,prod_price,prod_name
FROM products
ORDER BY prod_price DESC;
```

以降序排序产品 （最贵的在最前面），然后再对产品名排序：

```sql
SELECT prod_id,prod_price,prod_name
FROM products
ORDER BY prod_price DESC,prod_name;
```

- 如果想在多个列上进行降序排序，必须 对每个列指定DESC关键字。

使用ORDER BY和LIMIT的组合，能够找出一个列中最高或最低的值。下面的例子演示如何找出最昂贵物品的值：

```sql
SELECT prod_price
FROM products
ORDER BY prod_price DESC
LIMIT 1;
```



## ch6 过滤数据

使用SELECT语句的WHERE子句指定搜索条件。

### 使用 WHERE 子句

在SELECT语句中，数据根据WHERE子句中指定的搜索条件进行过滤。 WHERE子句在表名（FROM子句）之后给出，如下所示：

```sql 
SELECT prod_name,prod_price
FROM products
WHERE prod_price=2.50;
```

从products表中检索两个列，但不返回所有行，只返回prod_price值为2.50的行。这里采用了相等测试。

- 在同时使用ORDER BY和WHERE子句时，应 该让ORDER BY位于WHERE之后，否则将会产生错误

### WHERE子句操作符

= 等于、<> != 等于、BETWEEN 在指定的两个之间

#### 检查单个值

```sql
SELECT prod_name,prod_price
FROM products
WHERE prod_name='fuses';
```

返回prod_name的值 为Fuses的一行。MySQL在执行匹配时默认不区分大小写，所 以fuses与Fuses匹配。

列出价格小于10美元的所有产品：

```sql
SELECT prod_name,prod_price
FROM products
WHERE prod_price < 10;
```

句检索价格小于等于10美元的所有产品

```sql
SELECT prod_name,prod_price
FROM products
WHERE prod_price <= 10;
```

#### 不匹配检查

列出不是由供应商1003制造的所有产品

```sql
SELECT vend_id,prod_name
FROM products
WHERE vend_id <> 1003;
```

<> 可用 ！= 替换

- 单引号用来限定字符串。如果将值与串类型的 列进行比较，则需要限定引号。用来与数值列进行比较的值不 用引号。

#### 检查范围

检索价格在5美元和10 美元之间的所有产品：

```sql
SELECT prod_name,prod_price
FROM products
WHERE prod_price BETWWEN 5 AND 10;
```

这两个值必须用AND关键字 分隔。BETWEEN匹配范围中所有的值，包括指定的开始值和结束值。

#### 空值检查

NULL 无值（no value），它与字段包含0、空字符串或仅仅包含 空格不同。

检查具有NULL值的列。

```sql
SELECT cust_id
FROM customers
WHERE cust_email IS NILL;
```

- 在通过过滤选择出不具有特定值的行时，你 可能希望返回具有NULL值的行。但是，不行。因为未知具有特殊的含义，数据库不知道它们是否匹配，所以在匹配过滤或不匹配过滤时不返回它们。 因此，**在过滤数据时，一定要验证返回数据中确实给出了被过滤列具有NULL的行。**

## ch7 数据过滤

#### 组合WHERE子句

MySQL允许给出多个WHERE子句。这些子 句可以两种方式使用：以AND子句的方式或OR子句的方式使用。

**操作符**（operator） 用来联结或改变WHERE子句中的子句的关键 字。也称为逻辑操作符（logical operator）。

##### AND操作符

**AND 用在WHERE子句中的关键字，用来指示检索满足所有给定条件的行。**

由供应商1003制造且价格小于等于10美元的所有产品的名称和价格。

```sql
SELECT prod_id,prod_price,prod_name
FROM products
WHERE vend_id = 1003 AND prod_price <= 10;
```

##### OR操作符

**OR WHERE子句中使用的关键字，用来表示检索匹配任一给定条件的行。**

由任一个指定供应商制造的所有产品的产品 名和价格。

```sql
SELECT prod_name,prod_price
FROM products
WHERE vend_id = 1003 OR vend_id = 1004;
```

##### 计算次序

AND 高于 OR,圆括号具有较AND或OR操作符高 的计算次序。

列出价格为10美元（含）以上且由1002或1003制 造的所有产品

```sql
SELECT prod_name,prod_price
FROM products
WHERE (vend_id = 1002 OR vend_id = 1003) AND prod_price >= 10;
```

**任何时候使用具有AND和OR操作 符的WHERE子句，都应该使用圆括号明确地分组操作符。**

#### IN 操作符

IN操作符用来指定条件范围，范围中的每个条件都可以进行匹配。

检索供应商1002和1003制造的所有产品。

```sql
SELECT prod_name,prod_price
FROM products
WHERE vend_id IN(1003,1004)
ORDER BY prod_name;
```

**IN操作符完成与OR相同的功能,IN 有其自身的优点**

#### NOT 操作符

NOT WHERE子句中用来否定后跟条件的关键字。

列出除1002和1003之外的所有供应 商制造的产品

```sql
SELECT prod_name,prod_price
FROM products
WHERE vend_id NOT IN (1003,1004)
ORDER BY prod_name;
```

## ch8 用通配符进行过滤

#### LIKE 操作符

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

![img](img/vendors.png)

- 表示存储销售产品的供应商
- vend_id 列用来匹配产品和供应商
- vend_id 作为主键。vend_id为一个自动增量字段。

**products表**：

![img](img/products.png)

- 产品目录，每行一个产品
- prod_id 作为主键
- 为实施引用完整性，在vend_id 上定义一个外键，关联到vendors的vend_id。

**customers表：**

![img](img/customers.png)

- 存储所有顾客的信息
- cust_id 为主键，是一个自动增量字段。

**orders表：**

![img](img/orders.png)

- 存储顾客订单（不是订单细节）
- order_num为主键，是一个自动增量字段。
- 为实施引用完整性，在cust_id 上定义一个外键，关联到customers的cust_id。

**orderitems表：**

![img](img/orderitems.png)

- 存贮每个订单中的实际物品，每个订单的每个物品占一行。
- order_num和order_item 作为主键。
- 为实施引用完整性，在order_num上定义外键，关联它到 orders表的order_num，在prod_id上定义外键，关联它到products 表 的prod_id。

**productnotes表：**

![img](img/productnotes.png)

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

![img](img/string.png)

**既然变长数据类型这样灵活，为什么还要使用定长数据类型？**

答：因为性能。MySQL处理定长列远比处理变长列快得多。此外，MySQL不 允许对变长列（或一个列的可变部分）进行索引。这也会极大地影响性能。

**注：**不管使用何种形式的串数据类型，串值都必须括在 引号内（通常单引号更好）。电话号码和邮政编码使用串类型存储。

### 数值数据类型

![img](img/numeric.png)

### 日期和时间数据类型

![img](img/data_time.png)

### 二进制数据类型

二进制数据类型可存储任何数据（甚至包括二进制信息），如图像、 多媒体、字处理文档等

![img](img/binary.png)