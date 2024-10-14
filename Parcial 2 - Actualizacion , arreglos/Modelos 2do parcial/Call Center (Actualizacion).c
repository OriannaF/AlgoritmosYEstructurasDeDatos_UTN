Accion CallCenter_Act (H: arreglo [1...20] de entero) ES

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

	Arch_1: Archivo secuencial de Reclamos ordenado por Region y CodRecl
	reg_1: Reclamos

	Reporte=Registro
		Region: N(2)
		UltFecRec: Fecha
		UrgAlt: N(6)
		UrgMed: N(6)
		UrgBaj: N(6)
		NueAud: ("si","no")
	Fin Registro

	Arch_2: Archivo de Reporte indexado por Region
	reg_2: Reporte

	i,cont_aud: entero

Proceso
	Abrir E/(Arch_1)
	Leer(Arch_1,reg_1)
	Abrir E/S(Arch_2)

	cont_aud:= 0

	Mientras NFDA(Arch_1) hacer

		reg_2.Region:= reg_1.Region
		Leer(Arch_2,reg_2)
		Si (existe) entonces

			Mientras (reg_2.Region = reg_1.Region)

				Segun (reg_1.Urgencia) hacer
					"A": reg_2.UrgAlt:= reg_2.UrgAlt + 1
					"M": reg_2.UrgMed:= reg_2.UrgMed + 1
					"B": reg_2.UrgBaj:= reg_2.UrgBaj + 1
				Fin Segun 

				reg_2.UltFecRec:= reg_1.FecRecl

				Leer(Arch_1,reg_1)
			Fin Mientras

			Regrabar(Arch_2,reg_2)
		SiNo
			Esc("Error! Region ",reg_1.Region," no esta en el archivo")
		Fin Si

		Leer(Arch_1,reg_1)
	Fin Mientras
	
	Cerrar(Arch_1)
	Cerrar(Arch_2)

	Abrir E/S(Arch_2)
	Leer(Arch_2,reg_2)

	Mientras NFDA(Arch_1) hacer
		i:= reg_2.Region

		Si (H[i] < 10) y ((reg_2.UrgBaj * 2) < reg_2.UrgAlt) entonces

			Esc("Se requiere auditoria en la region: ",i)

			H[i]:= H[i] + 1

			cont_aud:= cont_aud + 1
		SiNo
			Esc("No se requiere auditoria en la region ",i)
		Fin Si
	Fin Mientras

	Esc("Se solicitaron ",cont_aud," auditorias en total")
	Cerrar(Arch_2)
Fin Accion 