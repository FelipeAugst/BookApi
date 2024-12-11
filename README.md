
# Book Api
Uma Api simples que raliza operações de crud em uma base de dados em memória.

## Instalação

Rodando o servidor(por padrão localmente na porta 5000)

```bash
  $git clone https://github.com/FelipeAugst/BookApi
  $ls BookApi
  $go run main.go
```
## Rotas

```bash
  /books GET- Retorna todos os livros

  /books/create POST- Cria um livros

  /books/{id}/delete DELETE - Deleta um livro

  /books/{id}/update PUT - Edita um livro

  /books/search?mode=&parameter=? GET -Busca um livro 
  por author ou title(definido na query mode)


```

