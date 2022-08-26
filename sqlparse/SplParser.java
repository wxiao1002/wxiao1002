package com.iwhalecloud.bassc.lcdp.engine.node.ds.sql;




import java.util.List;

/**
 * @author wang xiao
 * @date 2022/8/8
 */
public class SplParser {


    /**
     * 完成对#{}的解析工作：1.将#{}使用？进行代替，2.解析出#{}里面的值进行存储
     * @param sql
     * @return
     */
    private String getBoundSql(String sql) {
        //标记处理类：配置标记解析器来完成对占位符的解析处理工作
        ParameterMappingTokenHandler parameterMappingTokenHandler = new ParameterMappingTokenHandler();

        // 工具类：3个参数：占位符的 开始标记、结束标记、替换为指定 符号'?' 的处理工具
        GenericTokenParser genericTokenParser = new GenericTokenParser("#{", "}", parameterMappingTokenHandler);

        //解析出来的sql
        String parseSql = genericTokenParser.parse(sql);
        System.out.println(parseSql);

        //#{}里面解析出来的参数名称
        List<ParameterMapping> parameterMappings = parameterMappingTokenHandler.getParameterMappings();
    }


    public static void main(String[] args) {
        new SplParser().getBoundSql("select * from t_student where id = #{id}");
    }

}
