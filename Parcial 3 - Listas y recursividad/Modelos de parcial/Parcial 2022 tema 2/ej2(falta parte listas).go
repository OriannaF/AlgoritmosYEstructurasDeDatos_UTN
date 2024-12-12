//valor lógico
Función Presion(N: arreglo de [1…4] de entero; B: arreglo de [1...4] de entero ; i, promUno, promDos: entero): logica es
	Si ( i = 0 ) entonces //caso base
		Si (promUno/4) > 120 Y promDos/4 > 80 entonces
			Presion:= V
		Sino
			Presion:= F
		Fs
	Sino
	//caso Recursivo 
		Presion:= Presion(N; B; i-1 ; (promUno + N[i]) ; promDos + B[i])
	Fs
Ff

Mientras (p distinto NIL) entonces //lista
	Si ((presion(*p.presionUno; *p.presionDos; 4 ; 0; 0) = V) entonces
		AccionPorExito()
	Fs
	P:= *p.prox
Fm

//valor real con promedio

Función Presion(N: arreglo de [1…4] de entero; i,prom: entero): real es
	Si ( i = 0 ) entonces //caso base
		Presion:= prom / 4
	Sino
	//caso Recursivo 
		Presion:= Presion(N; i-1 ; ; prom + N[i] )
	Fs
Ff

Mientras (p distinto NIL) entonces //lista
	Si ((presion(*p.presionUno; 4 ; 0) < 120) Y ((presion(*p.presionUno; 4 ; 0) < 80) entonces
		AccionPorExito()
	Fs
	P:= *p.prox
Fm

//valor real sin prom (hizo juampy)
Funcion prom(A:Arreglo de [1..4] de enteros, i: entero): real es
	Si (i=1) entonces
		prom:= A[i]/4
	Sino
		prom:= A[i]/4 + prom(A, i-1)
	FS
FF

Mientras (p distinto NIL) entonces //lista
	Si ((presion(*p.presionUno; 4 ) < 120) Y ((presion(*p.presionUno; 4) < 80) entonces
		AccionPorExito()
	Fs
	P:= *p.prox
Fm