package excutor;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

/**
 * 线程池
 * @author wangxiao
 * @date 2022/3/12
 */
@Component
public class EngineExecutorService extends BaseEngineListeningExecutor {

    @Value("${api_server.executor.poll_size:5}")
    private int poolSize;

    @Override
    protected int getThreadPollSize() {
        return poolSize;
    }
}
