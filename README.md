# {{ .appName }}

[Writing Web Applications](https://golang.org/doc/articles/wiki/)

## Setup Local Development.

You can use the Docker environment to get going if you have Docker on your
computer.

You will need to set the `$Env:HOME` environment on Windows before you can start
the dev environment. It should be the same value as `$Env:USERPROFILE`, or
what you need it to be. Once you have it set and it shows up in Powershell when
you run `Get-ChildItem Env:`; you then do one of the following.

### Run Docker

1. Clone this repository.
2. `cd` into the clone directory and run: `docker-compose up`
3. Open another command prompt to login to the container: 
   `docker exec -it {{ .appName }} sh`
4. Execute a command such as `go test`
   ```output
   ~/src/github.com/kohirens/stdlib $ go test
   PASS
   ok      github.com/kohirens/stdlib      0.004s
   ```
5. You can now open your IDE and point it to the clone directory and begin coding.

### Run with VS Code

1. Install the VS code extention "Remote Container".
2. Clone this repository locally.
3. Open the cloned project in VS Code, which should ask to open the folder in a
   remote container.
4. Open a termianl in VS Code and type `go test`.
5. You can now open your IDE and point it to the clone directory and begin coding.
