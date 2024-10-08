Acción ejListas(PRIM: puntero a nodo) es
Ambiente
	NODO = registro
		//zona de datos
		Legajo: N(8)
		Dni: N(8)
		Fecha: fecha //registro con formato de fecha
		//zona de punteros
		prox: puntero a NODO
Fr
	PRIM,P: puntero a NODO //se define igualmente el puntero externo
					//PRIM apunta al primero pero P puede apuntar al primero o a otros elementos
Algoritmo
	P:= PRIM // apunto p al primer elemento
	P:= *P.prox //si usas sin el asterisco, es como acceder a un registro, es un error,
	// acceder al contenido del nodo, puntero o dato
	//para acceder al nodo se ocupa el asteristo o la conjunción

	Esc(“Ingrese una fecha”)
	Leer(fechaIngresada)
	Mientras P <> nil hacer //este te encuentra todos los elementosa repetidos, hay otro que busca el primero y para
		Si *P.fecha = fechaIngresada entonces
			Esc("Se encontró el dato", *P.Legajo) //Para acceder al dato del registro se pone *P.datoquequierodelregistro
		fs
		P:= *P.prox
	fm
Fa

//este es la otra busqueda que encuentra el primero y para:
//BUSQUEDA LINEAL CON CENTINELA

LEER (valor)

p:= Prim
MIENTRAS (p <> nil) y (*p.dato <> valor) HACER
    p = *p.prox
FIN MIENTRAS

SI (p <> nil) ENTONCES
    Escribir ('Se encontro el elemento')
SINO
    Escribir ('No se encontro el elemento')
FIN SI


