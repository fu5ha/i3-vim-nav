#!/bin/bash

if [ ! -d ~/bin ]; then
    mkdir ~/bin
fi

rm -rf ~/bin/i3-vim-nav
ln -s ./i3-vim-nav ~/bin
