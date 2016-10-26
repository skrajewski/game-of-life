package main;

import (
    "math/rand"
    "time"
    "bufio"
    "os"
    "strconv"
    tm "github.com/buger/goterm"
)

func printWorld(board [][]int) {
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            tm.MoveCursor(i+3, j+1);

            if board[i][j] == 0 {
                tm.Print(tm.Background(" ", tm.RED));
            } else {
                tm.Print(tm.Background(" ", tm.GREEN));
            }
        }
    }
}

func nextGeneration(board [][]int) [][]int {
    next := make([][]int, len(board));

    for i := 0; i < len(board); i++ {
        next[i] = make([]int, len(board[i]));

        for j := 0; j < len(next[i]); j++ {
            x := nextCellState(board, i, j);
            next[i][j] = x;
        }
    }

    return next;
}

func nextCellState(board [][]int, r int, c int) int {
    height := len(board);
    width := len(board[0]);
    state := board[r][c];
    neighbours := 0;

    for x := Max(0, r-1); x <= Min(r+1, height-1); x++ {
        for y := Max(0, c-1); y <= Min(c+1, width-1); y++ {
            if x != r || y != c {
                if (board[x][y] > 0) {
                    neighbours++
                }
            }
        }
    }

    if state == 0 && neighbours == 3 {
        return 1;
    }

    if state > 0 && neighbours >= 2 && neighbours <= 3 {
        return 1;
    }

    return 0;
}

func main()  {
    rand.Seed(time.Now().UTC().UnixNano());
    seed := rand.Int63();
    rand.Seed(seed);

    args := os.Args[1:];

    width := 80;
    height := 22;
    counter := 0;
    timer := time.Now().UTC();

    if len(args) > 0 {
        width, _ = strconv.Atoi(args[0]);
    }

    if len(args) > 1 {
        height, _ = strconv.Atoi(args[1]);
    }

    board := make([][]int, height);

    for i := 0; i < height; i++ {
        board[i] = make([]int, width);
        for j := 0; j < len(board[i]); j++ {
            if rand.Intn(8) == 0 {
                board[i][j] = 1;
            }
        }
    }

    tm.Clear();

    ticker := time.NewTicker(time.Millisecond * 100);

    go func() {
        for range ticker.C {
            tm.MoveCursor(0, 0);
            tm.Println("Generation: ", counter, "| Seed: ", seed, "| Time: ", time.Since(timer));
            tm.Println("Press [Enter] to finish");

            counter++;

            printWorld(board);
            board = nextGeneration(board);
            tm.Flush();
        }
    }()

    reader := bufio.NewReader(os.Stdin)
    reader.ReadString('\n')
    ticker.Stop();
}

func Min(x, y int) int {
    if x < y {
        return x;
    }
    return y;
}

func Max(x, y int) int {
    if x > y {
        return x;
    }
    return y;
}
