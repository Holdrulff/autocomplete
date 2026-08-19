# Autocomplete

Aplicação de autocomplete para tecnologias baseada nas tags mais populares do Stack Overflow.

O backend carrega um snapshot das tags em uma trie, busca termos por prefixo e devolve as sugestões ordenadas por popularidade. O frontend oferece busca com debounce, navegação por teclado e temas claro e escuro.

## Tecnologias

- Go
- GraphQL com gqlgen
- React, TypeScript e Vite
- Nginx
- Docker Compose

## Como executar

Pré-requisito: Docker com Docker Compose.

Na raiz do projeto, execute:

```bash
docker compose up --build
```

Acesse a aplicação em:

```text
http://localhost:8080
```

Para encerrar:

```bash
docker compose down
```

## API GraphQL

O endpoint está disponível em `POST /query`.

Exemplo:

```graphql
query {
  autocomplete(prefix: "java") {
    value
    score
  }
}
```

## Desenvolvimento

Execute os testes e as verificações do backend:

```bash
go test ./...
go vet ./...
```

Execute o frontend localmente:

```bash
npm --prefix ./frontend install
npm --prefix ./frontend run dev
```

Em outro terminal, inicie o backend:

```bash
go run ./cmd/server
```

O Vite disponibiliza o frontend em `http://localhost:5173` e encaminha `/query` para o backend local.

## Arquitetura

```text
React → Nginx → GraphQL → serviço de autocomplete → trie → snapshot de tags
```

Os dados foram obtidos pela Stack Exchange API. As contribuições do Stack Overflow são licenciadas sob CC BY-SA.
