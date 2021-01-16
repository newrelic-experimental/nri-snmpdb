--
-- Dump all sysName from DB
--
select MIB.NAME,OBJECT.KEY,OBJECT.NAME from OBJECT INNER JOIN MIB ON MIB.ID=OBJECT.MIB where OBJECT.NAME='sysName' ;
