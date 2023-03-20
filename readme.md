# dent
**dent** is a blockchain built using Cosmos SDK and Tendermint.

### Requirements 
- [Make](https://www.gnu.org/software/make/#download)
- [Go 1.18+](https://go.dev/doc/install)
### Devnet Setup

### 1. `make install` command installs dependencies and builds the executable binary and places it into `~/go/bin/`
```
$ make install
```
<img width="1180" alt="image" src="https://user-images.githubusercontent.com/43311385/226217730-3afb02d2-2c6b-429c-8bbf-b6ecef6c387d.png">


### 2. move the `dent` binary into the `/usr/local/bin` to make it globally accessible.

```
$ sudo mv ~/go/bin/dent /usr/local/bin
```
<img width="853" alt="image" src="https://user-images.githubusercontent.com/43311385/226217675-4ae31d04-b853-4a54-b0ac-4c297254ec21.png">

### 3. Validate if binary installation is properly done, as shown below.
```
$ dent
```
<img width="895" alt="image" src="https://user-images.githubusercontent.com/43311385/226217865-c124704e-f8ba-43f3-9118-351480959636.png">


### 4. Initialize the new blockchain chain with id `test-1`
```
$ rm -rf ~/.dent
$ dent init test --chain-id test-1
$ dent keys add cooluser --keyring-backend test
$ dent add-genesis-account $(dent keys show cooluser --keyring-backend test -a) 1000000000000000000udent,1000000000000000000stake
$ dent gentx cooluser 1000000000stake --chain-id test-1 --keyring-backend test
$ dent collect-gentxs
```
<img width="1440" alt="image" src="https://user-images.githubusercontent.com/43311385/226218606-e9b22a8a-5d95-4830-8efe-41b5668f0310.png">

### 5. Start Chain
```
$ dent start
```
<img width="1440" alt="image" src="https://user-images.githubusercontent.com/43311385/226218717-7327deaa-59b8-471c-9068-6159fb51ed7f.png">

### Configure

Your blockchain in development can be configured with `config.yml` and `app.toml` located under `~/.dent/config/`.