# Eco Medic ( Extensão API )

Este repositório contém a implementação do backend do sistema **Eco Medic**, desenvolvido como parte do trabalho da disciplina **Introdução a Sistemas de Informação** na PUC Minas. O sistema tem como objetivo fornecer uma API para autenticação de usuários, gerenciamento de descartes e recompensas de medicamentos.

## 🚀 Tecnologias Utilizadas
- **Linguagem:** Go (Golang)
- **Banco de Dados:** MySQL
- **Docker**: Para orquestração dos serviços.
- **Frameworks e Bibliotecas:** 
  - net/http
  - jwt-go
  - gorm (ORM)
- **Ferramentas:** Postman e Dbeaver

## 📽 Como Executar o Projeto

### Pré-requisitos
- Instale o [Go](https://go.dev/doc/install) (versão 1.23 ou superior).
- Docker e Docker Compose instalados.

## 🛠️ Configuração do Ambiente

### Configuração
1. Clone este repositório.
2. Ajuste variáveis de ambiente no `docker-compose.yml` se necessário.
3. Suba os serviços com:
   ```bash
   docker-compose up -d
   ```
4. Acesse o banco de dados via phpMyAdmin em http://localhost:8080 , ou pode também utilizar um outro SGBD para poder acessar o banco de dados, passando as credenciais. No projeto, tive como uso o Dbeaver e recomendo. Motivo de implementar o phpMyAdmin foi para outros integrantes que fizeram o front-end, poderem acessar o banco de dados para melhor integração.

## 🖥️ Endpoints da API

### Auth
1. **Registrar Usuário**
2. **Login**


### API
1. **Detalhes Usuários**
2. **Criar Descarte**
3. **Consultar Recompensas**


### 📄 Tabela de Endpoints

| **Endpoint**              | **Método** | **Descrição**                        | **Body de Exemplo**                                                                                                                                                   | **Headers Necessários**                  |
|---------------------------|------------|--------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------|
| `/auth/register`          | POST       | Registrar um novo usuário           | ```json { "name": "Pedro", "email": "pedro@gmail.com", "password": "senha123", "birth_date": "2002-12-13", "phone": "(31) 9 92051504" } ```                            | `Content-Type: application/json`         |
| `/auth/login`             | POST       | Realizar login                      | ```json { "email": "eduardoatenesilva@gmail.com", "password": "senha123" } ```                                                                                       | `Content-Type: application/json`         |
| `/api/user`            | GET           | Consultar detalhes do usuário logado   | N/A                                                                                                                                                                 | `Authorization: Bearer <seu_token>`      |
| `/api/discard`            | POST       | Criar um registro de descarte       | ```json { "city": "São Paulo", "location": "Av. Paulista, 1000", "date": "2024-11-25T10:00:00Z" } ```                                                                | `Authorization: Bearer <seu_token>`      |
| `/api/rewards`            | GET        | Consultar recompensas do usuário    | N/A                                                                                                                                                                 | `Authorization: Bearer <seu_token>`      |

### 🛠️ Exemplos de Uso

#### **1. Registrar Usuário**
**Requisição**:
```bash
curl -X POST http://localhost:3000/auth/register \
-H "Content-Type: application/json" \
-d '{
    "name": "Eduardo",
    "email": "eduardoatenesilvamarinha@gmail.com",
    "password": "senha123",
    "birth_date": "2002-12-13",
    "phone": "(31) 9 999999999"
}'
```

#### **2. Login**
**Requisição**:
```bash
curl -X POST http://localhost:3000/auth/login \
-H "Content-Type: application/json" \
-d '{
    "email": "eduardoatenesilvamarinha@gmail.com",
    "password": "senha123"
}'
```
#### **3. Get Detalhes Usuário**
**Requisição**:
```bash
curl --location 'http://localhost:3000/api/user' \
--header 'Authorization: ••••••'
```
#### **4. Criar Descarte**
**Requisição**:
```bash
curl -X POST http://localhost:3000/api/discard \
-H "Authorization: Bearer <seu_token>" \
-H "Content-Type: application/json" \
-d '{
    "city": "São Paulo",
    "location": "Av. Paulista, 1000",
    "date": "2024-11-25T10:00:00Z"
}'
```
#### **5. Consultar Recompensas**
**Requisição**:
```bash
curl -X GET http://localhost:3000/api/rewards \
-H "Authorization: Bearer <seu_token>"
```
---

## 📁 Estrutura do Projeto/

- extensao-web-api/         # Código da API
- docker-compose.yml        # Orquestração com Docker ( Onde deve ficar seu docker-compose.yml. Pode-se mudar também )

## Diretórios Principais
- **/cmd/api/**: Ponto de entrada da aplicação, onde o servidor é iniciado e as rotas são configuradas.
- **/internal/**: Código interno da aplicação, organizado em pacotes com funcionalidades específicas.
- **/config/api/**: Configurações do sistema, como banco de dados, tokens e variáveis de ambiente.
- **/domain/**: Define os modelos e interfaces do domínio da aplicação.
- **/handlers/**: Implementa os manipuladores HTTP (controladores das rotas).
- **/middleware/**: Contém os middlewares usados para processamento intermediário das requisições.
- **/repository/**: Camada de acesso a dados, como consultas ao banco.
- **/usecase/**: Lógica de negócios, conectando os repositórios aos manipuladores.
- **/utils/**: Funções utilitárias usadas em diferentes partes do sistema.

## Arquivos Importantes
- **extensao_database.sql**: Script SQL para criação do banco de dados.
- **Eco_Medic_API.postman_collection.json**: Coleção de requisições Postman para testar os endpoints.

---

## 💡 Observações
- Certifique-se de configurar corretamente o `docker-compose.yml`, de acordo com o que foi comentado aqui no MK.
- A aplicação está configurada para rodar na porta `3000` e o phpMyAdmin na porta `8080`.

---

## 🤝 Contribuições
Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e pull requests.
