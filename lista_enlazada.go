package lista

const (
	_CAPACIDAD_INICIAL = 10
	_MENSAJE_PANICO = "La lista está vacía"
	_MENSAJE_PANICO_ITERADOR = "El iterador termino de iterar"
)

type listaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
	cantidad int
}

type iteradorLista[T any] struct {
	actual *nodo[T]
	lista  *listaEnlazada[T]
}

type nodo[T any] struct {
	elemento T
	siguiente *nodo[T]
}

/*
En caso que se invoque a VerActual, Avanzar o Borrar sobre un iterador que ya haya iterado todos los elementos, debe entrar en pánico con un mensaje El iterador termino de iterar.
*/

func (iter *iteradorLista[T]) VerActual() T {
	if iter.actual == nil {
		panic(_MENSAJE_PANICO_ITERADOR)
	}
	return iter.actual.elemento
}

func (iter *iteradorLista[T]) HayAlgoMas() bool {
	return iter.actual != nil
}

func (iter *iteradorLista[T]) Avanzar() {
	if iter.actual == nil {
		panic(_MENSAJE_PANICO_ITERADOR)
	}
	iter.actual = iter.actual.siguiente
}

func (iter *iteradorLista[T]) Insertar(elemento T) {
	nuevoNodo := crearNodo(elemento)
	aux := iter.actual
	iter.actual = nuevoNodo
	nuevoNodo.siguiente = aux
	iter.lista.cantidad++
}

func (iter *iteradorLista[T]) Borrar() T {
	if iter.actual == nil {
		panic(_MENSAJE_PANICO_ITERADOR)
	} else if iter.actual == iter.lista.ultimo {
		
	}
	dato := iter.actual.elemento
	iter.actual = iter.actual.siguiente
	iter.lista.cantidad--
	return dato
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

//Iterador interno
func (lista *listaEnlazada[T]) Iterar() Iterador[T] {

}

//devuelve una instancia del iterador externo (el struct iteradorLista)
func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {

}
