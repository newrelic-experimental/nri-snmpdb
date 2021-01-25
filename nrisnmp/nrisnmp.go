package nrisnmp

import (
  "os"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "github.com/Jeffail/gabs"
)

var snmpdb = ""

type SNMPDB struct {
  TDB   *sql.DB
}

func Init() {
  if os.Getenv("NRISNMPDB") == "" {
    snmpdb = "/usr/share/nrisnmpdb/nrisnmp.db"
  } else {
    snmpdb = os.Getenv("NRISNMPDB")
  }
}

func DBName() string {
  return snmpdb
}

func SetDBName(n string) bool {
  _, err := os.Stat(n)
  if os.IsNotExist(err) {
      return false
  }
  snmpdb = n
  return true
}

func Open() *SNMPDB {
  var db SNMPDB
  if snmpdb == "" {
    return nil
  }
  _, err := os.Stat(snmpdb)
  if os.IsNotExist(err) {
    return nil
  }
  db.TDB, err = sql.Open("sqlite3", snmpdb)
  if err != nil {
    return nil
  }
  return &db
}

func (s *SNMPDB) Symbol(oid string) (res *gabs.Container, err error) {
  res = gabs.New()
  res.Array("data")
  if s.TDB == nil {
    return
  }
  stmt, err := s.TDB.Prepare("select distinct OBJECT.NAME from OBJECT where OBJECT.KEY= ?")
  if err != nil {
    return
  }
  rows, err := stmt.Query(oid)
  if err != nil {
    return
  }
  for rows.Next() {
    var n string
    rows.Scan(&n)
    res.ArrayAppend(n, "data")
  }
  return
}

func (s *SNMPDB) Close() bool {
  if s.TDB != nil {
    s.TDB.Close()
    return true
  }
  return false
}
