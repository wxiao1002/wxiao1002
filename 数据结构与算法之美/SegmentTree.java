package lc;

public class SegmentTree {

    private int m;
    public Segment [] segments ;

    public SegmentTree(int m) {
        this.segments = new Segment[4*m];
        this.m = m;
        buildSegmentTree(0,m);
    }
    private void buildSegmentTree(int i,int left,int right) {
        segments[i] = new Segment(left,m);
        if (left == right){
            return;
        }
        int mid  = (left+right)/2;
        buildSegmentTree(2i,left,mid);
        buildSegmentTree(2i+1,mid+1,right);
    }

    public void add(int data){
        int left = 0;
        int right = m;
        int i = 1;
        while (left != right){
            segments[i].count++;
            mid = (left+mid)/2;
            if (data > mid){
                i = 2i+1;
                left = mid+1;
            }else {
                i= 2i;
                right = mid;
            }
        }
        // 最后一个节点
        segments[i]++
    }

    public void remove(int data){
        int left = 0;
        int right = m;
        int i = 1;
        while (left != right){
            segments[i].count--;
            mid = (left+mid)/2;
            if (data > mid){
                i = 2i+1;
                left = mid+1;
            }else {
                i= 2i;
                right = mid;
            }
        }
        // 最后一个节点
        segments[i]--
    }

    public int count(int left,int right){
        return count(left,right,1);
    }

    private int count(int left,int rigth,int i){
        if (segments[i].left == left&& segments[i].right == rigth){
            return segments[i].count;
        }
        int mid = (segments[i].left+ segments[j].right)/2;
        if (left > mid){
            // 整个区间在左边
            return count(left,rigth,2i+1);
        }else if (rigth < mid){
            // 整个区间在右边
            return count(left,rigth,2i)
        }else {
            // 区间在中间
            return count(left,mid,2i)+ count(mid+1,rigth,2i+1)
        }
    }
    class Segment {

        private int left;
        private int right;

        private int count;

        public Segment(int left, int right) {
            this.left = left;
            this.right = right;
            this.count = 0;
        }
    }
}
