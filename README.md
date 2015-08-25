# c-go-zip
Go wrapper for zip package which you can link in C.

## Build Go wrapper
```sh
# go build will create zip.a and zip.h files.
$ go build -buildmode c-archive -o zip.a
```
## Link in C
```C
#include "zip.h"

// usage: ./a.out my.zip file1 file2 file3 ...
int main(int argc, char *argv[]) {
    GoSlice files = { &argv[2], argc - 2, argc };
    
    // ziping...
    Zip(argv[1], files);
    
    // unzipping...
    Unzip(argv[1]);
    return 0;
}
```

```sh
# Compile & Link with zip.a
$ gcc zip.c zip.a -Wl,-no_pie
# Usage
$ ./a.out my.zip file1 file2 file3
```

