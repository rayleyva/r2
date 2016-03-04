# p28: a DSL(Domain-Specific Language) for HTTP Auto Test

### Synposis

p28 means Perfect Numer 28. As known to all, 28 is the first perfect number.

### Installation

    go install github.com/brg-liuwei/p28

### Usage

* Write the content as follows into a text file:

        #!/usr/bin/env p28

        # this is comment
        echo send get request to github.com, check return code
        req get https://github.com
        ret 200

* Assume your file name is `001-ret.l` (see example/001-ret.l):

        chmod +x 001-ret.l
        ./001-ret.l

        // Output:
        send get request to github.com, check return code
        [PASS] <./t1.l:6> ret 200

