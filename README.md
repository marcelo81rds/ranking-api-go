# 🏆 Ranking API (Go + PostgreSQL)

API RESTful desenvolvida em Go (Golang) para gerenciamento de ranking de jogadores, utilizando PostgreSQL em nuvem. Permite inserção de pontuações e consulta dos Top 10 jogadores ordenados por score.

---

## 🚀 Tecnologias utilizadas

- Go (Golang)
- PostgreSQL (Neon - cloud)
- pgx (driver PostgreSQL)
- HTTP (net/http)
- Postman (testes)

---

## 📁 Estrutura do projeto
ranking-api-go/
├── db/
│ └── connection.go
├── handlers/
│ └── ranking.go
├── models/
│ └── ranking.go
├── main.go
├── go.mod
└── go.sum

---

## ⚙️ Configuração do ambiente

### 🔹 1. Clone o repositório

```marcelo81rds
git clone https://github.com/SEU_USUARIO/ranking-api-go.git
cd ranking-api-go

🔹 2. Configure a variável de ambiente
DATABASE_URL=postgresql://user:password@host/database?sslmode=require

🔹 3. Instale dependências
go mod tidy

🔹 4. Execute a API
go run main.go

Servidor rodando em:
http://localhost:8080

🔗 Endpoints
📥 Criar pontuação

POST /ranking

{
  "nome": "Marcelo",
  "pontuacao": 100
}
📤 Listar ranking (Top 10)

GET /ranking

📌 Retorna os 10 maiores scores ordenados:

[
  {
    "id": 1,
    "nome": "Player1",
    "pontuacao": 900
  }
]
🧪 Testes

A API pode ser testada utilizando o Postman ou qualquer cliente HTTP.

🎮 Integração

Projetada para integração com jogos desenvolvidos na Unity, funcionando como backend para sistemas de leaderboard.

📈 Melhorias futuras
 Atualizar score apenas se maior (high score)
 Autenticação de usuários
 Deploy da API (Render / Railway)
 Cache de ranking
 Versionamento da API (v1)

👨‍💻 Autor
Desenvolvido por Marcelo Ribeiro dos Santos

