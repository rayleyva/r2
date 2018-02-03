# r2: a DSL for HTTP Auto Test

What is a [DSL](https://en.wikipedia.org/wiki/Domain-specific_language) (Domain-Specific Language)?

### Synopsis

Found this at: github.com/brg-liuwei/r2

Looked interesting, but no updates since 2016. Decided to pick it up for a quick weekend project.

Don't expect much, if any activity on this, just have it around to kick the tires.

### Installation

    go get -u github.com/rayleyva/r2

### Usage

* Write the content as follows into a text file:

        #!/usr/bin/env r2

        # this is comment
        echo send get request to github.com, check return code
        req get https://github.com
        ret 200

* Assume your file name is `001-ret.r2` (see [examples/001-ret.r2](examples/001-ret.r2)):

        chmod +x 001-ret.r2
        ./001-ret.r2

        // Output:
        // send get request to github.com, check return code
        // [PASS] <./001-ret.r2:6> ret 200

### Commands

* Echo commands:

        // see examples/003-echo.r2
        echo <msg1> <msg2> ...

        // see examples/002-color.r2
        red <msg1> <msg2> ...
        green <msg1> <msg2> ...
        brown <msg1> <msg2> ...
        blue <msg1> <msg2> ...
        magenta <msg1> <msg2> ...
        cyan <msg1> <msg2> ...

* Generate requests:

        // see examples/001-ret.r2
        req <method> <url>

        // see examples/004-header-and-body.r2
        header <key> <value>
        body <data>

* Check responses:

        // see examples/001-ret.r2
        ret <status_code>

        // see examples/005-equal.r2
        header-equal <key> <value>
        body-equal <data>

        // see examples/006-match.r2
        header-match <key> <regexp-value>
        body-match <regexp-data>
        var-equal <var-name> <var-value>

### Examples

For detail, see dir `example`, run example code:

    // run 001-ret.r2
    r2 example/001-ret.r2

    // run all test case in example
    r2 example/*.r2

##### Examples of `header` and `body`:

* Edit your nginx config file as follows:

        server {
            listen 9999;
            server_name 127.0.0.1;
        
            location /test_r2/004 {
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

* Write test case (see [examples/004-header-and-body.r2](examples/004-header-and-body.r2) for detail):

        #!/usr/bin/env r2

        cyan case 1: HTTP GET
        req get http://127.0.0.1:9999/test_r2/004
        ret 403
        
        cyan case 2: HTTP HEADER
        req post http://127.0.0.1:9999/test_r2/004
        header Content-Type application/text
        ret 520
        
        cyan case 3: json decode 
        req post http://127.0.0.1:9999/test_r2/004
        header Content-Type application/json
        body {"ret":204B}
        ret 521
        
        cyan case 4: json key
        req post http://127.0.0.1:9999/test_r2/004
        header Content-Type application/json
        body {"return":204}
        ret 522
        
        cyan case 5: json data type
        req post http://127.0.0.1:9999/test_r2/004
        header Content-Type application/json
        body "{\"ret\":\"204\"}"
        ret 523
        
        cyan case 6: json data type
        req post http://127.0.0.1:9999/test_r2/004
        header Content-Type application/json
        body {"ret":204}
        ret 204

* Run example code:

        cd /path/to/example/004-header-body.r2
        ./004-header-body.r2

        // Outputs:
        // case 1: HTTP GET
        // [PASS] <./004-header-and-body.r2:49> ret 403
        // case 2: HTTP HEADER
        // [PASS] <./004-header-and-body.r2:54> ret 520
        // case 3: json decode
        // [PASS] <./004-header-and-body.r2:60> ret 521
        // case 4: json key
        // [PASS] <./004-header-and-body.r2:66> ret 522
        // case 5: json data type
        // [PASS] <./004-header-and-body.r2:72> ret 523
        // case 6: json data type
        // [PASS] <./004-header-and-body.r2:78> ret 204

### TODO

* Generate requests:

        auth <auth-name> <auth-interface>

* Check responses:

        latency <micro-second>

