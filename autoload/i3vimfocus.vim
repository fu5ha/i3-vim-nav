let s:plugin_path = escape(expand('<sfile>:p:h'), '\')

exe 'pyfile ' . escape(s:plugin_path, ' ') . '/i3vimfocus.py'

function! i3vimfocus#PythonExecProcess(name, args)
    python PythonExecSubprocess()
endfunction
