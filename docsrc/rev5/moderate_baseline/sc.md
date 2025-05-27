# sc - System and Communications Protection

* Controls: 25

## Controls

### sc-1: Policy and Procedures

Develop, document, and disseminate to {{ insert: param, sc-1_prm_1 }}:

 {{ insert: param, sc-01_odp.03 }} system and communications protection policy that:

Addresses purpose, scope, roles, responsibilities, management commitment, coordination among organizational entities, and compliance; and

Is consistent with applicable laws, executive orders, directives, regulations, policies, standards, and guidelines; and

Procedures to facilitate the implementation of the system and communications protection policy and the associated system and communications protection controls;

Designate an {{ insert: param, sc-01_odp.04 }} to manage the development, documentation, and dissemination of the system and communications protection policy and procedures; and

Review and update the current system and communications protection:

Policy {{ insert: param, sc-01_odp.05 }} and following {{ insert: param, sc-01_odp.06 }} ; and

Procedures {{ insert: param, sc-01_odp.07 }} and following {{ insert: param, sc-01_odp.08 }}.

System and communications protection policy and procedures address the controls in the SC family that are implemented within systems and organizations. The risk management strategy is an important factor in establishing such policies and procedures. Policies and procedures contribute to security and privacy assurance. Therefore, it is important that security and privacy programs collaborate on the development of system and communications protection policy and procedures. Security and privacy program policies and procedures at the organization level are preferable, in general, and may obviate the need for mission- or system-specific policies and procedures. The policy can be included as part of the general security and privacy policy or be represented by multiple policies that reflect the complex nature of organizations. Procedures can be established for security and privacy programs, for mission or business processes, and for systems, if needed. Procedures describe how the policies or controls are implemented and can be directed at the individual or role that is the object of the procedure. Procedures can be documented in system security and privacy plans or in one or more separate documents. Events that may precipitate an update to system and communications protection policy and procedures include assessment or audit findings, security incidents or breaches, or changes in applicable laws, executive orders, directives, regulations, policies, standards, and guidelines. Simply restating controls does not constitute an organizational policy or procedure.

a system and communications protection policy is developed and documented;

the system and communications protection policy is disseminated to {{ insert: param, sc-01_odp.01 }};

system and communications protection procedures to facilitate the implementation of the system and communications protection policy and associated system and communications protection controls are developed and documented;

the system and communications protection procedures are disseminated to {{ insert: param, sc-01_odp.02 }};

the {{ insert: param, sc-01_odp.03 }} system and communications protection policy addresses purpose;

the {{ insert: param, sc-01_odp.03 }} system and communications protection policy addresses scope;

the {{ insert: param, sc-01_odp.03 }} system and communications protection policy addresses roles;

the {{ insert: param, sc-01_odp.03 }} system and communications protection policy addresses responsibilities;

the {{ insert: param, sc-01_odp.03 }} system and communications protection policy addresses management commitment;

the {{ insert: param, sc-01_odp.03 }} system and communications protection policy addresses coordination among organizational entities;

the {{ insert: param, sc-01_odp.03 }} system and communications protection policy addresses compliance;

the {{ insert: param, sc-01_odp.03 }} system and communications protection policy is consistent with applicable laws, Executive Orders, directives, regulations, policies, standards, and guidelines;

the {{ insert: param, sc-01_odp.04 }} is designated to manage the development, documentation, and dissemination of the system and communications protection policy and procedures;

the current system and communications protection policy is reviewed and updated {{ insert: param, sc-01_odp.05 }};

the current system and communications protection policy is reviewed and updated following {{ insert: param, sc-01_odp.06 }};

the current system and communications protection procedures are reviewed and updated {{ insert: param, sc-01_odp.07 }};

the current system and communications protection procedures are reviewed and updated following {{ insert: param, sc-01_odp.08 }}.

System and communications protection policy

system and communications protection procedures

system security plan

privacy plan

risk management strategy documentation

audit findings

other relevant documents or records

Organizational personnel with system and communications protection responsibilities

organizational personnel with information security and privacy responsibilities

### sc-2: Separation of System and User Functionality

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

### sc-4: Information in Shared System Resources

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

### sc-5: Denial-of-service Protection

 {{ insert: param, sc-05_odp.02 }} the effects of the following types of denial-of-service events: {{ insert: param, sc-05_odp.01 }} ; and

Employ the following controls to achieve the denial-of-service objective: {{ insert: param, sc-05_odp.03 }}.

Denial-of-service events may occur due to a variety of internal and external causes, such as an attack by an adversary or a lack of planning to support organizational needs with respect to capacity and bandwidth. Such attacks can occur across a wide range of network protocols (e.g., IPv4, IPv6). A variety of technologies are available to limit or eliminate the origination and effects of denial-of-service events. For example, boundary protection devices can filter certain types of packets to protect system components on internal networks from being directly affected by or the source of denial-of-service attacks. Employing increased network capacity and bandwidth combined with service redundancy also reduces the susceptibility to denial-of-service events.

the effects of {{ insert: param, sc-05_odp.01 }} are {{ insert: param, sc-05_odp.02 }};

 {{ insert: param, sc-05_odp.03 }} are employed to achieve the denial-of-service protection objective.

System and communications protection policy

procedures addressing denial-of-service protection

system design documentation

list of denial-of-service attacks requiring employment of security safeguards to protect against or limit effects of such attacks

list of security safeguards protecting against or limiting the effects of denial-of-service attacks

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

organizational personnel with incident response responsibilities

system developer

Mechanisms protecting against or limiting the effects of denial-of-service attacks

### sc-7: Boundary Protection

Monitor and control communications at the external managed interfaces to the system and at key internal managed interfaces within the system;

Implement subnetworks for publicly accessible system components that are {{ insert: param, sc-07_odp }} separated from internal organizational networks; and

Connect to external networks or systems only through managed interfaces consisting of boundary protection devices arranged in accordance with an organizational security and privacy architecture.

Managed interfaces include gateways, routers, firewalls, guards, network-based malicious code analysis, virtualization systems, or encrypted tunnels implemented within a security architecture. Subnetworks that are physically or logically separated from internal networks are referred to as demilitarized zones or DMZs. Restricting or prohibiting interfaces within organizational systems includes restricting external web traffic to designated web servers within managed interfaces, prohibiting external traffic that appears to be spoofing internal addresses, and prohibiting internal traffic that appears to be spoofing external addresses. [SP 800-189](#f5edfe51-d1f2-422e-9b27-5d0e90b49c72) provides additional information on source address validation techniques to prevent ingress and egress of traffic with spoofed addresses. Commercial telecommunications services are provided by network components and consolidated management systems shared by customers. These services may also include third party-provided access lines and other service elements. Such services may represent sources of increased risk despite contract security provisions. Boundary protection may be implemented as a common control for all or part of an organizational network such that the boundary to be protected is greater than a system-specific boundary (i.e., an authorization boundary).

communications at external managed interfaces to the system are monitored;

communications at external managed interfaces to the system are controlled;

communications at key internal managed interfaces within the system are monitored;

communications at key internal managed interfaces within the system are controlled;

subnetworks for publicly accessible system components are {{ insert: param, sc-07_odp }} separated from internal organizational networks;

external networks or systems are only connected to through managed interfaces consisting of boundary protection devices arranged in accordance with an organizational security and privacy architecture.

System and communications protection policy

procedures addressing boundary protection

list of key internal boundaries of the system

system design documentation

boundary protection hardware and software

system configuration settings and associated documentation

enterprise security architecture documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

system developer

organizational personnel with boundary protection responsibilities

Mechanisms implementing boundary protection capabilities

### sc-7.3: Access Points

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

### sc-7.4: External Telecommunications Services

Implement a managed interface for each external telecommunication service;

Establish a traffic flow policy for each managed interface;

Protect the confidentiality and integrity of the information being transmitted across each interface;

Document each exception to the traffic flow policy with a supporting mission or business need and duration of that need;

Review exceptions to the traffic flow policy {{ insert: param, sc-07.04_odp }} and remove exceptions that are no longer supported by an explicit mission or business need;

Prevent unauthorized exchange of control plane traffic with external networks;

Publish information to enable remote networks to detect unauthorized control plane traffic from internal networks; and

Filter unauthorized control plane traffic from external networks.

External telecommunications services can provide data and/or voice communications services. Examples of control plane traffic include Border Gateway Protocol (BGP) routing, Domain Name System (DNS), and management protocols. See [SP 800-189](#f5edfe51-d1f2-422e-9b27-5d0e90b49c72) for additional information on the use of the resource public key infrastructure (RPKI) to protect BGP routes and detect unauthorized BGP announcements.

a managed interface is implemented for each external telecommunication service;

a traffic flow policy is established for each managed interface;

the confidentiality of the information being transmitted across each interface is protected;

the integrity of the information being transmitted across each interface is protected;

each exception to the traffic flow policy is documented with a supporting mission or business need and duration of that need;

exceptions to the traffic flow policy are reviewed {{ insert: param, sc-07.04_odp }};

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

### sc-7.5: Deny by Default — Allow by Exception

Deny network communications traffic by default and allow network communications traffic by exception {{ insert: param, sc-07.05_odp.01 }}.

Denying by default and allowing by exception applies to inbound and outbound network communications traffic. A deny-all, permit-by-exception network communications traffic policy ensures that only those system connections that are essential and approved are allowed. Deny by default, allow by exception also applies to a system that is connected to an external system.

network communications traffic is denied by default {{ insert: param, sc-07.05_odp.01 }};

network communications traffic is allowed by exception {{ insert: param, sc-07.05_odp.01 }}.

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

### sc-7.7: Split Tunneling for Remote Devices

Prevent split tunneling for remote devices connecting to organizational systems unless the split tunnel is securely provisioned using {{ insert: param, sc-07.07_odp }}.

Split tunneling is the process of allowing a remote user or device to establish a non-remote connection with a system and simultaneously communicate via some other connection to a resource in an external network. This method of network access enables a user to access remote devices and simultaneously, access uncontrolled networks. Split tunneling might be desirable by remote users to communicate with local system resources, such as printers or file servers. However, split tunneling can facilitate unauthorized external connections, making the system vulnerable to attack and to exfiltration of organizational information. Split tunneling can be prevented by disabling configuration settings that allow such capability in remote devices and by preventing those configuration settings from being configurable by users. Prevention can also be achieved by the detection of split tunneling (or of configuration settings that allow split tunneling) in the remote device, and by prohibiting the connection if the remote device is using split tunneling. A virtual private network (VPN) can be used to securely provision a split tunnel. A securely provisioned VPN includes locking connectivity to exclusive, managed, and named environments, or to a specific set of pre-approved addresses, without user control.

split tunneling is prevented for remote devices connecting to organizational systems unless the split tunnel is securely provisioned using {{ insert: param, sc-07.07_odp }}.

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

### sc-7.8: Route Traffic to Authenticated Proxy Servers

Route {{ insert: param, sc-07.08_odp.01 }} to {{ insert: param, sc-07.08_odp.02 }} through authenticated proxy servers at managed interfaces.

External networks are networks outside of organizational control. A proxy server is a server (i.e., system or application) that acts as an intermediary for clients requesting system resources from non-organizational or other organizational servers. System resources that may be requested include files, connections, web pages, or services. Client requests established through a connection to a proxy server are assessed to manage complexity and provide additional protection by limiting direct connectivity. Web content filtering devices are one of the most common proxy servers that provide access to the Internet. Proxy servers can support the logging of Transmission Control Protocol sessions and the blocking of specific Uniform Resource Locators, Internet Protocol addresses, and domain names. Web proxies can be configured with organization-defined lists of authorized and unauthorized websites. Note that proxy servers may inhibit the use of virtual private networks (VPNs) and create the potential for "man-in-the-middle" attacks (depending on the implementation).

 {{ insert: param, sc-07.08_odp.01 }} is routed to {{ insert: param, sc-07.08_odp.02 }} through authenticated proxy servers at managed interfaces.

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

### sc-8: Transmission Confidentiality and Integrity

Protect the {{ insert: param, sc-08_odp }} of transmitted information.

Protecting the confidentiality and integrity of transmitted information applies to internal and external networks as well as any system components that can transmit information, including servers, notebook computers, desktop computers, mobile devices, printers, copiers, scanners, facsimile machines, and radios. Unprotected communication paths are exposed to the possibility of interception and modification. Protecting the confidentiality and integrity of information can be accomplished by physical or logical means. Physical protection can be achieved by using protected distribution systems. A protected distribution system is a wireline or fiber-optics telecommunications system that includes terminals and adequate electromagnetic, acoustical, electrical, and physical controls to permit its use for the unencrypted transmission of classified information. Logical protection can be achieved by employing encryption techniques.

Organizations that rely on commercial providers who offer transmission services as commodity services rather than as fully dedicated services may find it difficult to obtain the necessary assurances regarding the implementation of needed controls for transmission confidentiality and integrity. In such situations, organizations determine what types of confidentiality or integrity services are available in standard, commercial telecommunications service packages. If it is not feasible to obtain the necessary controls and assurances of control effectiveness through appropriate contracting vehicles, organizations can implement appropriate compensating controls.

the {{ insert: param, sc-08_odp }} of transmitted information is/are protected.

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

### sc-8.1: Cryptographic Protection

Implement cryptographic mechanisms to {{ insert: param, sc-08.01_odp }} during transmission.

Encryption protects information from unauthorized disclosure and modification during transmission. Cryptographic mechanisms that protect the confidentiality and integrity of information during transmission include TLS and IPSec. Cryptographic mechanisms used to protect information integrity include cryptographic hash functions that have applications in digital signatures, checksums, and message authentication codes.

cryptographic mechanisms are implemented to {{ insert: param, sc-08.01_odp }} during transmission.

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

### sc-10: Network Disconnect

Terminate the network connection associated with a communications session at the end of the session or after {{ insert: param, sc-10_odp }} of inactivity.

Network disconnect applies to internal and external networks. Terminating network connections associated with specific communications sessions includes de-allocating TCP/IP address or port pairs at the operating system level and de-allocating the networking assignments at the application level if multiple application sessions are using a single operating system-level network connection. Periods of inactivity may be established by organizations and include time periods by type of network access or for specific network accesses.

the network connection associated with a communication session is terminated at the end of the session or after {{ insert: param, sc-10_odp }} of inactivity.

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

### sc-12: Cryptographic Key Establishment and Management

Establish and manage cryptographic keys when cryptography is employed within the system in accordance with the following key management requirements: {{ insert: param, sc-12_odp }}.

Cryptographic key management and establishment can be performed using manual procedures or automated mechanisms with supporting manual procedures. Organizations define key management requirements in accordance with applicable laws, executive orders, directives, regulations, policies, standards, and guidelines and specify appropriate options, parameters, and levels. Organizations manage trust stores to ensure that only approved trust anchors are part of such trust stores. This includes certificates with visibility external to organizational systems and certificates related to the internal operations of systems. [NIST CMVP](#1acdc775-aafb-4d11-9341-dc6a822e9d38) and [NIST CAVP](#84dc1b0c-acb7-4269-84c4-00dbabacd78c) provide additional information on validated cryptographic modules and algorithms that can be used in cryptographic key management and establishment.

cryptographic keys are established when cryptography is employed within the system in accordance with {{ insert: param, sc-12_odp }};

cryptographic keys are managed when cryptography is employed within the system in accordance with {{ insert: param, sc-12_odp }}.

System and communications protection policy

procedures addressing cryptographic key establishment and management

system design documentation

cryptographic mechanisms

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

organizational personnel with responsibilities for cryptographic key establishment and/or management

Mechanisms supporting and/or implementing cryptographic key establishment and management

### sc-13: Cryptographic Protection

Determine the {{ insert: param, sc-13_odp.01 }} ; and

Implement the following types of cryptography required for each specified cryptographic use: {{ insert: param, sc-13_odp.02 }}.

Cryptography can be employed to support a variety of security solutions, including the protection of classified information and controlled unclassified information, the provision and implementation of digital signatures, and the enforcement of information separation when authorized individuals have the necessary clearances but lack the necessary formal access approvals. Cryptography can also be used to support random number and hash generation. Generally applicable cryptographic standards include FIPS-validated cryptography and NSA-approved cryptography. For example, organizations that need to protect classified information may specify the use of NSA-approved cryptography. Organizations that need to provision and implement digital signatures may specify the use of FIPS-validated cryptography. Cryptography is implemented in accordance with applicable laws, executive orders, directives, regulations, policies, standards, and guidelines.

 {{ insert: param, sc-13_odp.01 }} are identified;

 {{ insert: param, sc-13_odp.02 }} for each specified cryptographic use (defined in SC-13_ODP[01]) are implemented.

System and communications protection policy

procedures addressing cryptographic protection

system design documentation

system configuration settings and associated documentation

cryptographic module validation certificates

list of FIPS-validated cryptographic modules

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

system developer

organizational personnel with responsibilities for cryptographic protection

Mechanisms supporting and/or implementing cryptographic protection

### sc-15: Collaborative Computing Devices and Applications

Prohibit remote activation of collaborative computing devices and applications with the following exceptions: {{ insert: param, sc-15_odp }} ; and

Provide an explicit indication of use to users physically present at the devices.

Collaborative computing devices and applications include remote meeting devices and applications, networked white boards, cameras, and microphones. The explicit indication of use includes signals to users when collaborative computing devices and applications are activated.

remote activation of collaborative computing devices and applications is prohibited except {{ insert: param, sc-15_odp }};

an explicit indication of use is provided to users physically present at the devices.

System and communications protection policy

procedures addressing collaborative computing

access control policy and procedures

system design documentation

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

system developer

organizational personnel with responsibilities for managing collaborative computing devices

Mechanisms supporting and/or implementing the management of remote activation of collaborative computing devices

mechanisms providing an indication of use of collaborative computing devices

### sc-17: Public Key Infrastructure Certificates

Issue public key certificates under an {{ insert: param, sc-17_odp }} or obtain public key certificates from an approved service provider; and

Include only approved trust anchors in trust stores or certificate stores managed by the organization.

Public key infrastructure (PKI) certificates are certificates with visibility external to organizational systems and certificates related to the internal operations of systems, such as application-specific time services. In cryptographic systems with a hierarchical structure, a trust anchor is an authoritative source (i.e., a certificate authority) for which trust is assumed and not derived. A root certificate for a PKI system is an example of a trust anchor. A trust store or certificate store maintains a list of trusted root certificates.

public key certificates are issued under {{ insert: param, sc-17_odp }} , or public key certificates are obtained from an approved service provider;

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

### sc-18: Mobile Code

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

### sc-20: Secure Name/Address Resolution Service (Authoritative Source)

Provide additional data origin authentication and integrity verification artifacts along with the authoritative name resolution data the system returns in response to external name/address resolution queries; and

Provide the means to indicate the security status of child zones and (if the child supports secure resolution services) to enable verification of a chain of trust among parent and child domains, when operating as part of a distributed, hierarchical namespace.

Providing authoritative source information enables external clients, including remote Internet clients, to obtain origin authentication and integrity verification assurances for the host/service name to network address resolution information obtained through the service. Systems that provide name and address resolution services include domain name system (DNS) servers. Additional artifacts include DNS Security Extensions (DNSSEC) digital signatures and cryptographic keys. Authoritative data includes DNS resource records. The means for indicating the security status of child zones include the use of delegation signer resource records in the DNS. Systems that use technologies other than the DNS to map between host and service names and network addresses provide other means to assure the authenticity and integrity of response data.

additional data origin authentication is provided along with the authoritative name resolution data that the system returns in response to external name/address resolution queries;

integrity verification artifacts are provided along with the authoritative name resolution data that the system returns in response to external name/address resolution queries;

the means to indicate the security status of child zones (and if the child supports secure resolution services) is provided when operating as part of a distributed, hierarchical namespace;

the means to enable verification of a chain of trust among parent and child domains when operating as part of a distributed, hierarchical namespace is provided.

System and communications protection policy

procedures addressing secure name/address resolution services (authoritative source)

system design documentation

system configuration settings and associated documentation

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

organizational personnel with responsibilities for managing DNS

Mechanisms supporting and/or implementing secure name/address resolution services

### sc-21: Secure Name/Address Resolution Service (Recursive or Caching Resolver)

Request and perform data origin authentication and data integrity verification on the name/address resolution responses the system receives from authoritative sources.

Each client of name resolution services either performs this validation on its own or has authenticated channels to trusted validation providers. Systems that provide name and address resolution services for local clients include recursive resolving or caching domain name system (DNS) servers. DNS client resolvers either perform validation of DNSSEC signatures, or clients use authenticated channels to recursive resolvers that perform such validations. Systems that use technologies other than the DNS to map between host and service names and network addresses provide some other means to enable clients to verify the authenticity and integrity of response data.

data origin authentication is requested for the name/address resolution responses that the system receives from authoritative sources;

data origin authentication is performed on the name/address resolution responses that the system receives from authoritative sources;

data integrity verification is requested for the name/address resolution responses that the system receives from authoritative sources;

data integrity verification is performed on the name/address resolution responses that the system receives from authoritative sources.

System and communications protection policy

procedures addressing secure name/address resolution services (recursive or caching resolver)

system design documentation

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

organizational personnel with responsibilities for managing DNS

Mechanisms supporting and/or implementing data origin authentication and data integrity verification for name/address resolution services

### sc-22: Architecture and Provisioning for Name/Address Resolution Service

Ensure the systems that collectively provide name/address resolution service for an organization are fault-tolerant and implement internal and external role separation.

Systems that provide name and address resolution services include domain name system (DNS) servers. To eliminate single points of failure in systems and enhance redundancy, organizations employ at least two authoritative domain name system servers—one configured as the primary server and the other configured as the secondary server. Additionally, organizations typically deploy the servers in two geographically separated network subnetworks (i.e., not located in the same physical facility). For role separation, DNS servers with internal roles only process name and address resolution requests from within organizations (i.e., from internal clients). DNS servers with external roles only process name and address resolution information requests from clients external to organizations (i.e., on external networks, including the Internet). Organizations specify clients that can access authoritative DNS servers in certain roles (e.g., by address ranges and explicit lists).

the systems that collectively provide name/address resolution services for an organization are fault-tolerant;

the systems that collectively provide name/address resolution services for an organization implement internal role separation;

the systems that collectively provide name/address resolution services for an organization implement external role separation.

System and communications protection policy

procedures addressing architecture and provisioning for name/address resolution services

access control policy and procedures

system design documentation

assessment results from independent testing organizations

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

organizational personnel with responsibilities for managing DNS

Mechanisms supporting and/or implementing name/address resolution services for fault tolerance and role separation

### sc-23: Session Authenticity

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

### sc-28: Protection of Information at Rest

Protect the {{ insert: param, sc-28_odp.01 }} of the following information at rest: {{ insert: param, sc-28_odp.02 }}.

Information at rest refers to the state of information when it is not in process or in transit and is located on system components. Such components include internal or external hard disk drives, storage area network devices, or databases. However, the focus of protecting information at rest is not on the type of storage device or frequency of access but rather on the state of the information. Information at rest addresses the confidentiality and integrity of information and covers user information and system information. System-related information that requires protection includes configurations or rule sets for firewalls, intrusion detection and prevention systems, filtering routers, and authentication information. Organizations may employ different mechanisms to achieve confidentiality and integrity protections, including the use of cryptographic mechanisms and file share scanning. Integrity protection can be achieved, for example, by implementing write-once-read-many (WORM) technologies. When adequate protection of information at rest cannot otherwise be achieved, organizations may employ other controls, including frequent scanning to identify malicious code at rest and secure offline storage in lieu of online storage.

the {{ insert: param, sc-28_odp.01 }} of {{ insert: param, sc-28_odp.02 }} is/are protected.

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

### sc-28.1: Cryptographic Protection

Implement cryptographic mechanisms to prevent unauthorized disclosure and modification of the following information at rest on {{ insert: param, sc-28.01_odp.02 }}: {{ insert: param, sc-28.01_odp.01 }}.

The selection of cryptographic mechanisms is based on the need to protect the confidentiality and integrity of organizational information. The strength of mechanism is commensurate with the security category or classification of the information. Organizations have the flexibility to encrypt information on system components or media or encrypt data structures, including files, records, or fields.

cryptographic mechanisms are implemented to prevent unauthorized disclosure of {{ insert: param, sc-28.01_odp.01 }} at rest on {{ insert: param, sc-28.01_odp.02 }};

cryptographic mechanisms are implemented to prevent unauthorized modification of {{ insert: param, sc-28.01_odp.01 }} at rest on {{ insert: param, sc-28.01_odp.02 }}.

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

### sc-39: Process Isolation

Maintain a separate execution domain for each executing system process.

Systems can maintain separate execution domains for each executing process by assigning each process a separate address space. Each system process has a distinct address space so that communication between processes is performed in a manner controlled through the security functions, and one process cannot modify the executing code of another process. Maintaining separate execution domains for executing processes can be achieved, for example, by implementing separate address spaces. Process isolation technologies, including sandboxing or virtualization, logically separate software and firmware from other software, firmware, and data. Process isolation helps limit the access of potentially untrusted software to other system resources. The capability to maintain separate execution domains is available in commercial operating systems that employ multi-state processor technologies.

a separate execution domain is maintained for each executing system process.

System design documentation

system architecture

independent verification and validation documentation

testing and evaluation documentation

other relevant documents or records

System developers/integrators

system security architect

Mechanisms supporting and/or implementing separate execution domains for each executing process

