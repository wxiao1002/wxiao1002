package excutor;

import com.google.common.util.concurrent.ListenableFuture;
import java.util.concurrent.Callable;
import java.util.concurrent.Executor;

/**
 * 具有监听功能的 executor
 * @author wangxiao
 * @date 2022/3/9
 */
public interface EngineListeningExecutor extends Executor {

    /**
     * 异步执行任务
     * @param task 任务
     * @return 结果
     */
    <T> ListenableFuture<T> executeAsync(Callable<T> task);

}
