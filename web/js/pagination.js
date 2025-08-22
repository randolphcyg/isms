/**
 * 通用分页组件
 */
class Pagination {
    /**
     * 构造函数
     * @param {Object} options - 配置选项
     * @param {string} options.containerId - 分页控件容器ID
     * @param {Function} options.onPageChange - 页面改变回调函数
     * @param {number} options.currentPage - 当前页码
     * @param {number} options.pageSize - 每页显示条数
     * @param {number} options.total - 总记录数
     */
    constructor(options) {
        this.containerId = options.containerId;
        this.onPageChange = options.onPageChange;
        this.currentPage = options.currentPage || 1;
        this.pageSize = options.pageSize || 10;
        this.total = options.total || 0;
        
        this.init();
    }
    
    /**
     * 初始化分页组件
     */
    init() {
        this.render();
        this.bindEvents();
    }
    
    /**
     * 渲染分页控件
     */
    render() {
        const container = document.getElementById(this.containerId);
        if (!container) return;
        
        // 清空容器内容，这会自动移除所有事件监听器
        container.innerHTML = '';
        
        const totalPages = Math.ceil(this.total / this.pageSize);
        
        container.innerHTML = `
            <div class="pagination">
                <button id="firstPage" class="pagination-btn">首页</button>
                <button id="prevPage" class="pagination-btn">上一页</button>
                <span>第 </span>
                <input type="number" id="pageInput" class="page-input" value="${this.currentPage}" min="1" max="${totalPages}">
                <span> 页 / 共 <span id="totalPages">${totalPages || 1}</span> 页</span>
                <button id="nextPage" class="pagination-btn">下一页</button>
                <button id="lastPage" class="pagination-btn">末页</button>
                <span>每页显示 </span>
                <input type="number" id="pageSizeInput" class="page-size-input" value="${this.pageSize}" min="1" max="100">
                <span> 条</span>
                <button id="goPage" class="pagination-btn">跳转</button>
                <span id="totalInfo" class="total-info">共 ${this.total} 条记录</span>
            </div>
        `;
        
        this.updateControls();
    }
    
    /**
     * 绑定事件
     */
    bindEvents() {
        // 先移除已存在的事件监听器
        this.removeEvents();
        
        // 保存事件处理函数的引用，以便后续可以移除
        this._firstPageHandler = () => this.goToPage(1);
        this._prevPageHandler = () => this.goToPage(this.currentPage - 1);
        this._nextPageHandler = () => this.goToPage(this.currentPage + 1);
        this._lastPageHandler = () => {
            const totalPages = Math.ceil((this.total || 0) / (this.pageSize || 1));
            this.goToPage(totalPages || 1);
        };
        this._pageInputHandler = (e) => {
            const page = parseInt(e.target.value);
            if (!isNaN(page)) {
                this.goToPage(page);
            }
        };
        this._pageSizeInputHandler = (e) => {
            const pageSize = parseInt(e.target.value);
            if (!isNaN(pageSize) && pageSize > 0) {
                this.pageSize = pageSize;
                this.goToPage(1); // 重置到第一页
            }
        };
        this._goPageHandler = () => {
            const pageInput = document.getElementById('pageInput');
            const page = parseInt(pageInput.value);
            if (!isNaN(page)) {
                this.goToPage(page);
            }
        };
        
        // 绑定事件监听器
        document.getElementById('firstPage')?.addEventListener('click', this._firstPageHandler);
        document.getElementById('prevPage')?.addEventListener('click', this._prevPageHandler);
        document.getElementById('nextPage')?.addEventListener('click', this._nextPageHandler);
        document.getElementById('lastPage')?.addEventListener('click', this._lastPageHandler);
        document.getElementById('pageInput')?.addEventListener('change', this._pageInputHandler);
        document.getElementById('pageSizeInput')?.addEventListener('change', this._pageSizeInputHandler);
        document.getElementById('goPage')?.addEventListener('click', this._goPageHandler);
    }
    
    /**
     * 移除事件监听器
     */
    removeEvents() {
        // 移除所有可能的事件监听器
        const firstPageBtn = document.getElementById('firstPage');
        const prevPageBtn = document.getElementById('prevPage');
        const nextPageBtn = document.getElementById('nextPage');
        const lastPageBtn = document.getElementById('lastPage');
        const pageInput = document.getElementById('pageInput');
        const pageSizeInput = document.getElementById('pageSizeInput');
        const goPageBtn = document.getElementById('goPage');
        
        if (firstPageBtn) firstPageBtn.removeEventListener('click', this._firstPageHandler);
        if (prevPageBtn) prevPageBtn.removeEventListener('click', this._prevPageHandler);
        if (nextPageBtn) nextPageBtn.removeEventListener('click', this._nextPageHandler);
        if (lastPageBtn) lastPageBtn.removeEventListener('click', this._lastPageHandler);
        if (pageInput) pageInput.removeEventListener('change', this._pageInputHandler);
        if (pageSizeInput) pageSizeInput.removeEventListener('change', this._pageSizeInputHandler);
        if (goPageBtn) goPageBtn.removeEventListener('click', this._goPageHandler);
    }
    
    /**
     * 销毁分页组件，清理事件监听器
     */
    destroy() {
        // 调用移除事件监听器的方法
        this.removeEvents();
    }
    
    /**
     * 跳转到指定页面
     * @param {number} page - 页码
     */
    goToPage(page) {
        const totalPages = Math.ceil((this.total || 0) / (this.pageSize || 1));
        
        // 确保页码在有效范围内
        page = Math.max(1, Math.min(page || 1, totalPages || 1));
        
        if (page !== this.currentPage) {
            this.currentPage = page;
            this.updateControls();
            
            // 调用回调函数
            if (this.onPageChange) {
                this.onPageChange(this.currentPage, this.pageSize);
            }
        }
    }
    
    /**
     * 更新分页控件状态
     */
    updateControls() {
        const totalPages = Math.ceil(this.total / this.pageSize);
        
        // 更新页码输入框
        const pageInput = document.getElementById('pageInput');
        if (pageInput) {
            pageInput.value = this.currentPage;
            pageInput.min = 1;
            pageInput.max = totalPages;
        }
        
        // 更新总页数显示
        const totalPagesEl = document.getElementById('totalPages');
        if (totalPagesEl) {
            totalPagesEl.textContent = totalPages || 1;
        }
        
        // 更新按钮状态
        const firstPageBtn = document.getElementById('firstPage');
        const prevPageBtn = document.getElementById('prevPage');
        const nextPageBtn = document.getElementById('nextPage');
        const lastPageBtn = document.getElementById('lastPage');
        
        if (firstPageBtn) firstPageBtn.disabled = this.currentPage <= 1;
        if (prevPageBtn) prevPageBtn.disabled = this.currentPage <= 1;
        if (nextPageBtn) nextPageBtn.disabled = this.currentPage >= totalPages;
        if (lastPageBtn) lastPageBtn.disabled = this.currentPage >= totalPages;
        
        // 更新每页条数输入框
        const pageSizeInput = document.getElementById('pageSizeInput');
        if (pageSizeInput) {
            // 添加错误检查
            if (this.pageSize && !isNaN(this.pageSize)) {
                pageSizeInput.value = this.pageSize;
            }
        }
        
        // 更新总记录数显示
        const totalInfo = document.getElementById('totalInfo');
        if (totalInfo) {
            totalInfo.textContent = `共 ${this.total} 条记录`;
        }
    }
    
    /**
     * 更新总记录数
     * @param {number} total - 总记录数
     */
    updateTotal(total) {
        this.total = total;
        this.totalPages = Math.ceil(this.total / this.pageSize);
        this.updateControls();
    }
    
    /**
     * 更新每页显示条数
     * @param {number} pageSize - 每页显示条数
     */
    updatePageSize(pageSize) {
        this.pageSize = pageSize;
        // 不再自动调用 goToPage(1)，由调用者决定是否需要跳转
        this.updateControls();
    }
}

// 导出Pagination类
window.Pagination = Pagination;