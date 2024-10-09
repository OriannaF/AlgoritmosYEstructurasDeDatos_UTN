//Apilar
Leer(valor)
Prim:= nil
Mientras (valor <> "*") hacer // * es el tope
	Nuevo(p) //p va a puntar al elemento nuevo cuando se crea
	*p.dato:= valor //al dato del nuevo nodo se le agrega lo que ingreso el usuario
	*p.prox:= Prim //apunta p a donde estaba anteriormente prim, o sea al elemento siguiente
	Prim:= p //apunto prim al nuevo elemento
	Leer(valor)
fm

//Encolar

LEER (valor)

Prim:= nil
MIENTRAS (valor <> '*') HACER
    NUEVO (p)
    *p.Dato:= valor
    *p.Prox:= nil

    SI Prim = nil ENTONCES
        Prim = p
    SINO
        *a.Prox:= p
    FIN SI
    a:=p
    LEER(valor)
FIN MIENTRAS

//Carga ordenada

Prim:= nil
A:= nil;

LEER (valor) 
NUEVO (p) 
*p.dato:= valor 
q:= Prim 

MIENTRAS (q <> nil) y (*q.dato < valor) HACER
    a:= q
    q:= *q.prox
FIN MIENTRAS

SI (a = nil) ENTONCES
    *p.prox:= Prim
    Prim:= p
SINO
    *p.prox:= q
    *a.prox:= p
FIN SI








