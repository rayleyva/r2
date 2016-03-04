# p28: a DSL for HTTP Auto Test

What is [DSL](https://en.wikipedia.org/wiki/Domain-specific_language) (Domain-Specific Language)?

### Synposis

p28 means Perfect Numer 28. As known to all, 28 is the first perfect number.

### Installation

    go get -u github.com/brg-liuwei/p28

### Usage

* Write the content as follows into a text file:

        #!/usr/bin/env p28

        # this is comment
        echo send get request to github.com, check return code
        req get https://github.com
        ret 200

* Assume your file name is `001-ret.l` (see [example/001-ret.l](https://github.com/brg-liuwei/p28/blob/master/examples/001-ret.l)):

        chmod +x 001-ret.l
        ./001-ret.l

        // Output:
        send get request to github.com, check return code
        [PASS] <./t1.l:6> ret 200

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

* test response

        ret <status_code>

### Examples

for detail, see dir `example`, run example code:

    // run 001-ret.l
    p28 example/001-ret.l

    // run all test case in example
    p28 example/*.l

### TODO

* generate a request:

        header <key> <value1> <value2> ...
        auth <auth-name> <auth-interface>
        body <data>

* test response

        header-equal <key> <value1> <value2> ...
        header-match <key> <regexp-value>

        body-equal <data>
        body-match <regexp-data>

        latency <micro-second>
