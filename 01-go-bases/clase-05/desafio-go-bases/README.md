# desafio-go-bases

Realizar un programa que sirva como herramienta para calcular diferentes datos estadísticos. Para lograrlo, debes clonar 
este repositorio que contiene un archivo .csv con datos generados y un esqueleto del proyecto.


¡Atención! Los ejemplos a continuación son sólo de guía, el desafío se puede resolver de múltiples maneras.

Requerimiento 1:
Una función que calcule cuántas personas viajan a un país determinado:
func GetTotalTickets(destination string) (int, error) {}

(ejemplo 1)
Tip: VS Code nos permite buscar una palabra en un archivo con Ctrl + F o ⌘ + F.

Requerimiento 2:
Una o varias funciones que calcule cuántas personas viajan en madrugada (0 → 6), mañana (7 → 12), tarde (13 → 19) 
y noche (20 → 23):
func GetCountByPeriod(time string) (int, error) {}

(ejemplo 2)
Tip: En Go para manipular caracteres tenemos el paquete strings.

Requerimiento 3:
Calcular el promedio de personas que viajan a un país determinado en un dia:
func AverageDestination(destination string, total int) (float64, error) {}

(ejemplo 3)
Tip: El promedio de x se calcula como: x̄ =  xn

Requerimiento 4:
Crear test unitarios para cada uno de los requerimientos anteriores, mínimo 2 casos por requerimiento:
import "testing"
func TestGetTotalTickets(t *testing.T) {}

(ejemplo 4)

Tip: Podemos testear múltiples casos en un mismo test.
