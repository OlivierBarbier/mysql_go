package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type DbWrapper struct {
    db *sql.DB
    insert *sql.Stmt
}

func main() {
    const cnxNum int =  10

    var dbs [cnxNum]DbWrapper

    for i:=0; i<cnxNum; i++ {
        var db *sql.DB
        var insert *sql.Stmt
        var err error

        db, err = sql.Open("mysql", "root:root@/tiller")
        if err != nil {
            panic(err.Error())
        }
        defer db.Close()

        insert, err = db.Prepare("INSERT INTO squareNum (val) VALUES(?)")
        if err != nil {
            panic(err.Error())
        }
        defer insert.Close()
        
        dbs[i] = DbWrapper{db, insert}

    }

    fmt.Println("We are good to go ;)")

    c := make(chan int, cnxNum)

    size := 5000

    for i:=0; i<cnxNum; i++ {
        k := i
        go func () { 
            for j := 0*(size/cnxNum); j < 0*(size/cnxNum)+(size/cnxNum); j++ {
                _, err := dbs[k].insert.Exec(1)
                if err != nil {
                    panic(err.Error())
                }
            }
            c <- 1
        }()
    }

    for i:=0; i<cnxNum; i++ {
        <- c
        fmt.Print(".")
    }
    fmt.Println(" Done!")
}
