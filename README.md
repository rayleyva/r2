# r2: a DSL for HTTP Auto Test

What is [DSL](https://en.wikipedia.org/wiki/Domain-specific_language) (Domain-Specific Language)?

### Synposis

[R2-D2](https://en.wikipedia.org/wiki/R2-D2) or Artoo-Detoo (called "R2" for short) is a fictional character in the Star Wars universe.

And I like this guy.

![R2-D2](R2-D2_Droid.png)

### Installation

    go get -u github.com/brg-liuwei/r2

### Usage

* Write the content as follows into a text file:

        #!/usr/bin/env r2

        # this is comment
        echo send get request to github.com, check return code
        req get https://github.com
        ret 200

* Assume your file name is `001-ret.l` (see [example/001-ret.l](https://github.com/brg-liuwei/r2/blob/master/examples/001-ret.l)):

        chmod +x 001-ret.l
        ./001-ret.l

        // Output:
        send get request to github.com, check return code
        [PASS] <./001-ret.l:6> ret 200

### Commands

* echo commands:

        echo <msg1> <msg2> ...
        red <msg1> <msg2> ...
        green <msg1> <msg2> ...
        brown <msg1> <msg2> ...
        blue <msg1> <msg2> ...
        magenta <msg1> <msg2> ...
        cyan <msg1> <msg2> ...

* generate a request:

        req <method> <url>
        header <key> <value>
        body <data>

* test response

        ret <status_code>

### Examples

for detail, see dir `example`, run example code:

    // run 001-ret.l
    r2 example/001-ret.l

    // run all test case in example
    r2 example/*.l

##### Examples of `header` and `body`:

* edit your nginx config file as follows:

        server {
            listen 9999;
            server_name 127.0.0.1;
        
            location /test_r2 {
                content_by_lua_block {
                    -- test case 1: http method
                    if ngx.req.get_method() ~= "POST" then
                        ngx.exit(ngx.HTTP_FORBIDDEN) -- ngx.HTTP_FORBIDDEN: 403
                    end
        
                    -- test case 2: header
                    local header = ngx.req.get_headers()
                    if header.content_type ~= "application/json" then
                        ngx.exit(520)
                    end
        
                    local cjson = require "cjson"
                    ngx.req.read_body()
                    local body = ngx.req.get_body_data()
                    local ok, data = pcall(cjson.decode, body)
        
                    -- test case 3: json decode
                    if not ok then
                        ngx.exit(521)
                    end
        
                    -- test case 4: json key
                    if not data.ret then
                        ngx.exit(522)
                    end
        
                    -- test case 5: json data type
                    if type(data.ret) ~= "number" then
                        ngx.exit(523)
                    end
        
                    -- test case 6: json data value
                    ngx.exit(data.ret)
                }
            }
        }

* write test case (see [004-header-body.l](example/004-header-body.l) for detail):

        #!/usr/bin/env r2

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

* run example code:

        cd /path/to/example/004-header-body.l
        ./004-header-body.l

        // Outputs:
        // case 1: HTTP GET
        // [PASS] <./004-header-and-body.l:49> ret 403
        // case 2: HTTP HEADER
        // [PASS] <./004-header-and-body.l:54> ret 520
        // case 3: json decode
        // [PASS] <./004-header-and-body.l:60> ret 521
        // case 4: json key
        // [PASS] <./004-header-and-body.l:66> ret 522
        // case 5: json data type
        // [PASS] <./004-header-and-body.l:72> ret 523
        // case 6: json data type
        // [PASS] <./004-header-and-body.l:78> ret 204

### TODO

* generate a request:

        auth <auth-name> <auth-interface>

* test response

        header-equal <key> <value>
        header-match <key> <regexp-value>

        body-equal <data>
        body-match <regexp-data>

        latency <micro-second>

