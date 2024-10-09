Accion listas_simples (prim: puntero a nodo) es 
 Ambiente 
 nodo = registro 
	 dato: entero //el tipo de dato puede ser entero, alfanumerico o incluso otro registro 
	 prox: puntero a nodo
	FinReg
 p,q,ant: puntero a nodo 
 //aclaracion: cuando se menciona al "proximo", hablo del puntero interno "prox" el cual está dentro de los nodos
 //aclaracion: p,q,k,ant,prim o todos los punteros los cuales se definan FUERA DEL NODO van a ser punteros externos
 //aclaracion: ant(solo en listas doblemente enlazadas) y prox, como son definidos DENTRO DEL NODO son punteros internos
 //aclaracion: solo en caso de que tengas que crear nuevas listas (prim1,prim2,...,primn) cada prim va a tener que apuntar a nil al comenzar el proceso 
 Proceso 
 	 Procedimiento carga_apilada() es 
		 Nuevo(q)
		 *q.dato:= valor //le asignas un valor ingresado por el usuario o un valor
	 		Si prim = nil entonces 
			 *q.prox:= nil   //el prox de "q" va a apuntar a nil en caso de que la lista este vacia
			Sino 
			 *q.prox:= prim  //el prox de "q" va a apuntar a donde este apuntando prim en caso de que la lista ya tenga elementos
			FinSi
		 prim:= q  //cuando se salga del condicional prim va a apuntar a "q"
		FP
	 Procedimiento carga_apilada2() es 
	 	 Nuevo(q)
		 *q.dato:= valor  
		 *q.prox:= prim //en caso de que prim este apuntando a nil entonces el proximo de "q" va a apuntar a nil, sino va a apuntar donde sea que apunte prim
		 prim:= q //prim va a apuntar a "q"
		FP
	 Procedimiento carga_encolada() es  
	 	 Nuevo(q)
		 *q.dato:= valor 
		 	Si prim = nil entonces  //primero preguntamos si es que la lista está vacia
			 *q.prox:= nil 
			 prim:= q  //si esta vacia entonces prim va a apuntar a "q" y el proximo de "q" va a apuntar a nil
			Sino 
			 p:= prim   //en caso de no estar vacia vamos a hacer que "p" apunte a donde este apuntando prim (el primer nodo de la lista)
			 Mientras (p<>nil) hacer  
			 	 ant:= p 
				 p:= *p.prox  //con la condicion del mientras y estas acciones vamos a recorrer la lista hasta que "p" apunte a nil y "ant" apunte a un nodo antes de nil
				FM 
			 *ant.prox:= q 
			 *q.prox:= p  //aca insertamos "q" haciendo que el proximo de "ant" apunte a "q" y el proximo de "q" apunte a "p"
			FinSi
		FP
	 Procedimiento insertar_elemento() es 
	 	 Nuevo(q)
		 *q.dato:= valor 
		 	Si prim = nil entonces  //si la lista esta vacia entonces prim apunta a "q" y el proximo de "q" apunta a nil
			 *q.prox:= nil 
			 prim:= q 
			Sino 
			 p:= prim; ant:= nil //hacemos que "p" apunte a prim para que al momento de poner el ciclo mientras "p" recorra desde ese primer elemento hasta nil 
			 Mientras (p<>nil) hacer 
			 	 ant:= p 
				 p:= *p.prox 
				FM
				Si (p = prim) entonces
				 *q.prox:= p 
				 prim:= q 
				Sino   //esto abarca las condiciones de (*p.dato > *q.dato) y la de (p = nil), por lo tanto se realizan las mismas acciones en ambas
				 *ant.prox:= q  
				 *q.prox:= p //para la condicion (p = nil) se puede poner *q.prox:= nil, pero como "p" ya está apuntando a nil entonces solamente se pone a "p"
				FinSi 
			FinSi
		FP	 
	 Procedimiento eliminar_elemento() es 
	 	 	Si prim = nil entonces 
			 Esc("No se puede borrar nada, lista vacia")
			Sino 
			 p:= prim; ant:= nil 
			 Mientras (p<>nil) hacer  //recorremos la lista
			 	 ant:= p 
				 p:= *p.prox 
				FM 
				Si (prim = p) entonces   //si prim y "p" estan apuntando al mismo nodo y queremos borrar dicho nodo, solo apuntamos prim a nil
				 prim:= nil 
				 //disponer(p)
				Sino 
					Si (*p.prox = nil) entonces  //eliminar el ultimo elemento: solo hacemos que el proximo de "ant" apunte al proximo nodo al que "p" está apuntando  
					 *ant.prox:= *p.prox 
					 //disponer(p)
					Sino 
						Si (*p.dato = valor) entonces  //en caso de que el usuario ingrese un valor a eliminar se hace la misma accion que en la condicion anterior
						 *ant.prox:= *p.prox 
						 //disponer(p)
						FinSi
					FinSi
				FinSi
			 disponer(p) //el disponer(p) se puede poner aca o dentro de cada condicional 
			FinSi
		FP
	 Procedimiento ingresar_a_un_nodo() es 
	 	 p:= prim 
		 Mientras (p<>nil) y (*p.dato <> datoingresado) hacer //busqueda lineal con centinela 
		 	 p:= *p.prox 
			FM
		 	Si (*p.dato <> datoingresado) entonces 
			 Esc("No se encontro el elemento")
			Sino 
			 Esc(*p.dato)
			FinSi
		FP
	FinProceso
FinAccion

	

