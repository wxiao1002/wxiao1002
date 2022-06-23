package 数据结构与算法之美;

/**
 * @author wang xiao
 * @date 2022/6/23
 */
public class MergeSort {

    public void mergeSort (int arr,int low,int high){
        while (low <high){
            int mid = (low+high)/2;
            mergeSort(arr,low,mid);
            mergeSort(arr,mid+1,high);
        }
    }

    public void merge(int [] arr,int left,int mid, int right){
        int [] temp = new int[arr.length];
        int i = left;
        int j = mid+1;
        int index = left;

        while (i<= mid && j<= right){
            if (arr[i] <=  arr[j]){
                temp[index++] = arr[i++];
            }else {
                temp[index++] = arr[j++];
            }
        }

        while (i <= mid){
            temp[index++]= arr[i++];
        }
        while (j<right){
            temp[index++] = arr[j++];
        }
        for (int c=left;c<=right;c++){
            arr[c] = temp[c];
        }
    }
}
