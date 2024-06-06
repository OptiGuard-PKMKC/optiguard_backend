package repo_intf

type FundusRepository interface {
	Create() error
	FindAll() error
	FindByID() error
	Delete() error
	UpdateVerifyDoctor() error
}
