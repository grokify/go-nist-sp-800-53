# IA - Identification and Authentication

* Controls: 8

## Controls

### IA-3: Device Identification and Authentication

Uniquely identify and authenticate {{ insert: param, ia-03_odp.01 }} before establishing a {{ insert: param, ia-03_odp.02 }} connection.

Devices that require unique device-to-device identification and authentication are defined by type, device, or a combination of type and device. Organization-defined device types include devices that are not owned by the organization. Systems use shared known information (e.g., Media Access Control [MAC], Transmission Control Protocol/Internet Protocol [TCP/IP] addresses) for device identification or organizational authentication solutions (e.g., Institute of Electrical and Electronics Engineers (IEEE) 802.1x and Extensible Authentication Protocol [EAP], RADIUS server with EAP-Transport Layer Security [TLS] authentication, Kerberos) to identify and authenticate devices on local and wide area networks. Organizations determine the required strength of authentication mechanisms based on the security categories of systems and mission or business requirements. Because of the challenges of implementing device authentication on a large scale, organizations can restrict the application of the control to a limited number/type of devices based on mission or business needs.

 {{ insert: param, ia-03_odp.01 }} are uniquely identified and authenticated before establishing a {{ insert: param, ia-03_odp.02 }} connection.

Identification and authentication policy

system security plan

procedures addressing device identification and authentication

system design documentation

list of devices requiring unique identification and authentication

device connection reports

system configuration settings and associated documentation

other relevant documents or records

Organizational personnel with operational responsibilities for device identification and authentication

organizational personnel with information security responsibilities

system/network administrators

system developers

Mechanisms supporting and/or implementing device identification and authentication capabilities

### IA-4 (4): Identify User Status

Manage individual identifiers by uniquely identifying each individual as {{ insert: param, ia-04.04_odp }}.

Characteristics that identify the status of individuals include contractors, foreign nationals, and non-organizational users. Identifying the status of individuals by these characteristics provides additional information about the people with whom organizational personnel are communicating. For example, it might be useful for a government employee to know that one of the individuals on an email message is a contractor.

individual identifiers are managed by uniquely identifying each individual as {{ insert: param, ia-04.04_odp }}.

Identification and authentication policy

system security plan

procedures addressing identifier management

procedures addressing account management

list of characteristics identifying individual status

other relevant documents or records

Organizational personnel with identifier management responsibilities

organizational personnel with information security responsibilities

system/network administrators

Mechanisms supporting and/or implementing identifier management

### IA-5 (2): Public Key-based Authentication

For public key-based authentication:

Enforce authorized access to the corresponding private key; and

Map the authenticated identity to the account of the individual or group; and

When public key infrastructure (PKI) is used:

Validate certificates by constructing and verifying a certification path to an accepted trust anchor, including checking certificate status information; and

Implement a local cache of revocation data to support path discovery and validation.

Public key cryptography is a valid authentication mechanism for individuals, machines, and devices. For PKI solutions, status information for certification paths includes certificate revocation lists or certificate status protocol responses. For PIV cards, certificate validation involves the construction and verification of a certification path to the Common Policy Root trust anchor, which includes certificate policy processing. Implementing a local cache of revocation data to support path discovery and validation also supports system availability in situations where organizations are unable to access revocation information via the network.

authorized access to the corresponding private key is enforced for public key-based authentication;

the authenticated identity is mapped to the account of the individual or group for public key-based authentication;

when public key infrastructure (PKI) is used, certificates are validated by constructing and verifying a certification path to an accepted trust anchor, including checking certificate status information;

when public key infrastructure (PKI) is used, a local cache of revocation data is implemented to support path discovery and validation.

Identification and authentication policy

procedures addressing authenticator management

system security plan

system design documentation

system configuration settings and associated documentation

PKI certification validation records

PKI certification revocation lists

other relevant documents or records

Organizational personnel with PKI-based, authenticator management responsibilities

organizational personnel with information security responsibilities

system/network administrators

system developers

Mechanisms supporting and/or implementing PKI-based, authenticator management capability

### IA-5 (6): Protection of Authenticators

Protect authenticators commensurate with the security category of the information to which use of the authenticator permits access.

For systems that contain multiple security categories of information without reliable physical or logical separation between categories, authenticators used to grant access to the systems are protected commensurate with the highest security category of information on the systems. Security categories of information are determined as part of the security categorization process.

authenticators are protected commensurate with the security category of the information to which use of the authenticator permits access.

Identification and authentication policy

procedures addressing authenticator management

security categorization documentation for the system

security assessments of authenticator protections

risk assessment results

system security plan

other relevant documents or records

Organizational personnel with authenticator management responsibilities

organizational personnel implementing and/or maintaining authenticator protections

organizational personnel with information security responsibilities

system/network administrators

Mechanisms supporting and/or implementing authenticator management capability

mechanisms protecting authenticators

### IA-12: Identity Proofing

Identity proof users that require accounts for logical access to systems based on appropriate identity assurance level requirements as specified in applicable standards and guidelines;

Resolve user identities to a unique individual; and

Collect, validate, and verify identity evidence.

Identity proofing is the process of collecting, validating, and verifying a user’s identity information for the purposes of establishing credentials for accessing a system. Identity proofing is intended to mitigate threats to the registration of users and the establishment of their accounts. Standards and guidelines specifying identity assurance levels for identity proofing include [SP 800-63-3](#737513fa-6758-403f-831d-5ddab5e23cb3) and [SP 800-63A](#9099ed2c-922a-493d-bcb4-d896192243ff) . Organizations may be subject to laws, executive orders, directives, regulations, or policies that address the collection of identity evidence. Organizational personnel consult with the senior agency official for privacy and legal counsel regarding such requirements.

users who require accounts for logical access to systems based on appropriate identity assurance level requirements as specified in applicable standards and guidelines are identity proofed;

user identities are resolved to a unique individual;

identity evidence is collected;

identity evidence is validated;

identity evidence is verified.

Identification and authentication policy

procedures addressing identity proofing

system security plan

privacy plan

other relevant documents or records

Organizational personnel with system operations responsibilities

organizational personnel with information security and privacy responsibilities

legal counsel

system/network administrators

system developers

organizational personnel with identification and authentication responsibilities

Mechanisms supporting and/or implementing identification and authentication capabilities

### IA-12 (2): Identity Evidence

Require evidence of individual identification be presented to the registration authority.

Identity evidence, such as documentary evidence or a combination of documents and biometrics, reduces the likelihood of individuals using fraudulent identification to establish an identity or at least increases the work factor of potential adversaries. The forms of acceptable evidence are consistent with the risks to the systems, roles, and privileges associated with the user’s account.

evidence of individual identification is presented to the registration authority.

Identification and authentication policy

procedures addressing identity proofing

system security plan

other relevant documents or records

Organizational personnel with system operations responsibilities

organizational personnel with information security responsibilities

system/network administrators

system developers

organizational personnel with identification and authentication responsibilities

Mechanisms supporting and/or implementing identification and authentication capabilities

### IA-12 (3): Identity Evidence Validation and Verification

Require that the presented identity evidence be validated and verified through {{ insert: param, ia-12.03_odp }}.

Validation and verification of identity evidence increases the assurance that accounts and identifiers are being established for the correct user and authenticators are being bound to that user. Validation refers to the process of confirming that the evidence is genuine and authentic, and the data contained in the evidence is correct, current, and related to an individual. Verification confirms and establishes a linkage between the claimed identity and the actual existence of the user presenting the evidence. Acceptable methods for validating and verifying identity evidence are consistent with the risks to the systems, roles, and privileges associated with the users account.

the presented identity evidence is validated and verified through {{ insert: param, ia-12.03_odp }}.

Identification and authentication policy

procedures addressing identity proofing

system security plan

other relevant documents or records

Organizational personnel with system operations responsibilities

organizational personnel with information security responsibilities

system/network administrators

system developers

organizational personnel with identification and authentication responsibilities

Mechanisms supporting and/or implementing identification and authentication capabilities

### IA-12 (5): Address Confirmation

Require that a {{ insert: param, ia-12.05_odp }} be delivered through an out-of-band channel to verify the users address (physical or digital) of record.

To make it more difficult for adversaries to pose as legitimate users during the identity proofing process, organizations can use out-of-band methods to ensure that the individual associated with an address of record is the same individual that participated in the registration. Confirmation can take the form of a temporary enrollment code or a notice of proofing. The delivery address for these artifacts is obtained from records and not self-asserted by the user. The address can include a physical or digital address. A home address is an example of a physical address. Email addresses and telephone numbers are examples of digital addresses.

a {{ insert: param, ia-12.05_odp }} is delivered through an out-of-band channel to verify the user’s address (physical or digital) of record.

Identification and authentication policy

procedures addressing identity proofing

system security plan

other relevant documents or records

Organizational personnel with system operations responsibilities

organizational personnel with information security responsibilities

system/network administrators

system developers

organizational personnel with identification and authentication responsibilities

Mechanisms supporting and/or implementing identification and authentication capabilities

