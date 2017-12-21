# bb

bb is a command line tool to replace hex string with Go byte slice.

# Installation

go get

# Use case

* command line

Run bb command with an argument.

```
# bb 16bf0f1e88de
```

You will get `[]byte{0x16, 0xbf, 0x0f, 0x1e, 0x88, 0xde, }`

* vim visual mode

Select the line and run bb as a external command `!bb`

```
16bf0f1e88de

:'<,'>!bb
```

The line `16bf0f1e88de` will be replaced with `[]byte{0x16, 0xbf, 0x0f, 0x1e, 0x88, 0xde, }`

# License

MIT
