# fullcycle-imersao13

Transações assíncronas são importantes, pois, garantimos que os dados não sejam perdidos no meio do caminho mesmo se diversos microsserviços fiquem fora do ar.
## Kafka
    Garante que os dados entre sistemas no sejam perdidos, além disso ele consegue persistir
    esses dados em disco, garantindo que o sistema tenha a chance de reprocessar essas informacoes.
    O kafka vai além do que só um sistema de filas.


Conceitos básicos:
- O que é um "Topic"?
  
  - Stream de dados que atua como um banco de dados(Recebe as informações e guarda)
  O tópico manda a mensagem para várias repartições. Essas repartições são enumeradas: 0, 1, 2... 
  Nós escolhemos quantas repartições terão o nosso tópico
  -  Producers: é o responsável por mandar msg em algum tópico do kafka (cara que esta produzindo a informação)
  -  Consumer:  responsável por ler a msg do kafka
  - 
    - posso ter vários consumers para ler o mesmo tópico - sistemas diferente(sistema finaceiro e de distribuicao)) OBS: esses consumers nn precisam esta lendo a mesma msg em forma simultanea
![Captura de Tela 2023-06-24 às 09.07.56.png](..%2F..%2FDesktop%2FCaptura%20de%20Tela%202023-06-24%20%C3%A0s%2009.07.56.png)
![Captura de Tela 2023-06-24 às 09.08.03.png](..%2F..%2FDesktop%2FCaptura%20de%20Tela%202023-06-24%20%C3%A0s%2009.08.03.png)
 
- Kafka Cluster
  - conjunto de Brokers(maquina que tem o kafka instalado e esse cara vai armazenar as msg,
  vai conseguir deixar essas msg serem consumidas e que alguem publique msg neles)
  - cada Broker é um servidor
  - cada Broken é responsavel por armazenar os dados de uma reparticao
  - cada particao de topic esta distribuido em diferente brokers
  
![Captura de Tela 2023-06-24 às 09.14.33.png](..%2F..%2FDesktop%2FCaptura%20de%20Tela%202023-06-24%20%C3%A0s%2009.14.33.png)
Se uma maquina cai, eu ainda tenho duas particoes funcionando para que o sistema continue lendo