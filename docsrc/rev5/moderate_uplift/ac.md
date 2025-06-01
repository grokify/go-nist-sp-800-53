# AC - Access Control

* Controls: 28

## Controls

### AC-2 (1): Automated System Account Management

Support the management of system accounts using automated mechanisms used to support the management of system accounts are defined;.

Automated system account management includes using automated mechanisms to create, enable, modify, disable, and remove accounts; notify account managers when an account is created, enabled, modified, disabled, or removed, or when users are terminated or transferred; monitor system account usage; and report atypical system account usage. Automated mechanisms can include internal system functions and email, telephonic, and text messaging notifications.

the management of system accounts is supported using automated mechanisms used to support the management of system accounts are defined;.

Access control policy

procedures for addressing account management

system design documentation

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

Organizational personnel with account management responsibilities

system/network administrators

organizational personnel with information security with information security responsibilities

system developers

Automated mechanisms for implementing account management functions

### AC-2 (2): Automated Temporary and Emergency Account Management

Automatically removeordisable temporary and emergency accounts after the time period after which to automatically remove or disable temporary or emergency accounts is defined;.

Management of temporary and emergency accounts includes the removal or disabling of such accounts automatically after a predefined time period rather than at the convenience of the system administrator. Automatic removal or disabling of accounts provides a more consistent implementation.

temporary and emergency accounts are automatically removeordisable after the time period after which to automatically remove or disable temporary or emergency accounts is defined;.

Access control policy

procedures for addressing account management

system design documentation

system configuration settings and associated documentation

system-generated list of temporary accounts removed and/or disabled

system-generated list of emergency accounts removed and/or disabled

system audit records

system security plan

other relevant documents or records

Organizational personnel with account management responsibilities

system/network administrators

organizational personnel with information security with information security responsibilities

system developers

Automated mechanisms for implementing account management functions

### AC-2 (3): Disable Accounts

Disable accounts within time period within which to disable accounts is defined; when the accounts:

Have expired;

Are no longer associated with a user or individual;

Are in violation of organizational policy; or

Have been inactive for time period for account inactivity before disabling is defined;.

Disabling expired, inactive, or otherwise anomalous accounts supports the concepts of least privilege and least functionality which reduce the attack surface of the system.

accounts are disabled within time period within which to disable accounts is defined; when the accounts have expired;

accounts are disabled within time period within which to disable accounts is defined; when the accounts are no longer associated with a user or individual;

accounts are disabled within time period within which to disable accounts is defined; when the accounts are in violation of organizational policy;

accounts are disabled within time period within which to disable accounts is defined; when the accounts have been inactive for time period for account inactivity before disabling is defined;.

Access control policy

procedures for addressing account management

system security plan

system design documentation

system configuration settings and associated documentation

system-generated list of accounts removed

system-generated list of emergency accounts disabled

system audit records

system security plan

other relevant documents or records

Organizational personnel with account management responsibilities

system/network administrators

organizational personnel with information security responsibilities

system developers

Mechanisms for implementing account management functions

### AC-2 (4): Automated Audit Actions

Automatically audit account creation, modification, enabling, disabling, and removal actions.

Account management audit records are defined in accordance with [AU-02](#au-2) and reviewed, analyzed, and reported in accordance with [AU-06](#au-6).

account creation is automatically audited;

account modification is automatically audited;

account enabling is automatically audited;

account disabling is automatically audited;

account removal actions are automatically audited.

Access control policy

procedures addressing account management

system design documentation

system configuration settings and associated documentation

notifications/alerts of account creation, modification, enabling, disabling, and removal actions

system audit records

system security plan

other relevant documents or records

Organizational personnel with account management responsibilities

system/network administrators

organizational personnel with information security responsibilities

Automated mechanisms implementing account management functions

### AC-2 (5): Inactivity Logout

Require that users log out when the time period of expected inactivity or description of when to log out is defined;.

Inactivity logout is behavior- or policy-based and requires users to take physical action to log out when they are expecting inactivity longer than the defined period. Automatic enforcement of inactivity logout is addressed by [AC-11](#ac-11).

users are required to log out when the time period of expected inactivity or description of when to log out is defined;.

Access control policy

procedures addressing account management

system design documentation

system configuration settings and associated documentation

security violation reports

system audit records

system security plan

other relevant documents or records

Organizational personnel with account management responsibilities

system/network administrators

organizational personnel with information security responsibilities

users that must comply with inactivity logout policy

### AC-2 (13): Disable Accounts for High-risk Individuals

Disable accounts of individuals within time period within which to disable accounts of individuals who are discovered to pose significant risk is defined; of discovery of significant risks leading to disabling accounts are defined;.

Users who pose a significant security and/or privacy risk include individuals for whom reliable evidence indicates either the intention to use authorized access to systems to cause harm or through whom adversaries will cause harm. Such harm includes adverse impacts to organizational operations, organizational assets, individuals, other organizations, or the Nation. Close coordination among system administrators, legal staff, human resource managers, and authorizing officials is essential when disabling system accounts for high-risk individuals.

accounts of individuals are disabled within time period within which to disable accounts of individuals who are discovered to pose significant risk is defined; of discovery of significant risks leading to disabling accounts are defined;.

Access control policy

procedures addressing account management

system design documentation

system configuration settings and associated documentation

system-generated list of disabled accounts

list of user activities posing significant organizational risk

system audit records

system security plan

other relevant documents or records

Organizational personnel with account management responsibilities

system/network administrators

organizational personnel with information security responsibilities

Mechanisms implementing account management functions

### AC-4: Information Flow Enforcement

Enforce approved authorizations for controlling the flow of information within the system and between connected systems based on information flow control policies within the system and between connected systems are defined;.

Information flow control regulates where information can travel within a system and between systems (in contrast to who is allowed to access the information) and without regard to subsequent accesses to that information. Flow control restrictions include blocking external traffic that claims to be from within the organization, keeping export-controlled information from being transmitted in the clear to the Internet, restricting web requests that are not from the internal web proxy server, and limiting information transfers between organizations based on data structures and content. Transferring information between organizations may require an agreement specifying how the information flow is enforced (see [CA-3](#ca-3) ). Transferring information between systems in different security or privacy domains with different security or privacy policies introduces the risk that such transfers violate one or more domain security or privacy policies. In such situations, information owners/stewards provide guidance at designated policy enforcement points between connected systems. Organizations consider mandating specific architectural solutions to enforce specific security and privacy policies. Enforcement includes prohibiting information transfers between connected systems (i.e., allowing access only), verifying write permissions before accepting information from another security or privacy domain or connected system, employing hardware mechanisms to enforce one-way information flows, and implementing trustworthy regrading mechanisms to reassign security or privacy attributes and labels.

Organizations commonly employ information flow control policies and enforcement mechanisms to control the flow of information between designated sources and destinations within systems and between connected systems. Flow control is based on the characteristics of the information and/or the information path. Enforcement occurs, for example, in boundary protection devices that employ rule sets or establish configuration settings that restrict system services, provide a packet-filtering capability based on header information, or provide a message-filtering capability based on message content. Organizations also consider the trustworthiness of filtering and/or inspection mechanisms (i.e., hardware, firmware, and software components) that are critical to information flow enforcement. Control enhancements 3 through 32 primarily address cross-domain solution needs that focus on more advanced filtering techniques, in-depth analysis, and stronger flow enforcement mechanisms implemented in cross-domain products, such as high-assurance guards. Such capabilities are generally not available in commercial off-the-shelf products. Information flow enforcement also applies to control plane traffic (e.g., routing and DNS).

approved authorizations are enforced for controlling the flow of information within the system and between connected systems based on information flow control policies within the system and between connected systems are defined;.

Access control policy

information flow control policies

procedures addressing information flow enforcement

security architecture documentation

privacy architecture documentation

system design documentation

system configuration settings and associated documentation

system baseline configuration

list of information flow authorizations

system audit records

system security plan

privacy plan

other relevant documents or records

System/network administrators

organizational personnel with information security and privacy architecture development responsibilities

organizational personnel with information security and privacy responsibilities

system developers

Mechanisms implementing information flow enforcement policy

### AC-5: Separation of Duties

Identify and document duties of individuals requiring separation are defined; ; and

Define system access authorizations to support separation of duties.

Separation of duties addresses the potential for abuse of authorized privileges and helps to reduce the risk of malevolent activity without collusion. Separation of duties includes dividing mission or business functions and support functions among different individuals or roles, conducting system support functions with different individuals, and ensuring that security personnel who administer access control functions do not also administer audit functions. Because separation of duty violations can span systems and application domains, organizations consider the entirety of systems and system components when developing policy on separation of duties. Separation of duties is enforced through the account management activities in [AC-2](#ac-2) , access control mechanisms in [AC-3](#ac-3) , and identity management activities in [IA-2](#ia-2), [IA-4](#ia-4) , and [IA-12](#ia-12).

duties of individuals requiring separation are defined; are identified and documented;

system access authorizations to support separation of duties are defined.

Access control policy

procedures addressing divisions of responsibility and separation of duties

system configuration settings and associated documentation

list of divisions of responsibility and separation of duties

system access authorizations

system audit records

system security plan

other relevant documents or records

Organizational personnel with responsibilities for defining appropriate divisions of responsibility and separation of duties

organizational personnel with information security responsibilities

system/network administrators

Mechanisms implementing separation of duties policy

### AC-6: Least Privilege

Employ the principle of least privilege, allowing only authorized accesses for users (or processes acting on behalf of users) that are necessary to accomplish assigned organizational tasks.

Organizations employ least privilege for specific duties and systems. The principle of least privilege is also applied to system processes, ensuring that the processes have access to systems and operate at privilege levels no higher than necessary to accomplish organizational missions or business functions. Organizations consider the creation of additional processes, roles, and accounts as necessary to achieve least privilege. Organizations apply least privilege to the development, implementation, and operation of organizational systems.

the principle of least privilege is employed, allowing only authorized accesses for users (or processes acting on behalf of users) that are necessary to accomplish assigned organizational tasks.

Access control policy

procedures addressing least privilege

list of assigned access authorizations (user privileges)

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

Organizational personnel with responsibilities for defining least privileges necessary to accomplish specified tasks

organizational personnel with information security responsibilities

system/network administrators

Mechanisms implementing least privilege functions

### AC-6 (1): Authorize Access to Security Functions

Authorize access for individuals and roles with authorized access to security functions and security-relevant information are defined; to:

organization-defined security functions (deployed in hardware, software, and firmware) ; and

 security-relevant information for authorized access is defined;.

Security functions include establishing system accounts, configuring access authorizations (i.e., permissions, privileges), configuring settings for events to be audited, and establishing intrusion detection parameters. Security-relevant information includes filtering rules for routers or firewalls, configuration parameters for security services, cryptographic key management information, and access control lists. Authorized personnel include security administrators, system administrators, system security officers, system programmers, and other privileged users.

access is authorized for individuals and roles with authorized access to security functions and security-relevant information are defined; to security functions (deployed in hardware) for authorized access are defined;;

access is authorized for individuals and roles with authorized access to security functions and security-relevant information are defined; to security functions (deployed in software) for authorized access are defined;;

access is authorized for individuals and roles with authorized access to security functions and security-relevant information are defined; to security functions (deployed in firmware) for authorized access are defined;;

access is authorized for individuals and roles with authorized access to security functions and security-relevant information are defined; to security-relevant information for authorized access is defined;.

Access control policy

procedures addressing least privilege

list of security functions (deployed in hardware, software, and firmware) and security-relevant information for which access must be explicitly authorized

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

Organizational personnel with responsibilities for defining least privileges necessary to accomplish specified tasks

organizational personnel with information security responsibilities

system/network administrators

Mechanisms implementing least privilege functions

### AC-6 (2): Non-privileged Access for Nonsecurity Functions

Require that users of system accounts (or roles) with access to security functions or security-relevant information, the access to which requires users to use non-privileged accounts to access non-security functions, are defined; use non-privileged accounts or roles, when accessing nonsecurity functions.

Requiring the use of non-privileged accounts when accessing nonsecurity functions limits exposure when operating from within privileged accounts or roles. The inclusion of roles addresses situations where organizations implement access control policies, such as role-based access control, and where a change of role provides the same degree of assurance in the change of access authorizations for the user and the processes acting on behalf of the user as would be provided by a change between a privileged and non-privileged account.

users of system accounts (or roles) with access to security functions or security-relevant information, the access to which requires users to use non-privileged accounts to access non-security functions, are defined; are required to use non-privileged accounts or roles when accessing non-security functions.

Access control policy

procedures addressing least privilege

list of system-generated security functions or security-relevant information assigned to system accounts or roles

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

Organizational personnel with responsibilities for defining least privileges necessary to accomplish specified tasks

organizational personnel with information security responsibilities

system/network administrators

Mechanisms implementing least privilege functions

### AC-6 (5): Privileged Accounts

Restrict privileged accounts on the system to personnel or roles to which privileged accounts on the system are to be restricted is/are defined;.

Privileged accounts, including super user accounts, are typically described as system administrator for various types of commercial off-the-shelf operating systems. Restricting privileged accounts to specific personnel or roles prevents day-to-day users from accessing privileged information or privileged functions. Organizations may differentiate in the application of restricting privileged accounts between allowed privileges for local accounts and for domain accounts provided that they retain the ability to control system configurations for key parameters and as otherwise necessary to sufficiently mitigate risk.

privileged accounts on the system are restricted to personnel or roles to which privileged accounts on the system are to be restricted is/are defined;.

Access control policy

procedures addressing least privilege

list of system-generated privileged accounts

list of system administration personnel

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

Organizational personnel with responsibilities for defining least privileges necessary to accomplish specified tasks

organizational personnel with information security responsibilities

system/network administrators

Mechanisms implementing least privilege functions

### AC-6 (7): Review of User Privileges

Review the frequency at which to review the privileges assigned to roles or classes of users is defined; the privileges assigned to roles or classes of users to which privileges are assigned are defined; to validate the need for such privileges; and

Reassign or remove privileges, if necessary, to correctly reflect organizational mission and business needs.

The need for certain assigned user privileges may change over time to reflect changes in organizational mission and business functions, environments of operation, technologies, or threats. A periodic review of assigned user privileges is necessary to determine if the rationale for assigning such privileges remains valid. If the need cannot be revalidated, organizations take appropriate corrective actions.

privileges assigned to roles or classes of users to which privileges are assigned are defined; are reviewed the frequency at which to review the privileges assigned to roles or classes of users is defined; to validate the need for such privileges;

privileges are reassigned or removed, if necessary, to correctly reflect organizational mission and business needs.

Access control policy

procedures addressing least privilege

list of system-generated roles or classes of users and assigned privileges

system design documentation

system configuration settings and associated documentation

validation reviews of privileges assigned to roles or classes or users

records of privilege removals or reassignments for roles or classes of users

system audit records

system security plan

other relevant documents or records

Organizational personnel with responsibilities for reviewing least privileges necessary to accomplish specified tasks

organizational personnel with information security responsibilities

system/network administrators

Mechanisms implementing review of user privileges

### AC-6 (9): Log Use of Privileged Functions

Log the execution of privileged functions.

The misuse of privileged functions, either intentionally or unintentionally by authorized users or by unauthorized external entities that have compromised system accounts, is a serious and ongoing concern and can have significant adverse impacts on organizations. Logging and analyzing the use of privileged functions is one way to detect such misuse and, in doing so, help mitigate the risk from insider threats and the advanced persistent threat.

the execution of privileged functions is logged.

Access control policy

procedures addressing least privilege

system design documentation

system configuration settings and associated documentation

list of privileged functions to be audited

list of audited events

system audit records

system security plan

other relevant documents or records

Organizational personnel with responsibilities for reviewing least privileges necessary to accomplish specified tasks

organizational personnel with information security responsibilities

system/network administrators

system developers

Mechanisms auditing the execution of least privilege functions

### AC-6 (10): Prohibit Non-privileged Users from Executing Privileged Functions

Prevent non-privileged users from executing privileged functions.

Privileged functions include disabling, circumventing, or altering implemented security or privacy controls, establishing system accounts, performing system integrity checks, and administering cryptographic key management activities. Non-privileged users are individuals who do not possess appropriate authorizations. Privileged functions that require protection from non-privileged users include circumventing intrusion detection and prevention mechanisms or malicious code protection mechanisms. Preventing non-privileged users from executing privileged functions is enforced by [AC-3](#ac-3).

non-privileged users are prevented from executing privileged functions.

Access control policy

procedures addressing least privilege

system design documentation

system configuration settings and associated documentation

list of privileged functions and associated user account assignments

system audit records

system security plan

other relevant documents or records

Organizational personnel with responsibilities for defining least privileges necessary to accomplish specified tasks

organizational personnel with information security responsibilities

system developers

Mechanisms implementing least privilege functions for non-privileged users

### AC-11: Device Lock

Prevent further access to the system by initiating a device lock after time period of inactivity after which a device lock is initiated is defined (if selected); of inactivityand/orrequiring the user to initiate a device lock before leaving the system unattended ; and

Retain the device lock until the user reestablishes access using established identification and authentication procedures.

Device locks are temporary actions taken to prevent logical access to organizational systems when users stop work and move away from the immediate vicinity of those systems but do not want to log out because of the temporary nature of their absences. Device locks can be implemented at the operating system level or at the application level. A proximity lock may be used to initiate the device lock (e.g., via a Bluetooth-enabled device or dongle). User-initiated device locking is behavior or policy-based and, as such, requires users to take physical action to initiate the device lock. Device locks are not an acceptable substitute for logging out of systems, such as when organizations require users to log out at the end of workdays.

further access to the system is prevented by initiating a device lock after time period of inactivity after which a device lock is initiated is defined (if selected); of inactivityand/orrequiring the user to initiate a device lock before leaving the system unattended;

device lock is retained until the user re-establishes access using established identification and authentication procedures.

Access control policy

procedures addressing session lock

procedures addressing identification and authentication

system design documentation

system configuration settings and associated documentation

security plan

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

system developers

Mechanisms implementing access control policy for session lock

### AC-11 (1): Pattern-hiding Displays

Conceal, via the device lock, information previously visible on the display with a publicly viewable image.

The pattern-hiding display can include static or dynamic images, such as patterns used with screen savers, photographic images, solid colors, clock, battery life indicator, or a blank screen with the caveat that controlled unclassified information is not displayed.

information previously visible on the display is concealed, via device lock, with a publicly viewable image.

Access control policy

procedures addressing session lock

display screen with session lock activated

system design documentation

system configuration settings and associated documentation

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

system developers

System session lock mechanisms

### AC-12: Session Termination

Automatically terminate a user session after conditions or trigger events requiring session disconnect are defined;.

Session termination addresses the termination of user-initiated logical sessions (in contrast to [SC-10](#sc-10) , which addresses the termination of network connections associated with communications sessions (i.e., network disconnect)). A logical session (for local, network, and remote access) is initiated whenever a user (or process acting on behalf of a user) accesses an organizational system. Such user sessions can be terminated without terminating network sessions. Session termination ends all processes associated with a user’s logical session except for those processes that are specifically created by the user (i.e., session owner) to continue after the session is terminated. Conditions or trigger events that require automatic termination of the session include organization-defined periods of user inactivity, targeted responses to certain types of incidents, or time-of-day restrictions on system use.

a user session is automatically terminated after conditions or trigger events requiring session disconnect are defined;.

Access control policy

procedures addressing session termination

system design documentation

system configuration settings and associated documentation

list of conditions or trigger events requiring session disconnect

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

system developers

Automated mechanisms implementing user session termination

### AC-17 (1): Monitoring and Control

Employ automated mechanisms to monitor and control remote access methods.

Monitoring and control of remote access methods allows organizations to detect attacks and help ensure compliance with remote access policies by auditing the connection activities of remote users on a variety of system components, including servers, notebook computers, workstations, smart phones, and tablets. Audit logging for remote access is enforced by [AU-2](#au-2) . Audit events are defined in [AU-2a](#au-2_smt.a).

automated mechanisms are employed to monitor remote access methods;

automated mechanisms are employed to control remote access methods.

Access control policy

procedures addressing remote access to the system

system design documentation

system configuration settings and associated documentation

system audit records

system monitoring records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

system developers

Automated mechanisms monitoring and controlling remote access methods

### AC-17 (2): Protection of Confidentiality and Integrity Using Encryption

Implement cryptographic mechanisms to protect the confidentiality and integrity of remote access sessions.

Virtual private networks can be used to protect the confidentiality and integrity of remote access sessions. Transport Layer Security (TLS) is an example of a cryptographic protocol that provides end-to-end communications security over networks and is used for Internet communications and online transactions.

cryptographic mechanisms are implemented to protect the confidentiality and integrity of remote access sessions.

Access control policy

procedures addressing remote access to the system

system design documentation

system configuration settings and associated documentation

cryptographic mechanisms and associated configuration documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

system developers

Cryptographic mechanisms protecting confidentiality and integrity of remote access sessions

### AC-17 (3): Managed Access Control Points

Route remote accesses through authorized and managed network access control points.

Organizations consider the Trusted Internet Connections (TIC) initiative [DHS TIC](#4f42ee6e-86cc-403b-a51f-76c2b4f81b54) requirements for external network connections since limiting the number of access control points for remote access reduces attack surfaces.

remote accesses are routed through authorized and managed network access control points.

Access control policy

procedures addressing remote access to the system

system design documentation

list of all managed network access control points

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

Mechanisms routing all remote accesses through managed network access control points

### AC-17 (4): Privileged Commands and Access

Authorize the execution of privileged commands and access to security-relevant information via remote access only in a format that provides assessable evidence and for the following needs: organization-defined needs ; and

Document the rationale for remote access in the security plan for the system.

Remote access to systems represents a significant potential vulnerability that can be exploited by adversaries. As such, restricting the execution of privileged commands and access to security-relevant information via remote access reduces the exposure of the organization and the susceptibility to threats by adversaries to the remote access capability.

the execution of privileged commands via remote access is authorized only in a format that provides assessable evidence;

access to security-relevant information via remote access is authorized only in a format that provides assessable evidence;

the execution of privileged commands via remote access is authorized only for the following needs: needs requiring execution of privileged commands via remote access are defined;;

access to security-relevant information via remote access is authorized only for the following needs: needs requiring access to security-relevant information via remote access are defined;;

the rationale for remote access is documented in the security plan for the system.

Access control policy

procedures addressing remote access to the system

system configuration settings and associated documentation

security plan

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

Mechanisms implementing remote access management

### AC-18 (1): Authentication and Encryption

Protect wireless access to the system using authentication of usersand/ordevices and encryption.

Wireless networking capabilities represent a significant potential vulnerability that can be exploited by adversaries. To protect systems with wireless access points, strong authentication of users and devices along with strong encryption can reduce susceptibility to threats by adversaries involving wireless technologies.

wireless access to the system is protected using authentication of usersand/ordevices;

wireless access to the system is protected using encryption.

Access control policy

procedures addressing wireless implementation and usage (including restrictions)

system design documentation

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

system developers

Mechanisms implementing wireless access protections to the system

### AC-18 (3): Disable Wireless Networking

Disable, when not intended for use, wireless networking capabilities embedded within system components prior to issuance and deployment.

Wireless networking capabilities that are embedded within system components represent a significant potential vulnerability that can be exploited by adversaries. Disabling wireless capabilities when not needed for essential organizational missions or functions can reduce susceptibility to threats by adversaries involving wireless technologies.

when not intended for use, wireless networking capabilities embedded within system components are disabled prior to issuance and deployment.

Access control policy

procedures addressing wireless implementation and usage (including restrictions)

system design documentation

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

Mechanisms managing the disabling of wireless networking capabilities internally embedded within system components

### AC-19 (5): Full Device or Container-based Encryption

Employ full-device encryptionorcontainer-based encryption to protect the confidentiality and integrity of information on mobile devices on which to employ encryption are defined;.

Container-based encryption provides a more fine-grained approach to data and information encryption on mobile devices, including encrypting selected data structures such as files, records, or fields.

 full-device encryptionorcontainer-based encryption is employed to protect the confidentiality and integrity of information on mobile devices on which to employ encryption are defined;.

Access control policy

procedures addressing access control for mobile devices

system design documentation

system configuration settings and associated documentation

encryption mechanisms and associated configuration documentation

system audit records

system security plan

other relevant documents or records

Organizational personnel with access control responsibilities for mobile devices

system/network administrators

organizational personnel with information security responsibilities

Encryption mechanisms protecting confidentiality and integrity of information on mobile devices

### AC-20 (1): Limits on Authorized Use

Permit authorized individuals to use an external system to access the system or to process, store, or transmit organization-controlled information only after:

Verification of the implementation of controls on the external system as specified in the organization’s security and privacy policies and security and privacy plans; or

Retention of approved system connection or processing agreements with the organizational entity hosting the external system.

Limiting authorized use recognizes circumstances where individuals using external systems may need to access organizational systems. Organizations need assurance that the external systems contain the necessary controls so as not to compromise, damage, or otherwise harm organizational systems. Verification that the required controls have been implemented can be achieved by external, independent assessments, attestations, or other means, depending on the confidence level required by organizations.

authorized individuals are permitted to use an external system to access the system or to process, store, or transmit organization-controlled information only after verification of the implementation of controls on the external system as specified in the organization’s security and privacy policies and security and privacy plans (if applicable);

authorized individuals are permitted to use an external system to access the system or to process, store, or transmit organization-controlled information only after retention of approved system connection or processing agreements with the organizational entity hosting the external system (if applicable).

Access control policy

procedures addressing the use of external systems

system connection or processing agreements

account management documents

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

Mechanisms implementing limits on use of external systems

### AC-20 (2): Portable Storage Devices — Restricted Use

Restrict the use of organization-controlled portable storage devices by authorized individuals on external systems using restrictions on the use of organization-controlled portable storage devices by authorized individuals on external systems are defined;.

Limits on the use of organization-controlled portable storage devices in external systems include restrictions on how the devices may be used and under what conditions the devices may be used.

the use of organization-controlled portable storage devices by authorized individuals is restricted on external systems using restrictions on the use of organization-controlled portable storage devices by authorized individuals on external systems are defined;.

Access control policy

procedures addressing the use of external systems

system configuration settings and associated documentation

system connection or processing agreements

account management documents

system security plan

other relevant documents or records

Organizational personnel with responsibilities for restricting or prohibiting the use of organization-controlled storage devices on external systems

system/network administrators

organizational personnel with information security responsibilities

Mechanisms implementing restrictions on the use of portable storage devices

### AC-21: Information Sharing

Enable authorized users to determine whether access authorizations assigned to a sharing partner match the information’s access and use restrictions for information-sharing circumstances where user discretion is required to determine whether access authorizations assigned to a sharing partner match the information’s access and use restrictions are defined; ; and

Employ automated mechanisms or manual processes that assist users in making information-sharing and collaboration decisions are defined; to assist users in making information sharing and collaboration decisions.

Information sharing applies to information that may be restricted in some manner based on some formal or administrative determination. Examples of such information include, contract-sensitive information, classified information related to special access programs or compartments, privileged information, proprietary information, and personally identifiable information. Security and privacy risk assessments as well as applicable laws, regulations, and policies can provide useful inputs to these determinations. Depending on the circumstances, sharing partners may be defined at the individual, group, or organizational level. Information may be defined by content, type, security category, or special access program or compartment. Access restrictions may include non-disclosure agreements (NDA). Information flow techniques and security attributes may be used to provide automated assistance to users making sharing and collaboration decisions.

authorized users are enabled to determine whether access authorizations assigned to a sharing partner match the information’s access and use restrictions for information-sharing circumstances where user discretion is required to determine whether access authorizations assigned to a sharing partner match the information’s access and use restrictions are defined;;

 automated mechanisms or manual processes that assist users in making information-sharing and collaboration decisions are defined; are employed to assist users in making information-sharing and collaboration decisions.

Access control policy

procedures addressing user-based collaboration and information sharing (including restrictions)

system design documentation

system configuration settings and associated documentation

list of users authorized to make information-sharing/collaboration decisions

list of information-sharing circumstances requiring user discretion

non-disclosure agreements

acquisitions/contractual agreements

system security plan

privacy plan

privacy impact assessment

security and privacy risk assessments

other relevant documents or records

Organizational personnel responsible for information-sharing/collaboration decisions

organizational personnel with responsibility for acquisitions/contractual agreements

system/network administrators

organizational personnel with information security and privacy responsibilities

Automated mechanisms or manual process implementing access authorizations supporting information-sharing/user collaboration decisions

