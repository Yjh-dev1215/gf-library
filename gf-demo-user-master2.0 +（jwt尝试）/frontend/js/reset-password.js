document.addEventListener('DOMContentLoaded', function() {
    const form = document.getElementById('resetPasswordForm');
    const sendCodeBtn = document.getElementById('sendCodeBtn');
    const emailInput = document.getElementById('email');
    const verificationCodeInput = document.getElementById('verificationCode');
    const newPasswordInput = document.getElementById('newPassword');
    const confirmPasswordInput = document.getElementById('confirmPassword');
    const errorMessage = document.getElementById('errorMessage');
    const successMessage = document.getElementById('successMessage');

    let countdown = 60;
    let timer = null;

    // 发送验证码
    sendCodeBtn.addEventListener('click', async function() {
        const email = emailInput.value.trim();
        if (!email) {
            showError('请输入邮箱地址');
            return;
        }

        if (!isValidEmail(email)) {
            showError('请输入有效的邮箱地址');
            return;
        }

        try {
            // 调用发送验证码API
            await api.sendEmailCode(email);
            
            // 开始倒计时
            startCountdown();
            showSuccess('验证码已发送，请查收邮箱');
        } catch (error) {
            showError(error.message || '发送验证码失败，请稍后重试');
        }
    });

    // 表单提交
    form.addEventListener('submit', async function(e) {
        e.preventDefault();
        clearMessages();

        const email = emailInput.value.trim();
        const verificationCode = verificationCodeInput.value.trim();
        const newPassword = newPasswordInput.value;
        const confirmPassword = confirmPasswordInput.value;

        // 表单验证
        if (!email || !verificationCode || !newPassword || !confirmPassword) {
            showError('请填写所有必填项');
            return;
        }

        if (!isValidEmail(email)) {
            showError('请输入有效的邮箱地址');
            return;
        }

        if (newPassword !== confirmPassword) {
            showError('两次输入的密码不一致');
            return;
        }

        if (newPassword.length < 6) {
            showError('密码长度不能少于6位');
            return;
        }

        try {
            // 调用重置密码API
            await api.resetPassword({
                email: email,
                verificationCode: verificationCode,
                newPassword: newPassword,
                confirmPassword: confirmPassword
            });

            showSuccess('密码重置成功！3秒后跳转到登录页面...');
            setTimeout(() => {
                window.location.href = 'index.html';
            }, 3000);
        } catch (error) {
            showError(error.message || '重置密码失败，请稍后重试');
        }
    });

    // 倒计时功能
    function startCountdown() {
        sendCodeBtn.disabled = true;
        countdown = 60;
        updateCountdownText();

        timer = setInterval(() => {
            countdown--;
            updateCountdownText();

            if (countdown <= 0) {
                clearInterval(timer);
                sendCodeBtn.disabled = false;
                sendCodeBtn.textContent = '发送验证码';
            }
        }, 1000);
    }

    function updateCountdownText() {
        sendCodeBtn.textContent = `重新发送(${countdown}s)`;
    }

    // 验证邮箱格式
    function isValidEmail(email) {
        return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
    }

    // 显示错误信息
    function showError(message) {
        errorMessage.textContent = message;
        errorMessage.style.display = 'block';
        successMessage.style.display = 'none';
    }

    // 显示成功信息
    function showSuccess(message) {
        successMessage.textContent = message;
        successMessage.style.display = 'block';
        errorMessage.style.display = 'none';
    }

    // 清除消息
    function clearMessages() {
        errorMessage.style.display = 'none';
        successMessage.style.display = 'none';
    }
});