# zen-go
A RESTFul Web Project Based On Golang Within gin And gorm

# run
```$xslt
    step 1: go mod tidy
    step 2: go build
    step 3: sudo chmod +x zen
    step 4: ./zen
    step 5: enjoy your RESTFul Http Apis
```


# data base
    author
```MYSQL
   CREATE TABLE `author` (
     `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
     `author_id` varchar(40) NOT NULL COMMENT '唯一id',
     `name` varchar(60) NOT NULL COMMENT '姓名',
     `sex` tinyint(4) NOT NULL COMMENT '性别: 1 男 2 女',
     `age` tinyint(4) DEFAULT NULL COMMENT '年龄',
     PRIMARY KEY (`id`),
     UNIQUE KEY `idx_only_id` (`author_id`)
   ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
```

    book
    
```MYSQL
CREATE TABLE `book` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(60) CHARACTER SET utf8mb4 NOT NULL,
  `book_id` varchar(40) NOT NULL,
  `author_id` varchar(40) DEFAULT NULL,
  `price` int(10) NOT NULL,
  `time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
```