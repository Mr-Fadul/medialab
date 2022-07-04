## TODO

- Criar pipeline de video e audio fake
- Verificar consumo excessivo de memória

## DONE

- Controle de slice de arquivos baseado no tempo estipulado
- Adicionar controle de framerate maximo no video
- Ajustes no dockerfile para layers
- Ajustes de controle de sinais com o dumb-init
- Criar funcao de envio de arquivos para azure
- Ajustes de verbosidade na lib zerolog
- Criar Função para desativar envio de arquivos
- Controlar o envio para azure com path na env CHANNEL_LOCATION
- Criar compilação do plugin interpipes no docker
- Usar plugin interpipes para parar o arquivo sem deletar o pipeline principal
- Criar Função para deletar arquivos enviados com xx data de vida
- Tratamento de erros http
- Adicionar metodo de envio de arquivos para GCP