// de Leonardo Brabo: https://github.com/LeonardoBrabo/Algoritmos_Resueltos/tree/master/10.Listas/Recestas%20Listas

Lista Simplemente enlazada Recetas:

datos:

* cuando el problema dice dada una lista que se recibe... el puntero externo (prim) se lo pasa por parametro "accion lista1 (prim: puntero a nodo) es "
* si tenes que crear de cero una lista definimos en el ambiente el puntero externo (prim)


Insertar elemento
------------------


ACCION insertar_simple ES
	
	AMBIENTE

	nodo= registro
		dato: entero
		prox: puntero a nodo
	fr

	p,ant,q, prim: puntero a nodo

	PROCESO

	prim:= nil		//solo si se lo crea de cero
	nuevo(q)
	esc("ingrese valor del elemento");
	leer(*q.dato)
	si prim= nil entonces		//unico elemento
		prim:= q
		*q.prox:= nil
	sino
		p:= prim
		mientras (p <> nil) y (*q.dato > *p.dato) hacer  	//carga ordenada ascendente, para que sea descendente: *q.dato < *p.dato
			ant:= p
			p:= *p.prox
		fm
		si p = prim entonces		//primer elemento
			*q.prox:=p
			prim:= q
		sino
			*ant.prox:= q 			//elemento medio o ultimo elemento
			*q.prox:= p
		fs
	fs
Fin_accion.


Eliminar Elemento
-----------------


Accion Eliminar_simple (prim: puntero a nodo) es
	
	AMBIENTE

	nodo= registro
		dato: entero
		prox: puntero a nodo
	fr

	p,ant,: puntero a nodo
	valor: entero

	PROCESO

	si prim = nil entonces
		esc("lista vacia")
	sino
		esc("ingrese el valor que desea eliminar:")
		leer(valor)
		p:= prim
		mientras (p<> nil) y (*p.dato<>valor)  //sale por terminar la lista o por encontrar el valor
			ant:= p
			p:= *p.prox
		fm
		si p = nil entonces						//si se termina la lista
			esc("el elemento no encontrado")
		sino
			si p = prim entonces
				prim:= *p.prox    //primer elemento o unico elemento
			sino
				*ant.prox:= *p.prox		//elemento medio o ultimo elemento
			fs
			disponer(p)		
		fs
	fs
Fin_accion.



