### mysql
  - tabelas
```
show tables
```

  - descrição da tabela
```
desc NomeTabela
```


### Docker
- Criar container a partir docker compose:
```
  docker compose up -d
```

- bash:
```
  docker compose exec mysql bash
```

- acessar a tabela pelo bash (pra executar alguma query)
```
  mysql -u root -p goexpert
```

### gomod
```
  go mod init NomeModulo
```

  - verifica pacotes a serem instalados
```
  go mod tidy
```


