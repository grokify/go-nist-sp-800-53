# au - Audit and Accountability

* Controls: 25

## Controls

### {au-01 au %!s(int=1) %!s(int=0)}: Policy and Procedures

Develop, document, and disseminate to {{ insert: param, au-1_prm_1 }}:

 {{ insert: param, au-01_odp.03 }} audit and accountability policy that:

Addresses purpose, scope, roles, responsibilities, management commitment, coordination among organizational entities, and compliance; and

Is consistent with applicable laws, executive orders, directives, regulations, policies, standards, and guidelines; and

Procedures to facilitate the implementation of the audit and accountability policy and the associated audit and accountability controls;

Designate an {{ insert: param, au-01_odp.04 }} to manage the development, documentation, and dissemination of the audit and accountability policy and procedures; and

Review and update the current audit and accountability:

Policy {{ insert: param, au-01_odp.05 }} and following {{ insert: param, au-01_odp.06 }} ; and

Procedures {{ insert: param, au-01_odp.07 }} and following {{ insert: param, au-01_odp.08 }}.

Audit and accountability policy and procedures address the controls in the AU family that are implemented within systems and organizations. The risk management strategy is an important factor in establishing such policies and procedures. Policies and procedures contribute to security and privacy assurance. Therefore, it is important that security and privacy programs collaborate on the development of audit and accountability policy and procedures. Security and privacy program policies and procedures at the organization level are preferable, in general, and may obviate the need for mission- or system-specific policies and procedures. The policy can be included as part of the general security and privacy policy or be represented by multiple policies that reflect the complex nature of organizations. Procedures can be established for security and privacy programs, for mission or business processes, and for systems, if needed. Procedures describe how the policies or controls are implemented and can be directed at the individual or role that is the object of the procedure. Procedures can be documented in system security and privacy plans or in one or more separate documents. Events that may precipitate an update to audit and accountability policy and procedures include assessment or audit findings, security incidents or breaches, or changes in applicable laws, executive orders, directives, regulations, policies, standards, and guidelines. Simply restating controls does not constitute an organizational policy or procedure.

an audit and accountability policy is developed and documented;

the audit and accountability policy is disseminated to {{ insert: param, au-01_odp.01 }};

audit and accountability procedures to facilitate the implementation of the audit and accountability policy and associated audit and accountability controls are developed and documented;

the audit and accountability procedures are disseminated to {{ insert: param, au-01_odp.02 }};

the {{ insert: param, au-01_odp.03 }} of the audit and accountability policy addresses purpose;

the {{ insert: param, au-01_odp.03 }} of the audit and accountability policy addresses scope;

the {{ insert: param, au-01_odp.03 }} of the audit and accountability policy addresses roles;

the {{ insert: param, au-01_odp.03 }} of the audit and accountability policy addresses responsibilities;

the {{ insert: param, au-01_odp.03 }} of the audit and accountability policy addresses management commitment;

the {{ insert: param, au-01_odp.03 }} of the audit and accountability policy addresses coordination among organizational entities;

the {{ insert: param, au-01_odp.03 }} of the audit and accountability policy addresses compliance;

the {{ insert: param, au-01_odp.03 }} of the audit and accountability policy is consistent with applicable laws, executive orders, directives, regulations, policies, standards, and guidelines;

the {{ insert: param, au-01_odp.04 }} is designated to manage the development, documentation, and dissemination of the audit and accountability policy and procedures;

the current audit and accountability policy is reviewed and updated {{ insert: param, au-01_odp.05 }};

the current audit and accountability policy is reviewed and updated following {{ insert: param, au-01_odp.06 }};

the current audit and accountability procedures are reviewed and updated {{ insert: param, au-01_odp.07 }};

the current audit and accountability procedures are reviewed and updated following {{ insert: param, au-01_odp.08 }}.

Audit and accountability policy and procedures

system security plan

privacy plan

other relevant documents or records

Organizational personnel with audit and accountability responsibilities

organizational personnel with information security and privacy responsibilities

### {au-02 au %!s(int=2) %!s(int=0)}: Event Logging

Identify the types of events that the system is capable of logging in support of the audit function: {{ insert: param, au-02_odp.01 }};

Coordinate the event logging function with other organizational entities requiring audit-related information to guide and inform the selection criteria for events to be logged;

Specify the following event types for logging within the system: {{ insert: param, au-2_prm_2 }};

Provide a rationale for why the event types selected for logging are deemed to be adequate to support after-the-fact investigations of incidents; and

Review and update the event types selected for logging {{ insert: param, au-02_odp.04 }}.

An event is an observable occurrence in a system. The types of events that require logging are those events that are significant and relevant to the security of systems and the privacy of individuals. Event logging also supports specific monitoring and auditing needs. Event types include password changes, failed logons or failed accesses related to systems, security or privacy attribute changes, administrative privilege usage, PIV credential usage, data action changes, query parameters, or external credential usage. In determining the set of event types that require logging, organizations consider the monitoring and auditing appropriate for each of the controls to be implemented. For completeness, event logging includes all protocols that are operational and supported by the system.

To balance monitoring and auditing requirements with other system needs, event logging requires identifying the subset of event types that are logged at a given point in time. For example, organizations may determine that systems need the capability to log every file access successful and unsuccessful, but not activate that capability except for specific circumstances due to the potential burden on system performance. The types of events that organizations desire to be logged may change. Reviewing and updating the set of logged events is necessary to help ensure that the events remain relevant and continue to support the needs of the organization. Organizations consider how the types of logging events can reveal information about individuals that may give rise to privacy risk and how best to mitigate such risks. For example, there is the potential to reveal personally identifiable information in the audit trail, especially if the logging event is based on patterns or time of usage.

Event logging requirements, including the need to log specific event types, may be referenced in other controls and control enhancements. These include [AC-2(4)](#ac-2.4), [AC-3(10)](#ac-3.10), [AC-6(9)](#ac-6.9), [AC-17(1)](#ac-17.1), [CM-3f](#cm-3_smt.f), [CM-5(1)](#cm-5.1), [IA-3(3)(b)](#ia-3.3_smt.b), [MA-4(1)](#ma-4.1), [MP-4(2)](#mp-4.2), [PE-3](#pe-3), [PM-21](#pm-21), [PT-7](#pt-7), [RA-8](#ra-8), [SC-7(9)](#sc-7.9), [SC-7(15)](#sc-7.15), [SI-3(8)](#si-3.8), [SI-4(22)](#si-4.22), [SI-7(8)](#si-7.8) , and [SI-10(1)](#si-10.1) . Organizations include event types that are required by applicable laws, executive orders, directives, policies, regulations, standards, and guidelines. Audit records can be generated at various levels, including at the packet level as information traverses the network. Selecting the appropriate level of event logging is an important part of a monitoring and auditing capability and can identify the root causes of problems. When defining event types, organizations consider the logging necessary to cover related event types, such as the steps in distributed, transaction-based processes and the actions that occur in service-oriented architectures.

 {{ insert: param, au-02_odp.01 }} that the system is capable of logging are identified in support of the audit logging function;

the event logging function is coordinated with other organizational entities requiring audit-related information to guide and inform the selection criteria for events to be logged;

 {{ insert: param, au-02_odp.02 }} are specified for logging within the system;

the specified event types are logged within the system {{ insert: param, au-02_odp.03 }};

a rationale is provided for why the event types selected for logging are deemed to be adequate to support after-the-fact investigations of incidents;

the event types selected for logging are reviewed and updated {{ insert: param, au-02_odp.04 }}.

Audit and accountability policy

procedures addressing auditable events

system security plan

privacy plan

system design documentation

system configuration settings and associated documentation

system audit records

system auditable events

other relevant documents or records

Organizational personnel with audit and accountability responsibilities

organizational personnel with information security and privacy responsibilities

system/network administrators

Mechanisms implementing system auditing

### {au-03 au %!s(int=3) %!s(int=0)}: Content of Audit Records

Ensure that audit records contain information that establishes the following:

What type of event occurred;

When the event occurred;

Where the event occurred;

Source of the event;

Outcome of the event; and

Identity of any individuals, subjects, or objects/entities associated with the event.

Audit record content that may be necessary to support the auditing function includes event descriptions (item a), time stamps (item b), source and destination addresses (item c), user or process identifiers (items d and f), success or fail indications (item e), and filenames involved (items a, c, e, and f) . Event outcomes include indicators of event success or failure and event-specific results, such as the system security and privacy posture after the event occurred. Organizations consider how audit records can reveal information about individuals that may give rise to privacy risks and how best to mitigate such risks. For example, there is the potential to reveal personally identifiable information in the audit trail, especially if the trail records inputs or is based on patterns or time of usage.

audit records contain information that establishes what type of event occurred;

audit records contain information that establishes when the event occurred;

audit records contain information that establishes where the event occurred;

audit records contain information that establishes the source of the event;

audit records contain information that establishes the outcome of the event;

audit records contain information that establishes the identity of any individuals, subjects, or objects/entities associated with the event.

Audit and accountability policy

system security plan

privacy plan

procedures addressing content of audit records

system design documentation

system configuration settings and associated documentation

list of organization-defined auditable events

system audit records

system incident reports

other relevant documents or records

Organizational personnel with audit and accountability responsibilities

organizational personnel with information security and privacy responsibilities

system/network administrators

Mechanisms implementing system auditing of auditable events

### {au-03.01 au %!s(int=3) %!s(int=1)}: Additional Audit Information

Generate audit records containing the following additional information: {{ insert: param, au-03.01_odp }}.

The ability to add information generated in audit records is dependent on system functionality to configure the audit record content. Organizations may consider additional information in audit records including, but not limited to, access control or flow control rules invoked and individual identities of group account users. Organizations may also consider limiting additional audit record information to only information that is explicitly needed for audit requirements. This facilitates the use of audit trails and audit logs by not including information in audit records that could potentially be misleading, make it more difficult to locate information of interest, or increase the risk to individuals' privacy.

generated audit records contain the following {{ insert: param, au-03.01_odp }}.

Audit and accountability policy

procedures addressing content of audit records

system security plan

privacy plan

system design documentation

system configuration settings and associated documentation

list of organization-defined auditable events

system audit records

other relevant documents or records

Organizational personnel with audit and accountability responsibilities

organizational personnel with information security and privacy responsibilities

system/network administrators

system developers

system audit capability

### {au-04 au %!s(int=4) %!s(int=0)}: Audit Log Storage Capacity

Allocate audit log storage capacity to accommodate {{ insert: param, au-04_odp }}.

Organizations consider the types of audit logging to be performed and the audit log processing requirements when allocating audit log storage capacity. Allocating sufficient audit log storage capacity reduces the likelihood of such capacity being exceeded and resulting in the potential loss or reduction of audit logging capability.

audit log storage capacity is allocated to accommodate {{ insert: param, au-04_odp }}.

Audit and accountability policy

procedures addressing audit storage capacity

system security plan

privacy plan

system design documentation

system configuration settings and associated documentation

audit record storage requirements

audit record storage capability for system components

system audit records

other relevant documents or records

Organizational personnel with audit and accountability responsibilities

organizational personnel with information security and privacy responsibilities

system/network administrators

system developers

Audit record storage capacity and related configuration settings

### {au-05 au %!s(int=5) %!s(int=0)}: Response to Audit Logging Process Failures

Alert {{ insert: param, au-05_odp.01 }} within {{ insert: param, au-05_odp.02 }} in the event of an audit logging process failure; and

Take the following additional actions: {{ insert: param, au-05_odp.03 }}.

Audit logging process failures include software and hardware errors, failures in audit log capturing mechanisms, and reaching or exceeding audit log storage capacity. Organization-defined actions include overwriting oldest audit records, shutting down the system, and stopping the generation of audit records. Organizations may choose to define additional actions for audit logging process failures based on the type of failure, the location of the failure, the severity of the failure, or a combination of such factors. When the audit logging process failure is related to storage, the response is carried out for the audit log storage repository (i.e., the distinct system component where the audit logs are stored), the system on which the audit logs reside, the total audit log storage capacity of the organization (i.e., all audit log storage repositories combined), or all three. Organizations may decide to take no additional actions after alerting designated roles or personnel.

 {{ insert: param, au-05_odp.01 }} are alerted in the event of an audit logging process failure within {{ insert: param, au-05_odp.02 }};

 {{ insert: param, au-05_odp.03 }} are taken in the event of an audit logging process failure.

Audit and accountability policy

procedures addressing response to audit processing failures

system design documentation

system security plan

privacy plan

system configuration settings and associated documentation

list of personnel to be notified in case of an audit processing failure

system audit records

other relevant documents or records

Organizational personnel with audit and accountability responsibilities

organizational personnel with information security and privacy responsibilities

system/network administrators

system developers

Mechanisms implementing system response to audit processing failures

### {au-05.01 au %!s(int=5) %!s(int=1)}: Storage Capacity Warning

Provide a warning to {{ insert: param, au-05.01_odp.01 }} within {{ insert: param, au-05.01_odp.02 }} when allocated audit log storage volume reaches {{ insert: param, au-05.01_odp.03 }} of repository maximum audit log storage capacity.

Organizations may have multiple audit log storage repositories distributed across multiple system components with each repository having different storage volume capacities.

a warning is provided to {{ insert: param, au-05.01_odp.01 }} within {{ insert: param, au-05.01_odp.02 }} when allocated audit log storage volume reaches {{ insert: param, au-05.01_odp.03 }} of repository maximum audit log storage capacity.

Audit and accountability policy

procedures addressing response to audit processing failures

system design documentation

system security plan

privacy system configuration settings and associated documentation

system audit records

other relevant documents or records

Organizational personnel with audit and accountability responsibilities

organizational personnel with information security and privacy responsibilities

system/network administrators

system developers

Mechanisms implementing audit storage limit warnings

### {au-05.02 au %!s(int=5) %!s(int=2)}: Real-time Alerts

Provide an alert within {{ insert: param, au-05.02_odp.01 }} to {{ insert: param, au-05.02_odp.02 }} when the following audit failure events occur: {{ insert: param, au-05.02_odp.03 }}.

Alerts provide organizations with urgent messages. Real-time alerts provide these messages at information technology speed (i.e., the time from event detection to alert occurs in seconds or less).

an alert is provided within {{ insert: param, au-05.02_odp.01 }} to {{ insert: param, au-05.02_odp.02 }} when {{ insert: param, au-05.02_odp.03 }} occur.

Audit and accountability policy

procedures addressing response to audit processing failures

system design documentation

system security plan

privacy plan

system configuration settings and associated documentation

system audit records

other relevant documents or records

Organizational personnel with audit and accountability responsibilities

organizational personnel with information security and privacy responsibilities

system/network administrators

system developers

### {au-06 au %!s(int=6) %!s(int=0)}: Audit Record Review, Analysis, and Reporting

Review and analyze system audit records {{ insert: param, au-06_odp.01 }} for indications of {{ insert: param, au-06_odp.02 }} and the potential impact of the inappropriate or unusual activity;

Report findings to {{ insert: param, au-06_odp.03 }} ; and

Adjust the level of audit record review, analysis, and reporting within the system when there is a change in risk based on law enforcement information, intelligence information, or other credible sources of information.

Audit record review, analysis, and reporting covers information security- and privacy-related logging performed by organizations, including logging that results from the monitoring of account usage, remote access, wireless connectivity, mobile device connection, configuration settings, system component inventory, use of maintenance tools and non-local maintenance, physical access, temperature and humidity, equipment delivery and removal, communications at system interfaces, and use of mobile code or Voice over Internet Protocol (VoIP). Findings can be reported to organizational entities that include the incident response team, help desk, and security or privacy offices. If organizations are prohibited from reviewing and analyzing audit records or unable to conduct such activities, the review or analysis may be carried out by other organizations granted such authority. The frequency, scope, and/or depth of the audit record review, analysis, and reporting may be adjusted to meet organizational needs based on new information received.

system audit records are reviewed and analyzed {{ insert: param, au-06_odp.01 }} for indications of {{ insert: param, au-06_odp.02 }} and the potential impact of the inappropriate or unusual activity;

findings are reported to {{ insert: param, au-06_odp.03 }};

the level of audit record review, analysis, and reporting within the system is adjusted when there is a change in risk based on law enforcement information, intelligence information, or other credible sources of information.

Audit and accountability policy

system security plan

privacy plan

procedures addressing audit review, analysis, and reporting

reports of audit findings

records of actions taken in response to reviews/analyses of audit records

other relevant documents or records

Organizational personnel with audit review, analysis, and reporting responsibilities

organizational personnel with information security and privacy responsibilities

### {au-06.01 au %!s(int=6) %!s(int=1)}: Automated Process Integration

Integrate audit record review, analysis, and reporting processes using {{ insert: param, au-06.01_odp }}.

Organizational processes that benefit from integrated audit record review, analysis, and reporting include incident response, continuous monitoring, contingency planning, investigation and response to suspicious activities, and Inspector General audits.

audit record review, analysis, and reporting processes are integrated using {{ insert: param, au-06.01_odp }}.

Audit and accountability policy

system security plan

privacy plan

procedures addressing audit review, analysis, and reporting

procedures addressing investigation and response to suspicious activities

system design documentation

system configuration settings and associated documentation

system audit records

other relevant documents or records

Organizational personnel with audit review, analysis, and reporting responsibilities

organizational personnel with information security and privacy responsibilities

Automated mechanisms integrating audit review, analysis, and reporting processes

### {au-06.03 au %!s(int=6) %!s(int=3)}: Correlate Audit Record Repositories

Analyze and correlate audit records across different repositories to gain organization-wide situational awareness.

Organization-wide situational awareness includes awareness across all three levels of risk management (i.e., organizational level, mission/business process level, and information system level) and supports cross-organization awareness.

audit records across different repositories are analyzed and correlated to gain organization-wide situational awareness.

Audit and accountability policy

system security plan

privacy plan

procedures addressing audit review, analysis, and reporting

system design documentation

system configuration settings and associated documentation

system audit records across different repositories

other relevant documents or records

Organizational personnel with audit review, analysis, and reporting responsibilities

organizational personnel with information security and privacy responsibilities

Mechanisms supporting the analysis and correlation of audit records

### {au-06.05 au %!s(int=6) %!s(int=5)}: Integrated Analysis of Audit Records

Integrate analysis of audit records with analysis of {{ insert: param, au-06.05_odp.01 }} to further enhance the ability to identify inappropriate or unusual activity.

Integrated analysis of audit records does not require vulnerability scanning, the generation of performance data, or system monitoring. Rather, integrated analysis requires that the analysis of information generated by scanning, monitoring, or other data collection activities is integrated with the analysis of audit record information. Security Information and Event Management tools can facilitate audit record aggregation or consolidation from multiple system components as well as audit record correlation and analysis. The use of standardized audit record analysis scripts developed by organizations (with localized script adjustments, as necessary) provides more cost-effective approaches for analyzing audit record information collected. The correlation of audit record information with vulnerability scanning information is important in determining the veracity of vulnerability scans of the system and in correlating attack detection events with scanning results. Correlation with performance data can uncover denial-of-service attacks or other types of attacks that result in the unauthorized use of resources. Correlation with system monitoring information can assist in uncovering attacks and in better relating audit information to operational situations.

analysis of audit records is integrated with analysis of {{ insert: param, au-06.05_odp.01 }} to further enhance the ability to identify inappropriate or unusual activity.

Audit and accountability policy

system security plan

privacy plan

procedures addressing audit review, analysis, and reporting

system design documentation

system configuration settings and associated documentation

integrated analysis of audit records, vulnerability scanning information, performance data, network monitoring information, and associated documentation

other relevant documents or records

Organizational personnel with audit review, analysis, and reporting responsibilities

organizational personnel with information security and privacy responsibilities

Mechanisms implementing the capability to integrate analysis of audit records with analysis of data/information sources

### {au-06.06 au %!s(int=6) %!s(int=6)}: Correlation with Physical Monitoring

Correlate information from audit records with information obtained from monitoring physical access to further enhance the ability to identify suspicious, inappropriate, unusual, or malevolent activity.

The correlation of physical audit record information and the audit records from systems may assist organizations in identifying suspicious behavior or supporting evidence of such behavior. For example, the correlation of an individual’s identity for logical access to certain systems with the additional physical security information that the individual was present at the facility when the logical access occurred may be useful in investigations.

information from audit records is correlated with information obtained from monitoring physical access to further enhance the ability to identify suspicious, inappropriate, unusual, or malevolent activity.

Audit and accountability policy

procedures addressing audit review, analysis, and reporting

procedures addressing physical access monitoring

system design documentation

system configuration settings and associated documentation

documentation providing evidence of correlated information obtained from audit records and physical access monitoring records

system security plan

privacy plan

other relevant documents or records

Organizational personnel with audit review, analysis, and reporting responsibilities

organizational personnel with physical access monitoring responsibilities

organizational personnel with information security and privacy responsibilities

Mechanisms implementing the capability to correlate information from audit records with information from monitoring physical access

### {au-07 au %!s(int=7) %!s(int=0)}: Audit Record Reduction and Report Generation

Provide and implement an audit record reduction and report generation capability that:

Supports on-demand audit record review, analysis, and reporting requirements and after-the-fact investigations of incidents; and

Does not alter the original content or time ordering of audit records.

Audit record reduction is a process that manipulates collected audit log information and organizes it into a summary format that is more meaningful to analysts. Audit record reduction and report generation capabilities do not always emanate from the same system or from the same organizational entities that conduct audit logging activities. The audit record reduction capability includes modern data mining techniques with advanced data filters to identify anomalous behavior in audit records. The report generation capability provided by the system can generate customizable reports. Time ordering of audit records can be an issue if the granularity of the timestamp in the record is insufficient.

an audit record reduction and report generation capability is provided that supports on-demand audit record review, analysis, and reporting requirements and after-the-fact investigations of incidents;

an audit record reduction and report generation capability is implemented that supports on-demand audit record review, analysis, and reporting requirements and after-the-fact investigations of incidents;

an audit record reduction and report generation capability is provided that does not alter the original content or time ordering of audit records;

an audit record reduction and report generation capability is implemented that does not alter the original content or time ordering of audit records.

Audit and accountability policy

system security plan

privacy plan

procedures addressing audit reduction and report generation

system design documentation

system configuration settings and associated documentation

audit reduction, review, analysis, and reporting tools

system audit records

other relevant documents or records

Organizational personnel with audit reduction and report generation responsibilities

organizational personnel with information security and privacy responsibilities

Audit reduction and report generation capability

### {au-07.01 au %!s(int=7) %!s(int=1)}: Automatic Processing

Provide and implement the capability to process, sort, and search audit records for events of interest based on the following content: {{ insert: param, au-07.01_odp }}.

Events of interest can be identified by the content of audit records, including system resources involved, information objects accessed, identities of individuals, event types, event locations, event dates and times, Internet Protocol addresses involved, or event success or failure. Organizations may define event criteria to any degree of granularity required, such as locations selectable by a general networking location or by specific system component.

the capability to process, sort, and search audit records for events of interest based on {{ insert: param, au-07.01_odp }} are provided;

the capability to process, sort, and search audit records for events of interest based on {{ insert: param, au-07.01_odp }} are implemented.

Audit and accountability policy

system security plan

privacy plan

procedures addressing audit reduction and report generation

system design documentation

system configuration settings and associated documentation

audit reduction, review, analysis, and reporting tools

audit record criteria (fields) establishing events of interest

system audit records

other relevant documents or records

Organizational personnel with audit reduction and report generation responsibilities

organizational personnel with information security and privacy responsibilities

system developers

Audit reduction and report generation capability

### {au-08 au %!s(int=8) %!s(int=0)}: Time Stamps

Use internal system clocks to generate time stamps for audit records; and

Record time stamps for audit records that meet {{ insert: param, au-08_odp }} and that use Coordinated Universal Time, have a fixed local time offset from Coordinated Universal Time, or that include the local time offset as part of the time stamp.

Time stamps generated by the system include date and time. Time is commonly expressed in Coordinated Universal Time (UTC), a modern continuation of Greenwich Mean Time (GMT), or local time with an offset from UTC. Granularity of time measurements refers to the degree of synchronization between system clocks and reference clocks (e.g., clocks synchronizing within hundreds of milliseconds or tens of milliseconds). Organizations may define different time granularities for different system components. Time service can be critical to other security capabilities such as access control and identification and authentication, depending on the nature of the mechanisms used to support those capabilities.

internal system clocks are used to generate timestamps for audit records;

timestamps are recorded for audit records that meet {{ insert: param, au-08_odp }} and that use Coordinated Universal Time, have a fixed local time offset from Coordinated Universal Time, or include the local time offset as part of the timestamp.

Audit and accountability policy

system security plan

privacy plan

procedures addressing timestamp generation

system design documentation

system configuration settings and associated documentation

system audit records

other relevant documents or records

Organizational personnel with information security and privacy responsibilities

system/network administrators

system developers

Mechanisms implementing timestamp generation

### {au-09 au %!s(int=9) %!s(int=0)}: Protection of Audit Information

Protect audit information and audit logging tools from unauthorized access, modification, and deletion; and

Alert {{ insert: param, au-09_odp }} upon detection of unauthorized access, modification, or deletion of audit information.

Audit information includes all information needed to successfully audit system activity, such as audit records, audit log settings, audit reports, and personally identifiable information. Audit logging tools are those programs and devices used to conduct system audit and logging activities. Protection of audit information focuses on technical protection and limits the ability to access and execute audit logging tools to authorized individuals. Physical protection of audit information is addressed by both media protection controls and physical and environmental protection controls.

audit information and audit logging tools are protected from unauthorized access, modification, and deletion;

 {{ insert: param, au-09_odp }} are alerted upon detection of unauthorized access, modification, or deletion of audit information.

Audit and accountability policy

system security plan

privacy plan

access control policy and procedures

procedures addressing protection of audit information

system design documentation

system configuration settings and associated documentation

system audit records

audit tools

other relevant documents or records

Organizational personnel with audit and accountability responsibilities

organizational personnel with information security and privacy responsibilities

system/network administrators

system developers

Mechanisms implementing audit information protection

### {au-09.02 au %!s(int=9) %!s(int=2)}: Store on Separate Physical Systems or Components

Store audit records {{ insert: param, au-09.02_odp }} in a repository that is part of a physically different system or system component than the system or component being audited.

Storing audit records in a repository separate from the audited system or system component helps to ensure that a compromise of the system being audited does not also result in a compromise of the audit records. Storing audit records on separate physical systems or components also preserves the confidentiality and integrity of audit records and facilitates the management of audit records as an organization-wide activity. Storing audit records on separate systems or components applies to initial generation as well as backup or long-term storage of audit records.

audit records are stored {{ insert: param, au-09.02_odp }} in a repository that is part of a physically different system or system component than the system or component being audited.

Audit and accountability policy

system security plan

privacy plan

procedures addressing protection of audit information

system design documentation

system configuration settings and associated documentation

system or media storing backups of system audit records

system audit records

other relevant documents or records

Organizational personnel with audit and accountability responsibilities

organizational personnel with information security and privacy responsibilities

system/network administrators

system developers

Mechanisms implementing the backing up of audit records

### {au-09.03 au %!s(int=9) %!s(int=3)}: Cryptographic Protection

Implement cryptographic mechanisms to protect the integrity of audit information and audit tools.

Cryptographic mechanisms used for protecting the integrity of audit information include signed hash functions using asymmetric cryptography. This enables the distribution of the public key to verify the hash information while maintaining the confidentiality of the secret key used to generate the hash.

cryptographic mechanisms to protect the integrity of audit information and audit tools are implemented.

Audit and accountability policy

system security plan

privacy plan

access control policy and procedures

procedures addressing protection of audit information

system design documentation

system hardware settings

system configuration settings and associated documentation

system audit records

other relevant documents or records

Organizational personnel with audit and accountability responsibilities

organizational personnel with information security and privacy responsibilities

system/network administrators

system developers

Cryptographic mechanisms protecting the integrity of audit information and tools

### {au-09.04 au %!s(int=9) %!s(int=4)}: Access by Subset of Privileged Users

Authorize access to management of audit logging functionality to only {{ insert: param, au-09.04_odp }}.

Individuals or roles with privileged access to a system and who are also the subject of an audit by that system may affect the reliability of the audit information by inhibiting audit activities or modifying audit records. Requiring privileged access to be further defined between audit-related privileges and other privileges limits the number of users or roles with audit-related privileges.

access to management of audit logging functionality is authorized only to {{ insert: param, au-09.04_odp }}.

Audit and accountability policy

system security plan

privacy plan

access control policy and procedures

procedures addressing protection of audit information

system design documentation

system configuration settings and associated documentation

system-generated list of privileged users with access to management of audit functionality

access authorizations

access control list

system audit records

other relevant documents or records

Organizational personnel with audit and accountability responsibilities

organizational personnel with information security and privacy responsibilities

system/network administrators

Mechanisms managing access to audit functionality

### {au-10 au %!s(int=10) %!s(int=0)}: Non-repudiation

Provide irrefutable evidence that an individual (or process acting on behalf of an individual) has performed {{ insert: param, au-10_odp }}.

Types of individual actions covered by non-repudiation include creating information, sending and receiving messages, and approving information. Non-repudiation protects against claims by authors of not having authored certain documents, senders of not having transmitted messages, receivers of not having received messages, and signatories of not having signed documents. Non-repudiation services can be used to determine if information originated from an individual or if an individual took specific actions (e.g., sending an email, signing a contract, approving a procurement request, or receiving specific information). Organizations obtain non-repudiation services by employing various techniques or mechanisms, including digital signatures and digital message receipts.

irrefutable evidence is provided that an individual (or process acting on behalf of an individual) has performed {{ insert: param, au-10_odp }}.

Audit and accountability policy

system security plan

privacy plan

procedures addressing non-repudiation

system design documentation

system configuration settings and associated documentation

system audit records

other relevant documents or records

Organizational personnel with information security and privacy responsibilities

system/network administrators

system developers

Mechanisms implementing non-repudiation capability

### {au-11 au %!s(int=11) %!s(int=0)}: Audit Record Retention

Retain audit records for {{ insert: param, au-11_odp }} to provide support for after-the-fact investigations of incidents and to meet regulatory and organizational information retention requirements.

Organizations retain audit records until it is determined that the records are no longer needed for administrative, legal, audit, or other operational purposes. This includes the retention and availability of audit records relative to Freedom of Information Act (FOIA) requests, subpoenas, and law enforcement actions. Organizations develop standard categories of audit records relative to such types of actions and standard response processes for each type of action. The National Archives and Records Administration (NARA) General Records Schedules provide federal policy on records retention.

audit records are retained for {{ insert: param, au-11_odp }} to provide support for after-the-fact investigations of incidents and to meet regulatory and organizational information retention requirements.

Audit and accountability policy

system security plan

privacy plan

audit record retention policy and procedures

security plan

organization-defined retention period for audit records

audit record archives

audit logs

audit records

other relevant documents or records

Organizational personnel with audit record retention responsibilities

organizational personnel with information security and privacy responsibilities

system/network administrators

### {au-12 au %!s(int=12) %!s(int=0)}: Audit Record Generation

Provide audit record generation capability for the event types the system is capable of auditing as defined in [AU-2a](#au-2_smt.a) on {{ insert: param, au-12_odp.01 }};

Allow {{ insert: param, au-12_odp.02 }} to select the event types that are to be logged by specific components of the system; and

Generate audit records for the event types defined in [AU-2c](#au-2_smt.c) that include the audit record content defined in [AU-3](#au-3).

Audit records can be generated from many different system components. The event types specified in [AU-2d](#au-2_smt.d) are the event types for which audit logs are to be generated and are a subset of all event types for which the system can generate audit records.

audit record generation capability for the event types the system is capable of auditing (defined in AU-02_ODP[01]) is provided by {{ insert: param, au-12_odp.01 }};

 {{ insert: param, au-12_odp.02 }} is/are allowed to select the event types that are to be logged by specific components of the system;

audit records for the event types defined in AU-02_ODP[02] that include the audit record content defined in AU-03 are generated.

Audit and accountability policy

procedures addressing audit record generation

system security plan

privacy plan

system design documentation

system configuration settings and associated documentation

list of auditable events

system audit records

other relevant documents or records

Organizational personnel with audit record generation responsibilities

organizational personnel with information security and privacy responsibilities

system/network administrators

system developers

Mechanisms implementing audit record generation capability

### {au-12.01 au %!s(int=12) %!s(int=1)}: System-wide and Time-correlated Audit Trail

Compile audit records from {{ insert: param, au-12.01_odp.01 }} into a system-wide (logical or physical) audit trail that is time-correlated to within {{ insert: param, au-12.01_odp.02 }}.

Audit trails are time-correlated if the time stamps in the individual audit records can be reliably related to the time stamps in other audit records to achieve a time ordering of the records within organizational tolerances.

audit records from {{ insert: param, au-12.01_odp.01 }} are compiled into a system-wide (logical or physical) audit trail that is time-correlated to within {{ insert: param, au-12.01_odp.02 }}.

Audit and accountability policy

system security plan

privacy plan

procedures addressing audit record generation

system design documentation

system configuration settings and associated documentation

system-wide audit trail (logical or physical)

system audit records

other relevant documents or records

Organizational personnel with audit record generation responsibilities

organizational personnel with information security and privacy responsibilities

system/network administrators

system developers

Mechanisms implementing audit record generation capability

### {au-12.03 au %!s(int=12) %!s(int=3)}: Changes by Authorized Individuals

Provide and implement the capability for {{ insert: param, au-12.03_odp.01 }} to change the logging to be performed on {{ insert: param, au-12.03_odp.02 }} based on {{ insert: param, au-12.03_odp.03 }} within {{ insert: param, au-12.03_odp.04 }}.

Permitting authorized individuals to make changes to system logging enables organizations to extend or limit logging as necessary to meet organizational requirements. Logging that is limited to conserve system resources may be extended (either temporarily or permanently) to address certain threat situations. In addition, logging may be limited to a specific set of event types to facilitate audit reduction, analysis, and reporting. Organizations can establish time thresholds in which logging actions are changed (e.g., near real-time, within minutes, or within hours).

the capability for {{ insert: param, au-12.03_odp.01 }} to change the logging to be performed on {{ insert: param, au-12.03_odp.02 }} based on {{ insert: param, au-12.03_odp.03 }} within {{ insert: param, au-12.03_odp.04 }} is provided;

the capability for {{ insert: param, au-12.03_odp.01 }} to change the logging to be performed on {{ insert: param, au-12.03_odp.02 }} based on {{ insert: param, au-12.03_odp.03 }} within {{ insert: param, au-12.03_odp.04 }} is implemented.

Audit and accountability policy

system security plan

privacy plan

procedures addressing audit record generation

system design documentation

system configuration settings and associated documentation

system-generated list of individuals or roles authorized to change auditing to be performed

system audit records

other relevant documents or records

Organizational personnel with audit record generation responsibilities

organizational personnel with information security and privacy responsibilities

system/network administrators

system developers

Mechanisms implementing audit record generation capability

