# p28: a DSL for HTTP Auto Test

### Synposis

p28 means Perfect Numer 28. As known to all, 28 is the first perfect number.

### Installation

    go install github.com/brg-liuwei/p28

### Usage

* Write the content as follows into a text file:

    #!/usr/bin/env p28

    # this is comment
    echo hello world

* Assume your file name is `hello.l`:

    chmod +x hello.l
    ./hello.l

    // Output:
    hello world
