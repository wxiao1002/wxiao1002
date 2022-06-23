package 数据结构与算法之美;

/**
 * @author wang xiao
 * @date 2022/6/23
 */
public class QuickSort {

    /**
     * 快速排序的思路
     * 找到一个基准元素pivot 从前往后 从后往前进行比较 然后交换
     * 最重要的是分区函数 partition
     */

    public void quickSort(int [] arr,int low,int high){

        if (low >= high){
            return;
        }
        int mid = partition(arr,low,high);
        quickSort(arr,low,mid-1);
        quickSort(arr,mid+1,high);
    }

    public int partition (int [] arr,int low,int high){
        int pivot = arr[low];
        while (low <high){
            while (low <high && arr[high] >= pivot){
                high --;
            }
            arr[low] = arr[high];
            while (low<high && arr[low] <= pivot){
                low++;
            }
            arr[high] = arr[low];

        }
        arr[low] = pivot;
        return  low;
    }


}
