package cola_prioridad_test

import (
	"math/rand"
	"strings"
	TDAColaPrioridad "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

func cmp(a, b int) int { return a - b }
func TestHeapVacio(t *testing.T) {
	t.Log("Test que comporueba que una cola de prioridad vacia se comporte como tal")
	heap := TDAColaPrioridad.CrearHeap[int](cmp)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestEncolarYDesencolarUnElemento(t *testing.T) {
	t.Log("Test de funcionamiento basico con un elemento de la cola de prioridad")
	heap := TDAColaPrioridad.CrearHeap[int](cmp)
	heap.Encolar(10)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 10, heap.VerMax())
	require.EqualValues(t, 10, heap.Desencolar())
	require.True(t, heap.EstaVacia())
}

func TestInvarienteColaPrioridad(t *testing.T) {
	t.Log("Se comprueba con pocos elementos que el funcionamiento de la cola de prioridad sea correcto, es decir" +
		"que se cumpla que lo que va saliendo sea lo que tenga mayor prioridad")
	heap := TDAColaPrioridad.CrearHeap[int](cmp) //heap de maximos
	heap.Encolar(4)
	require.EqualValues(t, 4, heap.VerMax())
	heap.Encolar(776)
	require.EqualValues(t, 776, heap.VerMax())
	heap.Encolar(-99)
	require.EqualValues(t, 776, heap.VerMax())
	heap.Encolar(7)
	require.EqualValues(t, 776, heap.VerMax())
	heap.Encolar(1000)
	require.EqualValues(t, 1000, heap.VerMax())

	orden := []int{1000, 776, 7, 4, -99}

	for i := 0; i < len(orden); i++ {
		require.False(t, heap.EstaVacia())
		require.EqualValues(t, 5-i, heap.Cantidad())
		require.EqualValues(t, orden[i], heap.VerMax())
		require.EqualValues(t, orden[i], heap.Desencolar())
	}
	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())

}

func TestVaciarColaPrioridad(t *testing.T) {
	t.Log("Se aniaden varios elementos a la cola, luego se la vacia, se comprueba que se comporto como una cola de prioridad" +
		"vacia y se le pueda volver a aniadir elemento sin problemas")
	heap := TDAColaPrioridad.CrearHeap[int](cmp) //heap maximos
	heap.Encolar(1)
	require.EqualValues(t, 1, heap.VerMax())
	heap.Encolar(123)
	require.EqualValues(t, 123, heap.VerMax())
	heap.Encolar(3)
	require.EqualValues(t, 123, heap.VerMax())
	heap.Encolar(500)
	require.EqualValues(t, 500, heap.VerMax())
	heap.Encolar(5)
	require.EqualValues(t, 500, heap.VerMax())
	orden := []int{500, 123, 5, 3, 1}
	i := 0
	for !heap.EstaVacia() {
		require.EqualValues(t, 5-i, heap.Cantidad())
		require.EqualValues(t, orden[i], heap.VerMax())
		require.EqualValues(t, orden[i], heap.Desencolar())
		i++
	}
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })

	heap.Encolar(50)
	heap.Encolar(0)

	require.EqualValues(t, 2, heap.Cantidad())
	require.EqualValues(t, 50, heap.VerMax())
	require.EqualValues(t, 50, heap.Desencolar())
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, 0, heap.VerMax())
	require.EqualValues(t, 0, heap.Desencolar())
	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
}

func TestHeapDeMinimos(t *testing.T) {
	t.Log("Se comprueba que la cola de prioridaad funcione tambien para minimos al cambiar la funcion de comparacion")
	heap := TDAColaPrioridad.CrearHeap[int](func(a, b int) int { return b - a })
	heap.Encolar(77)
	require.EqualValues(t, 77, heap.VerMax())
	heap.Encolar(49)
	require.EqualValues(t, 49, heap.VerMax())
	heap.Encolar(0)
	require.EqualValues(t, 0, heap.VerMax())
	heap.Encolar(12)
	require.EqualValues(t, 0, heap.VerMax())
	orden_minimo := []int{0, 12, 49, 77}
	for i := 0; i < len(orden_minimo); i++ {
		require.False(t, heap.EstaVacia())
		require.EqualValues(t, 4-i, heap.Cantidad())
		require.EqualValues(t, orden_minimo[i], heap.VerMax())
		require.EqualValues(t, orden_minimo[i], heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())

}

func TestVolumen(t *testing.T) {
	t.Log("Se prueba que la cola de prioridad funciona correctamente encolando y desencolando muchos elementos")
	heap := TDAColaPrioridad.CrearHeap[int](cmp) //heap de maximos
	for i := 0; i <= 1000000; i++ {
		heap.Encolar(i)
		require.EqualValues(t, i+1, heap.Cantidad())
		require.EqualValues(t, i, heap.VerMax())
	}

	for i := 1000000; !heap.EstaVacia(); i-- {
		require.EqualValues(t, i+1, heap.Cantidad())
		require.EqualValues(t, i, heap.VerMax())
		require.EqualValues(t, i, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
}

func TestHeapArr(t *testing.T) {
	t.Log("Se crear un heap desde un arreglo vacio y se comprueba que el heap se comporta como tal")
	arr := []int{}
	heap := TDAColaPrioridad.CrearHeapArr[int](arr, cmp)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	heap.Encolar(2)
}

func TestHeapArrPocosElementos(t *testing.T) {
	t.Log("Se crear un heap desde un arreglo no vacio de pocos elementos y se comprueba que al desencolar los elementos salen" +
		"Con la prioridad impuesta")
	arr := []int{1, 5, 66, 2, 3, 0}
	orden := []int{66, 5, 3, 2, 1, 0}
	heap := TDAColaPrioridad.CrearHeapArr[int](arr, cmp)
	require.False(t, heap.EstaVacia())
	for i := 0; i < len(orden); i++ {
		require.EqualValues(t, 6-i, heap.Cantidad())
		require.EqualValues(t, orden[i], heap.VerMax())
		require.EqualValues(t, orden[i], heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())

}

func TestHeapSort(t *testing.T) {
	t.Log("Prueba que el ordenamiento HeapSort funcione adecuadamente")
	arr := []int{10, 15, 7, 3, 21}
	arr_ord := []int{3, 7, 10, 15, 21}
	TDAColaPrioridad.HeapSort[int](arr, cmp)
	for i := 0; i < len(arr); i++ {
		require.EqualValues(t, arr_ord[i], arr[i])
	}
}

func TestHeapSortMayorAMenor(t *testing.T) {
	t.Log("Prueba que Heapsort funcione para ordenar un arreglo de mayor a menor")
	arr := []int{2, 11, 200, 55, 6, 9}
	arr_ord := []int{200, 55, 11, 9, 6, 2}
	TDAColaPrioridad.HeapSort[int](arr, func(a, b int) int { return b - a })
	for i := 0; i < len(arr); i++ {
		require.EqualValues(t, arr_ord[i], arr[i])
	}
}

func TestHeapConStrings(t *testing.T) {
	t.Log("Prueba que creado un heap a partir de un array de strings funcione adecuadamente")
	arr := []string{"Arbol", "Dario", "Boca", "Kennedy", "Fluminense", "Efemeride", "Club", "Alto", "Septima"}
	heap := TDAColaPrioridad.CrearHeapArr[string](arr, strings.Compare)

	arr_ord := []string{"Septima", "Kennedy", "Fluminense", "Efemeride", "Dario", "Club", "Boca", "Arbol", "Alto"}

	for i := 0; i < len(arr); i++ {
		require.EqualValues(t, arr_ord[i], heap.Desencolar())
	}
}

func TestHeapArrVolumen(t *testing.T) {
	t.Log("Comprueba que un heap apartir de un arreglo de muchos elementos se comporte adecuaddamente")
	arr := []int{}
	for i := 0; i < 100000; i++ {
		arr = append(arr, i)
	}
	heap := TDAColaPrioridad.CrearHeapArr[int](arr, cmp)

	for i := 99999; i >= 0; i-- {
		require.False(t, heap.EstaVacia())
		require.EqualValues(t, i+1, heap.Cantidad())
		require.EqualValues(t, i, heap.VerMax())
		require.EqualValues(t, i, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
}

func TestHeapSortVolumen(t *testing.T) {
	t.Log("Prueba que HepSort funcione adecuadamente con muchos elementos")
	arr := []int{}

	for i := 0; i < 100000; i++ {
		arr = append(arr, rand.Intn(1939924))
	}
	num_anterior := -1
	TDAColaPrioridad.HeapSort[int](arr, cmp)
	for i := 0; i < 100000; i++ {
		require.True(t, num_anterior <= arr[i])
		num_anterior = arr[i]
	}
}
