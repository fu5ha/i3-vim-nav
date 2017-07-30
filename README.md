# i3-vim-nav
Seamlessly navigate between i3 windows and vim splits.

# Installation

First, install the Vim plugin.

### Using vim-plug
In your .vimrc (vim) or .config/nvim/init.vim (neovim):

```vim
Plug 'termhn/i3-vim-nav'
	
" i3 integration
nnoremap <silent> <c-l> :call Focus('right', 'l')<CR>
nnoremap <silent> <c-h> :call Focus('left', 'h')<CR>
nnoremap <silent> <c-k> :call Focus('up', 'k')<CR>
nnoremap <silent> <c-j> :call Focus('down', 'j')<CR>
```
	
### Using Pathogen
1. cd ~/.vim/bundle
2. git clone https://github.com/termhn/i3-vim-nav
3. add the following to your .vimrc

```vim
" i3 integration
nnoremap <silent> <c-l> :call Focus('right', 'l')<CR>
nnoremap <silent> <c-h> :call Focus('left', 'h')<CR>
nnoremap <silent> <c-k> :call Focus('up', 'k')<CR>
nnoremap <silent> <c-j> :call Focus('down', 'j')<CR>
```
	
Next, install the binary on your PATH. If you have go installed, this can be done simply by

```
go get -u github.com/termhn/i3-vim-nav
```
If not, you can link the binary from its downloaded directory (this changes based on which plugin manager you used and if you're using vim or neovim into /usr/local/bin. For example, if you used Pathogen on default Vim, this would be:

```
ln -s ~/.vim/bundle/i3-vim-nav/i3-vim-nav /usr/local/bin
```

Then, in your i3 config (adjust the path to the executable as necessary if you installed it differently):

```
bindsym --release $mod+h exec --no-startup-id "i3-vim-nav h"
bindsym --release $mod+j exec --no-startup-id "i3-vim-nav j"
bindsym --release $mod+k exec --no-startup-id "i3-vim-nav k"
bindsym --release $mod+l exec --no-startup-id "i3-vim-nav l"
```
