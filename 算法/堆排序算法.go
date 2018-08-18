package main

import (
    "fmt"
    "math/rand"
)

var num int //用于统计次数
//堆排序,当0下标不参与排序时,左子节点角标为2*k(k<<1),右子节点角标为2*k+1,父节点角标为k/2(k>>1)
func HeapSort(s []int) {
    end := len(s) - 1 //s[0]不用，实际元素数量和最后一个元素的角标都为N
    //从最后一个非叶子节点开始, 向上构造最大堆(让最大值到根节点)
    //初始最大值
    for k := end>>1; k >= 1; k-- {
        num++
        sink(s, k, end)
    }
    //下沉排序
    //将最大值下沉到队列尾部, 即end, 然后将end排除在排序队列外,即 1到 end -1
    //直到end为1时终止
    for end > 1 {
        num++
        swap(s, 1, end) //将大的放在数组后面，升序排序
        end--
        sink(s, 1, end) //从根的左子节点开始找最大值, 并移动到根
    }
}
 
//下沉（由上至下的堆有序化）(二叉树原理)
func sink(s []int, k, end int) {
    for i := k<<1; i <= end; i <<= 2{
        num++
        if i < end && s[i+1] > s[i] { //选择较大的子节点
            i++
        }
        if s[k] >= s[i] { //没下沉到底就构造好堆了
            break
        }
        swap(s, k, i)
    }
}
 
func swap(s []int, i int, j int) {
    s[i], s[j] = s[j], s[i]
}

func main() {
    s := make([]int, 101)
    s[0] = 0
    for i := 100; i > 0; i-- {
        s = append(s, rand.Intn(10000))
    }
    fmt.Println(s[1:])
    HeapSort(s)
    fmt.Println(s[1:], len(s), num)
}
