package nrisnmp

import (
  "os"
  "fmt"
)

var snmpdb string

func Init() {
  if os.Getenv("NRISNMPDB") == nil {
    snmpdb = "/usr/share/nrisnmpdb/nrisnmp.db"
  } else {
    snmpdb = os.Getenv("NRISNMPDB")
  }
}

func DBName() string {
  return snmpdb
}

func SetDBName(n string) bool {
  info, err := os.Stat(n)
  if os.IsNotExist(err) {
      return false
  }
  snmpdb = n
  return true
}
