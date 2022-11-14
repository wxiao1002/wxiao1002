package com.baidu.ace.so.jmh;



import com.alibaba.fastjson2.JSONArray;
import com.alibaba.fastjson2.JSONPath;
import com.alibaba.fastjson2.JSONReader;
import org.openjdk.jmh.annotations.*;
import org.openjdk.jmh.infra.Blackhole;
import org.openjdk.jmh.results.format.ResultFormatType;
import org.openjdk.jmh.runner.Runner;
import org.openjdk.jmh.runner.RunnerException;
import org.openjdk.jmh.runner.options.Options;
import org.openjdk.jmh.runner.options.OptionsBuilder;

import java.util.concurrent.TimeUnit;


@BenchmarkMode(Mode.AverageTime) // 测试完成时间
@OutputTimeUnit(TimeUnit.NANOSECONDS)
@Warmup(iterations = 2, time = 1, timeUnit = TimeUnit.SECONDS) // 预热 2 轮，每次 1s
@Measurement(iterations = 5, time = 3, timeUnit = TimeUnit.SECONDS) // 测试 5 轮，每次 3s
@Fork(3) //Fork出3个线程来测试
@State(Scope.Thread) // 每个测试线程分配1个实例
public class JsonPathJmhTest {
    @Param({"2", "10", "100", "1000"})
    private int count; //指定添加元素的不同个数，便于分析结果

    public static void main(String[] args) throws RunnerException {
        //1、启动基准测试：输出普通文件
//        Options opt = new OptionsBuilder()
//                .include(ArrayAndLinkedJmhTest.class.getSimpleName()) //要导入的测试类
//                .output("C:\\Users\\Administrator\\Desktop\\StringAppendJmhTest.log") //输出测试结果的普通txt文件
//                .build();


        //1、启动基准测试：输出json结果文件（用于查看可视化图）
        Options opt = new OptionsBuilder()
                .include(com.baidu.ace.so.jmh.JsonPathJmhTest.class.getSimpleName()) //要导入的测试类
                .result("D:\\JSONPATH.json") //输出测试结果的json文件
                .resultFormat(ResultFormatType.JSON)//格式化json文件
                .build();

        //2、执行测试
        new Runner(opt).run();
    }

    @Setup(Level.Trial) // 初始化方法，在全部Benchmark运行之前进行
    public void init() {
        System.out.println("Start...");
    }



    @Benchmark
    public void jsonPath(Blackhole blackhole) {
        JSONArray jsonArray = JSONArray.parseArray("[{\"id\":1,\"name\":\"123\"},{\"id\":1,\"name\":\"345\"},{\"id\":2,\"name\":\"567\"},{\"id\":2,\"name\":\"78\"}]");
        for (int i = 0; i < count; i++) {
            // random ?
            JSONPath of = JSONPath.of("$.[0].id");
            // this need cache?
            of.extract(JSONReader.ofJSONB(jsonArray.toJSONBBytes()));
        }
        blackhole.consume(jsonArray);
    }

    @TearDown(Level.Trial) // 结束方法，在全部Benchmark运行之后进行
    public void clear() {
        System.out.println("End...");
    }

}
