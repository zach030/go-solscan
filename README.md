# go-solscan
_____
#### Golang sdk For Solscan.io API

##### Block -
- ```GetLastBlocks(limit=10)```  Get last **[limit]** blocks (Default Limit: **10**)
- ```GetBlockTxs(limit=10,offset=0,block=None)``` Get block transactions.
- ```GetBlockDetail(block=None)```  Detail information of given block.

##### Transaction -
- ```GetLastTxs(limit=10)``` Get last [limit] transactions
- ```GetTxWithSignature(signature=None)``` Detail information of given transaction signature

##### Account -
- ```GetAccountTokens(account=None)```  Get token balances of the given account
- ```GetAccountTxs(account=None,beforeHash=None,limit=10)``` Get list of transactions of the given account.
- ```GetAccountStakeAccounts(account=None)``` Get staking accounts of the given account
- ```GetAccountSplTransfers(account=None,limit=10,offset=0,fromTime=None,toTime=None)``` Get list of transactions make tokenBalance changes.
-  ```GetAccountSolTransfers(account=None,limit=10,offset=0,fromTime=None,toTime=None)``` Get list of SOL transfers
-  ```GetAccountExportTransactions(account=None,typ='all',fromTime=None,toTime=None)``` Export transactions to CSV. Returns **Blob** URL
-  ```GetAccountInfo(account=None)``` Get overall account information, including program account, NFT metadata information

##### Token -
- ``` GetTokenHolders(tokenAddr=None,limit=10,offset=0)``` Get token holders
- ``` GetTokenMeta(tokenAddr=None)``` Get metadata of given token
- ``` GetTokenList(sortBy='market_cap',direction='desc',limit=10,offset=0)``` Get list of tokens.

##### Market -
- ``` GetTokenMarket(tokenaddr=None)``` Get market information of the given token

##### Chain Information -
- ``` ChainInfo()``` Get market information of the given token


## Installation

```sh
go install github.com/zach030/go-solscan
```


## License
MIT
