- [第2章 指针操作](#org4c84ad9)
  - [指针基础](#org5316160)
  - [存储空间分配](#org2eb01c8)
  - [数据集合与指针的算数运算](#orgb52106d)
    - [结构](#org431d32b)
    - [数组](#org07c8540)
  - [作为函数参数的指针](#orgef02d2f)
  - [按引用调用传递参数](#orga356892)
  - [作为参数指向指针的指针](#orgfaf6319)
  - [泛型指针与类型转换](#org9a5bf47)
    - [泛型指针](#orgc5848a3)
    - [类型转换](#orgafd8c28)
  - [函数指针](#org3b6abe3)
- [第3章 递归](#org7582ca4)
  - [基本递归](#orgd6d41cd)
  - [尾递归](#orgace53d1)
- [第4章 算法分析](#orgf2c94f1)
  - [最坏情况分析](#orgb48fc62)
  - [大O表示法](#org152f333)
  - [计算的复杂度](#orgbffb6ca)
- [第5章 链表](#org198d317)
  - [单链表介绍](#orgff1f510)
  - [单链表接口的定义](#orgc80cd79)
    - [list<sub>init</sub>](#orga0ea03d)
    - [list<sub>destroy</sub>](#org1bfdc7b)
    - [list<sub>ins</sub><sub>next</sub>](#org531645a)
    - [list<sub>rem</sub><sub>next</sub>](#org435fd80)
    - [list<sub>size</sub>](#org24a161c)
    - [list<sub>head</sub>](#orgaf2bde8)
    - [list<sub>tail</sub>](#orgb704779)
    - [list<sub>is</sub><sub>head</sub>](#org26616f1)
    - [list<sub>is</sub><sub>tail</sub>](#org7feb8b3)
    - [list<sub>data</sub>](#org4c876c6)
    - [list<sub>next</sub>](#orgab6536f)
  - [单链表的实现](#org2818ced)
  - [双向链表介绍](#orgce96cf5)
  - [双向链表接口的定义](#orgaadef95)
    - [DList<sub>init</sub>](#orgc0e0254)
    - [Destroy](#org6db5178)
    - [Ins<sub>next</sub>](#org1e3d738)
    - [Ins<sub>prev</sub>](#orge01ce26)
    - [Remove](#org160c3d1)
  - [双向链表的实现](#org273f173)
  - [循环链表的介绍](#orge7b5ae5)
  - [循环链表接口的定义](#org177a047)
    - [CList<sub>init</sub>](#org7efc0ac)
    - [Ins<sub>next</sub>](#orgcf45beb)
    - [Rem<sub>next</sub>](#orgffc0085)
  - [循环链表的实现](#org27567d7)



<a id="org4c84ad9"></a>

# 第2章 指针操作


<a id="org5316160"></a>

## 指针基础

一个指针只是一个变量,它存储数据在内存中的地址.


<a id="org2eb01c8"></a>

## 存储空间分配

当声明一个指针时,仅仅只为该指针本身分配空间,并没有为指针所引用的数据分配空间.


<a id="orgb52106d"></a>

## 数据集合与指针的算数运算

数据集合: 结构和数组


<a id="org431d32b"></a>

### 结构

结构由任意的有序的元素构成,类似 Product Type.

```c
typedef struct ListElmt_{
void* data;
struct ListElmt_* next;
}
```

结构不允许包含自身的实例,但可以包含指向自身实例的指针.


<a id="org07c8540"></a>

### 数组

数组是内存中连续排列的同类元素的序列. a[i] `= *(a+i) a[i][j] =` \*(​\*(a + i) + j)​ ​


<a id="orgef02d2f"></a>

## 作为函数参数的指针

按引用传递参数时,当函数改变此参数时,这个被改变的值会一直存在,甚至函数退出后仍然存在. 使用指针传递大容量复制的函数参数.


<a id="orga356892"></a>

## 按引用调用传递参数

通常情况下,c语言只支持按值传参(pass by value).但我们可以模仿按引用调用传递参数(pass by reference)将一个指向参数的指针传递给函数.


<a id="orgfaf6319"></a>

## 作为参数指向指针的指针

作为参数指向指针的指针,是因为在函数想改变传递给它的指针.


<a id="org9a5bf47"></a>

## 泛型指针与类型转换

使用泛型指针模拟泛型.泛型指针并不指定具体的数据类型.


<a id="orgc5848a3"></a>

### 泛型指针

通常情况下,c语言只允许相同类型的指针之间进行转换.但是泛型指针能够转换为任何类型的指针,反之亦然.


<a id="orgafd8c28"></a>

### 类型转换

将类型T的变量t转换为类型S的变量s.只需要 (S)t. 可能破坏内存中的数据对齐.


<a id="org3b6abe3"></a>

## 函数指针

函数指针是指向可执行代码段或调用可执行代码段的信息块的指针,而不是指向某种数据的指针. 声明函数指针的形式: return-value (\*function-name) (args&#x2026;). 调用形式与调用函数相同. 函数指针的一个重要用途是将函数封装到数据结构中.


<a id="org7582ca4"></a>

# 第3章 递归


<a id="orgd6d41cd"></a>

## 基本递归

基本递归过程有两个基本阶段: 递推和回归 一个可执行程序由4个区域构成: 代码段,静态数据区,堆和栈. 代码段包含程序运行所执行的机器指令,静态数据区包含在程序生命周期内一直持久的数据,如全局变量和静态局部变量. 堆包含程序运行时动态分配的存储数据,栈包含函数调用的信息. 堆的增长方向由程序低地址向高地址增长,栈则相反. 当c程序调用了一个函数时,栈中会分配一块空间保存调用相关信息,称为栈帧.


<a id="orgace53d1"></a>

## 尾递归

在基本递归的基础上,尾递归在回归过程中不执行操作.编译器可以优化此函数.


<a id="orgf2c94f1"></a>

# 第4章 算法分析


<a id="orgb48fc62"></a>

## 最坏情况分析

分析算法的最佳情况的性能没有太多意义. 分析算法的平均情况的性能不那么容易. 分析算法的最坏情况可以告诉我们算法性能的上限.


<a id="org152f333"></a>

## 大O表示法

大O表示法的基本规则: 优化常数项和常数因子,只考虑高阶项的因子.


<a id="orgbffb6ca"></a>

## 计算的复杂度

使用上述两种方法完成计算的复杂度分析


<a id="org198d317"></a>

# 第5章 链表


<a id="orgff1f510"></a>

## 单链表介绍

单链表由各个元素之间通过一个指针彼此链接起来而组成.每个元素包含两部分: 数据成员和一个称为next的指针.将每一个元素next指针设置为指向后面的元素.最后一个元素的next指针指向NULL.


<a id="orgc80cd79"></a>

## 单链表接口的定义


<a id="orga0ea03d"></a>

### list<sub>init</sub>

```c
void list_init(List* list,void (*destroy) (void* data));
```

返回值: 无. 描述: 初始化有参数 list 指定的链表,该函数必须在链表做其他操作之前调用.destroy 参数提供了一种释放动态分配的数据的方法.如果链表包含不应该释放的数据或者不需要动态释放空间的数据时,destroy应该设置为 NULL. 复杂度: O(1).


<a id="org1bfdc7b"></a>

### list<sub>destroy</sub>

```c
void list_destroy(List* list);
```

返回值: 无. 描述: 销毁由参数list指定的链表,调用list<sub>destroy后不允许执行其他关于此list的操作.list</sub><sub>destroy</sub> 将链表中的所有元素都移除,如果list<sub>init中的destroy不为NULL,则移除链表中每个元素时都调用该函数一次</sub>. 复杂度: O(n),n为链表的长度.


<a id="org531645a"></a>

### list<sub>ins</sub><sub>next</sub>

```c
int list_ins_next(List *list,ListElmt *element,const void *data);
```

返回值: 插入元素成功则返回0,否则返回-1. 描述: 在list指定的链表中element后面插入一个新元素,如果element为NULL,则新链表插入链表头部.新元素包含一个指向data的指针,因此只要该元素还在链表中,data所引用的内存应该保持合法.管理data所引用的储存空间是调用者的责任. 复杂度: O(1).


<a id="org435fd80"></a>

### list<sub>rem</sub><sub>next</sub>

```c
int list_ins_next(List *list,ListElmt *element,void **data);
```

返回值: 删除元素成功则返回0,否则返回-1. 描述:与list<sub>ins</sub><sub>next,只不过由插入改为删除</sub>. 复杂度: O(1).


<a id="org24a161c"></a>

### list<sub>size</sub>

```c
int list_size(const List *list);
```

返回值: 链表中元素的个数. 描述: 这是一个宏,用来计算由参数list指定的链表中的元素的个数. 复杂度:O(1).


<a id="orgaf2bde8"></a>

### list<sub>head</sub>

```c
ListElmt *list_head(const List *list);
```

返回值: 指向链表中头元素的指针. 描述: 这是一个宏,返回由参数list指定的链表中头元素的指针. 复杂度: O(1).


<a id="orgb704779"></a>

### list<sub>tail</sub>

```c
ListElmt *list_tail(const List *list);
```

返回值: 指向链表中尾元素的指针. 描述: 这是一个宏,返回由参数list指定的链表中尾元素的指针. 复杂度: O(1).


<a id="org26616f1"></a>

### list<sub>is</sub><sub>head</sub>

```c
int list_is_head(const ListElmt *element);
```

返回值: 如果element所指定的元素是链表头节点则返回1;否则返回-1. 描述: 这是一个宏,用来判断element所指定的元素是否是链表的链表头结点. 复杂度: O(1).


<a id="org7feb8b3"></a>

### list<sub>is</sub><sub>tail</sub>

```c
int list_is_tail(const ListElmt *element);
```

返回值: 如果element所指定的元素是链表尾节点则返回1;否则返回-1. 描述: 这是一个宏,用来判断element所指定的元素是否是链表的链表头结点. 复杂度: O(1).


<a id="org4c876c6"></a>

### list<sub>data</sub>

```c
void *list_data(const ListElmt *element);
```

返回值: 节点中保存的数据. 描述: 这是一个宏,返回由element所指定的链表结点元素保存的数据. 复杂度: O(1).


<a id="orgab6536f"></a>

### list<sub>next</sub>

```c
ListElmt *list_next(const ListElmt *element);
```

返回值: 返回由参数element指定的节点的下一个节点. 描述: 这是一个宏,返回由参数element指定的节点的下一个节点. 复杂度: O(1).


<a id="org2818ced"></a>

## 单链表的实现

[单链表](src/ch5/list.go)


<a id="orgce96cf5"></a>

## 双向链表介绍

双向链表元素之间由两个指针链接,双向链表的每一个元素由三部分组成: data,prev,next. 为了标识链表的头与尾,将第一个元素的prev指针与最后一个元素的next指针设置为 nil.


<a id="orgaadef95"></a>

## 双向链表接口的定义


<a id="orgc0e0254"></a>

### DList<sub>init</sub>

```go
func DList_init() *DList
```

返回值: 指向DList的指针. 描述: 初始化双向链表. 复杂度: O(1).


<a id="org6db5178"></a>

### Destroy

```go
func (lst *DList) Destroy()
```

返回值: 无. 描述: 销毁由参数lst指定的链表. 复杂度: O(n).


<a id="org1e3d738"></a>

### Ins<sub>next</sub>

```go
func (lst *DList) Ins_next(element *DListElmt,data any) error
```

返回值: 如果插入成功返回 nil,否则返回具体错误. 描述: 将data插入由list指定的双向链表中element之后,element当且仅当lst为空列表时才能为nil. 复杂度: O(1).


<a id="orge01ce26"></a>

### Ins<sub>prev</sub>

```go
func (lst *DList) Ins_prev(element *DListElmt ,data any) error
```

返回值: 如果插入成功返回nil,否则返回具体错误. 描述: 将data插入由list指定的双向链表中element之前,element当且仅当lst为空列表时才能为nil. 复杂度: O(1).


<a id="org160c3d1"></a>

### Remove

```go
func (lst *DList) Remove(element *DListElmt) (any, error)
```

返回值: 如果移除成功,返回(value,nil),否则返回(nil,error) 描述: 从lst指向的双向链表中移除由element指定的元素. 复制度: O(1)


<a id="org273f173"></a>

## 双向链表的实现

[双向链表](src/ch5/dlist.go)


<a id="orge7b5ae5"></a>

## 循环链表的介绍

在单向循环链表中,最后一个元素的next指针又指回头元素而不是设置为nil. 在双向循环链表中,头元素的prev指针则指向最后一个元素.


<a id="org177a047"></a>

## 循环链表接口的定义


<a id="org7efc0ac"></a>

### CList<sub>init</sub>

```go
func CList_init() *CList
```

返回值: 指向CList的指针. 描述: 初始化双向链表. 复杂度: O(1).


<a id="orgcf45beb"></a>

### Ins<sub>next</sub>

```go
func (lst *CList) Ins_next(element *CListElmt,data any) error
```

返回值: 如果插入成功返回 nil,否则返回具体错误. 描述: 将data插入由list指定的循环链表中element之后,element当且仅当lst为空列表时才能为nil. 复杂度: O(1).


<a id="orgffc0085"></a>

### Rem<sub>next</sub>

```go
func (lst *CList) Rem_next(element *CListElmt) (any,error)
```

返回值: 如果移除成功,返回(value,nil),否则返回(nil,error) 描述: 从lst指向的循环链表中移除由element指定的元素. 复制度: O(1)


<a id="org27567d7"></a>

## 循环链表的实现

[循环链表](src/ch5/clist.go)
