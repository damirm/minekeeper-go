# Minekeeper Game

## Quickstart
```bash
./build.sh && ./main
```

## Keys

| Key | Description |
|-----|-------------|
| <kbd>h</kbd>   | Move left   |
| <kbd>j</kbd>   | Move down   |
| <kbd>k</kbd>   | Move up   |
| <kbd>l</kbd>   | Move right   |
| <kbd> </kbd> | open cell |
| <kbd>m</kbd> | mark cell as bomb (symbol "!"), mark cell as unknown (symbol "?") |
| <kbd>q</kbd> | quit |

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
