Accion ej1 es
ambiente

	fecha = registro
		dia: N(2)
		mes: N(2)
		anio: N(4)
	fr

	usuarios = registro
		clave = registro
			codRegion: N(4)
			codUsuario: N(10)
		fr
		correo: AN(50)
		experiencia: N(7)
	fr

	archMae, archSal: archivo de usuarios ordenado por clave
	mae,mov,act: usuarios

	capturas = registro
		clave = registro
			codRegion: N(4)
			codUsuario: N(10)
			codPokemon: 1...151
		fr
		PE: N(5)
		fechaCaptura: fecha
		estadoPokemon: ("E","I","D") //entrenandose, incubandose, descansando
		estadoUsuario: ("A","S","B") //activo, suspendido, baneado
	fr
	
	archMov: archivo de usuarios ordenado por clave
	mov: usuarios

	A: arreglo de [1...151] de alfanumerico
	Pokemon: arreglo de [1...151] de entero //para ver el pokemon descansando con mas uso
	i,mayor,posicionMayor: entero
	usuario: booleano

	procedimiento LeerMae() es
		leer(archMae,mae)
		si FDA(archMae) entonces
			mae.clave:= HV
		fs
	fp

	procedimiento LeerMov() es
		leer(archMov,mov)
		si FDA(archMov) entonces
			mov.clave:= HV
		fs
	fp

	procedimiento ProcesosIguales() es
		//alta ya se hizo, pero si el estado esta en activo le sumo experiencia
		si (mov.estadoUsuario = "A") entonces
			//punto 1, si estado = E, se duplican los puntos
			si (mov.estadoPokemon = "E") entonces
				aux.experiencia := aux.experiencia + ( mov.codPokemon * 2 )
			sino
				aux.experiencia := aux.experiencia + mov.codPokemon
			fs
			//punto 3, pokemon con más estado en descanso
			//nombre del pokemon y cantidad de personas que lo tienen en ese estado
			si (mov.estadoPokemon = "D") entonces
				i:= mov.codPokemon
				Pokemon[i] := Pokemon[i] + 1
			fs

		sino
			si (mov.estadoUsuario = "S") entonces
				usuario := falso //para que no se actualice la salida
			sino
				si (mov.estadoUsuario = "B") entonces
					//no se que hacer si está baneado, en todo caso capaz tampoco se pone en salida y entonces no estaria preguntando el estado igua a s o b, si no es igual a A, no lo pongo en salida.
				fs
			fs
		fs
	fp
 
algoritmo
	mientras (mae.clave <> HV) O (mov.clave <> HV) hacer
		si (mae.clave < mov.clave) entonces
			//maestro a salida
			act:= mae ; grabar(archAct,act) ; LeerMae()
		sino
			si (mae.clave = mov.clave) entonces
				aux:= mae
				mientras (aux.clave = mov.clave) hacer
					ProcesosIguales() ; LeerMov()
				fm
				act:= aux; grabar(archAct,act)
				LeerMae()
			sino
				// mae.clave > mov.clave
				//voy a suponer que baneado no pasa nada pq no dice que hacer si esta baneado.
				si (mov.estadoUsuario = "S") entonces
					esc("baja inválida, usuario inexistente")
				sino
					//movimiento a auxiliar
					//solo se hace el alta cuando estado = activo
					//la modificacion se puede hacer solo si el usuario existe por eso esta dentro de esta condicion
					si (mov.estadoUsuario = "A") entonces

						aux.codRegion:= mov.codRegion
						aux.codUsuario:= mov.codUsuario
						aux.correo:= " "
			
						mientras (aux.clave = mov.clave) hacer
							ProcesosIguales() ; LeerMov()
						fm
						//si en algun momento el estadoUs = suspendido, debe darse de baja fisica, no tiene que aparecer en el actualizado directamente (creo)
						si (usuario = verdadero) entonces
							act:= aux ; grabar(archAct,act)
						fs
					fs
				fs
			fs
		fs
	fm

	para (i:= 1 hasta 151) hacer
		si (Pokemon[i] > mayor) entonces
			mayor:= Pokemon[i]
			posicionMayor:= i
		fs
	fp

	//punto 4 informar
	Esc("El pokemon que más usuarios tienen en estado de descanso es ", A[posicionMayor], "un total de ", mayor, " usuarios lo tienen en ese estado")
fa