## Encrypt & Hash

### Usage

#### Generate Keys
```bash
go run main.go key:generate
```

#### Encrypt RSA
```bash
go run main.go rsa:encrypt <text>
```

#### Encrypt AES
```bash
go run main.go aes:encrypt <text>
```

#### Hash
```bash
go run main.go hash <text>
```

#### Example
```bash
go run main.go rsa:encrypt "YSS2955"

go run main.go aes:encrypt "YSS2955"

go run main.go hash "YSS2955"
```

