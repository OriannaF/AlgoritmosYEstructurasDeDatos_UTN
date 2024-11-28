// de Leonardo Brabo: https://github.com/LeonardoBrabo/Algoritmos_Resueltos/tree/master/10.Listas/Recestas%20Listas

Lista doblemente enlazada Recetas:

datos:

* cuando el problema dice dada una lista que se recibe... los punteros externos (prim y ult) se lo pasa por parametro "accion lista1 (prim,ult: puntero a nodo) es 
* si tenes que crear de cero una lista definimos en el ambiente los punteros externos (prim y ult)


Insertar elemento
------------------


ACCION insertar_doble ES
	
	AMBIENTE

	nodo= registro
		dato: entero
		prox: puntero a nodo
		ant: puntero a nodo
	fr

	p,q,prim,ult: puntero a nodo

	PROCESO
	prim:= nil ; ult:= nil
	nuevo(q)
	esc("ingrese el valor del elemento:")
	leer(*q.dato)
	si Prim = nil entonces			//unico elemento
		prim:= q
		ult:= q
		*q.prox:= nil
		*q.ant:= nil
	sino
		p:= prim
		mientas p<> nil y *q.dato > *p.dato hacer  // carga ordenada de forma ascendente, para descendente *q.dato < *p.dato
			p:= *p.prox
		fm
		si p = prim entonces  //primer elemento
			*q.prox:= p
			*p.ant:= q
			*q.ant:= nil
			prim:= q
		sino
			si P = nil entonces  //ultimo elemento
				ult.prox:= q
				*q.ant:= ult
				*q.prox:= nil
				ult:= q
			sino					//elemento medio
				*q.prox:= p
				*q.ant:= *p.ant
				*p.ant:= q		
				*(*q.ant).prox:= q  	//esto se lee como : al nodo que apunta *q.ant, de ese nodo que su puntero interno, prox, apunte a q.
			fs
		fs
	fs
fin_accion.




Eliminar Elemento
-----------------


Accion Eliminar_doble (prim,ult: puntero a nodo) es
	
	AMBIENTE

	nodo= registro
		dato: entero
		prox: puntero a nodo
		ant: puntero a nodo
	fr

	p: puntero a nodo
	valor: entero

	PROCESO
	si prim = nil entonces
		esc("lista vacia")
	sino
		esc("ingrese el valor del elemento a eliminar:")
		leer(valor)
		p:= prim
		mientras (p <> nil) y (*p.dato <> valor) hacer 
			p:= *p.prox
		fm
		si p= nil entonces
			esc("elemento no encontrado")
		sino
			si prim = ult entonces		//unico elemento
				prim:= nil
				ult:= nil
			sino
				si p = prim entonces  //primer elemento
					prim:= *p.prox
					*prim.ant:= nil
				sino
					si p = ult entonces		//ultimo elemento
						ult:= *p.ant
						*ult.prox:= nil
					sino							//elemento medio
						*(*p.ant).prox:= *p.prox
						*(*p.prox).ant:= *p.ant
					fs
				fs
			fs
			disponer(p)
		fs
	fs

Fin_accion.



