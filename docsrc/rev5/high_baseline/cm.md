# cm - Configuration Management

* Controls: 32

## Controls

### cm-1: Policy and Procedures

Develop, document, and disseminate to {{ insert: param, cm-1_prm_1 }}:

 {{ insert: param, cm-01_odp.03 }} configuration management policy that:

Addresses purpose, scope, roles, responsibilities, management commitment, coordination among organizational entities, and compliance; and

Is consistent with applicable laws, executive orders, directives, regulations, policies, standards, and guidelines; and

Procedures to facilitate the implementation of the configuration management policy and the associated configuration management controls;

Designate an {{ insert: param, cm-01_odp.04 }} to manage the development, documentation, and dissemination of the configuration management policy and procedures; and

Review and update the current configuration management:

Policy {{ insert: param, cm-01_odp.05 }} and following {{ insert: param, cm-01_odp.06 }} ; and

Procedures {{ insert: param, cm-01_odp.07 }} and following {{ insert: param, cm-01_odp.08 }}.

Configuration management policy and procedures address the controls in the CM family that are implemented within systems and organizations. The risk management strategy is an important factor in establishing such policies and procedures. Policies and procedures contribute to security and privacy assurance. Therefore, it is important that security and privacy programs collaborate on the development of configuration management policy and procedures. Security and privacy program policies and procedures at the organization level are preferable, in general, and may obviate the need for mission- or system-specific policies and procedures. The policy can be included as part of the general security and privacy policy or be represented by multiple policies that reflect the complex nature of organizations. Procedures can be established for security and privacy programs, for mission/business processes, and for systems, if needed. Procedures describe how the policies or controls are implemented and can be directed at the individual or role that is the object of the procedure. Procedures can be documented in system security and privacy plans or in one or more separate documents. Events that may precipitate an update to configuration management policy and procedures include, but are not limited to, assessment or audit findings, security incidents or breaches, or changes in applicable laws, executive orders, directives, regulations, policies, standards, and guidelines. Simply restating controls does not constitute an organizational policy or procedure.

a configuration management policy is developed and documented;

the configuration management policy is disseminated to {{ insert: param, cm-01_odp.01 }};

configuration management procedures to facilitate the implementation of the configuration management policy and associated configuration management controls are developed and documented;

the configuration management procedures are disseminated to {{ insert: param, cm-01_odp.02 }};

the {{ insert: param, cm-01_odp.03 }} of the configuration management policy addresses purpose;

the {{ insert: param, cm-01_odp.03 }} of the configuration management policy addresses scope;

the {{ insert: param, cm-01_odp.03 }} of the configuration management policy addresses roles;

the {{ insert: param, cm-01_odp.03 }} of the configuration management policy addresses responsibilities;

the {{ insert: param, cm-01_odp.03 }} of the configuration management policy addresses management commitment;

the {{ insert: param, cm-01_odp.03 }} of the configuration management policy addresses coordination among organizational entities;

the {{ insert: param, cm-01_odp.03 }} of the configuration management policy addresses compliance;

the configuration management policy is consistent with applicable laws, Executive Orders, directives, regulations, policies, standards, and guidelines;

the {{ insert: param, cm-01_odp.04 }} is designated to manage the development, documentation, and dissemination of the configuration management policy and procedures;

the current configuration management policy is reviewed and updated {{ insert: param, cm-01_odp.05 }}; 

the current configuration management policy is reviewed and updated following {{ insert: param, cm-01_odp.06 }};

the current configuration management procedures are reviewed and updated {{ insert: param, cm-01_odp.07 }}; 

the current configuration management procedures are reviewed and updated following {{ insert: param, cm-01_odp.08 }}.

Configuration management policy and procedures

security and privacy program policies and procedures

assessment or audit findings

documentation of security incidents or breaches

system security plan

privacy plan

risk management strategy

other relevant artifacts, documents, or records

Organizational personnel with configuration management responsibilities

organizational personnel with information security and privacy responsibilities

### cm-2: Baseline Configuration

Develop, document, and maintain under configuration control, a current baseline configuration of the system; and

Review and update the baseline configuration of the system:

 {{ insert: param, cm-02_odp.01 }};

When required due to {{ insert: param, cm-02_odp.02 }} ; and

When system components are installed or upgraded.

Baseline configurations for systems and system components include connectivity, operational, and communications aspects of systems. Baseline configurations are documented, formally reviewed, and agreed-upon specifications for systems or configuration items within those systems. Baseline configurations serve as a basis for future builds, releases, or changes to systems and include security and privacy control implementations, operational procedures, information about system components, network topology, and logical placement of components in the system architecture. Maintaining baseline configurations requires creating new baselines as organizational systems change over time. Baseline configurations of systems reflect the current enterprise architecture.

a current baseline configuration of the system is developed and documented;

a current baseline configuration of the system is maintained under configuration control;

the baseline configuration of the system is reviewed and updated {{ insert: param, cm-02_odp.01 }};

the baseline configuration of the system is reviewed and updated when required due to {{ insert: param, cm-02_odp.02 }};

the baseline configuration of the system is reviewed and updated when system components are installed or upgraded.

Configuration management policy

procedures addressing the baseline configuration of the system

configuration management plan

enterprise architecture documentation

system design documentation

system security plan

privacy plan

system architecture and configuration documentation

system configuration settings and associated documentation

system component inventory

change control records

other relevant documents or records

Organizational personnel with configuration management responsibilities

organizational personnel with information security and privacy responsibilities

system/network administrators

Organizational processes for managing baseline configurations

mechanisms supporting configuration control of the baseline configuration

### cm-2.2: Automation Support for Accuracy and Currency

Maintain the currency, completeness, accuracy, and availability of the baseline configuration of the system using {{ insert: param, cm-02.02_odp }}.

Automated mechanisms that help organizations maintain consistent baseline configurations for systems include configuration management tools, hardware, software, firmware inventory tools, and network management tools. Automated tools can be used at the organization level, mission and business process level, or system level on workstations, servers, notebook computers, network components, or mobile devices. Tools can be used to track version numbers on operating systems, applications, types of software installed, and current patch levels. Automation support for accuracy and currency can be satisfied by the implementation of [CM-8(2)](#cm-8.2) for organizations that combine system component inventory and baseline configuration activities.

the currency of the baseline configuration of the system is maintained using {{ insert: param, cm-02.02_odp }};

the completeness of the baseline configuration of the system is maintained using {{ insert: param, cm-02.02_odp }};

the accuracy of the baseline configuration of the system is maintained using {{ insert: param, cm-02.02_odp }};

the availability of the baseline configuration of the system is maintained using {{ insert: param, cm-02.02_odp }}.

Configuration management policy

procedures addressing the baseline configuration of the system

configuration management plan

system design documentation

system architecture and configuration documentation

system configuration settings and associated documentation

system component inventory

configuration change control records

system security plan

other relevant documents or records

Organizational personnel with configuration management responsibilities

organizational personnel with information security responsibilities

system/network administrators

Organizational processes for managing baseline configurations

automated mechanisms implementing baseline configuration maintenance

### cm-2.3: Retention of Previous Configurations

Retain {{ insert: param, cm-02.03_odp }} of previous versions of baseline configurations of the system to support rollback.

Retaining previous versions of baseline configurations to support rollback include hardware, software, firmware, configuration files, configuration records, and associated documentation.

 {{ insert: param, cm-02.03_odp }} of previous baseline configuration version(s) of the system is/are retained to support rollback.

Configuration management policy

procedures addressing the baseline configuration of the system

configuration management plan

system architecture and configuration documentation

system configuration settings and associated documentation

copies of previous baseline configuration versions

system security plan

other relevant documents or records

Organizational personnel with configuration management responsibilities

organizational personnel with information security responsibilities

system/network administrators

Organizational processes for managing baseline configurations

### cm-2.7: Configure Systems and Components for High-risk Areas

Issue {{ insert: param, cm-02.07_odp.01 }} with {{ insert: param, cm-02.07_odp.02 }} to individuals traveling to locations that the organization deems to be of significant risk; and

Apply the following controls to the systems or components when the individuals return from travel: {{ insert: param, cm-02.07_odp.03 }}.

When it is known that systems or system components will be in high-risk areas external to the organization, additional controls may be implemented to counter the increased threat in such areas. For example, organizations can take actions for notebook computers used by individuals departing on and returning from travel. Actions include determining the locations that are of concern, defining the required configurations for the components, ensuring that components are configured as intended before travel is initiated, and applying controls to the components after travel is completed. Specially configured notebook computers include computers with sanitized hard drives, limited applications, and more stringent configuration settings. Controls applied to mobile devices upon return from travel include examining the mobile device for signs of physical tampering and purging and reimaging disk drives. Protecting information that resides on mobile devices is addressed in the [MP](#mp) (Media Protection) family.

 {{ insert: param, cm-02.07_odp.01 }} with {{ insert: param, cm-02.07_odp.02 }} are issued to individuals traveling to locations that the organization deems to be of significant risk;

 {{ insert: param, cm-02.07_odp.03 }} are applied to the systems or system components when the individuals return from travel.

Configuration management policy

configuration management plan

procedures addressing the baseline configuration of the system

procedures addressing system component installations and upgrades

system architecture and configuration documentation

system configuration settings and associated documentation

system component inventory

records of system baseline configuration reviews and updates

system component installations/upgrades and associated records

change control records

system security plan

other relevant documents or records

Organizational personnel with configuration management responsibilities

organizational personnel with information security responsibilities

system/network administrators

Organizational processes for managing baseline configurations

### cm-3: Configuration Change Control

Determine and document the types of changes to the system that are configuration-controlled;

Review proposed configuration-controlled changes to the system and approve or disapprove such changes with explicit consideration for security and privacy impact analyses;

Document configuration change decisions associated with the system;

Implement approved configuration-controlled changes to the system;

Retain records of configuration-controlled changes to the system for {{ insert: param, cm-03_odp.01 }};

Monitor and review activities associated with configuration-controlled changes to the system; and

Coordinate and provide oversight for configuration change control activities through {{ insert: param, cm-03_odp.02 }} that convenes {{ insert: param, cm-03_odp.03 }}.

Configuration change control for organizational systems involves the systematic proposal, justification, implementation, testing, review, and disposition of system changes, including system upgrades and modifications. Configuration change control includes changes to baseline configurations, configuration items of systems, operational procedures, configuration settings for system components, remediate vulnerabilities, and unscheduled or unauthorized changes. Processes for managing configuration changes to systems include Configuration Control Boards or Change Advisory Boards that review and approve proposed changes. For changes that impact privacy risk, the senior agency official for privacy updates privacy impact assessments and system of records notices. For new systems or major upgrades, organizations consider including representatives from the development organizations on the Configuration Control Boards or Change Advisory Boards. Auditing of changes includes activities before and after changes are made to systems and the auditing activities required to implement such changes. See also [SA-10](#sa-10).

the types of changes to the system that are configuration-controlled are determined and documented;

proposed configuration-controlled changes to the system are reviewed;

proposed configuration-controlled changes to the system are approved or disapproved with explicit consideration for security and privacy impact analyses;

configuration change decisions associated with the system are documented;

approved configuration-controlled changes to the system are implemented;

records of configuration-controlled changes to the system are retained for {{ insert: param, cm-03_odp.01 }};

activities associated with configuration-controlled changes to the system are monitored;

activities associated with configuration-controlled changes to the system are reviewed;

configuration change control activities are coordinated and overseen by {{ insert: param, cm-03_odp.02 }};

the configuration control element convenes {{ insert: param, cm-03_odp.03 }}.

Configuration management policy

procedures addressing system configuration change control

configuration management plan

system architecture and configuration documentation

change control records

system audit records

change control audit and review reports

agenda/minutes/documentation from configuration change control oversight meetings

system security plan

privacy plan

privacy impact assessments

system of records notices

other relevant documents or records

Organizational personnel with configuration change control responsibilities

organizational personnel with information security and privacy responsibilities

system/network administrators

members of change control board or similar

Organizational processes for configuration change control

mechanisms that implement configuration change control

### cm-3.1: Automated Documentation, Notification, and Prohibition of Changes

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

### cm-3.2: Testing, Validation, and Documentation of Changes

Test, validate, and document changes to the system before finalizing the implementation of the changes.

Changes to systems include modifications to hardware, software, or firmware components and configuration settings defined in [CM-6](#cm-6) . Organizations ensure that testing does not interfere with system operations that support organizational mission and business functions. Individuals or groups conducting tests understand security and privacy policies and procedures, system security and privacy policies and procedures, and the health, safety, and environmental risks associated with specific facilities or processes. Operational systems may need to be taken offline, or replicated to the extent feasible, before testing can be conducted. If systems must be taken offline for testing, the tests are scheduled to occur during planned system outages whenever possible. If the testing cannot be conducted on operational systems, organizations employ compensating controls.

changes to the system are tested before finalizing the implementation of the changes;

changes to the system are validated before finalizing the implementation of the changes;

changes to the system are documented before finalizing the implementation of the changes.

Configuration management policy

configuration management plan

procedures addressing system configuration change control

system design documentation

system architecture and configuration documentation

system configuration settings and associated documentation

test records

validation records

change control records

system audit records

system security plan

other relevant documents or records

Organizational personnel with configuration change control responsibilities

organizational personnel with information security responsibilities

system/network administrators

system developers

members of change control board or similar

Organizational processes for configuration change control

mechanisms supporting and/or implementing, testing, validating, and documenting system changes

### cm-3.4: Security and Privacy Representatives

Require {{ insert: param, cm-3.4_prm_1 }} to be members of the {{ insert: param, cm-03.04_odp.03 }}.

Information security and privacy representatives include system security officers, senior agency information security officers, senior agency officials for privacy, or system privacy officers. Representation by personnel with information security and privacy expertise is important because changes to system configurations can have unintended side effects, some of which may be security- or privacy-relevant. Detecting such changes early in the process can help avoid unintended, negative consequences that could ultimately affect the security and privacy posture of systems. The configuration change control element referred to in the second organization-defined parameter reflects the change control elements defined by organizations in [CM-3g](#cm-3_smt.g).

 {{ insert: param, cm-03.04_odp.01 }} are required to be members of the {{ insert: param, cm-03.04_odp.03 }};

 {{ insert: param, cm-03.04_odp.02 }} are required to be members of the {{ insert: param, cm-03.04_odp.03 }}.

Configuration management policy

procedures addressing system configuration change control

configuration management plan

system security plan

privacy plan

other relevant documents or records

Organizational personnel with configuration change control responsibilities

organizational personnel with information security and privacy responsibilities

members of change control board or similar

Organizational processes for configuration change control

### cm-3.6: Cryptography Management

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

### cm-4: Impact Analyses

Analyze changes to the system to determine potential security and privacy impacts prior to change implementation.

Organizational personnel with security or privacy responsibilities conduct impact analyses. Individuals conducting impact analyses possess the necessary skills and technical expertise to analyze the changes to systems as well as the security or privacy ramifications. Impact analyses include reviewing security and privacy plans, policies, and procedures to understand control requirements; reviewing system design documentation and operational procedures to understand control implementation and how specific system changes might affect the controls; reviewing the impact of changes on organizational supply chain partners with stakeholders; and determining how potential changes to a system create new risks to the privacy of individuals and the ability of implemented controls to mitigate those risks. Impact analyses also include risk assessments to understand the impact of the changes and determine if additional controls are required.

changes to the system are analyzed to determine potential security impacts prior to change implementation;

changes to the system are analyzed to determine potential privacy impacts prior to change implementation.

Configuration management policy

procedures addressing security impact analyses for changes to the system

procedures addressing privacy impact analyses for changes to the system

configuration management plan

security impact analysis documentation

privacy impact analysis documentation

privacy impact assessment

privacy risk assessment documentation, system design documentation

analysis tools and associated outputs

change control records

system audit records

system security plan

privacy plan

other relevant documents or records

Organizational personnel with responsibility for conducting security impact analyses

organizational personnel with responsibility for conducting privacy impact analyses

organizational personnel with information security and privacy responsibilities

system developer

system/network administrators

members of change control board or similar

Organizational processes for security impact analyses

organizational processes for privacy impact analyses

### cm-4.1: Separate Test Environments

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

### cm-4.2: Verification of Controls

After system changes, verify that the impacted controls are implemented correctly, operating as intended, and producing the desired outcome with regard to meeting the security and privacy requirements for the system.

Implementation in this context refers to installing changed code in the operational system that may have an impact on security or privacy controls.

the impacted controls are implemented correctly with regard to meeting the security requirements for the system after system changes;

the impacted controls are implemented correctly with regard to meeting the privacy requirements for the system after system changes;

the impacted controls are operating as intended with regard to meeting the security requirements for the system after system changes;

the impacted controls are operating as intended with regard to meeting the privacy requirements for the system after system changes;

the impacted controls are producing the desired outcome with regard to meeting the security requirements for the system after system changes;

the impacted controls are producing the desired outcome with regard to meeting the privacy requirements for the system after system changes.

Configuration management policy

procedures addressing security impact analyses for changes to the system

procedures addressing privacy impact analyses for changes to the system

privacy risk assessment documentation

configuration management plan

security and privacy impact analysis documentation

privacy impact assessment

analysis tools and associated outputs

change control records

control assessment results

system audit records

system component inventory

system security plan

privacy plan

other relevant documents or records

Organizational personnel with responsibility for conducting security and privacy impact analyses

organizational personnel with information security and privacy responsibilities

system/network administrators

security and privacy assessors

Organizational processes for security and privacy impact analyses

mechanisms supporting and/or implementing security and privacy impact analyses of changes

### cm-5: Access Restrictions for Change

Define, document, approve, and enforce physical and logical access restrictions associated with changes to the system.

Changes to the hardware, software, or firmware components of systems or the operational procedures related to the system can potentially have significant effects on the security of the systems or individuals’ privacy. Therefore, organizations permit only qualified and authorized individuals to access systems for purposes of initiating changes. Access restrictions include physical and logical access controls (see [AC-3](#ac-3) and [PE-3](#pe-3) ), software libraries, workflow automation, media libraries, abstract layers (i.e., changes implemented into external interfaces rather than directly into systems), and change windows (i.e., changes occur only during specified times).

physical access restrictions associated with changes to the system are defined and documented;

physical access restrictions associated with changes to the system are approved;

physical access restrictions associated with changes to the system are enforced;

logical access restrictions associated with changes to the system are defined and documented;

logical access restrictions associated with changes to the system are approved;

logical access restrictions associated with changes to the system are enforced.

Configuration management policy

procedures addressing access restrictions for changes to the system

configuration management plan

system design documentation

system architecture and configuration documentation

system configuration settings and associated documentation

logical access approvals

physical access approvals

access credentials

change control records

system audit records

system security plan

other relevant documents or records

Organizational personnel with logical access control responsibilities

organizational personnel with physical access control responsibilities

organizational personnel with information security responsibilities

system/network administrators

Organizational processes for managing access restrictions to change

mechanisms supporting, implementing, or enforcing access restrictions associated with changes to the system

### cm-5.1: Automated Access Enforcement and Audit Records

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

### cm-6: Configuration Settings

Establish and document configuration settings for components employed within the system that reflect the most restrictive mode consistent with operational requirements using {{ insert: param, cm-06_odp.01 }};

Implement the configuration settings;

Identify, document, and approve any deviations from established configuration settings for {{ insert: param, cm-06_odp.02 }} based on {{ insert: param, cm-06_odp.03 }} ; and

Monitor and control changes to the configuration settings in accordance with organizational policies and procedures.

Configuration settings are the parameters that can be changed in the hardware, software, or firmware components of the system that affect the security and privacy posture or functionality of the system. Information technology products for which configuration settings can be defined include mainframe computers, servers, workstations, operating systems, mobile devices, input/output devices, protocols, and applications. Parameters that impact the security posture of systems include registry settings; account, file, or directory permission settings; and settings for functions, protocols, ports, services, and remote connections. Privacy parameters are parameters impacting the privacy posture of systems, including the parameters required to satisfy other privacy controls. Privacy parameters include settings for access controls, data processing preferences, and processing and retention permissions. Organizations establish organization-wide configuration settings and subsequently derive specific configuration settings for systems. The established settings become part of the configuration baseline for the system.

Common secure configurations (also known as security configuration checklists, lockdown and hardening guides, and security reference guides) provide recognized, standardized, and established benchmarks that stipulate secure configuration settings for information technology products and platforms as well as instructions for configuring those products or platforms to meet operational requirements. Common secure configurations can be developed by a variety of organizations, including information technology product developers, manufacturers, vendors, federal agencies, consortia, academia, industry, and other organizations in the public and private sectors.

Implementation of a common secure configuration may be mandated at the organization level, mission and business process level, system level, or at a higher level, including by a regulatory agency. Common secure configurations include the United States Government Configuration Baseline [USGCB](#98498928-3ca3-44b3-8b1e-f48685373087) and security technical implementation guides (STIGs), which affect the implementation of [CM-6](#cm-6) and other controls such as [AC-19](#ac-19) and [CM-7](#cm-7) . The Security Content Automation Protocol (SCAP) and the defined standards within the protocol provide an effective method to uniquely identify, track, and control configuration settings.

configuration settings that reflect the most restrictive mode consistent with operational requirements are established and documented for components employed within the system using {{ insert: param, cm-06_odp.01 }};

the configuration settings documented in CM-06a are implemented;

any deviations from established configuration settings for {{ insert: param, cm-06_odp.02 }} are identified and documented based on {{ insert: param, cm-06_odp.03 }};

any deviations from established configuration settings for {{ insert: param, cm-06_odp.02 }} are approved;

changes to the configuration settings are monitored in accordance with organizational policies and procedures;

changes to the configuration settings are controlled in accordance with organizational policies and procedures.

Configuration management policy

procedures addressing configuration settings for the system

configuration management plan

system design documentation

system configuration settings and associated documentation

common secure configuration checklists

system component inventory

evidence supporting approved deviations from established configuration settings

change control records

system data processing and retention permissions

system audit records

system security plan

privacy plan

other relevant documents or records

Organizational personnel with security configuration management responsibilities

organizational personnel with privacy configuration management responsibilities

organizational personnel with information security and privacy responsibilities

system/network administrators

Organizational processes for managing configuration settings

mechanisms that implement, monitor, and/or control system configuration settings

mechanisms that identify and/or document deviations from established configuration settings

### cm-6.1: Automated Management, Application, and Verification

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

### cm-6.2: Respond to Unauthorized Changes

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

### cm-7: Least Functionality

Configure the system to provide only {{ insert: param, cm-07_odp.01 }} ; and

Prohibit or restrict the use of the following functions, ports, protocols, software, and/or services: {{ insert: param, cm-7_prm_2 }}.

Systems provide a wide variety of functions and services. Some of the functions and services routinely provided by default may not be necessary to support essential organizational missions, functions, or operations. Additionally, it is sometimes convenient to provide multiple services from a single system component, but doing so increases risk over limiting the services provided by that single component. Where feasible, organizations limit component functionality to a single function per component. Organizations consider removing unused or unnecessary software and disabling unused or unnecessary physical and logical ports and protocols to prevent unauthorized connection of components, transfer of information, and tunneling. Organizations employ network scanning tools, intrusion detection and prevention systems, and end-point protection technologies, such as firewalls and host-based intrusion detection systems, to identify and prevent the use of prohibited functions, protocols, ports, and services. Least functionality can also be achieved as part of the fundamental design and development of the system (see [SA-8](#sa-8), [SC-2](#sc-2) , and [SC-3](#sc-3)).

the system is configured to provide only {{ insert: param, cm-07_odp.01 }};

the use of {{ insert: param, cm-07_odp.02 }} is prohibited or restricted;

the use of {{ insert: param, cm-07_odp.03 }} is prohibited or restricted;

the use of {{ insert: param, cm-07_odp.04 }} is prohibited or restricted;

the use of {{ insert: param, cm-07_odp.05 }} is prohibited or restricted;

the use of {{ insert: param, cm-07_odp.06 }} is prohibited or restricted.

Configuration management policy

procedures addressing least functionality in the system

configuration management plan

system design documentation

system configuration settings and associated documentation

system component inventory

common secure configuration checklists

system security plan

other relevant documents or records

Organizational personnel with security configuration management responsibilities

organizational personnel with information security responsibilities

system/network administrators

system developers

Organizational processes prohibiting or restricting functions, ports, protocols, software, and/or services

mechanisms implementing restrictions or prohibition of functions, ports, protocols, software, and/or services

### cm-7.1: Periodic Review

Review the system {{ insert: param, cm-07.01_odp.01 }} to identify unnecessary and/or nonsecure functions, ports, protocols, software, and services; and

Disable or remove {{ insert: param, cm-7.1_prm_2 }}.

Organizations review functions, ports, protocols, and services provided by systems or system components to determine the functions and services that are candidates for elimination. Such reviews are especially important during transition periods from older technologies to newer technologies (e.g., transition from IPv4 to IPv6). These technology transitions may require implementing the older and newer technologies simultaneously during the transition period and returning to minimum essential functions, ports, protocols, and services at the earliest opportunity. Organizations can either decide the relative security of the function, port, protocol, and/or service or base the security decision on the assessment of other entities. Unsecure protocols include Bluetooth, FTP, and peer-to-peer networking.

the system is reviewed {{ insert: param, cm-07.01_odp.01 }} to identify unnecessary and/or non-secure functions, ports, protocols, software, and services:

 {{ insert: param, cm-07.01_odp.02 }} deemed to be unnecessary and/or non-secure are disabled or removed;

 {{ insert: param, cm-07.01_odp.03 }} deemed to be unnecessary and/or non-secure are disabled or removed;

 {{ insert: param, cm-07.01_odp.04 }} deemed to be unnecessary and/or non-secure are disabled or removed;

 {{ insert: param, cm-07.01_odp.05 }} deemed to be unnecessary and/or non-secure is disabled or removed;

 {{ insert: param, cm-07.01_odp.06 }} deemed to be unnecessary and/or non-secure are disabled or removed.

Configuration management policy

procedures addressing least functionality in the system

configuration management plan

system design documentation

system configuration settings and associated documentation

common secure configuration checklists

documented reviews of functions, ports, protocols, and/or services

change control records

system audit records

system security plan

other relevant documents or records

Organizational personnel with responsibilities for reviewing functions, ports, protocols, and services on the system

organizational personnel with information security responsibilities

system/network administrators

system developers

Organizational processes for reviewing or disabling functions, ports, protocols, and services on the system

mechanisms implementing review and disabling of functions, ports, protocols, and/or services

### cm-7.2: Prevent Program Execution

Prevent program execution in accordance with {{ insert: param, cm-07.02_odp.01 }}.

Prevention of program execution addresses organizational policies, rules of behavior, and/or access agreements that restrict software usage and the terms and conditions imposed by the developer or manufacturer, including software licensing and copyrights. Restrictions include prohibiting auto-execute features, restricting roles allowed to approve program execution, permitting or prohibiting specific software programs, or restricting the number of program instances executed at the same time.

program execution is prevented in accordance with {{ insert: param, cm-07.02_odp.01 }}.

Configuration management policy

procedures addressing least functionality in the system

configuration management plan

system design documentation

system configuration settings and associated documentation

system component inventory

common secure configuration checklists

specifications for preventing software program execution

change control records

system audit records

system security plan

other relevant documents or records

Organizational personnel with information security responsibilities

system/network administrators

system developers

Organizational processes preventing program execution on the system

organizational processes for software program usage and restrictions

mechanisms preventing program execution on the system

mechanisms supporting and/or implementing software program usage and restrictions

### cm-7.5: Authorized Software — Allow-by-exception

Identify {{ insert: param, cm-07.05_odp.01 }};

Employ a deny-all, permit-by-exception policy to allow the execution of authorized software programs on the system; and

Review and update the list of authorized software programs {{ insert: param, cm-07.05_odp.02 }}.

Authorized software programs can be limited to specific versions or from a specific source. To facilitate a comprehensive authorized software process and increase the strength of protection for attacks that bypass application level authorized software, software programs may be decomposed into and monitored at different levels of detail. These levels include applications, application programming interfaces, application modules, scripts, system processes, system services, kernel functions, registries, drivers, and dynamic link libraries. The concept of permitting the execution of authorized software may also be applied to user actions, system ports and protocols, IP addresses/ranges, websites, and MAC addresses. Organizations consider verifying the integrity of authorized software programs using digital signatures, cryptographic checksums, or hash functions. Verification of authorized software can occur either prior to execution or at system startup. The identification of authorized URLs for websites is addressed in [CA-3(5)](#ca-3.5) and [SC-7](#sc-7).

 {{ insert: param, cm-07.05_odp.01 }} are identified;

a deny-all, permit-by-exception policy to allow the execution of authorized software programs on the system is employed;

the list of authorized software programs is reviewed and updated {{ insert: param, cm-07.05_odp.02 }}.

Configuration management policy

procedures addressing least functionality in the system

configuration management plan

system design documentation

system configuration settings and associated documentation

list of software programs authorized to execute on the system

system component inventory

common secure configuration checklists

review and update records associated with list of authorized software programs

change control records

system audit records

system security plan

other relevant documents or records

Organizational personnel with responsibilities for identifying software authorized to execute on the system

organizational personnel with information security responsibilities

system/network administrators

Organizational process for identifying, reviewing, and updating programs authorized to execute on the system

organizational process for implementing authorized software policy

mechanisms supporting and/or implementing authorized software policy

### cm-8: System Component Inventory

Develop and document an inventory of system components that:

Accurately reflects the system;

Includes all components within the system;

Does not include duplicate accounting of components or components assigned to any other system;

Is at the level of granularity deemed necessary for tracking and reporting; and

Includes the following information to achieve system component accountability: {{ insert: param, cm-08_odp.01 }} ; and

Review and update the system component inventory {{ insert: param, cm-08_odp.02 }}.

System components are discrete, identifiable information technology assets that include hardware, software, and firmware. Organizations may choose to implement centralized system component inventories that include components from all organizational systems. In such situations, organizations ensure that the inventories include system-specific information required for component accountability. The information necessary for effective accountability of system components includes the system name, software owners, software version numbers, hardware inventory specifications, software license information, and for networked components, the machine names and network addresses across all implemented protocols (e.g., IPv4, IPv6). Inventory specifications include date of receipt, cost, model, serial number, manufacturer, supplier information, component type, and physical location.

Preventing duplicate accounting of system components addresses the lack of accountability that occurs when component ownership and system association is not known, especially in large or complex connected systems. Effective prevention of duplicate accounting of system components necessitates use of a unique identifier for each component. For software inventory, centrally managed software that is accessed via other systems is addressed as a component of the system on which it is installed and managed. Software installed on multiple organizational systems and managed at the system level is addressed for each individual system and may appear more than once in a centralized component inventory, necessitating a system association for each software instance in the centralized inventory to avoid duplicate accounting of components. Scanning systems implementing multiple network protocols (e.g., IPv4 and IPv6) can result in duplicate components being identified in different address spaces. The implementation of [CM-8(7)](#cm-8.7) can help to eliminate duplicate accounting of components.

an inventory of system components that accurately reflects the system is developed and documented;

an inventory of system components that includes all components within the system is developed and documented;

an inventory of system components that does not include duplicate accounting of components or components assigned to any other system is developed and documented;

an inventory of system components that is at the level of granularity deemed necessary for tracking and reporting is developed and documented;

an inventory of system components that includes {{ insert: param, cm-08_odp.01 }} is developed and documented;

the system component inventory is reviewed and updated {{ insert: param, cm-08_odp.02 }}.

Configuration management policy

procedures addressing system component inventory

configuration management plan

system security plan

system design documentation

system component inventory

inventory reviews and update records

system security plan

other relevant documents or records

Organizational personnel with component inventory management responsibilities

organizational personnel with information security responsibilities

system/network administrators

Organizational processes for managing the system component inventory

mechanisms supporting and/or implementing system component inventory

### cm-8.1: Updates During Installation and Removal

Update the inventory of system components as part of component installations, removals, and system updates.

Organizations can improve the accuracy, completeness, and consistency of system component inventories if the inventories are updated as part of component installations or removals or during general system updates. If inventories are not updated at these key times, there is a greater likelihood that the information will not be appropriately captured and documented. System updates include hardware, software, and firmware components.

the inventory of system components is updated as part of component installations;

the inventory of system components is updated as part of component removals;

the inventory of system components is updated as part of system updates.

Configuration management policy

procedures addressing system component inventory

configuration management plan

system security plan

system component inventory

inventory reviews and update records

change control records

component installation records

component removal records

system security plan

other relevant documents or records

Organizational personnel with component inventory updating responsibilities

organizational personnel with information security responsibilities

system/network administrators

Organizational processes for updating the system component inventory

mechanisms supporting and/or implementing system component inventory updates

### cm-8.2: Automated Maintenance

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

### cm-8.3: Automated Unauthorized Component Detection

Detect the presence of unauthorized hardware, software, and firmware components within the system using {{ insert: param, cm-8.3_prm_1 }} {{ insert: param, cm-08.03_odp.04 }} ; and

Take the following actions when unauthorized components are detected: {{ insert: param, cm-08.03_odp.05 }}.

Automated unauthorized component detection is applied in addition to the monitoring for unauthorized remote connections and mobile devices. Monitoring for unauthorized system components may be accomplished on an ongoing basis or by the periodic scanning of systems for that purpose. Automated mechanisms may also be used to prevent the connection of unauthorized components (see [CM-7(9)](#cm-7.9) ). Automated mechanisms can be implemented in systems or in separate system components. When acquiring and implementing automated mechanisms, organizations consider whether such mechanisms depend on the ability of the system component to support an agent or supplicant in order to be detected since some types of components do not have or cannot support agents (e.g., IoT devices, sensors). Isolation can be achieved , for example, by placing unauthorized system components in separate domains or subnets or quarantining such components. This type of component isolation is commonly referred to as "sandboxing." 

the presence of unauthorized hardware within the system is detected using {{ insert: param, cm-08.03_odp.01 }} {{ insert: param, cm-08.03_odp.04 }};

the presence of unauthorized software within the system is detected using {{ insert: param, cm-08.03_odp.02 }} {{ insert: param, cm-08.03_odp.04 }};

the presence of unauthorized firmware within the system is detected using {{ insert: param, cm-08.03_odp.03 }} {{ insert: param, cm-08.03_odp.04 }};

 {{ insert: param, cm-08.03_odp.05 }} are taken when unauthorized hardware is detected;

 {{ insert: param, cm-08.03_odp.05 }} are taken when unauthorized software is detected;

 {{ insert: param, cm-08.03_odp.05 }} are taken when unauthorized firmware is detected.

Configuration management policy

procedures addressing system component inventory

configuration management plan

system design documentation

system security plan

system component inventory

change control records

alerts/notifications of unauthorized components within the system

system monitoring records

system maintenance records

system audit records

system security plan

other relevant documents or records

Organizational personnel with component inventory management responsibilities

organizational personnel with responsibilities for managing the automated mechanisms implementing unauthorized system component detection

organizational personnel with information security responsibilities

system/network administrators

system developers

Organizational processes for detection of unauthorized system components

organizational processes for taking action when unauthorized system components are detected

automated mechanisms supporting and/or implementing the detection of unauthorized system components

automated mechanisms supporting and/or implementing actions taken when unauthorized system components are detected

### cm-8.4: Accountability Information

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

### cm-9: Configuration Management Plan

Develop, document, and implement a configuration management plan for the system that:

Addresses roles, responsibilities, and configuration management processes and procedures;

Establishes a process for identifying configuration items throughout the system development life cycle and for managing the configuration of the configuration items;

Defines the configuration items for the system and places the configuration items under configuration management;

Is reviewed and approved by {{ insert: param, cm-09_odp }} ; and

Protects the configuration management plan from unauthorized disclosure and modification.

Configuration management activities occur throughout the system development life cycle. As such, there are developmental configuration management activities (e.g., the control of code and software libraries) and operational configuration management activities (e.g., control of installed components and how the components are configured). Configuration management plans satisfy the requirements in configuration management policies while being tailored to individual systems. Configuration management plans define processes and procedures for how configuration management is used to support system development life cycle activities.

Configuration management plans are generated during the development and acquisition stage of the system development life cycle. The plans describe how to advance changes through change management processes; update configuration settings and baselines; maintain component inventories; control development, test, and operational environments; and develop, release, and update key documents.

Organizations can employ templates to help ensure the consistent and timely development and implementation of configuration management plans. Templates can represent a configuration management plan for the organization with subsets of the plan implemented on a system by system basis. Configuration management approval processes include the designation of key stakeholders responsible for reviewing and approving proposed changes to systems, and personnel who conduct security and privacy impact analyses prior to the implementation of changes to the systems. Configuration items are the system components, such as the hardware, software, firmware, and documentation to be configuration-managed. As systems continue through the system development life cycle, new configuration items may be identified, and some existing configuration items may no longer need to be under configuration control.

a configuration management plan for the system is developed and documented;

a configuration management plan for the system is implemented;

the configuration management plan addresses roles;

the configuration management plan addresses responsibilities;

the configuration management plan addresses configuration management processes and procedures;

the configuration management plan establishes a process for identifying configuration items throughout the system development life cycle;

the configuration management plan establishes a process for managing the configuration of the configuration items;

the configuration management plan defines the configuration items for the system;

the configuration management plan places the configuration items under configuration management;

the configuration management plan is reviewed and approved by {{ insert: param, cm-09_odp }};

the configuration management plan is protected from unauthorized disclosure;

the configuration management plan is protected from unauthorized modification.

Configuration management policy

procedures addressing configuration management planning

configuration management plan

system design documentation

system security plan

privacy plan

other relevant documents or records

Organizational personnel with responsibilities for developing the configuration management plan

organizational personnel with responsibilities for implementing and managing processes defined in the configuration management plan

organizational personnel with responsibilities for protecting the configuration management plan

organizational personnel with information security and privacy responsibilities

system/network administrators

Organizational processes for developing and documenting the configuration management plan

organizational processes for identifying and managing configuration items

organizational processes for protecting the configuration management plan

mechanisms implementing the configuration management plan

mechanisms for managing configuration items

mechanisms for protecting the configuration management plan

### cm-10: Software Usage Restrictions

Use software and associated documentation in accordance with contract agreements and copyright laws;

Track the use of software and associated documentation protected by quantity licenses to control copying and distribution; and

Control and document the use of peer-to-peer file sharing technology to ensure that this capability is not used for the unauthorized distribution, display, performance, or reproduction of copyrighted work.

Software license tracking can be accomplished by manual or automated methods, depending on organizational needs. Examples of contract agreements include software license agreements and non-disclosure agreements.

software and associated documentation are used in accordance with contract agreements and copyright laws;

the use of software and associated documentation protected by quantity licenses is tracked to control copying and distribution;

the use of peer-to-peer file sharing technology is controlled and documented to ensure that peer-to-peer file sharing is not used for the unauthorized distribution, display, performance, or reproduction of copyrighted work.

Configuration management policy

software usage restrictions

software contract agreements and copyright laws

site license documentation

list of software usage restrictions

software license tracking reports

configuration management plan

system security plan

system security plan

other relevant documents or records

Organizational personnel operating, using, and/or maintaining the system

organizational personnel with software license management responsibilities

organizational personnel with information security responsibilities

system/network administrators

Organizational processes for tracking the use of software protected by quantity licenses

organizational processes for controlling/documenting the use of peer-to-peer file sharing technology

mechanisms implementing software license tracking

mechanisms implementing and controlling the use of peer-to-peer files sharing technology

### cm-11: User-installed Software

Establish {{ insert: param, cm-11_odp.01 }} governing the installation of software by users;

Enforce software installation policies through the following methods: {{ insert: param, cm-11_odp.02 }} ; and

Monitor policy compliance {{ insert: param, cm-11_odp.03 }}.

If provided the necessary privileges, users can install software in organizational systems. To maintain control over the software installed, organizations identify permitted and prohibited actions regarding software installation. Permitted software installations include updates and security patches to existing software and downloading new applications from organization-approved "app stores." Prohibited software installations include software with unknown or suspect pedigrees or software that organizations consider potentially malicious. Policies selected for governing user-installed software are organization-developed or provided by some external entity. Policy enforcement methods can include procedural methods and automated methods.

 {{ insert: param, cm-11_odp.01 }} governing the installation of software by users are established;

software installation policies are enforced through {{ insert: param, cm-11_odp.02 }};

compliance with {{ insert: param, cm-11_odp.01 }} is monitored {{ insert: param, cm-11_odp.03 }}.

Configuration management policy

procedures addressing user-installed software

configuration management plan

system security plan

system design documentation

system configuration settings and associated documentation

list of rules governing user installed software

system monitoring records

system audit records

continuous monitoring strategy

system security plan

other relevant documents or records

Organizational personnel with responsibilities for governing user-installed software

organizational personnel operating, using, and/or maintaining the system

organizational personnel monitoring compliance with user-installed software policy

organizational personnel with information security responsibilities

system/network administrators

Organizational processes governing user-installed software on the system

mechanisms enforcing policies and methods for governing the installation of software by users

mechanisms monitoring policy compliance

### cm-12: Information Location

Identify and document the location of {{ insert: param, cm-12_odp }} and the specific system components on which the information is processed and stored;

Identify and document the users who have access to the system and system components where the information is processed and stored; and

Document changes to the location (i.e., system or system components) where the information is processed and stored.

Information location addresses the need to understand where information is being processed and stored. Information location includes identifying where specific information types and information reside in system components and how information is being processed so that information flow can be understood and adequate protection and policy management provided for such information and system components. The security category of the information is also a factor in determining the controls necessary to protect the information and the system component where the information resides (see [FIPS 199](#628d22a1-6a11-4784-bc59-5cd9497b5445) ). The location of the information and system components is also a factor in the architecture and design of the system (see [SA-4](#sa-4), [SA-8](#sa-8), [SA-17](#sa-17)).

the location of {{ insert: param, cm-12_odp }} is identified and documented;

the specific system components on which {{ insert: param, cm-12_odp }} is processed are identified and documented;

the specific system components on which {{ insert: param, cm-12_odp }} is stored are identified and documented;

the users who have access to the system and system components where {{ insert: param, cm-12_odp }} is processed are identified and documented;

the users who have access to the system and system components where {{ insert: param, cm-12_odp }} is stored are identified and documented;

changes to the location (i.e., system or system components) where {{ insert: param, cm-12_odp }} is processed are documented;

changes to the location (i.e., system or system components) where {{ insert: param, cm-12_odp }} is stored are documented.

Configuration management policy

procedures addressing identification and documentation of information location

configuration management plan

system design documentation

system architecture documentation

PII inventory documentation

data mapping documentation

audit records

list of users with system and system component access

change control records

system component inventory

system security plan

privacy plan

other relevant documents or records

Organizational personnel with responsibilities for managing information location and user access to information

organizational personnel with responsibilities for operating, using, and/or maintaining the system

organizational personnel with information security and privacy responsibilities

system/network administrators

system developers

Organizational processes governing information location

mechanisms enforcing policies and methods for governing information location

### cm-12.1: Automated Tools to Support Information Location

Use automated tools to identify {{ insert: param, cm-12.01_odp.01 }} on {{ insert: param, cm-12.01_odp.02 }} to ensure controls are in place to protect organizational information and individual privacy.

The use of automated tools helps to increase the effectiveness and efficiency of the information location capability implemented within the system. Automation also helps organizations manage the data produced during information location activities and share such information across the organization. The output of automated information location tools can be used to guide and inform system architecture and design decisions.

automated tools are used to identify {{ insert: param, cm-12.01_odp.01 }} on {{ insert: param, cm-12.01_odp.02 }} to ensure that controls are in place to protect organizational information and individual privacy.

Configuration management policy

procedures addressing identification and documentation of information location

configuration management plan

system design documentation

PII inventory documentation

data mapping documentation

change control records

system component inventory

system security plan

privacy plan

other relevant documents or records

Organizational personnel with responsibilities for managing information location

organizational personnel with information security responsibilities

system/network administrators

system developers

Organizational processes governing information location

automated mechanisms enforcing policies and methods for governing information location

automated tools used to identify information on system components

