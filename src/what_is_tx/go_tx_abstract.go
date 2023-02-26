// abstract begin ?

userRepository.Update(ctx, User{ID: 1, Name: "Admin"})

eventPublisher.Publish(ctx, Event{Type: "user_renamed"})

// abstract commit ?