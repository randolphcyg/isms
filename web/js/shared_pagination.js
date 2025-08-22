/**
 * 通用分页初始化函数
 * @param {string} containerId - 分页控件容器ID
 * @param {number} currentPage - 当前页码
 * @param {number} pageSize - 每页显示条数
 * @param {number} total - 总记录数
 * @param {Function} fetchDataFunction - 获取数据的函数
 * @returns {Pagination} Pagination实例
 */
function initSharedPagination(containerId, currentPage, pageSize, total, fetchDataFunction) {
    return new Pagination({
        containerId: containerId,
        currentPage: currentPage,
        pageSize: pageSize,
        total: total,
        onPageChange: (page, size) => {
            fetchDataFunction(page, size);
        }
    });
}

// 导出函数以供其他模块使用
if (typeof module !== 'undefined' && module.exports) {
    module.exports = { initSharedPagination };
} else if (typeof window !== 'undefined') {
    window.initSharedPagination = initSharedPagination;
}