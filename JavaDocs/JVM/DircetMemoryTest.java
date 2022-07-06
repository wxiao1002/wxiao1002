package tes;

public class DircetMemoryTest {

    private static int _1M = 1024*1024;
    public static void main(String[] args) {
        Field declaredField = Unsafe.class.getDeclaredFields()[0];
        declaredField.setAccessible(true);
        Unsafe o = (Unsafe)declaredField.get(null);
        o.allocateMemory(_1M);
    }
}