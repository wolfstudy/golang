## 匿名函数

正常的函数调用前面已经说过，其实就是直接调用一个地址，中间外加一些参数传递的过程。

但是匿名函数并没有直接调用这个地址，而是返回一个复合结构体，这个结构体包含了一个指针对象
指向调用的那个东西，（包含的这个指针对象，这一步和正常的函数调用是一样的）

所以说，匿名函数它是一个二次寻址过程，第一次寻址寻找这个复合结构体的地址，第二次寻找这个结构体中真正指向函数调用的那个地址
也就是他有一个解包的过程。包装对象内部持有这个函数调用的地址。

这个包装对象，在运行期进行解包，返回的时候不是返回匿名函数的地址，而是返回一个复合结构体的对象。





```
key:  GoVersion   value:
key:  Version   value:
key:  BuildDate   value:
key:  DataDir   value:  data
key:  RPCListeners   value:
key:  RPCUser   value:
key:  RPCPass   value:
key:  RPCLimitUser   value:
key:  RPCLimitPass   value:
key:  RPCCert   value:
key:  RPCKey   value:
key:  RPCMaxClients   value:
key:  RPCMaxWebsockets   value:
key:  RPCMaxConcurrentReqs   value:
key:  RPCQuirks   value:
key:  Level   value:
key:  Module   value:
key:  FileName   value:
key:  MinFeeRate   value:
key:  LimitAncestorCount   value:
key:  LimitAncestorSize   value:
key:  LimitDescendantCount   value:
key:  LimitDescendantSize   value:
key:  MaxPoolSize   value:  300000000
key:  MaxPoolExpiry   value:
key:  ListenAddrs   value:  1234
key:  MaxPeers   value:  128
key:  TargetOutbound   value:  8
key:  ConnectPeersOnStart   value:
key:  DisableBanning   value:  true
key:  BanThreshold   value:
key:  SimNet   value:  false
key:  DisableListen   value:  true
key:  BlocksOnly   value:  true
key:  Proxy   value:
key:  UserAgentComments   value:
key:  DisableDNSSeed   value:
key:  DisableRPC   value:  false
key:  DisableTLS   value:  false
key:  NoOnion   value:  true
key:  Upnp   value:  false
key:  ExternalIPs   value:
key:  SimNet   value:
key:  ConnectPeers   value:
key:  NoPeerBloomFilters   value:  true
key:  DisableCheckpoints   value:  true
key:  AcceptDataCarrier   value:  true
key:  MaxDatacarrierBytes   value:  223
key:  IsBareMultiSigStd   value:  true
key:  PromiscuousMempoolFlags   value:
key:  DustRelayFee   value:  83
key:  AssumeValid   value:
key:  StartLogHeight   value:  2147483647
key:  BlockMinTxFee   value:
key:  BlockMaxSize   value:
key:  BlockVersion   value:  -1
key:  Strategy   value:  ancestorfeerate
```