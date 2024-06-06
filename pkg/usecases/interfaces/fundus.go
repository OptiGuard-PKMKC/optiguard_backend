package usecase_intf

type FundusUsecase interface {
	DetectImage() error
	FundusHistory() error
	ViewFundus() error
	VerifyFundusByDoctor() error
	DeleteFundus() error
}
