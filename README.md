<h1>API Rest - Pizzaria</h1>

- Esse projeto de API para uma pizzaria utiliza o framework Gin em Go para criar rotas e gerenciar as operações relacionadas ao cadastro e consulta de pizzas. Aqui estão os principais aspectos do código:
### Funcionalidades da API:

- GET /pizzas : Retorna todas as pizzas cadastradas no sistema em formato JSON.
- POST /pizzas : Permite a adição de uma nova pizza, recebendo os dados em JSON e armazenando no arquivo pizzas.json.
- GET /pizzas/ : Retorna uma pizza específica com base no ID fornecido.
### Estrutura do Código:
 - Modelos: As pizzas são representadas pela estrutura definida em models.Pizza, que contém os atributos como ID, Nome, Preço, etc.
### Rotas:
<ul>
  <li> O framework Gin é usado para definir rotas de API RESTful.</li>
  <li> A função getPizzas retorna a lista completa de pizzas.</li>
  <li>A função postPizzas adiciona uma nova pizza à lista e persiste os dados em um arquivo JSON.</li>
  <li>A função getPizzasById busca uma pizza pelo seu ID.</li>
 </ul>




### Persistência:
- Os dados das pizzas são armazenados em um arquivo JSON (pizzas.json), sendo carregados ao iniciar o serviço com a função loadPizzas e salvos com a função savePizza sempre que uma nova pizza é adicionada.
### Tecnologias Utilizadas:
- Linguagem: Go
- Framework Web: Gin
- Persistência de Dados: Arquivo JSON

<h6>Esse código implementa uma API REST simples e eficiente, ideal para o gerenciamento de pedidos de pizzas. O uso de persistência em arquivo JSON facilita a prototipagem e o desenvolvimento rápido, enquanto o Gin proporciona rotas rápidas e leves.</h6>
