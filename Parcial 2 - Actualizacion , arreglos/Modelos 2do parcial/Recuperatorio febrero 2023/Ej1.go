Accion Ej1 es
Ambiente

	productos = registro
		clave = registro
			idProducto: N(10)
		fr
		nombre: AN(60)
		descripcion: AN(160)
		categoria: {"C", "R", "G", "S"}
		nuevoLanzamiento: {"Si", "No"}
		porcentajeDescuento: N(1,2)
		stock: N(10)
	fr

	archMae, archAct: archivo de productos ordenado por clave
	mae, aux, act: productos

	preventas = registro
		clave = registro
			idProducto: N(10) 
			idCliente: N(10)
		fr
		cantidad: N(10)
		esPersonalizado: {"Si", "No"}
		nroJugador: [1...26]
		nombreJugador: AN(60)
		talle: N(2)
	fr

	archMov: archivo de preventas ordenado por clave
	mov: preventas

	A: arreglo de alfanumerico de [1...26]
	i: entero

	procedimiento LeerMae() es
		Leer(archMae,mae)
		si FDA(archMae) entonces
			mae.clave := HV
		fs
	fp

	procedimiento LeerMov() es
		Leer(archMov,mov)
		si FDA(archMov) entonces
			mov.clave := HV
		fs
	fp

Algoritmo
	mientras (mov.clave <> HV) O (mae.clave <> HV) entonces
		si ( mae.clave < mov.clave) entonces
			//maestro a salida
			act := mae
			LeerMae() ; Grabar(achAct)
		sino
			si ( mae.clave = mov.clave ) entonces
				Mientras ( mae.clave = mov.clave ) hacer
					actualizarAux(); LeerMov()
				fm

				si (aux.stock <> 0) entonces
					act := aux ; Grabar(archAct)
				sino
					Esc("Error, falta de stock")
				fs
				LeerMae()
			sino
				si ( mae.clave > mov.clave ) entonces

					mientras ( aux.clave = mov.clave ) hacer
						//mientras el auxiliar tenga mov actualizo
						actualizarAux(); LeerMov()
					fs
					si (aux.stock <> 0) entonces
						act := aux ; Grabar(archAct)
					sino
						Esc("Error, falta de stock")
					fs
					LeerMov()
				fs
			fs
		fs
	fm
fa