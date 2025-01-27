package utils

func IsDeadlockError(err error) bool {

	if err == nil {
		return false
	}
	return err.Error() == "database is locked"
}
