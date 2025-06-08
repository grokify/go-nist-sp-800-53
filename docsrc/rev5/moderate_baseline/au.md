# AU - Audit and Accountability

* Controls Count: 16
* Controls IDs: AU-1, AU-2, AU-3, AU-3 (1), AU-4, AU-5, AU-6, AU-6 (1), AU-6 (3), AU-7, AU-7 (1), AU-8, AU-9, AU-9 (4), AU-11, AU-12

## Controls

### AU-1: Policy and Procedures

Develop, document, and disseminate to organization-defined personnel or roles:

organization-level, mission/business process-level, and/or system-level audit and accountability policy that:

Addresses purpose, scope, roles, responsibilities, management commitment, coordination among organizational entities, and compliance; and

Is consistent with applicable laws, executive orders, directives, regulations, policies, standards, and guidelines; and

Procedures to facilitate the implementation of the audit and accountability policy and the associated audit and accountability controls;

Designate an an official to manage the audit and accountability policy and procedures is defined; to manage the development, documentation, and dissemination of the audit and accountability policy and procedures; and

Review and update the current audit and accountability:

Policy the frequency at which the current audit and accountability policy is reviewed and updated is defined; and following events that would require the current audit and accountability policy to be reviewed and updated are defined; ; and

Procedures the frequency at which the current audit and accountability procedures are reviewed and updated is defined; and following events that would require audit and accountability procedures to be reviewed and updated are defined;.

Audit and accountability policy and procedures address the controls in the AU family that are implemented within systems and organizations. The risk management strategy is an important factor in establishing such policies and procedures. Policies and procedures contribute to security and privacy assurance. Therefore, it is important that security and privacy programs collaborate on the development of audit and accountability policy and procedures. Security and privacy program policies and procedures at the organization level are preferable, in general, and may obviate the need for mission- or system-specific policies and procedures. The policy can be included as part of the general security and privacy policy or be represented by multiple policies that reflect the complex nature of organizations. Procedures can be established for security and privacy programs, for mission or business processes, and for systems, if needed. Procedures describe how the policies or controls are implemented and can be directed at the individual or role that is the object of the procedure. Procedures can be documented in system security and privacy plans or in one or more separate documents. Events that may precipitate an update to audit and accountability policy and procedures include assessment or audit findings, security incidents or breaches, or changes in applicable laws, executive orders, directives, regulations, policies, standards, and guidelines. Simply restating controls does not constitute an organizational policy or procedure.

an audit and accountability policy is developed and documented;

the audit and accountability policy is disseminated to personnel or roles to whom the audit and accountability policy is to be disseminated is/are defined;;

audit and accountability procedures to facilitate the implementation of the audit and accountability policy and associated audit and accountability controls are developed and documented;

the audit and accountability procedures are disseminated to personnel or roles to whom the audit and accountability procedures are to be disseminated is/are defined;;

the organization-level, mission/business process-level, and/or system-level of the audit and accountability policy addresses purpose;

the organization-level, mission/business process-level, and/or system-level of the audit and accountability policy addresses scope;

the organization-level, mission/business process-level, and/or system-level of the audit and accountability policy addresses roles;

the organization-level, mission/business process-level, and/or system-level of the audit and accountability policy addresses responsibilities;

the organization-level, mission/business process-level, and/or system-level of the audit and accountability policy addresses management commitment;

the organization-level, mission/business process-level, and/or system-level of the audit and accountability policy addresses coordination among organizational entities;

the organization-level, mission/business process-level, and/or system-level of the audit and accountability policy addresses compliance;

the organization-level, mission/business process-level, and/or system-level of the audit and accountability policy is consistent with applicable laws, executive orders, directives, regulations, policies, standards, and guidelines;

the an official to manage the audit and accountability policy and procedures is defined; is designated to manage the development, documentation, and dissemination of the audit and accountability policy and procedures;

the current audit and accountability policy is reviewed and updated the frequency at which the current audit and accountability policy is reviewed and updated is defined;;

the current audit and accountability policy is reviewed and updated following events that would require the current audit and accountability policy to be reviewed and updated are defined;;

the current audit and accountability procedures are reviewed and updated the frequency at which the current audit and accountability procedures are reviewed and updated is defined;;

the current audit and accountability procedures are reviewed and updated following events that would require audit and accountability procedures to be reviewed and updated are defined;.

Audit and accountability policy and procedures

system security plan

privacy plan

other relevant documents or records

Organizational personnel with audit and accountability responsibilities

organizational personnel with information security and privacy responsibilities

### AU-2: Event Logging

Identify the types of events that the system is capable of logging in support of the audit function: the event types that the system is capable of logging in support of the audit function are defined;;

Coordinate the event logging function with other organizational entities requiring audit-related information to guide and inform the selection criteria for events to be logged;

Specify the following event types for logging within the system: organization-defined event types (subset of the event types defined in [AU-2a.](#au-2_smt.a)) along with the frequency of (or situation requiring) logging for each identified event type;

Provide a rationale for why the event types selected for logging are deemed to be adequate to support after-the-fact investigations of incidents; and

Review and update the event types selected for logging the frequency of event types selected for logging are reviewed and updated;.

An event is an observable occurrence in a system. The types of events that require logging are those events that are significant and relevant to the security of systems and the privacy of individuals. Event logging also supports specific monitoring and auditing needs. Event types include password changes, failed logons or failed accesses related to systems, security or privacy attribute changes, administrative privilege usage, PIV credential usage, data action changes, query parameters, or external credential usage. In determining the set of event types that require logging, organizations consider the monitoring and auditing appropriate for each of the controls to be implemented. For completeness, event logging includes all protocols that are operational and supported by the system.

To balance monitoring and auditing requirements with other system needs, event logging requires identifying the subset of event types that are logged at a given point in time. For example, organizations may determine that systems need the capability to log every file access successful and unsuccessful, but not activate that capability except for specific circumstances due to the potential burden on system performance. The types of events that organizations desire to be logged may change. Reviewing and updating the set of logged events is necessary to help ensure that the events remain relevant and continue to support the needs of the organization. Organizations consider how the types of logging events can reveal information about individuals that may give rise to privacy risk and how best to mitigate such risks. For example, there is the potential to reveal personally identifiable information in the audit trail, especially if the logging event is based on patterns or time of usage.

Event logging requirements, including the need to log specific event types, may be referenced in other controls and control enhancements. These include [AC-2(4)](#ac-2.4), [AC-3(10)](#ac-3.10), [AC-6(9)](#ac-6.9), [AC-17(1)](#ac-17.1), [CM-3f](#cm-3_smt.f), [CM-5(1)](#cm-5.1), [IA-3(3)(b)](#ia-3.3_smt.b), [MA-4(1)](#ma-4.1), [MP-4(2)](#mp-4.2), [PE-3](#pe-3), [PM-21](#pm-21), [PT-7](#pt-7), [RA-8](#ra-8), [SC-7(9)](#sc-7.9), [SC-7(15)](#sc-7.15), [SI-3(8)](#si-3.8), [SI-4(22)](#si-4.22), [SI-7(8)](#si-7.8) , and [SI-10(1)](#si-10.1) . Organizations include event types that are required by applicable laws, executive orders, directives, policies, regulations, standards, and guidelines. Audit records can be generated at various levels, including at the packet level as information traverses the network. Selecting the appropriate level of event logging is an important part of a monitoring and auditing capability and can identify the root causes of problems. When defining event types, organizations consider the logging necessary to cover related event types, such as the steps in distributed, transaction-based processes and the actions that occur in service-oriented architectures.

the event types that the system is capable of logging in support of the audit function are defined; that the system is capable of logging are identified in support of the audit logging function;

the event logging function is coordinated with other organizational entities requiring audit-related information to guide and inform the selection criteria for events to be logged;

the event types (subset of AU-02_ODP[01]) for logging within the system are defined; are specified for logging within the system;

the specified event types are logged within the system the frequency or situation requiring logging for each specified event type is defined;;

a rationale is provided for why the event types selected for logging are deemed to be adequate to support after-the-fact investigations of incidents;

the event types selected for logging are reviewed and updated the frequency of event types selected for logging are reviewed and updated;.

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

### AU-3: Content of Audit Records

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

### AU-3 (1): Additional Audit Information

Generate audit records containing the following additional information: additional information to be included in audit records is defined;.

The ability to add information generated in audit records is dependent on system functionality to configure the audit record content. Organizations may consider additional information in audit records including, but not limited to, access control or flow control rules invoked and individual identities of group account users. Organizations may also consider limiting additional audit record information to only information that is explicitly needed for audit requirements. This facilitates the use of audit trails and audit logs by not including information in audit records that could potentially be misleading, make it more difficult to locate information of interest, or increase the risk to individuals' privacy.

generated audit records contain the following additional information to be included in audit records is defined;.

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

### AU-4: Audit Log Storage Capacity

Allocate audit log storage capacity to accommodate audit log retention requirements are defined;.

Organizations consider the types of audit logging to be performed and the audit log processing requirements when allocating audit log storage capacity. Allocating sufficient audit log storage capacity reduces the likelihood of such capacity being exceeded and resulting in the potential loss or reduction of audit logging capability.

audit log storage capacity is allocated to accommodate audit log retention requirements are defined;.

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

### AU-5: Response to Audit Logging Process Failures

Alert personnel or roles receiving audit logging process failure alerts are defined; within time period for personnel or roles receiving audit logging process failure alerts is defined; in the event of an audit logging process failure; and

Take the following additional actions: additional actions to be taken in the event of an audit logging process failure are defined;.

Audit logging process failures include software and hardware errors, failures in audit log capturing mechanisms, and reaching or exceeding audit log storage capacity. Organization-defined actions include overwriting oldest audit records, shutting down the system, and stopping the generation of audit records. Organizations may choose to define additional actions for audit logging process failures based on the type of failure, the location of the failure, the severity of the failure, or a combination of such factors. When the audit logging process failure is related to storage, the response is carried out for the audit log storage repository (i.e., the distinct system component where the audit logs are stored), the system on which the audit logs reside, the total audit log storage capacity of the organization (i.e., all audit log storage repositories combined), or all three. Organizations may decide to take no additional actions after alerting designated roles or personnel.

personnel or roles receiving audit logging process failure alerts are defined; are alerted in the event of an audit logging process failure within time period for personnel or roles receiving audit logging process failure alerts is defined;;

 additional actions to be taken in the event of an audit logging process failure are defined; are taken in the event of an audit logging process failure.

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

### AU-6: Audit Record Review, Analysis, and Reporting

Review and analyze system audit records frequency at which system audit records are reviewed and analyzed is defined; for indications of inappropriate or unusual activity is defined; and the potential impact of the inappropriate or unusual activity;

Report findings to personnel or roles to receive findings from reviews and analyses of system records is/are defined; ; and

Adjust the level of audit record review, analysis, and reporting within the system when there is a change in risk based on law enforcement information, intelligence information, or other credible sources of information.

Audit record review, analysis, and reporting covers information security- and privacy-related logging performed by organizations, including logging that results from the monitoring of account usage, remote access, wireless connectivity, mobile device connection, configuration settings, system component inventory, use of maintenance tools and non-local maintenance, physical access, temperature and humidity, equipment delivery and removal, communications at system interfaces, and use of mobile code or Voice over Internet Protocol (VoIP). Findings can be reported to organizational entities that include the incident response team, help desk, and security or privacy offices. If organizations are prohibited from reviewing and analyzing audit records or unable to conduct such activities, the review or analysis may be carried out by other organizations granted such authority. The frequency, scope, and/or depth of the audit record review, analysis, and reporting may be adjusted to meet organizational needs based on new information received.

system audit records are reviewed and analyzed frequency at which system audit records are reviewed and analyzed is defined; for indications of inappropriate or unusual activity is defined; and the potential impact of the inappropriate or unusual activity;

findings are reported to personnel or roles to receive findings from reviews and analyses of system records is/are defined;;

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

### AU-6 (1): Automated Process Integration

Integrate audit record review, analysis, and reporting processes using automated mechanisms used for integrating audit record review, analysis, and reporting processes are defined;.

Organizational processes that benefit from integrated audit record review, analysis, and reporting include incident response, continuous monitoring, contingency planning, investigation and response to suspicious activities, and Inspector General audits.

audit record review, analysis, and reporting processes are integrated using automated mechanisms used for integrating audit record review, analysis, and reporting processes are defined;.

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

### AU-6 (3): Correlate Audit Record Repositories

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

### AU-7: Audit Record Reduction and Report Generation

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

### AU-7 (1): Automatic Processing

Provide and implement the capability to process, sort, and search audit records for events of interest based on the following content: fields within audit records that can be processed, sorted, or searched are defined;.

Events of interest can be identified by the content of audit records, including system resources involved, information objects accessed, identities of individuals, event types, event locations, event dates and times, Internet Protocol addresses involved, or event success or failure. Organizations may define event criteria to any degree of granularity required, such as locations selectable by a general networking location or by specific system component.

the capability to process, sort, and search audit records for events of interest based on fields within audit records that can be processed, sorted, or searched are defined; are provided;

the capability to process, sort, and search audit records for events of interest based on fields within audit records that can be processed, sorted, or searched are defined; are implemented.

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

### AU-8: Time Stamps

Use internal system clocks to generate time stamps for audit records; and

Record time stamps for audit records that meet granularity of time measurement for audit record timestamps is defined; and that use Coordinated Universal Time, have a fixed local time offset from Coordinated Universal Time, or that include the local time offset as part of the time stamp.

Time stamps generated by the system include date and time. Time is commonly expressed in Coordinated Universal Time (UTC), a modern continuation of Greenwich Mean Time (GMT), or local time with an offset from UTC. Granularity of time measurements refers to the degree of synchronization between system clocks and reference clocks (e.g., clocks synchronizing within hundreds of milliseconds or tens of milliseconds). Organizations may define different time granularities for different system components. Time service can be critical to other security capabilities such as access control and identification and authentication, depending on the nature of the mechanisms used to support those capabilities.

internal system clocks are used to generate timestamps for audit records;

timestamps are recorded for audit records that meet granularity of time measurement for audit record timestamps is defined; and that use Coordinated Universal Time, have a fixed local time offset from Coordinated Universal Time, or include the local time offset as part of the timestamp.

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

### AU-9: Protection of Audit Information

Protect audit information and audit logging tools from unauthorized access, modification, and deletion; and

Alert personnel or roles to be alerted upon detection of unauthorized access, modification, or deletion of audit information is/are defined; upon detection of unauthorized access, modification, or deletion of audit information.

Audit information includes all information needed to successfully audit system activity, such as audit records, audit log settings, audit reports, and personally identifiable information. Audit logging tools are those programs and devices used to conduct system audit and logging activities. Protection of audit information focuses on technical protection and limits the ability to access and execute audit logging tools to authorized individuals. Physical protection of audit information is addressed by both media protection controls and physical and environmental protection controls.

audit information and audit logging tools are protected from unauthorized access, modification, and deletion;

 personnel or roles to be alerted upon detection of unauthorized access, modification, or deletion of audit information is/are defined; are alerted upon detection of unauthorized access, modification, or deletion of audit information.

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

### AU-9 (4): Access by Subset of Privileged Users

Authorize access to management of audit logging functionality to only a subset of privileged users or roles authorized to access management of audit logging functionality is defined;.

Individuals or roles with privileged access to a system and who are also the subject of an audit by that system may affect the reliability of the audit information by inhibiting audit activities or modifying audit records. Requiring privileged access to be further defined between audit-related privileges and other privileges limits the number of users or roles with audit-related privileges.

access to management of audit logging functionality is authorized only to a subset of privileged users or roles authorized to access management of audit logging functionality is defined;.

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

### AU-11: Audit Record Retention

Retain audit records for a time period to retain audit records that is consistent with the records retention policy is defined; to provide support for after-the-fact investigations of incidents and to meet regulatory and organizational information retention requirements.

Organizations retain audit records until it is determined that the records are no longer needed for administrative, legal, audit, or other operational purposes. This includes the retention and availability of audit records relative to Freedom of Information Act (FOIA) requests, subpoenas, and law enforcement actions. Organizations develop standard categories of audit records relative to such types of actions and standard response processes for each type of action. The National Archives and Records Administration (NARA) General Records Schedules provide federal policy on records retention.

audit records are retained for a time period to retain audit records that is consistent with the records retention policy is defined; to provide support for after-the-fact investigations of incidents and to meet regulatory and organizational information retention requirements.

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

### AU-12: Audit Record Generation

Provide audit record generation capability for the event types the system is capable of auditing as defined in [AU-2a](#au-2_smt.a) on system components that provide an audit record generation capability for the events types (defined in AU-02_ODP[02]) are defined;;

Allow personnel or roles allowed to select the event types that are to be logged by specific components of the system is/are defined; to select the event types that are to be logged by specific components of the system; and

Generate audit records for the event types defined in [AU-2c](#au-2_smt.c) that include the audit record content defined in [AU-3](#au-3).

Audit records can be generated from many different system components. The event types specified in [AU-2d](#au-2_smt.d) are the event types for which audit logs are to be generated and are a subset of all event types for which the system can generate audit records.

audit record generation capability for the event types the system is capable of auditing (defined in AU-02_ODP[01]) is provided by system components that provide an audit record generation capability for the events types (defined in AU-02_ODP[02]) are defined;;

 personnel or roles allowed to select the event types that are to be logged by specific components of the system is/are defined; is/are allowed to select the event types that are to be logged by specific components of the system;

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

