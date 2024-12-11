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
	B: arreglo de entero de [1...26]
	i,j: entero

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

	//Lo q no me queda claro: si la preventa es mayor del stock, le tomo como vendido lo que tengo de stock o directamente no se concetra el pedido?

	procedimiento actualizarAux() es
		j:= mov.nroJugador
		//tomo las ventas por las concretadas solamente
		si ( mov.cantidad > aux.stock ) entonces 
			B[j]:= B[j] + aux.stock
			aux.stock:= 0
			Esc("Error, falta de stock")
			productosFaltantes:= productosFaltantes + (mov.cantidad - aux.stock)	
		sino
			si ( mov.cantidad = aux.stock ) entonces
				B[j]:= B[j] + aux.stock
				aux.stock:= 0
			sino
				si (mov.cantidad < aux.stock) entonces
					B[j]:= B[j] + mov.cantidad
					aux.stock:= aux.stock - mov.cantidad 
				fs
			fs
		fs
		LeerMov()
	fp

Algoritmo

	abrirE/(archMae) ; abrirE/(archMov) ; abrir/S(archAct)
	productosFaltantes:= 0

	mientras (mov.clave <> HV) O (mae.clave <> HV) entonces
		si ( mae.clave < mov.clave) entonces
			//maestro a salida
			act := mae
			LeerMae() ; Grabar(achAct,act)
		sino
			si ( mae.clave = mov.clave ) entonces
				Mientras ( mae.clave = mov.clave ) hacer
					actualizarAux(); LeerMov()
				fm

				si ( aux.stock <> 0 ) entonces
					act:= aux ; Grabar(archAct,act) 
				fs

				LeerMae()
			sino
				si ( mae.clave > mov.clave ) entonces
					productosFaltantes:= productosFaltantes + mov.cantidad
					LeerMov()
				fs
			fs
		fs
	fm

	Para j = 1 hasta 26 hacer
		si (B[j] > masVendido) entonces
			masVendido:= aux.stock
			posicionMasVendido:= aux.nroJugador //yo supongo que quiere qe busque en el arreglo el nombre aunque en el archivo hay un dato que dice nombre jugador
		fs
	fp
fa