#!/usr/bin/env r2

# nginx config:
# 
#     location = /test_r2/005 {
#         content_by_lua_block {
#             -- test case 1: header-equal
#             ngx.header["MyAwesomeHeader"] = "TestForR2D2"
# 
#             local cjson = require "cjson"
#
#             -- test case 2: body-equal
#             local expect_result = ngx.var.arg_result
#             local rc = {
#                 errno = 0,
#                 errmsg = "ok",
#                 result = expect_result and expect_result or "empty",
#             }
#
#             ngx.say(cjson.encode(rc))
#         }
#     }

req get http://127.0.0.1:9999/test_r2/005?result=hello_r2
header-match Content-Type application
header-match MyAwesomeHeader R2D2
body-match {"errno":(?P<error_number>\\d+).+
var-equal error_number 0

req get http://127.0.0.1:9999/test_r2/005
body-match "\"result\":\"(?P<res>\\w+)\""
var-equal res empty

