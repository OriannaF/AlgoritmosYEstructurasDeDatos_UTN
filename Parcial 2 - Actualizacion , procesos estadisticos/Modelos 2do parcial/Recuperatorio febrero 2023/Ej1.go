Accion Ej1(A: arreglo de alfanumerico de [1...26]) es
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
	i,j,categoria,categoriaMenosVendida,totalVenta: entero
	nombreCategoria, nombreCategoriaMenosVendida: AN(70)

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


	procedimiento actualizarAux() es
		j:= mov.nroJugador
		categoria:= categoria + 1 
		nombreCategoria:= mov.categoria
		si ( mov.cantidad > aux.stock ) entonces 
			B[j]:= B[j] + aux.stock
			totalVenta:= totalVenta * (((precio * aux.porcentajeDescuento) - precio) * aux.stock) 
			aux.stock:= 0
			Esc("Error, falta de stock")
			productosFaltantes:= productosFaltantes + (mov.cantidad - aux.stock)	
		sino
			si ( mov.cantidad = aux.stock ) entonces
				//calcular venta
				// no se de donde sacar el precio del producto asi q voy a usar la variable precio y finjir demencia :)
				totalVenta:= totalVenta + (((precio * aux.porcentajeDescuento) - precio) * mov.cantidad)
				B[j]:= B[j] + aux.stock
				aux.stock:= 0
			sino
				si (mov.cantidad < aux.stock) entonces
					totalVenta:= totalVenta + (((precio * aux.porcentajeDescuento) - precio) * mov.cantidad)
					B[j]:= B[j] + mov.cantidad
					aux.stock:= aux.stock - mov.cantidad 
				fs
			fs
		fs
		LeerMov()
	fp

Algoritmo

	abrirE/(archMae) ; abrirE/(archMov) ; abrir/S(archAct)
	productosFaltantes:= 0 ; totalVenta:= 0 ; categoria:= 0
	categoriaMenosVendida:= 0

	LeerMae() ; LeerMov()

	mientras (mov.clave <> HV) O (mae.clave <> HV) hacer
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

				si ( categoria < categoriaMenosVendida ) entonces
					categoriaMenosVendida:= categoria
					nombreCategoriaMenosVendida:= nombreCategoria 
				fs

				LeerMae()
			sino
				si ( mae.clave > mov.clave ) entonces //esta condicion está de as pero la pongo para guiarme, si el maestro es mayor no entra en ninguna de las anteriores y esta aca
					//como no debe estar en el archivo si no existe, no hice nada aca para actualizar
					productosFaltantes:= productosFaltantes + mov.cantidad
					LeerMov()
				fs
			fs
		fs
	fm

	Para j = 1 hasta 26 hacer
		si (B[j] > masVendido) entonces
			masVendido:= B[j]
			posicionMasVendido:= j //yo supongo que quiere que busque en el arreglo el nombre aunque en el archivo hay un dato que dice nombre jugador asi q no se A
		fs
	fp

	i:= posicionMasVendido
	//act1
	Esc("El jugador con más camisetas vendidas es ", A[i])
	//act2
	Esc("La cantidad de productos con faltante de stock es de ", productosFaltantes)
	//act3
	Esc("La categoria menos vendida es ", nombreCategoriaMenosVendida)

	//total
	Esc("El importe total de ventas concretadas es de: ", totalVenta)

fa

// correcciones 
// Una aclaración nomas, en el campo esPersonalizado, podes tomarlo como un booleano, pero de la forma en la que está tampoco estaría mal.
//Arreglo de los nombres lo recibís como parámetro.: Accion_nmb(Noms:Arreglo...)
//Lo del punto 2 sería que si todavía tenes stock, pero la cantidad que te piden es mayor, vendes lo que tenes y el restante lo contas. Ponele, si tengo en stock 2 camisetas, y me piden 3, cuento 1 en ese punto. Tenes que hacer la diferencia, justo lo que no querías jaja.
//Al comienzo te falto leer los archivos.
//No haría falta considerar casos separados para mov.cantidad = aux.stock, ó mov.cantidad < aux.stock, podes juntarlos en uno solo.
//Esta condición está demás. Si las claves no son iguales, y tampoco la clave del maestro no es menor que el de movimiento, no hay otra chance jaja
