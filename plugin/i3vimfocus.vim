func! Focus(comando,vim_comando)
  let oldw = winnr()
  silent exe 'wincmd ' . a:vim_comando
  let neww = winnr()
  if oldw == neww
      " Use python to invoke the i3-msg command so that vim doesn't need to be
      " redrawn.
      call i3vimfocus#PythonExecProcess("i3-msg", ["-q", "focus", a:comando])
  endif
endfunction
