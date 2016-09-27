# go_mysql_tools
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
$ ./querymultiplier "insert ignore into table select * from table2" 10 301 id
insert ignore into table select * from table2 WHERE id between 1 and 10;
insert ignore into table select * from table2 WHERE id between 11 and 20;
insert ignore into table select * from table2 WHERE id between 21 and 30;
insert ignore into table select * from table2 WHERE id between 31 and 40;
insert ignore into table select * from table2 WHERE id between 41 and 50;
insert ignore into table select * from table2 WHERE id between 51 and 60;
insert ignore into table select * from table2 WHERE id between 61 and 70;
insert ignore into table select * from table2 WHERE id between 71 and 80;
insert ignore into table select * from table2 WHERE id between 81 and 90;
insert ignore into table select * from table2 WHERE id between 91 and 100;
insert ignore into table select * from table2 WHERE id between 101 and 110;
insert ignore into table select * from table2 WHERE id between 111 and 120;
insert ignore into table select * from table2 WHERE id between 121 and 130;
insert ignore into table select * from table2 WHERE id between 131 and 140;
insert ignore into table select * from table2 WHERE id between 141 and 150;
insert ignore into table select * from table2 WHERE id between 151 and 160;
insert ignore into table select * from table2 WHERE id between 161 and 170;
insert ignore into table select * from table2 WHERE id between 171 and 180;
insert ignore into table select * from table2 WHERE id between 181 and 190;
insert ignore into table select * from table2 WHERE id between 191 and 200;
insert ignore into table select * from table2 WHERE id between 201 and 210;
insert ignore into table select * from table2 WHERE id between 211 and 220;
insert ignore into table select * from table2 WHERE id between 221 and 230;
insert ignore into table select * from table2 WHERE id between 231 and 240;
insert ignore into table select * from table2 WHERE id between 241 and 250;
insert ignore into table select * from table2 WHERE id between 251 and 260;
insert ignore into table select * from table2 WHERE id between 261 and 270;
insert ignore into table select * from table2 WHERE id between 271 and 280;
insert ignore into table select * from table2 WHERE id between 281 and 290;
insert ignore into table select * from table2 WHERE id between 291 and 300;
insert ignore into table select * from table2 WHERE id between 301 and 301;
```
