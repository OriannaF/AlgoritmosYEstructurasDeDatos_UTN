Accion Ej es
Ambiente
	MAE_LIBROS = registro
		clave = registro
			ID_Libro: N()
		fr
		Tema: AN(20)
		Titulo: AN(30)
		Cant_Actual: N(5)
	fr

	MOVIMIENTOS = registro
		ID_Libro: N(5)
		Cod_Mov: AN(1)
		Tema: AN(20)
		Titulo: AN(30)
	
	fr

	Maestro: archivo de MAE_LIBROS ordenado por ID_Libro
	Movimientos: archivo de MOVIMIENTOS ordenado por ID_Libro, Cod_Mov

	Mat_Basicas,Libros_ISI,Libros_IEM,Libros_IQ: archivo de MAE_LIBROS

	Mae,aux: MAE_LIBROS
	Mov: MOVIMIENTOS

	procedimiento leerMae() es
		leer(maestro,mae)
		si FDA (maestro) entones
			mae.ID_Libros := HV
		finsi
	fp

	procedimiento leerMov() es
		leer(movimiento,mov)
		si FDA (movimientos) enttones
			mov.ID_libro := HV
		finsi
	fp
	
	procedimiento AbrirArchivos() es
		abrir E/(Maestro)
		LeerMae()
		abrir E/(Movimientos)
		LeerMov()
		abrir S/(Mat_Basicas)
		abrir S/(Libros_ISI)
		abrir S/(Libros_IEM)
		abrir S/(Libros_IQ)
	fp

	procedimiento CerrarArchivos() es
		cerrar(Maestro)
		cerrar(Movimientos)
		cerrar(Mat_Basicas)
		cerrar(Libros_ISI)
		cerrar(Libros_IEM)
		cerrar(Libros_IQ)
	fp

Algoritmo
	AbrirArchivos()
	MIENTRAS (mae.id_libro <> HV) O (mov.id_libro <> HV) HACER
		SI mae.id_libro < Mov.id_libro ENTONCES //no hubo modificaciones
			//maestro
			SEGUN mae.tema HACER
				= “MB”: Grabar(Mat_Basicas, mae)
				= “ISI”: Grabar(Libros_ISI, mae)
				= “IEM”: Grabar(Libros_IEM, mae)
				= “IQ”: Grabar(Libros_IQ, mae)
			Fs
			LeerMae()
		Sino 
			Si mae.clave = mov.clave entonces 
				SI mov.Cod_Mov = “A”ENTONCES
					aux := mae
					MIENTRAS aux.id_libro = mov.id_libro HACER
						//SE PUEDE AGREGAR CONTROL DE BAJAS
						aux.cant_actual := aux.cant_actual + 1
						leer_Mov()
					FINMIENTRAS

					SEGUN mae.tema HACER
						= “MB”: Grabar(Mat_Basicas, mae)
						= “ISI”: Grabar(Libros_ISI, mae)
						= “IEM”: Grabar(Libros_IEM, mae)
						= “IQ”: Grabar(Libros_IQ, mae)
					FINSEGUN
				finsi
				//en caso de baja no se debe copiar a la salida
				leer_mae()
			Sino
				SI mae.id_libro > mov.id_libro ENTONCE//solo puede ser alta
					SI mov.Cod_Mov = "A" ENTONCES
						aux.id_Libro:= mov.id_Libro
						aux.tema := mov.tema
						aux.titulo := mov.titulo
						aux.cant_actual:= 1
						leer_Mov()
						MIENTRAS aux.id_Libro = mov.id_Libro HACER
							aux.cant_actual := aux.cant_actual + 1
							leer_Mov()
							//SE PODRIA AGREGAR CONTROL DE BAJAS
						FINMIENTRAS
						SEGUN mae.tema HACER
							= “MB”: Grabar(Mat_Basicas, mae)
							= “ISI”: Grabar(Libros_ISI, mae)
							= “IEM”: Grabar(Libros_IEM, mae)
							= “IQ”: Grabar(Libros_IQ, mae)
						FINSEGUN
					SINO
						ESC("ERROR")
					fs
				fs
			fs
		fs
	fm
	LeerMae()
	CerrarArchivos()
fa