Accion Ej1(PRIM:puntero a Nodo) es
Ambiente
	proyectos = registro
		clave = registro
			codigo: N(10)
		fr
		nombre: AN(60)
		tipo: ("D" , "W")
		cantErrores: N(3) //cantidad de veces que hay que avanzar en la lista
	fr

	archProy : archivo de proyectos ordenado por clave
	proy: proyectos

	listaEntrada = registro
		descripcion: AN(250)
		estado: ("N","O","R") //(N: No resuelto | O: En Proceso | R: Resuelto)
		prox: puntero a listaEntrada
	fr

	prim,p: puntero a listaEntrada

	//listas de salida (ambas dobles)

	//La primera contendrá el código y nombre de los proyectos que tengan todos sus errores en estado resueltos
	//tengo que sacar los errores, ir viendo en la lista con la cantidad de errores y verificar si estan resueltos

	//La segunda contendrá el código y nombre de los proyectos que tengan al menos el 50% de sus errores en estado resueltos o en Proceso. (Ordenada por dicho porcentaje de mayor a menor)


Algoritmo

	AbrirE/(archProy) ; Leer(archProy,proy)

	//inicializar
	contador:= 0 ; resuelto:= 0 ; noResuelto:= 0; proyectosNoResueltos:= 0

	p:= PRIM

	mientras NFDA(archProy) hacer

		erroresTotales:= proy.cantErrores
		contador:= proy.cantErrores

		mientras (contador <> 0) hacer
			si (*p.estado = "R") entonces
				resuelto:= resuelto + 1
			sino
				si (*p.estado = "N") entonces //última actividad (hay que informar)
					noResuelto:= noResuelto + 1
				sino //en proceso
					enProceso:= enProceso + 1
				fs
			fs
			p:= *p.prox
			contador:= contador - 1
		fm

		//primera lista
		si (erroresTotales = resuelto) entonces 
			//pongo en primer lista cod y nom del archivo
		fs

		//segunda lista
		si ((((resuelto + proceso)/erroresTotales)* 100) > 50 ) entonces
			//carga ordena descendente
		fs

		//última consigna (se informa dsp)
		si (erroresTotales = noResuelto) entonces
			proyectosNoResueltos:= proyectosNoResueltos + 1
		fs
		
	fm

	esc("la cantidad de proyectos no resueltos es de: ", proyectosNoResueltos)

	cerrar(archProy)
fa