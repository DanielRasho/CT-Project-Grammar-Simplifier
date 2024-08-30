# Implementación de Algoritmos Básicos de Autómatas Finitos y Expresiones Regulares 🌟

Este proyecto tiene como objetivo la implementación de algoritmos fundamentales para trabajar con autómatas finitos y expresiones regulares, incluyendo la generación de autómatas, su simulación, y la verificación de pertenencia de cadenas al lenguaje descrito por una expresión regular.

## 📝 Entrada

- **Expresión Regular** `r`  
  - **Ejemplo:** `r = (b|b)* abb(a|b)*`

- **Cadena** `w`  
  - **Ejemplo:** `w = babbaaaaa`

- **Nota:** La cadena vacía `ε` se representará como `ε` en este proyecto.

## 📤 Salida

- **Generación de Imágenes:**
  - **Grafo del AFN:** Autómata Finito No Determinista.
  - **Grafo del AFD:** Autómata Finito Determinista, construido mediante el método de subconjuntos y su minimización.

- **Simulación:**
  - **Simulación AFN y AFD:** El programa debe indicar si `w ∈ L(r)` con un **"Sí"** si la cadena pertenece al lenguaje de la expresión regular, o **"No"** si no pertenece.

- **Lectura de Archivos:**
  - El programa también puede leer un archivo de texto, procesando cada línea como una entrada y aplicando las opciones descritas anteriormente.

## ⚙️ Especificaciones Técnicas

1. **Algoritmo Shunting Yard:** Conversión de expresiones de infix a postfix.
2. **Algoritmo de Construcción de Thompson:** Creación de AFN a partir de expresiones regulares.
3. **Algoritmo de Construcción de Subconjuntos:** Conversión de AFN a AFD.
4. **Algoritmo de Minimización de AFD:** Reducción de estados en el AFD.
5. **Simulación de AFN:** Verificación de pertenencia de la cadena `w` en el AFN.
6. **Simulación de AFD:** Verificación de pertenencia de la cadena `w` en el AFD.

## 🔗 Recursos y Referencias

- [Shunting Yard](https://www.youtube.com/watch?v=j5_cEkciqSc)
- [Balanceo de Expresiones](https://www.youtube.com/watch?v=jzJVkGRze2Y)
- [Algoritmo de Thompson para AST](https://youtu.be/UMoSHemFSx0)
- [Creación de AFN](https://youtu.be/VYDXzB57Of8)

## 💻 Ejecución del Proyecto

### 1. Árbol de Sintaxis Abstracta (AST)

Este comando generará una serie de imágenes en el directorio `./graphs` representando el AST de cada expresión regular.

```bash
nix run .#ast --experimental-features 'nix-command flakes'
```

### 2. Autómata Finito No Determinista (AFN)

Este comando ejecuta la construcción del AFN.

```bash
nix run .#afn --experimental-features 'nix-command flakes'
```

### 3. Balanceo de Expresiones

```bash
nix run .#balancer --experimental-features 'nix-command flakes'
```

### 4. Algoritmo Shunting Yard

```bash
nix run .#shuntingyard --experimental-features 'nix-command flakes'
```

### Nota 🗒️

Todos los grafos e imágenes generados, se guardarán automáticamente en la carpeta `./graphs`.

## 🚀 Getting Started

### Instalación

Para ejecutar el proyecto de manera sencilla, solo necesitas tener el gestor de paquetes [Nix](https://nixos.org/download/#nix-install-linux) instalado en tu sistema. Puedes hacerlo ejecutando el siguiente comando:

**Linux & Windows**

```bash
$ sudo sh <(curl -L https://nixos.org/nix/install) --daemon
```

**MacOS**

```bash
$ sh <(curl -L https://nixos.org/nix/install)
```

### Ejecución

Una vez que tengas Nix instalado, puedes ejecutar el resto de los ejercicios de este laboratorio.

Los siguientes comandos crearán un entorno shell con todas las dependencias necesarias para ejecutar el proyecto, de manera similar a lo que hace Docker.

## 🛠️ Troubleshooting

Dependiendo de la shell que estés usando para ejecutar Nix, podrías necesitar ajustar el comando mostrado anteriormente. Algunas variantes incluyen:

```bash
nix run .\#project --experimental-features 'nix-command flakes'
nix run '.#project' --experimental-features 'nix-command flakes'
```