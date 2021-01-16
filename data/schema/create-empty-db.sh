#!/bin/sh

NRISCHEMAPATH=`dirname $0`

if [[ $1 == "" ]]; then
  echo "$0 <database filename>"
  exit 1
fi
if [ ! -f $NRISCHEMAPATH/nri-snmpdb.sql ]; then
  echo "nri-snmpdb.sql not in the current directory"
  exit 1
fi
if [ -f $1 ]; then
  echo "Removing old database in: $1"
  rm -f $1
fi
cat nri-snmpdb.sql | sqlite3 $1
