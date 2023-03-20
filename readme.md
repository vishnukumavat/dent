# Index
- [Project Explanation](#project-explanation)
    - [Network Explanation](#what-is-dent)
    - [DApp Explanation](#what-is-vox-20)
    - [Problems - Traditional Voting](#drawbacks-in-the-traditional-voting-system)
    - [Why Blockchain](#why-only-blockchain-why-not-traditional-web20-tech)

- [Blockchain Setup](#project-setup)
- [Architectural Flow](#architectural-diagram)
- [Demo](#demo)

# Project Explanation

#### `Domain` - BlockChain
#### `Project Name` -  VOX 2.0 (Voting Experience 2.0)
#### `Network Name` - DENT (Decentralized Network)


### What is DENT?  

DENT is a third-generation blockchain while Ethereum and bitcoin are 2nd and 1st generation. So apparently DENT has many differences and advantages over 2nd and 1st gen blockchains. DENT is built using set of open source tools like Tendermint, the Cosmos SDK and IBC designed to let people build custom, secure, scalable and interoperable blockchain applications quickly. Below attached visual can help in better understanding of where DENT lies in the Blockchain Ecosystem.
<br>DENT consist of all the features from its descendant blockchains.

<img width="799" alt="image" src="https://user-images.githubusercontent.com/43311385/226431531-d0ac558a-749a-42c1-a336-f4c0879013c4.png">


### What is VOX 2.0?
	
VOX 2.0 or Voting Experience 2.0 is dApp and new generation voting platform which allows to register existing voters on the blockchain, create new polls and allow voters to vote on active polls. All the data and transactions made are fully secured with blockchain technology.  This is built on a sovereign blockchain network called [DENT](#what-is-dent).

<img width="424" alt="image" src="https://user-images.githubusercontent.com/43311385/226430447-cef638e6-e2dc-4c1e-a088-a3dd6dd4fb57.png">


### Drawbacks in the traditional voting system.

- `Voter suppression:` In India, there have been reports of voter suppression due to factors such as caste, religion, and gender other factors are Voter ID laws, Voter purging, Voter intimidation, Inadequate polling locations

- `Long queues and waiting times:` During elections, long queues and waiting times can discourage people from voting. In India, voters often have to stand in long queues for hours to cast their vote.

- `Voter fraud:` India has a history of voter fraud, including booth capturing, vote rigging, and bribing voters.

- `Human error:` The traditional voting system in India is dependent on human intervention, which can result in human errors that may affect the accuracy of the vote count.

- `Cost:` Traditional voting systems can be expensive to implement and maintain, especially in a country with a large population like India.

- `Limited accessibility:` In India, some people may not be able to vote due to physical disabilities or geographical location, as they may not be able to access a polling station.

- `Limited flexibility:` Traditional voting systems in India may not be able to accommodate unforeseen circumstances, such as unexpected changes in polling locations or inclement weather.

- `Limited transparency:` Traditional voting systems in India may not provide enough transparency to ensure that the electoral process is fair and unbiased. There have been instances of vote-counting discrepancies and allegations of vote tampering.


### For all the above problems - Introducing VOX2.0 (Voting Experience 2.0) powered by DENT (Decentralized Network), A blockchain-based voting solution.

### Why only blockchain? Why not traditional Web2.0 tech?

- `Security:` Blockchain technology is inherently secure and tamper-proof. Each transaction is verified by multiple nodes on the network, making it almost impossible to manipulate the data or tamper with the results. This makes it an ideal technology for secure and transparent voting.

- `Transparency:` Blockchain technology provides a high level of transparency, which is crucial in any voting system. Each transaction on the blockchain is public and can be traced back to its origin, making it easy to detect any fraudulent activities.

- `Decentralization:` Blockchain technology allows for a decentralized voting system, where each node on the network has a copy of the entire database. This eliminates the need for a central authority to manage the voting process, making it more resistant to hacking or other types of attacks.

- `Immutability:` Once a vote is cast on the blockchain, it cannot be changed or deleted. This ensures that the results of the election are final and cannot be manipulated after the fact.

- `Accessibility:` Blockchain technology can make voting more accessible to people who are unable to vote in traditional ways. For example, people with disabilities, those living in remote areas, or people who are unable to travel to a polling station can vote securely and conveniently from their homes.

<br>

# Project Setup

## Requirements 
- [Make](https://www.gnu.org/software/make/#download)
- [Go 1.18+](https://go.dev/doc/install)

## Blockchain setup

- ###  `make install` command installs dependencies and builds the executable binary and places it into `~/go/bin/`
```
$ make install
```
<img width="1180" alt="image" src="https://user-images.githubusercontent.com/43311385/226217730-3afb02d2-2c6b-429c-8bbf-b6ecef6c387d.png">


- ### Move the `dent` binary into the `/usr/local/bin` to make it globally accessible.

```
$ sudo mv ~/go/bin/dent /usr/local/bin
```
<img width="853" alt="image" src="https://user-images.githubusercontent.com/43311385/226217675-4ae31d04-b853-4a54-b0ac-4c297254ec21.png">


- ### Validate if binary installation is properly done, as shown below.
```
$ dent
```
<img width="895" alt="image" src="https://user-images.githubusercontent.com/43311385/226217865-c124704e-f8ba-43f3-9118-351480959636.png">


- ### Initialize the new blockchain chain with id `test-1`
```
$ rm -rf ~/.dent
$ dent init test --chain-id test-1
$ dent keys add cooluser --keyring-backend test
$ dent add-genesis-account $(dent keys show cooluser --keyring-backend test -a) 1000000000000000000udent,1000000000000000000stake
$ dent gentx cooluser 1000000000stake --chain-id test-1 --keyring-backend test
$ dent collect-gentxs
```
<img width="1440" alt="image" src="https://user-images.githubusercontent.com/43311385/226218606-e9b22a8a-5d95-4830-8efe-41b5668f0310.png">


- ### Start Chain
```
$ dent start
```
<img width="1440" alt="image" src="https://user-images.githubusercontent.com/43311385/226218717-7327deaa-59b8-471c-9068-6159fb51ed7f.png">


- ### Configure

Your blockchain in development can be configured with `config.yml` and `app.toml` located under `~/.dent/config/`.


# Architectural Diagram

<img width="833" alt="image" src="https://user-images.githubusercontent.com/43311385/226439106-ade99357-957c-4306-b90f-465aca0f7487.png">


# Demo

[https://youtu.be/bGW0NY06h2s](https://youtu.be/bGW0NY06h2s)