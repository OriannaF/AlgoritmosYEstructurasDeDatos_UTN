Accion Figuritas_Ejer_2 ES

Ambiente
	
	Album=Registro
		CodUser: N(5)
		CodFig: N(5)
		Cantidad: N(2)
		Tipo: ("D","C","V")
	Fin Registro

	Arch_Al: Archivo secuencial de Album ordenado por CodUser
	regAl: Album

	Amigos=Registro
		CodUser: N(5)
		Apellido: AN(30)
		Nombre: AN(30)
		Celular: N(10)
	Fin Registro

	Arch_AM: Archivo de Amigos indexado por CodUser
	regAM: Amigos

	Matriz: arreglo [1...11,1...4] de entero

	i,j: entero

	res_CodUser,mayor_cantidad: entero
	res_appellido,res_nombre: alfanumerico

	Procedimiento Corte_User() ES
		Si (Matriz[i,4] > mayor_cantidad) entonces

			regAM.CodUser:= regAl.CodUser
			Leer(Arch_AM,regAM)
			Si (existe) entonces

				res_appellido:= regAM.Apellido
				res_nombre:= regAM.Nombre

				mayor_cantidad:= Matriz[i,4]
			Fin Si
		Fin Si

		i:= i + 1
	Fin Procedimiento

Proceso
	Abrir E/(Arch_Al)
	Leer(Arch_Al,regAl)
	Abrir E/(Arch_AM)

	Para (i:=1 hasta 11) hacer
		Para (i:=1 hasta 4) hacer
			Matriz[i,j]:= 0
		Fin Para
	Fin Para

	res_CodUser:= regAl.CodUser
	i:= 1
	mayor_cantidad:= LV

	Mientras NFDA(Arch_Al) hacer
		Si (res_CodUser <> regAl.CodUser) entonces
			Corte_User()
		Fin Si

		Segun (regAl.Tipo) hacer
			"D": j:=1
			"C": j:=2
			"V": j:=3
		Fin Segun

		Matriz[i,j]:= Matriz[i,j] + regAl.Cantidad
		Matriz[i,4]:= Matriz[i,4] + regAl.Cantidad
		Matriz[11,j]:= Matriz[11,j] + regAl.Cantidad
		Matriz[11,4]:= Matriz[11,4] + regAl.Cantidad

		Leer(Arch_Al,regAl)
	Fin Mientras 

	Corte_User()

	Cerrar(Arch_Al)
	Cerrar(Arch_AM)

	Esc("Usuario	Album Dorado	Album Comun		Album Virtual		Total x Usuario")
	Para (i:=1 hasta 11) hacer
		Si (i < 11) entonces
			Esc(i)
		SiNo
			Esc("Total x Album")
		Fin Si

		Para (i:=1 hasta 4) hacer
			Esc(Matriz[i,j])
		Fin Para
	Fin Para

	Esc("El usuario que mas figuritas coleccionó fue: ",res_nombre," ",res_appellido)
	Para (i:=1 hasta 10) hacer
		Esc("Usuario ",i,"; porsentaje de figuritas: ", (Matriz[i,4] / 10) * 100)
	Fin Para
Fin Accion 