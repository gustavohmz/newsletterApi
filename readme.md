# Newsletter App API

Esta API proporciona servicios para la gestión de boletines informativos y suscriptores.

## Ejecución con Docker

Este proyecto incluye un archivo Dockerfile para facilitar la creación de un contenedor.

```bash
docker build -t newsletter-app .
docker run -p 8080:8080 --env-file .env newsletter-app
```

## Configuración del Entorno

Este proyecto utiliza variables de entorno para su configuración. Asegúrate de definir las siguientes variables en tu entorno de ejecución o archivo `.env`:

- `mongoUrl`: URL de conexión a MongoDB.
- `mongoDb`: Nombre de la base de datos MongoDB.
- `mongoNewsletterCollection`: Nombre de la colección de boletines en MongoDB.
- `mongoSubscriberCollection`: Nombre de la colección de suscriptores en MongoDB.
- `emailSender`: Dirección de correo electrónico para el envío de boletines.
- `emailPass`: Contraseña del correo electrónico para el envío de boletines.
- `smtpServer`: Servidor SMTP para el envío de correos electrónicos.
- `smtpPort`: Puerto SMTP para el envío de correos electrónicos.

## Funcionalidades

### Boletines (Newsletters)

#### Obtener lista de boletines

- **Método:** GET
- **Ruta:** `/api/v1/newsletters`
- **Descripción:** Obtiene una lista de boletines con parámetros opcionales de búsqueda y paginación.

  **Parámetros:**

  - `name` (string, query): Nombre del boletín a buscar.
  - `page` (integer, query): Número de página para la paginación.
  - `pageSize` (integer, query): Número de elementos por página para la paginación.

  **Respuestas:**

  - Código 200 (OK)
  - Código 400 (Bad Request)
  - Código 500 (Internal Server Error)

#### Actualizar un boletín existente

- **Método:** PUT
- **Ruta:** `/api/v1/newsletters`
- **Descripción:** Permite a un usuario administrador actualizar un boletín existente.

  **Parámetros:**

  - `updateRequest` (objeto, body): Detalles actualizados del boletín.

  **Respuestas:**

  - Código 200 (OK)
  - Código 400 (Bad Request)
  - Código 500 (Internal Server Error)

#### Crear un nuevo boletín

- **Método:** POST
- **Ruta:** `/api/v1/newsletters`
- **Descripción:** Permite a un usuario administrador crear un nuevo boletín.

  **Parámetros:**

  - `newsletter` (objeto, body): Detalles del nuevo boletín.

  **Respuestas:**

  - Código 201 (Created)
  - Código 400 (Bad Request)
  - Código 500 (Internal Server Error)

#### Enviar boletín a suscriptores

- **Método:** POST
- **Ruta:** `/api/v1/newsletters/send/{newsletterID}`
- **Descripción:** Permite a un usuario administrador enviar un boletín a una lista de suscriptores.

  **Parámetros:**

  - `newsletterID` (string, path): ID del boletín a enviar.

  **Respuestas:**

  - Código 200 (OK)
  - Código 400 (Bad Request)
  - Código 500 (Internal Server Error)

#### Eliminar un boletín

- **Método:** DELETE
- **Ruta:** `/api/v1/newsletters/{id}`
- **Descripción:** Permite a un usuario administrador eliminar un boletín.

  **Parámetros:**

  - `id` (string, path): ID del boletín a eliminar.

  **Respuestas:**

  - Código 200 (OK)
  - Código 400 (Bad Request)
  - Código 500 (Internal Server Error)

### Suscriptores (Subscribers)

#### Suscribirse al boletín

- **Método:** POST
- **Ruta:** `/api/v1/subscribe/{email}/{category}`
- **Descripción:** Permite a un usuario suscribirse al boletín.

  **Parámetros:**

  - `email` (string, path): Dirección de correo electrónico para la suscripción.
  - `category` (string, path): Categoría a la que suscribirse.

  **Respuestas:**

  - Código 200 (OK)
  - Código 400 (Bad Request)
  - Código 500 (Internal Server Error)

#### Obtener lista de suscriptores

- **Método:** GET
- **Ruta:** `/api/v1/subscribers`
- **Descripción:** Obtiene una lista de suscriptores con parámetros opcionales de búsqueda y paginación.

  **Parámetros:**

  - `email` (string, query): Dirección de correo electrónico del suscriptor a buscar.
  - `category` (string, query): Categoría del suscriptor a buscar.
  - `page` (integer, query): Número de página para la paginación.
  - `pageSize` (integer, query): Número de elementos por página para la paginación.

  **Respuestas:**

  - Código 400 (Bad Request)
  - Código 500 (Internal Server Error)

#### Obtener suscriptor por correo electrónico y categoría

- **Método:** GET
- **Ruta:** `/api/v1/subscribers/{email}/{category}`
- **Descripción:** Obtiene detalles de un suscriptor por dirección de correo electrónico.

  **Parámetros:**

  - `email` (string, path): Dirección de correo electrónico para obtener detalles.
  - `category` (string, path): Categoría a la que está suscrito.

  **Respuestas:**

  - Código 200 (OK)
  - Código 404 (Subscriber not found)
  - Código 500 (Internal Server Error)

#### Cancelar suscripción al boletín

- **Método:** DELETE
- **Ruta:** `/api/v1/unsubscribe/{email}/{category}`
- **Descripción:** Permite a un usuario cancelar la suscripción al boletín.

  **Parámetros:**

  - `email` (string, path): Dirección de correo electrónico para cancelar la suscripción.
  - `category` (string, path): Categoría a la que está suscrito.

  **Respuestas:**

  - Código 200 (OK)
  - Código 400 (Bad Request)
  - Código 500 (Internal Server Error)
