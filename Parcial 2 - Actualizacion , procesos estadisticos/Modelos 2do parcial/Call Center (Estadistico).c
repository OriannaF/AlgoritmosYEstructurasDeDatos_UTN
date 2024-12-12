Accion CallCente_Arreglo ES

Ambiente
	
	Fecha=Registro
		AA: 1900...2100
		MM: 1...12
		DD: 1...31
	Fin Registro

	Reclamos=Registro
		Region: N(2)
		CodRecl: N(10)
		FecRecl: Fecha
		MailCliente: AN(20)
		Urgencia: AN(1)
		Detalle: AN(100)
	Fin Registro

	Arch: Archivo secuencial de Reclamos ordenado por Region
	reg: Reclamos

	Matriz: arreglo [1...13,1...4]

	i,j,Enero_Alt,Urg_men10_Mes: entero

Proceso
	Abrir E/(Arch)
	Leer(Arch,reg)

	Para (i:=1 hasta 13) hacer
		Para (j:=1 hasta 4) hacer
			Matriz[i,j]:= 0
		Fin Para
	Fin Para

	Mientras NFDA(Arch) hacer
		i:= reg.FecRecl.MM

		Segun (reg.Urgencia) hacer
			"A": j:= 1
			"M": j:= 2
			"B": j:= 3
		Fin Segun

		Matriz[i,j]:= Matri[i,j] + 1
		Matriz[i,4]:= Matri[i,4] + 1
		Matriz[13,j]:= Matri[13,j] + 1
		Matriz[13,4]:= Matri[13,4] + 1

		Leer(Arch,reg)
	Fin Mientras

	Cerrar(Arch)

	Enero_Alt:= 0
	Urg_men10_Mes:= 0
	
	Esc("Meses		Urgencia Alta 	Urgencia Media 	Urgencia Baja	Total por mes")

	Para (i:=1 hasta 13) hacer
		Si (i < 13) entonces	
			Esc("Mes ",i)
		SiNo
			Esc("Total por Urgencia")
		Fin Si

		Para (i:=1 hasta 4) hacer

			Esc(Matriz[i,j])

			Si (i = 1) y (j = 1) entonces
				Enero_Alt:= Enero_Alt + Matriz[i,j]
			Fin Si

			Si (j = 2) y (Matriz[i,j] < 10) entonces
				Urg_men10_Mes:= Urg_men10_Mes + 1
			Fin Si
		Fin Para

		Esc("En Enero se registraron ",Enero_Alt," urgencias Altas")
		Esc("La cantidad de meses que registraron menos de 10 urgencias medias fueron: ",Urg_men10_Mes)
	Fin Para
Fin Accion 