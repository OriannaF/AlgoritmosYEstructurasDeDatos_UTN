Ambiente
	Maestro = registro
		clave = registro
			campoClave: dato
		fr
		campo1: dato
	fr
	archMae, archAct: archivo de Maestro ordenado por clave
	mae, aux, act: Maestro

	Movimiento = registro
		clave = registro
			campoClave: dato
		fr
		campo1: dato
	fr

	archMov: archivo de Movimiento ordenado por clave
	mov: Movimiento

	Procedimiento LeerMaestro() es
		Leer(archMae,mae)
		Si FDA(archMae) entonces
			mae.clave:= HV
		fs
	fp

	Procedimiento LeerMovimiento() es
		Leer(archMov,mov)
		Si FDA(archMov) entonces
			mov.clave:= HV
		fs
	fp

	Procedimiento Modificaciones() es
		Si mov.codigo = “ALTA” entonces
			Esc("Error, alta inválida, registro ya existente")
		sino
			Si mov.codigo = “BAJA” entonces
				//aca depende del ejercicio, suele ser una baja logica, donde se pone una marca como esta, o la fecha de baja
				aux.baja := "SI"
			sino
				Si mov.codigo = “MODIFICACION” entonces
					//modificaciones correspondientes... ej:
					aux.monto:= mov.monto
				fs
			fs
		fs
	fp

Algoritmo
	AbrirE/(archMae) ; AbrirE/(archMov) ; Abrir/S(archAct)
	LeerMae() ; LeerMov()

	Mientras ( mov.clave <> HV ) o ( mae.clave <> HV ) hacer
		Si mae.clave < mov.clave entonces
		//salida desde maestro, no hay modificaciones para este registro
			act:=mae
			Grabar(archAct,aux) ; LeerMaestro()
		sino
			Si mae.clave = mov.clave entonces
			//hay modificaciones para el registro
				Mientras mae.clave = mov.clave hacer
					Modificar() ; LeerMovimiento()
				fm
				Grabar(archAux,aux)
				LeerMaestro() 
			sino
				Si mae.clave > mov.clave entonces
				//no existe el reg en maestro
				//primero doy de alta desde movimiento
				//después pongo un mientras para aplicar modificaciones para ese registro que ahora si se dio de alta
					Si mov.codigo = “ALTA” entonces
						//aux:= mov campo por campo porque no suelen tener el mismo formato
						aux.clave:=mov.clave
						aux.campo1:=mov.campo1
						aux.campo2:=mov.campo2

						//ahora que ya se dio de alta, se hacen todas las modificaciones
						Mientras aux.clave = mov.clave hacer
							Modificar() ; LeerMovimiento()
						fm
						act:= aux
						Grabar(archAct, act)
					sino
						Si mov.codigo = “BAJA” entonces
							Esc(“Error, baja inválida, registro inexistente”)
						sino
							Si mov.codigo = “MODIFICACION” entonces
								Esc(“Error, modificación inválida, registro inexistente”)
							fs
						fs
					fs
					
				fs
			fs
		fs
	fm

	Cerrar(archMae) ; Cerrar(archMov) ; Cerrar(archAct)
Fa