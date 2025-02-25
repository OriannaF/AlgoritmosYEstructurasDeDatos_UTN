-- Excluyente


Mientras nfda(arch1) y nfda(arch2) hacer 
    Si r.nro_pelicula < r2.nro_pelicula entonces 
        Grabar(salida, r)
        Leer(arch1, r)
    Sino 

        Grabar(salida, r2)
        Leer(arch2, r2)
    FinSi

FinMientras

Mientras nfda(arch1) hacer 
    Grabar(salida, r)
    Leer(arch1, r)
FinMientras

Mientras nfda(arch2) hacer 
    Grabar(salida, r2)
    Leer(arch2, r2)
FinMientras


Cerrar(salida)


-- Incluyente

Procedimiento LeerArch1() entonces
    Leer(arch, r)
    Si FDA(arch) hacer
        r.nro_pelicula:= HV
    FinHacer
FinProcedimiento

Procedimiento LeerArch2() entonces
    Leer(arch2, r2)
    Si FDA(arch2) hacer
        r2.nro_pelicula:= HV
    FinHacer
FinProcedimiento

Mientras r.nro_pelicula <> HV o r2.nro_pelicula <> HV hacer 
    Si r.nro_pelicula < r2.nro_pelicula entonces 
        Grabar(salida, r)
        LeerArch1()
    Sino 

        Grabar(salida, r2)
        LeerArch2()
    FinSi

FinMientras

