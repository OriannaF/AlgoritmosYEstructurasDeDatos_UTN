//pueden los campos de distintos archivos llamarse igual?
//se puede asignar el registro de salida de la manera que puse?, tipo de una asignar el registro al registro del otro? o es siempre campo por campo?
Accion EJ1 es
Ambiente
	clientes = registro //maestro
		clave = registro
			nroCliente: N(3) //ordenado por
		fr
		apellidoNombre: AN(30)
		dni: N(8)
		idPaquete:N(10)
		saldo: N(10)
		estado: {"SALDADO","SALDO A FAVOR","DEUDOR"}
		categoria:{"SIMPLE","PLATA","ORO","DIAMANTE"}
		puntos: N(10)
		fechaBaja = registro
			dia: N(2)
			mes: N(2)
			ano: N(4)
		fr
	fr

	archClientes: archivo de clientes ordenado por nroCliente
	cl: clientes

	novedades = registro //movimiento
		clave = registro
			nroCliente: N(3)
			nroNovedad: N(3)
		fr
		apellidoNombre: AN(30)
		dni: N(8)
		idPaquete: N(10)
		fechaNovedad = registro
			dia: N(2)
			mes: N(2)
			ano: N(4)
		fr
		monto: N(10)
	fr

	archNovedades: archivo de novedades ordenado por clave
	nv: novedades

	paquetesTuristicos = registro
		idPaquete: N(10)
		nombre: AN(10)
		costo: N(10)
		destino: AN(30)
	fr

	archPaquetes: archivo de paquetesTuristicos indexado por idPaquete
	pq: paquetesTuristicos

	Procedimiento LeerMae() es
		Leer(archClientes,cl)
		Si FDA(archClientes) entonces
			cl.nroCliente := HV
		Fs
	fp

	Procedimiento LeerMov() es
		Leer(archNovedades,nv)
		Si FDA(archNovedades) entonces
			nv.clave := HV
		Fs
	fp

	actClientes = registro
		nroCliente: N(3) //ordenado por
		apellidoNombre: AN(30)
		dni: N(8)
		idPaquete:N(10)
		saldo: N(10)
		estado: {"SALDADO","SALDO A FAVOR","DEUDOR"}
		categoria:{"SIMPLE","PLATA","ORO","DIAMANTE"}
		puntos: N(10)
		fechaBaja = registro
			dia: N(2)
			mes: N(2)
			ano: N(4)
		fr		
	fr

	archActClientes: archivo de actClientes ordenado por nroCliente
	actCl: actClientes

	Procedimiento Porcentaje() es
		Segun cl.categoria hacer
			"SIMPLE": simple:= simple + 1
			"PLATA": plata:= plata + 1
			"ORO": oro:= oro + 1
			"DIAMANTE": diamante:= diamante + 1
		fs
	fp

	Procedimiento Modificar() es
		//AGREGAR mientras nv = actCl guardo modificaciones, PUNTOS PAGOS SE ACTUALIZA
		//Leo movimiento dentro de ese ciclo tmb
		Mientras actCl.clave.nroCliente = nv.clave.nroCliente hacer
			
			Si nv.nroNovedad = 0 entonces
				Porcentaje()
				actCl.estado:= "DEUDOR"
				actCl.saldo:= saldo
				actCl.puntos:= 0
				actCl.nroCliente:= nv.nroCliente
				actCl.apellidoNombre:= nv.apellidoNombre
				actCl.dni:= nv.dni
				actCl.idPaquete:= nv.idPaquete
				
			Sino
				Si nv.nroNovedad = 999 entonces //baja
					//actualizo salida
					actCl.nroCliente:= nv.nroCliente
					actCl.apellidoNombre:= nv.apellidoNombre
					actCl.dni:= nv.dni
					actCl.idPaquete:= nv.idPaquete
					actCl.estado:= "SALDO A FAVOR"
					actCl.saldo:= saldo
					actCl.fechaBaja:= cl.fechaNovedad
					actCl.puntos:= cl.puntos
					//act A
					cantidadBaja:= cantidadBaja + 1
					montoReintegrar:= montoReintegrar + nv.monto
				Sino
					Mientras actCl.clave = cl. hacer
						//entre 1...998
						Porcentaje()
						Segun cl.categoria hacer
							"PLATA": saldo:= cl.saldo - ((nv.nroNovedad * 0,10) +nv.nroNovedad)
							actClientes.puntos:= 10
							"ORO": saldo:= cl.saldo - ((nv.nroNovedad * 0,15) + nv.nroNovedad)
							actClientes.puntos:= 20
							"DIAMANTE": saldo:= cl.saldo - ((nv.nroNovedad * 0,20) + nv.nroNovedad)
							actClientes.puntos:= 30
						Fs

						Si ( saldo = 0 ) entonces
							actCl.saldo:= 0
							actCl.estado:= "SALDADO"
						Sino

							Si saldo < 0 entonces
								actCl.saldo:= 0
								actCl.estado:= "A FAVOR"
								montoReintegrar:= montoReintegrar + saldo
							Sino
							// > 0
								actCl.saldo:= SALDO
								actCl.saldo:= "DEUDOR"
							Fs
						Fs

						actCl.nroCliente:= nv.nroCliente
						actCl.apellidoNombre:= nv.apellidoNombre
						actCl.dni:= nv.dni
						actCl.idPaquete:= nv.idPaquete
					Fm
				Fs
			Fs
			LeerMov()
		fm
	fp

	

Algoritmo

	AbrirE/(clientes) ; AbrirE/(novedades) ; AbrirE/(paquetesTuristicos) ; Abrir/S(actClientes)
	LeerMae() ; LeerMov()

	//majo dijo que hay que comparar clave no importa si tiene mas de una clave, no son tan especificos con eso
	Mientras (nv.clave <> HV) o (cl.clave <> HV) hacer
		Si cl.nroCliente < nv.clave.nroCliente entonces
			// sal := mae
			//campo por campo, salida igual a maestro
			Porcentaje()
			actCl:= cl
			LeerMae()
		Sino
			Si cl.nroCliente = nv.clave.nroCliente
				Mientras cl.nroCliente=nv.clave.nroCliente hacer //este no se si va
					Si nv.nroNovedad = 0 entonces
						Porcentaje()
						actCl.estado:= "DEUDOR"
						actCl.saldo:= cl.saldo
						actCl.puntos:= 0
						actCl.nroCliente:= nv.nroCliente
						actCl.apellidoNombre:= nv.apellidoNombre
						actCl.dni:= nv.dni
						actCl.idPaquete:= nv.idPaquete
						
					Sino
						Si nv.nroNovedad = 999 entonces //baja
							//actualizo salida
							actCl.nroCliente:= nv.nroCliente
							actCl.apellidoNombre:= nv.apellidoNombre
							actCl.dni:= nv.dni
							actCl.idPaquete:= nv.idPaquete
							actCl.estado:= "SALDO A FAVOR"
							actCl.saldo:= cl.saldo
							actCl.fechaBaja:= cl.fechaNovedad
							actCl.puntos:= cl.puntos
							//act A
							cantidadBaja:= cantidadBaja + 1
							montoReintegrar:= montoReintegrar + nv.monto
						Sino
							//entre 998 y 0
							actCl.saldo:= cl.saldo
							Modificar()
						Fs
				Fm
				LeerMov() ; LeerMae()
			Sino
				// mae>mov
				// alta a nuevo cliente
				// sal := mov
				// salida igual al movimiento, salida = novedades campo por campo.

				simple:= simple + 1 //los nuevos clientes son categoria simple
				actClientes:= nv

				
				LeerMov()
				Modificar()
				
			fs
		fs

	Fm
	CerrarArchivos()

	SacarPorcentajesConElProc() ; MostrarPorcentajeClientesDeBajaYMontoTotal()

Fa