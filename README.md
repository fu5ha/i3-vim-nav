# i3-vim-nav
Seamlessly navigate between i3 windows and vim splits. 

## Installation

Using vim-plug (recommended), in vimrc/init.vim

```vim
Plug 'termhn/i3-vim-nav', { 'do': 'bash install.sh' }


" i3 integration
nnoremap <c-l> :call Focus('right', 'l')<CR>
nnoremap <c-h> :call Focus('left', 'h')<CR>
nnoremap <c-k> :call Focus('up', 'k')<CR>
nnoremap <c-j> :call Focus('down', 'j')<CR>
```

You can also install using pathogen/vundle and then run the install script manually, or simply put the binary somewhere on your path, or compile it yourself using 

```
go get -u github.com/termhn/i3-vim-nav
```
if you have go installed.

Then, in your i3 config:

```
bindsym --release $mod+h exec --no-startup-id "~/bin/i3-vim-nav h"
bindsym --release $mod+j exec --no-startup-id "~/bin/i3-vim-nav j"
bindsym --release $mod+k exec --no-startup-id "~/bin/i3-vim-nav k"
bindsym --release $mod+l exec --no-startup-id "~/bin/i3-vim-nav l"
```

