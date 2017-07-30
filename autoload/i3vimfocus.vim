let s:plugin_path = escape(expand('<sfile>:p:h'), '\')

if has('python')
	exe 'pyfile ' . escape(s:plugin_path, ' ') . '/i3vimfocus.py'
elseif has('python3')
	exe 'py3file ' . escape(s:plugin_path, ' ') . '/i3vimfocus.py'
else
	echom "Your vim installation does not support +python or +python3."
endif

function! i3vimfocus#PythonExecProcess(name, args)
	if has('python')
		python PythonExecSubprocess()
	elseif has('python3')
		python3 PythonExecSubprocess()
	else
		echom "Your vim installation does not support +python or +python3."
	endif
endfunction
