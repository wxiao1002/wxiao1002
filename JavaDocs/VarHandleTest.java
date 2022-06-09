public class VarHandleTest {
    volatile int x;

    public static VarHandle varHandle;

    static {
        varHandle  = MethodHandles.lookup().
                findVarHandle(VarHandleTest.class, "x", int.class)
    }

    public static void main(String[] args) {

    }
}