# ac - Access Control

## Base Controls

### ac-1: Policy and Procedures

Develop, document, and disseminate to {{ insert: param, ac-1_prm_1 }}:

 {{ insert: param, ac-01_odp.03 }} access control policy that:

Addresses purpose, scope, roles, responsibilities, management commitment, coordination among organizational entities, and compliance; and

Is consistent with applicable laws, executive orders, directives, regulations, policies, standards, and guidelines; and

Procedures to facilitate the implementation of the access control policy and the associated access controls;

Designate an {{ insert: param, ac-01_odp.04 }} to manage the development, documentation, and dissemination of the access control policy and procedures; and

Review and update the current access control:

Policy {{ insert: param, ac-01_odp.05 }} and following {{ insert: param, ac-01_odp.06 }} ; and

Procedures {{ insert: param, ac-01_odp.07 }} and following {{ insert: param, ac-01_odp.08 }}.

Access control policy and procedures address the controls in the AC family that are implemented within systems and organizations. The risk management strategy is an important factor in establishing such policies and procedures. Policies and procedures contribute to security and privacy assurance. Therefore, it is important that security and privacy programs collaborate on the development of access control policy and procedures. Security and privacy program policies and procedures at the organization level are preferable, in general, and may obviate the need for mission- or system-specific policies and procedures. The policy can be included as part of the general security and privacy policy or be represented by multiple policies reflecting the complex nature of organizations. Procedures can be established for security and privacy programs, for mission or business processes, and for systems, if needed. Procedures describe how the policies or controls are implemented and can be directed at the individual or role that is the object of the procedure. Procedures can be documented in system security and privacy plans or in one or more separate documents. Events that may precipitate an update to access control policy and procedures include assessment or audit findings, security incidents or breaches, or changes in laws, executive orders, directives, regulations, policies, standards, and guidelines. Simply restating controls does not constitute an organizational policy or procedure.

an access control policy is developed and documented;

the access control policy is disseminated to {{ insert: param, ac-01_odp.01 }};

access control procedures to facilitate the implementation of the access control policy and associated controls are developed and documented;

the access control procedures are disseminated to {{ insert: param, ac-01_odp.02 }};

the {{ insert: param, ac-01_odp.03 }} access control policy addresses purpose;

the {{ insert: param, ac-01_odp.03 }} access control policy addresses scope;

the {{ insert: param, ac-01_odp.03 }} access control policy addresses roles;

the {{ insert: param, ac-01_odp.03 }} access control policy addresses responsibilities;

the {{ insert: param, ac-01_odp.03 }} access control policy addresses management commitment;

the {{ insert: param, ac-01_odp.03 }} access control policy addresses coordination among organizational entities;

the {{ insert: param, ac-01_odp.03 }} access control policy addresses compliance;

the {{ insert: param, ac-01_odp.03 }} access control policy is consistent with applicable laws, Executive Orders, directives, regulations, policies, standards, and guidelines;

the {{ insert: param, ac-01_odp.04 }} is designated to manage the development, documentation, and dissemination of the access control policy and procedures;

the current access control policy is reviewed and updated {{ insert: param, ac-01_odp.05 }};

the current access control policy is reviewed and updated following {{ insert: param, ac-01_odp.06 }};

the current access control procedures are reviewed and updated {{ insert: param, ac-01_odp.07 }};

the current access control procedures are reviewed and updated following {{ insert: param, ac-01_odp.08 }}.

Access control policy and procedures

system security plan

privacy plan

other relevant documents or records

Organizational personnel with access control responsibilities

organizational personnel with information security with information security and privacy responsibilities

### ac-2: Account Management

Define and document the types of accounts allowed and specifically prohibited for use within the system;

Assign account managers;

Require {{ insert: param, ac-02_odp.01 }} for group and role membership;

Specify:

Authorized users of the system;

Group and role membership; and

Access authorizations (i.e., privileges) and {{ insert: param, ac-02_odp.02 }} for each account;

Require approvals by {{ insert: param, ac-02_odp.03 }} for requests to create accounts;

Create, enable, modify, disable, and remove accounts in accordance with {{ insert: param, ac-02_odp.04 }};

Monitor the use of accounts;

Notify account managers and {{ insert: param, ac-02_odp.05 }} within:

 {{ insert: param, ac-02_odp.06 }} when accounts are no longer required;

 {{ insert: param, ac-02_odp.07 }} when users are terminated or transferred; and

 {{ insert: param, ac-02_odp.08 }} when system usage or need-to-know changes for an individual;

Authorize access to the system based on:

A valid access authorization;

Intended system usage; and

 {{ insert: param, ac-02_odp.09 }};

Review accounts for compliance with account management requirements {{ insert: param, ac-02_odp.10 }};

Establish and implement a process for changing shared or group account authenticators (if deployed) when individuals are removed from the group; and

Align account management processes with personnel termination and transfer processes.

Examples of system account types include individual, shared, group, system, guest, anonymous, emergency, developer, temporary, and service. Identification of authorized system users and the specification of access privileges reflect the requirements in other controls in the security plan. Users requiring administrative privileges on system accounts receive additional scrutiny by organizational personnel responsible for approving such accounts and privileged access, including system owner, mission or business owner, senior agency information security officer, or senior agency official for privacy. Types of accounts that organizations may wish to prohibit due to increased risk include shared, group, emergency, anonymous, temporary, and guest accounts.

Where access involves personally identifiable information, security programs collaborate with the senior agency official for privacy to establish the specific conditions for group and role membership; specify authorized users, group and role membership, and access authorizations for each account; and create, adjust, or remove system accounts in accordance with organizational policies. Policies can include such information as account expiration dates or other factors that trigger the disabling of accounts. Organizations may choose to define access privileges or other attributes by account, type of account, or a combination of the two. Examples of other attributes required for authorizing access include restrictions on time of day, day of week, and point of origin. In defining other system account attributes, organizations consider system-related requirements and mission/business requirements. Failure to consider these factors could affect system availability.

Temporary and emergency accounts are intended for short-term use. Organizations establish temporary accounts as part of normal account activation procedures when there is a need for short-term accounts without the demand for immediacy in account activation. Organizations establish emergency accounts in response to crisis situations and with the need for rapid account activation. Therefore, emergency account activation may bypass normal account authorization processes. Emergency and temporary accounts are not to be confused with infrequently used accounts, including local logon accounts used for special tasks or when network resources are unavailable (may also be known as accounts of last resort). Such accounts remain available and are not subject to automatic disabling or removal dates. Conditions for disabling or deactivating accounts include when shared/group, emergency, or temporary accounts are no longer required and when individuals are transferred or terminated. Changing shared/group authenticators when members leave the group is intended to ensure that former group members do not retain access to the shared or group account. Some types of system accounts may require specialized training.

account types allowed for use within the system are defined and documented;

account types specifically prohibited for use within the system are defined and documented;

account managers are assigned;

 {{ insert: param, ac-02_odp.01 }} for group and role membership are required;

authorized users of the system are specified;

group and role membership are specified;

access authorizations (i.e., privileges) are specified for each account;

 {{ insert: param, ac-02_odp.02 }} are specified for each account;

approvals are required by {{ insert: param, ac-02_odp.03 }} for requests to create accounts;

accounts are created in accordance with {{ insert: param, ac-02_odp.04 }};

accounts are enabled in accordance with {{ insert: param, ac-02_odp.04 }};

accounts are modified in accordance with {{ insert: param, ac-02_odp.04 }};

accounts are disabled in accordance with {{ insert: param, ac-02_odp.04 }};

accounts are removed in accordance with {{ insert: param, ac-02_odp.04 }};

the use of accounts is monitored; 

account managers and {{ insert: param, ac-02_odp.05 }} are notified within {{ insert: param, ac-02_odp.06 }} when accounts are no longer required;

account managers and {{ insert: param, ac-02_odp.05 }} are notified within {{ insert: param, ac-02_odp.07 }} when users are terminated or transferred;

account managers and {{ insert: param, ac-02_odp.05 }} are notified within {{ insert: param, ac-02_odp.08 }} when system usage or the need to know changes for an individual;

access to the system is authorized based on a valid access authorization;

access to the system is authorized based on intended system usage;

access to the system is authorized based on {{ insert: param, ac-02_odp.09 }};

accounts are reviewed for compliance with account management requirements {{ insert: param, ac-02_odp.10 }};

a process is established for changing shared or group account authenticators (if deployed) when individuals are removed from the group;

a process is implemented for changing shared or group account authenticators (if deployed) when individuals are removed from the group;

account management processes are aligned with personnel termination processes;

account management processes are aligned with personnel transfer processes.

Access control policy

personnel termination policy and procedure

personnel transfer policy and procedure

procedures for addressing account management

system design documentation

system configuration settings and associated documentation

list of active system accounts along with the name of the individual associated with each account

list of recently disabled system accounts and the name of the individual associated with each account

list of conditions for group and role membership

notifications of recent transfers, separations, or terminations of employees

access authorization records

account management compliance reviews

system monitoring records

system audit records

system security plan

privacy plan

other relevant documents or records

Organizational personnel with account management responsibilities

system/network administrators

organizational personnel with information security with information security and privacy responsibilities

Organizational processes for account management on the system

mechanisms for implementing account management

### ac-3: Access Enforcement

Enforce approved authorizations for logical access to information and system resources in accordance with applicable access control policies.

Access control policies control access between active entities or subjects (i.e., users or processes acting on behalf of users) and passive entities or objects (i.e., devices, files, records, domains) in organizational systems. In addition to enforcing authorized access at the system level and recognizing that systems can host many applications and services in support of mission and business functions, access enforcement mechanisms can also be employed at the application and service level to provide increased information security and privacy. In contrast to logical access controls that are implemented within the system, physical access controls are addressed by the controls in the Physical and Environmental Protection ( [PE](#pe) ) family.

approved authorizations for logical access to information and system resources are enforced in accordance with applicable access control policies.

Access control policy

procedures addressing access enforcement

system design documentation

system configuration settings and associated documentation

list of approved authorizations (user privileges)

system audit records

system security plan

privacy plan

other relevant documents or records

Organizational personnel with access enforcement responsibilities

system/network administrators

organizational personnel with information security and privacy responsibilities

system developers

Mechanisms implementing access control policy

### ac-4: Information Flow Enforcement

Enforce approved authorizations for controlling the flow of information within the system and between connected systems based on {{ insert: param, ac-04_odp }}.

Information flow control regulates where information can travel within a system and between systems (in contrast to who is allowed to access the information) and without regard to subsequent accesses to that information. Flow control restrictions include blocking external traffic that claims to be from within the organization, keeping export-controlled information from being transmitted in the clear to the Internet, restricting web requests that are not from the internal web proxy server, and limiting information transfers between organizations based on data structures and content. Transferring information between organizations may require an agreement specifying how the information flow is enforced (see [CA-3](#ca-3) ). Transferring information between systems in different security or privacy domains with different security or privacy policies introduces the risk that such transfers violate one or more domain security or privacy policies. In such situations, information owners/stewards provide guidance at designated policy enforcement points between connected systems. Organizations consider mandating specific architectural solutions to enforce specific security and privacy policies. Enforcement includes prohibiting information transfers between connected systems (i.e., allowing access only), verifying write permissions before accepting information from another security or privacy domain or connected system, employing hardware mechanisms to enforce one-way information flows, and implementing trustworthy regrading mechanisms to reassign security or privacy attributes and labels.

Organizations commonly employ information flow control policies and enforcement mechanisms to control the flow of information between designated sources and destinations within systems and between connected systems. Flow control is based on the characteristics of the information and/or the information path. Enforcement occurs, for example, in boundary protection devices that employ rule sets or establish configuration settings that restrict system services, provide a packet-filtering capability based on header information, or provide a message-filtering capability based on message content. Organizations also consider the trustworthiness of filtering and/or inspection mechanisms (i.e., hardware, firmware, and software components) that are critical to information flow enforcement. Control enhancements 3 through 32 primarily address cross-domain solution needs that focus on more advanced filtering techniques, in-depth analysis, and stronger flow enforcement mechanisms implemented in cross-domain products, such as high-assurance guards. Such capabilities are generally not available in commercial off-the-shelf products. Information flow enforcement also applies to control plane traffic (e.g., routing and DNS).

approved authorizations are enforced for controlling the flow of information within the system and between connected systems based on {{ insert: param, ac-04_odp }}.

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

### ac-5: Separation of Duties

Identify and document {{ insert: param, ac-05_odp }} ; and

Define system access authorizations to support separation of duties.

Separation of duties addresses the potential for abuse of authorized privileges and helps to reduce the risk of malevolent activity without collusion. Separation of duties includes dividing mission or business functions and support functions among different individuals or roles, conducting system support functions with different individuals, and ensuring that security personnel who administer access control functions do not also administer audit functions. Because separation of duty violations can span systems and application domains, organizations consider the entirety of systems and system components when developing policy on separation of duties. Separation of duties is enforced through the account management activities in [AC-2](#ac-2) , access control mechanisms in [AC-3](#ac-3) , and identity management activities in [IA-2](#ia-2), [IA-4](#ia-4) , and [IA-12](#ia-12).

 {{ insert: param, ac-05_odp }} are identified and documented;

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

### ac-6: Least Privilege

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

### ac-7: Unsuccessful Logon Attempts

Enforce a limit of {{ insert: param, ac-07_odp.01 }} consecutive invalid logon attempts by a user during a {{ insert: param, ac-07_odp.02 }} ; and

Automatically {{ insert: param, ac-07_odp.03 }} when the maximum number of unsuccessful attempts is exceeded.

The need to limit unsuccessful logon attempts and take subsequent action when the maximum number of attempts is exceeded applies regardless of whether the logon occurs via a local or network connection. Due to the potential for denial of service, automatic lockouts initiated by systems are usually temporary and automatically release after a predetermined, organization-defined time period. If a delay algorithm is selected, organizations may employ different algorithms for different components of the system based on the capabilities of those components. Responses to unsuccessful logon attempts may be implemented at the operating system and the application levels. Organization-defined actions that may be taken when the number of allowed consecutive invalid logon attempts is exceeded include prompting the user to answer a secret question in addition to the username and password, invoking a lockdown mode with limited user capabilities (instead of full lockout), allowing users to only logon from specified Internet Protocol (IP) addresses, requiring a CAPTCHA to prevent automated attacks, or applying user profiles such as location, time of day, IP address, device, or Media Access Control (MAC) address. If automatic system lockout or execution of a delay algorithm is not implemented in support of the availability objective, organizations consider a combination of other actions to help prevent brute force attacks. In addition to the above, organizations can prompt users to respond to a secret question before the number of allowed unsuccessful logon attempts is exceeded. Automatically unlocking an account after a specified period of time is generally not permitted. However, exceptions may be required based on operational mission or need.

a limit of {{ insert: param, ac-07_odp.01 }} consecutive invalid logon attempts by a user during {{ insert: param, ac-07_odp.02 }} is enforced;

automatically {{ insert: param, ac-07_odp.03 }} when the maximum number of unsuccessful attempts is exceeded.

Access control policy

procedures addressing unsuccessful logon attempts

system design documentation

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

Organizational personnel with information security responsibilities

system developers

system/network administrators

Mechanisms implementing access control policy for unsuccessful logon attempts

### ac-8: System Use Notification

Display {{ insert: param, ac-08_odp.01 }} to users before granting access to the system that provides privacy and security notices consistent with applicable laws, executive orders, directives, regulations, policies, standards, and guidelines and state that:

Users are accessing a U.S. Government system;

System usage may be monitored, recorded, and subject to audit;

Unauthorized use of the system is prohibited and subject to criminal and civil penalties; and

Use of the system indicates consent to monitoring and recording;

Retain the notification message or banner on the screen until users acknowledge the usage conditions and take explicit actions to log on to or further access the system; and

For publicly accessible systems:

Display system use information {{ insert: param, ac-08_odp.02 }} , before granting further access to the publicly accessible system;

Display references, if any, to monitoring, recording, or auditing that are consistent with privacy accommodations for such systems that generally prohibit those activities; and

Include a description of the authorized uses of the system.

System use notifications can be implemented using messages or warning banners displayed before individuals log in to systems. System use notifications are used only for access via logon interfaces with human users. Notifications are not required when human interfaces do not exist. Based on an assessment of risk, organizations consider whether or not a secondary system use notification is needed to access applications or other system resources after the initial network logon. Organizations consider system use notification messages or banners displayed in multiple languages based on organizational needs and the demographics of system users. Organizations consult with the privacy office for input regarding privacy messaging and the Office of the General Counsel or organizational equivalent for legal review and approval of warning banner content.

 {{ insert: param, ac-08_odp.01 }} is displayed to users before granting access to the system that provides privacy and security notices consistent with applicable laws, Executive Orders, directives, regulations, policies, standards, and guidelines;

the system use notification states that users are accessing a U.S. Government system;

the system use notification states that system usage may be monitored, recorded, and subject to audit;

the system use notification states that unauthorized use of the system is prohibited and subject to criminal and civil penalties; and

the system use notification states that use of the system indicates consent to monitoring and recording;

the notification message or banner is retained on the screen until users acknowledge the usage conditions and take explicit actions to log on to or further access the system;

for publicly accessible systems, system use information {{ insert: param, ac-08_odp.02 }} is displayed before granting further access to the publicly accessible system;

for publicly accessible systems, any references to monitoring, recording, or auditing that are consistent with privacy accommodations for such systems that generally prohibit those activities are displayed;

for publicly accessible systems, a description of the authorized uses of the system is included.

Access control policy

privacy and security policies, procedures addressing system use notification

documented approval of system use notification messages or banners

system audit records

user acknowledgements of notification message or banner

system design documentation

system configuration settings and associated documentation

system use notification messages

system security plan

privacy plan

privacy impact assessment

privacy assessment report

other relevant documents or records

System/network administrators

organizational personnel with information security and privacy responsibilities

legal counsel

system developers

Mechanisms implementing system use notification

### ac-10: Concurrent Session Control

Limit the number of concurrent sessions for each {{ insert: param, ac-10_odp.01 }} to {{ insert: param, ac-10_odp.02 }}.

Organizations may define the maximum number of concurrent sessions for system accounts globally, by account type, by account, or any combination thereof. For example, organizations may limit the number of concurrent sessions for system administrators or other individuals working in particularly sensitive domains or mission-critical applications. Concurrent session control addresses concurrent sessions for system accounts. It does not, however, address concurrent sessions by single users via multiple system accounts.

the number of concurrent sessions for each {{ insert: param, ac-10_odp.01 }} is limited to {{ insert: param, ac-10_odp.02 }}.

Access control policy

procedures addressing concurrent session control

system design documentation

system configuration settings and associated documentation

security plan

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

system developers

Mechanisms implementing access control policy for concurrent session control

### ac-11: Device Lock

Prevent further access to the system by {{ insert: param, ac-11_odp.01 }} ; and

Retain the device lock until the user reestablishes access using established identification and authentication procedures.

Device locks are temporary actions taken to prevent logical access to organizational systems when users stop work and move away from the immediate vicinity of those systems but do not want to log out because of the temporary nature of their absences. Device locks can be implemented at the operating system level or at the application level. A proximity lock may be used to initiate the device lock (e.g., via a Bluetooth-enabled device or dongle). User-initiated device locking is behavior or policy-based and, as such, requires users to take physical action to initiate the device lock. Device locks are not an acceptable substitute for logging out of systems, such as when organizations require users to log out at the end of workdays.

further access to the system is prevented by {{ insert: param, ac-11_odp.01 }};

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

### ac-12: Session Termination

Automatically terminate a user session after {{ insert: param, ac-12_odp }}.

Session termination addresses the termination of user-initiated logical sessions (in contrast to [SC-10](#sc-10) , which addresses the termination of network connections associated with communications sessions (i.e., network disconnect)). A logical session (for local, network, and remote access) is initiated whenever a user (or process acting on behalf of a user) accesses an organizational system. Such user sessions can be terminated without terminating network sessions. Session termination ends all processes associated with a user’s logical session except for those processes that are specifically created by the user (i.e., session owner) to continue after the session is terminated. Conditions or trigger events that require automatic termination of the session include organization-defined periods of user inactivity, targeted responses to certain types of incidents, or time-of-day restrictions on system use.

a user session is automatically terminated after {{ insert: param, ac-12_odp }}.

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

### ac-14: Permitted Actions Without Identification or Authentication

Identify {{ insert: param, ac-14_odp }} that can be performed on the system without identification or authentication consistent with organizational mission and business functions; and

Document and provide supporting rationale in the security plan for the system, user actions not requiring identification or authentication.

Specific user actions may be permitted without identification or authentication if organizations determine that identification and authentication are not required for the specified user actions. Organizations may allow a limited number of user actions without identification or authentication, including when individuals access public websites or other publicly accessible federal systems, when individuals use mobile phones to receive calls, or when facsimiles are received. Organizations identify actions that normally require identification or authentication but may, under certain circumstances, allow identification or authentication mechanisms to be bypassed. Such bypasses may occur, for example, via a software-readable physical switch that commands bypass of the logon functionality and is protected from accidental or unmonitored use. Permitting actions without identification or authentication does not apply to situations where identification and authentication have already occurred and are not repeated but rather to situations where identification and authentication have not yet occurred. Organizations may decide that there are no user actions that can be performed on organizational systems without identification and authentication, and therefore, the value for the assignment operation can be "none." 

 {{ insert: param, ac-14_odp }} that can be performed on the system without identification or authentication consistent with organizational mission and business functions are identified;

user actions not requiring identification or authentication are documented in the security plan for the system;

a rationale for user actions not requiring identification or authentication is provided in the security plan for the system.

Access control policy

procedures addressing permitted actions without identification or authentication

system configuration settings and associated documentation

security plan

list of user actions that can be performed without identification or authentication

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

### ac-17: Remote Access

Establish and document usage restrictions, configuration/connection requirements, and implementation guidance for each type of remote access allowed; and

Authorize each type of remote access to the system prior to allowing such connections.

Remote access is access to organizational systems (or processes acting on behalf of users) that communicate through external networks such as the Internet. Types of remote access include dial-up, broadband, and wireless. Organizations use encrypted virtual private networks (VPNs) to enhance confidentiality and integrity for remote connections. The use of encrypted VPNs provides sufficient assurance to the organization that it can effectively treat such connections as internal networks if the cryptographic mechanisms used are implemented in accordance with applicable laws, executive orders, directives, regulations, policies, standards, and guidelines. Still, VPN connections traverse external networks, and the encrypted VPN does not enhance the availability of remote connections. VPNs with encrypted tunnels can also affect the ability to adequately monitor network communications traffic for malicious code. Remote access controls apply to systems other than public web servers or systems designed for public access. Authorization of each remote access type addresses authorization prior to allowing remote access without specifying the specific formats for such authorization. While organizations may use information exchange and system connection security agreements to manage remote access connections to other systems, such agreements are addressed as part of [CA-3](#ca-3) . Enforcing access restrictions for remote access is addressed via [AC-3](#ac-3).

usage restrictions are established and documented for each type of remote access allowed;

configuration/connection requirements are established and documented for each type of remote access allowed;

implementation guidance is established and documented for each type of remote access allowed;

each type of remote access to the system is authorized prior to allowing such connections.

Access control policy

procedures addressing remote access implementation and usage (including restrictions)

configuration management plan

system configuration settings and associated documentation

remote access authorizations

system audit records

system security plan

other relevant documents or records

Organizational personnel with responsibilities for managing remote access connections

system/network administrators

organizational personnel with information security responsibilities

Remote access management capability for the system

### ac-18: Wireless Access

Establish configuration requirements, connection requirements, and implementation guidance for each type of wireless access; and

Authorize each type of wireless access to the system prior to allowing such connections.

Wireless technologies include microwave, packet radio (ultra-high frequency or very high frequency), 802.11x, and Bluetooth. Wireless networks use authentication protocols that provide authenticator protection and mutual authentication.

configuration requirements are established for each type of wireless access;

connection requirements are established for each type of wireless access;

implementation guidance is established for each type of wireless access;

each type of wireless access to the system is authorized prior to allowing such connections.

Access control policy

procedures addressing wireless access implementation and usage (including restrictions)

configuration management plan

system design documentation

system configuration settings and associated documentation

wireless access authorizations

system audit records

system security plan

other relevant documents or records

Organizational personnel with responsibilities for managing wireless access connections

organizational personnel with information security responsibilities

Wireless access management capability for the system

### ac-19: Access Control for Mobile Devices

Establish configuration requirements, connection requirements, and implementation guidance for organization-controlled mobile devices, to include when such devices are outside of controlled areas; and

Authorize the connection of mobile devices to organizational systems.

A mobile device is a computing device that has a small form factor such that it can easily be carried by a single individual; is designed to operate without a physical connection; possesses local, non-removable or removable data storage; and includes a self-contained power source. Mobile device functionality may also include voice communication capabilities, on-board sensors that allow the device to capture information, and/or built-in features for synchronizing local data with remote locations. Examples include smart phones and tablets. Mobile devices are typically associated with a single individual. The processing, storage, and transmission capability of the mobile device may be comparable to or merely a subset of notebook/desktop systems, depending on the nature and intended purpose of the device. Protection and control of mobile devices is behavior or policy-based and requires users to take physical action to protect and control such devices when outside of controlled areas. Controlled areas are spaces for which organizations provide physical or procedural controls to meet the requirements established for protecting information and systems.

Due to the large variety of mobile devices with different characteristics and capabilities, organizational restrictions may vary for the different classes or types of such devices. Usage restrictions and specific implementation guidance for mobile devices include configuration management, device identification and authentication, implementation of mandatory protective software, scanning devices for malicious code, updating virus protection software, scanning for critical software updates and patches, conducting primary operating system (and possibly other resident software) integrity checks, and disabling unnecessary hardware.

Usage restrictions and authorization to connect may vary among organizational systems. For example, the organization may authorize the connection of mobile devices to its network and impose a set of usage restrictions, while a system owner may withhold authorization for mobile device connection to specific applications or impose additional usage restrictions before allowing mobile device connections to a system. Adequate security for mobile devices goes beyond the requirements specified in [AC-19](#ac-19) . Many safeguards for mobile devices are reflected in other controls. [AC-20](#ac-20) addresses mobile devices that are not organization-controlled.

configuration requirements are established for organization-controlled mobile devices, including when such devices are outside of the controlled area;

connection requirements are established for organization-controlled mobile devices, including when such devices are outside of the controlled area;

implementation guidance is established for organization-controlled mobile devices, including when such devices are outside of the controlled area;

the connection of mobile devices to organizational systems is authorized.

Access control policy

procedures addressing access control for mobile device usage (including restrictions)

configuration management plan

system design documentation

system configuration settings and associated documentation

authorizations for mobile device connections to organizational systems

system audit records

system security plan

other relevant documents or records

Organizational personnel using mobile devices to access organizational systems

system/network administrators

organizational personnel with information security responsibilities

Access control capability for mobile device connections to organizational systems

configurations of mobile devices

### ac-20: Use of External Systems

 {{ insert: param, ac-20_odp.01 }} , consistent with the trust relationships established with other organizations owning, operating, and/or maintaining external systems, allowing authorized individuals to:

Access the system from external systems; and

Process, store, or transmit organization-controlled information using external systems; or

Prohibit the use of {{ insert: param, ac-20_odp.04 }}.

External systems are systems that are used by but not part of organizational systems, and for which the organization has no direct control over the implementation of required controls or the assessment of control effectiveness. External systems include personally owned systems, components, or devices; privately owned computing and communications devices in commercial or public facilities; systems owned or controlled by nonfederal organizations; systems managed by contractors; and federal information systems that are not owned by, operated by, or under the direct supervision or authority of the organization. External systems also include systems owned or operated by other components within the same organization and systems within the organization with different authorization boundaries. Organizations have the option to prohibit the use of any type of external system or prohibit the use of specified types of external systems, (e.g., prohibit the use of any external system that is not organizationally owned or prohibit the use of personally-owned systems).

For some external systems (i.e., systems operated by other organizations), the trust relationships that have been established between those organizations and the originating organization may be such that no explicit terms and conditions are required. Systems within these organizations may not be considered external. These situations occur when, for example, there are pre-existing information exchange agreements (either implicit or explicit) established between organizations or components or when such agreements are specified by applicable laws, executive orders, directives, regulations, policies, or standards. Authorized individuals include organizational personnel, contractors, or other individuals with authorized access to organizational systems and over which organizations have the authority to impose specific rules of behavior regarding system access. Restrictions that organizations impose on authorized individuals need not be uniform, as the restrictions may vary depending on trust relationships between organizations. Therefore, organizations may choose to impose different security restrictions on contractors than on state, local, or tribal governments.

External systems used to access public interfaces to organizational systems are outside the scope of [AC-20](#ac-20) . Organizations establish specific terms and conditions for the use of external systems in accordance with organizational security policies and procedures. At a minimum, terms and conditions address the specific types of applications that can be accessed on organizational systems from external systems and the highest security category of information that can be processed, stored, or transmitted on external systems. If the terms and conditions with the owners of the external systems cannot be established, organizations may impose restrictions on organizational personnel using those external systems.

 {{ insert: param, ac-20_odp.01 }} is/are consistent with the trust relationships established with other organizations owning, operating, and/or maintaining external systems, allowing authorized individuals to access the system from external systems (if applicable);

 {{ insert: param, ac-20_odp.01 }} is/are consistent with the trust relationships established with other organizations owning, operating, and/or maintaining external systems, allowing authorized individuals to process, store, or transmit organization-controlled information using external systems (if applicable);

the use of {{ insert: param, ac-20_odp.04 }} is prohibited (if applicable).

Access control policy

procedures addressing the use of external systems

external systems terms and conditions

list of types of applications accessible from external systems

maximum security categorization for information processed, stored, or transmitted on external systems

system configuration settings and associated documentation

system security plan

other relevant documents or records

Organizational personnel with responsibilities for defining terms and conditions for use of external systems to access organizational systems

system/network administrators

organizational personnel with information security responsibilities

Mechanisms implementing terms and conditions on the use of external systems

### ac-21: Information Sharing

Enable authorized users to determine whether access authorizations assigned to a sharing partner match the information’s access and use restrictions for {{ insert: param, ac-21_odp.01 }} ; and

Employ {{ insert: param, ac-21_odp.02 }} to assist users in making information sharing and collaboration decisions.

Information sharing applies to information that may be restricted in some manner based on some formal or administrative determination. Examples of such information include, contract-sensitive information, classified information related to special access programs or compartments, privileged information, proprietary information, and personally identifiable information. Security and privacy risk assessments as well as applicable laws, regulations, and policies can provide useful inputs to these determinations. Depending on the circumstances, sharing partners may be defined at the individual, group, or organizational level. Information may be defined by content, type, security category, or special access program or compartment. Access restrictions may include non-disclosure agreements (NDA). Information flow techniques and security attributes may be used to provide automated assistance to users making sharing and collaboration decisions.

authorized users are enabled to determine whether access authorizations assigned to a sharing partner match the information’s access and use restrictions for {{ insert: param, ac-21_odp.01 }};

 {{ insert: param, ac-21_odp.02 }} are employed to assist users in making information-sharing and collaboration decisions.

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

### ac-22: Publicly Accessible Content

Designate individuals authorized to make information publicly accessible;

Train authorized individuals to ensure that publicly accessible information does not contain nonpublic information;

Review the proposed content of information prior to posting onto the publicly accessible system to ensure that nonpublic information is not included; and

Review the content on the publicly accessible system for nonpublic information {{ insert: param, ac-22_odp }} and remove such information, if discovered.

In accordance with applicable laws, executive orders, directives, policies, regulations, standards, and guidelines, the public is not authorized to have access to nonpublic information, including information protected under the [PRIVACT](#18e71fec-c6fd-475a-925a-5d8495cf8455) and proprietary information. Publicly accessible content addresses systems that are controlled by the organization and accessible to the public, typically without identification or authentication. Posting information on non-organizational systems (e.g., non-organizational public websites, forums, and social media) is covered by organizational policy. While organizations may have individuals who are responsible for developing and implementing policies about the information that can be made publicly accessible, publicly accessible content addresses the management of the individuals who make such information publicly accessible.

designated individuals are authorized to make information publicly accessible;

authorized individuals are trained to ensure that publicly accessible information does not contain non-public information;

the proposed content of information is reviewed prior to posting onto the publicly accessible system to ensure that non-public information is not included;

the content on the publicly accessible system is reviewed for non-public information {{ insert: param, ac-22_odp }};

non-public information is removed from the publicly accessible system, if discovered.

Access control policy

procedures addressing publicly accessible content

list of users authorized to post publicly accessible content on organizational systems

training materials and/or records

records of publicly accessible information reviews

records of response to non-public information on public websites

system audit logs

security awareness training records

system security plan

other relevant documents or records

Organizational personnel with responsibilities for managing publicly accessible information posted on organizational systems

organizational personnel with information security responsibilities

Mechanisms implementing management of publicly accessible content

## Control Enhancements

