package chessboard

type File = []bool

type Chessboard = map[string]File

const BoardSize = 8

func CountInFile(cb Chessboard, file string) int {
	sum := 0
    for i:= 0; i < BoardSize; i++ {
        if v, ok := cb[file]; ok && v[i] {
            sum++
        }
    }
    return sum
}

func CountInRank(cb Chessboard, rank int) int {
    if rank < 1 || rank > BoardSize {
        return 0
    }
	sum := 0
    for _, f := range cb {
        if f[rank-1] {
            sum++
        }
    }
    return sum
}

func CountAll(cb Chessboard) int {
	sum := 0
    for _, f := range cb {
        for range f {
            sum++
        }
    }
    return sum
}

func CountOccupied(cb Chessboard) int {
	sum := 0
    for _, f := range cb {
        for _, cell := range f {
            if cell {
                sum++
            }
        }
    }
    return sum
}
