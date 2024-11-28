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
	//osea, Si todos los errores dan *p.estado = "R" entonces se pone en la lista: proy.clave.cod proy.nombre

	//La segunda contendrá el código y nombre de los proyectos que tengan al menos el 50% de sus errores en estado resueltos o en Proceso. (Ordenada por dicho porcentaje de mayor a menor)


Algoritmo
fa