# Desarrollo de una API REST para gestión de productos

La empresa necesita una API REST para gestionar productos en su plataforma de e-commerce. La API debe permitir crear, leer, actualizar y eliminar productos. Los productos tienen atributos como nombre, precio, stock y categoría. La API debe validar que los nombres de los productos no sean duplicados y que los precios no sean negativos. Además, debe manejar adecuadamente los errores y proporcionar respuestas coherentes al cliente.

## Informacion General

| Campo | Valor |
|-------|-------|
| **Tema** | Go Gin API |
| **Nivel** | junior-l1 |
| **Tipo** | practical |
| **Tiempo estimado** | 8 horas |

## Fases del Reto

### Fase 0: Configuración del Proyecto

**Objetivo:** Obtener el proyecto base funcional enviando el Código Base a un asistente de IA, que lo analizará, corregirá errores y generará un ZIP listo para usar.

**Tiempo estimado:** 15-30 minutos

**Instrucciones:**

- Asegúrate de tener instalado para ejecutar el proyecto: Go 1.21+, VS Code o GoLand.
- Copia todo el contenido del campo **Código Base** de este reto — incluyendo el texto de instrucciones que aparece al inicio.
- Abre un asistente de IA (Claude en claude.ai, ChatGPT o Gemini — se recomienda Claude), pega el contenido copiado en el chat y envíalo.
- El asistente analizará los archivos, corregirá errores y generará un archivo ZIP descargable. Descárgalo y extráelo en la carpeta donde quieras trabajar.
- Ejecuta `go build ./...`. Si no hay errores, estás listo.

**Entregable:** El proyecto compila/arranca sin errores.

<details>
<summary>Pistas de conocimiento</summary>

- Copia el Código Base completo incluyendo el texto de instrucciones al inicio — esas instrucciones le indican al asistente exactamente qué hacer con los archivos.
- Si el asistente no genera el ZIP automáticamente al terminar el análisis, escríbele: "genera el ZIP ahora".
- Si el proyecto tiene errores al arrancar, comparte el mensaje de error con el mismo asistente para que lo corrija.

</details>

### Fase 1: Creación de la estructura básica de la API

**Objetivo:** Definir y crear la estructura básica de la API, incluyendo la configuración del enrutamiento y la conexión a la base de datos.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- Configurar las rutas básicas para crear, leer, actualizar y eliminar productos.
- Conectar la API a una base de datos para persistir los productos.
- Asegurar que la API devuelva respuestas adecuadas para cada operación.

**Entregable:** API con rutas básicas y conexión a la base de datos.

<details>
<summary>Pistas de conocimiento</summary>

- Considera cómo estructurar la API para que sea escalable y mantenible.
- Piensa en cómo manejarás las validaciones de los productos.

</details>

### Fase 2: Implementación de validaciones y manejo de errores

**Objetivo:** Implementar las validaciones necesarias para los productos y manejar adecuadamente los errores que puedan ocurrir durante las operaciones.

**Tiempo estimado:** 3 horas

**Instrucciones:**

- Validar que los nombres de los productos no sean duplicados.
- Validar que los precios de los productos no sean negativos.
- Manejar los errores que puedan ocurrir durante las operaciones y proporcionar respuestas coherentes al cliente.

**Entregable:** API con validaciones y manejo de errores implementados.

<details>
<summary>Pistas de conocimiento</summary>

- Considera cómo puedes mejorar la experiencia del usuario al proporcionar mensajes de error claros y útiles.
- Piensa en cómo puedes hacer que las validaciones sean eficientes y no afecten el rendimiento de la API.

</details>

### Fase 3: Optimización y refactorización de la API

**Objetivo:** Optimizar y refactorizar la API para mejorar su rendimiento y mantenibilidad.

**Tiempo estimado:** 3 horas

**Instrucciones:**

- Identificar y corregir cualquier ineficiencia en la API.
- Refactorizar el código para mejorar su legibilidad y mantenibilidad.
- Asegurar que la API cumpla con las mejores prácticas de desarrollo.

**Entregable:** API optimizada y refactorizada.

<details>
<summary>Pistas de conocimiento</summary>

- Considera cómo puedes mejorar el rendimiento de la API sin sacrificar su funcionalidad.
- Piensa en cómo puedes hacer que el código sea más legible y fácil de mantener.

</details>

## Dimensiones Evaluadas

- **queEs**: ¿Qué es una API REST y cuál es su propósito en este contexto?
- **paraQueSirve**: ¿Para qué sirve la validación de nombres y precios de los productos en la API?
- **comoSeUsa**: ¿Cómo se usa la API para crear, leer, actualizar y eliminar productos?
- **erroresComunes**: ¿Cuáles son los errores comunes que pueden ocurrir durante las operaciones de la API y cómo se manejan?
- **queDecisionesImplica**: ¿Qué decisiones implica la optimización y refactorización de la API?

## Criterios de Evaluacion

- Implementar las rutas básicas de la API y conectarla a una base de datos.
- Validar que los nombres de los productos no sean duplicados y que los precios no sean negativos.
- Manejar adecuadamente los errores y proporcionar respuestas coherentes al cliente.
- Optimizar y refactorizar la API para mejorar su rendimiento y mantenibilidad.

## Como trabajar con un asistente de IA

Hay dos caminos, elegi uno:

- **AGENTS.md** (recomendado) — instrucciones nativas del repo. Abri esta carpeta con tu agente local (Claude Code, Cursor, Codex, Copilot, Gemini) y las carga solo. Sabe que archivos faltan y con que comando se verifica, y completa el scaffold escribiendo en disco.
- **PROMPT_MEJORA.md** — para copiar y pegar en un chat (claude.ai, ChatGPT). Devuelve un ZIP con el proyecto. Sirve si no tenes un agente en el IDE.

Ninguno de los dos resuelve las fases del reto: eso es tu trabajo.

## Verificacion

El proyecto esta listo para trabajar cuando este comando corre sin errores:

```bash
el comando de build o arranque canonico del stack elegido
```

---

*Reto generado automaticamente por Challenge Generator - Pragma*
