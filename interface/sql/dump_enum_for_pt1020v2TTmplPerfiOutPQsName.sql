--
-- Dump all sysName from DB
--
select ENUM.NAME,MIB.NAME,OBJECT.NAME,ENUM.VALUE from ENUM inner join OBJECT,MIB on ENUM.ID=OBJECT.EKEY and MIB.ID=OBJECT.MIB  where OBJECT.NAME='pt1020v2TTmplPerfiOutPQsName';
