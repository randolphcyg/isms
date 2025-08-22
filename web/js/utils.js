// 日期格式化函数
function formatDate(dateString) {
    if (!dateString) return '-';
    
    try {
        const date = new Date(dateString);
        if (isNaN(date.getTime())) return '-'; // 无效日期
        
        // 使用中文格式化日期（包含时间）
        return new Intl.DateTimeFormat('zh-CN', {
            year: 'numeric',
            month: '2-digit',
            day: '2-digit',
            hour: '2-digit',
            minute: '2-digit',
            second: '2-digit',
            hour12: false
        }).format(date);
    } catch (error) {
        console.error('日期格式化错误:', error);
        return '-'; // 格式化失败时返回默认值
    }
}

// 显示消息函数
function showMessage(text, type) {
    const messageEl = document.getElementById('message');
    if (!messageEl) {
        console.warn('未找到消息元素');
        return;
    }
    
    messageEl.textContent = text;
    messageEl.className = `message ${type}`;

    // 显示消息
    messageEl.style.display = 'block';

    // 3秒后隐藏消息
    setTimeout(() => {
        messageEl.style.display = 'none';
        messageEl.textContent = '';
        messageEl.className = 'message';
    }, 3000);
}

function formatSize(sizeBytes) {
    if (!sizeBytes || sizeBytes <= 0) return '0 B';

    const units = ['B', 'KB', 'MB', 'GB', 'TB'];
    let size = sizeBytes;
    let unitIndex = 0;

    while (size >= 1024 && unitIndex < units.length - 1) {
        size /= 1024;
        unitIndex++;
    }

    return size.toFixed(2) + ' ' + units[unitIndex];
}

// 导出函数以便在其他文件中使用
window.Utils = {
    formatDate,
    formatSize,
    showMessage
};