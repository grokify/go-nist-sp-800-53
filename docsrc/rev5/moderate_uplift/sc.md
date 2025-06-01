# SC - System and Communications Protection

* Controls: 15

## Controls

### SC-2: Separation of System and User Functionality

Separate user functionality, including user interface services, from system management functionality.

System management functionality includes functions that are necessary to administer databases, network components, workstations, or servers. These functions typically require privileged user access. The separation of user functions from system management functions is physical or logical. Organizations may separate system management functions from user functions by using different computers, instances of operating systems, central processing units, or network addresses; by employing virtualization techniques; or some combination of these or other methods. Separation of system management functions from user functions includes web administrative interfaces that employ separate authentication methods for users of any other system resources. Separation of system and user functions may include isolating administrative interfaces on different domains and with additional access controls. The separation of system and user functionality can be achieved by applying the systems security engineering design principles in [SA-8](#sa-8) , including [SA-8(1)](#sa-8.1), [SA-8(3)](#sa-8.3), [SA-8(4)](#sa-8.4), [SA-8(10)](#sa-8.10), [SA-8(12)](#sa-8.12), [SA-8(13)](#sa-8.13), [SA-8(14)](#sa-8.14) , and [SA-8(18)](#sa-8.18).

user functionality, including user interface services, is separated from system management functionality.

System and communications protection policy

procedures addressing application partitioning

system design documentation

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

system developer

Separation of user functionality from system management functionality

### SC-4: Information in Shared System Resources

Prevent unauthorized and unintended information transfer via shared system resources.

Preventing unauthorized and unintended information transfer via shared system resources stops information produced by the actions of prior users or roles (or the actions of processes acting on behalf of prior users or roles) from being available to current users or roles (or current processes acting on behalf of current users or roles) that obtain access to shared system resources after those resources have been released back to the system. Information in shared system resources also applies to encrypted representations of information. In other contexts, control of information in shared system resources is referred to as object reuse and residual information protection. Information in shared system resources does not address information remanence, which refers to the residual representation of data that has been nominally deleted; covert channels (including storage and timing channels), where shared system resources are manipulated to violate information flow restrictions; or components within systems for which there are only single users or roles.

unauthorized information transfer via shared system resources is prevented;

unintended information transfer via shared system resources is prevented.

System and communications protection policy

procedures addressing information protection in shared system resources

system design documentation

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

system developer

Mechanisms preventing the unauthorized and unintended transfer of information via shared system resources

### SC-7 (3): Access Points

Limit the number of external network connections to the system.

Limiting the number of external network connections facilitates monitoring of inbound and outbound communications traffic. The Trusted Internet Connection [DHS TIC](#4f42ee6e-86cc-403b-a51f-76c2b4f81b54) initiative is an example of a federal guideline that requires limits on the number of external network connections. Limiting the number of external network connections to the system is important during transition periods from older to newer technologies (e.g., transitioning from IPv4 to IPv6 network protocols). Such transitions may require implementing the older and newer technologies simultaneously during the transition period and thus increase the number of access points to the system.

the number of external network connections to the system is limited.

System and communications protection policy

procedures addressing boundary protection

system design documentation

boundary protection hardware and software

system architecture and configuration documentation

system configuration settings and associated documentation

communications and network traffic monitoring logs

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

organizational personnel with boundary protection responsibilities

Mechanisms implementing boundary protection capabilities

mechanisms limiting the number of external network connections to the system

### SC-7 (4): External Telecommunications Services

Implement a managed interface for each external telecommunication service;

Establish a traffic flow policy for each managed interface;

Protect the confidentiality and integrity of the information being transmitted across each interface;

Document each exception to the traffic flow policy with a supporting mission or business need and duration of that need;

Review exceptions to the traffic flow policy the frequency at which to review exceptions to traffic flow policy is defined; and remove exceptions that are no longer supported by an explicit mission or business need;

Prevent unauthorized exchange of control plane traffic with external networks;

Publish information to enable remote networks to detect unauthorized control plane traffic from internal networks; and

Filter unauthorized control plane traffic from external networks.

External telecommunications services can provide data and/or voice communications services. Examples of control plane traffic include Border Gateway Protocol (BGP) routing, Domain Name System (DNS), and management protocols. See [SP 800-189](#f5edfe51-d1f2-422e-9b27-5d0e90b49c72) for additional information on the use of the resource public key infrastructure (RPKI) to protect BGP routes and detect unauthorized BGP announcements.

a managed interface is implemented for each external telecommunication service;

a traffic flow policy is established for each managed interface;

the confidentiality of the information being transmitted across each interface is protected;

the integrity of the information being transmitted across each interface is protected;

each exception to the traffic flow policy is documented with a supporting mission or business need and duration of that need;

exceptions to the traffic flow policy are reviewed the frequency at which to review exceptions to traffic flow policy is defined;;

exceptions to the traffic flow policy that are no longer supported by an explicit mission or business need are removed;

unauthorized exchanges of control plan traffic with external networks are prevented;

information is published to enable remote networks to detect unauthorized control plane traffic from internal networks;

unauthorized control plane traffic is filtered from external networks.

System and communications protection policy

traffic flow policy

information flow control policy

procedures addressing boundary protection

system security architecture

system design documentation

boundary protection hardware and software

system architecture and configuration documentation

system configuration settings and associated documentation

records of traffic flow policy exceptions

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

organizational personnel with boundary protection responsibilities

Organizational processes for documenting and reviewing exceptions to the traffic flow policy

organizational processes for removing exceptions to the traffic flow policy

mechanisms implementing boundary protection capabilities

managed interfaces implementing traffic flow policy

### SC-7 (5): Deny by Default — Allow by Exception

Deny network communications traffic by default and allow network communications traffic by exception at managed interfacesand/orfor systems for which network communications traffic is denied by default and network communications traffic is allowed by exception are defined (if selected)..

Denying by default and allowing by exception applies to inbound and outbound network communications traffic. A deny-all, permit-by-exception network communications traffic policy ensures that only those system connections that are essential and approved are allowed. Deny by default, allow by exception also applies to a system that is connected to an external system.

network communications traffic is denied by default at managed interfacesand/orfor systems for which network communications traffic is denied by default and network communications traffic is allowed by exception are defined (if selected).;

network communications traffic is allowed by exception at managed interfacesand/orfor systems for which network communications traffic is denied by default and network communications traffic is allowed by exception are defined (if selected)..

System and communications protection policy

procedures addressing boundary protection

system design documentation

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

system developer

organizational personnel with boundary protection responsibilities

Mechanisms implementing traffic management at managed interfaces

### SC-7 (7): Split Tunneling for Remote Devices

Prevent split tunneling for remote devices connecting to organizational systems unless the split tunnel is securely provisioned using safeguards to securely provision split tunneling are defined;.

Split tunneling is the process of allowing a remote user or device to establish a non-remote connection with a system and simultaneously communicate via some other connection to a resource in an external network. This method of network access enables a user to access remote devices and simultaneously, access uncontrolled networks. Split tunneling might be desirable by remote users to communicate with local system resources, such as printers or file servers. However, split tunneling can facilitate unauthorized external connections, making the system vulnerable to attack and to exfiltration of organizational information. Split tunneling can be prevented by disabling configuration settings that allow such capability in remote devices and by preventing those configuration settings from being configurable by users. Prevention can also be achieved by the detection of split tunneling (or of configuration settings that allow split tunneling) in the remote device, and by prohibiting the connection if the remote device is using split tunneling. A virtual private network (VPN) can be used to securely provision a split tunnel. A securely provisioned VPN includes locking connectivity to exclusive, managed, and named environments, or to a specific set of pre-approved addresses, without user control.

split tunneling is prevented for remote devices connecting to organizational systems unless the split tunnel is securely provisioned using safeguards to securely provision split tunneling are defined;.

System and communications protection policy

procedures addressing boundary protection

system design documentation

system hardware and software

system architecture

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

system developer

organizational personnel with boundary protection responsibilities

Mechanisms implementing boundary protection capabilities

mechanisms supporting/restricting non-remote connections

### SC-7 (8): Route Traffic to Authenticated Proxy Servers

Route internal communications traffic to be routed to external networks is defined; to external networks to which internal communications traffic is to be routed are defined; through authenticated proxy servers at managed interfaces.

External networks are networks outside of organizational control. A proxy server is a server (i.e., system or application) that acts as an intermediary for clients requesting system resources from non-organizational or other organizational servers. System resources that may be requested include files, connections, web pages, or services. Client requests established through a connection to a proxy server are assessed to manage complexity and provide additional protection by limiting direct connectivity. Web content filtering devices are one of the most common proxy servers that provide access to the Internet. Proxy servers can support the logging of Transmission Control Protocol sessions and the blocking of specific Uniform Resource Locators, Internet Protocol addresses, and domain names. Web proxies can be configured with organization-defined lists of authorized and unauthorized websites. Note that proxy servers may inhibit the use of virtual private networks (VPNs) and create the potential for "man-in-the-middle" attacks (depending on the implementation).

 internal communications traffic to be routed to external networks is defined; is routed to external networks to which internal communications traffic is to be routed are defined; through authenticated proxy servers at managed interfaces.

System and communications protection policy

procedures addressing boundary protection

system design documentation

system hardware and software

system architecture

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

system developer

organizational personnel with boundary protection responsibilities

Mechanisms implementing traffic management through authenticated proxy servers at managed interfaces

### SC-8: Transmission Confidentiality and Integrity

Protect the confidentialityand/orintegrity of transmitted information.

Protecting the confidentiality and integrity of transmitted information applies to internal and external networks as well as any system components that can transmit information, including servers, notebook computers, desktop computers, mobile devices, printers, copiers, scanners, facsimile machines, and radios. Unprotected communication paths are exposed to the possibility of interception and modification. Protecting the confidentiality and integrity of information can be accomplished by physical or logical means. Physical protection can be achieved by using protected distribution systems. A protected distribution system is a wireline or fiber-optics telecommunications system that includes terminals and adequate electromagnetic, acoustical, electrical, and physical controls to permit its use for the unencrypted transmission of classified information. Logical protection can be achieved by employing encryption techniques.

Organizations that rely on commercial providers who offer transmission services as commodity services rather than as fully dedicated services may find it difficult to obtain the necessary assurances regarding the implementation of needed controls for transmission confidentiality and integrity. In such situations, organizations determine what types of confidentiality or integrity services are available in standard, commercial telecommunications service packages. If it is not feasible to obtain the necessary controls and assurances of control effectiveness through appropriate contracting vehicles, organizations can implement appropriate compensating controls.

the confidentialityand/orintegrity of transmitted information is/are protected.

System and communications protection policy

procedures addressing transmission confidentiality and integrity

system design documentation

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

system developer

Mechanisms supporting and/or implementing transmission confidentiality and/or integrity

### SC-8 (1): Cryptographic Protection

Implement cryptographic mechanisms to prevent unauthorized disclosure of informationand/ordetect changes to information during transmission.

Encryption protects information from unauthorized disclosure and modification during transmission. Cryptographic mechanisms that protect the confidentiality and integrity of information during transmission include TLS and IPSec. Cryptographic mechanisms used to protect information integrity include cryptographic hash functions that have applications in digital signatures, checksums, and message authentication codes.

cryptographic mechanisms are implemented to prevent unauthorized disclosure of informationand/ordetect changes to information during transmission.

System and communications protection policy

procedures addressing transmission confidentiality and integrity

system design documentation

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

system developer

Cryptographic mechanisms supporting and/or implementing transmission confidentiality and/or integrity

mechanisms supporting and/or implementing alternative physical safeguards

organizational processes for defining and implementing alternative physical safeguards

### SC-10: Network Disconnect

Terminate the network connection associated with a communications session at the end of the session or after a time period of inactivity after which the system terminates a network connection associated with a communication session is defined; of inactivity.

Network disconnect applies to internal and external networks. Terminating network connections associated with specific communications sessions includes de-allocating TCP/IP address or port pairs at the operating system level and de-allocating the networking assignments at the application level if multiple application sessions are using a single operating system-level network connection. Periods of inactivity may be established by organizations and include time periods by type of network access or for specific network accesses.

the network connection associated with a communication session is terminated at the end of the session or after a time period of inactivity after which the system terminates a network connection associated with a communication session is defined; of inactivity.

System and communications protection policy

procedures addressing network disconnect

system design documentation

security plan

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

system developer

Mechanisms supporting and/or implementing a network disconnect capability

### SC-17: Public Key Infrastructure Certificates

Issue public key certificates under an a certificate policy for issuing public key certificates is defined; or obtain public key certificates from an approved service provider; and

Include only approved trust anchors in trust stores or certificate stores managed by the organization.

Public key infrastructure (PKI) certificates are certificates with visibility external to organizational systems and certificates related to the internal operations of systems, such as application-specific time services. In cryptographic systems with a hierarchical structure, a trust anchor is an authoritative source (i.e., a certificate authority) for which trust is assumed and not derived. A root certificate for a PKI system is an example of a trust anchor. A trust store or certificate store maintains a list of trusted root certificates.

public key certificates are issued under a certificate policy for issuing public key certificates is defined; , or public key certificates are obtained from an approved service provider;

only approved trust anchors are included in trust stores or certificate stores managed by the organization.

System and communications protection policy

procedures addressing public key infrastructure certificates

public key certificate policy or policies

public key issuing process

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

organizational personnel with responsibilities for issuing public key certificates

service providers

Mechanisms supporting and/or implementing the management of public key infrastructure certificates

### SC-18: Mobile Code

Define acceptable and unacceptable mobile code and mobile code technologies; and

Authorize, monitor, and control the use of mobile code within the system.

Mobile code includes any program, application, or content that can be transmitted across a network (e.g., embedded in an email, document, or website) and executed on a remote system. Decisions regarding the use of mobile code within organizational systems are based on the potential for the code to cause damage to the systems if used maliciously. Mobile code technologies include Java applets, JavaScript, HTML5, WebGL, and VBScript. Usage restrictions and implementation guidelines apply to both the selection and use of mobile code installed on servers and mobile code downloaded and executed on individual workstations and devices, including notebook computers and smart phones. Mobile code policy and procedures address specific actions taken to prevent the development, acquisition, and introduction of unacceptable mobile code within organizational systems, including requiring mobile code to be digitally signed by a trusted source.

acceptable mobile code is defined;

unacceptable mobile code is defined;

acceptable mobile code technologies are defined;

unacceptable mobile code technologies are defined;

the use of mobile code is authorized within the system;

the use of mobile code is monitored within the system;

the use of mobile code is controlled within the system.

System and communications protection policy

procedures addressing mobile code

mobile code implementation policy and procedures

list of acceptable mobile code and mobile code technologies

list of unacceptable mobile code and mobile technologies

authorization records

system monitoring records

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

organizational personnel with responsibilities for managing mobile code

Organizational process for authorizing, monitoring, and controlling mobile code

mechanisms supporting and/or implementing the management of mobile code

mechanisms supporting and/or implementing the monitoring of mobile code

### SC-23: Session Authenticity

Protect the authenticity of communications sessions.

Protecting session authenticity addresses communications protection at the session level, not at the packet level. Such protection establishes grounds for confidence at both ends of communications sessions in the ongoing identities of other parties and the validity of transmitted information. Authenticity protection includes protecting against "man-in-the-middle" attacks, session hijacking, and the insertion of false information into sessions.

the authenticity of communication sessions is protected.

System and communications protection policy

procedures addressing session authenticity

system design documentation

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

Mechanisms supporting and/or implementing session authenticity

### SC-28: Protection of Information at Rest

Protect the confidentialityand/orintegrity of the following information at rest: information at rest requiring protection is defined;.

Information at rest refers to the state of information when it is not in process or in transit and is located on system components. Such components include internal or external hard disk drives, storage area network devices, or databases. However, the focus of protecting information at rest is not on the type of storage device or frequency of access but rather on the state of the information. Information at rest addresses the confidentiality and integrity of information and covers user information and system information. System-related information that requires protection includes configurations or rule sets for firewalls, intrusion detection and prevention systems, filtering routers, and authentication information. Organizations may employ different mechanisms to achieve confidentiality and integrity protections, including the use of cryptographic mechanisms and file share scanning. Integrity protection can be achieved, for example, by implementing write-once-read-many (WORM) technologies. When adequate protection of information at rest cannot otherwise be achieved, organizations may employ other controls, including frequent scanning to identify malicious code at rest and secure offline storage in lieu of online storage.

the confidentialityand/orintegrity of information at rest requiring protection is defined; is/are protected.

System and communications protection policy

procedures addressing the protection of information at rest

system design documentation

system configuration settings and associated documentation

cryptographic mechanisms and associated configuration documentation

list of information at rest requiring confidentiality and integrity protections

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

system developer

Mechanisms supporting and/or implementing confidentiality and integrity protections for information at rest

### SC-28 (1): Cryptographic Protection

Implement cryptographic mechanisms to prevent unauthorized disclosure and modification of the following information at rest on system components or media requiring cryptographic protection is/are defined;: information requiring cryptographic protection is defined;.

The selection of cryptographic mechanisms is based on the need to protect the confidentiality and integrity of organizational information. The strength of mechanism is commensurate with the security category or classification of the information. Organizations have the flexibility to encrypt information on system components or media or encrypt data structures, including files, records, or fields.

cryptographic mechanisms are implemented to prevent unauthorized disclosure of information requiring cryptographic protection is defined; at rest on system components or media requiring cryptographic protection is/are defined;;

cryptographic mechanisms are implemented to prevent unauthorized modification of information requiring cryptographic protection is defined; at rest on system components or media requiring cryptographic protection is/are defined;.

System and communications protection policy

procedures addressing the protection of information at rest

system design documentation

system configuration settings and associated documentation

cryptographic mechanisms and associated configuration documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

system developer

Cryptographic mechanisms implementing confidentiality and integrity protections for information at rest

