//1. Si el pais no tiene historial destacado, historial -> P. Si el deporte del atleta comienza con vocal.
// Secuencia de salida con nombre del pais, finalizado con "#". Separar info de cada jugador con "%" y finalizar cada seccion con "@"
// pais#nombreApellidoAteta%otroAtleta@paisDos#atletaDos%otroAtleta

//2. Si el pais no tiene historial destacado, historial -> P. Si el atleta supera o iguala la edadUsuario
// Contar las veces que esto sucede, por delegacion

// Deporte(3caracteres)Edad(2caracteres)NombreyApellidoAtleta#lDeporte(3caracteres)Edad(2caracteres)NombreApellidoAtleta#@empieza otra delegacion

Accion ej1 es 
Ambiente
	secHistorial: secuencia de caracteres
	h,d: caracter
	secDeportistas: secuencia de caracteres

	funcion covertir(d: caracter): entero
		segun d hacer
			"1": convertir := 1
			"2": convertir := 2
			"3": convertir := 3
			"4": convertir := 4
			"5": convertir := 5
			"6": convertir := 6
			"7": convertir := 7
			"8": convertir := 8
			"9": convertir := 9
		fs
	ff

	procedimiento saltarDeporteYVerificar() es
		deporte:= d
		Avz(deportes,d) ; Avz(deportes,d)
	fp

	procedimiento edadAtleta() es
		Avz(deportes,d)
		edad:= convertir(d) * 10
		Avz(deportes,d)
		edad:= edad + convertir(d)
	fp

Algoritmo

	Iniciar() ; Crear(salida) 

	si (historial = "P") entonces
		salidaPais() ; Esc(salida,"#")
		saltarDeporteYVerificar() ; edadAtleta() 
		si (deporte EN vocales) entonces
			salidaAtleta() ; Esc(salida,"%")
		sino
			saltarAtleta()
		fs
	fs

Fa