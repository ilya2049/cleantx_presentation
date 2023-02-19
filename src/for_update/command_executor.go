func (e *DoctorCommandExecutor) FinishShift(ctx context.Context, doctorID int) error {
	// Елвин и Ада одновременно хотят закончить смену.
	// ListDoctorsOnCall считает два доктора в каждом запросе.
	doctorsOnCall, err := e.repository.ListDoctorsOnCall(ctx)
	if err != nil {
		return err
	}

	doctor := doctorsOnCall.Get(doctorID)
	if doctor == nil {
		return domain.ErrDoctorNotFound
	}

	// Затем, у каждого запроса проверяется количество докторов на смене.
	// И каждый запрос считал по два доктора!
	if err := doctor.FinishShift(len(doctorsOnCall)); err != nil {
		return err
	}

	// Теперь Елвин и Ада заканчивают смены.
	// Все доктора закончили смены!
	if err := e.repository.Update(ctx, doctor); err != nil {
		return err
	}

	return nil
}
