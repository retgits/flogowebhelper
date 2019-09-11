# Flogo Web UI - Helper

_This code is no longer maintained. Feel free to fork and use it if you want though :smile:_

Flogo Web UI - Helper is a utility to help work with the Flogo Web UI. Especially tasks like importing and exporting apps are considerably easier...

You can build an executable out of this app using the command
```
go build
```

## Usage
```
CLI to interact with the Project Flogo Web UI

Usage:
  flogowebhelper [command]

Available Commands:
  apps        Apps management for the Project Flogo Web UI
  docker      Docker container management for the Project Flogo Web UI
  help        Help about any command

Flags:
  -h, --help   help for flogowebhelper

Use "flogowebhelper [command] --help" for more information about a command.
```

### apps
```
The Apps command supports the app management capabilities.
The commands available are:

export................... Exports Flogo apps from the Flogo Web UI
import................... Imports Flogo apps into the Flogo Web UI
```

### apps - export
```
Exports Flogo apps from the Flogo Web UI

Usage:
  flogowebhelper apps export [flags]

Flags:
  -h, --help          help for export
      --host string   The URL for the Flogo Web UI (default "http://localhost:3303")
```

### apps - import
```
Imports Flogo apps into the Flogo Web UI

Usage:
  flogowebhelper apps import [flags]

Flags:
      --dir               import all JSON files in the current directory
      --filename string   The name of the file you want to import if you do not specify 'dir' (default "flogo.json")
  -h, --help              help for import
      --host string       The URL for the Flogo Web UI (default "http://localhost:3303")
```

### docker
```
The Docker command supports the container management capabilities.
The commands available are:

build.................... Builds a new docker image
latest................... Pulls the latest version of the Flogo Web UI from Docker Hub
start.................... Starts a new instance of the Flogo Web UI with default settings
```

### docker - start
```
Starts a new instance of the Flogo Web UI with default settings

Usage:
  flogowebhelper docker start [flags]

Flags:
  -h, --help           help for start
      --image string   The image name for the Flogo Web UI container (default "flogo/flogo-docker")
```

### docker - latest
```
Pulls the latest version of the Flogo Web UI from Docker Hub

Usage:
  flogowebhelper docker latest [flags]

Flags:
  -h, --help   help for latest
```

### docker - build
```
Builds a new docker image

Usage:
  flogowebhelper docker build [flags]

Flags:
  -h, --help             help for build
      --image string     The image name for the Flogo Web UI container (default "flogo/flogo-docker")
      --imports string   An imports file in case you want to add additional activities (like /home/user/Downloads/imports.go)
```

_Check out [this example](https://github.com/retgits/dockerfiles/blob/master/flogoweb/imports.go) of an imports.go_

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
