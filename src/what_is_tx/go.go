db.Exec(ctx, `update users set name = $1 where id = $2;`, "Admin", 1)

db.Exec(ctx, `insert into events (type) values ($1);`, "user_renamed")