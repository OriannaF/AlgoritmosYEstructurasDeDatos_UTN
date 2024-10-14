Accion Pokemon_Go (PokeDex: arreglo [1...151] de alfanumerico) ES

Ambiente

	Usuario=Registro
		Clave=Registro
			Cod_Region: N(4)
			Cod_Usuario: N(10)
		Fin Registro
		Correo: AN(50)
		Experiencia: N(7)
	Fin Registro

	Arch_Usuario,Nuevo_Maestro: Archivo secuencial de Usuario ordenado por Clave
	reg_user,reg_sal,aux: Usuario

	Capturas=Registro
		Clave=Registro
			Cod_Region: N(4)
			Cod_Usuario: N(10)
		Fin Registro
		Cod_Pokemon: (1...151)
		Puntos_de_Experiencia: N(5)
		Fecha_de_Captura: N(8)
		Estado_Pokemon: ("E","I","D")
		Estado_Usuario: ("A","S","B")
	Fin Registro

	Arch_Capturas: Archivo secuencial de Capturas ordenado por Clave
	reg_cap: Capturas

	Procedimiento Leer_Usuario() ES
		Leer(Arch_Usuario,reg_user)
		Si FDA(Arch_Usuario) entonces
			reg_user.Clave:= HV
		Fin Si
	Fin Procedimiento

	Procedimiento Leer_Capturas() ES
		Leer(Arch_Capturas,reg_cap)

		Si (Cod_Pokemon <> reg_cap.Cod_Pokemon) entonces
			Corte_Pokemon()
		Fin Si

		Si FDA(Arch_Capturas) entonces
			reg_cap.Clave:= HV
		Fin Si
	Fin Procedimiento

	cont_desc,aux_desc: entero
	Poke_Descanso: alfanumerico
	Bandera: booleano

Proceso
	Abrir E/(Arch_Usuario)
	Leer_Usuario()
	Abrir E/(Arch_Capturas)
	Leer_Capturas()
	Abrir /S(Nuevo_Maestro)

	Cod_Pokemon:= reg_cap.Cod_Pokemon
	cont_desc:= 0
	aux_desc:= 0

	Mientras (reg_user.Clave <> HV) o (reg_cap.Clave <> HV) hacer

		Si (reg_user.Clave < reg_cap.Clave) entonces

			reg_sal:= reg_user
			Grabar(Nuevo_Maestro,reg_sal)
			Leer_Capturas()
		SiNo
			Si (reg_user.Clave = reg_cap.Clave) entonces
				Bandera:= Verdadero
				aux:= reg_user

				Mientras (aux.Clave = reg_cap.Clave) hacer
					Movimientos()
					Leer_Capturas()
				Fin Mientras

				Si (Bandera) entonces
					reg_sal:= aux
					Grabar(Nuevo_Maestro,reg_sal)
				Fin Si
				
				Leer_Usuario()
			SiNo
				Si (reg_cap.Estado_Usuario <> "A") entonces

					aux.Clave:= reg_cap.Clave
					Correo: ""
					Experiencia:= 0

					Mientras (aux.Clave = reg_cap.Clave) hacer
						Movimientos()
						Leer_Capturas()
					Fin Mientras

					Si (Bandera) entonces
						reg_sal:= aux
						Grabar(Nuevo_Maestro,reg_sal)
					Fin Si
				SiNo
					Leer_Capturas()
				Fin Si
			Fin Si
		Fin Si
	Fin Mientras

	Esc("El Pokemon en Descanso que pertenece a mas usuarios es: ",Poke_Descanso," con una cantidad de: ",cont_desc)

	Cerrar(Arch_Usuario)
	Cerrar(Arch_Capturas)
	Cerrar(Nuevo_Maestro)
Fin Accion 

Procedimiento Corte_Pokemon() ES

	Si (aux_desc > cont_desc) entonces

		i:= reg_cap.Cod_Pokemon
		Poke_Descanso:= PokeDex[i]
		cont_desc:= aux_desc
	Fin Si

	aux_desc:= 0
	Cod_Pokemon:= reg_cap.Cod_Pokemon
Fin Procedimiento

Procedimiento Movimientos() ES
	
	Si (reg_cap.Estado_Usuario <> "S") entonces

		Si (reg_cap.Estado_Pokemon = "E") entonces
	 
			aux.Experiencia:= aux.Experiencia + (reg_cap.Puntos_de_Experiencia * 2)
		SiNo
			aux.Experiencia:= aux.Experiencia + reg_cap.Puntos_de_Experiencia

			Si (reg_cap.Estado_Pokemon = "D")

				aux_desc:= aux_desc + 1
				
			Fin Si
		Fin Si
	SiNo
		Bandera:= Falso
	Fin Si
Fin Procedimiento 


/*
Cod_Pokemon:= reg_cap.Cod_Pokemon
i:= reg_cap.Cod_Pokemon
...
Si (Cod_Pokemon <> reg_cap.Cod_Pokemon) entonces
	aux_desc:= 0
	Cod_Pokemon:= reg_cap.Cod_Pokemon
	i:= reg_cap.Cod_Pokemon
Fin Si
...
aux_desc:= aux_desc + 1
...
Si (cantidad_obtenida > cont_mayor) entonces
	cont_desc:= aux_desc
	Poke_Descanso:= PokeDex[i]
Fin Si 
*/