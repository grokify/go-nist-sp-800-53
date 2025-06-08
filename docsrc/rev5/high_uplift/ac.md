# AC - Access Control

* Controls Count: 7
* Controls IDs: AC-10, AC-18 (4), AC-18 (5), AC-2 (11), AC-2 (12), AC-4 (4), AC-6 (3)

## Controls

### AC-2 (11): Usage Conditions

Enforce circumstances and/or usage conditions to be enforced for system accounts are defined; for system accounts subject to enforcement of circumstances and/or usage conditions are defined;.

Specifying and enforcing usage conditions helps to enforce the principle of least privilege, increase user accountability, and enable effective account monitoring. Account monitoring includes alerts generated if the account is used in violation of organizational parameters. Organizations can describe specific conditions or circumstances under which system accounts can be used, such as by restricting usage to certain days of the week, time of day, or specific durations of time.

 circumstances and/or usage conditions to be enforced for system accounts are defined; for system accounts subject to enforcement of circumstances and/or usage conditions are defined; are enforced.

Access control policy

procedures addressing account management

system design documentation

system configuration settings and associated documentation

system-generated list of system accounts and associated assignments of usage circumstances and/or usage conditions

system audit records

system security plan

other relevant documents or records

Organizational personnel with account management responsibilities

system/network administrators

organizational personnel with information security responsibilities

system developers

Mechanisms implementing account management functions

### AC-2 (12): Account Monitoring for Atypical Usage

Monitor system accounts for atypical usage for which to monitor system accounts is defined; ; and

Report atypical usage of system accounts to personnel or roles to report atypical usage is/are defined;.

Atypical usage includes accessing systems at certain times of the day or from locations that are not consistent with the normal usage patterns of individuals. Monitoring for atypical usage may reveal rogue behavior by individuals or an attack in progress. Account monitoring may inadvertently create privacy risks since data collected to identify atypical usage may reveal previously unknown information about the behavior of individuals. Organizations assess and document privacy risks from monitoring accounts for atypical usage in their privacy impact assessment and make determinations that are in alignment with their privacy program plan.

system accounts are monitored for atypical usage for which to monitor system accounts is defined;; 

atypical usage of system accounts is reported to personnel or roles to report atypical usage is/are defined;.

Access control policy

procedures addressing account management

system design documentation

system configuration settings and associated documentation

system monitoring records

system audit records

audit tracking and monitoring reports

privacy impact assessment

system security plan

privacy plan

other relevant documents or records

Organizational personnel with account management responsibilities

system/network administrators

organizational personnel with information security responsibilities

Mechanisms implementing account management functions

### AC-4 (4): Flow Control of Encrypted Information

Prevent encrypted information from bypassing information flow control mechanisms that encrypted information is prevented from bypassing are defined; by decrypting the information, blocking the flow of the encrypted information, terminating communications sessions attempting to pass encrypted information, and/or  the organization-defined procedure or method used to prevent encrypted information from bypassing information flow control mechanisms is defined (if selected);.

Flow control mechanisms include content checking, security policy filters, and data type identifiers. The term encryption is extended to cover encoded data not recognized by filtering mechanisms.

encrypted information is prevented from bypassing information flow control mechanisms that encrypted information is prevented from bypassing are defined; by decrypting the information, blocking the flow of the encrypted information, terminating communications sessions attempting to pass encrypted information, and/or  the organization-defined procedure or method used to prevent encrypted information from bypassing information flow control mechanisms is defined (if selected);.

Access control policy

information flow control policies

procedures addressing information flow enforcement

system design documentation

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

system developers

Mechanisms implementing information flow enforcement policy

### AC-6 (3): Network Access to Privileged Commands

Authorize network access to privileged commands to which network access is to be authorized only for compelling operational needs are defined; only for compelling operational needs necessitating network access to privileged commands are defined; and document the rationale for such access in the security plan for the system.

Network access is any access across a network connection in lieu of local access (i.e., user being physically present at the device).

network access to privileged commands to which network access is to be authorized only for compelling operational needs are defined; is authorized only for compelling operational needs necessitating network access to privileged commands are defined;;

the rationale for authorizing network access to privileged commands is documented in the security plan for the system.

Access control policy

procedures addressing least privilege

system configuration settings and associated documentation

system audit records

list of operational needs for authorizing network access to privileged commands

system security plan

other relevant documents or records

Organizational personnel with responsibilities for defining least privileges necessary to accomplish specified tasks

organizational personnel with information security responsibilities

Mechanisms implementing least privilege functions

### AC-10: Concurrent Session Control

Limit the number of concurrent sessions for each accounts and/or account types for which to limit the number of concurrent sessions is defined; to the number of concurrent sessions to be allowed for each account and/or account type is defined;.

Organizations may define the maximum number of concurrent sessions for system accounts globally, by account type, by account, or any combination thereof. For example, organizations may limit the number of concurrent sessions for system administrators or other individuals working in particularly sensitive domains or mission-critical applications. Concurrent session control addresses concurrent sessions for system accounts. It does not, however, address concurrent sessions by single users via multiple system accounts.

the number of concurrent sessions for each accounts and/or account types for which to limit the number of concurrent sessions is defined; is limited to the number of concurrent sessions to be allowed for each account and/or account type is defined;.

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

### AC-18 (4): Restrict Configurations by Users

Identify and explicitly authorize users allowed to independently configure wireless networking capabilities.

Organizational authorizations to allow selected users to configure wireless networking capabilities are enforced, in part, by the access enforcement mechanisms employed within organizational systems.

users allowed to independently configure wireless networking capabilities are identified;

users allowed to independently configure wireless networking capabilities are explicitly authorized.

Access control policy

procedures addressing wireless implementation and usage (including restrictions)

system design documentation

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

Mechanisms authorizing independent user configuration of wireless networking capabilities

### AC-18 (5): Antennas and Transmission Power Levels

Select radio antennas and calibrate transmission power levels to reduce the probability that signals from wireless access points can be received outside of organization-controlled boundaries.

Actions that may be taken to limit unauthorized use of wireless communications outside of organization-controlled boundaries include reducing the power of wireless transmissions so that the transmissions are less likely to emit a signal that can be captured outside of the physical perimeters of the organization, employing measures such as emissions security to control wireless emanations, and using directional or beamforming antennas that reduce the likelihood that unintended receivers will be able to intercept signals. Prior to taking such mitigating actions, organizations can conduct periodic wireless surveys to understand the radio frequency profile of organizational systems as well as other systems that may be operating in the area.

radio antennas are selected to reduce the probability that signals from wireless access points can be received outside of organization-controlled boundaries;

transmission power levels are calibrated to reduce the probability that signals from wireless access points can be received outside of organization-controlled boundaries.

Access control policy

procedures addressing wireless implementation and usage (including restrictions)

system design documentation

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

Calibration of transmission power levels for wireless access

radio antenna signals for wireless access

wireless access reception outside of organization-controlled boundaries

