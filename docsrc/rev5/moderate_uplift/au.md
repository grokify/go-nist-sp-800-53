# AU - Audit and Accountability

* Controls: 6

## Controls

### AU-3 (1): Additional Audit Information

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

### AU-6 (1): Automated Process Integration

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

### AU-9 (4): Access by Subset of Privileged Users

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

