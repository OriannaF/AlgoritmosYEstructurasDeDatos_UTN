Accion Ej2 es
Ambiente

	Totales = registro
		barco: {1,2,3}
		fecha = registro
			dia: N(2)
			mes: N(2)
			ano: N(2)
		fr
		obra: {1,2,3,4,5}
		espectadores: N(3)
		total: N(5.2)
	fr

	archTot: archivo de Totales
	tot: Totales

	Arreglo = registro
		recaudado: N(5.2)
		espectadores: N(3)
	fr
	archArreglo: archivo de registroArreglo
	reg: registroArreglo

	a: arreglo de [1..3,1...6,1..13] de Arreglo
	i,j,k,obraMayor,barcoMayor,mesMayor,mayor:entero

	funcion menor(x:entero): entero es
		Para j=1 hasta 12 hacer
			Si A[x,6,j].recaudado < menor entonces
				menor:= j
			fs
		fp

		Esc("Para el barco",x,"el mes de menor recaudación es", menor,".")
	fp
	
Algoritmo

	AbrirE/(archTot) ; Abrir/S(Arreglo)

	Para i=1 hasta 4 hacer
		Para j=1 hasta 5 hacer
			Para k=1 hasta 13 hacer
				A[i,j,k]:= 0
			fp
		fp
	fp

	Mientras NFDA(archTot) hacer
		i:= tot.barco
		j:= tot.obra
		k:= tot.fecha.mes

		A[i,j,k].Arreglo.recaudado:= A[i,j,k] + tot.total
		A[i,j,13].Arreglo.recaudado:= A[i,j,13] + tot.total
		A[3,j,k].Arreglo.recaudado:= A[3,j,k] + tot.total
		A[i,6,k].Arreglo.recaudado:= A[i,6,k] + tot.total
		A[3,6,13].Arreglo.recaudado:= A[3,6,13] + tot.total

		A[i,j,k].Arreglo.espectadores:= A[i,j,k] + tot.espectadores
		A[i,6,k].Arreglo.espectadores:= A[i,6,k] + tot.espectadores
		A[i,j,13].Arreglo.espectadores:= A[i,j,13] + tot.espectadores
		A[3,j,k].Arreglo.espectadores:= A[3,j,k] + tot.espectadores
		A[3,6,13].Arreglo.espectadores:= A[3,6,13] + tot.espectadores
	fm

	mayor:= 0

	Para i=1 hasta 3 hacer
		menor(i)
	fp
	

	Para i=1 hasta 4 hacer
		Para j=1 hasta 6 hacer
			Para k=1 hasta 13 hacer
				Si A[i,j,k].espectadores > mayor entonces
					mayor:= A[i,j,k].espectadores 
					obraMayor:= j 
					mesMayor:= k
					barcoMayor:= i 
				fs
			fp
		fp
	fp

	Para i=1 hasta 4 hacer
		Para j=1 hasta 6 hacer
			Para k=1 hasta 13 hacer
				Si A[i,j,k].espectadores > mayor entonces
					mayor:= A[i,j,k].espectadores 
					obraMayor:= j 
					mesMayor:= k
					barcoMayor:= i 
				fs
			fp
		fp
	fp



Fa