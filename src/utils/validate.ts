// 表单验证工具函数

// 验证邮箱
export const validateEmail = (email: string): boolean => {
  const reg = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/
  return reg.test(email)
}

// 验证手机号
export const validatePhone = (phone: string): boolean => {
  const reg = /^1[3-9]\d{9}$/
  return reg.test(phone)
}

// 验证身份证号
export const validateIdCard = (idCard: string): boolean => {
  const reg = /(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)/
  return reg.test(idCard)
}

// 验证密码强度
export const validatePassword = (password: string): {
  isValid: boolean
  strength: 'weak' | 'medium' | 'strong'
  message: string
} => {
  if (password.length < 6) {
    return {
      isValid: false,
      strength: 'weak',
      message: '密码长度至少6位'
    }
  }

  const hasLetter = /[a-zA-Z]/.test(password)
  const hasNumber = /\d/.test(password)
  const hasSpecial = /[!@#$%^&*(),.?":{}|<>]/.test(password)

  if (hasLetter && hasNumber && hasSpecial) {
    return {
      isValid: true,
      strength: 'strong',
      message: '密码强度：强'
    }
  } else if ((hasLetter && hasNumber) || (hasLetter && hasSpecial) || (hasNumber && hasSpecial)) {
    return {
      isValid: true,
      strength: 'medium',
      message: '密码强度：中'
    }
  } else {
    return {
      isValid: true,
      strength: 'weak',
      message: '密码强度：弱'
    }
  }
}

// 验证URL
export const validateUrl = (url: string): boolean => {
  try {
    new URL(url)
    return true
  } catch {
    return false
  }
}

// 验证IP地址
export const validateIp = (ip: string): boolean => {
  const reg = /^(\d{1,3}\.){3}\d{1,3}$/
  if (!reg.test(ip)) return false
  
  const parts = ip.split('.')
  return parts.every(part => {
    const num = parseInt(part)
    return num >= 0 && num <= 255
  })
}

// 验证中文姓名
export const validateChineseName = (name: string): boolean => {
  const reg = /^[\u4e00-\u9fa5]{2,4}$/
  return reg.test(name)
}

// 验证金额
export const validateAmount = (amount: string): boolean => {
  const reg = /^[1-9]\d*(\.\d{1,2})?$|^0\.\d{1,2}$/
  return reg.test(amount)
}

// 验证正整数
export const validatePositiveInteger = (num: string): boolean => {
  const reg = /^[1-9]\d*$/
  return reg.test(num)
}

// 验证文件大小
export const validateFileSize = (file: File, maxSize: number): boolean => {
  return file.size <= maxSize * 1024 * 1024 // maxSize in MB
}

// 验证文件类型
export const validateFileType = (file: File, allowedTypes: string[]): boolean => {
  return allowedTypes.includes(file.type)
}

// 验证图片尺寸
export const validateImageSize = (
  file: File,
  maxWidth: number,
  maxHeight: number
): Promise<boolean> => {
  return new Promise((resolve) => {
    const img = new Image()
    img.onload = () => {
      resolve(img.width <= maxWidth && img.height <= maxHeight)
    }
    img.onerror = () => {
      resolve(false)
    }
    img.src = URL.createObjectURL(file)
  })
}

// 生成验证规则
export const generateRules = {
  // 必填
  required: (message = '此项为必填项') => ({
    required: true,
    message,
    trigger: 'blur'
  }),

  // 邮箱
  email: (message = '请输入正确的邮箱格式') => ({
    type: 'email',
    message,
    trigger: 'blur'
  }),

  // 手机号
  phone: (message = '请输入正确的手机号') => ({
    pattern: /^1[3-9]\d{9}$/,
    message,
    trigger: 'blur'
  }),

  // 长度限制
  length: (min: number, max: number, message?: string) => ({
    min,
    max,
    message: message || `长度在 ${min} 到 ${max} 个字符`,
    trigger: 'blur'
  }),

  // 数字范围
  range: (min: number, max: number, message?: string) => ({
    type: 'number',
    min,
    max,
    message: message || `数值在 ${min} 到 ${max} 之间`,
    trigger: 'blur'
  }),

  // 自定义正则
  pattern: (pattern: RegExp, message: string) => ({
    pattern,
    message,
    trigger: 'blur'
  })
}