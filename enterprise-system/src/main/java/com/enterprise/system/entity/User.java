package com.enterprise.system.entity;

import com.baomidou.mybatisplus.annotation.*;
import lombok.Data;
import lombok.EqualsAndHashCode;

import jakarta.persistence.*;
import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.Size;
import java.time.LocalDateTime;

/**
 * 用户实体
 *
 * @author enterprise
 * @since 2023-12-01
 */
@Data
@EqualsAndHashCode(callSuper = false)
@Entity
@Table(name = "sys_user")
@TableName("sys_user")
public class User {

    /**
     * 用户ID
     */
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @TableId(type = IdType.AUTO)
    private Long id;

    /**
     * 用户名
     */
    @NotBlank(message = "用户名不能为空")
    @Size(min = 3, max = 50, message = "用户名长度必须在3-50个字符之间")
    @Column(unique = true, nullable = false, length = 50)
    private String username;

    /**
     * 密码
     */
    @NotBlank(message = "密码不能为空")
    @Size(min = 6, max = 100, message = "密码长度必须在6-100个字符之间")
    @Column(nullable = false, length = 100)
    private String password;

    /**
     * 昵称
     */
    @Size(max = 50, message = "昵称长度不能超过50个字符")
    @Column(length = 50)
    private String nickname;

    /**
     * 邮箱
     */
    @Email(message = "邮箱格式不正确")
    @Size(max = 100, message = "邮箱长度不能超过100个字符")
    @Column(length = 100)
    private String email;

    /**
     * 手机号
     */
    @Size(max = 20, message = "手机号长度不能超过20个字符")
    @Column(length = 20)
    private String phone;

    /**
     * 头像
     */
    @Size(max = 255, message = "头像URL长度不能超过255个字符")
    @Column(length = 255)
    private String avatar;

    /**
     * 状态（0：禁用，1：启用）
     */
    @Column(nullable = false)
    private Integer status = 1;

    /**
     * 性别（0：未知，1：男，2：女）
     */
    @Column
    private Integer gender = 0;

    /**
     * 部门ID
     */
    @Column
    private Long deptId;

    /**
     * 最后登录时间
     */
    @Column
    private LocalDateTime lastLoginTime;

    /**
     * 最后登录IP
     */
    @Size(max = 50, message = "IP地址长度不能超过50个字符")
    @Column(length = 50)
    private String lastLoginIp;

    /**
     * 备注
     */
    @Size(max = 500, message = "备注长度不能超过500个字符")
    @Column(length = 500)
    private String remark;

    /**
     * 创建时间
     */
    @Column(nullable = false, updatable = false)
    @TableField(fill = FieldFill.INSERT)
    private LocalDateTime createTime;

    /**
     * 更新时间
     */
    @Column(nullable = false)
    @TableField(fill = FieldFill.INSERT_UPDATE)
    private LocalDateTime updateTime;

    /**
     * 创建人
     */
    @Column(updatable = false)
    @TableField(fill = FieldFill.INSERT)
    private String createBy;

    /**
     * 更新人
     */
    @Column
    @TableField(fill = FieldFill.INSERT_UPDATE)
    private String updateBy;

    /**
     * 逻辑删除标识（0：未删除，1：已删除）
     */
    @Column(nullable = false)
    @TableLogic
    private Integer deleted = 0;
}