# Lab 7 - Gramatica 🌟

Este proyecto busca simplificar y verificar gramaticas libres de contexto

[🔴VIDEO DE EJECUCION AQUI](https://youtu.be/TqcEflv9wao)

[🔴VIDEO DE EXPLICACIÓN DE CODIGO](https://www.youtube.com/watch?v=RfeBYK0hnwU)

## 📝 Entrada
El programa recibe la ruta de un archivo con un listado de producciones, con un formato como el siguiente
```
A -> b{A}|{B}|i
B -> m|ε
C -> ?|!
```
- **Nota:** 
  - La cadena vacía `ε` se representará como `ε` en este proyecto.
  - Los " " entre producciones serán tomados como cualquier caracter.
  - Los No terminales deben escribirse dentro de llaves "{}", __por tanto las llaves no pueden formar parte del lenguaje__

## 📤 Salida

- **Simplificacion de gramatica:**
  Si la gramatica esta bien expresada, el programa se encargara de remover producciones-ε mostrando el proceso paso a paso.

- **Verificacion:**
  El programa verificara, si la gramatica se encuentra bien escrita usando algoritmo CYK.

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

### 💻 Ejecución del Proyecto
Una vez que tengas Nix instalado, puedes ejecutar el resto de los ejercicios de este laboratorio.

Los siguientes comandos crearán un entorno shell con todas las dependencias necesarias para ejecutar el proyecto, de manera similar a lo que hace Docker.

```bash
nix run .#grammar --experimental-features 'nix-command flakes'
```

## 🛠️ Troubleshooting

Dependiendo de la shell que estés usando para ejecutar Nix, podrías necesitar ajustar el comando mostrado anteriormente. Algunas variantes incluyen:

```bash
nix run .\#grammar --experimental-features 'nix-command flakes'
nix run '.#grammar' --experimental-features 'nix-command flakes'
```

# Diseño de la Apliación

# Discusión
El principal contratiempo en el desarrollo fue un mal diseño preliminar. Al principio se definio un agramatica como un diccionario donde las llaves eran los NO terminales y los cuerpos eran un lista de los cuerpos.

```go
type Grammar map[string][]string
```
Pero muy tarde nos dimos cuenta que el diseño tenia muchas falencias y tuvimos que definir un tipo custom para ello.
```go
type Symbol struct {
	IsTerminal bool
	Value      string
	Id         int
}
type Grammar struct {
	terminals    []Symbol              // List of all cached terminals in the grammar.
	NonTerminals []Symbol              // List of all cached NON terminals in the grammar.
	Productions  map[Symbol][][]Symbol // The actual productions.
}
```