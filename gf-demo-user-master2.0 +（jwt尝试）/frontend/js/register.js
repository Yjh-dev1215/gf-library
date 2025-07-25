// 全局变量用于存储倒计时
let countdown = 0;
let countdownInterval;

// 发送验证码按钮的倒计时处理
function updateCountdown() {
    const sendCodeBtn = document.getElementById('sendCodeBtn');
    if (countdown > 0) {
        sendCodeBtn.disabled = true;
        sendCodeBtn.textContent = `${countdown}秒后重试`;
        countdown--;
    } else {
        sendCodeBtn.disabled = false;
        sendCodeBtn.textContent = '发送验证码';
        clearInterval(countdownInterval);
    }
}

// 处理发送验证码
async function handleSendCode(event) {
    event.preventDefault();
    
    const emailInput = document.getElementById('email');
    const emailError = document.getElementById('emailError');
    const email = emailInput.value.trim();
    
    // 验证邮箱格式
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(email)) {
        emailError.textContent = '请输入有效的邮箱地址';
        return;
    }

    try {
        const response = await api.sendEmailCode(email);
        if (response.code === 0) {
            // 发送成功，开始倒计时
            countdown = 60;
            countdownInterval = setInterval(updateCountdown, 1000);
            emailError.textContent = '';
            // 显示成功消息
            document.getElementById('successMessage').textContent = '验证码已发送到您的邮箱';
            setTimeout(() => {
                document.getElementById('successMessage').textContent = '';
            }, 3000);
        } else {
            emailError.textContent = response.message || '发送验证码失败';
        }
    } catch (error) {
        emailError.textContent = '发送验证码失败，请稍后重试';
        console.error('Send code error:', error);
    }
}

// 处理注册表单提交
async function handleRegister(event) {
    event.preventDefault();
    
    // 清除所有错误消息
    const errorElements = document.querySelectorAll('.error-message');
    errorElements.forEach(element => element.textContent = '');
    
    // 获取表单数据
    const passport = document.getElementById('passport').value.trim();
    const nickname = document.getElementById('nickname').value.trim();
    const email = document.getElementById('email').value.trim();
    const code = document.getElementById('code').value.trim();
    const password = document.getElementById('password').value;
    const confirmPassword = document.getElementById('confirmPassword').value;
    
    // 基本验证
    if (!passport || !nickname || !email || !code || !password || !confirmPassword) {
        document.getElementById('errorMessage').textContent = '请填写所有必填字段';
        return;
    }
    
    // 验证密码匹配
    if (password !== confirmPassword) {
        document.getElementById('passwordError').textContent = '两次输入的密码不一致';
        return;
    }
    
    // 验证码格式验证
    if (!/^\d{4}$/.test(code)) {
        document.getElementById('codeError').textContent = '验证码必须是4位数字';
        return;
    }

    try {
        const response = await api.register(passport, password, confirmPassword, nickname, email, code);
        if (response.code === 0) {
            // 注册成功
            document.getElementById('successMessage').textContent = '注册成功！即将跳转到登录页面...';
            setTimeout(() => {
                window.location.href = 'index.html';
            }, 2000);
        } else {
            // 显示错误消息
            document.getElementById('errorMessage').textContent = response.message || '注册失败';
        }
    } catch (error) {
        document.getElementById('errorMessage').textContent = '注册失败，请稍后重试';
        console.error('Register error:', error);
    }
}

// 实时验证用户名是否可用
async function checkPassportAvailability() {
    const passport = document.getElementById('passport').value.trim();
    const passportError = document.getElementById('passportError');
    
    if (passport.length < 1) {
        passportError.textContent = '';
        return;
    }

    try {
        const response = await api.checkPassport(passport);
        if (response.code !== 0) {
            passportError.textContent = response.message || '用户名已被使用';
        } else {
            passportError.textContent = '';
        }
    } catch (error) {
        console.error('Check passport error:', error);
    }
}

// 实时验证昵称是否可用
async function checkNicknameAvailability() {
    const nickname = document.getElementById('nickname').value.trim();
    const nicknameError = document.getElementById('nicknameError');
    
    if (nickname.length < 1) {
        nicknameError.textContent = '';
        return;
    }

    try {
        const response = await api.checkNickname(nickname);
        if (response.code !== 0) {
            nicknameError.textContent = response.message || '昵称已被使用';
        } else {
            nicknameError.textContent = '';
        }
    } catch (error) {
        console.error('Check nickname error:', error);
    }
}

// 添加事件监听器
document.addEventListener('DOMContentLoaded', () => {
    // 注册表单提交
    document.getElementById('registerForm').addEventListener('submit', handleRegister);
    
    // 发送验证码按钮点击
    document.getElementById('sendCodeBtn').addEventListener('click', handleSendCode);
    
    // 用户名输入框失去焦点时验证
    document.getElementById('passport').addEventListener('blur', checkPassportAvailability);
    
    // 昵称输入框失去焦点时验证
    document.getElementById('nickname').addEventListener('blur', checkNicknameAvailability);
});