package com.enterprise.admin.controller;

import com.enterprise.common.core.Result;
import com.enterprise.common.utils.AssertUtils;
import com.enterprise.security.utils.JwtUtils;
import com.enterprise.system.entity.User;
import com.enterprise.system.service.UserService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.web.bind.annotation.*;

import java.util.HashMap;
import java.util.Map;

/**
 * 认证控制器
 *
 * @author enterprise
 * @since 2023-12-01
 */
@Slf4j
@RestController
@RequestMapping("/auth")
@RequiredArgsConstructor
@Tag(name = "认证管理", description = "用户认证相关接口")
public class AuthController {

    private final AuthenticationManager authenticationManager;
    private final JwtUtils jwtUtils;
    private final UserService userService;
    private final PasswordEncoder passwordEncoder;

    /**
     * 用户登录
     */
    @PostMapping("/login")
    @Operation(summary = "用户登录", description = "用户登录并返回JWT令牌")
    public Result<Map<String, Object>> login(@Valid @RequestBody LoginRequest loginRequest) {
        AssertUtils.notEmpty(loginRequest.getUsername(), "用户名不能为空");
        AssertUtils.notEmpty(loginRequest.getPassword(), "密码不能为空");

        try {
            // 认证用户
            Authentication authentication = authenticationManager.authenticate(
                    new UsernamePasswordAuthenticationToken(loginRequest.getUsername(), loginRequest.getPassword())
            );

            SecurityContextHolder.getContext().setAuthentication(authentication);

            // 获取用户信息
            User user = userService.getByUsername(loginRequest.getUsername());
            if (user == null) {
                return Result.error("用户不存在");
            }

            if (user.getStatus() == 0) {
                return Result.error("用户已被禁用");
            }

            // 生成JWT令牌
            String token = jwtUtils.generateToken(user.getUsername());

            // 更新最后登录信息
            userService.updateLastLoginInfo(user.getId(), getClientIp());

            // 返回结果
            Map<String, Object> result = new HashMap<>();
            result.put("token", token);
            result.put("user", user);

            return Result.success("登录成功", result);
        } catch (Exception e) {
            log.error("用户登录失败: {}", e.getMessage());
            return Result.error("用户名或密码错误");
        }
    }

    /**
     * 用户注册
     */
    @PostMapping("/register")
    @Operation(summary = "用户注册", description = "新用户注册")
    public Result<Void> register(@Valid @RequestBody RegisterRequest registerRequest) {
        AssertUtils.notEmpty(registerRequest.getUsername(), "用户名不能为空");
        AssertUtils.notEmpty(registerRequest.getPassword(), "密码不能为空");
        AssertUtils.notEmpty(registerRequest.getConfirmPassword(), "确认密码不能为空");

        if (!registerRequest.getPassword().equals(registerRequest.getConfirmPassword())) {
            return Result.error("两次输入的密码不一致");
        }

        // 检查用户名是否已存在
        if (userService.getByUsername(registerRequest.getUsername()) != null) {
            return Result.error("用户名已存在");
        }

        // 检查邮箱是否已存在
        if (registerRequest.getEmail() != null && userService.getByEmail(registerRequest.getEmail()) != null) {
            return Result.error("邮箱已存在");
        }

        // 创建用户
        User user = new User();
        user.setUsername(registerRequest.getUsername());
        user.setPassword(registerRequest.getPassword());
        user.setNickname(registerRequest.getNickname());
        user.setEmail(registerRequest.getEmail());
        user.setPhone(registerRequest.getPhone());

        if (userService.createUser(user)) {
            return Result.success("注册成功");
        } else {
            return Result.error("注册失败");
        }
    }

    /**
     * 刷新令牌
     */
    @PostMapping("/refresh")
    @Operation(summary = "刷新令牌", description = "刷新JWT令牌")
    public Result<Map<String, String>> refreshToken(@RequestHeader("Authorization") String authorization) {
        String token = jwtUtils.extractTokenFromHeader(authorization);
        if (token == null) {
            return Result.error("无效的令牌");
        }

        String refreshedToken = jwtUtils.refreshToken(token);
        if (refreshedToken == null) {
            return Result.error("令牌刷新失败");
        }

        Map<String, String> result = new HashMap<>();
        result.put("token", refreshedToken);

        return Result.success("令牌刷新成功", result);
    }

    /**
     * 获取当前用户信息
     */
    @GetMapping("/profile")
    @Operation(summary = "获取用户信息", description = "获取当前登录用户信息")
    public Result<User> getProfile() {
        Authentication authentication = SecurityContextHolder.getContext().getAuthentication();
        if (authentication == null || !authentication.isAuthenticated()) {
            return Result.error("用户未登录");
        }

        String username = authentication.getName();
        User user = userService.getByUsername(username);
        if (user == null) {
            return Result.error("用户不存在");
        }

        // 清除敏感信息
        user.setPassword(null);

        return Result.success(user);
    }

    /**
     * 修改密码
     */
    @PostMapping("/change-password")
    @Operation(summary = "修改密码", description = "修改当前用户密码")
    public Result<Void> changePassword(@Valid @RequestBody ChangePasswordRequest request) {
        AssertUtils.notEmpty(request.getOldPassword(), "旧密码不能为空");
        AssertUtils.notEmpty(request.getNewPassword(), "新密码不能为空");
        AssertUtils.notEmpty(request.getConfirmPassword(), "确认密码不能为空");

        if (!request.getNewPassword().equals(request.getConfirmPassword())) {
            return Result.error("两次输入的新密码不一致");
        }

        Authentication authentication = SecurityContextHolder.getContext().getAuthentication();
        if (authentication == null || !authentication.isAuthenticated()) {
            return Result.error("用户未登录");
        }

        String username = authentication.getName();
        User user = userService.getByUsername(username);
        if (user == null) {
            return Result.error("用户不存在");
        }

        if (userService.changePassword(user.getId(), request.getOldPassword(), request.getNewPassword())) {
            return Result.success("密码修改成功");
        } else {
            return Result.error("密码修改失败");
        }
    }

    /**
     * 获取客户端IP地址
     */
    private String getClientIp() {
        // 这里可以根据实际情况获取客户端IP
        return "127.0.0.1";
    }

    /**
     * 登录请求
     */
    public static class LoginRequest {
        private String username;
        private String password;

        public String getUsername() {
            return username;
        }

        public void setUsername(String username) {
            this.username = username;
        }

        public String getPassword() {
            return password;
        }

        public void setPassword(String password) {
            this.password = password;
        }
    }

    /**
     * 注册请求
     */
    public static class RegisterRequest {
        private String username;
        private String password;
        private String confirmPassword;
        private String nickname;
        private String email;
        private String phone;

        public String getUsername() {
            return username;
        }

        public void setUsername(String username) {
            this.username = username;
        }

        public String getPassword() {
            return password;
        }

        public void setPassword(String password) {
            this.password = password;
        }

        public String getConfirmPassword() {
            return confirmPassword;
        }

        public void setConfirmPassword(String confirmPassword) {
            this.confirmPassword = confirmPassword;
        }

        public String getNickname() {
            return nickname;
        }

        public void setNickname(String nickname) {
            this.nickname = nickname;
        }

        public String getEmail() {
            return email;
        }

        public void setEmail(String email) {
            this.email = email;
        }

        public String getPhone() {
            return phone;
        }

        public void setPhone(String phone) {
            this.phone = phone;
        }
    }

    /**
     * 修改密码请求
     */
    public static class ChangePasswordRequest {
        private String oldPassword;
        private String newPassword;
        private String confirmPassword;

        public String getOldPassword() {
            return oldPassword;
        }

        public void setOldPassword(String oldPassword) {
            this.oldPassword = oldPassword;
        }

        public String getNewPassword() {
            return newPassword;
        }

        public void setNewPassword(String newPassword) {
            this.newPassword = newPassword;
        }

        public String getConfirmPassword() {
            return confirmPassword;
        }

        public void setConfirmPassword(String confirmPassword) {
            this.confirmPassword = confirmPassword;
        }
    }
}