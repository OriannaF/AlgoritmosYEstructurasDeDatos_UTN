//En  el  restaurante  ÑOQUIS  se  está  pensando  en  una  solución  informática  para  el  soporte  de  datos  del  nuevo sistema de atención a clientes. Se han decidido por LISTAS por su dinamismo en cuanto a la cantidad de elementos. 
//Diseñe un algoritmo que realice las siguientes funciones: 
//	 Añadir cliente al ser atendido (lista simple ordenada por Nombre del Cliente). 
//	 Registrar su consumo (Acumular el Total Consumido en valores de montos). 
//	 Realizar el cobro (emitir ticket con Nombre, Fecha, Número de Mesa y Total). 
//	 Eliminar del listado de atención. 
//La información almacenada debe mantenerse ordenada por Nombre del cliente.

Accion Ej6(PRIM: puntero a Nodo) es
	Ambiente
		fecha = registro
			dia: N(2)
			mes: N(2)
			anio: N(2)
		fr

		Nodo = registro
			nombreCliente: AN(50)
			fecha: fecha
			numeroMesa: N(7)
			totalConsumido: N(7)
			prox: puntero a Nodo
		fr
		p,q,PRIM: puntero a Nodo

		nombre,encontrado: AN(30)
		numeroMesa, consumo: N(3)

		Procedimiento AgregarCliente() es
			Nuevo(q)

			Esc("Ingrese los datos solicitados")
			*q.fecha:= fechaSistema() //supongo que con esto saco la fecha de la compu xd
			Esc("Nombre") ; Leer(nombre) ; *q.nombreCliente:= nombre
			Esc("Mesa") ; Leer(numeroMesa) ; *q.numeroMesa:= numeroMesa

			Si PRIM = NIL entonces
				*q.prox:= NIL
			sino
				*q.prox:= PRIM
			fs
			PRIM:= q
		fp

		Procedimiento Buscar() es
			Esc("Ingrese el nombre del cliente")
			Leer(nombre) ; Esc("Buscando!!")
			//búsqueda lineal con centinela, busca el primero y para, ya que supongo que tienen que ingresar un solo cliente por vez.
			ant:= NIL //POR SI HAY UN UNICO NODO
			p:= PRIM

			Si PRIM = NIL entonces
				Esc("Error, lista vacía")
			fs

			Mientras (p <> NIL) Y (*p.nombreCliente <> nombre) hacer
				ant := p
				p := *p.prox
			fm
			
			si (p <> NIL) entonces
				Escribir ('Se encontró al cliente')
				encontrado:= "SI"
			sino
				Escribir ('Error, cliente inexistente')
			fs
		fp

		Procedimiento AgregarConsumo() es
			Buscar()
			Si encontrado = "SI" entonces
				Esc("Agregue el consumo") ; Leer(consumo)
				*p.totalConsumido:= *p.totalConsumido + consumo
			sino
				Esc("Porfavor, agregue el cliente antes de ingresar el consumo")
			fs
		fp

		Procedimiento EliminarCliente() es
			// p -> clienteParaEliminar ant -> anterior a este cliente
			//  a -> b -> c
			Si *p.prox = NIL entonces // p -> c, PRIM -> a, ant -> b
				*ant.prox := NIL
			sino 
				//*p.prox <> NIL 

				//un caso: p -> b, PRIM -> a, ant -> a
				*ant.prox:= *p.prox 
				//pasa a estar asi a -> c (c es el prox de p. y la flecha es el prox de ant)

				//otro caso: p -> a, PRIM -> a, ant -> PRIM (que seria a). Pasaria a estar igual que el anterior
			fs
			Disponer(p)
		fp

		Procedimiento Cobrar() es
			Buscar()
			Si encontrado = "SI" entonces
				Esc("Fecha: ", *p.fecha.dia, "/", *p.fecha.mes, "/", *p.fecha.anio)
				Esc("Nombre: ", *p.nombreCliente)
				Esc("Mesa: ", *p.numeroMesa)
				Esc("Total Consumido: $", *p.totalConsumido)
				EliminarCliente()
			Sino
				Escribir("Ingrese al cliente antes de cobrar")
			fs
		fp

		

	Algoritmo


		PRIM:= NIL

		Esc("Sistema ÑOQUIS ;)")
		Esc("Ingrese una opción")

		Repetir
			Esc("1. Nuevo cliente")
			Esc("2. Nuevo consumo")
			Esc("3. Cuenta")
			Esc("4. Salir")

			Leer(op)

			Segun op hacer
				1: AgregarCliente()
				2: AgregarConsumo()
				3: Cobrar()
			fs
		Hasta op = 4
	Fa
Fa