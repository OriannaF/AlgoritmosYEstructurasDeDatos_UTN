// Escribir un algoritmo que permita buscar, insertar o borrar un elemento identificado con una clave determinada en 
// una lista circular simplemente encadenada.

Accion ListaSimpleCircular_4.8 ES
	AMBIENTE

	Nodo = registro
		dato: entero
		prox: puntero a nodo
	FR

	prim,p,q,ant,k:puntero a nodo

	op: ("s","n")
	elec:1..3
	elemento: entero

	PROCEDIMIENTO buscar() es

		ESC("ingrese valor de elemento a buscar:")
		LEER(elemento)

		SI prim = nil ENT
			ESC("lista vacia")
		SINO
			p:= prim
			MIENTRAS (*p.prox <> prim) y (*p.dato <> elemento) HACER
				p:= *p.prox
			FM

			SI *p.prox = prim ENT
				ESC("Elemento no encontrado")
			SINO
				ESC("se encontro el elemento:",elemento)
			FS
		FS
	FP

		PROCEDIMIENTO insertar() es
		    ESC("Ingrese el valor que quiere insertar:")  // Solicitar al usuario el dato a insertar.
		    nuevo(q);                                    // Crear un nuevo nodo `q`.
		    LEER(*q.dato)                                // Leer el valor ingresado por el usuario y almacenarlo en `q.dato`.
		
		    SI prim = nil entonces                            // Caso 1: Si la lista está vacía:
		        prim := q                                //    - `q` se convierte en el primer nodo (prim).
		        *q.prox := q                             //    - El nodo apunta a sí mismo, formando un ciclo.
		    SINO                                         // Caso 2: Si la lista no está vacía:
		        p := prim                                //    - Inicializar `p` en el primer nodo.
		
		        // Recorrer la lista mientras no se vuelva al inicio (*p.prox <> prim) y
		        // el dato del nodo actual (*p.dato) sea menor que el nuevo dato (*q.dato).
		       
										 MIENTRAS (*p.prox <> prim) y (*p.dato < *q.dato) HACER
		            ant := p                             // Guardar el nodo actual como `ant` (anterior).
		            p := *p.prox                         // Avanzar al siguiente nodo.
		        	FM
		
		        // Caso 2a: Insertar antes del primer nodo (nuevo valor es el menor de todos).
		        SI p = prim ENT
		            k := prim                            // Inicializar `k` para buscar el último nodo.
		            MIENTRAS *k.prox <> prim HACER       // Recorrer la lista hasta encontrar el último nodo.
		                k := *k.prox                     // Avanzar al siguiente nodo.
		            FM
		            *k.prox := q                         // El último nodo apunta al nuevo nodo `q`.
		            *q.prox := p                         // El nuevo nodo `q` apunta al primer nodo actual (`p`).
		            prim := q                            // Actualizar `prim` para que apunte al nuevo nodo.
		
		        // Caso 2b: Insertar después del último nodo.
		        SINO SI *p.prox = prim ENT
		            *p.prox := q                         // El nodo actual `p` apunta al nuevo nodo `q`.
		            *q.prox := prim                      // El nuevo nodo `q` apunta al primer nodo (`prim`).
		
		        // Caso 2c: Insertar entre dos nodos.
		        SINO
		            *ant.prox := q                       // El nodo anterior (`ant`) apunta al nuevo nodo `q`.
		            *q.prox := p                         // El nuevo nodo `q` apunta al nodo actual (`p`).
		        FS
		    FS
		    ESC("Elemento agregado a la lista")          // Confirmar que el nodo fue insertado correctamente.
		FP
		
	PROCEDIMIENTO borrar() es
		ESC("ingrese el valor del elemento a eliminar:")
		LEER(elemento)
		SI prim= nil ENT
			esc("lista vacia")
		SINO
			p:= prim
			MIENTRAS (*p.prox <> prim) y (*p.dato <> elemento ) HACER
				ant:= p
				p:= *p.prox
			FM
			SI p.dato <> elemento ENTO
				esc("elemento no encontrado")
			SINO
				SI p= *p.prox ENTO
					prim:=nil
				SINO
					SI p = prim ENT  		//primerElemento
						k:= prim
						MIENTRAS *k.prox <> prim HACER
							k:= *k.prox
						FM
						*k.prox:=*p.prox
						prim:= *p.prox
					SINO
						*ant.prox:=*p.prox  //Medio y UltimoElemento
					FS
				FS
				DISPONER(P)
			FS
		FS
	FP



	PROCESO	
	
	prim:= nil
	repetir

		esc("1-Buscar")
		esc("2-Insertar")
		esc("3-Borrar")

		esc("que desea realizar:")leer(elec)
		SEGUN elec HACER
			1: buscar()
			2: insertar()	
			3: borrar()
		FS

		esc("quiere continuar? s/n")
		leer(op);

	hasta (op = "n")
fa













		