# IA - Identification and Authentication

* Controls Count: 24
* Controls IDs: IA-1, IA-2, IA-2 (1), IA-2 (2), IA-2 (8), IA-2 (12), IA-3, IA-4, IA-4 (4), IA-5, IA-5 (1), IA-5 (2), IA-5 (6), IA-6, IA-7, IA-8, IA-8 (1), IA-8 (2), IA-8 (4), IA-11, IA-12, IA-12 (2), IA-12 (3), IA-12 (5)

## Controls

### IA-1: Policy and Procedures

Develop, document, and disseminate to organization-defined personnel or roles:

organization-level, mission/business process-level, and/or system-level identification and authentication policy that:

Addresses purpose, scope, roles, responsibilities, management commitment, coordination among organizational entities, and compliance; and

Is consistent with applicable laws, executive orders, directives, regulations, policies, standards, and guidelines; and

Procedures to facilitate the implementation of the identification and authentication policy and the associated identification and authentication controls;

Designate an an official to manage the identification and authentication policy and procedures is defined; to manage the development, documentation, and dissemination of the identification and authentication policy and procedures; and

Review and update the current identification and authentication:

Policy the frequency at which the current identification and authentication policy is reviewed and updated is defined; and following events that would require the current identification and authentication policy to be reviewed and updated are defined; ; and

Procedures the frequency at which the current identification and authentication procedures are reviewed and updated is defined; and following events that would require identification and authentication procedures to be reviewed and updated are defined;.

Identification and authentication policy and procedures address the controls in the IA family that are implemented within systems and organizations. The risk management strategy is an important factor in establishing such policies and procedures. Policies and procedures contribute to security and privacy assurance. Therefore, it is important that security and privacy programs collaborate on the development of identification and authentication policy and procedures. Security and privacy program policies and procedures at the organization level are preferable, in general, and may obviate the need for mission- or system-specific policies and procedures. The policy can be included as part of the general security and privacy policy or be represented by multiple policies that reflect the complex nature of organizations. Procedures can be established for security and privacy programs, for mission or business processes, and for systems, if needed. Procedures describe how the policies or controls are implemented and can be directed at the individual or role that is the object of the procedure. Procedures can be documented in system security and privacy plans or in one or more separate documents. Events that may precipitate an update to identification and authentication policy and procedures include assessment or audit findings, security incidents or breaches, or changes in applicable laws, executive orders, directives, regulations, policies, standards, and guidelines. Simply restating controls does not constitute an organizational policy or procedure.

an identification and authentication policy is developed and documented;

the identification and authentication policy is disseminated to personnel or roles to whom the identification and authentication policy is to be disseminated are defined;;

identification and authentication procedures to facilitate the implementation of the identification and authentication policy and associated identification and authentication controls are developed and documented;

the identification and authentication procedures are disseminated to personnel or roles to whom the identification and authentication procedures are to be disseminated is/are defined;;

the organization-level, mission/business process-level, and/or system-level identification and authentication policy addresses purpose;

the organization-level, mission/business process-level, and/or system-level identification and authentication policy addresses scope;

the organization-level, mission/business process-level, and/or system-level identification and authentication policy addresses roles;

the organization-level, mission/business process-level, and/or system-level identification and authentication policy addresses responsibilities;

the organization-level, mission/business process-level, and/or system-level identification and authentication policy addresses management commitment;

the organization-level, mission/business process-level, and/or system-level identification and authentication policy addresses coordination among organizational entities;

the organization-level, mission/business process-level, and/or system-level identification and authentication policy addresses compliance;

the organization-level, mission/business process-level, and/or system-level identification and authentication policy is consistent with applicable laws, executive orders, directives, regulations, policies, standards, and guidelines;

the an official to manage the identification and authentication policy and procedures is defined; is designated to manage the development, documentation, and dissemination of the identification and authentication policy and procedures;

the current identification and authentication policy is reviewed and updated the frequency at which the current identification and authentication policy is reviewed and updated is defined;;

the current identification and authentication policy is reviewed and updated following events that would require the current identification and authentication policy to be reviewed and updated are defined;;

the current identification and authentication procedures are reviewed and updated the frequency at which the current identification and authentication procedures are reviewed and updated is defined;;

the current identification and authentication procedures are reviewed and updated following events that would require identification and authentication procedures to be reviewed and updated are defined;.

Identification and authentication policy and procedures

system security plan

privacy plan

risk management strategy documentation

list of events requiring identification and authentication procedures to be reviewed and updated (e.g., audit findings)

other relevant documents or records

Organizational personnel with identification and authentication responsibilities

organizational personnel with information security and privacy responsibilities

### IA-2: Identification and Authentication (Organizational Users)

Uniquely identify and authenticate organizational users and associate that unique identification with processes acting on behalf of those users.

Organizations can satisfy the identification and authentication requirements by complying with the requirements in [HSPD 12](#f16e438e-7114-4144-bfe2-2dfcad8cb2d0) . Organizational users include employees or individuals who organizations consider to have an equivalent status to employees (e.g., contractors and guest researchers). Unique identification and authentication of users applies to all accesses other than those that are explicitly identified in [AC-14](#ac-14) and that occur through the authorized use of group authenticators without individual authentication. Since processes execute on behalf of groups and roles, organizations may require unique identification of individuals in group accounts or for detailed accountability of individual activity.

Organizations employ passwords, physical authenticators, or biometrics to authenticate user identities or, in the case of multi-factor authentication, some combination thereof. Access to organizational systems is defined as either local access or network access. Local access is any access to organizational systems by users or processes acting on behalf of users, where access is obtained through direct connections without the use of networks. Network access is access to organizational systems by users (or processes acting on behalf of users) where access is obtained through network connections (i.e., nonlocal accesses). Remote access is a type of network access that involves communication through external networks. Internal networks include local area networks and wide area networks.

The use of encrypted virtual private networks for network connections between organization-controlled endpoints and non-organization-controlled endpoints may be treated as internal networks with respect to protecting the confidentiality and integrity of information traversing the network. Identification and authentication requirements for non-organizational users are described in [IA-8](#ia-8).

organizational users are uniquely identified and authenticated;

the unique identification of authenticated organizational users is associated with processes acting on behalf of those users.

Identification and authentication policy

procedures addressing user identification and authentication

system security plan, system design documentation

system configuration settings and associated documentation

system audit records

list of system accounts

other relevant documents or records

Organizational personnel with system operations responsibilities

organizational personnel with information security responsibilities

system/network administrators

organizational personnel with account management responsibilities

system developers

Organizational processes for uniquely identifying and authenticating users

mechanisms supporting and/or implementing identification and authentication capabilities

### IA-2 (1): Multi-factor Authentication to Privileged Accounts

Implement multi-factor authentication for access to privileged accounts.

Multi-factor authentication requires the use of two or more different factors to achieve authentication. The authentication factors are defined as follows: something you know (e.g., a personal identification number [PIN]), something you have (e.g., a physical authenticator such as a cryptographic private key), or something you are (e.g., a biometric). Multi-factor authentication solutions that feature physical authenticators include hardware authenticators that provide time-based or challenge-response outputs and smart cards such as the U.S. Government Personal Identity Verification (PIV) card or the Department of Defense (DoD) Common Access Card (CAC). In addition to authenticating users at the system level (i.e., at logon), organizations may employ authentication mechanisms at the application level, at their discretion, to provide increased security. Regardless of the type of access (i.e., local, network, remote), privileged accounts are authenticated using multi-factor options appropriate for the level of risk. Organizations can add additional security measures, such as additional or more rigorous authentication mechanisms, for specific types of access.

multi-factor authentication is implemented for access to privileged accounts.

Identification and authentication policy

procedures addressing user identification and authentication

system security plan

system design documentation

system configuration settings and associated documentation

system audit records

list of system accounts

other relevant documents or records

Organizational personnel with system operations responsibilities

organizational personnel with account management responsibilities

organizational personnel with information security responsibilities

system/network administrators

system developers

Mechanisms supporting and/or implementing a multi-factor authentication capability

### IA-2 (2): Multi-factor Authentication to Non-privileged Accounts

Implement multi-factor authentication for access to non-privileged accounts.

Multi-factor authentication requires the use of two or more different factors to achieve authentication. The authentication factors are defined as follows: something you know (e.g., a personal identification number [PIN]), something you have (e.g., a physical authenticator such as a cryptographic private key), or something you are (e.g., a biometric). Multi-factor authentication solutions that feature physical authenticators include hardware authenticators that provide time-based or challenge-response outputs and smart cards such as the U.S. Government Personal Identity Verification card or the DoD Common Access Card. In addition to authenticating users at the system level, organizations may also employ authentication mechanisms at the application level, at their discretion, to provide increased information security. Regardless of the type of access (i.e., local, network, remote), non-privileged accounts are authenticated using multi-factor options appropriate for the level of risk. Organizations can provide additional security measures, such as additional or more rigorous authentication mechanisms, for specific types of access.

multi-factor authentication for access to non-privileged accounts is implemented.

Identification and authentication policy

system security plan

procedures addressing user identification and authentication

system design documentation

system configuration settings and associated documentation

system audit records

list of system accounts

other relevant documents or records

Organizational personnel with system operations responsibilities

organizational personnel with account management responsibilities

organizational personnel with information security responsibilities

system/network administrators

system developers

Mechanisms supporting and/or implementing a multi-factor authentication capability

### IA-2 (8): Access to Accounts — Replay Resistant

Implement replay-resistant authentication mechanisms for access to privileged accountsand/ornon-privileged accounts.

Authentication processes resist replay attacks if it is impractical to achieve successful authentications by replaying previous authentication messages. Replay-resistant techniques include protocols that use nonces or challenges such as time synchronous or cryptographic authenticators.

replay-resistant authentication mechanisms for access to privileged accountsand/ornon-privileged accounts are implemented.

Identification and authentication policy

system security plan

procedures addressing user identification and authentication

system design documentation

system configuration settings and associated documentation

system audit records

list of privileged system accounts

other relevant documents or records

Organizational personnel with system operations responsibilities

organizational personnel with account management responsibilities

organizational personnel with information security responsibilities

system/network administrators

system developers

Mechanisms supporting and/or implementing identification and authentication capabilities

Mechanisms supporting and/or implementing replay-resistant authentication mechanisms

### IA-2 (12): Acceptance of PIV Credentials

Accept and electronically verify Personal Identity Verification-compliant credentials.

Acceptance of Personal Identity Verification (PIV)-compliant credentials applies to organizations implementing logical access control and physical access control systems. PIV-compliant credentials are those credentials issued by federal agencies that conform to FIPS Publication 201 and supporting guidance documents. The adequacy and reliability of PIV card issuers are authorized using [SP 800-79-2](#10963761-58fc-4b20-b3d6-b44a54daba03) . Acceptance of PIV-compliant credentials includes derived PIV credentials, the use of which is addressed in [SP 800-166](#e8552d48-cf41-40aa-8b06-f45f7fb4706c) . The DOD Common Access Card (CAC) is an example of a PIV credential.

Personal Identity Verification-compliant credentials are accepted and electronically verified.

Identification and authentication policy

system security plan

procedures addressing user identification and authentication

system design documentation

system configuration settings and associated documentation

system audit records

PIV verification records

evidence of PIV credentials

PIV credential authorizations

other relevant documents or records

Organizational personnel with system operations responsibilities

organizational personnel with account management responsibilities

organizational personnel with information security responsibilities

system/network administrators

system developers

Mechanisms supporting and/or implementing acceptance and verification of PIV credentials

### IA-3: Device Identification and Authentication

Uniquely identify and authenticate devices and/or types of devices to be uniquely identified and authenticated before establishing a connection are defined; before establishing a local, remote, and/or network connection.

Devices that require unique device-to-device identification and authentication are defined by type, device, or a combination of type and device. Organization-defined device types include devices that are not owned by the organization. Systems use shared known information (e.g., Media Access Control [MAC], Transmission Control Protocol/Internet Protocol [TCP/IP] addresses) for device identification or organizational authentication solutions (e.g., Institute of Electrical and Electronics Engineers (IEEE) 802.1x and Extensible Authentication Protocol [EAP], RADIUS server with EAP-Transport Layer Security [TLS] authentication, Kerberos) to identify and authenticate devices on local and wide area networks. Organizations determine the required strength of authentication mechanisms based on the security categories of systems and mission or business requirements. Because of the challenges of implementing device authentication on a large scale, organizations can restrict the application of the control to a limited number/type of devices based on mission or business needs.

 devices and/or types of devices to be uniquely identified and authenticated before establishing a connection are defined; are uniquely identified and authenticated before establishing a local, remote, and/or network connection.

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

### IA-4: Identifier Management

Manage system identifiers by:

Receiving authorization from personnel or roles from whom authorization must be received to assign an identifier are defined; to assign an individual, group, role, service, or device identifier;

Selecting an identifier that identifies an individual, group, role, service, or device;

Assigning the identifier to the intended individual, group, role, service, or device; and

Preventing reuse of identifiers for a time period for preventing reuse of identifiers is defined;.

Common device identifiers include Media Access Control (MAC) addresses, Internet Protocol (IP) addresses, or device-unique token identifiers. The management of individual identifiers is not applicable to shared system accounts. Typically, individual identifiers are the usernames of the system accounts assigned to those individuals. In such instances, the account management activities of [AC-2](#ac-2) use account names provided by [IA-4](#ia-4) . Identifier management also addresses individual identifiers not necessarily associated with system accounts. Preventing the reuse of identifiers implies preventing the assignment of previously used individual, group, role, service, or device identifiers to different individuals, groups, roles, services, or devices.

system identifiers are managed by receiving authorization from personnel or roles from whom authorization must be received to assign an identifier are defined; to assign to an individual, group, role, or device identifier;

system identifiers are managed by selecting an identifier that identifies an individual, group, role, service, or device;

system identifiers are managed by assigning the identifier to the intended individual, group, role, service, or device;

system identifiers are managed by preventing reuse of identifiers for a time period for preventing reuse of identifiers is defined;.

Identification and authentication policy

procedures addressing identifier management

procedures addressing account management

system security plan

system design documentation

system configuration settings and associated documentation

list of system accounts

list of identifiers generated from physical access control devices

other relevant documents or records

Organizational personnel with identifier management responsibilities

organizational personnel with information security responsibilities

system/network administrators

system developers

Mechanisms supporting and/or implementing identifier management

### IA-4 (4): Identify User Status

Manage individual identifiers by uniquely identifying each individual as characteristics used to identify individual status is defined;.

Characteristics that identify the status of individuals include contractors, foreign nationals, and non-organizational users. Identifying the status of individuals by these characteristics provides additional information about the people with whom organizational personnel are communicating. For example, it might be useful for a government employee to know that one of the individuals on an email message is a contractor.

individual identifiers are managed by uniquely identifying each individual as characteristics used to identify individual status is defined;.

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

### IA-5: Authenticator Management

Manage system authenticators by:

Verifying, as part of the initial authenticator distribution, the identity of the individual, group, role, service, or device receiving the authenticator;

Establishing initial authenticator content for any authenticators issued by the organization;

Ensuring that authenticators have sufficient strength of mechanism for their intended use;

Establishing and implementing administrative procedures for initial authenticator distribution, for lost or compromised or damaged authenticators, and for revoking authenticators;

Changing default authenticators prior to first use;

Changing or refreshing authenticators a time period for changing or refreshing authenticators by authenticator type is defined; or when events that trigger the change or refreshment of authenticators are defined; occur;

Protecting authenticator content from unauthorized disclosure and modification;

Requiring individuals to take, and having devices implement, specific controls to protect authenticators; and

Changing authenticators for group or role accounts when membership to those accounts changes.

Authenticators include passwords, cryptographic devices, biometrics, certificates, one-time password devices, and ID badges. Device authenticators include certificates and passwords. Initial authenticator content is the actual content of the authenticator (e.g., the initial password). In contrast, the requirements for authenticator content contain specific criteria or characteristics (e.g., minimum password length). Developers may deliver system components with factory default authentication credentials (i.e., passwords) to allow for initial installation and configuration. Default authentication credentials are often well known, easily discoverable, and present a significant risk. The requirement to protect individual authenticators may be implemented via control [PL-4](#pl-4) or [PS-6](#ps-6) for authenticators in the possession of individuals and by controls [AC-3](#ac-3), [AC-6](#ac-6) , and [SC-28](#sc-28) for authenticators stored in organizational systems, including passwords stored in hashed or encrypted formats or files containing encrypted or hashed passwords accessible with administrator privileges.

Systems support authenticator management by organization-defined settings and restrictions for various authenticator characteristics (e.g., minimum password length, validation time window for time synchronous one-time tokens, and number of allowed rejections during the verification stage of biometric authentication). Actions can be taken to safeguard individual authenticators, including maintaining possession of authenticators, not sharing authenticators with others, and immediately reporting lost, stolen, or compromised authenticators. Authenticator management includes issuing and revoking authenticators for temporary access when no longer needed.

system authenticators are managed through the verification of the identity of the individual, group, role, service, or device receiving the authenticator as part of the initial authenticator distribution;

system authenticators are managed through the establishment of initial authenticator content for any authenticators issued by the organization;

system authenticators are managed to ensure that authenticators have sufficient strength of mechanism for their intended use;

system authenticators are managed through the establishment and implementation of administrative procedures for initial authenticator distribution; lost, compromised, or damaged authenticators; and the revocation of authenticators;

system authenticators are managed through the change of default authenticators prior to first use;

system authenticators are managed through the change or refreshment of authenticators a time period for changing or refreshing authenticators by authenticator type is defined; or when events that trigger the change or refreshment of authenticators are defined; occur;

system authenticators are managed through the protection of authenticator content from unauthorized disclosure and modification;

system authenticators are managed through the requirement for individuals to take specific controls to protect authenticators;

system authenticators are managed through the requirement for devices to implement specific controls to protect authenticators;

system authenticators are managed through the change of authenticators for group or role accounts when membership to those accounts changes.

Identification and authentication policy

system security plan

addressing authenticator management

system design documentation

system configuration settings and associated documentation

list of system authenticator types

change control records associated with managing system authenticators

system audit records

other relevant documents or records

Organizational personnel with authenticator management responsibilities

organizational personnel with information security responsibilities

system/network administrators

Mechanisms supporting and/or implementing authenticator management capability

### IA-5 (1): Password-based Authentication

For password-based authentication:

Maintain a list of commonly-used, expected, or compromised passwords and update the list the frequency at which to update the list of commonly used, expected, or compromised passwords is defined; and when organizational passwords are suspected to have been compromised directly or indirectly;

Verify, when users create or update passwords, that the passwords are not found on the list of commonly-used, expected, or compromised passwords in IA-5(1)(a);

Transmit passwords only over cryptographically-protected channels;

Store passwords using an approved salted key derivation function, preferably using a keyed hash;

Require immediate selection of a new password upon account recovery;

Allow user selection of long passwords and passphrases, including spaces and all printable characters;

Employ automated tools to assist the user in selecting strong password authenticators; and

Enforce the following composition and complexity rules: authenticator composition and complexity rules are defined;.

Password-based authentication applies to passwords regardless of whether they are used in single-factor or multi-factor authentication. Long passwords or passphrases are preferable over shorter passwords. Enforced composition rules provide marginal security benefits while decreasing usability. However, organizations may choose to establish certain rules for password generation (e.g., minimum character length for long passwords) under certain circumstances and can enforce this requirement in IA-5(1)(h). Account recovery can occur, for example, in situations when a password is forgotten. Cryptographically protected passwords include salted one-way cryptographic hashes of passwords. The list of commonly used, compromised, or expected passwords includes passwords obtained from previous breach corpuses, dictionary words, and repetitive or sequential characters. The list includes context-specific words, such as the name of the service, username, and derivatives thereof.

for password-based authentication, a list of commonly used, expected, or compromised passwords is maintained and updated the frequency at which to update the list of commonly used, expected, or compromised passwords is defined; and when organizational passwords are suspected to have been compromised directly or indirectly;

for password-based authentication when passwords are created or updated by users, the passwords are verified not to be found on the list of commonly used, expected, or compromised passwords in IA-05(01)(a);

for password-based authentication, passwords are only transmitted over cryptographically protected channels;

for password-based authentication, passwords are stored using an approved salted key derivation function, preferably using a keyed hash;

for password-based authentication, immediate selection of a new password is required upon account recovery;

for password-based authentication, user selection of long passwords and passphrases is allowed, including spaces and all printable characters;

for password-based authentication, automated tools are employed to assist the user in selecting strong password authenticators;

for password-based authentication, authenticator composition and complexity rules are defined; are enforced.

Identification and authentication policy

password policy

procedures addressing authenticator management

system security plan

system design documentation

system configuration settings and associated documentation

password configurations and associated documentation

other relevant documents or records

Organizational personnel with authenticator management responsibilities

organizational personnel with information security responsibilities

system/network administrators

system developers

Mechanisms supporting and/or implementing password-based authenticator management capability

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

### IA-6: Authentication Feedback

Obscure feedback of authentication information during the authentication process to protect the information from possible exploitation and use by unauthorized individuals.

Authentication feedback from systems does not provide information that would allow unauthorized individuals to compromise authentication mechanisms. For some types of systems, such as desktops or notebooks with relatively large monitors, the threat (referred to as shoulder surfing) may be significant. For other types of systems, such as mobile devices with small displays, the threat may be less significant and is balanced against the increased likelihood of typographic input errors due to small keyboards. Thus, the means for obscuring authentication feedback is selected accordingly. Obscuring authentication feedback includes displaying asterisks when users type passwords into input devices or displaying feedback for a very limited time before obscuring it.

the feedback of authentication information is obscured during the authentication process to protect the information from possible exploitation and use by unauthorized individuals.

Identification and authentication policy

system security plan

procedures addressing authenticator feedback

system design documentation

system configuration settings and associated documentation

system audit records

other relevant documents or records

Organizational personnel with information security responsibilities

system/network administrators

system developers

Mechanisms supporting and/or implementing the obscuring of feedback of authentication information during authentication

### IA-7: Cryptographic Module Authentication

Implement mechanisms for authentication to a cryptographic module that meet the requirements of applicable laws, executive orders, directives, policies, regulations, standards, and guidelines for such authentication.

Authentication mechanisms may be required within a cryptographic module to authenticate an operator accessing the module and to verify that the operator is authorized to assume the requested role and perform services within that role.

mechanisms for authentication to a cryptographic module are implemented that meet the requirements of applicable laws, executive orders, directives, policies, regulations, standards, and guidelines for such authentication.

Identification and authentication policy

system security plan

procedures addressing cryptographic module authentication

system design documentation

system configuration settings and associated documentation

system audit records

other relevant documents or records

Organizational personnel with responsibility for cryptographic module authentication

organizational personnel with information security responsibilities

system/network administrators

system developers

Mechanisms supporting and/or implementing cryptographic module authentication

### IA-8: Identification and Authentication (Non-organizational Users)

Uniquely identify and authenticate non-organizational users or processes acting on behalf of non-organizational users.

Non-organizational users include system users other than organizational users explicitly covered by [IA-2](#ia-2) . Non-organizational users are uniquely identified and authenticated for accesses other than those explicitly identified and documented in [AC-14](#ac-14) . Identification and authentication of non-organizational users accessing federal systems may be required to protect federal, proprietary, or privacy-related information (with exceptions noted for national security systems). Organizations consider many factors—including security, privacy, scalability, and practicality—when balancing the need to ensure ease of use for access to federal information and systems with the need to protect and adequately mitigate risk.

non-organizational users or processes acting on behalf of non-organizational users are uniquely identified and authenticated.

Identification and authentication policy

system security plan

privacy plan

procedures addressing user identification and authentication

system design documentation

system configuration settings and associated documentation

system audit records

list of system accounts

other relevant documents or records

Organizational personnel with system operations responsibilities

organizational personnel with information security and privacy responsibilities

system/network administrators

organizational personnel with account management responsibilities

Mechanisms supporting and/or implementing identification and authentication capabilities

### IA-8 (1): Acceptance of PIV Credentials from Other Agencies

Accept and electronically verify Personal Identity Verification-compliant credentials from other federal agencies.

Acceptance of Personal Identity Verification (PIV) credentials from other federal agencies applies to both logical and physical access control systems. PIV credentials are those credentials issued by federal agencies that conform to FIPS Publication 201 and supporting guidelines. The adequacy and reliability of PIV card issuers are addressed and authorized using [SP 800-79-2](#10963761-58fc-4b20-b3d6-b44a54daba03).

Personal Identity Verification-compliant credentials from other federal agencies are accepted;

Personal Identity Verification-compliant credentials from other federal agencies are electronically verified.

Identification and authentication policy

system security plan

procedures addressing user identification and authentication

system design documentation

system configuration settings and associated documentation

system audit records

PIV verification records

evidence of PIV credentials

PIV credential authorizations

other relevant documents or records

Organizational personnel with system operations responsibilities

organizational personnel with information security responsibilities

system/network administrators

system developers

organizational personnel with account management responsibilities

Mechanisms supporting and/or implementing identification and authentication capabilities

mechanisms that accept and verify PIV credentials

### IA-8 (2): Acceptance of External Authenticators

Accept only external authenticators that are NIST-compliant; and

Document and maintain a list of accepted external authenticators.

Acceptance of only NIST-compliant external authenticators applies to organizational systems that are accessible to the public (e.g., public-facing websites). External authenticators are issued by nonfederal government entities and are compliant with [SP 800-63B](#e59c5a7c-8b1f-49ca-8de0-6ee0882180ce) . Approved external authenticators meet or exceed the minimum Federal Government-wide technical, security, privacy, and organizational maturity requirements. Meeting or exceeding Federal requirements allows Federal Government relying parties to trust external authenticators in connection with an authentication transaction at a specified authenticator assurance level.

only external authenticators that are NIST-compliant are accepted;

a list of accepted external authenticators is documented;

a list of accepted external authenticators is maintained.

Identification and authentication policy

system security plan

procedures addressing user identification and authentication

system design documentation

system configuration settings and associated documentation

system audit records

list of third-party credentialing products, components, or services procured and implemented by organization

third-party credential verification records

evidence of third-party credentials

third-party credential authorizations

other relevant documents or records

Organizational personnel with system operations responsibilities

organizational personnel with information security responsibilities

system/network administrators

system developers

organizational personnel with account management responsibilities

Mechanisms supporting and/or implementing identification and authentication capabilities

mechanisms that accept external credentials

### IA-8 (4): Use of Defined Profiles

Conform to the following profiles for identity management identity management profiles are defined;.

Organizations define profiles for identity management based on open identity management standards. To ensure that open identity management standards are viable, robust, reliable, sustainable, and interoperable as documented, the Federal Government assesses and scopes the standards and technology implementations against applicable laws, executive orders, directives, policies, regulations, standards, and guidelines.

there is conformance with identity management profiles are defined; for identity management.

Identification and authentication policy

system security plan

system design documentation

system configuration settings and associated documentation

system audit records

other relevant documents or records

Organizational personnel with system operations responsibilities

organizational personnel with information security responsibilities

system/network administrators

system developers

organizational personnel with account management responsibilities

Mechanisms supporting and/or implementing identification and authentication capabilities

mechanisms supporting and/or implementing conformance with profiles

### IA-11: Re-authentication

Require users to re-authenticate when circumstances or situations requiring re-authentication are defined;.

In addition to the re-authentication requirements associated with device locks, organizations may require re-authentication of individuals in certain situations, including when roles, authenticators or credentials change, when security categories of systems change, when the execution of privileged functions occurs, after a fixed time period, or periodically.

users are required to re-authenticate when circumstances or situations requiring re-authentication are defined;.

Identification and authentication policy

procedures addressing user and device re-authentication

system security plan

system design documentation

system configuration settings and associated documentation

list of circumstances or situations requiring re-authentication

system audit records

other relevant documents or records

Organizational personnel with system operations responsibilities

organizational personnel with information security responsibilities

system/network administrators

system developers

organizational personnel with identification and authentication responsibilities

Mechanisms supporting and/or implementing identification and authentication capabilities

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

Require that the presented identity evidence be validated and verified through methods of validation and verification of identity evidence are defined;.

Validation and verification of identity evidence increases the assurance that accounts and identifiers are being established for the correct user and authenticators are being bound to that user. Validation refers to the process of confirming that the evidence is genuine and authentic, and the data contained in the evidence is correct, current, and related to an individual. Verification confirms and establishes a linkage between the claimed identity and the actual existence of the user presenting the evidence. Acceptable methods for validating and verifying identity evidence are consistent with the risks to the systems, roles, and privileges associated with the users account.

the presented identity evidence is validated and verified through methods of validation and verification of identity evidence are defined;.

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

Require that a registration codeornotice of proofing be delivered through an out-of-band channel to verify the users address (physical or digital) of record.

To make it more difficult for adversaries to pose as legitimate users during the identity proofing process, organizations can use out-of-band methods to ensure that the individual associated with an address of record is the same individual that participated in the registration. Confirmation can take the form of a temporary enrollment code or a notice of proofing. The delivery address for these artifacts is obtained from records and not self-asserted by the user. The address can include a physical or digital address. A home address is an example of a physical address. Email addresses and telephone numbers are examples of digital addresses.

a registration codeornotice of proofing is delivered through an out-of-band channel to verify the user’s address (physical or digital) of record.

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

