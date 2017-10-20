                   HOME.BIN
                 =============

SYNOPSIS
    Utilities for making computering more tolerable

INSTALL
    go install github.com/wnh/home.bin/...

TOOLS
     swd    Short working Directory
            Prints the current working directory with all but the
            last path element shortened to the first character.
            Ex: ~/src/github.com/wnh/home.bin  -> ~/s/g/w/home.bin

   gitst    Print git branch for the current directory for use in your path

      mk    Walk up the directory tree until a 'Makefile' is found and execute
            make in that directory using the given arguments.
