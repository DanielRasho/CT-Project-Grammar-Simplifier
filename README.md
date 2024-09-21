# Lab 7 - Gramatica 🌟

Este proyecto busca simplificar y verificar gramaticas libres de contexto

[🔴VIDEO DE EJECUCION AQUI](https://youtu.be/TqcEflv9wao)

## 📝 Entrada
El programa recibe la ruta de un archivo con un listado de producciones, con un formato como el siguiente
```
A -> bA|B|i
B -> m|ε
C -> ?|!
```
- **Nota:** La cadena vacía `ε` se representará como `ε` en este proyecto.
- **Nota:** Los " " entre producciones serán tomados como cualquier caracter.

## 📤 Salida

- **Verificacion:**
  El programa verificara, si la gramatica se encuentra bien escrita.

- **Simplificacion de gramatica:**
  Si la gramatica esta bien expresada, el programa se encargara de remover producciones-ε mostrando el proceso paso a paso.


## 🔗 Recursos y Referencias

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

## 💻 Ejecución del Proyecto
Una vez que tengas Nix instalado, puedes ejecutar el resto de los ejercicios de este laboratorio.

Los siguientes comandos crearán un entorno shell con todas las dependencias necesarias para ejecutar el proyecto, de manera similar a lo que hace Docker.

### 1. Árbol de Sintaxis Abstracta (AST)

Este comando generará una serie de imágenes en el directorio `./graphs` representando el AST de cada expresión regular.

```bash
nix run .#grammar --experimental-features 'nix-command flakes'
```

## 🛠️ Troubleshooting

Dependiendo de la shell que estés usando para ejecutar Nix, podrías necesitar ajustar el comando mostrado anteriormente. Algunas variantes incluyen:

```bash
nix run .\#grammar --experimental-features 'nix-command flakes'
nix run '.#grammar' --experimental-features 'nix-command flakes'
```