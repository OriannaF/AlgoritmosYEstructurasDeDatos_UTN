accion ej2 es 
ambiente
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
	archMae: archivo de clientes ordenado por clave
	mae: clientes

	sucursales = registro
		idSucursal: N(10)
		nombreSucursal: AN(60)
		direccion: AN(60)
		localidad: AN(60)
	fr

	archIndexado: archivo de sucursales indezado por idCliente
	ind: sucursales

	A: arreglo de [1..16,1...4] de entero
	i,j: entero

	


algoritmo
	AbrirE/(archMae) ; AbrirE/(archIndexado)  
	leer(archMae,mae)

	//pongo a cero la matriz
	para (i:=1 hasta 16) hacer
		para (i:=1 hasta 4) hacer
			A[i,j]:= 0
		fp
	fp

	mientras NFDA(archMae) hacer

		si (fechaBaja <> 0) entonces //si baja distinta de vacio

			segun (mae.saldo) hacer
				<100000: j:= 3
				>100000 Y <1500000: j:= 2
				>1500000: j:= 1
			fs

			i:= mae.idSucursal

			A[i,j] := A[i,j] + 1
			A[16,j] := A[16,j] + 1
			A[i,4] := A[i,4] + 1
			A[16,4] := A[16,4] + 1
		fs

		leer(archMae,mae)
	fm

	esc("| sucursal | categoria diamante | categoria oro | categoria estandar | totales por sucursal |")
	
	para (i:=1 hasta 15) hacer
		ind.idSucursal:= i ; leer(archIndexado,ind)
		si EXISTE entonces
			esc("|",ind.nombreSucursal,"|", A[i,1], "|", A[i,2],"|", A[i,3],"|", A[i,4],)
		fs
	fp

	esc("| totales por categoria |", A[16,1], "|", A[16,2],"|", A[16,3],"|", A[16,4],"|" )

	cerrar(archMae) ; cerrar(archIndexado)
fa