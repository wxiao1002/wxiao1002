package com.iwhalecloud.bassc.lcdp.engine.node.ds.sql;

import java.util.ArrayList;
import java.util.List;

/**
 * @author wang xiao
 * @date 2022/8/8
 */
public class ParameterMappingTokenHandler implements TokenHandler{

    // 将解析出来的字段放入到集合中，面向对象的思想：把字段封装为对象ParameterMapping
    // 方便后续对占位符参数的使用
    private List<ParameterMapping> parameterMappings = new ArrayList<ParameterMapping>();

    // ParameterMapping.context是参数名称 #{id}、#{username}中的 id、username
    // 注意！注意！注意！=========================================
    // TODO 将占位符中的变量用 ？代替
    public String handleToken(String content) {
        parameterMappings.add(buildParameterMapping(content));
        return "?";
    }

    // 将占位符中的变量封装为对象
    private ParameterMapping buildParameterMapping(String content) {
        ParameterMapping parameterMapping = new ParameterMapping(content);
        return parameterMapping;
    }

    // 成员变量get方法
    public List<ParameterMapping> getParameterMappings() {
        return parameterMappings;
    }

    // 成员变量set方法
    public void setParameterMappings(List<ParameterMapping> parameterMappings) {
        this.parameterMappings = parameterMappings;
    }
}
