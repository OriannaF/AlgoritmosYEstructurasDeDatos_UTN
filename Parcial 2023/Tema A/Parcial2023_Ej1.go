Accion Ej1 es
Ambiente
	fecha = registro
		dia: N(2)
		mes: N(2)
		ano: N(4)
	fr

	hora = registro
		hora: N(2)
		minutos: N(2)
	fr

	Bicicletas = registro
		clave = registro
			nroSerie: N(5)
			modelo: AN(15)
		fr
		fechaAdquisicion: fecha
		fechaUltMantenimiento: fecha
		disponibilidad: booleano
	fr

	archMae: archivo de Bicicletas ordenado por clave
	mae: Bicicletas

	novedades = registro
		clave = registro
			nroSerie: N(5)
			modelo: AN(15)
			tipoNovedad: {1,2,3}
			fechaNovedad: fecha
		fr
		horaInicio: hora
		horaFin: hora
		circuitoNro: {1,2,3,4,5,6}
		idUsuario: N(8)
	fr

	archMov: archivo de novedades ordenado por clave
	mov: novedades

	salida = registro
		clave = registro
			nroSerie: N(5)
			modelo: AN(15)
		fr
		fechaAdquisicion: fecha
		fechaUltMantenimiento: fecha
		disponibilidad: booleano
	fr

	archSalida: archivo de salida ordenado por clave
	sal: salida

	usuarios = registro
		idUsuario: 	novedades(8)
		DNI: N(8)
		sexo: "M" , "F"
		apellidoNombre: AN(40)
		domicilio: AN(40)
		localidad: AN(40)
		provincia: AN(40)
		edad: AN(40)
	fr

	archIndexado: archivo de usuarios indexado por idUsuario
	ind: usuarios

	Procedimiento LeerMae() es
		Leer(archMae,mae)
		Si FDA(archMae) entonces
			mae.clave:= HV
		fs
	fp

	
	Procedimiento LeerMov() es
		Leer(archMov,mov)
		Si FDA(archMov) entonces
			mov.clave:= HV
		fs
	fp

	Funcion genero(x: entero):entero es
		ind.idUsuario:= x
		Leer(archIndexado,ind)
		Si EXISTE entonces
			Segun ind.sexo hacer
				="F": genero:= 1
				="M": genero:= 0
			fs
		sino
			Esc("Error préstamo, usuario no encontrado")
		fs
	fp

	Procedimiento modificarAuxiliar() es
		Si mov.clave.tipoNovedad = 1 entonces
			//alta
			Esc("Erro, alta inválida, esta bici ya se encuentra en el archivo")
		sino
			Si mov.clave.tipoNovedad = 2 entonces
				//prestamo
				genero(mov.clave.idUsuario)
				Segun genero hacer
					=1:= prestamosFemeninos:= prestamosFemeninos +1
					=0:= prestamosMasculino:= prestamosMasculino +1
				fs
			sino
				Si mov.clave.tipoNovedad = 3 entonces
					//mantenimiento
					sal.disponibilidad:= false
					sal.fechaUltMantenimiento:= mov.clave.fechaNovedad

					sal.clave.nroSerie:= mov.clave.nroSerie
					sal.clave.modelo:= mov.clave.modelo

					// esto creo que no va, con que le actualice esto cuando se hace el alta ya estaría creo.
					// sal.clave.fechaAdquisicion:= mov.clave.fechaNovedad
					
				Fs
			fs
		fs
	fp

		prestamosFemeninos , prestamosMasculino, genero, x : entero


Algoritmo

	AbrirE/(archMae) ; AbrirE/(archMov) ; AbrirE/(archIndexado) ; Abrir/S(archSalida)

	LeerMae() ; LeerMov()

	prestamosMasculino:= 0
	prestamosFemeninos:= 0

	Mientras mae.clave <> HV o mov.clave <> HV hacer
		Si mae.clave<mov.clave entonces
			//sal:= mae
			sal:=mae
			LeerMae()
		sino
			si mae.clave=mov.clave entonces
				//mod o baja
				//mod desde auxiliar
				//sal:= mov
				sal:=mae
				Mientras sal.clave.nroSerie = mov.clave.nroSerie hacer
					modificarAuxiliar()
					LeerMov()
				fm
				LeerMae()
			sino
				si mae.clave>mov.clave entonces
					//alta
					//sal:= mov
					//mod desde auxiliar
					//sal:= mov
					sal.clave.nroSerie:= mov.clave.nroSerie
					sal.clave.modelo:= mov.clave.modelo
					sal.clave.fechaAdquisicion:= mov.clave.fechaNovedad
					sal.fechaUltMantenimiento:= 0
					sal.disponibilidad:= true

					Mientras sal.clave.nroSerie = mov.clave.nroSerie hacer
						modificarAuxiliar()
						LeerMov()
					fm
				fs
			fs
		fs
	fm

	Esc("El total de prestamos para usuarios femeninos es de", prestamosFemeninos, "el total para usuarios masculinos es de", prestamosMasculinos, ".")

	Cerrar(archMae) ; Cerrar(archMov) ; Cerrar(archIndexado) ; Cerrar(archSalida)
Fa



