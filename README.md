# lil-github-deployer

If you make a get request 4 things are done

1. `git pull` in the TARGET_DIR 
2. All `tmux` sessions are passed CTRL-C
3. A new `tmux` session is created
4. The `shell script` is executed in the new session


```bash
go run main.go NAME ../TARGET_DIR YOUR_SHELL_LOCATION
```