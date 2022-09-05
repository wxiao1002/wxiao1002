package com.github.xiao.JMH;

import org.openjdk.jmh.annotations.*;
import org.openjdk.jmh.results.format.ResultFormatType;
import org.openjdk.jmh.runner.Runner;
import org.openjdk.jmh.runner.RunnerException;
import org.openjdk.jmh.runner.options.Options;
import org.openjdk.jmh.runner.options.OptionsBuilder;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.concurrent.TimeUnit;


@BenchmarkMode(Mode.AverageTime) // 测试完成时间
@OutputTimeUnit(TimeUnit.NANOSECONDS)
@Warmup(iterations = 2, time = 1, timeUnit = TimeUnit.SECONDS) // 预热 2 轮，每次 1s
@Measurement(iterations = 5, time = 3, timeUnit = TimeUnit.SECONDS) // 测试 5 轮，每次 3s
@Fork(1) // fork 1 个线程
@State(Scope.Thread) // 每个测试线程一个实例
@RestController
@RequestMapping("/log")
public class JmhLogPrint {

    private final Logger log = LoggerFactory.getLogger(JmhLogPrint.class);

    public static void main(String[] args) throws RunnerException {
        //1、启动基准测试：输出json结果文件（用于查看可视化图）
        Options opt = new OptionsBuilder()
                .include(JmhLogPrint.class.getSimpleName()) //要导入的测试类
                .result("D:\\JmhLogPrint.json") //输出测试结果的json文件
                .resultFormat(ResultFormatType.JSON)//格式化json文件
                .build();

        //2、执行测试
        new Runner(opt).run();
    }

    @Benchmark
    public void appendLogPrint() {
        StringBuilder sb = new StringBuilder();
        sb.append("Hello, ");
        sb.append("Java");
        sb.append(".");
        sb.append("Hello, ");
        sb.append("Redis");
        sb.append(".");
        sb.append("Hello, ");
        sb.append("MySQL");
        sb.append(".");
        log.info(sb.toString());
    }

    @Benchmark
    public void logPrintWithPlaceholder() {
        log.info("Hello, {}.Hello, {}.Hello, {}.", "Java", "Redis", "MySQL");
    }
}

