# stress-tester
O stress-tester é uma ferramenta para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.
## Exemplo de uso
### Com o docker rodando, utilize o seguinte comando:
#### docker run stress-tester-app --url=http://google.com.br --requests=1000 --concurrency=10
### O resultado deve será um relatório dos testes realizados, como o seguinte:
#### - Quantidade total de requests realizadas: 1000
#### - Quantidade de requests com status HTTP 200: 1000
#### - Tempo total gasto na execução: 46.15455764s
### Detalhamento dos parâmetros:
#### --url: URL do serviço a ser testado.
#### --requests: Número total de requests.
#### --concurrency: Número de chamadas simultâneas.
### É sempre possível consultar o help da ferramenta por meio do seguinte comando:
#### docker run stress-tester-app -h
