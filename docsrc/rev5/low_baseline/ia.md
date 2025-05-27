# ia - Identification and Authentication

## Base Controls

### ia-1: Policy and Procedures

Develop, document, and disseminate to {{ insert: param, ia-1_prm_1 }}:

 {{ insert: param, ia-01_odp.03 }} identification and authentication policy that:

Addresses purpose, scope, roles, responsibilities, management commitment, coordination among organizational entities, and compliance; and

Is consistent with applicable laws, executive orders, directives, regulations, policies, standards, and guidelines; and

Procedures to facilitate the implementation of the identification and authentication policy and the associated identification and authentication controls;

Designate an {{ insert: param, ia-01_odp.04 }} to manage the development, documentation, and dissemination of the identification and authentication policy and procedures; and

Review and update the current identification and authentication:

Policy {{ insert: param, ia-01_odp.05 }} and following {{ insert: param, ia-01_odp.06 }} ; and

Procedures {{ insert: param, ia-01_odp.07 }} and following {{ insert: param, ia-01_odp.08 }}.

Identification and authentication policy and procedures address the controls in the IA family that are implemented within systems and organizations. The risk management strategy is an important factor in establishing such policies and procedures. Policies and procedures contribute to security and privacy assurance. Therefore, it is important that security and privacy programs collaborate on the development of identification and authentication policy and procedures. Security and privacy program policies and procedures at the organization level are preferable, in general, and may obviate the need for mission- or system-specific policies and procedures. The policy can be included as part of the general security and privacy policy or be represented by multiple policies that reflect the complex nature of organizations. Procedures can be established for security and privacy programs, for mission or business processes, and for systems, if needed. Procedures describe how the policies or controls are implemented and can be directed at the individual or role that is the object of the procedure. Procedures can be documented in system security and privacy plans or in one or more separate documents. Events that may precipitate an update to identification and authentication policy and procedures include assessment or audit findings, security incidents or breaches, or changes in applicable laws, executive orders, directives, regulations, policies, standards, and guidelines. Simply restating controls does not constitute an organizational policy or procedure.

an identification and authentication policy is developed and documented;

the identification and authentication policy is disseminated to {{ insert: param, ia-01_odp.01 }};

identification and authentication procedures to facilitate the implementation of the identification and authentication policy and associated identification and authentication controls are developed and documented;

the identification and authentication procedures are disseminated to {{ insert: param, ia-01_odp.02 }};

the {{ insert: param, ia-01_odp.03 }} identification and authentication policy addresses purpose;

the {{ insert: param, ia-01_odp.03 }} identification and authentication policy addresses scope;

the {{ insert: param, ia-01_odp.03 }} identification and authentication policy addresses roles;

the {{ insert: param, ia-01_odp.03 }} identification and authentication policy addresses responsibilities;

the {{ insert: param, ia-01_odp.03 }} identification and authentication policy addresses management commitment;

the {{ insert: param, ia-01_odp.03 }} identification and authentication policy addresses coordination among organizational entities;

the {{ insert: param, ia-01_odp.03 }} identification and authentication policy addresses compliance;

the {{ insert: param, ia-01_odp.03 }} identification and authentication policy is consistent with applicable laws, executive orders, directives, regulations, policies, standards, and guidelines;

the {{ insert: param, ia-01_odp.04 }} is designated to manage the development, documentation, and dissemination of the identification and authentication policy and procedures;

the current identification and authentication policy is reviewed and updated {{ insert: param, ia-01_odp.05 }};

the current identification and authentication policy is reviewed and updated following {{ insert: param, ia-01_odp.06 }};

the current identification and authentication procedures are reviewed and updated {{ insert: param, ia-01_odp.07 }};

the current identification and authentication procedures are reviewed and updated following {{ insert: param, ia-01_odp.08 }}.

Identification and authentication policy and procedures

system security plan

privacy plan

risk management strategy documentation

list of events requiring identification and authentication procedures to be reviewed and updated (e.g., audit findings)

other relevant documents or records

Organizational personnel with identification and authentication responsibilities

organizational personnel with information security and privacy responsibilities

### ia-2: Identification and Authentication (Organizational Users)

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

### ia-4: Identifier Management

Manage system identifiers by:

Receiving authorization from {{ insert: param, ia-04_odp.01 }} to assign an individual, group, role, service, or device identifier;

Selecting an identifier that identifies an individual, group, role, service, or device;

Assigning the identifier to the intended individual, group, role, service, or device; and

Preventing reuse of identifiers for {{ insert: param, ia-04_odp.02 }}.

Common device identifiers include Media Access Control (MAC) addresses, Internet Protocol (IP) addresses, or device-unique token identifiers. The management of individual identifiers is not applicable to shared system accounts. Typically, individual identifiers are the usernames of the system accounts assigned to those individuals. In such instances, the account management activities of [AC-2](#ac-2) use account names provided by [IA-4](#ia-4) . Identifier management also addresses individual identifiers not necessarily associated with system accounts. Preventing the reuse of identifiers implies preventing the assignment of previously used individual, group, role, service, or device identifiers to different individuals, groups, roles, services, or devices.

system identifiers are managed by receiving authorization from {{ insert: param, ia-04_odp.01 }} to assign to an individual, group, role, or device identifier;

system identifiers are managed by selecting an identifier that identifies an individual, group, role, service, or device;

system identifiers are managed by assigning the identifier to the intended individual, group, role, service, or device;

system identifiers are managed by preventing reuse of identifiers for {{ insert: param, ia-04_odp.02 }}.

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

### ia-5: Authenticator Management

Manage system authenticators by:

Verifying, as part of the initial authenticator distribution, the identity of the individual, group, role, service, or device receiving the authenticator;

Establishing initial authenticator content for any authenticators issued by the organization;

Ensuring that authenticators have sufficient strength of mechanism for their intended use;

Establishing and implementing administrative procedures for initial authenticator distribution, for lost or compromised or damaged authenticators, and for revoking authenticators;

Changing default authenticators prior to first use;

Changing or refreshing authenticators {{ insert: param, ia-05_odp.01 }} or when {{ insert: param, ia-05_odp.02 }} occur;

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

system authenticators are managed through the change or refreshment of authenticators {{ insert: param, ia-05_odp.01 }} or when {{ insert: param, ia-05_odp.02 }} occur;

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

### ia-6: Authentication Feedback

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

### ia-7: Cryptographic Module Authentication

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

### ia-8: Identification and Authentication (Non-organizational Users)

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

### ia-11: Re-authentication

Require users to re-authenticate when {{ insert: param, ia-11_odp }}.

In addition to the re-authentication requirements associated with device locks, organizations may require re-authentication of individuals in certain situations, including when roles, authenticators or credentials change, when security categories of systems change, when the execution of privileged functions occurs, after a fixed time period, or periodically.

users are required to re-authenticate when {{ insert: param, ia-11_odp }}.

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

## Control Enhancements

