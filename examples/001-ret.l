#!/usr/bin/env p28

# file: 001-ret.l
echo send get request to github.com, check return code
req get https://github.com
ret 200