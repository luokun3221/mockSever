package dbOpr

import (
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
   )


func CreateDbLink()(error){
   var err error = nil
   Db,err = sql.Open("mysql",DBConf)
   Db.SetMaxOpenConns(2000)
   Db.SetMaxIdleConns(1000)
   Db.Ping()
   if err != nil {
      return err
   }
   return nil
}

