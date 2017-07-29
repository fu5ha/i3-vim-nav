# i3-vim-nav
Seamlessly navigate between i3 windows and vim splits. 

## Installation

Using vim-plug (recommended), in vimrc/init.vim

```vim
Plug 'termhn/i3-vim-nav', { 'do': 'mkdir ~/bin && ln -s i3-vim-nav ~/bin' }


" i3 integration
nnoremap <c-l> :call Focus('right', 'l')<CR>
nnoremap <c-h> :call Focus('left', 'h')<CR>
nnoremap <c-k> :call Focus('up', 'k')<CR>
nnoremap <c-j> :call Focus('down', 'j')<CR>
```

In i3 config:

```
bindsym --release $mod+h exec --no-startup-id "~/bin/i3-vim-nav h"
bindsym --release $mod+j exec --no-startup-id "~/bin/i3-vim-nav j"
bindsym --release $mod+k exec --no-startup-id "~/bin/i3-vim-nav k"
bindsym --release $mod+l exec --no-startup-id "~/bin/i3-vim-nav l"
```

