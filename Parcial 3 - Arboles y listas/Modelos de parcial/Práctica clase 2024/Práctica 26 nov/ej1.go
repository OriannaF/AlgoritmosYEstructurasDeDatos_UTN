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

	listaUno = registro
		codigo: N(10)
		nombre: AN(60)
		prox: puntero a listaUno
		ant: puntero a listaUno
	fr

	//La segunda contendrá el código y nombre de los proyectos que tengan al menos el 50% de sus errores en estado resueltos o en Proceso. (Ordenada por dicho porcentaje de mayor a menor)

	listaDos = registro
		codigo: N(10)
		nombre: AN(60)
		prox: puntero a listaUno
		ant: puntero a listaUno
	fr

	primDos, qDos: puntero a


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
			nuevo(qUno);
			*qUno.nombre:= proy.nombre                                    // Crear un nuevo nodo `q`.
			*qUno.codigo:= proy.clave.codigo                               // Leer el valor ingresado por el usuario y almacenarlo en `q.dato`.
		
			si (prim = nil) entonces                            // Caso 1: Si la lista está vacía:
				prim := q                                //    - `q` se convierte en el primer nodo (prim).
				*qUno.prox := *qUno                             //    - El nodo apunta a sí mismo, formando un ciclo.
			sino                                         // Caso 2: Si la lista no está vacía:
				p := prim                                //    - Inicializar `p` en el primer nodo.
		
				// Recorrer la lista mientras no se vuelva al inicio (*p.prox <> prim) y
				// el dato del nodo actual (*p.dato) sea menor que el nuevo dato (*q.dato).
				
				meintras (*p.prox <> prim) y (*p.dato < *q.dato) hacer
					ant := p                             // Guardar el nodo actual como `ant` (anterior).
					p := *p.prox                         // Avanzar al siguiente nodo.
				fm
		
				// Caso 2a: Insertar antes del primer nodo (nuevo valor es el menor de todos).
				si (p = prim) entonces
					k := prim                            // Inicializar `k` para buscar el último nodo.
					MIENTRAS *k.prox <> prim HACER       // Recorrer la lista hasta encontrar el último nodo.
						k := *k.prox                     // Avanzar al siguiente nodo.
					FM
					*k.prox := q                         // El último nodo apunta al nuevo nodo `q`.
					*qUno.prox := p                         // El nuevo nodo `q` apunta al primer nodo actual (`p`).
					prim := q                            // Actualizar `prim` para que apunte al nuevo nodo.
		
				 // Caso 2b: Insertar después del último nodo.
				sino
					si (*p.prox = prim) entonces
						*p.prox := *qUno                        // El nodo actual `p` apunta al nuevo nodo `q`.
						*qUno.prox := prim                      // El nuevo nodo `q` apunta al primer nodo (`prim`).
		
					 // Caso 2c: Insertar entre dos nodos.
					sino
						*ant.prox := *qUno                      // El nodo anterior (`ant`) apunta al nuevo nodo `q`.
						*qUno.prox := p 
					fs                        // El nuevo nodo `q` apunta al nodo actual (`p`).
				fs
			fs
		fs

		//segunda lista
		si ((((resuelto + proceso)/erroresTotales)* 100) > 50 ) entonces
			//carga ordena descendente
			*qDos.nombre:= proy.nombre                                    // Crear un nuevo nodo `q`.
			*qDos.codigo:= proy.clave.codigo  
			mientas (p <> nil) Y (*q.dato > *p.dato) hacer  // carga ordenada de forma ascendente, para descendente *q.dato < *p.dato
				p:= *p.prox
			fm
			si (p = prim) entonces  //primer elemento
				*qDos.prox:= p
				*p.ant:= *qDos
				*qDos.ant:= nil
				prim:= *qDos
			sino
				si (p = nil) entonces  //ultimo elemento
					ult.prox:= *qDos
					*qDos.ant:= ult
					*qDos.prox:= nil
					ult:= *qDos
				sino					//elemento medio
					*qDos.prox:= p
					*qDos.ant:= *p.ant
					*p.ant:= *qDos		
					*(*qDos.ant).prox:= *qDos 	//esto se lee como : al nodo que apunta *q.ant, de ese nodo que su puntero interno, prox, apunte a q.
				fs
			fs
		fs

		//última consigna (se informa dsp)
		si (erroresTotales = noResuelto) entonces
			proyectosNoResueltos:= proyectosNoResueltos + 1
		fs
		
	fm

	esc("la cantidad de proyectos no resueltos es de: ", proyectosNoResueltos)

	cerrar(archProy)
fa