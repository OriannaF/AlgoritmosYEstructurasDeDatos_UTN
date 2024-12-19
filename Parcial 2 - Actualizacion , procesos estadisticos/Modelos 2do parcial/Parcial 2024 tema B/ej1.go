accion ej1 es
ambiente
	clientesBaja: entero

	fecha = registro
		dia: N(2)
		mes: N(2)
		anio: N(2)
	fr

	clientes = registro
		clave = registro
			idSucursal: N(10)
			idCliente: N(10)
		fr
		nombreApellido: AN(60)
		saldo: N(2)
		fechaAlta: fecha
		fechaBaja: fecha
	fr
	archMae, archSal: archivo de clientes ordenado por clave
	mae, sal, aux: clientes

	movimientos = registro
		clave = registro
			idSucursal: N(10)
			idCliente: N(10)
			codMovimiento: (0...99) //0= alta 99= baja otro= transaccion
		fr
		nombreApellido: AN(60)
		fechaMovimiento: fecha
		monto: N(10)
		detalle: AN(60) //descripcion
		categoria: (1...6) 
		tipo: {"I","E"} //I= ingreso E= egreso
	fr
	archMov: archivo de movimientos ordenado por clave
	mov: movimientos

	procedimiento leerMae() es
		leer(archMae,mae)
		si FDA(archMae) entonces
			mae.clave:= HV
		fs
	fp

	procedimiento leerMov() es
		leer(archMov,mov)
		si FDA(archMov) entonces
			mov.clave:= HV
		fs
	fp

	procedimiento modificar() es
		si (mov.clave.codMovimiento = 0) entonces
			//alta
			esc("Alta inválida, cliente existente")
		sino
			si (mov.clave.codMovimiento = 99) entonces
				//baja
				aux.fechaBaja:= mov.fechaMovimiento
				si (aux.fechaBaja = mes) entonces
					clientesBaja:= clientesBaja + 1 
				fs 
			sino
				//1...98 modificion
				si (mov.tipo = "I") entonces //ingreso
					aux.saldo:= aux.saldo + mov.monto
				sino //egreso
					aux.saldo:= aux.saldo - mov.monto
				fs
			fs
		fs
	fp

algoritmo
	abrirE/(archMae) ; abrirE/(archMov) ; abrir/S(archSal) 
	leerMae() ; leerMov()

	clientesBaja:= 0

	esc("ingrese el mes actual") //para saber cual es el mes actual
	leer("mes")
	
	mientras (mae.clave <> HV) O (mov.clave <> HV) hacer
		si (mae.clave < mov.clave) entonces
			sal:= mae ; grabar(archSal,sal) ; leerMae()
		sino
			si (mae.clave = mov.clave) entonces
				aux:= mae
				mientras (aux.clave = mov.clave) hacer
					modificar() ; leerMov()
				fm
				sal:= aux ; grabar(archSal,sal) ; leerMae()
			sino
				//mae.clave > mov.clave
				si (mov.codMovimiento = 0) entonces
					//alta
					aux.idSucursal:= mov.idSucursal
					aux.idCliente:= mov.idCliente
					aux.nombreApellido:= mov.nombreApellido
					auc.saldo:= 0
					aux.fechaAlta:= mov.fechaMovimiento

					mientras (aux.clave = mov.clave) hacer
						modificar() ; leerMov()
					fm
					sal:= aux ; grabar(archSal,sal) ; leerMae()
				sino
					si (mov.codMovimiento = 99) entonces
						//baja
						esc("Baja inválida, cliente inexistente")
					sino
						//1...98 modificion
						esc("Moficación inválida, cliente inexistente")
					fs
				fs
			fs
		fs
	fm

	esc("clientes dados de baja: ", clientesBaja)

	cerrar(archMae) ; cerrar(archMov) ; cerrar(archSal)
fm