accion ej1 es
ambiente
	anual = registro
		clave = registro 
			codigoPcia: N(4)
			nroCentral: N(10)
		fr
		cantVoz: N(10)
		cantDatos: N(10)
		cantSMS: N(10)
	fr

	archMae, archAct: archivo de anual ordenado por clave
	mae, aux, act: anual

	comunicaciones = registro
		clave = registro 
			codigoPcia: N(4)
			nroCentral: N(10)
		fr
		codMovimiento: {"alta", "modificacion"}
		tipoComunicacion: 1..33
		datosTransmitidos: N(10)
	fr

	archMov: archivo de comunicaciones ordenado por clave
	mov: comunicaciones

	provincias = registro
		codigoPcia: N(4)
		nombrePcia: AN(60)
	fr

	archProv: archivo de provincias indexado por codigoPcia
	prov: provincias

	procedimiento LeerMae() es
		Leer(archMae,mae)
		si FDA(archMae) entonces
			mae.clave:= HV
		fs
	fp

	procedimiento LeerMov() es
		Leer(archMov,mov)
		si FDA(archMov) entonces
			mov.clave:= HV
		fs
	fp

	procedimiento modificar() es
		si (mov.codMovimiento = "alta") entonces
			esc("Alta inválida, usuario existe")
		sino //modificacion
			segun aux.tipoComunicacion hacer
				1: aux.cantVoz + mov.datosTransmitidos
				2: aux.cantDatos + mov.datosTransmitidos
				3: aux.cantSMS + mov.datosTransmitidos
			fs
		fs
		LeerMov()
	fp


algoritmo

	abrirE/(archMae) ; abrirE/(archMov) ; abrir/S(archAct)
	abrirE/(archProv)

	LeerMae() ; LeerMov()

	mientras (mae.clave <> HV) O (mov.clave <> HV) hacer
		si (mae.clave > mov.clave) entonces
			act:= mae ; LeerMae() ; grabar(archAct,act)
		sino
			si (mae.clave = mov.clave) entonces
				aux:= mae
				mientras (aux.clave = mov.clave) hacer
					modificar()
				fm
				act:= aux ; grabar(archAct,act)
			sino
				//mae.clave < mov.clave
				si (mov.codMovimiento = "alta") entonces
					aux := mov //es campo por campo y dejar en 0 los de cant voz y demas
					mientras (aux.clave = mov.clave) hacer
						modificar()
					fm
					act:= aux ; grabar(archAct,act)
				sino
					esc("Modificacion inválida, usuario inexiste")
				fs
			fs
		fs
	fm

fa