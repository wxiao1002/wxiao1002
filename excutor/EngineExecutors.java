package excutor;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.ForkJoinPool;

/**
 * @author wang xiao
 * @date 2022/8/3
 */
public class EngineExecutors {

    /**
     * 创建 线程 工厂
     * @param parallelism 目标并发级别
     * @param namePrefix 用于定义线程名称
     * @return 线程工厂
     */
    public static ExecutorService newWorkStealingPool(int parallelism, String namePrefix) {
        return new ForkJoinPool(parallelism,
                new EngineForkJoinWorkerThreadFactory(namePrefix),
                null, true);
    }

    public static ExecutorService newWorkStealingPool(int parallelism, Class<?> clazz) {
        return newWorkStealingPool(parallelism, clazz.getSimpleName());
    }
}
