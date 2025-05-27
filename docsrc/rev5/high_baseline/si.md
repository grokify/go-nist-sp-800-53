# SI - System and Information Integrity

* Controls: 28

## Controls

### SI-1: Policy and Procedures

Develop, document, and disseminate to {{ insert: param, si-1_prm_1 }}:

 {{ insert: param, si-01_odp.03 }} system and information integrity policy that:

Addresses purpose, scope, roles, responsibilities, management commitment, coordination among organizational entities, and compliance; and

Is consistent with applicable laws, executive orders, directives, regulations, policies, standards, and guidelines; and

Procedures to facilitate the implementation of the system and information integrity policy and the associated system and information integrity controls;

Designate an {{ insert: param, si-01_odp.04 }} to manage the development, documentation, and dissemination of the system and information integrity policy and procedures; and

Review and update the current system and information integrity:

Policy {{ insert: param, si-01_odp.05 }} and following {{ insert: param, si-01_odp.06 }} ; and

Procedures {{ insert: param, si-01_odp.07 }} and following {{ insert: param, si-01_odp.08 }}.

System and information integrity policy and procedures address the controls in the SI family that are implemented within systems and organizations. The risk management strategy is an important factor in establishing such policies and procedures. Policies and procedures contribute to security and privacy assurance. Therefore, it is important that security and privacy programs collaborate on the development of system and information integrity policy and procedures. Security and privacy program policies and procedures at the organization level are preferable, in general, and may obviate the need for mission- or system-specific policies and procedures. The policy can be included as part of the general security and privacy policy or be represented by multiple policies that reflect the complex nature of organizations. Procedures can be established for security and privacy programs, for mission or business processes, and for systems, if needed. Procedures describe how the policies or controls are implemented and can be directed at the individual or role that is the object of the procedure. Procedures can be documented in system security and privacy plans or in one or more separate documents. Events that may precipitate an update to system and information integrity policy and procedures include assessment or audit findings, security incidents or breaches, or changes in applicable laws, executive orders, directives, regulations, policies, standards, and guidelines. Simply restating controls does not constitute an organizational policy or procedure.

a system and information integrity policy is developed and documented;

the system and information integrity policy is disseminated to {{ insert: param, si-01_odp.01 }};

system and information integrity procedures to facilitate the implementation of the system and information integrity policy and associated system and information integrity controls are developed and documented;

the system and information integrity procedures are disseminated to {{ insert: param, si-01_odp.02 }};

the {{ insert: param, si-01_odp.03 }} system and information integrity policy addresses purpose;

the {{ insert: param, si-01_odp.03 }} system and information integrity policy addresses scope;

the {{ insert: param, si-01_odp.03 }} system and information integrity policy addresses roles;

the {{ insert: param, si-01_odp.03 }} system and information integrity policy addresses responsibilities;

the {{ insert: param, si-01_odp.03 }} system and information integrity policy addresses management commitment;

the {{ insert: param, si-01_odp.03 }} system and information integrity policy addresses coordination among organizational entities;

the {{ insert: param, si-01_odp.03 }} system and information integrity policy addresses compliance;

the {{ insert: param, si-01_odp.03 }} system and information integrity policy is consistent with applicable laws, Executive Orders, directives, regulations, policies, standards, and guidelines;

the {{ insert: param, si-01_odp.04 }} is designated to manage the development, documentation, and dissemination of the system and information integrity policy and procedures;

the current system and information integrity policy is reviewed and updated {{ insert: param, si-01_odp.05 }};

the current system and information integrity policy is reviewed and updated following {{ insert: param, si-01_odp.06 }};

the current system and information integrity procedures are reviewed and updated {{ insert: param, si-01_odp.07 }};

the current system and information integrity procedures are reviewed and updated following {{ insert: param, si-01_odp.08 }}.

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

Install security-relevant software and firmware updates within {{ insert: param, si-02_odp }} of the release of the updates; and

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

security-relevant software updates are installed within {{ insert: param, si-02_odp }} of the release of the updates;

security-relevant firmware updates are installed within {{ insert: param, si-02_odp }} of the release of the updates;

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

### SI-2 (2): Automated Flaw Remediation Status

Determine if system components have applicable security-relevant software and firmware updates installed using {{ insert: param, si-02.02_odp.01 }} {{ insert: param, si-02.02_odp.02 }}.

Automated mechanisms can track and determine the status of known flaws for system components.

system components have applicable security-relevant software and firmware updates installed {{ insert: param, si-02.02_odp.02 }} using {{ insert: param, si-02.02_odp.01 }}.

System and information integrity policy

system and information integrity procedures

procedures addressing flaw remediation

automated mechanisms supporting centralized management of flaw remediation

system design documentation

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

organizational personnel installing, configuring, and/or maintaining the system

organizational personnel responsible for flaw remediation

Automated mechanisms used to determine the state of system components with regard to flaw remediation

### SI-3: Malicious Code Protection

Implement {{ insert: param, si-03_odp.01 }} malicious code protection mechanisms at system entry and exit points to detect and eradicate malicious code;

Automatically update malicious code protection mechanisms as new releases are available in accordance with organizational configuration management policy and procedures;

Configure malicious code protection mechanisms to:

Perform periodic scans of the system {{ insert: param, si-03_odp.02 }} and real-time scans of files from external sources at {{ insert: param, si-03_odp.03 }} as the files are downloaded, opened, or executed in accordance with organizational policy; and

 {{ insert: param, si-03_odp.04 }} ; and send alert to {{ insert: param, si-03_odp.06 }} in response to malicious code detection; and

Address the receipt of false positives during malicious code detection and eradication and the resulting potential impact on the availability of the system.

System entry and exit points include firewalls, remote access servers, workstations, electronic mail servers, web servers, proxy servers, notebook computers, and mobile devices. Malicious code includes viruses, worms, Trojan horses, and spyware. Malicious code can also be encoded in various formats contained within compressed or hidden files or hidden in files using techniques such as steganography. Malicious code can be inserted into systems in a variety of ways, including by electronic mail, the world-wide web, and portable storage devices. Malicious code insertions occur through the exploitation of system vulnerabilities. A variety of technologies and methods exist to limit or eliminate the effects of malicious code.

Malicious code protection mechanisms include both signature- and nonsignature-based technologies. Nonsignature-based detection mechanisms include artificial intelligence techniques that use heuristics to detect, analyze, and describe the characteristics or behavior of malicious code and to provide controls against such code for which signatures do not yet exist or for which existing signatures may not be effective. Malicious code for which active signatures do not yet exist or may be ineffective includes polymorphic malicious code (i.e., code that changes signatures when it replicates). Nonsignature-based mechanisms also include reputation-based technologies. In addition to the above technologies, pervasive configuration management, comprehensive software integrity controls, and anti-exploitation software may be effective in preventing the execution of unauthorized code. Malicious code may be present in commercial off-the-shelf software as well as custom-built software and could include logic bombs, backdoors, and other types of attacks that could affect organizational mission and business functions.

In situations where malicious code cannot be detected by detection methods or technologies, organizations rely on other types of controls, including secure coding practices, configuration management and control, trusted procurement processes, and monitoring practices to ensure that software does not perform functions other than the functions intended. Organizations may determine that, in response to the detection of malicious code, different actions may be warranted. For example, organizations can define actions in response to malicious code detection during periodic scans, the detection of malicious downloads, or the detection of maliciousness when attempting to open or execute files.

 {{ insert: param, si-03_odp.01 }} malicious code protection mechanisms are implemented at system entry and exit points to detect malicious code;

 {{ insert: param, si-03_odp.01 }} malicious code protection mechanisms are implemented at system entry and exit points to eradicate malicious code;

malicious code protection mechanisms are updated automatically as new releases are available in accordance with organizational configuration management policy and procedures;

malicious code protection mechanisms are configured to perform periodic scans of the system {{ insert: param, si-03_odp.02 }};

malicious code protection mechanisms are configured to perform real-time scans of files from external sources at {{ insert: param, si-03_odp.03 }} as the files are downloaded, opened, or executed in accordance with organizational policy;

malicious code protection mechanisms are configured to {{ insert: param, si-03_odp.04 }} in response to malicious code detection;

malicious code protection mechanisms are configured to send alerts to {{ insert: param, si-03_odp.06 }} in response to malicious code detection;

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

Attacks and indicators of potential attacks in accordance with the following monitoring objectives: {{ insert: param, si-04_odp.01 }} ; and

Unauthorized local, network, and remote connections;

Identify unauthorized use of the system through the following techniques and methods: {{ insert: param, si-04_odp.02 }};

Invoke internal monitoring capabilities or deploy monitoring devices:

Strategically within the system to collect organization-determined essential information; and

At ad hoc locations within the system to track specific types of transactions of interest to the organization;

Analyze detected events and anomalies;

Adjust the level of system monitoring activity when there is a change in risk to organizational operations and assets, individuals, other organizations, or the Nation;

Obtain legal opinion regarding system monitoring activities; and

Provide {{ insert: param, si-04_odp.03 }} to {{ insert: param, si-04_odp.04 }} {{ insert: param, si-04_odp.05 }}.

System monitoring includes external and internal monitoring. External monitoring includes the observation of events occurring at external interfaces to the system. Internal monitoring includes the observation of events occurring within the system. Organizations monitor systems by observing audit activities in real time or by observing other system aspects such as access patterns, characteristics of access, and other actions. The monitoring objectives guide and inform the determination of the events. System monitoring capabilities are achieved through a variety of tools and techniques, including intrusion detection and prevention systems, malicious code protection software, scanning tools, audit record monitoring software, and network monitoring software.

Depending on the security architecture, the distribution and configuration of monitoring devices may impact throughput at key internal and external boundaries as well as at other locations across a network due to the introduction of network throughput latency. If throughput management is needed, such devices are strategically located and deployed as part of an established organization-wide security architecture. Strategic locations for monitoring devices include selected perimeter locations and near key servers and server farms that support critical applications. Monitoring devices are typically employed at the managed interfaces associated with controls [SC-7](#sc-7) and [AC-17](#ac-17) . The information collected is a function of the organizational monitoring objectives and the capability of systems to support such objectives. Specific types of transactions of interest include Hypertext Transfer Protocol (HTTP) traffic that bypasses HTTP proxies. System monitoring is an integral part of organizational continuous monitoring and incident response programs, and output from system monitoring serves as input to those programs. System monitoring requirements, including the need for specific types of system monitoring, may be referenced in other controls (e.g., [AC-2g](#ac-2_smt.g), [AC-2(7)](#ac-2.7), [AC-2(12)(a)](#ac-2.12_smt.a), [AC-17(1)](#ac-17.1), [AU-13](#au-13), [AU-13(1)](#au-13.1), [AU-13(2)](#au-13.2), [CM-3f](#cm-3_smt.f), [CM-6d](#cm-6_smt.d), [MA-3a](#ma-3_smt.a), [MA-4a](#ma-4_smt.a), [SC-5(3)(b)](#sc-5.3_smt.b), [SC-7a](#sc-7_smt.a), [SC-7(24)(b)](#sc-7.24_smt.b), [SC-18b](#sc-18_smt.b), [SC-43b](#sc-43_smt.b) ). Adjustments to levels of system monitoring are based on law enforcement information, intelligence information, or other sources of information. The legality of system monitoring activities is based on applicable laws, executive orders, directives, regulations, policies, standards, and guidelines.

the system is monitored to detect attacks and indicators of potential attacks in accordance with {{ insert: param, si-04_odp.01 }};

the system is monitored to detect unauthorized local connections;

the system is monitored to detect unauthorized network connections;

the system is monitored to detect unauthorized remote connections;

unauthorized use of the system is identified through {{ insert: param, si-04_odp.02 }};

internal monitoring capabilities are invoked or monitoring devices are deployed strategically within the system to collect organization-determined essential information;

internal monitoring capabilities are invoked or monitoring devices are deployed at ad hoc locations within the system to track specific types of transactions of interest to the organization;

detected events are analyzed;

detected anomalies are analyzed;

the level of system monitoring activity is adjusted when there is a change in risk to organizational operations and assets, individuals, other organizations, or the Nation;

a legal opinion regarding system monitoring activities is obtained;

 {{ insert: param, si-04_odp.03 }} is provided to {{ insert: param, si-04_odp.04 }} {{ insert: param, si-04_odp.05 }}.

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

### SI-4 (2): Automated Tools and Mechanisms for Real-time Analysis

Employ automated tools and mechanisms to support near real-time analysis of events.

Automated tools and mechanisms include host-based, network-based, transport-based, or storage-based event monitoring tools and mechanisms or security information and event management (SIEM) technologies that provide real-time analysis of alerts and notifications generated by organizational systems. Automated monitoring techniques can create unintended privacy risks because automated controls may connect to external or otherwise unrelated systems. The matching of records between these systems may create linkages with unintended consequences. Organizations assess and document these risks in their privacy impact assessment and make determinations that are in alignment with their privacy program plan.

automated tools and mechanisms are employed to support a near real-time analysis of events.

System and information integrity policy

system and information integrity procedures

procedures addressing system monitoring tools and techniques

system design documentation

system monitoring tools and techniques documentation

system configuration settings and associated documentation

system audit records

system security plan

privacy plan

privacy program plan

privacy impact assessment

privacy risk management documentation

other relevant documents or records

System/network administrators

organizational personnel with information security and privacy responsibilities

organizational personnel installing, configuring, and/or maintaining the system

organizational personnel responsible for monitoring the system

organizational personnel responsible for incident response/management

Organizational processes for the near real-time analysis of events

organizational processes for system monitoring

mechanisms supporting and/or implementing system monitoring

mechanisms/tools supporting and/or implementing an analysis of events

### SI-4 (4): Inbound and Outbound Communications Traffic

Determine criteria for unusual or unauthorized activities or conditions for inbound and outbound communications traffic;

Monitor inbound and outbound communications traffic {{ insert: param, si-4.4_prm_1 }} for {{ insert: param, si-4.4_prm_2 }}.

Unusual or unauthorized activities or conditions related to system inbound and outbound communications traffic includes internal traffic that indicates the presence of malicious code or unauthorized use of legitimate code or credentials within organizational systems or propagating among system components, signaling to external systems, and the unauthorized exporting of information. Evidence of malicious code or unauthorized use of legitimate code or credentials is used to identify potentially compromised systems or system components.

criteria for unusual or unauthorized activities or conditions for inbound communications traffic are defined;

criteria for unusual or unauthorized activities or conditions for outbound communications traffic are defined;

inbound communications traffic is monitored {{ insert: param, si-04.04_odp.01 }} for {{ insert: param, si-04.04_odp.02 }};

outbound communications traffic is monitored {{ insert: param, si-04.04_odp.03 }} for {{ insert: param, si-04.04_odp.04 }}.

System and information integrity policy

system and information integrity procedures

procedures addressing system monitoring tools and techniques

system design documentation

system monitoring tools and techniques documentation

system configuration settings and associated documentation

system protocols

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

organizational personnel installing, configuring, and/or maintaining the system

organizational personnel responsible for monitoring the system

organizational personnel responsible for the intrusion detection system

Organizational processes for intrusion detection and system monitoring

mechanisms supporting and/or implementing intrusion detection and system monitoring capabilities

mechanisms supporting and/or implementing the monitoring of inbound and outbound communications traffic

### SI-4 (5): System-generated Alerts

Alert {{ insert: param, si-04.05_odp.01 }} when the following system-generated indications of compromise or potential compromise occur: {{ insert: param, si-04.05_odp.02 }}.

Alerts may be generated from a variety of sources, including audit records or inputs from malicious code protection mechanisms, intrusion detection or prevention mechanisms, or boundary protection devices such as firewalls, gateways, and routers. Alerts can be automated and may be transmitted telephonically, by electronic mail messages, or by text messaging. Organizational personnel on the alert notification list can include system administrators, mission or business owners, system owners, information owners/stewards, senior agency information security officers, senior agency officials for privacy, system security officers, or privacy officers. In contrast to alerts generated by the system, alerts generated by organizations in [SI-4(12)](#si-4.12) focus on information sources external to the system, such as suspicious activity reports and reports on potential insider threats.

 {{ insert: param, si-04.05_odp.01 }} are alerted when system-generated {{ insert: param, si-04.05_odp.02 }} occur.

System and information integrity policy

system and information integrity procedures

procedures addressing system monitoring tools and techniques

system monitoring tools and techniques documentation

system configuration settings and associated documentation

list of personnel selected to receive alerts

documentation of alerts generated based on compromise indicators

system audit records

system security plan

privacy plan

other relevant documents or records

System/network administrators

organizational personnel with information security and privacy responsibilities

system developers

organizational personnel installing, configuring, and/or maintaining the system

organizational personnel responsible for monitoring the system

organizational personnel on the system alert notification list

organizational personnel responsible for the intrusion detection system

Organizational processes for intrusion detection and system monitoring

mechanisms supporting and/or implementing intrusion detection and system monitoring capabilities

mechanisms supporting and/or implementing alerts for compromise indicators

### SI-4 (10): Visibility of Encrypted Communications

Make provisions so that {{ insert: param, si-04.10_odp.01 }} is visible to {{ insert: param, si-04.10_odp.02 }}.

Organizations balance the need to encrypt communications traffic to protect data confidentiality with the need to maintain visibility into such traffic from a monitoring perspective. Organizations determine whether the visibility requirement applies to internal encrypted traffic, encrypted traffic intended for external destinations, or a subset of the traffic types.

provisions are made so that {{ insert: param, si-04.10_odp.01 }} is visible to {{ insert: param, si-04.10_odp.02 }}.

System and information integrity policy

system and information integrity procedures

procedures addressing system monitoring tools and techniques

system design documentation

system monitoring tools and techniques documentation

system configuration settings and associated documentation

system protocols

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

organizational personnel installing, configuring, and/or maintaining the system

organizational personnel responsible for monitoring the system

organizational personnel responsible for the intrusion detection system

Organizational processes for intrusion detection and system monitoring

mechanisms supporting and/or implementing intrusion detection and system monitoring capabilities

mechanisms supporting and/or implementing the visibility of encrypted communications traffic to monitoring tools

### SI-4 (12): Automated Organization-generated Alerts

Alert {{ insert: param, si-04.12_odp.01 }} using {{ insert: param, si-04.12_odp.02 }} when the following indications of inappropriate or unusual activities with security or privacy implications occur: {{ insert: param, si-04.12_odp.03 }}.

Organizational personnel on the system alert notification list include system administrators, mission or business owners, system owners, senior agency information security officer, senior agency official for privacy, system security officers, or privacy officers. Automated organization-generated alerts are the security alerts generated by organizations and transmitted using automated means. The sources for organization-generated alerts are focused on other entities such as suspicious activity reports and reports on potential insider threats. In contrast to alerts generated by the organization, alerts generated by the system in [SI-4(5)](#si-4.5) focus on information sources that are internal to the systems, such as audit records.

 {{ insert: param, si-04.12_odp.01 }} is/are alerted using {{ insert: param, si-04.12_odp.02 }} when {{ insert: param, si-04.12_odp.03 }} indicate inappropriate or unusual activities with security or privacy implications.

System and information integrity policy

system and information integrity procedures

procedures addressing system monitoring tools and techniques

system design documentation

system monitoring tools and techniques documentation

system configuration settings and associated documentation

list of inappropriate or unusual activities with security and privacy implications that trigger alerts

suspicious activity reports

alerts provided to security and privacy personnel

system monitoring logs or records

system audit records

system security plan

privacy plan

other relevant documents or records

System/network administrators

organizational personnel with information security and privacy responsibilities

system developers

organizational personnel installing, configuring, and/or maintaining the system

organizational personnel responsible for monitoring the system

organizational personnel responsible for the intrusion detection system

Organizational processes for intrusion detection and system monitoring

automated mechanisms supporting and/or implementing intrusion detection and system monitoring capabilities

automated mechanisms supporting and/or implementing automated alerts to security personnel

### SI-4 (14): Wireless Intrusion Detection

Employ a wireless intrusion detection system to identify rogue wireless devices and to detect attack attempts and potential compromises or breaches to the system.

Wireless signals may radiate beyond organizational facilities. Organizations proactively search for unauthorized wireless connections, including the conduct of thorough scans for unauthorized wireless access points. Wireless scans are not limited to those areas within facilities containing systems but also include areas outside of facilities to verify that unauthorized wireless access points are not connected to organizational systems.

a wireless intrusion detection system is employed to identify rogue wireless devices;

a wireless intrusion detection system is employed to detect attack attempts on the system;

a wireless intrusion detection system is employed to detect potential compromises or breaches to the system.

System and information integrity policy

system and information integrity procedures

procedures addressing system monitoring tools and techniques

system design documentation

system monitoring tools and techniques documentation

system configuration settings and associated documentation

system protocols

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

organizational personnel installing, configuring, and/or maintaining the system

organizational personnel responsible for monitoring the system

organizational personnel responsible for the intrusion detection system

Organizational processes for intrusion detection

mechanisms supporting and/or implementing a wireless intrusion detection capability

### SI-4 (20): Privileged Users

Implement the following additional monitoring of privileged users: {{ insert: param, si-04.20_odp }}.

Privileged users have access to more sensitive information, including security-related information, than the general user population. Access to such information means that privileged users can potentially do greater damage to systems and organizations than non-privileged users. Therefore, implementing additional monitoring on privileged users helps to ensure that organizations can identify malicious activity at the earliest possible time and take appropriate actions.

 {{ insert: param, si-04.20_odp }} of privileged users is implemented.

System and information integrity policy

system and information integrity procedures

procedures addressing system monitoring tools and techniques

system design documentation

system monitoring tools and techniques documentation

system configuration settings and associated documentation

system monitoring logs or records

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

organizational personnel installing, configuring, and/or maintaining the system

organizational personnel responsible for monitoring the system

Organizational processes for system monitoring

mechanisms supporting and/or implementing a system monitoring capability

### SI-4 (22): Unauthorized Network Services

Detect network services that have not been authorized or approved by {{ insert: param, si-04.22_odp.01 }} ; and

 {{ insert: param, si-04.22_odp.02 }} when detected.

Unauthorized or unapproved network services include services in service-oriented architectures that lack organizational verification or validation and may therefore be unreliable or serve as malicious rogues for valid services.

network services that have not been authorized or approved by {{ insert: param, si-04.22_odp.01 }} are detected;

 {{ insert: param, si-04.22_odp.02 }} is/are initiated when network services that have not been authorized or approved by authorization or approval processes are detected.

System and information integrity policy

system and information integrity procedures

procedures addressing system monitoring tools and techniques

system design documentation

system monitoring tools and techniques documentation

system configuration settings and associated documentation

documented authorization/approval of network services

notifications or alerts of unauthorized network services

system monitoring logs or records

system audit records

system security plan

other relevant documents or records

System/network administrators

organizational personnel with information security responsibilities

system developer

organizational personnel installing, configuring, and/or maintaining the system

organizational personnel responsible for monitoring the system

Organizational processes for system monitoring

mechanisms supporting and/or implementing a system monitoring capability

mechanisms for auditing network services

mechanisms for providing alerts

### SI-5: Security Alerts, Advisories, and Directives

Receive system security alerts, advisories, and directives from {{ insert: param, si-05_odp.01 }} on an ongoing basis;

Generate internal security alerts, advisories, and directives as deemed necessary;

Disseminate security alerts, advisories, and directives to: {{ insert: param, si-05_odp.02 }} ; and

Implement security directives in accordance with established time frames, or notify the issuing organization of the degree of noncompliance.

The Cybersecurity and Infrastructure Security Agency (CISA) generates security alerts and advisories to maintain situational awareness throughout the Federal Government. Security directives are issued by OMB or other designated organizations with the responsibility and authority to issue such directives. Compliance with security directives is essential due to the critical nature of many of these directives and the potential (immediate) adverse effects on organizational operations and assets, individuals, other organizations, and the Nation should the directives not be implemented in a timely manner. External organizations include supply chain partners, external mission or business partners, external service providers, and other peer or supporting organizations.

system security alerts, advisories, and directives are received from {{ insert: param, si-05_odp.01 }} on an ongoing basis;

internal security alerts, advisories, and directives are generated as deemed necessary;

security alerts, advisories, and directives are disseminated to {{ insert: param, si-05_odp.02 }};

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

### SI-5 (1): Automated Alerts and Advisories

Broadcast security alert and advisory information throughout the organization using {{ insert: param, si-05.01_odp }}.

The significant number of changes to organizational systems and environments of operation requires the dissemination of security-related information to a variety of organizational entities that have a direct interest in the success of organizational mission and business functions. Based on information provided by security alerts and advisories, changes may be required at one or more of the three levels related to the management of risk, including the governance level, mission and business process level, and the information system level.

 {{ insert: param, si-05.01_odp }} are used to broadcast security alert and advisory information throughout the organization.

System and information integrity policy

system and information integrity procedures

procedures addressing security alerts, advisories, and directives

system design documentation

system configuration settings and associated documentation

automated mechanisms supporting the distribution of security alert and advisory information

records of security alerts and advisories

system audit records

system security plan

other relevant documents or records

Organizational personnel with security alert and advisory responsibilities

organizational personnel implementing, operating, maintaining, and using the system

organizational personnel, organizational elements, and/or external organizations to whom alerts and advisories are to be disseminated

system/network administrators

organizational personnel with information security responsibilities

Organizational processes for defining, receiving, generating, and disseminating security alerts and advisories

automated mechanisms supporting and/or implementing the dissemination of security alerts and advisories

### SI-6: Security and Privacy Function Verification

Verify the correct operation of {{ insert: param, si-6_prm_1 }};

Perform the verification of the functions specified in SI-6a {{ insert: param, si-06_odp.03 }};

Alert {{ insert: param, si-06_odp.06 }} to failed security and privacy verification tests; and

 {{ insert: param, si-06_odp.07 }} when anomalies are discovered.

Transitional states for systems include system startup, restart, shutdown, and abort. System notifications include hardware indicator lights, electronic alerts to system administrators, and messages to local computer consoles. In contrast to security function verification, privacy function verification ensures that privacy functions operate as expected and are approved by the senior agency official for privacy or that privacy attributes are applied or used as expected.

 {{ insert: param, si-06_odp.01 }} are verified to be operating correctly;

 {{ insert: param, si-06_odp.02 }} are verified to be operating correctly;

 {{ insert: param, si-06_odp.01 }} are verified {{ insert: param, si-06_odp.03 }};

 {{ insert: param, si-06_odp.02 }} are verified {{ insert: param, si-06_odp.03 }};

 {{ insert: param, si-06_odp.06 }} is/are alerted to failed security verification tests;

 {{ insert: param, si-06_odp.06 }} is/are alerted to failed privacy verification tests;

 {{ insert: param, si-06_odp.07 }} is/are initiated when anomalies are discovered.

System and information integrity policy

system and information integrity procedures

procedures addressing security and privacy function verification

system design documentation

system configuration settings and associated documentation

alerts/notifications of failed security verification tests

list of system transition states requiring security functionality verification

system audit records

system security plan

privacy plan

other relevant documents or records

Organizational personnel with security and privacy function verification responsibilities

organizational personnel implementing, operating, and maintaining the system

system/network administrators

organizational personnel with information security and privacy responsibilities

system developer

Organizational processes for security and privacy function verification

mechanisms supporting and/or implementing the security and privacy function verification capability

### SI-7: Software, Firmware, and Information Integrity

Employ integrity verification tools to detect unauthorized changes to the following software, firmware, and information: {{ insert: param, si-7_prm_1 }} ; and

Take the following actions when unauthorized changes to the software, firmware, and information are detected: {{ insert: param, si-7_prm_2 }}.

Unauthorized changes to software, firmware, and information can occur due to errors or malicious activity. Software includes operating systems (with key internal components, such as kernels or drivers), middleware, and applications. Firmware interfaces include Unified Extensible Firmware Interface (UEFI) and Basic Input/Output System (BIOS). Information includes personally identifiable information and metadata that contains security and privacy attributes associated with information. Integrity-checking mechanisms—including parity checks, cyclical redundancy checks, cryptographic hashes, and associated tools—can automatically monitor the integrity of systems and hosted applications.

integrity verification tools are employed to detect unauthorized changes to {{ insert: param, si-07_odp.01 }};

integrity verification tools are employed to detect unauthorized changes to {{ insert: param, si-07_odp.02 }};

integrity verification tools are employed to detect unauthorized changes to {{ insert: param, si-07_odp.03 }};

 {{ insert: param, si-07_odp.04 }} are taken when unauthorized changes to the software, are detected;

 {{ insert: param, si-07_odp.05 }} are taken when unauthorized changes to the firmware are detected;

 {{ insert: param, si-07_odp.06 }} are taken when unauthorized changes to the information are detected.

System and information integrity policy

system and information integrity procedures

procedures addressing software, firmware, and information integrity

personally identifiable information processing policy

system design documentation

system configuration settings and associated documentation

integrity verification tools and associated documentation

records generated or triggered by integrity verification tools regarding unauthorized software, firmware, and information changes

system audit records

system security plan

privacy plan

other relevant documents or records

Organizational personnel responsible for software, firmware, and/or information integrity

organizational personnel with information security and privacy responsibilities

system/network administrators

Software, firmware, and information integrity verification tools

### SI-7 (1): Integrity Checks

Perform an integrity check of {{ insert: param, si-7.1_prm_1 }} {{ insert: param, si-7.1_prm_2 }}.

Security-relevant events include the identification of new threats to which organizational systems are susceptible and the installation of new hardware, software, or firmware. Transitional states include system startup, restart, shutdown, and abort.

an integrity check of {{ insert: param, si-07.01_odp.01 }} is performed {{ insert: param, si-07.01_odp.02 }};

an integrity check of {{ insert: param, si-07.01_odp.05 }} is performed {{ insert: param, si-07.01_odp.06 }};

an integrity check of {{ insert: param, si-07.01_odp.09 }} is performed {{ insert: param, si-07.01_odp.10 }}.

System and information integrity policy

system and information integrity procedures

procedures addressing software, firmware, and information integrity testing

system design documentation

system configuration settings and associated documentation

integrity verification tools and associated documentation

records of integrity scans

system security plan

other relevant documents or records

Organizational personnel responsible for software, firmware, and/or information integrity

organizational personnel with information security responsibilities

system/network administrators

system developer

Software, firmware, and information integrity verification tools

### SI-7 (2): Automated Notifications of Integrity Violations

Employ automated tools that provide notification to {{ insert: param, si-07.02_odp }} upon discovering discrepancies during integrity verification.

The employment of automated tools to report system and information integrity violations and to notify organizational personnel in a timely matter is essential to effective risk response. Personnel with an interest in system and information integrity violations include mission and business owners, system owners, senior agency information security official, senior agency official for privacy, system administrators, software developers, systems integrators, information security officers, and privacy officers.

automated tools that provide notification to {{ insert: param, si-07.02_odp }} upon discovering discrepancies during integrity verification are employed.

System and information integrity policy

system and information integrity procedures

procedures addressing software, firmware, and information integrity

personally identifiable information processing policy

system design documentation

system configuration settings and associated documentation

integrity verification tools and associated documentation

records of integrity scans

automated tools supporting alerts and notifications for integrity discrepancies

notifications provided upon discovering discrepancies during integrity verifications

system audit records

system security plan

privacy plan

other relevant documents or records

Organizational personnel responsible for software, firmware, and/or information integrity

organizational personnel with information security and privacy responsibilities

system administrators

software developers

Software, firmware, and information integrity verification tools

mechanisms providing integrity discrepancy notifications

### SI-7 (5): Automated Response to Integrity Violations

Automatically {{ insert: param, si-07.05_odp.01 }} when integrity violations are discovered.

Organizations may define different integrity-checking responses by type of information, specific information, or a combination of both. Types of information include firmware, software, and user data. Specific information includes boot firmware for certain types of machines. The automatic implementation of controls within organizational systems includes reversing the changes, halting the system, or triggering audit alerts when unauthorized modifications to critical security files occur.

 {{ insert: param, si-07.05_odp.01 }} are automatically performed when integrity violations are discovered.

System and information integrity policy

system and information integrity procedures

procedures addressing software, firmware, and information integrity

system design documentation

system configuration settings and associated documentation

integrity verification tools and associated documentation

records of integrity scans

records of integrity checks and responses to integrity violations

audit records

system security plan

other relevant documents or records

Organizational personnel responsible for software, firmware, and/or information integrity

organizational personnel with information security responsibilities

system/network administrators

system developer

Software, firmware, and information integrity verification tools

mechanisms providing an automated response to integrity violations

mechanisms supporting and/or implementing security safeguards to be implemented when integrity violations are discovered

### SI-7 (7): Integration of Detection and Response

Incorporate the detection of the following unauthorized changes into the organizational incident response capability: {{ insert: param, si-07.07_odp }}.

Integrating detection and response helps to ensure that detected events are tracked, monitored, corrected, and available for historical purposes. Maintaining historical records is important for being able to identify and discern adversary actions over an extended time period and for possible legal actions. Security-relevant changes include unauthorized changes to established configuration settings or the unauthorized elevation of system privileges.

the detection of {{ insert: param, si-07.07_odp }} are incorporated into the organizational incident response capability.

System and information integrity policy

system and information integrity procedures

procedures addressing software, firmware, and information integrity

procedures addressing incident response

system design documentation

system configuration settings and associated documentation

incident response records

audit records

system security plan

other relevant documents or records

Organizational personnel responsible for software, firmware, and/or information integrity

organizational personnel with information security responsibilities

organizational personnel with incident response responsibilities

Organizational processes for incorporating the detection of unauthorized security-relevant changes into the incident response capability

software, firmware, and information integrity verification tools

mechanisms supporting and/or implementing the incorporation of detection of unauthorized security-relevant changes into the incident response capability

### SI-7 (15): Code Authentication

Implement cryptographic mechanisms to authenticate the following software or firmware components prior to installation: {{ insert: param, si-07.15_odp }}.

Cryptographic authentication includes verifying that software or firmware components have been digitally signed using certificates recognized and approved by organizations. Code signing is an effective method to protect against malicious code. Organizations that employ cryptographic mechanisms also consider cryptographic key management solutions.

cryptographic mechanisms are implemented to authenticate {{ insert: param, si-07.15_odp }} prior to installation.

System and information integrity policy

system and information integrity procedures

procedures addressing software, firmware, and information integrity

system design documentation

system configuration settings and associated documentation

cryptographic mechanisms and associated documentation

system audit records

system security plan

other relevant documents or records

Organizational personnel responsible for software, firmware, and/or information integrity

organizational personnel with information security responsibilities

system/network administrators

system developer

Cryptographic mechanisms authenticating software and firmware prior to installation

### SI-8: Spam Protection

Employ spam protection mechanisms at system entry and exit points to detect and act on unsolicited messages; and

Update spam protection mechanisms when new releases are available in accordance with organizational configuration management policy and procedures.

System entry and exit points include firewalls, remote-access servers, electronic mail servers, web servers, proxy servers, workstations, notebook computers, and mobile devices. Spam can be transported by different means, including email, email attachments, and web accesses. Spam protection mechanisms include signature definitions.

spam protection mechanisms are employed at system entry points to detect unsolicited messages;

spam protection mechanisms are employed at system exit points to detect unsolicited messages;

spam protection mechanisms are employed at system entry points to act on unsolicited messages;

spam protection mechanisms are employed at system exit points to act on unsolicited messages;

spam protection mechanisms are updated when new releases are available in accordance with organizational configuration management policies and procedures.

System and information integrity policy

system and information integrity procedures

configuration management policies and procedures (CM-01)

procedures addressing spam protection

spam protection mechanisms

records of spam protection updates

system design documentation

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

Organizational personnel responsible for spam protection

organizational personnel with information security responsibilities

system/network administrators

system developer

Organizational processes for implementing spam protection

mechanisms supporting and/or implementing spam protection

### SI-8 (2): Automatic Updates

Automatically update spam protection mechanisms {{ insert: param, si-08.02_odp }}.

Using automated mechanisms to update spam protection mechanisms helps to ensure that updates occur on a regular basis and provide the latest content and protection capabilities.

spam protection mechanisms are automatically updated {{ insert: param, si-08.02_odp }}.

System and information integrity policy

system and information integrity procedures

procedures addressing spam protection

spam protection mechanisms

records of spam protection updates

system design documentation

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

Organizational personnel responsible for spam protection

organizational personnel with information security responsibilities

system/network administrators

system developer

Organizational processes for spam protection

mechanisms supporting and/or implementing automatic updates to spam protection mechanisms

### SI-10: Information Input Validation

Check the validity of the following information inputs: {{ insert: param, si-10_odp }}.

Checking the valid syntax and semantics of system inputs—including character set, length, numerical range, and acceptable values—verifies that inputs match specified definitions for format and content. For example, if the organization specifies that numerical values between 1-100 are the only acceptable inputs for a field in a given application, inputs of "387," "abc," or "%K%" are invalid inputs and are not accepted as input to the system. Valid inputs are likely to vary from field to field within a software application. Applications typically follow well-defined protocols that use structured messages (i.e., commands or queries) to communicate between software modules or system components. Structured messages can contain raw or unstructured data interspersed with metadata or control information. If software applications use attacker-supplied inputs to construct structured messages without properly encoding such messages, then the attacker could insert malicious commands or special characters that can cause the data to be interpreted as control information or metadata. Consequently, the module or component that receives the corrupted output will perform the wrong operations or otherwise interpret the data incorrectly. Prescreening inputs prior to passing them to interpreters prevents the content from being unintentionally interpreted as commands. Input validation ensures accurate and correct inputs and prevents attacks such as cross-site scripting and a variety of injection attacks.

the validity of the {{ insert: param, si-10_odp }} is checked.

System and information integrity policy

system and information integrity procedures

access control policy and procedures

separation of duties policy and procedures

procedures addressing information input validation

documentation for automated tools and applications to verify the validity of information

list of information inputs requiring validity checks

system design documentation

system configuration settings and associated documentation

system audit records

system security plan

other relevant documents or records

Organizational personnel responsible for information input validation

organizational personnel with information security responsibilities

system/network administrators

system developer

Mechanisms supporting and/or implementing validity checks on information inputs

### SI-11: Error Handling

Generate error messages that provide information necessary for corrective actions without revealing information that could be exploited; and

Reveal error messages only to {{ insert: param, si-11_odp }}.

Organizations consider the structure and content of error messages. The extent to which systems can handle error conditions is guided and informed by organizational policy and operational requirements. Exploitable information includes stack traces and implementation details; erroneous logon attempts with passwords mistakenly entered as the username; mission or business information that can be derived from, if not stated explicitly by, the information recorded; and personally identifiable information, such as account numbers, social security numbers, and credit card numbers. Error messages may also provide a covert channel for transmitting information.

error messages that provide the information necessary for corrective actions are generated without revealing information that could be exploited;

error messages are revealed only to {{ insert: param, si-11_odp }}.

System and information integrity policy

system and information integrity procedures

procedures addressing system error handling

system design documentation

system configuration settings and associated documentation

documentation providing the structure and content of error messages

system audit records

system security plan

other relevant documents or records

Organizational personnel responsible for information input validation

organizational personnel with information security responsibilities

system/network administrators

system developer

Organizational processes for error handling

automated mechanisms supporting and/or implementing error handling

automated mechanisms supporting and/or implementing the management of error messages

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

### SI-16: Memory Protection

Implement the following controls to protect the system memory from unauthorized code execution: {{ insert: param, si-16_odp }}.

Some adversaries launch attacks with the intent of executing code in non-executable regions of memory or in memory locations that are prohibited. Controls employed to protect memory include data execution prevention and address space layout randomization. Data execution prevention controls can either be hardware-enforced or software-enforced with hardware enforcement providing the greater strength of mechanism.

 {{ insert: param, si-16_odp }} are implemented to protect the system memory from unauthorized code execution.

System and information integrity policy

system and information integrity procedures

procedures addressing memory protection for the system

system design documentation

system configuration settings and associated documentation

list of security safeguards protecting system memory from unauthorized code execution

system audit records

system security plan

other relevant documents or records

Organizational personnel responsible for memory protection

organizational personnel with information security responsibilities

system/network administrators

system developer

Automated mechanisms supporting and/or implementing safeguards to protect the system memory from unauthorized code execution

