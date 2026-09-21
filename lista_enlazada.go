package lista

const (
	_CAPACIDAD_INICIAL = 10
	_MENSAJE_PANICO = "La lista está vacía"
	_MENSAJE_PANICO_ITERADOR = "El iterador está al final de la lista"
)

type listaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
	cantidad int
}

type iteradorLista[T any] struct {
	actual *nodo[T]
}

type nodo[T any] struct {
	elemento T
	siguiente *nodo[T]
}

func crearNodo[T any](elemento T) *nodo[T] {
	return &nodo[T]{elemento: elemento}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return new(listaEnlazada[T])
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.cantidad == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nuevoNodo := crearNodo(elemento)
	lista.primero = nuevoNodo
	if lista.ultimo == nil {
		lista.ultimo = nuevoNodo
	}
	lista.cantidad++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nuevoNodo := crearNodo(elemento)
	if lista.ultimo != nil {
		lista.ultimo.siguiente = nuevoNodo
	}
	lista.ultimo = nuevoNodo
	lista.cantidad++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic(_MENSAJE_PANICO)
	}
	elemento := lista.primero.elemento
	lista.primero = lista.primero.siguiente
	lista.cantidad--
	if lista.EstaVacia() {
		lista.ultimo = nil
	}
	return elemento
}

func (lista *listaEnlazada[T]) BorrarUltimo() T {
	if lista.EstaVacia() {
		panic(_MENSAJE_PANICO)
	}
	var elemento T
	if lista.primero == lista.ultimo {
		elemento = lista.primero.elemento
		lista.primero = nil
		lista.ultimo = nil
	} else {
		actual := lista.primero
		for actual.siguiente != lista.ultimo {
			actual = actual.siguiente
		}
		elemento = lista.ultimo.elemento
		actual.siguiente = nil
		lista.ultimo = actual
	}
	lista.cantidad--
	return elemento
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic(_MENSAJE_PANICO)
	}
	return lista.primero.elemento
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic(_MENSAJE_PANICO)
	}
	return lista.ultimo.elemento
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.cantidad
}
