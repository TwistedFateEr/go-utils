package doc



//DHCPv4(xid=0xec92bac0 hwaddr=00:0c:29:00:a4:0b msg_type=DISCOVER, your_ip=0.0.0.0, server_ip=0.0.0.0)
//DHCPv4(xid=0xec92bac0 hwaddr=00:0c:29:00:a4:0b msg_type=OFFER, your_ip=192.168.147.154, server_ip=192.168.147.254)
//DHCPv4(xid=0xec92bac0 hwaddr=00:0c:29:00:a4:0b msg_type=REQUEST, your_ip=0.0.0.0, server_ip=192.168.147.254)
//DHCPv4(xid=0xec92bac0 hwaddr=00:0c:29:00:a4:0b msg_type=ACK, your_ip=192.168.147.154, server_ip=192.168.147.254)

//DHCP在工作过程中涉及到的报文种类及其作用如下：
//
//1、DHCP DISCOVER：客户端开始DHCP过程的第一个报文，是请求IP地址和其它配置参数的广播报文。
//
//2、DHCP OFFER：服务器对DHCP DISCOVER报文的响应，是包含有效IP地址及配置的单播（或广播）报文。
//
//3、DHCP REQUEST：客户端对DHCP OFFER报文的响应，表示接受相关配置。客户端续延IP地址租期时也会发出该报文。
//
//4、DHCP DECLINE：当客户端发现服务器分配的IP地址无法使用（如IP地址冲突时），将发出此报文，通知服务器禁止使用该IP地址。
//
//5、DHCP ACK ：服务器对客户端的DHCP REQUEST报文的确认响应报文。客户端收到此报文后，才真正获得了IP地址和相关的配置信息。
//
//6、DHCP NAK：服务器对客户端的DHCP REQUEST报文的拒绝响应报文。客户端收到此报文后，会重新开始新的DHCP过程。
//
//7、DHCP RELEASE：客户端主动释放服务器分配的IP地址。当服务器收到此报文后，则回收该IP地址，并可以将其分配给其它的客户端。
//
//8、DHCP INFORM：客户端获得IP地址后，发送此报文请求获取服务器的其它一些网络配置信息，如DNS等。
