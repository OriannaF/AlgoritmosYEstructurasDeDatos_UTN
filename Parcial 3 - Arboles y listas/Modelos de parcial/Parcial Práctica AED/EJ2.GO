//Tenemos una lista circular de entrada donde cada nodo contiene códigos binarios. Se desea copiar a una nueva lista doble aquellos nodos donde la cantidad de 0 no sea múltiplo de 3.
//Para el control, realizar una función recursivo que retorne un valor lógico, verdadero si los 0 son múltiplos de 3 y falso si no lo son. Por finalizar, mostrar la lista generada.

Función ceros(n,contador: entero): logica es
	Si ( n = 0 ) entonces // caso base
		Si ((contador MOD 3) = 0 ) entonces
			Ceros:= V //es múltiplo
		Sino
			Ceros:= F
		Fs
	Sino // caso recursivo
		Si ((n MOD 10) = 0) entonces
			Ceros(n DIV 10,contador +1)
		Sino
			Ceros(n DIV 10,contador)
		Fs
	Fs
Ff

Mientras *p.prox distinto PRIM hacer
	Si Ceros(*p.numero;0) = F entonces //si es falso no es múltiplo de tres
		Cargar()
	Fs
	P:= *p.prox
Fm

MostrarLista()