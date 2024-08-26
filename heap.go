package cola_prioridad

const (
	CAPACIDAD_INICIAL  = 4
	INICIO             = 0
	FACTOR_REDIMENSION = 2
	AGRANDAMIENTO      = 4
	ERROR_COLA_VACIA   = "La cola esta vacia"
)

type heap[T any] struct {
	datos    []T
	cantidad int
	cmp      func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{datos: make([]T, CAPACIDAD_INICIAL), cmp: funcion_cmp}
}

func heapify[T any](arr *[]T, cmp func(T, T) int, largo int) {
	for i := largo - 1; i >= 0; i-- {
		downheap(i, *arr, largo, cmp)
	}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	if len(arreglo) == 0 {
		return CrearHeap[T](funcion_cmp)
	}

	arr := make([]T, len(arreglo))
	copy(arr, arreglo)

	heapify(&arr, funcion_cmp, len(arreglo))

	heap := &heap[T]{datos: arr, cantidad: len(arreglo), cmp: funcion_cmp}
	return heap
}

func (heap heap[T]) EstaVacia() bool {
	return heap.cantidad == INICIO
}

func (heap *heap[T]) Encolar(dato T) {
	if heap.cantidad == cap(heap.datos) {
		heap.redimension(heap.cantidad * FACTOR_REDIMENSION)
	}
	heap.datos[heap.cantidad] = dato
	upheap(heap.cantidad, heap.datos, heap.cmp)
	heap.cantidad++
}

func (heap heap[T]) VerMax() T {
	if heap.EstaVacia() {
		panic(ERROR_COLA_VACIA)
	}
	return heap.datos[INICIO]
}

func (heap *heap[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic(ERROR_COLA_VACIA)
	}
	if cap(heap.datos) > ((heap.cantidad) * AGRANDAMIENTO) {
		heap.redimension(cap(heap.datos) / FACTOR_REDIMENSION)
	}
	dato := heap.datos[INICIO]
	heap.cantidad--
	swap(&heap.datos[INICIO], &heap.datos[heap.cantidad])
	downheap(INICIO, heap.datos, heap.cantidad, heap.cmp)
	return dato
}

func (heap heap[T]) Cantidad() int {
	return heap.cantidad
}

func (heap *heap[T]) redimension(nuevoLargo int) {
	nuevo := make([]T, nuevoLargo)
	copy(nuevo, heap.datos)
	heap.datos = nuevo
}

func swap[T any](a, b *T) {
	*a, *b = *b, *a
}

func upheap[T any](indice int, datos []T, cmp func(T, T) int) {
	padre := (indice - 1) / 2
	if cmp(datos[indice], datos[padre]) <= 0 {
		return
	}
	swap(&datos[indice], &datos[padre])
	upheap(padre, datos, cmp)
}

func downheap[T any](indice int, arr []T, largo int, cmp func(T, T) int) {
	hijo_izq := 2*indice + 1
	hijo_der := 2*indice + 2
	if largo <= hijo_izq {
		return
	}

	reemplazante := hijo_izq

	if largo > hijo_der {
		reemplazante = max(arr, cmp, hijo_izq, hijo_der)
	}

	if cmp(arr[indice], arr[reemplazante]) < 0 {
		swap(&arr[indice], &arr[reemplazante])
		downheap(reemplazante, arr, largo, cmp)
	}
}

func max[T any](arr []T, cmp func(T, T) int, izq, der int) int {
	if cmp(arr[izq], arr[der]) > 0 {
		return izq
	}
	return der
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	heapify(&elementos, funcion_cmp, len(elementos))
	heapSort_aux(elementos, INICIO, len(elementos)-1, funcion_cmp)
}

func heapSort_aux[T any](elementos []T, inicio, fin int, funcion_cmp func(T, T) int) {
	if inicio > fin {
		return
	}
	swap(&elementos[inicio], &elementos[fin])
	downheap(inicio, elementos, fin, funcion_cmp)
	heapSort_aux(elementos, inicio, fin-1, funcion_cmp)
}

func (heap *heap[T]) Invertir() {
	nueva_comparacion := func(a T, b T) int {
		nueva := heap.cmp(a, b)
		return nueva * -1
	}
	heap.cmp = nueva_comparacion
	heapify(&heap.datos, nueva_comparacion, heap.cantidad)
}
