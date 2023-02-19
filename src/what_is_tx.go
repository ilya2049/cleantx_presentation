// Получаем объект транзакции, чтобы делать запросы к БД в рамках транзакции.
tx, err := db.BeginTx(ctx, nil)
if err != nil {
	return err
}

// Откатываем изменения, если возникла ошибка.
defer tx.Rollback()

// Фиксируем транзакцию.
if err = tx.Commit(); err != nil {
	return err
}