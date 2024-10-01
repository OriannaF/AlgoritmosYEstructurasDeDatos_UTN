Accion ej1 es
Ambiente
	Stock = registro
		clave = registro
			productoId: N(5)
		fr
		stock: N(5)
	fr
	archMae: archivo de stock ordenado por clave
	mae: stock

	AUXILIAR = registro
		clave = registro
			productoId: N(5)
		fr
		stock: N(5)
	fr
	archAux: archivo de stock ordenado por clave
	aux: stock

	movimientosDiarios = registro
		clave = registro
			productoId: N(5)
		fr
		clienteId: N(5)
		tipo: "C","D" //compra devolucion
		cantidad: N(4)
		precioUnitario: N(10)
		total: N(10)
		tipoEnvio: "ENVIO","RETIRO"
	fr

	archMov: archivo de movimientosDiarios ordenado por clave
	mov: movimientosDiarios

	procedimiento LeerMae() es
		Leer(archMae,mae)
		Si FDA(archMae) entonces
			mae.clave:= HV
		fs
	fp

	procedimiento LeerMov() es
		Leer(archMov,mov)
		Si FDA(archMov) entonces
			mov.clave:= HV
		fs
	fp

	productos = registro
		clave = registro
			productoId: nombre(5)
		fr
		nombre: AN(40)
		descripcion: AN(80)
		rubro: "LIMPIEZA", "CARNICERIA", "VERDULERIA", "BAZAR", "PANADERIA"
	fr

	archProd: archivo de productos indexado por clave
	prod: productos

	domicilio: entero
Algoritmo

	AbrirE/(archMae) ; AbrirE/(archMov) ; AbrirE/(archProd) ; Abrir/S(archAux)
	domicilio:= 0
	Mientras (mae.clave <> HV) o (mov.clave <> HV) hacer
		Si mae.clave<mov.clave entonces
			aux:=mae
			Grabar(archAux,aux)
		sino
			Si mae.clave=mov.clave entonces
				//hay actualizacion para el prod
				Mientras mae.clave=mov.clave hacer
					Si mov.tipo="C" hacer
						Si mae.stock >= mov.cantidad entonces
							aux.stock:= mae.stock - mov.cantidad
						sino
							prod.productoId:= mae.clave
							Leer(archProd,prod)
							Si EXISTE entonces	
								Esc("Error, falta stock para" prod.nombre " de " mae.stock - mov.cantidad " unidades.")
							sino
								Esc(Error, producto no encontrado)
							fs

						fs
						LeerMae() ; LeerMov()
					fs

					Si mov.tipo = "D" hacer
						aux.stock:= mae.stock + mov.cantidad
						domicilio:= domicilio + 1 
						LeerMae() ; LeerMov()
					fs
				fm

				
				
			sino
				Si mae.clave>mov.clave entonces
					Si mov.tipo = "C" entonces
						prod.productoId:= mov.clave
						Leer(archProd,prod)
						Si EXISTE entonces	
							Esc("Error, falta stock para" prod.nombre " de " mov.cantidad " unidades.")
						sino
							Esc(Error, producto no encontrado)
						fs
					sino
						Si mov.tipo = "D" entonces
							aux.productoId:= mov.productoId
							aux.stock:= mov.cantidad
						fs
					fs
					LeerMov()
				fs
			fs
		fs
	fm
fa

Esc("La cantidad a enviar a domicilio es ", domicilio)


