# SC - System and Communications Protection

* Controls Count: 10
* Controls IDs: SC-1, SC-5, SC-7, SC-12, SC-13, SC-15, SC-20, SC-21, SC-22, SC-39

## Controls

### SC-1: Policy and Procedures

Develop, document, and disseminate to organization-defined personnel or roles:

organization-level, mission/business-process-level, and/or system-level system and communications protection policy that:

Addresses purpose, scope, roles, responsibilities, management commitment, coordination among organizational entities, and compliance; and

Is consistent with applicable laws, executive orders, directives, regulations, policies, standards, and guidelines; and

Procedures to facilitate the implementation of the system and communications protection policy and the associated system and communications protection controls;

Designate an an official to manage the system and communications protection policy and procedures is defined; to manage the development, documentation, and dissemination of the system and communications protection policy and procedures; and

Review and update the current system and communications protection:

Policy the frequency at which the current system and communications protection policy is reviewed and updated is defined; and following events that would require the current system and communications protection policy to be reviewed and updated are defined; ; and

Procedures the frequency at which the current system and communications protection procedures are reviewed and updated is defined; and following events that would require the system and communications protection procedures to be reviewed and updated are defined;.

System and communications protection policy and procedures address the controls in the SC family that are implemented within systems and organizations. The risk management strategy is an important factor in establishing such policies and procedures. Policies and procedures contribute to security and privacy assurance. Therefore, it is important that security and privacy programs collaborate on the development of system and communications protection policy and procedures. Security and privacy program policies and procedures at the organization level are preferable, in general, and may obviate the need for mission- or system-specific policies and procedures. The policy can be included as part of the general security and privacy policy or be represented by multiple policies that reflect the complex nature of organizations. Procedures can be established for security and privacy programs, for mission or business processes, and for systems, if needed. Procedures describe how the policies or controls are implemented and can be directed at the individual or role that is the object of the procedure. Procedures can be documented in system security and privacy plans or in one or more separate documents. Events that may precipitate an update to system and communications protection policy and procedures include assessment or audit findings, security incidents or breaches, or changes in applicable laws, executive orders, directives, regulations, policies, standards, and guidelines. Simply restating controls does not constitute an organizational policy or procedure.

a system and communications protection policy is developed and documented;

the system and communications protection policy is disseminated to personnel or roles to whom the system and communications protection policy is to be disseminated is/are defined;;

system and communications protection procedures to facilitate the implementation of the system and communications protection policy and associated system and communications protection controls are developed and documented;

the system and communications protection procedures are disseminated to personnel or roles to whom the system and communications protection procedures are to be disseminated is/are defined;;

the organization-level, mission/business-process-level, and/or system-level system and communications protection policy addresses purpose;

the organization-level, mission/business-process-level, and/or system-level system and communications protection policy addresses scope;

the organization-level, mission/business-process-level, and/or system-level system and communications protection policy addresses roles;

the organization-level, mission/business-process-level, and/or system-level system and communications protection policy addresses responsibilities;

the organization-level, mission/business-process-level, and/or system-level system and communications protection policy addresses management commitment;

the organization-level, mission/business-process-level, and/or system-level system and communications protection policy addresses coordination among organizational entities;

the organization-level, mission/business-process-level, and/or system-level system and communications protection policy addresses compliance;

the organization-level, mission/business-process-level, and/or system-level system and communications protection policy is consistent with applicable laws, Executive Orders, directives, regulations, policies, standards, and guidelines;

the an official to manage the system and communications protection policy and procedures is defined; is designated to manage the development, documentation, and dissemination of the system and communications protection policy and procedures;

the current system and communications protection policy is reviewed and updated the frequency at which the current system and communications protection policy is reviewed and updated is defined;;

the current system and communications protection policy is reviewed and updated following events that would require the current system and communications protection policy to be reviewed and updated are defined;;

the current system and communications protection procedures are reviewed and updated the frequency at which the current system and communications protection procedures are reviewed and updated is defined;;

the current system and communications protection procedures are reviewed and updated following events that would require the system and communications protection procedures to be reviewed and updated are defined;.

System and communications protection policy

system and communications protection procedures

system security plan

privacy plan

risk management strategy documentation

audit findings

other relevant documents or records

Organizational personnel with system and communications protection responsibilities

organizational personnel with information security and privacy responsibilities

### SC-5: Denial-of-service Protection

protect againstorlimit the effects of the following types of denial-of-service events: types of denial-of-service events to be protected against or limited are defined; ; and

Employ the following controls to achieve the denial-of-service objective: controls to achieve the denial-of-service objective by type of denial-of-service event are defined;.

Denial-of-service events may occur due to a variety of internal and external causes, such as an attack by an adversary or a lack of planning to support organizational needs with respect to capacity and bandwidth. Such attacks can occur across a wide range of network protocols (e.g., IPv4, IPv6). A variety of technologies are available to limit or eliminate the origination and effects of denial-of-service events. For example, boundary protection devices can filter certain types of packets to protect system components on internal networks from being directly affected by or the source of denial-of-service attacks. Employing increased network capacity and bandwidth combined with service redundancy also reduces the susceptibility to denial-of-service events.

the effects of types of denial-of-service events to be protected against or limited are defined; are protect againstorlimit;

 controls to achieve the denial-of-service objective by type of denial-of-service event are defined; are employed to achieve the denial-of-service protection objective.

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

### SC-7: Boundary Protection

Monitor and control communications at the external managed interfaces to the system and at key internal managed interfaces within the system;

Implement subnetworks for publicly accessible system components that are physicallyorlogically separated from internal organizational networks; and

Connect to external networks or systems only through managed interfaces consisting of boundary protection devices arranged in accordance with an organizational security and privacy architecture.

Managed interfaces include gateways, routers, firewalls, guards, network-based malicious code analysis, virtualization systems, or encrypted tunnels implemented within a security architecture. Subnetworks that are physically or logically separated from internal networks are referred to as demilitarized zones or DMZs. Restricting or prohibiting interfaces within organizational systems includes restricting external web traffic to designated web servers within managed interfaces, prohibiting external traffic that appears to be spoofing internal addresses, and prohibiting internal traffic that appears to be spoofing external addresses. [SP 800-189](#f5edfe51-d1f2-422e-9b27-5d0e90b49c72) provides additional information on source address validation techniques to prevent ingress and egress of traffic with spoofed addresses. Commercial telecommunications services are provided by network components and consolidated management systems shared by customers. These services may also include third party-provided access lines and other service elements. Such services may represent sources of increased risk despite contract security provisions. Boundary protection may be implemented as a common control for all or part of an organizational network such that the boundary to be protected is greater than a system-specific boundary (i.e., an authorization boundary).

communications at external managed interfaces to the system are monitored;

communications at external managed interfaces to the system are controlled;

communications at key internal managed interfaces within the system are monitored;

communications at key internal managed interfaces within the system are controlled;

subnetworks for publicly accessible system components are physicallyorlogically separated from internal organizational networks;

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

### SC-12: Cryptographic Key Establishment and Management

Establish and manage cryptographic keys when cryptography is employed within the system in accordance with the following key management requirements: requirements for key generation, distribution, storage, access, and destruction are defined;.

Cryptographic key management and establishment can be performed using manual procedures or automated mechanisms with supporting manual procedures. Organizations define key management requirements in accordance with applicable laws, executive orders, directives, regulations, policies, standards, and guidelines and specify appropriate options, parameters, and levels. Organizations manage trust stores to ensure that only approved trust anchors are part of such trust stores. This includes certificates with visibility external to organizational systems and certificates related to the internal operations of systems. [NIST CMVP](#1acdc775-aafb-4d11-9341-dc6a822e9d38) and [NIST CAVP](#84dc1b0c-acb7-4269-84c4-00dbabacd78c) provide additional information on validated cryptographic modules and algorithms that can be used in cryptographic key management and establishment.

cryptographic keys are established when cryptography is employed within the system in accordance with requirements for key generation, distribution, storage, access, and destruction are defined;;

cryptographic keys are managed when cryptography is employed within the system in accordance with requirements for key generation, distribution, storage, access, and destruction are defined;.

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

### SC-13: Cryptographic Protection

Determine the cryptographic uses are defined; ; and

Implement the following types of cryptography required for each specified cryptographic use: types of cryptography for each specified cryptographic use are defined;.

Cryptography can be employed to support a variety of security solutions, including the protection of classified information and controlled unclassified information, the provision and implementation of digital signatures, and the enforcement of information separation when authorized individuals have the necessary clearances but lack the necessary formal access approvals. Cryptography can also be used to support random number and hash generation. Generally applicable cryptographic standards include FIPS-validated cryptography and NSA-approved cryptography. For example, organizations that need to protect classified information may specify the use of NSA-approved cryptography. Organizations that need to provision and implement digital signatures may specify the use of FIPS-validated cryptography. Cryptography is implemented in accordance with applicable laws, executive orders, directives, regulations, policies, standards, and guidelines.

cryptographic uses are defined; are identified;

 types of cryptography for each specified cryptographic use are defined; for each specified cryptographic use (defined in SC-13_ODP[01]) are implemented.

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

### SC-15: Collaborative Computing Devices and Applications

Prohibit remote activation of collaborative computing devices and applications with the following exceptions: exceptions where remote activation is to be allowed are defined; ; and

Provide an explicit indication of use to users physically present at the devices.

Collaborative computing devices and applications include remote meeting devices and applications, networked white boards, cameras, and microphones. The explicit indication of use includes signals to users when collaborative computing devices and applications are activated.

remote activation of collaborative computing devices and applications is prohibited except exceptions where remote activation is to be allowed are defined;;

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

### SC-20: Secure Name/Address Resolution Service (Authoritative Source)

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

### SC-21: Secure Name/Address Resolution Service (Recursive or Caching Resolver)

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

### SC-22: Architecture and Provisioning for Name/Address Resolution Service

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

### SC-39: Process Isolation

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

