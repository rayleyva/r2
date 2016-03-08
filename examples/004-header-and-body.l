#!/usr/bin/env r2

# nginx config:
# server {
#     listen 9999;
#     server_name 127.0.0.1;
# 
#     location /test_r2 {
#         content_by_lua_block {
#             -- test case 1: http method
#             if ngx.req.get_method() ~= "POST" then
#                 ngx.exit(ngx.HTTP_FORBIDDEN) -- ngx.HTTP_FORBIDDEN: 403
#             end
#
#             -- test case 2: header
#             local header = ngx.req.get_headers()
#             if header.content_type ~= "application/json" then
#                 ngx.exit(520)
#             end
# 
#             local cjson = require "cjson"
#             ngx.req.read_body()
#             local body = ngx.req.get_body_data()
#             local ok, data = pcall(cjson.decode, body)
#
#             -- test case 3: json decode
#             if not ok then
#                 ngx.exit(521)
#             end
# 
#             -- test case 4: json key
#             if not data.ret then
#                 ngx.exit(522)
#             end
# 
#             -- test case 5: json data type
#             if type(data.ret) ~= "number" then
#                 ngx.exit(523)
#             end
# 
#             -- test case 6: json data value
#             ngx.exit(data.ret)
#         }
#     }
# }

cyan case 1: HTTP GET
req get http://127.0.0.1:9999/test_r2
ret 403

cyan case 2: HTTP HEADER
req post http://127.0.0.1:9999/test_r2
header Content-Type application/text
ret 520

cyan case 3: json decode 
req post http://127.0.0.1:9999/test_r2
header Content-Type application/json
body {"ret":204B}
ret 521

cyan case 4: json key
req post http://127.0.0.1:9999/test_r2
header Content-Type application/json
body {"return":204}
ret 522

cyan case 5: json data type
req post http://127.0.0.1:9999/test_r2
header Content-Type application/json
body "{\"ret\":\"204\"}"
ret 523

cyan case 6: json data type
req post http://127.0.0.1:9999/test_r2
header Content-Type application/json
body {"ret":204}
ret 204

