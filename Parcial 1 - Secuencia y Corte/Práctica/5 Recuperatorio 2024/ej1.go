//1. Si el pais no tiene historial destacado, historial -> P. Si el deporte del atleta comienza con vocal.
// Secuencia de salida con nombre del pais, finalizado con "#". Separar info de cada jugador con "%" y finalizar cada seccion con "@"
// pais#nombreApellidoAteta%otroAtleta@paisDos#atletaDos%otroAtleta

//2. Si el pais no tiene historial destacado, historial -> P. Si el atleta supera o iguala la edadUsuario
// Contar las veces que esto sucede, por delegacion

// Deporte(3caracteres)Edad(2caracteres)NombreyApellidoAtleta#lDeporte(3caracteres)Edad(2caracteres)NombreApellidoAtleta#@empieza otra delegacion

Accion ej1 es 
Ambiente
	historial: secuencia de caracteres
	h,d: caracter
	deportes: secuencia de caracteres
	igualan, superan

	funcion covertir(d: caracter): entero // no estoy segura de esta funcion tt
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

		si (edad > edadUsuario) entonces
			superan:= superan + 1
		sino
			si (edad = edadUsuario) entonces
				igualan:= igualan + 1 
			fs
		fs
	fp

	procedimiento salidaPais() es
		Repetir
			Esc(salida,h) ; Avz(historial,h)
		hasta h = "#"
	fp

	procedimiento saltarDeporteYVerificar() es
		Avz(deportes,d)
		si (d EN vocales) entonces
			Avz(deportes,d) ; Avz(deportes,d)
			deporte:= true
		fs	
	fp

	procedimiento salidaAtleta() es
		Repetir
			Esc(salida,d) ; Avz(deporte,d)
		hasta h = "#"
		Avz(deporte,d)
	fp

	procedimiento saltarAtleta() es
		Repetir
			Avz(deporte,d)
		hasta d = "#"
	fp

	
Algoritmo

	Iniciar() ; Crear(salida) 

	mientras NFDS(historial) hacer
		si (historial = "P") entonces
			salidaPais() ; Esc(salida,"#")
			saltarDeporteYVerificar() ; edadAtleta() 
			si (deporte = true) entonces
				salidaAtleta() ; Esc(salida,"%")
			sino
				saltarAtleta()
			fs
		fs
	fm

Fa