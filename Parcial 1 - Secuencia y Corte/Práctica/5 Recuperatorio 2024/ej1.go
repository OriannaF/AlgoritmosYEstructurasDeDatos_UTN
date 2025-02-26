//26/02/25 preparando el final :)

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
	igualan, superan, edad, deporte: entero

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
		si (d EN vocales) entonces
			Avz(deportes,d) ; Avz(deportes,d)
			deporte:= true
		fs	
		Avz(deportes,d)
	fp

	procedimiento salidaAtleta() es
		Repetir
			Avz(deporte,d) ; Esc(salida,d)
		hasta h = "#"
		Avz(deporte,d)
	fp

	procedimiento saltarAtleta() es
		Repetir
			Avz(deporte,d)
		hasta d = "#"
	fp

	procedimiento Iniciar() es
		Arr(historial); Arr(deportes)
		Avz(deporte,d) ; Avz(historial,h)
	fp

	
Algoritmo

	Iniciar() ; Crear(salida) 

	Esc("Ingrese la edad") ; leer(edadUsuario)

	mientras NFDS(historial) hacer
		si (historial = "P") entonces
			salidaPais() ; Esc(salida,"#")
			Repetir
				saltarDeporteYVerificar()  
				edadAtleta() //punto b
				si (deporte = true) entonces
					salidaAtleta() ; Esc(salida,"%")
				sino
					saltarAtleta()
				fs
			Hasta d = "@"

			Esc(salida,"@")

			//informar act b
			Esc("En esta delegacion, hay", igualan ,"deportistas de edad igual a la ingresada", superan ,"deportistas supera la edad ingresada.")
			superan:= 0 ; igualan := 0

		sino
			Repetir
				Avz(deporte,d) ; Avz(historial,h)
			Hasta d = "#"
			Avz(deporte,d) ; Avz(historial,h)
		fs

	fm

	cerrar(deporte) ; cerrar(historial)

Fa