// de Leonardo Brabo: https://github.com/LeonardoBrabo/Algoritmos_Resueltos/tree/master/10.Listas/Recestas%20Listas

Pilas y Colas Recetas

*Pilas(LIFO): last in first out: ultimo en entrar ,primero en salir: se va agregando los nuevos nodos de derecha a izquiera, es decir el puntero prim siempre va a apuntar al ultimo nodo en entrar

*Colas(FIFO): first in first out: primero en entrar, primero en salir: se va agregando los nuevos nodos de izquierda a derecha, es la carga mas comun.


Insertar como Pila
------------------

Accion cargar_pila (prim: puntero a nodo)es
	
	AMBIENTE
	nodo = registro
		dato: entero
		prox: puntero a nodo
	fr
	q: puntero a nodo

	PROCESO

	nuevo(q)
	esc("ingrese el valor que desea insertar")
	leer(*q.dato)
	*q.prox:= prim
	prim:= q

Fin_accion.



Insertar como Cola
------------------


Accion cargar_cola (prim: puntero a nodo)es
	
	AMBIENTE
	nodo = registro
		dato: entero
		prox: puntero a nodo
	fr
	q,p,ant: puntero a nodo

	PROCESO

	nuevo(q)
	esc("ingrese el valor que desea insertar")
	leer(*q.dato)
	si prim = nil entonces
		prim:= q
		*q.prox:= nil
	sino
		p:= prim
		mientras (p<>nil) hacer
			ant:= p
			p:= *p.prox
		fm
		*ant.prox:=q
		q.prox:= p
	fs	
Fin_accion.

otra opcion:
sin usar un ciclo =>	
			si prim = nil entonces
				*q.prox:= nil
				prim:= q
				P:= prim
			sino				//carga encolada 
				*P.prox:= q
				*q.prox:= nil
				P:= q
			fs


extraer los datos
------------------

PROCEDIMIENTO extraer() es
	si prim = nil entonces
		esc("lista vacia")
	sino
		p:= prim
		esc(*P.dato)
		prim:= *p.prox
		disponer(p)
		esc("elemento extraido")		
	fs
FP

