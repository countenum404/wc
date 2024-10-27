# WC - print lines, words, chars counts

This is the one of the projects written in one evening, it buggy and not stable. But it's a good practice, instad of writing yet another REST API.   
The https://codingchallenges.fyi/challenges/challenge-wc challange inspired me to implement simple CLI tool.   
Now it counts a lines, visible chars, words in text files.
 

### Build executable binary

#### For windows:
```ps
    go build -o wc.exe .\cmd\main.go
```

#### For Unix:
```bash
    go build -o wc cmd/main.go
```
