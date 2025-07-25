// API基础URL
const API_BASE_URL = 'http://localhost:8000';

// 封装API请求函数
const api = {
    // 登录请求
    login: async (identifier, password) => {
        try {
            const response = await fetch(`${API_BASE_URL}/user/sign-in`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    identifier,
                    password
                }),
                credentials: 'include' // 包含cookies
            });
            
            return await response.json();
        } catch (error) {
            console.error('Login error:', error);
            throw error;
        }
    },

    // 发送邮箱验证码
    sendEmailCode: async (email) => {
        try {
            const response = await fetch(`${API_BASE_URL}/user/email-code`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    email
                }),
                credentials: 'include'
            });
            
            return await response.json();
        } catch (error) {
            console.error('Send email code error:', error);
            throw error;
        }
    },

    // 注册请求
    register: async (passport, password, confirmPassword, nickname, email, code) => {
        try {
            const response = await fetch(`${API_BASE_URL}/user/sign-up`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    passport,
                    password,
                    password2: confirmPassword, // 后端API期望的字段名仍然是 password2
                    nickname,
                    email,
                    code
                }),
                credentials: 'include'
            });
            
            return await response.json();
        } catch (error) {
            console.error('Register error:', error);
            throw error;
        }
    },

    // 获取用户信息
    getUserProfile: async () => {
        try {
            const response = await fetch(`${API_BASE_URL}/user/profile`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                },
                credentials: 'include'
            });
            
            return await response.json();
        } catch (error) {
            console.error('Get profile error:', error);
            throw error;
        }
    },

    // 退出登录
    logout: async () => {
        try {
            const response = await fetch(`${API_BASE_URL}/user/sign-out`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                credentials: 'include'
            });
            
            return await response.json();
        } catch (error) {
            console.error('Logout error:', error);
            throw error;
        }
    },

    // 检查用户名是否可用
    checkPassport: async (passport) => {
        try {
            const response = await fetch(`${API_BASE_URL}/user/check-passport?passport=${passport}`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                },
                credentials: 'include'
            });
            
            return await response.json();
        } catch (error) {
            console.error('Check passport error:', error);
            throw error;
        }
    },

    // 检查昵称是否可用
    checkNickname: async (nickname) => {
        try {
            const response = await fetch(`${API_BASE_URL}/user/check-nick-name?nickname=${nickname}`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                },
                credentials: 'include'
            });
            
            return await response.json();
        } catch (error) {
            console.error('Check nickname error:', error);
            throw error;
        }
    },

    // 检查用户是否已登录
    isSignedIn: async () => {
        try {
            const response = await fetch(`${API_BASE_URL}/user/is-signed-in`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                },
                credentials: 'include'
            });
            
            return await response.json();
        } catch (error) {
            console.error('Check sign in status error:', error);
            throw error;
        }
    },

    // 重置密码
    resetPassword: async ({ email, verificationCode, newPassword, confirmPassword }) => {
        try {
            const response = await fetch(`${API_BASE_URL}/user/reset-password`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    Email: email,
                    Code: verificationCode,
                    Password: newPassword,
                    Password2: confirmPassword
                }),
                credentials: 'include'
            });
            
            return await response.json();
        } catch (error) {
            console.error('Reset password error:', error);
            throw error;
        }
    }
};