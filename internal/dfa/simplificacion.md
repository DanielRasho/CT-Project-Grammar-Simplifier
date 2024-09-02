# Algoritmo de minimización

1. **Partición Inicial**:
   - Divide los estados del AFD en dos grupos: estados de aceptación y estados no de aceptación.

2. **Refinamiento de Grupos**:
   - Para cada símbolo de entrada, divide los grupos actuales en subgrupos, basándote en las transiciones de los estados bajo el símbolo.
   - Si dos estados pertenecen al mismo grupo, pero se dirigen a diferentes grupos en la transición, se dividen en subgrupos separados.

3. **Repetición**:
   - Repite el proceso hasta que no sea posible dividir más los grupos.
   - El resultado será un conjunto de grupos donde cada grupo representa un estado minimizado del AFD.
