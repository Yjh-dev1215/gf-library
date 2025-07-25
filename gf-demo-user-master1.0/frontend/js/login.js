document.addEventListener('DOMContentLoaded', () => {
    // 检查用户是否已登录
    checkLoginStatus();
    
    // 登录表单提交事件
    const loginForm = document.getElementById('loginForm');
    loginForm.addEventListener('submit', handleLogin);
});

// 检查登录状态
async function checkLoginStatus() {
    try {
        const response = await api.isSignedIn();
        if (response.data && response.data.isSignedIn) {
            // 用户已登录，跳转到主页面
            window.location.href = 'dashboard.html';
        }
    } catch (error) {
        console.error('Error checking login status:', error);
    }
}

// 处理登录表单提交
async function handleLogin(event) {
    event.preventDefault();
    
    const identifier = document.getElementById('identifier').value;
    const password = document.getElementById('password').value;
    const errorMessage = document.getElementById('errorMessage');
    
    // 简单的表单验证
    if (!identifier || !password) {
        errorMessage.textContent = '请填写所有必填字段';
        return;
    }
    
    try {
        // 调用登录API
        const response = await api.login(identifier, password);
        
        if (response.code === 0) {
            // 登录成功，跳转到主页面
            window.location.href = 'dashboard.html';
        } else {
            // 显示错误信息
            errorMessage.textContent = response.message || '登录失败，请检查用户名和密码';
        }
    } catch (error) {
        console.error('Login error:', error);
        errorMessage.textContent = '登录过程中发生错误，请稍后再试';
    }
}