package talk_training

import (
    _ "github.com/lib/pq"
)

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
