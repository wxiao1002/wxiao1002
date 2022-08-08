package com.iwhalecloud.bassc.lcdp.engine.node.ds.sql;

/**
 * @author wang xiao
 * @date 2022/8/8
 */
public class ParameterMapping {

    private String content;

    // 有参构造
    public ParameterMapping(String content) {
        this.content = content;
    }
    // 成员变量get方法
    public String getContent() {
        return content;
    }
    // 成员变量set方法
    public void setContent(String content) {
        this.content = content;
    }
}
