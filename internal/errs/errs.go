package errs

type Err string

func (e Err) Error() string {
	return string(e)
}

const (
	ObjectNotFound   = Err("object_not_found")
	ServiceNA        = Err("service_na")
	NotAuthorized    = Err("not_authorized")
	SqlRequestNotCorrect = Err("SqlRequestNotCorrect")
	TableNotExist  = Err("TableNotExist or Not Find")
	IDNotFind = Err("IDNotFind")
	CantScanValues = Err("CantScanValues")
	ValuesNotFilled = Err("ValuesNotFilled")
	CantConnectToDb = Err("CantConnectToDb")
	CantDeleteThisID = Err("CantDeleteThisID")
	CantUpdate = Err("CantUpdate")
)