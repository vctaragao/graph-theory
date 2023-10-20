package helpers

func Chk(err error) {
	if err != nil {
		panic(err)
	}
}
