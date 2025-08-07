# 使用OpenJDK 17作为基础镜像
FROM openjdk:17-jdk-slim

# 设置工作目录
WORKDIR /app

# 设置环境变量
ENV JAVA_OPTS="-Xms512m -Xmx1024m -Djava.security.egd=file:/dev/./urandom"

# 复制Maven配置文件
COPY pom.xml .
COPY enterprise-common/pom.xml enterprise-common/
COPY enterprise-security/pom.xml enterprise-security/
COPY enterprise-system/pom.xml enterprise-system/
COPY enterprise-admin/pom.xml enterprise-admin/

# 复制源代码
COPY enterprise-common/src enterprise-common/src
COPY enterprise-security/src enterprise-security/src
COPY enterprise-system/src enterprise-system/src
COPY enterprise-admin/src enterprise-admin/src

# 安装Maven并构建项目
RUN apt-get update && apt-get install -y maven && \
    mvn clean package -DskipTests && \
    apt-get remove -y maven && \
    apt-get autoremove -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# 复制构建好的jar文件
COPY enterprise-admin/target/enterprise-admin-*.jar app.jar

# 暴露端口
EXPOSE 8080

# 设置启动命令
ENTRYPOINT ["sh", "-c", "java $JAVA_OPTS -jar app.jar"]