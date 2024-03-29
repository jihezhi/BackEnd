# 挑战任务集合：熟悉语言

有小伙伴表示不知道该从何做起，因此我设计了一系列挑战任务，来帮助大家逐步、连续地完成go语言的学习过程。

本挑战任务系列指导思想：

* 简洁、内聚、解耦：每个挑战任务只传达一个核心思想。
* 全面、连续、线性：该系列任务的目标是令新手从0到精通，获得所需的全部信息。

## 动机：为什么“重新造轮子”？

有小伙伴指责我说，go官方有教学文档，也有中文版，你为何要重新造轮子？

我造轮子的原因是，官方文档归根结底只是文档，它是资料的大集合。而我要做的，是符合我们学习理念的教程，我们的教程力求全面、完整，补全新手最头疼的“路线图”问题。我们追求唯一正解，因此我们只得出一个唯一的、线性的学习路径。而这是官方分散式的文档所不具备的性质。

当然，这次项目也是我们第一次尝试，效果如何还要看发展。

本系列挑战采取的格式为：

**想要传达的核心思想：可以体会到思想的具体任务**

## 系列1：熟悉语言（请在完成golang.org的tour后再尝试该系列）

* Go是比C更高级的语言：创建5个变量（随意赋值），其类型分别是string int float64 [3]bool和error。用最简洁的方式将它们输出到一行中。禁止使用Printf。
* Go具有几乎可以完成一切任务的标准库：使用Go编写一个排序程序。请采用最简单的写法（调用语言自带的package）
* Go使用最标准的方式做事：1.任意定义一个10进制数字（如123456789），以16进制的形式转换成string，然后打印至屏幕。2.任意定义一个写有一个10进制数字的string（比如“3.1415926”），将其转换为float64，然后打印至屏幕。
* interface是Go最重要的功能没有之一，尤其是io.Reader和io.Writer：请将一个TXT文本文档的内容输出至屏幕。禁止使用ioutil库。禁止使用Print系列命令。
* interface也是type：将一个string、一个[]byte和一个文件的内容输出到屏幕上。输出相关命令只允许在程序中出现一次，并且禁止使用Print系列命令。
* Go不带观点、没有捷径，什么工具干什么事：创造一个key为float64的map，赋予多个值，然后将该map按照key从小到大的顺序变成一个slice。
* Go routine不是多线程，而是描述实际问题逻辑的工具：创造一万个go routine，每个go routine负责将自己收到的数字加一再传出去。使用这一万个go routine完成将任意数字加一万的功能。

## 系列2：熟悉理念（难度跃升，请在完成系列1后进行）
* Don't communicate by sharing memory, share memory by communicating. 创造一个HTTP服务器，并开放一个API，该API接收一个int后，返回迄今为止接收到的所有int之和。注意：如果多个线程同时读写同一段内存，程序会崩溃。
* Concurrency is not parallelism. 写一个行列式求值算法，然后随机生成100个行列式，求其值的总和。通过合理使用go routine分割程序来加快你的程序的运行速度（比如说，如果你使用8核CPU，那么你应该能够通过分割你的程序结构来让你的算法和单线程运行相比提速6倍）
* The bigger the interface, the weaker the abstraction. 写一个函数，使它接受任意个在执行evaluate()方法后返回一个int、在执行describe()方法后返回一个string的type，将所有接收到的参数的evaluate和describe后的结果输出到屏幕上。在主程序中调用该函数，给它提供三个（或者任意个）不同的type（请随意自造type）。
* Make the zero values useful. 修改上面的练习，使得当我们传入了0个参数时函数将（你自己造的）使用说明输出到屏幕上。
* interface{} says nothing. 修改上面的练习，使得它接受的参数type为interface{}（其余不变），然后任意传入任意type的参数看看会发生什么。
* Errors are values. 修改上面的练习，使得当我们传入了0个参数时返回一种错误（里面包含你自己写的函数说明）；当我们传入了无法执行的type时返回另一种错误（里面包含你传入的type名和函数所期待的type的形式）。在主函数中执行该函数多次后将错误按照倒序输出到屏幕上。
* Don't just check errors, handle them gracefully. 修改上面的练习，使得当我们传入0个参数时，提示错误信息，然后要求用户从键盘输入一个int然后打印到屏幕上；当我们传入无法执行的type时，提示错误信息，然后要求用户从键盘输入一个int和一个string，然后程序使用这两个值替换掉原来无法执行的type重新执行。程序重复这个步骤直到所有参数都合格为止。
* Don't panic. 除非你在debug，不然不要在正式程序中使用任何panic。

## 系列3：熟悉实际项目开发中的元素（进入真正有用的实战，请在完成系列12后进行）
* 熟悉服务器：写一个文件服务器，使得别人可以通过链接访问你电脑上的本地文件。限制条件：一共只允许写一行代码。
* 熟悉数据传输：写一个聊天室程序，分为服务器端和客户端，当服务器端启动程序后，任何启动客户端程序的用户都会首先收到服务器端启动后的所有历史数据，然后用户可以输入聊天信息，发送后所有其他用户收到该消息。给每个用户分配一个随机的名字，数据请采用gzip压缩，UTF-8编码，JSON格式传输。这个练习可能相当之难，但对于实际应用水平来说仅仅算是入门。
* 熟悉数据库：修改上面的聊天室程序，使得历史消息全部保存入你的本地运行的数据库中（如果你没有偏好，请使用MySql）。这样一来，用户可以收到你的服务器端首次运行以来的所有历史消息。也就是说我们加了一个云同步功能。这个练习也可能相当之难，因为你会需要学习SQL语句。
