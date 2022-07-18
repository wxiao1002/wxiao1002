package com.baoan.sk.jvm;

/**
 * 对象的一次 自我拯救
 * @author wang xiao
 * @date 2022/7/18
 */
public class FinalizeEscapeGC {

    public static FinalizeEscapeGC SAVE_HOOK = null;

    public void isActive () {
        System.out.println("I`m alive");
    }

    @Override
    protected void finalize() throws Throwable {
        super.finalize();
        System.out.println("finalize exec");
        SAVE_HOOK = this;
    }

    public static void main(String[] args) throws InterruptedException {
        SAVE_HOOK = new FinalizeEscapeGC();
        SAVE_HOOK = null;
        System.gc();
        Thread.sleep(500);

        if (SAVE_HOOK != null){
            SAVE_HOOK.isActive();
        }else {
            System.out.println("dead");
        }

    }
}
