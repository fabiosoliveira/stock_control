# Gerenciador de Estoque

Um sistema de gerenciamento de estoque desenvolvido em Go, utilizando tecnologias modernas e práticas de desenvolvimento de software.

## Tecnologias Utilizadas

- **Go**: Linguagem de programação utilizada para desenvolver o sistema.
- **SQLite**: Banco de dados relacional utilizado para armazenar os dados do estoque.
- **HTML/CSS**: Utilizados para criar a interface do usuário e estilizar a aplicação.
- **JavaScript**: Utilizado para criar interações dinâmicas na interface do usuário.
- **Template Engine**: Utilizado para renderizar templates HTML com dados dinâmicos.
- **Middleware**: Utilizado para implementar funcionalidades de cache e autenticação.
- **API**: Utilizada para criar endpoints para criar, ler, atualizar e excluir produtos.

## Funcionalidades

- **Cadastro de Produtos**: Permite criar, ler, atualizar e excluir produtos.
- **Listagem de Produtos**: Exibe uma lista de produtos com suas respectivas informações.
- **Modal de Confirmação**: Exibe um modal de confirmação antes de excluir um produto.
- **Toast de Notificação**: Exibe uma notificação toast ao criar, atualizar ou excluir um produto.

## Estrutura do Projeto

- **internal**: Pasta que contém os arquivos de código Go.
  - **config**: Pasta que contém os arquivos de configuração do banco de dados.
  - **controller**: Pasta que contém os arquivos de controle do sistema.
  - **middleware**: Pasta que contém os arquivos de middleware.
  - **product**: Pasta que contém os arquivos de modelo de produto.
  - **repository**: Pasta que contém os arquivos de repositório de produtos.
- **internal/template**: Pasta que contém os arquivos de template HTML.
- **cmd**: Pasta que contém os arquivos de comando do sistema.

## Como Executar o Projeto

1. Clone o repositório.
2. Execute o comando `go build` para compilar o projeto.
3. Execute o comando `go run` para executar o projeto.
4. Acesse o sistema através do navegador em [http://localhost:8080](http://localhost:8080).

## Contribuições

Contribuições são bem-vindas! Se você tiver alguma sugestão ou correção, por favor, abra uma issue ou envie um pull request.
