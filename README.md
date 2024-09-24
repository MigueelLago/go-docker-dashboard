
# Docker Manager API

Esta é uma aplicação que gerencia containers Docker via API. A aplicação é construída com Go e utiliza o framework Gin para lidar com as rotas HTTP, além da biblioteca `docker/client` para interagir com a API Docker.

## Bibliotecas Utilizadas

- **Gin**: Framework web para Go. Usado para lidar com rotas e requisições HTTP.
  - Link: [Gin](https://github.com/gin-gonic/gin)
  
- **Docker Client**: Biblioteca oficial do Docker para interagir com a API Docker.
  - Link: [Docker Client](https://github.com/moby/moby/tree/master/client)

## Pré-requisitos

Antes de executar o projeto, você precisa ter:

- **Go**: Instale a versão mais recente do Go em [golang.org](https://golang.org/doc/install).
- **Docker**: Certifique-se de que o Docker esteja instalado e rodando. Instale o Docker em [docker.com](https://www.docker.com/get-started).

## Instalação

1. Clone o repositório:

```bash
git clone https://github.com/seu-usuario/docker-manager-api.git
cd docker-manager-api
```

2. Instale as dependências do projeto:

```bash
go mod tidy
```

Isso instalará as dependências listadas no arquivo `go.mod`.

## Configuração

Certifique-se de que o Docker esteja rodando no seu sistema, pois a API interage diretamente com a API Docker.

## Executando o Projeto

Para iniciar o servidor, execute o seguinte comando:

```bash
go run main.go
```

A aplicação irá rodar na porta `8080` por padrão. O servidor estará acessível em `http://localhost:8080`.

## Endpoints Disponíveis

### 1. Listar Containers

Lista todos os containers disponíveis ou apenas os que estão em execução, dependendo do parâmetro `only_running`.

- **Método**: `GET`
- **Rota**: `/api/containers`
- **Query Params**:
  - `only_running` (opcional): Se for `true`, retorna apenas os containers em execução.
- **Exemplo**:

  ```bash
  curl "http://localhost:8080/api/containers?only_running=true"
  ```

- **Resposta** (Exemplo):

  ```json
  {
    "containers": [
      {
        "id": "ab123cd456ef",
        "image": "nginx",
        "command": "/bin/bash",
        "created": 1695834345,
        "names": ["/example-container"],
        "status": "Up 5 minutes"
      }
    ]
  }
  ```

### 2. Buscar Containers por Imagem

Busca containers que utilizam uma determinada imagem.

- **Método**: `GET`
- **Rota**: `/api/containers/search`
- **Query Params**:
  - `image`: Nome da imagem Docker.
- **Exemplo**:

  ```bash
  curl "http://localhost:8080/api/containers/search?image=nginx"
  ```

- **Resposta** (Exemplo):

  ```json
  {
    "containers": [
      {
        "id": "ab123cd456ef",
        "image": "nginx",
        "command": "/bin/bash",
        "created": 1695834345,
        "names": ["/example-container"],
        "status": "Up 5 minutes"
      }
    ]
  }
  ```

### 3. Deletar Container

Deleta um container pelo seu ID.

- **Método**: `DELETE`
- **Rota**: `/api/containers/:id`
- **Parâmetros da URL**:
  - `id`: ID do container a ser deletado.
- **Exemplo**:

  ```bash
  curl -X DELETE "http://localhost:8080/api/containers/ab123cd456ef"
  ```

- **Resposta** (Exemplo):

  ```json
  {
    "message": "Container deleted successfully"
  }
  ```

### 4. Iniciar Container

Inicia um container pelo seu ID.

- **Método**: `POST`
- **Rota**: `/api/containers/:id/start`
- **Parâmetros da URL**:
  - `id`: ID do container a ser iniciado.
- **Exemplo**:

  ```bash
  curl -X POST "http://localhost:8080/api/containers/ab123cd456ef/start"
  ```

- **Resposta** (Exemplo):

  ```json
  {
    "message": "Container started successfully"
  }
  ```

## Estrutura do Projeto

```
.
├── Dockerfile
├── go.mod
├── go.sum
├── handlers
│   ├── container_handler.go
├── main.go
├── routes
│   ├── router.go
├── docker
│   ├── docker_usecase.go
```

### Descrição das Pastas:

- **handlers**: Contém os handlers responsáveis por receber as requisições HTTP e chamar os casos de uso apropriados.
- **routes**: Define as rotas da aplicação.
- **docker**: Contém os casos de uso responsáveis pela interação com a API Docker.

## Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou pull requests.

## Licença

Este projeto está licenciado sob a licença MIT. Para mais informações, consulte o arquivo `LICENSE`.
