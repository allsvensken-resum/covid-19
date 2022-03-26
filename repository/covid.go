package repository

type CovidPatient struct {
	ConfirmDate     *string
	No              *int
	Age             *int
	Gender          *string
	GenderEn        *string
	Nation          *string
	NationEn        *string
	Province        *string
	ProvinceId      *int
	District        *string
	ProvinceEn      *string
	StateQuarantine *string
}

type ICovidRepo interface {
	GetAllCovidPatient() ([]CovidPatient, error)
}
