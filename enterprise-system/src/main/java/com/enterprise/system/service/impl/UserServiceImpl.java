package com.enterprise.system.service.impl;

import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.enterprise.common.core.ResultCode;
import com.enterprise.common.exception.BusinessException;
import com.enterprise.common.utils.AssertUtils;
import com.enterprise.system.entity.User;
import com.enterprise.system.mapper.UserMapper;
import com.enterprise.system.service.UserService;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.time.LocalDateTime;

/**
 * 用户服务实现类
 *
 * @author enterprise
 * @since 2023-12-01
 */
@Slf4j
@Service
@RequiredArgsConstructor
public class UserServiceImpl extends ServiceImpl<UserMapper, User> implements UserService {

    private final PasswordEncoder passwordEncoder;

    @Override
    public User getByUsername(String username) {
        AssertUtils.notEmpty(username, "用户名不能为空");
        return baseMapper.selectByUsername(username);
    }

    @Override
    public User getByEmail(String email) {
        AssertUtils.notEmpty(email, "邮箱不能为空");
        return baseMapper.selectByEmail(email);
    }

    @Override
    public User getByPhone(String phone) {
        AssertUtils.notEmpty(phone, "手机号不能为空");
        return baseMapper.selectByPhone(phone);
    }

    @Override
    @Transactional(rollbackFor = Exception.class)
    public boolean createUser(User user) {
        AssertUtils.notNull(user, "用户信息不能为空");
        AssertUtils.notEmpty(user.getUsername(), "用户名不能为空");
        AssertUtils.notEmpty(user.getPassword(), "密码不能为空");

        // 检查用户名是否已存在
        if (getByUsername(user.getUsername()) != null) {
            throw new BusinessException(ResultCode.DATA_ALREADY_EXISTS, "用户名已存在");
        }

        // 检查邮箱是否已存在
        if (user.getEmail() != null && getByEmail(user.getEmail()) != null) {
            throw new BusinessException(ResultCode.DATA_ALREADY_EXISTS, "邮箱已存在");
        }

        // 检查手机号是否已存在
        if (user.getPhone() != null && getByPhone(user.getPhone()) != null) {
            throw new BusinessException(ResultCode.DATA_ALREADY_EXISTS, "手机号已存在");
        }

        // 加密密码
        user.setPassword(passwordEncoder.encode(user.getPassword()));
        user.setCreateTime(LocalDateTime.now());
        user.setUpdateTime(LocalDateTime.now());

        return save(user);
    }

    @Override
    @Transactional(rollbackFor = Exception.class)
    public boolean updateUser(User user) {
        AssertUtils.notNull(user, "用户信息不能为空");
        AssertUtils.notNull(user.getId(), "用户ID不能为空");

        User existingUser = getById(user.getId());
        if (existingUser == null) {
            throw new BusinessException(ResultCode.DATA_NOT_FOUND, "用户不存在");
        }

        // 检查用户名是否已被其他用户使用
        if (user.getUsername() != null && !user.getUsername().equals(existingUser.getUsername())) {
            User userWithSameUsername = getByUsername(user.getUsername());
            if (userWithSameUsername != null && !userWithSameUsername.getId().equals(user.getId())) {
                throw new BusinessException(ResultCode.DATA_ALREADY_EXISTS, "用户名已存在");
            }
        }

        // 检查邮箱是否已被其他用户使用
        if (user.getEmail() != null && !user.getEmail().equals(existingUser.getEmail())) {
            User userWithSameEmail = getByEmail(user.getEmail());
            if (userWithSameEmail != null && !userWithSameEmail.getId().equals(user.getId())) {
                throw new BusinessException(ResultCode.DATA_ALREADY_EXISTS, "邮箱已存在");
            }
        }

        // 检查手机号是否已被其他用户使用
        if (user.getPhone() != null && !user.getPhone().equals(existingUser.getPhone())) {
            User userWithSamePhone = getByPhone(user.getPhone());
            if (userWithSamePhone != null && !userWithSamePhone.getId().equals(user.getId())) {
                throw new BusinessException(ResultCode.DATA_ALREADY_EXISTS, "手机号已存在");
            }
        }

        user.setUpdateTime(LocalDateTime.now());
        return updateById(user);
    }

    @Override
    @Transactional(rollbackFor = Exception.class)
    public boolean deleteUser(Long id) {
        AssertUtils.notNull(id, "用户ID不能为空");

        User user = getById(id);
        if (user == null) {
            throw new BusinessException(ResultCode.DATA_NOT_FOUND, "用户不存在");
        }

        return removeById(id);
    }

    @Override
    @Transactional(rollbackFor = Exception.class)
    public boolean enableUser(Long id) {
        AssertUtils.notNull(id, "用户ID不能为空");

        User user = getById(id);
        if (user == null) {
            throw new BusinessException(ResultCode.DATA_NOT_FOUND, "用户不存在");
        }

        user.setStatus(1);
        user.setUpdateTime(LocalDateTime.now());
        return updateById(user);
    }

    @Override
    @Transactional(rollbackFor = Exception.class)
    public boolean disableUser(Long id) {
        AssertUtils.notNull(id, "用户ID不能为空");

        User user = getById(id);
        if (user == null) {
            throw new BusinessException(ResultCode.DATA_NOT_FOUND, "用户不存在");
        }

        user.setStatus(0);
        user.setUpdateTime(LocalDateTime.now());
        return updateById(user);
    }

    @Override
    @Transactional(rollbackFor = Exception.class)
    public boolean resetPassword(Long id, String newPassword) {
        AssertUtils.notNull(id, "用户ID不能为空");
        AssertUtils.notEmpty(newPassword, "新密码不能为空");

        User user = getById(id);
        if (user == null) {
            throw new BusinessException(ResultCode.DATA_NOT_FOUND, "用户不存在");
        }

        user.setPassword(passwordEncoder.encode(newPassword));
        user.setUpdateTime(LocalDateTime.now());
        return updateById(user);
    }

    @Override
    @Transactional(rollbackFor = Exception.class)
    public boolean changePassword(Long id, String oldPassword, String newPassword) {
        AssertUtils.notNull(id, "用户ID不能为空");
        AssertUtils.notEmpty(oldPassword, "旧密码不能为空");
        AssertUtils.notEmpty(newPassword, "新密码不能为空");

        User user = getById(id);
        if (user == null) {
            throw new BusinessException(ResultCode.DATA_NOT_FOUND, "用户不存在");
        }

        // 验证旧密码
        if (!passwordEncoder.matches(oldPassword, user.getPassword())) {
            throw new BusinessException(ResultCode.USERNAME_OR_PASSWORD_ERROR, "旧密码错误");
        }

        user.setPassword(passwordEncoder.encode(newPassword));
        user.setUpdateTime(LocalDateTime.now());
        return updateById(user);
    }

    @Override
    @Transactional(rollbackFor = Exception.class)
    public boolean updateLastLoginInfo(Long id, String loginIp) {
        AssertUtils.notNull(id, "用户ID不能为空");

        User user = getById(id);
        if (user == null) {
            throw new BusinessException(ResultCode.DATA_NOT_FOUND, "用户不存在");
        }

        user.setLastLoginTime(LocalDateTime.now());
        user.setLastLoginIp(loginIp);
        return updateById(user);
    }
}