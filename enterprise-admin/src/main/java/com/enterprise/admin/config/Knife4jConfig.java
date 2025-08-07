package com.enterprise.admin.config;

import io.swagger.v3.oas.models.OpenAPI;
import io.swagger.v3.oas.models.info.Contact;
import io.swagger.v3.oas.models.info.Info;
import io.swagger.v3.oas.models.info.License;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

/**
 * Knife4j API文档配置
 *
 * @author enterprise
 * @since 2023-12-01
 */
@Configuration
public class Knife4jConfig {

    @Bean
    public OpenAPI customOpenAPI() {
        return new OpenAPI()
                .info(new Info()
                        .title("企业级管理后台API文档")
                        .description("基于Spring Boot 3.x的企业级多项目管理基座")
                        .version("1.0.0")
                        .contact(new Contact()
                                .name("Enterprise Team")
                                .email("enterprise@example.com")
                                .url("https://github.com/enterprise"))
                        .license(new License()
                                .name("MIT License")
                                .url("https://opensource.org/licenses/MIT")));
    }
}