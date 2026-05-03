# 🚀 Space Shooter

Um jogo de nave espacial desenvolvido em Go com a biblioteca [Ebiten](https://ebitengine.org/). Desvie e destrua meteoros, acumule pontos e tente bater seu recorde!

![Go](https://img.shields.io/badge/Go-1.22-00ADD8?style=flat&logo=go)
![Ebiten](https://img.shields.io/badge/Ebiten-v2-orange?style=flat)
![License](https://img.shields.io/badge/license-MIT-green?style=flat)

---

## 🎮 Gameplay

- Controle sua nave e destrua os meteoros antes que eles te atinjam
- Cada meteoro destruído vale 1 ponto
- O jogo reinicia automaticamente ao colidir com um meteoro
- Estrelas ao fundo criam o efeito de movimento espacial

---

## 🕹️ Controles

| Tecla | Ação |
|-------|------|
| ← Seta Esquerda | Mover nave para a esquerda |
| → Seta Direita | Mover nave para a direita |
| Espaço | Atirar laser |

---

## 🛠️ Como rodar

### Pré-requisitos

- [Go 1.22+](https://golang.org/dl/)
- Dependências do Ebiten (Linux):

```bash
sudo apt-get install libasound2-dev libgl1-mesa-dev libxcursor-dev \
  libxi-dev libxinerama-dev libxrandr-dev libxxf86vm-dev
```

### Rodando o jogo

```bash
git clone https://github.com/Criserick/gamego.git
cd gamego
go mod tidy
go run main.go
```

---

## 📦 Download

Baixe o executável mais recente na página de [Releases](https://github.com/Criserick/gamego/releases) — sem precisar instalar Go!

---

## 🏗️ Estrutura do projeto

```
gamego/
├── assets/          # Imagens, fontes e sprites
│   ├── meteors/     # Sprites dos meteoros
│   ├── planets/     # Sprites dos planetas
│   ├── stars/       # Sprites das estrelas
│   └── assets.go    # Carregamento dos assets
├── game/            # Lógica do jogo
│   ├── game.go      # Loop principal
│   ├── player.go    # Jogador
│   ├── meteor.go    # Meteoros
│   ├── laser.go     # Lasers
│   ├── star.go      # Estrelas de fundo
│   ├── timer.go     # Timer genérico
│   └── utils.go     # Utilitários (Vector, Rect)
├── main.go          # Entrada do programa
├── go.mod
└── go.sum
```

---

## 🤝 Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir uma _issue_ ou enviar um _pull request_.

---

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
