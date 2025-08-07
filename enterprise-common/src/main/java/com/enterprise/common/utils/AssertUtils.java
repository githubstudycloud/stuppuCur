package com.enterprise.common.utils;

import com.enterprise.common.core.ResultCode;
import com.enterprise.common.exception.BusinessException;
import org.springframework.util.StringUtils;

import java.util.Collection;

/**
 * 断言工具类
 *
 * @author enterprise
 * @since 2023-12-01
 */
public class AssertUtils {

    /**
     * 断言对象不为空
     */
    public static void notNull(Object object, String message) {
        if (object == null) {
            throw new BusinessException(ResultCode.PARAM_ERROR, message);
        }
    }

    /**
     * 断言对象不为空
     */
    public static void notNull(Object object, ResultCode resultCode) {
        if (object == null) {
            throw new BusinessException(resultCode);
        }
    }

    /**
     * 断言对象不为空
     */
    public static void notNull(Object object, ResultCode resultCode, String message) {
        if (object == null) {
            throw new BusinessException(resultCode, message);
        }
    }

    /**
     * 断言字符串不为空
     */
    public static void notEmpty(String str, String message) {
        if (!StringUtils.hasText(str)) {
            throw new BusinessException(ResultCode.PARAM_ERROR, message);
        }
    }

    /**
     * 断言字符串不为空
     */
    public static void notEmpty(String str, ResultCode resultCode) {
        if (!StringUtils.hasText(str)) {
            throw new BusinessException(resultCode);
        }
    }

    /**
     * 断言字符串不为空
     */
    public static void notEmpty(String str, ResultCode resultCode, String message) {
        if (!StringUtils.hasText(str)) {
            throw new BusinessException(resultCode, message);
        }
    }

    /**
     * 断言集合不为空
     */
    public static void notEmpty(Collection<?> collection, String message) {
        if (collection == null || collection.isEmpty()) {
            throw new BusinessException(ResultCode.PARAM_ERROR, message);
        }
    }

    /**
     * 断言集合不为空
     */
    public static void notEmpty(Collection<?> collection, ResultCode resultCode) {
        if (collection == null || collection.isEmpty()) {
            throw new BusinessException(resultCode);
        }
    }

    /**
     * 断言集合不为空
     */
    public static void notEmpty(Collection<?> collection, ResultCode resultCode, String message) {
        if (collection == null || collection.isEmpty()) {
            throw new BusinessException(resultCode, message);
        }
    }

    /**
     * 断言条件为真
     */
    public static void isTrue(boolean condition, String message) {
        if (!condition) {
            throw new BusinessException(ResultCode.PARAM_ERROR, message);
        }
    }

    /**
     * 断言条件为真
     */
    public static void isTrue(boolean condition, ResultCode resultCode) {
        if (!condition) {
            throw new BusinessException(resultCode);
        }
    }

    /**
     * 断言条件为真
     */
    public static void isTrue(boolean condition, ResultCode resultCode, String message) {
        if (!condition) {
            throw new BusinessException(resultCode, message);
        }
    }

    /**
     * 断言条件为假
     */
    public static void isFalse(boolean condition, String message) {
        if (condition) {
            throw new BusinessException(ResultCode.PARAM_ERROR, message);
        }
    }

    /**
     * 断言条件为假
     */
    public static void isFalse(boolean condition, ResultCode resultCode) {
        if (condition) {
            throw new BusinessException(resultCode);
        }
    }

    /**
     * 断言条件为假
     */
    public static void isFalse(boolean condition, ResultCode resultCode, String message) {
        if (condition) {
            throw new BusinessException(resultCode, message);
        }
    }

    /**
     * 断言对象相等
     */
    public static void equals(Object obj1, Object obj2, String message) {
        if (!obj1.equals(obj2)) {
            throw new BusinessException(ResultCode.PARAM_ERROR, message);
        }
    }

    /**
     * 断言对象相等
     */
    public static void equals(Object obj1, Object obj2, ResultCode resultCode) {
        if (!obj1.equals(obj2)) {
            throw new BusinessException(resultCode);
        }
    }

    /**
     * 断言对象相等
     */
    public static void equals(Object obj1, Object obj2, ResultCode resultCode, String message) {
        if (!obj1.equals(obj2)) {
            throw new BusinessException(resultCode, message);
        }
    }

    /**
     * 断言对象不相等
     */
    public static void notEquals(Object obj1, Object obj2, String message) {
        if (obj1.equals(obj2)) {
            throw new BusinessException(ResultCode.PARAM_ERROR, message);
        }
    }

    /**
     * 断言对象不相等
     */
    public static void notEquals(Object obj1, Object obj2, ResultCode resultCode) {
        if (obj1.equals(obj2)) {
            throw new BusinessException(resultCode);
        }
    }

    /**
     * 断言对象不相等
     */
    public static void notEquals(Object obj1, Object obj2, ResultCode resultCode, String message) {
        if (obj1.equals(obj2)) {
            throw new BusinessException(resultCode, message);
        }
    }
}