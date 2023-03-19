### Todo-CLI

### How to use

#### Add a new todo item
```
$ go run . add "wash dish"
2023/03/18 22:09:36 Todo item added.
```

#### List todo items with limit(default 3)
```
$ go run . list
2023/03/19 15:01:43 Todo list:
2023/03/19 15:01:43 1. go to sleep
2023/03/19 15:01:43 2. take a shower

go run . list -l 4
2023/03/19 15:04:33 Todo list:
2023/03/19 15:04:33 1. go to sleep
2023/03/19 15:04:33 2. take a shower
2023/03/19 15:04:33 4. go to office
2023/03/19 15:04:33 5. pay phone bill
```

#### Search for todo items containing a specific string
```
$ go run . search take
2023/03/19 15:02:03 Todo items containing 'take':
2023/03/19 15:02:03 2. take a shower
```

### Build guide

#### Preparation
- Install cobra-cli to generate cobra applications and command files
- `go install github.com/spf13/cobra-cli@latest`

#### Generate a cobra application
```
mkdir <repo name>
cd <repo name>
go mod init github.com/<your github name>/<repo name>
cobra-cli init
```

#### Generate cobra command files
```
cobra-cli add add
cobra-cli add list
cobra-cli add search
```

#### Work with flags
- https://github.com/spf13/cobra/blob/main/user_guide.md#working-with-flags

#### Try out each command
- use `go run .` + a specific command
- For example
  - `go run . list`
  - `go run . list --list 4` -> with flag
