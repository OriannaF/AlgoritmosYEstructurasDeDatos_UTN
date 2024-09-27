Accion EJ1 es
Ambiente
	clientes = registro
		nroCliente: N(3)
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

	novedades = registro
		clave = registro
			nroCliente: N(3)
			nroNovedad: N(3)
		fr
		apellidoNombre: AN(30)
		dni: N(8)
		idPaquete: 
	fr
Algoritmo
Fa