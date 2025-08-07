package com.enterprise.system.service;

import com.baomidou.mybatisplus.extension.service.IService;
import com.enterprise.system.entity.User;

/**
 * 用户服务接口
 *
 * @author enterprise
 * @since 2023-12-01
 */
public interface UserService extends IService<User> {

    /**
     * 根据用户名查询用户
     */
    User getByUsername(String username);

    /**
     * 根据邮箱查询用户
     */
    User getByEmail(String email);

    /**
     * 根据手机号查询用户
     */
    User getByPhone(String phone);

    /**
     * 创建用户
     */
    boolean createUser(User user);

    /**
     * 更新用户
     */
    boolean updateUser(User user);

    /**
     * 删除用户
     */
    boolean deleteUser(Long id);

    /**
     * 启用用户
     */
    boolean enableUser(Long id);

    /**
     * 禁用用户
     */
    boolean disableUser(Long id);

    /**
     * 重置密码
     */
    boolean resetPassword(Long id, String newPassword);

    /**
     * 修改密码
     */
    boolean changePassword(Long id, String oldPassword, String newPassword);

    /**
     * 更新最后登录信息
     */
    boolean updateLastLoginInfo(Long id, String loginIp);
}