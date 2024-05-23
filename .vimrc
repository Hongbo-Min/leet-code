set number "设置行号"

set nocompatible "设置不与Vi兼容"

syntax on

set showmode "底部显示当前模式"

set showcmd "显示命令"

set mouse=a "使用鼠标"

set encoding=utf-8 "编码方式"

set t_Co=256 "使用256色"

filetype indent on

set autoindent "自动缩进"

set tabstop=2 "tab时显示的空格数"

set expandtab "自动将Tab转为空格"

set softtabstop=2 "Tab转换为多少个空格"

set cursorline "高亮当前行"

set textwidth=80 "设置行宽"

set wrap "自动拆行，太长的行自动拆分"

set nowrap "关闭自动折行"

set laststatus=2 "是否显示底部0不显示1只有在多窗口时显示2表示显示"

set ruler "状态栏显示当前行列"

set showmatch "自动匹配括号"

set hlsearch "高亮匹配搜索"

set incsearch "每输入一个字符都自动进行匹配"

set spell spelllang=en_us "打开英语单词拼写检查"

set nobackup "不创建备份文件"

set noswapfile "不创建交换文件"

set autochdir "自动切换工作目录"

set noerrorbells "出错时不响"

set visualbell "出错时视觉提醒"

set autoread "文件监控"

set listchars=tab:»■,trail:■
set list "末尾有多余的空格时替换成小方块"

set wildmenu
set wildmode=longest:list,full "命令模式下，Tab自动补全命令"
