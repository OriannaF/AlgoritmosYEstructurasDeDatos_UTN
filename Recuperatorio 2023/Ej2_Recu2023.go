accion parcial es
	ambiente 
		fecha = registro 
			dia: N(2)
			mes: N(2)
			año: N(4)
		finregistro
		clientes = registro
			nro_cliente: N(5)
			apyno: AN(20)
			dni: N(8)
			id_paquete: N(5)
			saldo: N(6)
			estado: ["SALDADO", "SALDO A FAVOR", "DEUDOR"]
			categoria:["SIMPLE", "PLATA", "ORO", "DIAMANTE"]
			puntos: N(3)
			fecha_b: fecha
		finregistro
		
		mae: archivo de clientes ordenado por nro_cliente
		r: clientes


		paquetes = registro
			id_paquete: N(5)
			nombre: AN(20)
			costo: N(6)
			destino: AN(15)
		finregistro
		paq: archivo de paquetes indexado por id_paquete
		p: paquetes

		a:arreglo[1..5, 1..13]
		cat: alfanumerico
		i, j, mayor: entero

		funcion fila(x: alfanumerico):entero es
			segun x hacer
				"simple": i:=1
				"plata": i:=2
				"oro": i:=3
				"diamante": i:=4
			finsegun
		finfuncion


	proceso
	abrir e/ (mae)
	leer(mae, r)
	para i:=1 a 5 hacer
		para j:= 13 hacer 
			a[i,j]:=0
		finpara
	finpara
	
	mientras nfda(mae) hacer
		si (r.estado = "SALDADO") o (r.estado = "SALDO A FAVOR") entonces
			i:= fila(r.categoria)
			j:= r.id_paquete
			a[i,j]:= a[i,j] + 1
			a[i,13]:= a[i,13] + 1
			a[5,j]:= a[5,j] + 1
			a[5,13]:= a[5,13] + 1
		finsi
		leer(mae, r)
	finsi

	mayor:=1	
	para j:=2 a 13 hacer
		si a[5,j] > a[mayor] entonces
			mayor:=j
		finsi
	finpara

	para i=1 hasta 4 hacer
		segun i hacer 
			1: cat:="SIMPLE"
			2: cat:="PLATA"
			3: cat:="ORO"
			4: cat:="DIAMANTE"
		finsegun
		para j=1 hasta 12 hacer
			p.idPaquete:= j
			Leer(paq,p)
			si EXISTE entonces
				escribir("para la categoria ", cat)
				escribir("Para el paquete: ",p.nombre,"de ID",j,"hubo",a[i,j]"paquetes adquiridos")
			finsi
		finpara
	finpara
	
	escribir("el paquete mas vendido es el paquete numero: ", mayor)

	escribir("cantidad total de paquetes saldados: " a[7,13])
finproceso	


		

