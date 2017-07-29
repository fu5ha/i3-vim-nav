let s:plugin_path = escape(expand('<sfile>:p:h'), '\')

exe 'py3file ' . escape(s:plugin_path, ' ') . '/i3vimfocus.py'

function! i3vimfocus#PythonExecProcess(name, args)
    python3 PythonExecSubprocess()
endfunction
