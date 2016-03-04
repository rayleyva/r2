#!/usr/bin/env p28

# file: 001-ret.l

# test case 1
# cyan start to test github.com
magenta case 1: req get https://github.com
req get https://github.com
ret 200
echo
# magenta test github.com end

# test case 2
# cyan start test github.com again
magenta case 2: req get https://github.com
req get https://github.com
ret 400
echo
# cyan test github.com end

# test case 3
# echo test github.com once again and again
magenta case 3: req get https://github.com
req get https://github.com
ret 200
echo

magenta case 4: req post https://github.com
req post https://github.com
ret 404
echo

magenta case 5: req post https://www.sina.com
req post http://www.sina.com
ret 403
echo

# echo test github.com thirdly ok

# echo hello fish
# 
# red hello fish
# 
# green hello fish
# 
# blue hello blue fish
# 
# brown hello brown fish
# 
# magenta hello fat fish
# 
# cyan hello lovely fish
