package main

import (
	"database/sql"
	"flag"
	"fmt"
	"mysql/ex/config"
	"mysql/ex/logic"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

var configFile = flag.String("config", "./etc/mysql.yaml", "config file")

func main() {
	flag.Parse()
	cfg, err := config.Load(*configFile)
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		fmt.Printf("Open Mysql error: %v\n", err)
		return
	}
	defer db.Close()

	l := logic.NewLogic(db)
	results, err := l.Select([]string{"rdID", "rdName"}, "reader")
	if err != nil {
		fmt.Println(err)
		return
	}
	PrintResult(results)

	l.Txn([]logic.TxItem{{
		Sql: "update reader set rdType = 1 where rdName like \"__\"",
	}, {
		Sql: "update reader set rdType = 2 where rdName like \"___\"",
	},
	})
}

func PrintResult(data *sql.Rows) {
	defer data.Close()
	fmt.Println("===================")
	columnsTypes, _ := data.ColumnTypes()
	fmt.Printf("Columns Types:")
	var args []any
	for _, item := range columnsTypes {
		fmt.Printf("%s-%s\t", item.Name(), item.DatabaseTypeName())
		inter := reflect.New(item.ScanType()).Interface()
		args = append(args, inter)
	}
	fmt.Println()
	cnt := 0
	for data.Next() {
		data.Scan(args...)
		fmt.Printf("Result %d:", cnt)
		for _, val := range args {
			v := reflect.Indirect(reflect.ValueOf(val)).Interface()
			fmt.Printf("%v \t", v)
		}
		fmt.Println()
		cnt++
	}
}
