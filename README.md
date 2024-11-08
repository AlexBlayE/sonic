# sonic

## Description
A library that simplificates the management of tcp and tls sockets.💻

## How to use
Basic tcp:
```go
    reciever := sonic.NewReciever[string](codec.NewGobCodec(), 10)
	sender := sonic.NewSender[string](codec.NewGobCodec())

	listener, dialer := dialers.MakeTcpListenerDialer(":8080")		
    manager := sonic.NewManager(reciever, sender, listener, dialer)

	for msg := range manager.Recv() {
		fmt.Println("Recieved -> " + msg)
		manager.Send("Response", "localhost:3000")
	}
```
