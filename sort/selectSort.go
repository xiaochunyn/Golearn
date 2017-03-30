package main

import "fmt"

func main() {
	a := []int{7, 3, 6, 9, 45, 31, 4, 6, 8, 5}
	sort(a)
	fmt.Printf("%v", a)
}

func sort(a []int) {
	var min int
	var pos int
	for i := 0; i < len(a); i++ {
		min = a[i]
		pos = i
		for j := i + 1; j < len(a); j++ {
			if a[j] < min {
				min = a[j]
				pos = j
			}
		}

		a[pos] = a[i]
		a[i] = min
	}
}

/**
 * 简单选择排序：
 * 基本思想：
 * 	在要排序的一组数中，选出最小的一个数与第一个位置的数交换
 * 	然后在剩下的数当中再找最小的数与第二个位置的数交换，
 * 	如此循环到倒数第二个数和最后一个数比较为止。
 *
 * 	选择排序的主要优点与数据移动有关，
 * 	如果某个元素位于正确的位置，则它不会被移动。
 * 	选择排序每次交换一对元素，他们当中至少有一个将被移到其最终位置上，
 * 	因此，对n个元素的序列进行排序总共进行至多n-1次交换。
 *  在所有的完全依靠交换去移动元的排序方法中，选择排序属于非常好的一种，
 *  因为它在跑完一趟之后才进行元素的移动。

实例分析
以数组 arr = [8, 5, 2, 6, 9, 3, 1, 4, 0, 7] 为例，先直观看一下每一步的变化，后面再介绍细节
第一次从数组 [8, 5, 2, 6, 9, 3, 1, 4, 0, 7] 中找到最小的数 0，放到数组的最前面（与第一个元素进行交换）：

                               min
                                ↓
8   5   2   6   9   3   1   4   0   7
↑                               ↑
└───────────────────────────────┘
交换后：
1
0   5   2   6   9   3   1   4   8   7

在剩余的序列中 [5, 2, 6, 9, 3, 1, 4, 8, 7] 中找到最小的数 1，与该序列的第一个个元素进行位置交换：

                       min
                        ↓
0   5   2   6   9   3   1   4   8   7
    ↑                   ↑
    └───────────────────┘
交换后：
1
0   1   2   6   9   3   5   4   8   7

在剩余的序列中 [2, 6, 9, 3, 5, 4, 8, 7] 中找到最小的数 2，与该序列的第一个个元素进行位置交换（实际上不需要交换）：
       min
        ↓
0   1   2   6   9   3   5   4   8   7
        ↑
重复上述过程，直到最后一个元素就完成了排序。

                   min
                    ↓
0   1   2   6   9   3   5   4   8   7
            ↑       ↑
            └───────┘
                           min
                            ↓
0   1   2   3   9   6   5   4   8   7
                ↑           ↑
                └───────────┘
                       min
                        ↓
0   1   2   3   4   6   5   9   8   7
                    ↑   ↑
                    └───┘
                       min
                        ↓
0   1   2   3   4   5   6   9   8   7
                        ↑
                                   min
                                    ↓
0   1   2   3   4   5   6   9   8   7
                            ↑       ↑
                            └───────┘
                               min
                                ↓
0   1   2   3   4   5   6   7   8   9
                                ↑
                                   min
                                    ↓
0   1   2   3   4   5   6   7   8   9
                                    ↑

*/