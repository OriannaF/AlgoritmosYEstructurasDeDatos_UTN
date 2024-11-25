// Curso 1:
// Nodo cabecera: [CursoNro: 1 | Cantidad: 2 | Legajo: 0].
// Alumnos:
// [CursoNro: 1 | Cantidad: 1 | Legajo: 123]
// [CursoNro: 1 | Cantidad: 1 | Legajo: 124].

// Curso 2:
// Nodo cabecera: [CursoNro: 2 | Cantidad: 1 | Legajo: 0].
// Alumnos:
// [CursoNro: 2 | Cantidad: 1 | Legajo: 125].

Accion Procesar_Lista_Alumnos(prim: puntero a NodoAlumno) es
Ambiente

	NodoAlumno = registro //nodo entrada
		legajo: entero
		apellido_nombre: an(30)
		comision: entero
		notas: arreglo[1..5] de entero
		prox: puntero a NodoAlumno
	fr

	NodoAula = registro // nodo salida
		cursoNro: entero
		cantidad: N(3)
		legajo: N(8)
		prox: puntero a NodoAula
		ant: puntero a NodoAula
	fr

	prim, p: puntero a NodoAlumno
	primAula, q, ultimo: puntero a NodoAula
	totalAlumnos, aulaMax, maxAlumnos: entero
	aulaAsignada: entero

Proceso

	// Inicialización de variables
	primAula := nil          // Lista doble enlazada de salida comienza vacía (lista de salida)
	totalAlumnos := 0        // Total de alumnos procesados
	aulaMax := 0             // Número del aula con más alumnos
	maxAlumnos := 0          // Máximo de alumnos en una sola aula

	// Iterar por la lista original de alumnos
	p := prim                // inicio de la lista original
	mientras (p <> nil) hacer // Recorremos la lista mientras existan nodos

		// Verificar si el alumno no está libre  
		si EsAlumnoLibre(p.legajo) = falso entonces  //si no esta libre es regular 
			// Asignar aula  
			aulaAsignada:= AsignarAula(p.legajo)
	
			// Buscar el nodo cabecera del aula en la lista doble enlazada
			q := primAula                 // Iniciamos la búsqueda desde el inicio de la lista doble
			mientras (*q.prox <> nil) Y (*q.cursoNro <> aulaAsignada) hacer // Recorremos hasta encontrar el aula
				q := *q.prox              // Avanzamos al siguiente nodo
			fm

			// Incrementar la cantidad de alumnos en el nodo cabecera
			*q.cantidad := q.cantidad + 1

			// Si no se encuentra el nodo cabecera del aula, se crea uno nuevo
			si *q.prox = nil entonces
				nuevo(q)                // Creamos un nuevo nodo
				*q.cursoNro := aulaAsignada // Asignamos el número del aula
				*q.cantidad := 0         // Inicializamos la cantidad de alumnos a 0
				*q.legajo := 0           // Los nodos cabecera no tienen legajos

				// Insertar el nodo cabecera en la lista doble enlazada
				si primAula = nil entonces  // Si la lista está vacía
					primAula := q           // Este nodo será el primero
					ultimo := q             // También será el último
					*q.prox := primAula
					*q.ant := primAula
				sino                      // Si ya hay nodos en la lista
					*ultimo.prox := q       // Lo enlazamos al final
					*q.ant := ultimo
					*q.prox := primAula
					ultimo := q            // Actualizamos el último nodo
				fs
			fs

	
			

			// Crear un nodo intermedio para el alumno
			nuevo(nuevoNodo)
			*nuevoNodo.cursoNro := aulaAsignada   // Aula asignada
			*nuevoNodo.cantidad := 1             // Nodo intermedio siempre representa 1 alumno
			nuevoNodo.legajo := p.legajo        // Asignamos el legajo del alumno

			// Insertar el nodo intermedio después del nodo cabecera
			*nuevoNodo.prox := q.prox            // Conectamos al siguiente nodo
			si q.prox <> nil entonces           // Si no es el último nodo
				*q.prox.ant := nuevoNodo        // Ajustamos el enlace del siguiente nodo
			fs
			*q.prox := nuevoNodo                 // Enlazamos el nodo cabecera al nuevo nodo
			*nuevoNodo.ant := q                  // Conectamos al nodo anterior

	
			// Incrementar contador de alumnos  
			totalAlumnos := totalAlumnos + 1  
		fs  
	
		// Eliminar el nodo actual de la lista original  
		p := p.prox  
	fm  
	
	// Calcular promedio de alumnos por aula  
	promedio := totalAlumnos / 6  
	
	// Determinar aula con mayor cantidad de alumnos  
	// Calcular promedio de alumnos por aula
	promedio := totalAlumnos / 6

	// Determinar el aula con mayor cantidad de alumnos
	q := primAula                     // Iniciamos desde el primer nodo cabecera
	mientras (q <> nil) hacer          // Recorremos la lista doble enlazada
		si (q.cantidad > maxAlumnos) entonces // Si el aula tiene más alumnos que el máximo actual
			maxAlumnos := q.cantidad       // Actualizamos el máximo
			aulaMax := q.cursoNro          // Guardamos el número del aula
		fs
		q := q.prox                     // Avanzamos al siguiente nodo
	fm

	
	// Mostrar resultados  
	Esc("Promedio de alumnos por aula: ", promedio)  
	Esc("Aula con mayor cantidad de alumnos: ", aulaMax)  
	

fp

