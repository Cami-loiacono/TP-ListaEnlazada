package lista_test

import (
	"testing"
	TDALista "lista" 
	"github.com/stretchr/testify/require"
)

const _MENSAJE_LISTA_VACIA = "La lista esta vacia"
const _MENSAJE_ITERADOR_TERMINADO = "El iterador termino de iterar"

// ------------------------- Lista vacía -------------------------

func TestListaVaciaEsVacia(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
}

func TestListaVaciaPanicVerPrimero(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	require.PanicsWithValue(t, _MENSAJE_LISTA_VACIA, func() { lista.VerPrimero() })
}

func TestListaVaciaPanicVerUltimo(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	require.PanicsWithValue(t, _MENSAJE_LISTA_VACIA, func() { lista.VerUltimo() })
}

func TestListaVaciaPanicBorrarPrimero(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	require.PanicsWithValue(t, _MENSAJE_LISTA_VACIA, func() { lista.BorrarPrimero() })
}

// ------------------------- InsertarPrimero -------------------------

func TestInsertarPrimeroUnElemento(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
}

func TestInsertarPrimeroVariosElementosQuedanEnOrdenInverso(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)

	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
}

// ------------------------- InsertarUltimo -------------------------

func TestInsertarUltimoUnElemento(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
}

func TestInsertarUltimoVariosElementosQuedanEnOrden(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())
}

func TestInsertarPrimeroYUltimoCombinados(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(3)

	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())
	require.Equal(t, 3, lista.Largo())
}

// ------------------------- BorrarPrimero -------------------------

func TestBorrarPrimeroDevuelveYQuitaElElemento(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	require.Equal(t, 1, lista.BorrarPrimero())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 2, lista.VerPrimero())

	require.Equal(t, 2, lista.BorrarPrimero())
	require.Equal(t, 3, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
}

func TestBorrarPrimeroHastaVaciarYVolverAUsar(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia())

	// La lista se puede seguir usando normalmente tras vaciarse.
	lista.InsertarUltimo(10)
	require.Equal(t, 10, lista.VerPrimero())
	require.Equal(t, 10, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
}

// ------------------------- Genericidad -------------------------

func TestListaDeStrings(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarUltimo("a")
	lista.InsertarUltimo("b")
	require.Equal(t, "a", lista.VerPrimero())
	require.Equal(t, "b", lista.VerUltimo())
}

type punto struct {
	x, y int
}

func TestListaDeStructs(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[punto]()
	lista.InsertarUltimo(punto{1, 2})
	lista.InsertarUltimo(punto{3, 4})
	require.Equal(t, punto{1, 2}, lista.VerPrimero())
	require.Equal(t, punto{3, 4}, lista.VerUltimo())
}

// ------------------------- Volumen -------------------------

func TestVolumenInsertarUltimoYBorrarPrimero(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	const n = 10000

	for i := 0; i < n; i++ {
		lista.InsertarUltimo(i)
	}
	require.Equal(t, n, lista.Largo())
	require.Equal(t, 0, lista.VerPrimero())
	require.Equal(t, n-1, lista.VerUltimo())

	for i := 0; i < n; i++ {
		require.Equal(t, i, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())
}

func TestVolumenInsertarPrimero(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	const n = 10000

	for i := 0; i < n; i++ {
		lista.InsertarPrimero(i)
	}
	require.Equal(t, n, lista.Largo())
	require.Equal(t, n-1, lista.VerPrimero())
	require.Equal(t, 0, lista.VerUltimo())
}

// ------------------------- Iterador interno -------------------------

func TestIterarListaVaciaNoLlamaFuncion(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	llamado := false
	lista.Iterar(func(v int) bool {
		llamado = true
		return true
	})
	require.False(t, llamado)
}

func TestIterarRecorreTodosLosElementosEnOrden(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 5; i++ {
		lista.InsertarUltimo(i)
	}

	var recorridos []int
	lista.Iterar(func(v int) bool {
		recorridos = append(recorridos, v)
		return true
	})

	require.Equal(t, []int{1, 2, 3, 4, 5}, recorridos)
}

func TestIterarConCorteAnticipado(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 5; i++ {
		lista.InsertarUltimo(i)
	}

	var recorridos []int
	lista.Iterar(func(v int) bool {
		recorridos = append(recorridos, v)
		return v < 3 // corta apenas visita el 3
	})

	require.Equal(t, []int{1, 2, 3}, recorridos)
}

func TestIterarCorteEnElPrimerElemento(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 5; i++ {
		lista.InsertarUltimo(i)
	}

	var recorridos []int
	lista.Iterar(func(v int) bool {
		recorridos = append(recorridos, v)
		return false // corta inmediatamente
	})

	require.Equal(t, []int{1}, recorridos)
}

func TestIterarPermiteAcumularConVariableExterna(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 4; i++ {
		lista.InsertarUltimo(i)
	}

	suma := 0
	lista.Iterar(func(v int) bool {
		suma += v
		return true
	})
	require.Equal(t, 10, suma)
}

func TestIterarNoModificaLaLista(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)

	lista.Iterar(func(v int) bool { return true })

	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
}

// ------------------------- Iterador externo: básico -------------------------

func TestIteradorExternoListaVacia(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	require.False(t, iter.HayAlgoMas())
}

func TestIteradorExternoListaVaciaPanics(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	require.PanicsWithValue(t, _MENSAJE_ITERADOR_TERMINADO, func() { iter.VerActual() })
	require.PanicsWithValue(t, _MENSAJE_ITERADOR_TERMINADO, func() { iter.Avanzar() })
	require.PanicsWithValue(t, _MENSAJE_ITERADOR_TERMINADO, func() { iter.Borrar() })
}

func TestIteradorExternoRecorreTodosLosElementosEnOrden(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 5; i++ {
		lista.InsertarUltimo(i)
	}

	iter := lista.Iterador()
	var recorridos []int
	for iter.HayAlgoMas() {
		recorridos = append(recorridos, iter.VerActual())
		iter.Avanzar()
	}
	require.Equal(t, []int{1, 2, 3, 4, 5}, recorridos)
	require.False(t, iter.HayAlgoMas())
}

func TestIteradorExternoAvanzarCambiaElActual(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)

	iter := lista.Iterador()
	require.Equal(t, 1, iter.VerActual())
	iter.Avanzar()
	require.Equal(t, 2, iter.VerActual())
	iter.Avanzar()
	require.False(t, iter.HayAlgoMas())
}

func TestIteradorExternoPanicVerActualAlTerminar(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	iter := lista.Iterador()
	iter.Avanzar()
	require.PanicsWithValue(t, _MENSAJE_ITERADOR_TERMINADO, func() { iter.VerActual() })
}

func TestIteradorExternoPanicAvanzarAlTerminar(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	iter := lista.Iterador()
	iter.Avanzar()
	require.PanicsWithValue(t, _MENSAJE_ITERADOR_TERMINADO, func() { iter.Avanzar() })
}

func TestIteradorExternoPanicBorrarAlTerminar(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	iter := lista.Iterador()
	iter.Avanzar()
	require.PanicsWithValue(t, _MENSAJE_ITERADOR_TERMINADO, func() { iter.Borrar() })
}

// ------------------------- Iterador externo: casos pedidos por la consigna -------------------------

// 1. Insertar en la posición en la que se crea el iterador inserta al principio.
func TestIteradorExternoInsertarAlCrearloInsertaAlPrincipio(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	iter.Insertar(1)

	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, iter.VerActual())
}

// 2. Insertar con el iterador al final es equivalente a insertar al final de la lista.
func TestIteradorExternoInsertarAlFinalEquivaleAInsertarUltimo(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)

	iter := lista.Iterador()
	for iter.HayAlgoMas() {
		iter.Avanzar()
	}
	iter.Insertar(3)

	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 3, lista.VerUltimo())
	// El iterador, tras insertar al final, vuelve a tener un elemento actual.
	require.True(t, iter.HayAlgoMas())
	require.Equal(t, 3, iter.VerActual())
}

// 3. Insertar un elemento en el medio lo deja en la posición correcta.
func TestIteradorExternoInsertarEnElMedio(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	iter.Avanzar() // ahora apunta a 3
	iter.Insertar(2)

	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 2, iter.VerActual())

	var recorridos []int
	lista.Iterar(func(v int) bool {
		recorridos = append(recorridos, v)
		return true
	})
	require.Equal(t, []int{1, 2, 3}, recorridos)
}

// 4. Borrar el elemento al crear el iterador cambia el primero de la lista.
func TestIteradorExternoBorrarAlCrearloCambiaElPrimero(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	borrado := iter.Borrar()

	require.Equal(t, 1, borrado)
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 2, iter.VerActual())
}

// 5. Borrar el último elemento con el iterador cambia el último de la lista.
func TestIteradorExternoBorrarUltimoCambiaElUltimo(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	iter.Avanzar()
	iter.Avanzar() // apunta al último (3)
	borrado := iter.Borrar()

	require.Equal(t, 3, borrado)
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 2, lista.VerUltimo())
	require.False(t, iter.HayAlgoMas())
}

// 6. Borrar un elemento del medio hace que ya no esté en la lista.
func TestIteradorExternoBorrarEnElMedioLoQuitaDeLaLista(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	iter.Avanzar() // apunta a 2
	borrado := iter.Borrar()

	require.Equal(t, 2, borrado)
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 3, iter.VerActual())

	var recorridos []int
	lista.Iterar(func(v int) bool {
		recorridos = append(recorridos, v)
		return true
	})
	require.Equal(t, []int{1, 3}, recorridos)
	// El elemento borrado ya no aparece en el recorrido.
	require.NotContains(t, recorridos, 2)
}

// 7. Casos borde adicionales del iterador externo.

func TestIteradorExternoBorrarHastaVaciarLaLista(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)

	iter := lista.Iterador()
	iter.Borrar()
	iter.Borrar()

	require.True(t, lista.EstaVacia())
	require.False(t, iter.HayAlgoMas())
}

func TestIteradorExternoInsertarEnListaVacia(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(1)

	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 1, iter.VerActual())
}

func TestIteradorExternoInsertarVariasVecesEnLaMismaPosicion(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	iter.Insertar(2)
	iter.Insertar(1)

	require.Equal(t, 3, lista.Largo())
	var recorridos []int
	lista.Iterar(func(v int) bool {
		recorridos = append(recorridos, v)
		return true
	})
	require.Equal(t, []int{1, 2, 3}, recorridos)
}

func TestDosIteradoresIndependientesSobreLaMismaLista(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iter1 := lista.Iterador()
	iter2 := lista.Iterador()

	iter1.Avanzar()
	// iter2 no debería verse afectado por el avance de iter1.
	require.Equal(t, 1, iter2.VerActual())
	require.Equal(t, 2, iter1.VerActual())
}

func TestIteradorExternoRecorrerBorrandoTodosLosElementos(t *testing.T) {
	t.Parallel()
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 5; i++ {
		lista.InsertarUltimo(i)
	}

	iter := lista.Iterador()
	for iter.HayAlgoMas() {
		iter.Borrar()
	}

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
}
