package test;

public class Graph{

    private int v;

    private LinkedList<Integer> adj [] ;// 领接表


   public Graph (int v) {
       this.v = v;
       adj = new LinkedList[v];
       for (int i =0;i<v;i++){
           adj[i] = new LinkedList();
       }
   }


   public void bfs (int s,int t){

       if (s== t){
           return;
       }
       Queue<Integer> searchQueue = new LinkedList<>();
       searchQueue.add(s);
       while (!searchQueue.isEmpty()){
           int temp = searchQueue.poll();
           LinkedList<Integer> integerLinkedList = adj[temp];
           for (Integer w : integerLinkedList){
               if (w == t){
                   return;
               }
               searchQueue.add(w);
           }
       }
   }
}