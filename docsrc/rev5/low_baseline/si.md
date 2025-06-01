# SI - System and Information Integrity

* Controls: 6

## Controls

### SI-1: Policy and Procedures

Develop, document, and disseminate to organization-defined personnel or roles:

organization-level, mission/business process-level, and/or system-level system and information integrity policy that:

Addresses purpose, scope, roles, responsibilities, management commitment, coordination among organizational entities, and compliance; and

Is consistent with applicable laws, executive orders, directives, regulations, policies, standards, and guidelines; and

Procedures to facilitate the implementation of the system and information integrity policy and the associated system and information integrity controls;

Designate an an official to manage the system and information integrity policy and procedures is defined; to manage the development, documentation, and dissemination of the system and information integrity policy and procedures; and

Review and update the current system and information integrity:

Policy the frequency at which the current system and information integrity policy is reviewed and updated is defined; and following events that would require the current system and information integrity policy to be reviewed and updated are defined; ; and

Procedures the frequency at which the current system and information integrity procedures are reviewed and updated is defined; and following events that would require the system and information integrity procedures to be reviewed and updated are defined;.

System and information integrity policy and procedures address the controls in the SI family that are implemented within systems and organizations. The risk management strategy is an important factor in establishing such policies and procedures. Policies and procedures contribute to security and privacy assurance. Therefore, it is important that security and privacy programs collaborate on the development of system and information integrity policy and procedures. Security and privacy program policies and procedures at the organization level are preferable, in general, and may obviate the need for mission- or system-specific policies and procedures. The policy can be included as part of the general security and privacy policy or be represented by multiple policies that reflect the complex nature of organizations. Procedures can be established for security and privacy programs, for mission or business processes, and for systems, if needed. Procedures describe how the policies or controls are implemented and can be directed at the individual or role that is the object of the procedure. Procedures can be documented in system security and privacy plans or in one or more separate documents. Events that may precipitate an update to system and information integrity policy and procedures include assessment or audit findings, security incidents or breaches, or changes in applicable laws, executive orders, directives, regulations, policies, standards, and guidelines. Simply restating controls does not constitute an organizational policy or procedure.

a system and information integrity policy is developed and documented;

the system and information integrity policy is disseminated to personnel or roles to whom the system and information integrity policy is to be disseminated is/are defined;;

system and information integrity procedures to facilitate the implementation of the system and information integrity policy and associated system and information integrity controls are developed and documented;

the system and information integrity procedures are disseminated to personnel or roles to whom the system and information integrity procedures are to be disseminated is/are defined;;

the organization-level, mission/business process-level, and/or system-level system and information integrity policy addresses purpose;

the organization-level, mission/business process-level, and/or system-level system and information integrity policy addresses scope;

the organization-level, mission/business process-level, and/or system-level system and information integrity policy addresses roles;

the organization-level, mission/business process-level, and/or system-level system and information integrity policy addresses responsibilities;

the organization-level, mission/business process-level, and/or system-level system and information integrity policy addresses management commitment;

the organization-level, mission/business process-level, and/or system-level system and information integrity policy addresses coordination among organizational entities;

the organization-level, mission/business process-level, and/or system-level system and information integrity policy addresses compliance;

the organization-level, mission/business process-level, and/or system-level system and information integrity policy is consistent with applicable laws, Executive Orders, directives, regulations, policies, standards, and guidelines;

the an official to manage the system and information integrity policy and procedures is defined; is designated to manage the development, documentation, and dissemination of the system and information integrity policy and procedures;

the current system and information integrity policy is reviewed and updated the frequency at which the current system and information integrity policy is reviewed and updated is defined;;

the current system and information integrity policy is reviewed and updated following events that would require the current system and information integrity policy to be reviewed and updated are defined;;

the current system and information integrity procedures are reviewed and updated the frequency at which the current system and information integrity procedures are reviewed and updated is defined;;

the current system and information integrity procedures are reviewed and updated following events that would require the system and information integrity procedures to be reviewed and updated are defined;.

System and information integrity policy

system and information integrity procedures

system security plan

privacy plan

other relevant documents or records

Organizational personnel with system and information integrity responsibilities

organizational personnel with information security and privacy responsibilities

### SI-2: Flaw Remediation

Identify, report, and correct system flaws;

Test software and firmware updates related to flaw remediation for effectiveness and potential side effects before installation;

Install security-relevant software and firmware updates within time period within which to install security-relevant software updates after the release of the updates is defined; of the release of the updates; and

Incorporate flaw remediation into the organizational configuration management process.

The need to remediate system flaws applies to all types of software and firmware. Organizations identify systems affected by software flaws, including potential vulnerabilities resulting from those flaws, and report this information to designated organizational personnel with information security and privacy responsibilities. Security-relevant updates include patches, service packs, and malicious code signatures. Organizations also address flaws discovered during assessments, continuous monitoring, incident response activities, and system error handling. By incorporating flaw remediation into configuration management processes, required remediation actions can be tracked and verified.

Organization-defined time periods for updating security-relevant software and firmware may vary based on a variety of risk factors, including the security category of the system, the criticality of the update (i.e., severity of the vulnerability related to the discovered flaw), the organizational risk tolerance, the mission supported by the system, or the threat environment. Some types of flaw remediation may require more testing than other types. Organizations determine the type of testing needed for the specific type of flaw remediation activity under consideration and the types of changes that are to be configuration-managed. In some situations, organizations may determine that the testing of software or firmware updates is not necessary or practical, such as when implementing simple malicious code signature updates. In testing decisions, organizations consider whether security-relevant software or firmware updates are obtained from authorized sources with appropriate digital signatures.

system flaws are identified;

system flaws are reported;

system flaws are corrected;

software updates related to flaw remediation are tested for effectiveness before installation;

software updates related to flaw remediation are tested for potential side effects before installation;

firmware updates related to flaw remediation are tested for effectiveness before installation;

firmware updates related to flaw remediation are tested for potential side effects before installation;

security-relevant software updates are installed within time period within which to install security-relevant software updates after the release of the updates is defined; of the release of the updates;

security-relevant firmware updates are installed within time period within which to install security-relevant software updates after the release of the updates is defined; of the release of the updates;

flaw remediation is incorporated into the organizational configuration management process.

System and information integrity policy

system and information integrity procedures

procedures addressing flaw remediation

procedures addressing configuration management

list of flaws and vulnerabilities potentially affecting the system

list of recent security flaw remediation actions performed on the system (e.g., list of installed patches, service packs, hot fixes, and other software updates to correct system flaws)

test results from the installation of software and firmware updates to correct system flaws

installation/change control records for security-relevant software and firmware updates

system security plan

privacy plan

other relevant documents or records

System/network administrators

organizational personnel with information security and privacy responsibilities

organizational personnel responsible for installing, configuring, and/or maintaining the system

organizational personnel responsible for flaw remediation

organizational personnel with configuration management responsibilities

Organizational processes for identifying, reporting, and correcting system flaws

organizational process for installing software and firmware updates

mechanisms supporting and/or implementing the reporting and correcting of system flaws

mechanisms supporting and/or implementing testing software and firmware updates

### SI-3: Malicious Code Protection

Implement signature-basedand/ornon-signature-based malicious code protection mechanisms at system entry and exit points to detect and eradicate malicious code;

Automatically update malicious code protection mechanisms as new releases are available in accordance with organizational configuration management policy and procedures;

Configure malicious code protection mechanisms to:

Perform periodic scans of the system the frequency at which malicious code protection mechanisms perform scans is defined; and real-time scans of files from external sources at endpointand/ornetwork entry and exit points as the files are downloaded, opened, or executed in accordance with organizational policy; and

 block malicious code, quarantine malicious code, and/or take action to be taken in response to malicious code detection are defined (if selected); ; and send alert to personnel or roles to be alerted when malicious code is detected is/are defined; in response to malicious code detection; and

Address the receipt of false positives during malicious code detection and eradication and the resulting potential impact on the availability of the system.

System entry and exit points include firewalls, remote access servers, workstations, electronic mail servers, web servers, proxy servers, notebook computers, and mobile devices. Malicious code includes viruses, worms, Trojan horses, and spyware. Malicious code can also be encoded in various formats contained within compressed or hidden files or hidden in files using techniques such as steganography. Malicious code can be inserted into systems in a variety of ways, including by electronic mail, the world-wide web, and portable storage devices. Malicious code insertions occur through the exploitation of system vulnerabilities. A variety of technologies and methods exist to limit or eliminate the effects of malicious code.

Malicious code protection mechanisms include both signature- and nonsignature-based technologies. Nonsignature-based detection mechanisms include artificial intelligence techniques that use heuristics to detect, analyze, and describe the characteristics or behavior of malicious code and to provide controls against such code for which signatures do not yet exist or for which existing signatures may not be effective. Malicious code for which active signatures do not yet exist or may be ineffective includes polymorphic malicious code (i.e., code that changes signatures when it replicates). Nonsignature-based mechanisms also include reputation-based technologies. In addition to the above technologies, pervasive configuration management, comprehensive software integrity controls, and anti-exploitation software may be effective in preventing the execution of unauthorized code. Malicious code may be present in commercial off-the-shelf software as well as custom-built software and could include logic bombs, backdoors, and other types of attacks that could affect organizational mission and business functions.

In situations where malicious code cannot be detected by detection methods or technologies, organizations rely on other types of controls, including secure coding practices, configuration management and control, trusted procurement processes, and monitoring practices to ensure that software does not perform functions other than the functions intended. Organizations may determine that, in response to the detection of malicious code, different actions may be warranted. For example, organizations can define actions in response to malicious code detection during periodic scans, the detection of malicious downloads, or the detection of maliciousness when attempting to open or execute files.

signature-basedand/ornon-signature-based malicious code protection mechanisms are implemented at system entry and exit points to detect malicious code;

 signature-basedand/ornon-signature-based malicious code protection mechanisms are implemented at system entry and exit points to eradicate malicious code;

malicious code protection mechanisms are updated automatically as new releases are available in accordance with organizational configuration management policy and procedures;

malicious code protection mechanisms are configured to perform periodic scans of the system the frequency at which malicious code protection mechanisms perform scans is defined;;

malicious code protection mechanisms are configured to perform real-time scans of files from external sources at endpointand/ornetwork entry and exit points as the files are downloaded, opened, or executed in accordance with organizational policy;

malicious code protection mechanisms are configured to block malicious code, quarantine malicious code, and/or take action to be taken in response to malicious code detection are defined (if selected); in response to malicious code detection;

malicious code protection mechanisms are configured to send alerts to personnel or roles to be alerted when malicious code is detected is/are defined; in response to malicious code detection;

the receipt of false positives during malicious code detection and eradication and the resulting potential impact on the availability of the system are addressed.

System and information integrity policy

system and information integrity procedures

configuration management policy and procedures

procedures addressing malicious code protection

malicious code protection mechanisms

records of malicious code protection updates

system design documentation

system configuration settings and associated documentation

scan results from malicious code protection mechanisms

record of actions initiated by malicious code protection mechanisms in response to malicious code detection

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

organizational personnel installing, configuring, and/or maintaining the system

organizational personnel responsible for malicious code protection

organizational personnel with configuration management responsibilities

Organizational processes for employing, updating, and configuring malicious code protection mechanisms

organizational processes for addressing false positives and resulting potential impacts

mechanisms supporting and/or implementing, employing, updating, and configuring malicious code protection mechanisms

mechanisms supporting and/or implementing malicious code scanning and subsequent actions

### SI-4: System Monitoring

Monitor the system to detect:

Attacks and indicators of potential attacks in accordance with the following monitoring objectives: monitoring objectives to detect attacks and indicators of potential attacks on the system are defined; ; and

Unauthorized local, network, and remote connections;

Identify unauthorized use of the system through the following techniques and methods: techniques and methods used to identify unauthorized use of the system are defined;;

Invoke internal monitoring capabilities or deploy monitoring devices:

Strategically within the system to collect organization-determined essential information; and

At ad hoc locations within the system to track specific types of transactions of interest to the organization;

Analyze detected events and anomalies;

Adjust the level of system monitoring activity when there is a change in risk to organizational operations and assets, individuals, other organizations, or the Nation;

Obtain legal opinion regarding system monitoring activities; and

Provide system monitoring information to be provided to personnel or roles is defined; to personnel or roles to whom system monitoring information is to be provided is/are defined; as neededand/or a frequency for providing system monitoring to personnel or roles is defined (if selected);.

System monitoring includes external and internal monitoring. External monitoring includes the observation of events occurring at external interfaces to the system. Internal monitoring includes the observation of events occurring within the system. Organizations monitor systems by observing audit activities in real time or by observing other system aspects such as access patterns, characteristics of access, and other actions. The monitoring objectives guide and inform the determination of the events. System monitoring capabilities are achieved through a variety of tools and techniques, including intrusion detection and prevention systems, malicious code protection software, scanning tools, audit record monitoring software, and network monitoring software.

Depending on the security architecture, the distribution and configuration of monitoring devices may impact throughput at key internal and external boundaries as well as at other locations across a network due to the introduction of network throughput latency. If throughput management is needed, such devices are strategically located and deployed as part of an established organization-wide security architecture. Strategic locations for monitoring devices include selected perimeter locations and near key servers and server farms that support critical applications. Monitoring devices are typically employed at the managed interfaces associated with controls [SC-7](#sc-7) and [AC-17](#ac-17) . The information collected is a function of the organizational monitoring objectives and the capability of systems to support such objectives. Specific types of transactions of interest include Hypertext Transfer Protocol (HTTP) traffic that bypasses HTTP proxies. System monitoring is an integral part of organizational continuous monitoring and incident response programs, and output from system monitoring serves as input to those programs. System monitoring requirements, including the need for specific types of system monitoring, may be referenced in other controls (e.g., [AC-2g](#ac-2_smt.g), [AC-2(7)](#ac-2.7), [AC-2(12)(a)](#ac-2.12_smt.a), [AC-17(1)](#ac-17.1), [AU-13](#au-13), [AU-13(1)](#au-13.1), [AU-13(2)](#au-13.2), [CM-3f](#cm-3_smt.f), [CM-6d](#cm-6_smt.d), [MA-3a](#ma-3_smt.a), [MA-4a](#ma-4_smt.a), [SC-5(3)(b)](#sc-5.3_smt.b), [SC-7a](#sc-7_smt.a), [SC-7(24)(b)](#sc-7.24_smt.b), [SC-18b](#sc-18_smt.b), [SC-43b](#sc-43_smt.b) ). Adjustments to levels of system monitoring are based on law enforcement information, intelligence information, or other sources of information. The legality of system monitoring activities is based on applicable laws, executive orders, directives, regulations, policies, standards, and guidelines.

the system is monitored to detect attacks and indicators of potential attacks in accordance with monitoring objectives to detect attacks and indicators of potential attacks on the system are defined;;

the system is monitored to detect unauthorized local connections;

the system is monitored to detect unauthorized network connections;

the system is monitored to detect unauthorized remote connections;

unauthorized use of the system is identified through techniques and methods used to identify unauthorized use of the system are defined;;

internal monitoring capabilities are invoked or monitoring devices are deployed strategically within the system to collect organization-determined essential information;

internal monitoring capabilities are invoked or monitoring devices are deployed at ad hoc locations within the system to track specific types of transactions of interest to the organization;

detected events are analyzed;

detected anomalies are analyzed;

the level of system monitoring activity is adjusted when there is a change in risk to organizational operations and assets, individuals, other organizations, or the Nation;

a legal opinion regarding system monitoring activities is obtained;

 system monitoring information to be provided to personnel or roles is defined; is provided to personnel or roles to whom system monitoring information is to be provided is/are defined; as neededand/or a frequency for providing system monitoring to personnel or roles is defined (if selected);.

System and information integrity policy

system and information integrity procedures

procedures addressing system monitoring tools and techniques

continuous monitoring strategy

facility diagram/layout

system design documentation

system monitoring tools and techniques documentation

locations within the system where monitoring devices are deployed

system configuration settings and associated documentation

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

organizational personnel installing, configuring, and/or maintaining the system

organizational personnel responsible for monitoring the system

Organizational processes for system monitoring

mechanisms supporting and/or implementing system monitoring capabilities

### SI-5: Security Alerts, Advisories, and Directives

Receive system security alerts, advisories, and directives from external organizations from whom system security alerts, advisories, and directives are to be received on an ongoing basis are defined; on an ongoing basis;

Generate internal security alerts, advisories, and directives as deemed necessary;

Disseminate security alerts, advisories, and directives to: personnel or roles to whom security alerts, advisories, and directives are to be disseminated is/are defined (if selected); ,  elements within the organization to whom security alerts, advisories, and directives are to be disseminated are defined (if selected); , and/or  external organizations to whom security alerts, advisories, and directives are to be disseminated are defined (if selected); ; and

Implement security directives in accordance with established time frames, or notify the issuing organization of the degree of noncompliance.

The Cybersecurity and Infrastructure Security Agency (CISA) generates security alerts and advisories to maintain situational awareness throughout the Federal Government. Security directives are issued by OMB or other designated organizations with the responsibility and authority to issue such directives. Compliance with security directives is essential due to the critical nature of many of these directives and the potential (immediate) adverse effects on organizational operations and assets, individuals, other organizations, and the Nation should the directives not be implemented in a timely manner. External organizations include supply chain partners, external mission or business partners, external service providers, and other peer or supporting organizations.

system security alerts, advisories, and directives are received from external organizations from whom system security alerts, advisories, and directives are to be received on an ongoing basis are defined; on an ongoing basis;

internal security alerts, advisories, and directives are generated as deemed necessary;

security alerts, advisories, and directives are disseminated to personnel or roles to whom security alerts, advisories, and directives are to be disseminated is/are defined (if selected); ,  elements within the organization to whom security alerts, advisories, and directives are to be disseminated are defined (if selected); , and/or  external organizations to whom security alerts, advisories, and directives are to be disseminated are defined (if selected);;

security directives are implemented in accordance with established time frames or if the issuing organization is notified of the degree of noncompliance.

System and information integrity policy

system and information integrity procedures

procedures addressing security alerts, advisories, and directives

records of security alerts and advisories

system security plan

other relevant documents or records

Organizational personnel with security alert and advisory responsibilities

organizational personnel implementing, operating, maintaining, and using the system

organizational personnel, organizational elements, and/or external organizations to whom alerts, advisories, and directives are to be disseminated

system/network administrators

organizational personnel with information security responsibilities

Organizational processes for defining, receiving, generating, disseminating, and complying with security alerts, advisories, and directives

mechanisms supporting and/or implementing the definition, receipt, generation, and dissemination of security alerts, advisories, and directives

mechanisms supporting and/or implementing security directives

### SI-12: Information Management and Retention

Manage and retain information within the system and information output from the system in accordance with applicable laws, executive orders, directives, regulations, policies, standards, guidelines and operational requirements.

Information management and retention requirements cover the full life cycle of information, in some cases extending beyond system disposal. Information to be retained may also include policies, procedures, plans, reports, data output from control implementation, and other types of administrative information. The National Archives and Records Administration (NARA) provides federal policy and guidance on records retention and schedules. If organizations have a records management office, consider coordinating with records management personnel. Records produced from the output of implemented controls that may require management and retention include, but are not limited to: All XX-1, [AC-6(9)](#ac-6.9), [AT-4](#at-4), [AU-12](#au-12), [CA-2](#ca-2), [CA-3](#ca-3), [CA-5](#ca-5), [CA-6](#ca-6), [CA-7](#ca-7), [CA-8](#ca-8), [CA-9](#ca-9), [CM-2](#cm-2), [CM-3](#cm-3), [CM-4](#cm-4), [CM-6](#cm-6), [CM-8](#cm-8), [CM-9](#cm-9), [CM-12](#cm-12), [CM-13](#cm-13), [CP-2](#cp-2), [IR-6](#ir-6), [IR-8](#ir-8), [MA-2](#ma-2), [MA-4](#ma-4), [PE-2](#pe-2), [PE-8](#pe-8), [PE-16](#pe-16), [PE-17](#pe-17), [PL-2](#pl-2), [PL-4](#pl-4), [PL-7](#pl-7), [PL-8](#pl-8), [PM-5](#pm-5), [PM-8](#pm-8), [PM-9](#pm-9), [PM-18](#pm-18), [PM-21](#pm-21), [PM-27](#pm-27), [PM-28](#pm-28), [PM-30](#pm-30), [PM-31](#pm-31), [PS-2](#ps-2), [PS-6](#ps-6), [PS-7](#ps-7), [PT-2](#pt-2), [PT-3](#pt-3), [PT-7](#pt-7), [RA-2](#ra-2), [RA-3](#ra-3), [RA-5](#ra-5), [RA-8](#ra-8), [SA-4](#sa-4), [SA-5](#sa-5), [SA-8](#sa-8), [SA-10](#sa-10), [SI-4](#si-4), [SR-2](#sr-2), [SR-4](#sr-4), [SR-8](#sr-8).

information within the system is managed in accordance with applicable laws, Executive Orders, directives, regulations, policies, standards, guidelines, and operational requirements;

information within the system is retained in accordance with applicable laws, Executive Orders, directives, regulations, policies, standards, guidelines, and operational requirements;

information output from the system is managed in accordance with applicable laws, Executive Orders, directives, regulations, policies, standards, guidelines, and operational requirements;

information output from the system is retained in accordance with applicable laws, Executive Orders, directives, regulations, policies, standards, guidelines, and operational requirements.

System and information integrity policy

system and information integrity procedures

personally identifiable information processing policy

records retention and disposition policy

records retention and disposition procedures

federal laws, Executive Orders, directives, policies, regulations, standards, and operational requirements applicable to information management and retention

media protection policy

media protection procedures

audit findings

system security plan

privacy plan

privacy program plan

personally identifiable information inventory

privacy impact assessment

privacy risk assessment documentation

other relevant documents or records

Organizational personnel with information and records management, retention, and disposition responsibilities

organizational personnel with information security and privacy responsibilities

network administrators

Organizational processes for information management, retention, and disposition

automated mechanisms supporting and/or implementing information management, retention, and disposition

