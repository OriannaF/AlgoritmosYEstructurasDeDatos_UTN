//Genere  un  algoritmo  que  recorra  una  secuencia  texto  y  genere  una  lista  simplemente  encadenada  con  la 
//frecuencia de utilización de cada letra. La lista debe mantenerse ordenada alfabéticamente y al final informar cual fue 
//la frecuencia de cada letra y cuáles fueron la de mayor y menor frecuencia

Accion Ej7 es
Ambiente
	Nodo = registro
		letra: caracter
		frecuencia: N(3)
		valor: N(2) //valor de la letra para poder ordenar
		prox: puntero a Nodo
	fr
	p,PRIM: puntero a Nodo
	sec: secuencia de caracter
	v: caracter
	total: entero

	Funcion Mapeo(x:caracter):entero es
		Segun x hacer
			"A" : Mapeo:= 1
			...
			"Z": Mapeo:= 27
		fs	
	fp

	Procedimiento CargaOrdenada() es
		a:= NIL
		q:= PRIM 

		mientras (q <> nil) y (*q.valor < valor) hacer 
			a:= q
			q:= *q.prox
		fm

		si (a = nil) entonces
			*p.prox:= Prim
			Prim:= p
		sino
			*p.prox:= q
			*a.prox:= p
		fs
	fp




Algoritmo
	p:= PRIM
	PRIM:= NIL
	
	Arr(sec,v) ; Avz(sec,v)

	Mientras NFDS(sec) hacer
		
		valor:= Mapeo(v)

		//busco la letra en la lista
		mientras (*p.letra <> v) hacer
			p:= *p.prox
		fm

		si *p.letra = v entonces //si está, añado una nueva aparición
			*p.frecuencia:= *p.frecuencia + 1 
		sino //si no está creo nuevo nodo
			//Nodo nuevo
			Nuevo(p)
			*p.letra:= v
			*p.frecuencia:= *p.frecuencia + 1 
			*p.valor:= valor
			CargaOrdenada()
		fs 
		
		Avz(sec,v)
	fm

	//menor y mayor frecuencia
	si (PRIM = NIL) entonces
		esc("La lista está vacía")
	fs

	mayor := 0
	menor := 0

	p := PRIM

	mientras (p <> NIL) hacer
		si (*p.frecuencia > mayor) entonces
			mayor := *p.frecuencia
		fs

		si (*p.frecuencia < menor) entonces
			menor := *p.frecuencia
		fs

		p := *p.prox
	fm

	esc("El mayor valor de frecuencia es: ", mayor)
	esc("El menor valor de frecuencia es: ", menor)

	Cerrar(sec)
Fa