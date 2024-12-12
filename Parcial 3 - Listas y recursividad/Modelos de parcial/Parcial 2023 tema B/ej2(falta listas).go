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