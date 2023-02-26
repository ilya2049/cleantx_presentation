tx, _ := db.BeginTx(ctx, nil)

tx.Exec(ctx, `update users set name = $1 where id = $2;`, "Admin", 1)

tx.Exec(ctx, `insert into events (type) values ($1);`, "user_renamed")

tx.Commit()