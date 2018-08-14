# 深入理解if和for

首先，让我们来看azgy在尝试for循环时提交的问题代码。

```
package main

import (
	"fmt"
)

type Vertex struct{
	a int
	b int
	c int
	d int
	e int
}

func main(){
	v := Vertex{2,5,3,6,1}
	for i:=1;v.a>v.b&&v.b>v.c&&v.c>v.d&&v.d>v.e;i++ {
		if v.a<v.b {
		}
		if v.b<v.c {
			v.b,v.c = v.c,v.b
		}
		if v.c<v.d {
			v.c,v.d=v.d,v.c
		}
		if v.d<v.e {
			v.d,v.e=v.e,v.d
		}
		fmt.Println(v)
	}	
}
```
该代码有如下问题：
* import一个package的时候，可以写成import "fmt"，不需要写括号。引入多个package的时候为了避免反复写import才写括号。
* azgy并没有搞清楚struct和slice的区别。struct用来定义数据结构，而slice用来存放一系列具有相同类型的值。
* Don't Repeat Yourself。azgy在书写排序算法时，反复写了同一个逻辑多次。

然而这些都不重要，因为我们可以看出，azgy对编程语言尚未熟悉。因此在这个问题程序中，最严重的问题当属for循环中condition expression写反了。可以看出azgy的目的是希望for一直循环直到满足条件（也就是说相当于C中until的效果），然而实际上她的写法令循环一开始就跳出了。

可能大家觉得这个错误很可笑，然而我们不妨想一想：为什么初学者会搞不清楚在for循环中，condition expression到底该为真还是为假呢？这篇文章，就是解释这个问题的。我相信你们除非认真学过计算机科学，不然你们应该只会用，而不知其所以然。

注：本文已经对信息进行了筛选，希望大家不要跳步，而是阅读所有内容。编程语言最重要的就是思想。

## 道生一：编程语言的前身，flowchart

那么我们就按照我们一贯的理念，从0开始简洁地解释这个问题。

大家都知道，编程语言的流程控制分为三种结构：顺序、选择和循环。然而，大家有没有想过这三种结构是怎么来的？

实际上，这三种结构并不是显而易见的，而是在经过了长期的发展后最终定格的。而大家都知道，早在第一台计算机出现之前，类似编程语言的概念就已经以flowchart（流程图）的形式出现了，指导人们按照步骤做事。

最初的flowchart叫做flow process chart，由Frank和Lillian Gilbreth于1921年公开。这是当时flow process chart的样子：

![flow process chart](https://upload.wikimedia.org/wikipedia/commons/8/87/Subway_Fare_Card_Machine_Flow_Process_Chart.jpg)

大家可以注意到，在最初的flow process chart中已经有了流程控制的概念，有顺序、选择和循环结构。可以说，这三种结构自然而然就可以被想到，而且有了这三种结构，就可以表达一切逻辑顺序。

## 一生二：从flowchart到算法逻辑

Herman Goldstine和John von Neumann（即大家都知道的冯·诺伊曼）于1947年写了一篇未公开的报告"Planning and coding of problems for an electronic computing instrument"，开始使用flowchart来描述计算机算法。

随着第三代编程语言的出现，我们需要使用文字而不是流程图来描述算法，因为文字描述更加精确。这时候，我们就需要将流程图的概念使用文字来表达出来。

对于普通的顺序命令，当然是没有任何问题的；然而问题就在于，我们如何表现选择和循环呢？这里就牵扯到逻辑问题，我们需要逻辑工具。

Tony Hoare于1969年公开了一套理论，叫做Hoare logic，来精确地描述计算机算法的逻辑。在这里不深入讲解，给大家看一下选择和循环的逻辑表达，看不懂没关系（估计你也看不懂）。

### 从顺序结构到选择结构

逻辑描述：
`{P} if B then S else T endif {Q}`

逻辑表达：
`{B∧P}S{Q}, {¬B∧P}T{Q}`

大家可以注意到，**无论是for loop, do...until loop, switch, break, continue，归根结底都是选择结构**。

### 从选择结构到循环结构

逻辑描述：
`{P} while B do S done {¬B ∧ P}`

逻辑表达：
`{P∧B}S{P}`

大家可以注意到选择结构反而比顺序结构更简单了。简单而言，从逻辑的角度来说，循环结构的逻辑就是在循环前，条件B被满足；循环完毕的时候B不再被满足。

**因此大家可以注意到，从逻辑学的角度而言，循环结构需要的condition expression表现的是进入循环的条件，而不是跳出循环的条件。这也是azgy犯错的原因所在。**

**为了避免混乱，Go总是采用最标准的形式。因此在Go语言中，不存在do...until这种把condition expression反过来用的可能。大家仔细体会一下Go语言的设计思想。**

## 二生三：从算法逻辑到编程语言

我们有了指导思想，也有了具体的逻辑工具，接下来就要实际创造语言了。而根据Hoare logic被创造的语言，就是Guarded Command Language，由Edsger Dijkstra于2006年创造。

大家可能奇怪怎么是2006年创造的。实际上大家也知道，顺序、选择和循环这三种控制结构是自然而然就能想到的。然而，C语言这种语言也创造了一大堆乱七八糟的for, do...until, while, do...while, break, continue一类乱七八糟的控制语句。各位小伙伴如果有使用C的经验，有没有觉得容易混乱？

前面已经说过，无论顺序、选择和循环，归根结底其实都是选择，即满足一个条件，就执行某一个语句(if 条件 then 命令)。既然逻辑只有一个，我们为什么要创造一大堆多余的控制语句？

Guarded Command Language提供了一种工具语言来描述程序的算法。

### 从顺序、选择、循环三种结构合三为一：guarded command

如果上面的解释让你晕，编程逻辑发展到这里你应该就不晕了，因为就像hoare logic所证明的那样，一切结构控制从逻辑上讲归根结底都是一个if 条件 then 命令。那么，我们进一步将结构控制简化成一个概念行不行？

Guarded command:
`G → S`
其中G是guard，S是statement

于是一切结构控制都可以用一个非常简单的`G → S`来表示。也就是说，`if G then S`。

大家可以注意的是，这里的概念不再将G视作一个条件，而是一个guard。这下azgy所犯的错误就更加容易理解：一个guard，就是“保卫”的意思，这个保卫的作用就是，你让它同意，它才允许你执行它所保卫的代码。因此你无论怎么想，都是guard为真时才能执行代码。这样for中的语句肯定是为真时才能循环，azgy的认知混乱就不会再出现。

#### 使用guarded command表达的选择结构

```
if G0 → S0
| G1 → S1
...
| Gn → Sn
fi
```

这段guarded command的意思就是if G0 then S0 else if G1 then S1 else if ... else if Gn then Sn。从guarded command的角度来解释，选择结构就是一大堆guard(G0, G1, G2)，每个guard分别“保卫”一些语句。

#### 使用guarded command表达的循环结构

```
do G0 → S0
| G1 → S1
...
| Gn → Sn
od
```

即while G0 do S0，用选择结构表达就是if G0 then S0 if G0 then S0 if G0 then S0...

当大家将所有的控制语句都简化成一个if G then S(while G do S)时，各位还有可能出现azgy那样的真假混乱问题吗？

**Go的理论是从上述的Guarded Command Language衍生而来。大家可以在使用中体会Go的这个设计思想。**
