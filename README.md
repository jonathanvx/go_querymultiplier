# querymultiplier
Golang tools to help with maintaining MySQL

Usage:
```
$ ./querymultiplier "insert ignore into table select * from table2" 10 301 LIMIT
insert ignore into table select * from table2 LIMIT 0,10;
insert ignore into table select * from table2 LIMIT 10,10;
insert ignore into table select * from table2 LIMIT 20,10;
insert ignore into table select * from table2 LIMIT 30,10;
insert ignore into table select * from table2 LIMIT 40,10;
insert ignore into table select * from table2 LIMIT 50,10;
insert ignore into table select * from table2 LIMIT 60,10;
insert ignore into table select * from table2 LIMIT 70,10;
insert ignore into table select * from table2 LIMIT 80,10;
insert ignore into table select * from table2 LIMIT 90,10;
insert ignore into table select * from table2 LIMIT 100,10;
insert ignore into table select * from table2 LIMIT 110,10;
insert ignore into table select * from table2 LIMIT 120,10;
insert ignore into table select * from table2 LIMIT 130,10;
insert ignore into table select * from table2 LIMIT 140,10;
insert ignore into table select * from table2 LIMIT 150,10;
insert ignore into table select * from table2 LIMIT 160,10;
insert ignore into table select * from table2 LIMIT 170,10;
insert ignore into table select * from table2 LIMIT 180,10;
insert ignore into table select * from table2 LIMIT 190,10;
insert ignore into table select * from table2 LIMIT 200,10;
insert ignore into table select * from table2 LIMIT 210,10;
insert ignore into table select * from table2 LIMIT 220,10;
insert ignore into table select * from table2 LIMIT 230,10;
insert ignore into table select * from table2 LIMIT 240,10;
insert ignore into table select * from table2 LIMIT 250,10;
insert ignore into table select * from table2 LIMIT 260,10;
insert ignore into table select * from table2 LIMIT 270,10;
insert ignore into table select * from table2 LIMIT 280,10;
insert ignore into table select * from table2 LIMIT 290,10;
insert ignore into table select * from table2 LIMIT 300,1;
```

```
$ ./querymultiplier "insert ignore into table select * from table2" -rows=10 -min=1 -max=301 -keytype=id
insert ignore into table select * from table2 WHERE id between 1 and 11;
insert ignore into table select * from table2 WHERE id between 11 and 21;
insert ignore into table select * from table2 WHERE id between 21 and 31;
insert ignore into table select * from table2 WHERE id between 31 and 41;
insert ignore into table select * from table2 WHERE id between 41 and 51;
insert ignore into table select * from table2 WHERE id between 51 and 61;
insert ignore into table select * from table2 WHERE id between 61 and 71;
insert ignore into table select * from table2 WHERE id between 71 and 81;
insert ignore into table select * from table2 WHERE id between 81 and 91;
insert ignore into table select * from table2 WHERE id between 91 and 101;
insert ignore into table select * from table2 WHERE id between 101 and 111;
insert ignore into table select * from table2 WHERE id between 111 and 121;
insert ignore into table select * from table2 WHERE id between 121 and 131;
insert ignore into table select * from table2 WHERE id between 131 and 141;
insert ignore into table select * from table2 WHERE id between 141 and 151;
insert ignore into table select * from table2 WHERE id between 151 and 161;
insert ignore into table select * from table2 WHERE id between 161 and 171;
insert ignore into table select * from table2 WHERE id between 171 and 181;
insert ignore into table select * from table2 WHERE id between 181 and 191;
insert ignore into table select * from table2 WHERE id between 191 and 201;
insert ignore into table select * from table2 WHERE id between 201 and 211;
insert ignore into table select * from table2 WHERE id between 211 and 221;
insert ignore into table select * from table2 WHERE id between 221 and 231;
insert ignore into table select * from table2 WHERE id between 231 and 241;
insert ignore into table select * from table2 WHERE id between 241 and 251;
insert ignore into table select * from table2 WHERE id between 251 and 261;
insert ignore into table select * from table2 WHERE id between 261 and 271;
insert ignore into table select * from table2 WHERE id between 271 and 281;
insert ignore into table select * from table2 WHERE id between 281 and 291;
insert ignore into table select * from table2 WHERE id between 291 and 301;
```
