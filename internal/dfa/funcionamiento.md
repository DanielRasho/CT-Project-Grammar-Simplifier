# Proceso de Conversión de AFN a AFD

## Primer Paso: Construcción de la Tabla de Transiciones

1. **Estados y Transiciones:**
   - Identifica todos los estados y transiciones del AFN.
   - Para cada estado, determina a qué otros estados puede llegar mediante cada transición, incluyendo posibles transiciones vacías (ε-transiciones).

2. **Transiciones con ε:**
   - Para cada estado, calcula el conjunto de estados a los que se puede llegar utilizando exclusivamente ε-transiciones (denominado ε-cierre). Recuerda que un estado siempre puede llegar a sí mismo mediante ε, por lo que este conjunto nunca estará vacío.

3. **Identificación de Estados Inicial y Final:**
   - El estado inicial del AFN es el primer estado en la lista.
   - El estado final es aquel que está designado como tal en el AFN original.

**Ejemplo de Tabla de Transiciones del AFN:**

| **Estado** | **Transición 1** | **Transición N** | **ε-cierre**  |
|------------|------------------|------------------|---------------|
| Estado1    | {Estado2}        | {Estado3}        | {Estado1, ...}|
| Estado2    | {Estado3}        | {}               | {Estado2, ...}|
| ...        | ...              | ...              | ...           |
| EstadoN    | {}               | {Estado1}        | {EstadoN, ...}|

## Segundo Paso: Construcción del AFD

1. **Inicio con el Conjunto del Estado Inicial:**
   - Comienza con el ε-cierre del estado inicial del AFN y asígnale una etiqueta (por ejemplo, "A").

2. **Generación de Nuevos Estados:**
   - Para cada transición:
     - Examina todos los estados en el conjunto actual y determina a qué estados se puede llegar con esa transición.
     - Calcula el ε-cierre para cada estado alcanzado y combina estos resultados para formar un nuevo conjunto de estados.
   - Si el conjunto de estados generado no ha sido etiquetado previamente, asígnale una nueva etiqueta y añádelo a la lista de estados del AFD.

3. **Determinación de Estados Vacíos:**
   - Si, al procesar una transición, no se alcanzan nuevos estados (el conjunto está vacío), crea un nuevo estado en el AFD que representará esta situación. Este estado vacío permanecerá como un conjunto fijo y no se desarrollará más.

**Ejemplo de Tabla de Transiciones del AFD:**

| **Estados del AFN** | **Transición 1** | **Transición 2** | **Transición N** |
|---------------------|------------------|------------------|------------------|
| {Estado1, ...} = A  | {Estado2, ...} = B | {Estado3, ...} = C | ... |
| {Estado2, ...} = B  | {Estado4, ...} = D | {} = E          | ... |
| ...                 | ...              | ...              | ...              |

## Tercer Paso: Diagramación del AFD

1. **Creación del Diagrama:**
   - Utiliza la tabla resultante para construir el diagrama del AFD.
   - El estado inicial del AFD será el correspondiente al conjunto de estados iniciales del AFN.
   - Los estados de aceptación en el AFD serán aquellos cuyos conjuntos contienen el estado de aceptación original del AFN.

2. **Conclusión:**
   - El AFD final tendrá estados claramente etiquetados (A, B, C, ...) que representan conjuntos de estados del AFN original.
   - Las transiciones en el AFD indicarán a qué estado etiquetado se llega a partir de un estado dado y una entrada específica.
