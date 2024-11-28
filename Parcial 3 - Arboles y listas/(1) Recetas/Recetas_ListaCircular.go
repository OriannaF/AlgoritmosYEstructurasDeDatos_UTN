// de Leonardo Brabo: https://github.com/LeonardoBrabo/Algoritmos_Resueltos/tree/master/10.Listas/Recestas%20Listas

Lista circular simplemente enlazada Recetas:

datos:

* no existe el puntero nulo(nil)
* el ultimo nodo apunta al primero, y se accede a este ultimo a traves del puntero externo prim

Insertar elemento
------------------


ACCION insertar_Circular ES
	
	AMBIENTE

	nodo= registro
		dato: entero
		prox: puntero a nodo	
	fr

	p,q,ant,k,prim: puntero a nodo

	PROCESO
	prim:= nil ; 
	nuevo(q)
	esc("ingrese el valor del elemento:")
	leer(*q.dato)
	si Prim = nil entonces			//unico elemento
		prim:= q
		*q.prox:= q		
	sino
		p:= prim
		mientas (*p.prox <> prim) y (*q.dato > *p.dato hacer)  // carga ordenada de forma ascendente, para descendente *q.dato < *p.dato
			ant:= p
			p:= *p.prox
		fm
		si p = prim entonces  //primer elemento
			k:= prim 			
			mientras *k.prox<> prim hacer			//k se utiliza para llegar hasta el ultimo elemento, y luego conectarlo con el nuevo primer elemento.
				k:= *k.prox
			fm
			*k.prox:= q
			*q.prox:= prim
			prim:= q		
		sino
			si *p.prox = prim entonces  //ultimo elemento

				*p.prox:= q
				*q.prox:= prim
			sino					//elemento medio
				ant.prox:= q
				*q.prox:= p	 	
			fs
		fs
	fs
fin_accion.




Eliminar Elemento
-----------------


Accion Eliminar_circular (prim: puntero a nodo) es
	
	AMBIENTE

	nodo= registro
		dato: entero
		prox: puntero a nodo	
	fr

	p,ant,k: puntero a nodo
	valor: entero

	PROCESO
	si prim = nil entonces
		esc("lista vacia")
	sino
		esc("ingrese elemento a eliminar")
		leer(valor)
		p:= prim
		mientras *p.prox<> prim  y *p.dato <> valor hacer
			ant:= p
			p:= *p.prox
		fm
		si (*p.dato<> valor) entonces
			esc("elemento no encontrado")
		sino
			si p = *p.prox entonces 				// unico elemento
				prim:= nil
			sino
				si p = prim entonces				//primer elemento
					k:= prim
					mientras *k.prox<>prim hacer
						k:= *k.prox
					fm
					*k.prox:= *p.prox
					prim:= *p.prox
				sino
					*ant.prox:= *p.prox			//elemento medio o ultimo elemento
				fs
			fs
			disponer(p)
		fs
	fs
Fin_accion.




