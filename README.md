# Minekeeper Game

## Quickstart
```bash
./build.sh && ./main
```

## Keyboard

* hjkl - movements
* space - open cell
* m - mark cell as bomb (symbol "!"), mark cell as unknown (symbol "?")
* q - quit

```
$ ./build.sh && ./main
+ go build -o main .
    2  .  .  .  .  .  .  .  .  3  1  1
    2  .  .  .  .  .  .  .  .  5  !  1
 2  3  .  .  .  .  .  .  .  .  !  3  1
 .  .  .  .  . [.] .  .  .  .  !  2
 .  .  .  .  .  .  .  .  .  .  3  1  1  1  1
 .  .  .  .  .  .  .  .  .  .  1     1  !  1
 .  .  .  .  .  .  .  .  .  .  2  1  3  2  2
 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
 .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
```