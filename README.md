# mockSever

create table MockInfo(
  id  int not null auto_increment,
  ip varchar(32),
	host     varchar(32),
	requestUrl varchar(1024),
	mockResp   varchar(10240),
	status     varchar(2),
	primary key(id) 
)engine=INNODB  default charset=utf8 auto_increment=1;


drop table httpRequest;
create table httpRequest(
    id  int auto_increment not null ,
    uuid  varchar(32),
	method varchar(8),
	host varchar(32),
	remoteIp varchar(32),
	requestUrl varchar(1024),
	requestBoby varchar(1024),
	primary key(id) 
)engine=INNODB  default charset=utf8 auto_increment=1;

drop table httpRespone;
create table httpRespone(
    uuid  varchar(32), 
	responeBoby varchar(10240)
)engine=INNODB  default charset=utf8 auto_increment=1;
