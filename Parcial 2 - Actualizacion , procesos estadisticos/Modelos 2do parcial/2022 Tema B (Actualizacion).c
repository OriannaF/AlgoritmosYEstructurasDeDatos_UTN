Accion Figuritas_Ejer_1 ES

Ambiente
	
	Fecha=Registro
		AA: 2000...2100
		MM: 1...12
		DD: 1...31
	Fin Registro

	Album=Registro
		CodFig: N(5)
		Cantidad: N(2)
		PerDup: ("si","no")
	Fin Registro

	NewMae,Arch_A: Archivo secuencial de Album ordenado por CodFig
	reg_A,reg_S,aux: Album

	Intercambio=Registro
		CodFig: N(5)
		CodAmig: N(5)
		FechSol: Fecha
	Fin Registro

	Arch_I: Archivo secuencial de Intercambio ordenado por CodFig y CodAmig
	reg_I: Intercambio

	Procedimiento Leer_Album() ES
		Leer(Arch_A,reg_A)
		Si FDS(reg_A) entonces
			reg_A.CodFig:= HV
		Fin Si
	Fin Procedimiento

	Procedimiento Leer_Inter() ES
		Leer(Arch_I,reg_I)
		Si FDS(reg_I) entonces
			reg_I.CodFig:= HV
		Fin Si
	Fin Procedimiento

	cont_dup: entero

Proceso
	Abrir E/(Arch_A)
	Leer_Album()
	Abrir E/(Arch_I)
	Leer_Inter()
	Abrir /S(NewMae)

	cont_dup:= 0

	Mientras (reg_A.CodFig <> HV) o (reg_I.CodFig <> HV) hacer

		Si (reg_A.CodFig < reg_I.CodFig) entonces

			reg_S:= reg_A
			Grabar(NewMae,reg_S)
			Leer_Album()
		SiNo
			Si (reg_A.CodFig = reg_I.CodFig) entonces

				aux_cod:= reg_I

				Si (reg_A.PerDup = "si") entonces

					Mientras (aux.CodFig = reg_I.CodFig) hacer
						aux.Cantidad:= aux.Cantidad + 1
						cont_dup:= cont_dup + 1

						Leer_Inter()
					Fin Mientras
				SiNo
					Mientras (aux.CodFig = reg_I.CodFig) hacer
						Leer_Inter()
					Fin Mientras
				Fin Si

				reg_S:= aux
				Grabar(NewMae,reg_S)
				Leer_Album()
			SiNo
				aux.CodFig:= reg_I.CodFig

				Si (diff_dias(7,reg_I.FechSol)) entonces

					aux.Cantidad:= 1
					aux.PerDup:= "no"

					reg_S:= aux
					Grabar(NewMae,reg_S)
				Fin Si

				Mientras (aux.CodFig = reg_I.CodFig) hacer
					Leer_Inter()
				Fin Mientras
			Fin Si
		Fin Si
	Fin Mientras 

	Esc("La cantidad de figuritas duplicadas que se admitieron fueron: ",cont_dup)

	Cerrar(Arch_A)
	Cerrar(Arch_I)
	Cerrar(NewMae)
Fin Accion 