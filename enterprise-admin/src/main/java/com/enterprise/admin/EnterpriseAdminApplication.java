package com.enterprise.admin;

import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.transaction.annotation.EnableTransactionManagement;

/**
 * 企业级管理后台启动类
 *
 * @author enterprise
 * @since 2023-12-01
 */
@SpringBootApplication
@EnableTransactionManagement
@MapperScan("com.enterprise.system.mapper")
public class EnterpriseAdminApplication {

    public static void main(String[] args) {
        SpringApplication.run(EnterpriseAdminApplication.class, args);
        System.out.println("=================================");
        System.out.println("企业级管理后台启动成功！");
        System.out.println("=================================");
    }
}