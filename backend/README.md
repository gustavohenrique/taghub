## Sobre

Ecommerce é um serviço para lidar com catálogo de produtos e pagamentos. Suas principais funcionalidades são:

- Gestão de catálogo (cadastro de produtos, lotes, cupons)
- Integração com a Mundipagg para pagamentos via cartão de crédito e boleto
- Acompanhamento de pedidos
- Criação de cupons para um ou mais produtos e alunos
- Cobrança recorrente para produtos do tipo assinatura
- Fornecer informações sobre o que o aluno pode acessar de acordo com o que comprou

## Funcionamento

### Compra de um lote

1. Site obtém ou cria um customer na mundipagg de acordo com o `student_id` do usuário autenticado;
2. Site exibe os cartões e endereços cadastrados na mundipagg e permite que o aluno crie um novo cartão e endereço;
3. O aluno pode inserir um cupom e ao clicar no botão, o site apresenta o valor com desconto;
4. Ao finalizar o pedido, o site envia dados do cartão (apenas card_id), do aluno autenticado, do lote e do cupom;
5. API verifica os dados, aplica novamente o cupom para obter o desconto e cria um pedido na mundipagg. Se o pedido for criado com sucesso, armazena as informações no banco de dados;
6. Site exibe mensagem de sucesso para o aluno e o código de barras caso o pagamento escolhido seja boleto;
7. Após um período (+/- 1 minuto), a mundipagg envia as informações sobre o pedido para a API (webhook) que atualiza o status e outras informações sobre o pedido no banco de dados;
8. API envia um email ao aluno informando sobre a confirmação ou falha no pagamento;
9. API envia uma requisição ao CRM para gerar a NFe;
10. Todo mundo fica feliz =)

## Setup

```
go get github.com/golang/mock/mockgen@latest
go get golang.org/x/tools/cmd/goimports
go get -u github.com/rakyll/gotest
```

Para rodar em localhost:

```sh
make createdb  # cria um container docker do Postgres com o database ecommerce
make seed      # cria as tabelas no container
make run       # roda o serviço escutando em localhost:3008 com as variáveis de ambiente padrão
```

## Organização

```
.
├── build                # scripts de automatização
├── libs                 # libs privadas
│   ├── configuration    # obtém configurações do Consul ou de envvars
│   ├── errors           # cria erros com códigos específicos
│   ├── filter           # transforma requests de search em cláusulas WHERE do SQL
│   ├── httpclient       # implementa o FastHTTP para consumir APIs de terceiros
│   ├── httpvalidator    # implementa o validator.v9 para validar structs anotadas
│   ├── logger           # implementa o logrus para log de mensagens
│   ├── mundipagg        # implementa chamadas às APIs da mundipagg
│   ├── stringutils      # funções para manipulação e tratamento de strings
│   └── testutils        # facilita a escrita de testes
├── mocks                # mocks gerados pelo gomock a partir das interfaces definidas
├── pkg                  # atualiza a versão do produto durante o build de acordo com a branch
├── sql                  # schema do database e dados iniciais utilizados pelos testes
└── src                  # contém o código da aplicação
    ├── containers       # containers dos services e repositories injetados pelo main.go
    ├── domain           # structs, interfaces e adapters referentes ao domínio da aplicação
    ├── handlers         # implementa APIs e rotas REST utilizando o Echo Framework
    ├── http             # configura o servidor HTTP e adicionas os handlers
    ├── postgres         # acesso ao banco de dados
    └── services         # lógica de neǵocio e transformação de dados
```

### Arquitetura do Código

O código está baseado em uma arquitetura de 3 camadas:

- **handlers**: API Rest, porta de entrada da aplicação, que recebe requests HTTP, valida, converte para structs do projeto e encaminnha para a próxima camada, autando basicamente como um orquestrador.
- **services**: Aplica as regras de negócio e transformações necessárias para preparar os dados para serem armazenados em algum repositório de dados.
- **repositories**: Armazena os dados em banco de dados ou em outro repositório de dados (mundigpagg por exemplo).

O arquivo `main.go` é responsável principalmente pela injeção de dependências. Os services e repositories possuem um container contendo todos os existentes. O container dos repositories são injetados em cada service e o containers contendo todos os services é injetado em cada handler.  
Mantendo essa estrutura, fica simples evoluir a aplicação, criar mocks e separar os testes automatizados.

Cada handler está separado em um pacote diferente. Cada um também é responsável por montar sua estrutura de rotas. As rotas foram separadas em 2 grupos: /api para acesso externo e /ws para chamadas feitas pelo BO. Assim, cada grupo pode ter seu próprio middleware e aplicar regras separadas de autorização para alunos e usuários BO (managers).

### Testes Automatizados

O testes automatizados foram separados em 1 arquivo por função. Isso evita arquivos gigantes contendo centenas de linhas de código. Existem 3 tipos de teste:

1. **unidade**: Usado nos testes de handlers apenas para verificar o status code retornado. O services ficam mockados nesse caso.
2. **híbridos**: Usado nos services no qual, dependendo do caso, pode ser um teste unitário com os repositories mockados ou de integração, utilizando os repositories reais e se conectando ao banco de dados de testes.
3. **integração**: Usado nos repositories, visando TDD, permitindo que o código escrito acesse o banco de dados para ter certeza que funciona com o Postgres. As tabelas são recriadas a cada teste.


### Linter

O projeto utiliza o *goimports* como linter padrão. Para que não configurou o editor para isso, basta executar o comando `make lint` que as alterações serão aplicadas.

### Libs e Utils

Nenhuma dependência de terceiros deve ser importada diretamente no código. Cada dependência deve ser encapsulada em uma lib interna do projeto. Isso facilita a migração de uma determinada lib por outra, sem precisar modificar a implementação utilizada no restante do código.  

Códigos repetitivos foram extraídos para um pacote utils. Ex.: converter structs de/para JSON string, validar datas, gerar strings aleatórias, etc.

### Scripts Utilitários

A pasta `build` contém scripts shell para automatizar tarefas cotidianas do desenvolvimento.

### Gerador de CRUD

O target `crud` cria as interfaces, handlers, services e repositories de uma entidade a partir de outra. Ideal para preguiçosos =) Ex.:

```
make crud from=products to=orders
```

## Contribuição

Antes de abrir um PR:

- Não crie um mega commit contendo diversas alterações. Prefira vários commits menores.
- Não use `git commit -m ""`. Descreva direito o motivo disso ter sido feito.
- Não escreva mensagens de commit em inglês. Use português. É mais rápido para a maioria descrever os motivos no idioma principal.
- Não faça push antes de rodar os testes localmente. Execute `make tests`.
- Não adicione qualquer dependência porcaria no projeto. Quanto menos dependência, melhor. Se você só precisa de uma ou outra função fornecida por uma lib, copie o código dela e crie uma lib privada.

Ao ler (ou não ler) as recomendações acima, você está dando direito ao restante do time de praticar o devido bullying com sua falta de atenção nos canais do Slack ou a qualquer momento no escritório da Estratégia.

![image](https://www.itu.com.br/img/conteudo/54760.jpg)
