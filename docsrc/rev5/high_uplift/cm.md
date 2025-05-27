# CM - Configuration Management

* Controls: 8

## Controls

### CM-3 (1): Automated Documentation, Notification, and Prohibition of Changes

Use {{ insert: param, cm-03.01_odp.01 }} to:

Document proposed changes to the system;

Notify {{ insert: param, cm-03.01_odp.02 }} of proposed changes to the system and request change approval;

Highlight proposed changes to the system that have not been approved or disapproved within {{ insert: param, cm-03.01_odp.03 }};

Prohibit changes to the system until designated approvals are received;

Document all changes to the system; and

Notify {{ insert: param, cm-03.01_odp.04 }} when approved changes to the system are completed.

None.

 {{ insert: param, cm-03.01_odp.01 }} are used to document proposed changes to the system;

 {{ insert: param, cm-03.01_odp.01 }} are used to notify {{ insert: param, cm-03.01_odp.02 }} of proposed changes to the system and request change approval;

 {{ insert: param, cm-03.01_odp.01 }} are used to highlight proposed changes to the system that have not been approved or disapproved within {{ insert: param, cm-03.01_odp.03 }};

 {{ insert: param, cm-03.01_odp.01 }} are used to prohibit changes to the system until designated approvals are received;

 {{ insert: param, cm-03.01_odp.01 }} are used to document all changes to the system;

 {{ insert: param, cm-03.01_odp.01 }} are used to notify {{ insert: param, cm-03.01_odp.04 }} when approved changes to the system are completed.

Configuration management policy

procedures addressing system configuration change control

configuration management plan

system design documentation

system architecture and configuration documentation

automated configuration control mechanisms

system configuration settings and associated documentation

change control records

system audit records

change approval requests

change approvals

system security plan

other relevant documents or records

Organizational personnel with configuration change control responsibilities

organizational personnel with information security responsibilities

system/network administrators

system developers

members of change control board or similar

Organizational processes for configuration change control

automated mechanisms implementing configuration change control activities

### CM-3 (6): Cryptography Management

Ensure that cryptographic mechanisms used to provide the following controls are under configuration management: {{ insert: param, cm-03.06_odp }}.

The controls referenced in the control enhancement refer to security and privacy controls from the control catalog. Regardless of the cryptographic mechanisms employed, processes and procedures are in place to manage those mechanisms. For example, if system components use certificates for identification and authentication, a process is implemented to address the expiration of those certificates.

cryptographic mechanisms used to provide {{ insert: param, cm-03.06_odp }} are under configuration management.

Configuration management policy

procedures addressing system configuration change control

configuration management plan

system design documentation

system architecture and configuration documentation

system configuration settings and associated documentation

system security plan

other relevant documents or records

Organizational personnel with configuration change control responsibilities

organizational personnel with information security responsibilities

system/network administrators

system developers

members of change control board or similar

Organizational processes for configuration change control

cryptographic mechanisms implementing organizational security safeguards (controls)

### CM-4 (1): Separate Test Environments

Analyze changes to the system in a separate test environment before implementation in an operational environment, looking for security and privacy impacts due to flaws, weaknesses, incompatibility, or intentional malice.

A separate test environment requires an environment that is physically or logically separate and distinct from the operational environment. The separation is sufficient to ensure that activities in the test environment do not impact activities in the operational environment and that information in the operational environment is not inadvertently transmitted to the test environment. Separate environments can be achieved by physical or logical means. If physically separate test environments are not implemented, organizations determine the strength of mechanism required when implementing logical separation.

changes to the system are analyzed in a separate test environment before implementation in an operational environment;

changes to the system are analyzed for security impacts due to flaws;

changes to the system are analyzed for privacy impacts due to flaws;

changes to the system are analyzed for security impacts due to weaknesses;

changes to the system are analyzed for privacy impacts due to weaknesses;

changes to the system are analyzed for security impacts due to incompatibility;

changes to the system are analyzed for privacy impacts due to incompatibility;

changes to the system are analyzed for security impacts due to intentional malice;

changes to the system are analyzed for privacy impacts due to intentional malice.

Configuration management policy

procedures addressing security impact analyses for changes to the system

procedures addressing privacy impact analyses for changes to the system

configuration management plan

security impact analysis documentation

privacy impact analysis documentation

privacy impact assessment

privacy risk assessment documentation

analysis tools and associated outputs system design documentation

system architecture and configuration documentation

change control records

procedures addressing the authority to test with PII

system audit records

documentation of separate test and operational environments

system security plan

privacy plan

other relevant documents or records

Organizational personnel with responsibility for conducting security and privacy impact analyses

organizational personnel with information security and privacy responsibilities

system/network administrators

members of change control board or similar

Organizational processes for security and privacy impact analyses

mechanisms supporting and/or implementing security and privacy impact analyses of changes

### CM-5 (1): Automated Access Enforcement and Audit Records

Enforce access restrictions using {{ insert: param, cm-05.01_odp }} ; and

Automatically generate audit records of the enforcement actions.

Organizations log system accesses associated with applying configuration changes to ensure that configuration change control is implemented and to support after-the-fact actions should organizations discover any unauthorized changes.

access restrictions for change are enforced using {{ insert: param, cm-05.01_odp }};

audit records of enforcement actions are automatically generated.

Configuration management policy

procedures addressing access restrictions for changes to the system

system design documentation

system architecture and configuration documentation

system configuration settings and associated documentation

change control records

system audit records

system security plan

other relevant documents or records

Organizational personnel with logical access control responsibilities

organizational personnel with physical access control responsibilities

organizational personnel with information security responsibilities

system/network administrators

Organizational processes for managing access restrictions to change

automated mechanisms implementing the enforcement of access restrictions for changes to the system

automated mechanisms supporting auditing of enforcement actions

### CM-6 (1): Automated Management, Application, and Verification

Manage, apply, and verify configuration settings for {{ insert: param, cm-06.01_odp.01 }} using {{ insert: param, cm-6.1_prm_2 }}.

Automated tools (e.g., hardening tools, baseline configuration tools) can improve the accuracy, consistency, and availability of configuration settings information. Automation can also provide data aggregation and data correlation capabilities, alerting mechanisms, and dashboards to support risk-based decision-making within the organization.

configuration settings for {{ insert: param, cm-06.01_odp.01 }} are managed using {{ insert: param, cm-06.01_odp.02 }};

configuration settings for {{ insert: param, cm-06.01_odp.01 }} are applied using {{ insert: param, cm-06.01_odp.03 }};

configuration settings for {{ insert: param, cm-06.01_odp.01 }} are verified using {{ insert: param, cm-06.01_odp.04 }}.

Configuration management policy

procedures addressing configuration settings for the system

configuration management plan

system design documentation

system configuration settings and associated documentation

system component inventory

common secure configuration checklists

change control records

system audit records

system security plan

privacy plan

other relevant documents or records

Organizational personnel with security configuration management responsibilities

organizational personnel with information security and privacy responsibilities

system/network administrators

system developers

Organizational processes for managing configuration settings

automated mechanisms implemented to manage, apply, and verify system configuration settings

### CM-6 (2): Respond to Unauthorized Changes

Take the following actions in response to unauthorized changes to {{ insert: param, cm-06.02_odp.02 }}: {{ insert: param, cm-06.02_odp.01 }}.

Responses to unauthorized changes to configuration settings include alerting designated organizational personnel, restoring established configuration settings, or—in extreme cases—halting affected system processing.

 {{ insert: param, cm-06.02_odp.01 }} are taken in response to unauthorized changes to {{ insert: param, cm-06.02_odp.02 }}.

System security plan

privacy plan

configuration management policy

procedures addressing configuration settings for the system

configuration management plan

system design documentation

system configuration settings and associated documentation

alerts/notifications of unauthorized changes to system configuration settings

system component inventory

documented responses to unauthorized changes to system configuration settings

change control records

system audit records

other relevant documents or records

Organizational personnel with security configuration management responsibilities

organizational personnel with security and privacy responsibilities

system/network administrators

Organizational process for responding to unauthorized changes to system configuration settings

mechanisms supporting and/or implementing actions in response to unauthorized changes

### CM-8 (2): Automated Maintenance

Maintain the currency, completeness, accuracy, and availability of the inventory of system components using {{ insert: param, cm-8.2_prm_1 }}.

Organizations maintain system inventories to the extent feasible. For example, virtual machines can be difficult to monitor because such machines are not visible to the network when not in use. In such cases, organizations maintain as up-to-date, complete, and accurate an inventory as is deemed reasonable. Automated maintenance can be achieved by the implementation of [CM-2(2)](#cm-2.2) for organizations that combine system component inventory and baseline configuration activities.

 {{ insert: param, cm-08.02_odp.01 }} are used to maintain the currency of the system component inventory;

 {{ insert: param, cm-08.02_odp.02 }} are used to maintain the completeness of the system component inventory;

 {{ insert: param, cm-08.02_odp.03 }} are used to maintain the accuracy of the system component inventory;

 {{ insert: param, cm-08.02_odp.04 }} are used to maintain the availability of the system component inventory.

Configuration management policy

procedures addressing system component inventory

configuration management plan

system design documentation

system security plan

system component inventory

change control records

system maintenance records

system audit records

system security plan

other relevant documents or records

Organizational personnel with component inventory management responsibilities

organizational personnel with information security responsibilities

system/network administrators

system developers

Organizational processes for maintaining the system component inventory

automated mechanisms supporting and/or implementing the system component inventory

### CM-8 (4): Accountability Information

Include in the system component inventory information, a means for identifying by {{ insert: param, cm-08.04_odp }} , individuals responsible and accountable for administering those components.

Identifying individuals who are responsible and accountable for administering system components ensures that the assigned components are properly administered and that organizations can contact those individuals if some action is required (e.g., when the component is determined to be the source of a breach, needs to be recalled or replaced, or needs to be relocated).

individuals responsible and accountable for administering system components are identified by {{ insert: param, cm-08.04_odp }} in the system component inventory.

Configuration management policy

procedures addressing system component inventory

configuration management plan

system security plan

system component inventory

system security plan

other relevant documents or records

Organizational personnel with component inventory management responsibilities

organizational personnel with information security responsibilities

system/network administrators

Organizational processes for managing the system component inventory

mechanisms supporting and/or implementing the system component inventory

