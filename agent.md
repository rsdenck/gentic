# Projeto: Infrastructure AI Agents
- Os quatro agentes
01. VMware Agent

Virtualização e ambiente VMware
- vCenter e ESXi
- Máquinas virtuais
- Datastores e snapshots
- Clusters, hosts e recursos

02. Linux Agent

Sistemas operacionais Linux
- CPU, memória e disco
- systemd e serviços
- Logs e processos
- Diagnóstico e troubleshooting

03. Monitoring Agent

Observabilidade e análise de eventos
- Graylog
- Pesquisa e correlação de logs
- Identificação de erros
- Análise de incidentes

04. Cloud Agent

Cloud privada e infraestrutura
- Apache CloudStack
- Instâncias e volumes
- Redes e ofertas
- Recursos e quotas
----------------------------------
# TASK: Desenvolver a plataforma Infrastructure AI Agents

Atue como um engenheiro de software sênior, arquiteto de sistemas distribuídos, especialista em Go, Firebase Genkit, Ollama, automação de infraestrutura e segurança.

Sua responsabilidade é IMPLEMENTAR o projeto completo. Não quero apenas um planejamento, exemplos isolados ou documentação conceitual. Analise o repositório, implemente os componentes, execute os testes, corrija os problemas e documente o resultado.

## 1. Contexto do ambiente

Ambiente de desenvolvimento:
- Sistema operacional: AlmaLinux 9.8.
- Linguagem: Go.
- Go instalado: 1.26.7.
- Framework de IA: Firebase Genkit Go.
- Plugin de inferência: Ollama.
- Diretório atual do projeto: /root/projects/gentic.
- Endpoint remoto do Ollama: https://api.denck.sbs/
- Autenticação: token disponibilizado por variável de ambiente.
- Variável esperada: OLLAMA_API_KEY.

O projeto Go já foi inicializado anteriormente. Existe uma tentativa de integração com Firebase Genkit e Ollama.

IMPORTANTE:
- Antes de alterar qualquer arquivo, inspecione o projeto existente.
- Preserve configurações e implementações válidas.
- Não presuma que o endpoint remoto utiliza todas as rotas nativas do Ollama.
- Verifique a compatibilidade da API e do plugin instalado.
- Não utilize o Ollama local como fallback silencioso.
- Não exponha tokens, senhas ou credenciais em logs, arquivos ou documentação.
- Não invente modelos disponíveis no endpoint remoto. Descubra-os quando possível ou permita configuração explícita.

## 2. Objetivo

Criar uma plataforma modular chamada:

Infrastructure AI Agents

A plataforma terá quatro agentes independentes:

1. VMware Agent.
2. Linux Agent.
3. Monitoring Agent.
4. Cloud Agent.

Todos devem utilizar o mesmo núcleo de IA, mas possuir instruções, ferramentas, responsabilidades e fluxos independentes.

## 3. Arquitetura

Comece com um único serviço Go modular.

Estrutura desejada:

infrastructure-ai/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── agents/
│   │   ├── vmware/
│   │   ├── linux/
│   │   ├── monitoring/
│   │   └── cloud/
│   ├── tools/
│   │   ├── vmware/
│   │   ├── linux/
│   │   ├── graylog/
│   │   └── cloudstack/
│   ├── ai/
│   ├── api/
│   ├── config/
│   ├── security/
│   ├── audit/
│   └── workflows/
├── configs/
├── deployments/
├── docs/
├── scripts/
├── tests/
├── go.mod
├── go.sum
├── .env.example
├── .gitignore
└── README.md

Adapte a estrutura ao repositório existente quando necessário. Evite duplicar arquivos ou criar estruturas desnecessárias.

## 4. Núcleo de IA

Implementar um núcleo compartilhado utilizando Firebase Genkit Go.

Responsabilidades:
- Inicialização do Genkit.
- Configuração do plugin Ollama.
- Comunicação com endpoint remoto.
- Autenticação por token.
- Configuração do modelo por variável de ambiente.
- Timeout configurável.
- Tratamento de erros.
- Logs estruturados.
- Validação de conectividade.
- Separação entre configuração e código.

Variáveis de ambiente desejadas:

OLLAMA_SERVER_ADDRESS=https://api.denck.sbs
OLLAMA_API_KEY=
OLLAMA_MODEL=
OLLAMA_TIMEOUT=300

Não inserir valores secretos no código.

Utilizar a versão compatível do plugin oficial do Genkit. Consultar a documentação e verificar as APIs reais disponíveis na versão instalada antes de implementar.

Criar uma interface interna que permita aos agentes utilizar o mecanismo de inferência sem conhecer detalhes de conexão.

## 5. VMware Agent

Criar um agente especializado em infraestrutura VMware.

Responsabilidades:
- Consultar vCenter.
- Consultar ESXi.
- Listar máquinas virtuais.
- Consultar estado das VMs.
- Consultar CPU e memória.
- Consultar datastores.
- Consultar snapshots.
- Consultar hosts e clusters.
- Analisar problemas de capacidade.
- Gerar diagnósticos técnicos.

Ferramentas devem ser implementadas de forma modular.

Preparar integração com VMware vSphere API/SDK oficial para Go.

Configurações:

VMWARE_VCENTER_URL=
VMWARE_USERNAME=
VMWARE_PASSWORD=
VMWARE_INSECURE=false

Não inventar resultados quando o vCenter não estiver configurado.

## 6. Linux Agent

Criar um agente especializado em administração e diagnóstico Linux.

Responsabilidades:
- Consultar CPU.
- Consultar memória.
- Consultar filesystem.
- Consultar utilização de disco.
- Consultar processos.
- Consultar serviços systemd.
- Consultar logs.
- Analisar processos com consumo elevado.
- Diagnosticar problemas de recursos.
- Identificar serviços indisponíveis.

Implementar ferramentas com coleta controlada.

Não executar comandos arbitrários recebidos diretamente do modelo.

Criar ferramentas específicas, com argumentos validados e permissões restritas.

A execução de comandos deve ter lista de operações permitidas, timeout e registro de auditoria.

Preparar suporte a execução local e integração remota SSH, sem exigir que ambas estejam configuradas para iniciar o serviço.

## 7. Monitoring Agent

Criar um agente especializado em monitoramento e observabilidade.

Integração prioritária:
- Graylog 7.x.
- API REST do Graylog.
- Pesquisa de mensagens.
- Consulta de streams.
- Análise de eventos.
- Correlação temporal.
- Agrupamento de erros.
- Identificação de padrões.
- Geração de relatórios técnicos.

Configuração:

GRAYLOG_URL=
GRAYLOG_TOKEN=
GRAYLOG_TIMEOUT=

Criar ferramentas para:
- Pesquisar logs por período.
- Pesquisar por hostname.
- Pesquisar por aplicação.
- Pesquisar por nível de severidade.
- Consultar mensagens relacionadas a incidentes.
- Correlacionar eventos.

Não classificar alertas normais, flapping ou eventos esperados como incidentes sem evidências.

Toda conclusão deve distinguir:
- Evidência encontrada.
- Hipótese.
- Causa confirmada.
- Recomendação.

## 8. Cloud Agent

Criar um agente especializado em cloud privada.

Integração prioritária:
- Apache CloudStack.
- API de gerenciamento.
- Instâncias.
- Volumes.
- Redes.
- Templates.
- Service Offerings.
- Recursos e quotas.
- Domínios e contas.
- Estado dos recursos.

Implementar assinatura/autenticação da API de acordo com a documentação oficial do CloudStack.

Configuração:

CLOUDSTACK_URL=
CLOUDSTACK_API_KEY=
CLOUDSTACK_SECRET_KEY=
CLOUDSTACK_TIMEOUT=

Criar ferramentas específicas para consultas e diagnósticos.

Não inventar dados quando a API não estiver disponível.

## 9. API central

Criar API HTTP REST para disponibilizar os quatro agentes.

Endpoints mínimos:

GET /health
GET /api/v1/agents
POST /api/v1/agents/vmware/chat
POST /api/v1/agents/linux/chat
POST /api/v1/agents/monitoring/chat
POST /api/v1/agents/cloud/chat

Criar também endpoints de diagnóstico e documentação da API quando apropriado.

Requisitos:
- JSON.
- Validação de entrada.
- Respostas padronizadas.
- Tratamento de erros.
- Timeout.
- Context cancellation.
- Logs estruturados.
- Documentação OpenAPI.
- Versionamento da API.

Não retornar credenciais ou configurações sensíveis.

## 10. Segurança

Implementar desde o início:

- Autenticação da API.
- RBAC ou estrutura equivalente de autorização.
- Separação de permissões por agente.
- Auditoria das solicitações.
- Registro de ferramentas executadas.
- Proteção de credenciais.
- Validação dos argumentos.
- Timeouts.
- Limites de execução.
- Proteção contra prompt injection proveniente de logs e dados externos.

Política inicial: SOMENTE LEITURA.

Nenhum agente poderá reiniciar máquinas, desligar VMs, alterar redes, apagar arquivos, modificar configurações ou executar ações destrutivas.

Deixar a arquitetura preparada para operações futuras com aprovação humana explícita, mas não habilitar essas operações nesta versão.

Não registrar tokens, senhas ou dados secretos nos logs de auditoria.

## 11. Workflows

Criar fluxos Genkit que combinem agentes e ferramentas.

Implementar pelo menos:

- Diagnóstico de VM.
- Diagnóstico de servidor Linux.
- Análise de incidentes Graylog.
- Diagnóstico de infraestrutura CloudStack.
- Diagnóstico correlacionado entre VMware, Linux e Monitoring.

Cada workflow deve:
1. Receber uma solicitação.
2. Validar os parâmetros.
3. Consultar ferramentas autorizadas.
4. Reunir evidências.
5. Utilizar o modelo para analisar os dados.
6. Retornar diagnóstico estruturado.
7. Registrar a execução na auditoria.

Não permitir que o modelo invente evidências ou resultados de consultas.

## 12. Testes

Implementar testes unitários e de integração.

Os testes devem cobrir:
- Inicialização do serviço.
- Inicialização do Genkit.
- Configuração do Ollama.
- Ausência de token.
- Falha de conexão.
- Timeout.
- Respostas inválidas.
- API REST.
- Autenticação.
- Autorização.
- Ferramentas de cada agente.
- Auditoria.
- Workflows.

Utilizar mocks para integrações externas.

Não exigir VMware, Graylog ou CloudStack reais para executar os testes unitários.

Executar:

go fmt ./...
go vet ./...
go test ./...

Corrigir os erros encontrados.

Não afirmar que uma integração real foi testada se apenas foram executados testes com mocks.

## 13. Documentação

Criar documentação completa:

README.md:
- Descrição.
- Arquitetura.
- Requisitos.
- Instalação.
- Configuração.
- Execução.
- API.
- Agentes.
- Segurança.
- Testes.

Criar também:

docs/ARCHITECTURE.md
docs/VMWARE_AGENT.md
docs/LINUX_AGENT.md
docs/MONITORING_AGENT.md
docs/CLOUD_AGENT.md
docs/API.md
docs/SECURITY.md
docs/DEPLOYMENT.md

Criar .env.example somente com valores fictícios ou vazios.

## 14. Implantação

Preparar implantação nativa no AlmaLinux 9.8.

Criar:
- Script de build.
- Script de execução.
- Arquivo systemd.
- Configuração de usuário de serviço não privilegiado.
- Diretórios de configuração e logs.
- Instruções de instalação e atualização.

Não utilizar Docker ou Podman como dependência obrigatória.

Não instalar pacotes no sistema nem alterar serviços existentes sem necessidade e sem autorização.

## 15. Regras de execução

Trabalhe de maneira autônoma dentro do escopo autorizado.

Ordem obrigatória:

1. Inspecionar todos os arquivos existentes.
2. Verificar Go, módulos e dependências.
3. Validar a integração existente com Genkit.
4. Planejar as alterações internamente.
5. Implementar o núcleo de IA.
6. Implementar os quatro agentes.
7. Implementar as ferramentas.
8. Implementar API.
9. Implementar segurança e auditoria.
10. Implementar workflows.
11. Criar testes.
12. Executar testes e corrigir problemas.
13. Criar documentação.
14. Preparar implantação.
15. Apresentar relatório final.

Não interrompa o trabalho para pedir confirmação de decisões técnicas comuns.

Se faltar alguma credencial ou integração externa, implemente a interface e os mocks, documente exatamente o que falta e continue o restante do projeto.

Não simule sucesso de integrações externas.

Não substitua bibliotecas oficiais por implementações improvisadas sem justificativa técnica.

Não deixe funções vazias, TODOs críticos ou código fictício apresentado como funcional.

Não sobrescreva arquivos válidos sem necessidade.

Não altere o ambiente externo, infraestrutura real ou sistemas de produção durante os testes.

## 16. Critérios de conclusão

O projeto só deve ser considerado concluído quando:

- Os quatro agentes estiverem implementados.
- O núcleo Genkit estiver funcional.
- A configuração do Ollama remoto estiver implementada.
- A API REST estiver funcional.
- A segurança básica estiver implementada.
- A auditoria estiver funcionando.
- Os workflows estiverem implementados.
- Os testes automatizados estiverem executados.
- A documentação estiver criada.
- A implantação nativa estiver preparada.

## 17. Relatório final

Ao concluir, apresentar:

1. Arquitetura implementada.
2. Estrutura final de diretórios.
3. Arquivos criados e modificados.
4. Dependências utilizadas e versões.
5. Situação de cada agente.
6. Situação da integração com Ollama.
7. Endpoints disponíveis.
8. Resultado de go vet e go test.
9. Integrações testadas com serviços reais.
10. Integrações que ainda dependem de credenciais.
11. Como executar o projeto.
12. Pendências reais e limitações.

Se algum requisito não puder ser implementado, explicar objetivamente o motivo.

**Comece inspecionando o repositório atual e execute a implementação completa.**
