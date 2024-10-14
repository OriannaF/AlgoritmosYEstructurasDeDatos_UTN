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

	archMae: archivo de clientes ordenado por nroCliente
	mae: clientes

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

	archMov: archivo de novedades ordenado por clave
	mov: novedades

	paquetesTuristicos = registro
		idPaquete: N(10)
		nombre: AN(10)
		costo: N(10)
		destino: AN(30)
	fr

	archPaquetes: archivo de paquetesTuristicos indexado por idPaquete
	Paq: paquetesTuristicos

	Procedimiento LeerMae() es
		Leer(mae,regMae)
		Si FDA(mae) entonces
			regMae.clave:= HV
		Fs
	fp

	Procedimiento LeerMov() es
		Leer(mov,regMov)
		Si FDA(mov) entonces
			regMov.clave.idPaquete:= HV
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

	archAux: archivo de actClientes ordenado por nroCliente
	aux: actClientes

	Funcion Porcentaje(x:alfanumerico):entero es
		Segun x hacer
			="SIMPLE": simple:= simple + 1
			="PLATA": plata:= plata + 1
			="ORO": oro:= oro + 1
			="DIAMANTE": diamante:= diamante + 1
		fs
	fp

	cantidadBaja, montoReintegrar,x,simple,oro,plata,diamante,total: entero
	
	Funcion calculoPagos(x: entero):entero es //x es mov.monto
		Segun mae.categoria hacer
			="SIMPLE": pago := mae.saldo - x  
			="PLATA": pago := mae.saldo - (x + (x * 0.10))
					  aux.puntos := mae.punto + 10
			="ORO": pago := mae.saldo - (x + (x * 0.15))
					aux.puntos := mae.punto + 20
			="DIAMANTE": pago := mae.saldo - (x + (x * 0.20))
					aux.puntos := mae.punto + 30
		Fs
		Segun pago hacer
			=0: aux.estado:= "SALDADO"
			>0: aux.estado:= "DEUDOR"
			<0: aux.estado:= "A FAVOR"

		Fs
	Fp
Algoritmo

	AbrirE/(archMae) ; AbrirE/(archMov) ; AbrirE/(archPaquetes)
	Abrir/S(archAux) ; LeerMae() ; LeerMov()
	cantidadBaja:= 0 ; montoReintegrar:= 0
	simple:= 0 ; plata:= 0 ; oro:= 0 ; diamante:= 0
	
	Mientras (mae.clave <> HV) O (mov.clave.nroCliente <> HV) hacer
		Si mae.clave < mov.clave.nroCliente entonces
			//auxiliar := maestro
			aux:= mae
			Grabar(archAux,aux)
			Porcentaje(mae.categoria)
			LeerMae()
			//pongo lo mismo que el maestro pq no hay movimiento para este id
		Sino
			Si mae.clave = mov.clave.nroCliente entonces
				//alta o baja, hay cambios para el cliente
				Porcentaje(mae.categoria)
				Si mov.nroNovedad = 0 entonces //ALTA
					//aux:= mov
					Esc("Error, alta invalida, el cliente ya existe")
				Sino
					Si mov.nroNovedad = 999 //BAJA
						aux.nroCliente:= mov.nroCliente
						aux.apellidoNombre:= mov.apellidoNombre
						aux.dni:=mov.dni
						aux.idPaquete:=mov.idPaquete
						aux.saldo:= mae.saldo
						aux.estado:= "SALDO A FAVOR"
						aux.categoria:= mae.categoria
						aux.puntos:= 0
						aux.fechaBaja:= mov.fechaNovedad
						cantidadBaja:= cantidadBaja + 1
						montoReintegrar := montoReintegrar + mae.saldo
					Sino
				 		//entre 1...998 PAGOS
						Mientras mov.nroNovedad >0 Y <999 hacer
							calculoPagos(mov.monto) //pago y estado se ven aca
							aux.saldo:= pago
							aux.nroCliente:= mov.nroCliente
							aux.apellidoNombre:= mov.apellidoNombre
							aux.dni:=mov.dni
							aux.idPaquete:=mov.idPaquete
							aux.categoria:= mae.categoria
							aux.puntos:= 0
							
							Si aux.estado = "A FAVOR" entonces // para reintegrar
								montoReintegrar := montoReintegrar + aux.saldo
							Fs
	

						Fm
					Fs
				Fs
				LeerMae() ; LeerMov()
			Sino
				Si mae.clave > mov.clave.nroCliente entonces
					//alta
					Si mov.clave.nroNovedad = 0 entonces
						//busco costo en archivo indexado
						paq.idPaquete:= mov.idPaquete
						Leer(archPaquetes,paq)
						Si EXISTE entonces //aux:=mov
							aux.saldo:= paq.costo
							actCl.estado:= "DEUDOR"
							actCl.puntos:= 0
							actCl.nroCliente:= mov.nroCliente
							actCl.apellidoNombre:= mov.apellidoNombre
							actCl.dni:= mov.dni
							actCl.idPaquete:= mov.idPaquete
							simple:= simple + 1

						Sino // no guardo el alta ya que no encuentra el costo
							Esc("Error, ID Paquete no encontrado")
						Fs
					Fs
					LeerMov()
				Fs
			Fs
		Fs
	Fm

	//informo
	Esc("La cantidad de clientes que se dieron de baja es de:", cantidadBaja, " clientes. El monto a reintegrar, tanto de saldos a favor, como de saldos de clientes que se dieron de baja es de: $", montoReintegrar, ".")

	total:= simple + plata + oro + diamante

	Esc("Porcentajes de categorias de clientes:")
	Esc("-------------------------------------")
	Esc("SIMPLE:",(100* (simple / total)),"%")
	Esc("-------------------------------------")
	Esc("ORO:",(100* (oro / total)),"%")
	Esc("-------------------------------------")
	Esc("PLATA:",(100* (plata / total)),"%")
	Esc("-------------------------------------")
	Esc("DIAMANTE:",(100* (diamante / total)),"%")

