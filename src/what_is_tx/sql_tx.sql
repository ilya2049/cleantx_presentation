begin;

update users set name = 'Admin' where id = 1;

insert into events (type) values ('user_renamed');

commit;