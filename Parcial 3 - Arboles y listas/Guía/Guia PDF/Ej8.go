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
	elec: 1..3
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

	PROCEDIMIENTO insertar() es //CARGA ORDENADA
		ESC("Ingrese el valor que quiere insertar:")
		nuevo(q); LEER(*q.dato)
		SI prim= nil ENT
			prim:= q
			*q.prox:= q
		SINO	
			p:=prim
			MIENTRAS (*p.prox <> prim) y (*p.dato < *q.dato) HACER //BUSCAR COMO LUCAS HACEEE (LO MISMO QUE CARGA LISTA SIMPLE)
				ant:= p
				p:= *p.prox
			FM 
			SI p = prim ENTONCES  			//primerElemento
				k:= prim
				MIENTRAS *k.prox <> prim HACER
					k:=*k.prox
				FM
				*k.prox:=q
				*q.prox:=p
				prim:=q
			SINO
				SI *p.prox = prim ENTONCES  //ultimoElemento
					*p.prox:=q
					*q.prox:=prim
				SINO
					*ant.prox:=q       //Medio
					*q.prox:= p
				FS
			FS
		FS
		ESC("elemento agregado a la lista")
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

		esc("que desea realizar:") ; leer(elec)
		SEGUN elec HACER
			1: buscar()
			2: insertar()	
			3: borrar()
		FS

		esc("quiere continuar? s/n")
		leer(op);

	hasta (op = "n")
fa













		