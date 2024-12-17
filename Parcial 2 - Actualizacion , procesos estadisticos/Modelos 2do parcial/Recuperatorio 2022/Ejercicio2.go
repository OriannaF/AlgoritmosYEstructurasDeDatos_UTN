accion ej2 es
ambiente
    papers = registro
        dni: N(8)
        apellidoNombre: AN(50)
        email: AN(50)
        regional: AN(50)
        categoria: AN(50)
        tituloTrabajo: AN(50)
    fr

    archPaper: archivo de papers ordenado por dni
    p: papers

    regionales: arreglo de [1...31] de alfanumerico
    categorias: arreglo de [1...16] de alfanumerico
    A: arreglo de [1...32,1...17] de entero
    r,c,j,k,x,z,mayorPaper,posicionMayor: entero


algoritmo

    AbrirE/(archPaper) ; Leer(archPaper,p)
    mientras NFDA(archPaper) hacer

        //buscar regional en arreglo
        mientras (x < 31) Y ( p.regional <> regionales[x]) hacer
            x := x + 1
        fm
        //si se encuentra la regional, se busca categoria
        si (p.regional = regionales[x]) entonces
            //buscar categoria
            mientras ( z < 16) Y ( p.categoria <> categorias[z]) hacer
                z := z + 1
            fm

            si (p.categoria = categorias[z]) entonces
                j:= x ; k:= z
                A[j,k] := A[j,k] + 1 
                A[32,k] := A[32,k] + 1
                A[j,17] := A[j,17] + 1
                A[32,17] := A[32,17] + 1
            sino
                esc("error, categoria inexistente")
            fs
        sino
            esc("error, regional inexistente")
        fs
        
        Leer(archPaper,p)
    fm

    //punto 1, mostrar totales de papers por regional y categoria
    para (j:= 1 hasta 31) hacer
        para (k:= 1 hasta 16) hacer
            esc("Regional: ", regional[j] , "Categoria:" ,categoria[k], "Total:",A[j,k],".")
        fp
    fp

    //punto 2, regional con mas papers enviados
    para (j:=1 hasta 31) hacer
        si (A[j,17] > mayorPaper) entonces
            mayorPaper := A[j,17]
            posicionMayor := j
        fs
    fp

    esc("La regional con más papers es, " , A[j], "con ", mayorPaper, " papers.")

    cerrar(archPaper)
    
fa