# Flogo Web UI - Loader

Flogo Web UI - Loader is a utility to load Flogo apps into the Flogo Web UI using the commandline as opposed to using the import button. Especially when creating a new docker image, it helps to speed up the process of loading apps considerably.

You can build an executable out of this app using the command
```
go build
```

## Usage of the app:
```
$ flogowebloader [-dir | -filename] [-host]

  -dir
        Upload all JSON files in the current directory
  -filename string
        The name of the file you want to upload if you do not specify 'dir' (default "flogo.json")
  -host string
        The URL for Flogo Web UI (default "http://localhost:3303")
```

## License
The MIT License (MIT)

Copyright (c) 2018 retgits

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.