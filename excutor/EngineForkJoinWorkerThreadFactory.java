package excutor;

import java.util.concurrent.ForkJoinPool;
import java.util.concurrent.ForkJoinWorkerThread;
import java.util.concurrent.atomic.AtomicLong;

/**
 *  fork join 工厂 类
 * @author wangxiao
 * @date 2022/3/8
 */
public class EngineForkJoinWorkerThreadFactory implements ForkJoinPool.ForkJoinWorkerThreadFactory{

    private final String namePrefix;
    private final AtomicLong threadNumber = new AtomicLong(1);

    public EngineForkJoinWorkerThreadFactory(String namePrefix) {
        this.namePrefix = namePrefix;
    }

    @Override
    public final ForkJoinWorkerThread newThread(ForkJoinPool pool) {
        ForkJoinWorkerThread thread = ForkJoinPool.defaultForkJoinWorkerThreadFactory.newThread(pool);
        thread.setContextClassLoader(this.getClass().getClassLoader());
        thread.setName(namePrefix +"-"+thread.getPoolIndex()+"-"+threadNumber.getAndIncrement());
        return thread;
    }
}
