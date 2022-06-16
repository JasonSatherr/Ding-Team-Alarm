package convenience

func HandlePotErr(err error) {
	if err != nil {
		panic(err)
	}
}
