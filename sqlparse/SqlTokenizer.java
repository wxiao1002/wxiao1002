package com.iwhalecloud.bassc.lcdp.engine.node.ds.sql;

import java.util.Iterator;

/**
 * @author wang xiao
 * @date 2022/8/8
 */
public class SqlTokenizer implements Iterator<String> {

    String[] tokens;

    int index = 0;

    public SqlTokenizer(String sql) throws RuntimeException {
        sql = sql.trim();
        //去除多余的空格
        sql = sql.replace(";", " ;").replaceAll("\\s+", " ");
        //分词
        tokens = sql.split(" ");

    }


    @Override

    public boolean hasNext() {
        return index < tokens.length;
    }

    @Override
    public String next() {
        return tokens[index++];
    }
}
