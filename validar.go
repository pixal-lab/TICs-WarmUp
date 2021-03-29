package main

func RegistroUsuario(user string, password string) bool {
	if existUser(user) {
		return false
	}
	addUser(user, password)
	return true
}

func validarDatosCalculo(monto int, cuota int, periodo int) bool {
	if monto <= 0 || cuota <= 0 || periodo < 2 || periodo > 48 || periodo*cuota < monto || float64(monto)/float64(cuota) <= 1 {
		return false
	}
	return true
}
