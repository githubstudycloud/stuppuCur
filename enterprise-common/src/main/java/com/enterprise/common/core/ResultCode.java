package com.enterprise.common.core;

import lombok.AllArgsConstructor;
import lombok.Getter;

/**
 * 响应码枚举
 *
 * @author enterprise
 * @since 2023-12-01
 */
@Getter
@AllArgsConstructor
public enum ResultCode {

    /**
     * 成功
     */
    SUCCESS(200, "操作成功"),

    /**
     * 失败
     */
    ERROR(500, "操作失败"),

    /**
     * 参数错误
     */
    PARAM_ERROR(400, "参数错误"),

    /**
     * 未授权
     */
    UNAUTHORIZED(401, "未授权"),

    /**
     * 禁止访问
     */
    FORBIDDEN(403, "禁止访问"),

    /**
     * 资源不存在
     */
    NOT_FOUND(404, "资源不存在"),

    /**
     * 请求方法不允许
     */
    METHOD_NOT_ALLOWED(405, "请求方法不允许"),

    /**
     * 请求超时
     */
    REQUEST_TIMEOUT(408, "请求超时"),

    /**
     * 冲突
     */
    CONFLICT(409, "资源冲突"),

    /**
     * 服务器内部错误
     */
    INTERNAL_SERVER_ERROR(500, "服务器内部错误"),

    /**
     * 服务不可用
     */
    SERVICE_UNAVAILABLE(503, "服务不可用"),

    /**
     * 业务异常
     */
    BUSINESS_ERROR(1000, "业务异常"),

    /**
     * 数据不存在
     */
    DATA_NOT_FOUND(1001, "数据不存在"),

    /**
     * 数据已存在
     */
    DATA_ALREADY_EXISTS(1002, "数据已存在"),

    /**
     * 数据状态错误
     */
    DATA_STATUS_ERROR(1003, "数据状态错误"),

    /**
     * 文件上传失败
     */
    FILE_UPLOAD_ERROR(2001, "文件上传失败"),

    /**
     * 文件下载失败
     */
    FILE_DOWNLOAD_ERROR(2002, "文件下载失败"),

    /**
     * 文件不存在
     */
    FILE_NOT_FOUND(2003, "文件不存在"),

    /**
     * 文件格式不支持
     */
    FILE_FORMAT_NOT_SUPPORTED(2004, "文件格式不支持"),

    /**
     * 文件大小超限
     */
    FILE_SIZE_EXCEEDED(2005, "文件大小超限"),

    /**
     * 用户不存在
     */
    USER_NOT_FOUND(3001, "用户不存在"),

    /**
     * 用户名或密码错误
     */
    USERNAME_OR_PASSWORD_ERROR(3002, "用户名或密码错误"),

    /**
     * 用户已被禁用
     */
    USER_DISABLED(3003, "用户已被禁用"),

    /**
     * 用户已锁定
     */
    USER_LOCKED(3004, "用户已锁定"),

    /**
     * 用户已过期
     */
    USER_EXPIRED(3005, "用户已过期"),

    /**
     * 用户凭证已过期
     */
    USER_CREDENTIALS_EXPIRED(3006, "用户凭证已过期"),

    /**
     * 用户权限不足
     */
    USER_INSUFFICIENT_PERMISSIONS(3007, "用户权限不足"),

    /**
     * 验证码错误
     */
    CAPTCHA_ERROR(4001, "验证码错误"),

    /**
     * 验证码已过期
     */
    CAPTCHA_EXPIRED(4002, "验证码已过期"),

    /**
     * 短信验证码错误
     */
    SMS_CODE_ERROR(4003, "短信验证码错误"),

    /**
     * 短信验证码已过期
     */
    SMS_CODE_EXPIRED(4004, "短信验证码已过期"),

    /**
     * 邮箱验证码错误
     */
    EMAIL_CODE_ERROR(4005, "邮箱验证码错误"),

    /**
     * 邮箱验证码已过期
     */
    EMAIL_CODE_EXPIRED(4006, "邮箱验证码已过期");

    /**
     * 响应码
     */
    private final Integer code;

    /**
     * 响应消息
     */
    private final String message;
}