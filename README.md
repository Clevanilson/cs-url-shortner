# CS URL Shortener

Este é um serviço de encurtamento de URLs desenvolvido em Go. O objetivo do projeto é fornecer uma maneira eficiente e rápida de transformar URLs longas em URLs curtas e gerenciar o redirecionamento.

## Requisitos Funcionais

1.  **Encurtamento de URL**: Dado um URL longo, o sistema deve retornar um URL curto correspondente.
2.  **Redirecionamento**: Ao acessar um URL curto, o sistema deve redirecionar o usuário para o URL longo original.

## Requisitos Não Funcionais

-   **Tamanho da URL**: O tamanho da URL encurtada deve ser o mais curto possível.
-   **Caracteres**: A URL curta deve conter apenas caracteres alfanuméricos (`0-9`, `a-z`, `A-Z`).
-   **Alta Disponibilidade**: O sistema deve estar disponível 24 horas por dia, 7 dias por semana (24/7).
-   **Escalabilidade e Performance**: O sistema é projetado para suportar uma carga de **1 operação de leitura para cada 10 operações de escrita** (sistema focado em escrita).
