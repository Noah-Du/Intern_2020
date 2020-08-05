### cmd结构体
  type Cmd struct {
    
    // 命令
    Path string
    
    // 命令参数
    Args []string
    
    // 进程执行环境
    Env []string
   
    // 命令执行目录, 默认当前目录
    Dir string
    
    
    // 命令输入
    Stdin io.Reader
    
    // 命令输出
    Stdout io.Writer
    
    // 命令报错
    Stderr io.Writer
   
   
    ExtraFiles []*os.File
    SysProcAttr *syscall.SysProcAttr
    Process *os.Process
    ProcessState *os.ProcessState
 
  }

