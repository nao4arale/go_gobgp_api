package client

import (
        "gopkg.in/cheggaaa/pb.v1"
        "time"
)

func prg() {
        count := 100
        bar := pb.StartNew(count)
        for i := 0; i < count; i++ {
                bar.Increment()
                time.Sleep(10 * time.Millisecond)
        }
        bar.FinishPrint(" Check is done.")
}
