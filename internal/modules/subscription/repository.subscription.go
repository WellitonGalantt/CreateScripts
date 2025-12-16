package subscription

import "database/sql"

type SubscriptionRepository interface {
	Create(input CreateSubscriptionDtoInput) (*Subscription, error)
	Edit(input EditSubscriptionDtoInput) (*Subscription, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewSubscriptionRepository(db *sql.DB) SubscriptionRepository {
	return &postgresRepository{
		db,
	}
}

func (r *postgresRepository) Create(input CreateSubscriptionDtoInput) (*Subscription, error) {
	return nil, nil
}

func (r *postgresRepository) Edit(input EditSubscriptionDtoInput) (*Subscription, error) {
	return nil, nil
}
